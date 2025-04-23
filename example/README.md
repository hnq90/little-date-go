# Little Date Go Examples

This directory contains examples demonstrating how to use the Little Date Go library.

## Running the Examples

You can run the examples with the following command from the root of the repository:

```bash
go run example/main.go
```

Or from within the `example` directory:

```bash
go run main.go
```

## Example Overview

The `main.go` file demonstrates various use cases of the Little Date Go library, including:

1. Formatting date ranges within the same month
2. Formatting date ranges spanning different months
3. Formatting date ranges spanning different years
4. Formatting full day displays
5. Formatting date ranges with time on the same day
6. Formatting full month ranges
7. Formatting full year ranges
8. Formatting quarters
9. Formatting today with time
10. Using different locales (en_US and en_GB)
11. Using custom separators
12. Formatting without time information
13. Individual time formatting

Each example shows the input dates and the resulting formatted string, making it easy to understand how the library formats different kinds of date ranges.

## Customization

The examples demonstrate the use of `DateRangeFormatOptions` to customize the output. You can modify these options to see how they affect the formatted result:

- `Today`: Reference date for determining relative dates
- `Locale`: Language and regional formatting (currently supports "en_US" and "en_GB")
- `IncludeTime`: Whether to include time information in the output
- `Separator`: String used to separate date ranges 