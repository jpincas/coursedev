package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

// DB is the global database connection
var db *sql.DB

// InitDB initializes the database connection and creates tables
func InitDB(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	// Enable foreign keys
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}

	// Create tables
	if err := createTables(); err != nil {
		return err
	}

	return nil
}

// createTables creates all required database tables
func createTables() error {
	schema := `
	-- Cohorts define groups of students
	CREATE TABLE IF NOT EXISTS cohorts (
		id TEXT PRIMARY KEY,
		code TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		expires_at DATETIME NOT NULL,
		max_students INTEGER NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	-- Students are uniquely identified by cohort + email
	CREATE TABLE IF NOT EXISTS students (
		id TEXT PRIMARY KEY,
		cohort_id TEXT NOT NULL,
		email TEXT NOT NULL,
		name TEXT NOT NULL,
		current_module TEXT NOT NULL DEFAULT '',
		current_page INTEGER NOT NULL DEFAULT 0,
		font_size TEXT NOT NULL DEFAULT 'medium',
		theme TEXT NOT NULL DEFAULT 'light',
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(cohort_id, email),
		FOREIGN KEY (cohort_id) REFERENCES cohorts(id) ON DELETE CASCADE
	);

	-- Maps browser sessions to students
	CREATE TABLE IF NOT EXISTS sessions (
		session_id TEXT PRIMARY KEY,
		student_id TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
	);

	-- Progress tracking tables
	CREATE TABLE IF NOT EXISTS module_progress (
		student_id TEXT NOT NULL,
		module_id TEXT NOT NULL,
		started INTEGER NOT NULL DEFAULT 0,
		completed_at DATETIME,
		PRIMARY KEY (student_id, module_id),
		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS page_views (
		student_id TEXT NOT NULL,
		module_id TEXT NOT NULL,
		page_index INTEGER NOT NULL,
		viewed_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (student_id, module_id, page_index),
		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS quiz_scores (
		student_id TEXT NOT NULL,
		quiz_id TEXT NOT NULL,
		correct INTEGER NOT NULL DEFAULT 0,
		attempts INTEGER NOT NULL DEFAULT 1,
		answered_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (student_id, quiz_id),
		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
	);

	-- Indexes
	CREATE INDEX IF NOT EXISTS idx_cohorts_code ON cohorts(code);
	CREATE INDEX IF NOT EXISTS idx_students_cohort ON students(cohort_id);
	CREATE INDEX IF NOT EXISTS idx_sessions_student ON sessions(student_id);
	CREATE INDEX IF NOT EXISTS idx_module_progress_student ON module_progress(student_id);
	CREATE INDEX IF NOT EXISTS idx_page_views_student ON page_views(student_id);
	CREATE INDEX IF NOT EXISTS idx_quiz_scores_student ON quiz_scores(student_id);
	`

	_, err := db.Exec(schema)
	return err
}

// ============================================================================
// Cohort Types and Functions
// ============================================================================

// Cohort represents a training cohort
type Cohort struct {
	ID          uuid.UUID
	Code        string
	Name        string
	ExpiresAt   time.Time
	MaxStudents int
	CreatedAt   time.Time
}

// IsExpired returns true if the cohort has expired
func (c *Cohort) IsExpired() bool {
	return time.Now().After(c.ExpiresAt)
}

// GenerateCohortCode generates a unique cohort code in format xxx-xxxx-xxx
func GenerateCohortCode() (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	for i := range b {
		b[i] = letters[int(b[i])%len(letters)]
	}

	// Format as xxx-xxxx-xxx
	return fmt.Sprintf("%s-%s-%s", string(b[0:3]), string(b[3:7]), string(b[7:10])), nil
}

// CreateCohort creates a new cohort with a generated code
func CreateCohort(name string, expiresAt time.Time, maxStudents int) (*Cohort, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	// Generate unique code
	var code string
	var err error
	for attempts := 0; attempts < 10; attempts++ {
		code, err = GenerateCohortCode()
		if err != nil {
			return nil, err
		}

		// Check if code already exists
		var exists int
		err = db.QueryRow("SELECT 1 FROM cohorts WHERE code = ?", code).Scan(&exists)
		if err == sql.ErrNoRows {
			break // Code is unique
		}
		if err != nil {
			return nil, err
		}
		// Code exists, try again
	}

	cohort := &Cohort{
		ID:          uuid.New(),
		Code:        code,
		Name:        name,
		ExpiresAt:   expiresAt,
		MaxStudents: maxStudents,
		CreatedAt:   time.Now(),
	}

	_, err = db.Exec(`
		INSERT INTO cohorts (id, code, name, expires_at, max_students, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, cohort.ID.String(), cohort.Code, cohort.Name, cohort.ExpiresAt, cohort.MaxStudents, cohort.CreatedAt)

	if err != nil {
		return nil, err
	}

	return cohort, nil
}

// GetCohortByCode returns a cohort by its code, or nil if not found
func GetCohortByCode(code string) (*Cohort, error) {
	if db == nil {
		return nil, nil
	}

	var cohort Cohort
	var idStr string
	err := db.QueryRow(`
		SELECT id, code, name, expires_at, max_students, created_at
		FROM cohorts WHERE code = ?
	`, code).Scan(&idStr, &cohort.Code, &cohort.Name, &cohort.ExpiresAt, &cohort.MaxStudents, &cohort.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	cohort.ID, err = uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	return &cohort, nil
}

// GetCohortByID returns a cohort by its ID
func GetCohortByID(id uuid.UUID) (*Cohort, error) {
	if db == nil {
		return nil, nil
	}

	var cohort Cohort
	var idStr string
	err := db.QueryRow(`
		SELECT id, code, name, expires_at, max_students, created_at
		FROM cohorts WHERE id = ?
	`, id.String()).Scan(&idStr, &cohort.Code, &cohort.Name, &cohort.ExpiresAt, &cohort.MaxStudents, &cohort.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	cohort.ID, err = uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	return &cohort, nil
}

// GetAllCohorts returns all cohorts
func GetAllCohorts() ([]*Cohort, error) {
	if db == nil {
		return nil, nil
	}

	rows, err := db.Query(`
		SELECT id, code, name, expires_at, max_students, created_at
		FROM cohorts ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cohorts []*Cohort
	for rows.Next() {
		var cohort Cohort
		var idStr string
		if err := rows.Scan(&idStr, &cohort.Code, &cohort.Name, &cohort.ExpiresAt, &cohort.MaxStudents, &cohort.CreatedAt); err != nil {
			return nil, err
		}
		cohort.ID, err = uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}
		cohorts = append(cohorts, &cohort)
	}

	return cohorts, nil
}

// UpdateCohort updates a cohort's details
func UpdateCohort(id uuid.UUID, name string, expiresAt time.Time, maxStudents int) error {
	if db == nil {
		return nil
	}

	_, err := db.Exec(`
		UPDATE cohorts
		SET name = ?, expires_at = ?, max_students = ?
		WHERE id = ?
	`, name, expiresAt, maxStudents, id.String())

	return err
}

// DeleteCohort deletes a cohort (cascades to students and their progress)
func DeleteCohort(id uuid.UUID) error {
	if db == nil {
		return nil
	}

	_, err := db.Exec("DELETE FROM cohorts WHERE id = ?", id.String())
	return err
}

// GetCohortStudentCount returns the number of students in a cohort
func GetCohortStudentCount(cohortID uuid.UUID) (int, error) {
	if db == nil {
		return 0, nil
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM students WHERE cohort_id = ?", cohortID.String()).Scan(&count)
	return count, err
}

// ============================================================================
// Student Types and Functions
// ============================================================================

// Student represents a student in a cohort
type Student struct {
	ID            uuid.UUID
	CohortID      uuid.UUID
	Email         string
	Name          string
	CurrentModule string
	CurrentPage   int
	FontSize      string
	Theme         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// GetStudentByEmail returns a student by cohort and email, or nil if not found
func GetStudentByEmail(cohortID uuid.UUID, email string) (*Student, error) {
	if db == nil {
		return nil, nil
	}

	var student Student
	var idStr, cohortStr string
	err := db.QueryRow(`
		SELECT id, cohort_id, email, name, current_module, current_page, font_size, theme, created_at, updated_at
		FROM students WHERE cohort_id = ? AND email = ?
	`, cohortID.String(), email).Scan(
		&idStr, &cohortStr, &student.Email, &student.Name,
		&student.CurrentModule, &student.CurrentPage,
		&student.FontSize, &student.Theme,
		&student.CreatedAt, &student.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	student.ID, _ = uuid.Parse(idStr)
	student.CohortID, _ = uuid.Parse(cohortStr)
	return &student, nil
}

// GetStudentByID returns a student by ID
func GetStudentByID(id uuid.UUID) (*Student, error) {
	if db == nil {
		return nil, nil
	}

	var student Student
	var idStr, cohortStr string
	err := db.QueryRow(`
		SELECT id, cohort_id, email, name, current_module, current_page, font_size, theme, created_at, updated_at
		FROM students WHERE id = ?
	`, id.String()).Scan(
		&idStr, &cohortStr, &student.Email, &student.Name,
		&student.CurrentModule, &student.CurrentPage,
		&student.FontSize, &student.Theme,
		&student.CreatedAt, &student.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	student.ID, _ = uuid.Parse(idStr)
	student.CohortID, _ = uuid.Parse(cohortStr)
	return &student, nil
}

// CreateStudent creates a new student in a cohort
func CreateStudent(cohortID uuid.UUID, email, name string) (*Student, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	student := &Student{
		ID:            uuid.New(),
		CohortID:      cohortID,
		Email:         email,
		Name:          name,
		CurrentModule: "",
		CurrentPage:   0,
		FontSize:      "medium",
		Theme:         "light",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err := db.Exec(`
		INSERT INTO students (id, cohort_id, email, name, current_module, current_page, font_size, theme, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, student.ID.String(), student.CohortID.String(), student.Email, student.Name,
		student.CurrentModule, student.CurrentPage, student.FontSize, student.Theme,
		student.CreatedAt, student.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return student, nil
}

// ============================================================================
// Session Mapping Functions
// ============================================================================

// CreateSessionMapping creates a mapping from browser session to student
func CreateSessionMapping(sessionID, studentID uuid.UUID) error {
	if db == nil {
		return nil
	}

	_, err := db.Exec(`
		INSERT INTO sessions (session_id, student_id, created_at)
		VALUES (?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(session_id) DO UPDATE SET
			student_id = excluded.student_id,
			created_at = CURRENT_TIMESTAMP
	`, sessionID.String(), studentID.String())

	return err
}

// GetStudentBySessionID returns the student for a browser session, or nil if not found
func GetStudentBySessionID(sessionID uuid.UUID) (*Student, error) {
	if db == nil {
		return nil, nil
	}

	var studentIDStr string
	err := db.QueryRow("SELECT student_id FROM sessions WHERE session_id = ?", sessionID.String()).Scan(&studentIDStr)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	studentID, err := uuid.Parse(studentIDStr)
	if err != nil {
		return nil, err
	}

	return GetStudentByID(studentID)
}

// DeleteSessionMapping removes a session mapping (logout)
func DeleteSessionMapping(sessionID uuid.UUID) error {
	if db == nil {
		return nil
	}

	_, err := db.Exec("DELETE FROM sessions WHERE session_id = ?", sessionID.String())
	return err
}

// ============================================================================
// Student State Loading and Saving
// ============================================================================

// StudentDBState holds the full student state loaded from DB
type StudentDBState struct {
	Student    *Student
	Progress   map[string]*ModuleProgressDB
	QuizScores []QuizScoreDB
}

// ModuleProgressDB holds module progress from DB
type ModuleProgressDB struct {
	Started     bool
	PagesViewed map[int]bool
	QuizScores  map[string]QuizScore
	CompletedAt *time.Time
}

// QuizScoreDB holds quiz score from DB
type QuizScoreDB struct {
	QuizID   string
	Correct  bool
	Attempts int
}

// LoadStudentState loads full student state including progress
func LoadStudentState(studentID uuid.UUID) (*StudentDBState, error) {
	if db == nil {
		return nil, nil
	}

	student, err := GetStudentByID(studentID)
	if err != nil || student == nil {
		return nil, err
	}

	state := &StudentDBState{
		Student:  student,
		Progress: make(map[string]*ModuleProgressDB),
	}

	// Load module progress
	rows, err := db.Query(`
		SELECT module_id, started, completed_at
		FROM module_progress WHERE student_id = ?
	`, studentID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var moduleID string
		var started int
		var completedAt sql.NullTime

		if err := rows.Scan(&moduleID, &started, &completedAt); err != nil {
			return nil, err
		}

		prog := &ModuleProgressDB{
			Started:     started == 1,
			PagesViewed: make(map[int]bool),
			QuizScores:  make(map[string]QuizScore),
		}
		if completedAt.Valid {
			prog.CompletedAt = &completedAt.Time
		}
		state.Progress[moduleID] = prog
	}

	// Load page views
	rows, err = db.Query(`
		SELECT module_id, page_index
		FROM page_views WHERE student_id = ?
	`, studentID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var moduleID string
		var pageIndex int

		if err := rows.Scan(&moduleID, &pageIndex); err != nil {
			return nil, err
		}

		if state.Progress[moduleID] == nil {
			state.Progress[moduleID] = &ModuleProgressDB{
				PagesViewed: make(map[int]bool),
				QuizScores:  make(map[string]QuizScore),
			}
		}
		state.Progress[moduleID].PagesViewed[pageIndex] = true
	}

	// Load quiz scores
	rows, err = db.Query(`
		SELECT quiz_id, correct, attempts
		FROM quiz_scores WHERE student_id = ?
	`, studentID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var quizID string
		var correct int
		var attempts int

		if err := rows.Scan(&quizID, &correct, &attempts); err != nil {
			return nil, err
		}

		state.QuizScores = append(state.QuizScores, QuizScoreDB{
			QuizID:   quizID,
			Correct:  correct == 1,
			Attempts: attempts,
		})
	}

	return state, nil
}

// SaveStudentNavigation updates the student's current position
func SaveStudentNavigation(studentID uuid.UUID, currentModule string, currentPage int) error {
	if db == nil {
		return nil
	}

	_, err := db.Exec(`
		UPDATE students
		SET current_module = ?, current_page = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`, currentModule, currentPage, studentID.String())

	return err
}

// SavePageView records that a page was viewed
func SavePageView(studentID uuid.UUID, moduleID string, pageIndex int) error {
	if db == nil {
		return nil
	}

	// Ensure module progress exists
	_, err := db.Exec(`
		INSERT INTO module_progress (student_id, module_id, started)
		VALUES (?, ?, 1)
		ON CONFLICT(student_id, module_id) DO UPDATE SET started = 1
	`, studentID.String(), moduleID)
	if err != nil {
		return err
	}

	// Record page view
	_, err = db.Exec(`
		INSERT INTO page_views (student_id, module_id, page_index)
		VALUES (?, ?, ?)
		ON CONFLICT(student_id, module_id, page_index) DO NOTHING
	`, studentID.String(), moduleID, pageIndex)

	return err
}

// SaveQuizScore records a quiz attempt
func SaveQuizScore(studentID uuid.UUID, quizID string, correct bool, attempts int) error {
	if db == nil {
		return nil
	}

	correctInt := 0
	if correct {
		correctInt = 1
	}

	_, err := db.Exec(`
		INSERT INTO quiz_scores (student_id, quiz_id, correct, attempts)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(student_id, quiz_id) DO UPDATE SET
			correct = excluded.correct,
			attempts = excluded.attempts,
			answered_at = CURRENT_TIMESTAMP
	`, studentID.String(), quizID, correctInt, attempts)

	return err
}

// SaveModuleCompletion records module completion
func SaveModuleCompletion(studentID uuid.UUID, moduleID string) error {
	if db == nil {
		return nil
	}

	_, err := db.Exec(`
		UPDATE module_progress
		SET completed_at = CURRENT_TIMESTAMP
		WHERE student_id = ? AND module_id = ? AND completed_at IS NULL
	`, studentID.String(), moduleID)

	return err
}

// SaveStudentPreferences updates student preferences
func SaveStudentPreferences(studentID uuid.UUID, fontSize, theme string) error {
	if db == nil {
		return nil
	}

	_, err := db.Exec(`
		UPDATE students
		SET font_size = ?, theme = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`, fontSize, theme, studentID.String())

	return err
}

// ============================================================================
// Utility Functions
// ============================================================================

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// LogDBError logs database errors without crashing the app
func LogDBError(operation string, err error) {
	if err != nil {
		log.Printf("DB error during %s: %v", operation, err)
	}
}
