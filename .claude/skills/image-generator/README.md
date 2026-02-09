# Image Generator Skill

This skill enables Claude Code to generate infographic-style images for training course modules using OpenAI's image generation API.

## Setup

### 1. Install Dependencies

```bash
pip install -r .claude/skills/image-generator/requirements.txt
```

Or install individually:
```bash
pip install openai requests pillow
```

### 2. Set API Key

Add your OpenAI API key to your environment:

```bash
export OPENAI_API_KEY="sk-..."
```

To make it permanent, add to your shell profile (`~/.zshrc`, `~/.bash_profile`, etc.):

```bash
echo 'export OPENAI_API_KEY="sk-..."' >> ~/.zshrc
source ~/.zshrc
```

### 3. Test the Skill

```bash
python .claude/skills/image-generator/scripts/generate_image.py \
  "Clean infographic on dark background showing three connected boxes" \
  --output test-image.png
```

## Usage with Claude Code

Once set up, Claude will automatically use this skill when you ask to generate images:

```
Generate an infographic for the context window module showing how different content fills the window
```

Claude will:
1. Understand the teaching concept
2. Craft a prompt matching the app's dark theme and infographic style
3. Generate the image
4. Save it to the module's `images/` directory
5. Provide the markdown reference to embed it

## Model Configuration

The script defaults to `gpt-image-1.5` which produces high-quality images. You can use alternative models if needed:

```bash
python .claude/skills/image-generator/scripts/generate_image.py \
  "Your prompt here" \
  --model dall-e-3 \
  --output image.png
```

**Available models**:
- `gpt-image-1.5` (default, high quality)
- `dall-e-3` (alternative OpenAI model)
- `dall-e-2` (older, faster, cheaper)

Ensure your API key has access to the model you choose.

## Command Line Options

```
usage: generate_image.py [-h] [--output OUTPUT] [--size SIZE] [--model MODEL]
                         [--quality QUALITY] [--json]
                         prompt

Arguments:
  prompt                Text description of the image to generate

Options:
  --output, -o          Path to save the image
  --size, -s            Image size: 1024x1024, 1024x1536, 1536x1024, auto
  --model, -m           Model to use (default: gpt-image-1.5)
  --quality, -q         Quality: low, medium, high, auto
  --json                Output result as JSON
```

## Examples

### Generate infographic for a module
```bash
python .claude/skills/image-generator/scripts/generate_image.py \
  "Clean infographic on dark background (#18181b) showing AI delegation: a person directing three AI agents, each handling a different task. Flat design, teal accent (#00d9c0), white text labels." \
  --output content/module-delegation/images/delegation-overview.png \
  --size 1536x1024 \
  --quality high
```

### Generate square image
```bash
python .claude/skills/image-generator/scripts/generate_image.py \
  "Infographic icon showing a context window on dark background" \
  --output content/module-context/images/context-icon.png \
  --size 1024x1024
```

## Integration with Content Workflow

This skill integrates with the module-builder subagent:

1. **module-builder** writes narrative markdown pages
2. **`/mermaid`** creates structured diagrams (flowcharts, sequences)
3. **`/excalidraw`** creates freeform sketches and drawings
4. **`/image-generator`** (this skill) creates infographic-style visuals
5. **`/quiz`** and **`/agent-demo`** add interactive elements

## Troubleshooting

### "OPENAI_API_KEY environment variable not set"
- Ensure you've exported the API key in your current shell session
- Check: `echo $OPENAI_API_KEY`

### "openai package not installed"
- Run: `pip install openai`

### Model not found error
- Verify your API key has access to `gpt-image-1.5`
- Try with `--model dall-e-3` as an alternative to test basic functionality

### Image quality issues
- Use `--quality high` for best results
- Ensure prompts are detailed and specific
- Include colour palette and style instructions in the prompt
- Specify the dark background explicitly (`#18181b`)
