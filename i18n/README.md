# Internationalization (i18n) for little-date-go

This directory contains translation files for the little-date-go library.

## Structure

- `locales/`: Contains JSON translation files for each supported language
  - `en.json`: English translations
  - `fr.json`: French translations
  - `es.json`: Spanish translations
  - `de.json`: German translations

## Adding a New Locale

To add support for a new language:

1. Create a new JSON file in the `locales` directory with the language code as the filename (e.g., `it.json` for Italian)

2. Add translations for all month and weekday names in both long and short forms, following this format:

```json
{
  "month.long.1": {
    "description": "Full name of January",
    "other": "January"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "Jan"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "Sunday"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "Sun"
  }
}
```

3. Update the `supportedLocales` slice in `littedate.go` to include your new language code.

## JSON Format

Each translation file must follow the go-i18n format:

- The key is the message ID (e.g., "month.long.1")
- The "description" field provides context for translators
- The "other" field contains the actual translation

## Fallback Mechanism

If a translation file is not found or a specific translation is missing:

1. The library will first look for external JSON files in one of these paths:
   - `i18n/locales` (from the repo root)
   - `../i18n/locales` (from the example directory)
   - `../../../i18n/locales` (when used as a Go module)

2. If no file is found, it will use the built-in translations.

3. If a specific translation is missing, it will fall back to English.

## Building with Embedded Translations

For applications where you want to embed translations into the binary:

```bash
go build -tags=embed_resources
```

This will use the translations in `resources_embed.go` instead of loading external files. 