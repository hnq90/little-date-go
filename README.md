# little-date-go

Small & sweet date-range formatting for Go.

This is a Go implementation of the original [little-date](https://github.com/vercel/little-date) library, which is written in TypeScript.

## Introduction

When displaying date-ranges in UI, they are often too long, repetitive, and hard to read. `little-date-go` solves this problem by making date ranges **short**, **readable**, and **easy to understand**.

Consider this example:

```go
from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
to := time.Date(2024, 1, 12, 23, 59, 59, 999999999, time.UTC)

// Typical long format
fmt.Printf("%s - %s", from.Format(time.RFC1123), to.Format(time.RFC1123))
// Output: "Mon, 01 Jan 2024 00:00:00 UTC - Fri, 12 Jan 2024 23:59:59 UTC"

// With little-date-go
options := littledate.DateRangeFormatOptions{Locale: "en_US"}
fmt.Println(littledate.FormatDateRange(from, to, options))
// Output: "Jan 1 - 12"
```

**Example formats ✨**

- `Jan 1 - 12`
- `Jan 3 - Apr 20`
- `January 2023`
- `Q1 2023`

Wasn't that easy to read? You can find a full list of formatting examples in the section below.

## Installation

```sh
go get github.com/hnq90/little-date-go
```

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/hnq90/little-date-go"
)

func main() {
    from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
    to := time.Date(2023, 1, 12, 23, 59, 59, 999999999, time.UTC)

    options := littledate.DateRangeFormatOptions{
        Locale: "en_US",
        IncludeTime: true,
    }

    result := littledate.FormatDateRange(from, to, options)
    fmt.Println(result) // Outputs: "Jan 1 - 12"
}
```

## Formatting Examples

| Description                               | Output                                   |
| ----------------------------------------- | ---------------------------------------- |
| Multiple days, same month                 | `Jan 1 - 12`                             |
| Multiple days, different months           | `Jan 3 - Apr 20`                         |
| Full day                                  | `Sun, Jan 1`                             |
| Range spanning different years            | `Jan 1 '22 - Jan 20 '23`                 |
| Multiple days, same month, past year      | `Jan 1 - 12, 2022`                       |
| Full day, past year                       | `Sat, Jan 1, 2022`                       |
| **Special cases**                         |                                          |
| Full year                                 | `2023`                                   |
| Quarter range                             | `Q1 2023`                                |
| Full month                                | `January 2023`                           |
| Full months                               | `Jan - Feb 2023`                         |
| **With time**                             |                                          |
| Today, different hours                    | `12pm - 1pm`                             |
| Same day, different hours                 | `Jan 1, 12:11am - 2:30pm`                |
| Same day, different hours, 24-hour format | `Jan 1, 0:11 - 14:00`                    |
| Hour difference within a day              | `Jan 1, 12pm - 12:59pm`                  |
| Different days with time included         | `Jan 1, 12:11am - Jan 2, 2:30pm`         |
| Different years with time                 | `Jan 1 '22, 12:11am - Jan 2 '23, 2:30pm` |
| Different years, no time                  | `Jan 1 '22 - Jan 2 '23`                  |

## Options

The formatting behavior can be customized with the following options:

```go
options := littledate.DateRangeFormatOptions{
    Today:      time.Now(),  // Override the default "today" date (useful for testing)
    Locale:     "en_US",     // The locale to use for formatting (e.g., "en_US", "en_GB")
    IncludeTime: true,       // Whether to include time in the formatted output
    Separator:   "-",        // The separator to use between the dates (e.g., "-", "to")
}

result := littledate.FormatDateRange(from, to, options)
```

## Localization Support

The library supports internationalization (i18n) and localization using the following approaches:

1. **External JSON Files**: Translations can be defined in JSON files located in the `i18n/locales` directory.
2. **Embedded Translations**: Translations can be embedded into the binary when building with `-tags=embed_resources`.
3. **Fallback Mechanism**: If a translation is not found, the library falls back to built-in translations.

Currently supported languages:
- English (`en`)
- French (`fr`)
- Spanish (`es`)
- German (`de`)
- Japanese (`ja`)
- Korean (`ko`)
- Chinese Simplified (`zh-CN`)
- Chinese Traditional (`zh-TW`)
- Vietnamese (`vi`)

Adding a new language is as simple as creating a new JSON file in the `i18n/locales` directory. See the [i18n README](i18n/README.md) for more details.

The library also intelligently determines the time format (12-hour vs 24-hour) based on the locale, following regional standards.

## Features

- Zero external dependencies
- Simple and straightforward API
- Automatically shortens date ranges in an intelligent way
- Properly handles different time formats based on locale
- Supports different display formats for various date range scenarios

## License

MIT