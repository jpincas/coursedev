import { writable } from 'svelte/store';
import type { Locale } from '$lib/content/types.js';
import { setLocale as setLoaderLocale, getLocale } from '$lib/content/loader.js';

export type LocaleInfo = {
	code: Locale;
	emoji: string;
	label: string;
};

const LOCALE_INFO: Record<Locale, LocaleInfo> = {
	en: { code: 'en', emoji: '🇬🇧', label: 'English' },
	es: { code: 'es', emoji: '🇪🇸', label: 'Español' }
};

export function getLocaleInfo(locale: Locale): LocaleInfo {
	return LOCALE_INFO[locale];
}

// Plain writable store so $language works in templates
export const language = writable<Locale>(getLocale() || 'en');

export function setLanguage(locale: Locale) {
	setLoaderLocale(locale);
	language.set(locale);
}

export function toggleLanguage() {
	language.update((l) => {
		const next = l === 'en' ? 'es' : 'en';
		setLoaderLocale(next);
		return next;
	});
}
