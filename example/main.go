package main

import (
	"fmt"
	"time"

	littledate "github.com/hnq90/little-date-go"
)

func main() {
	fmt.Println("Little Date Go Examples")
	fmt.Println("======================")

	// Example 1: Basic usage with default options
	fmt.Println("\n1. Basic usage (default options):")
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 3), littledate.DateRangeFormatOptions{}))

	// Example 2: Same day
	fmt.Println("\n2. Same day:")
	today := time.Now()
	fmt.Println(littledate.FormatDateRange(today, today, littledate.DateRangeFormatOptions{}))

	// Example 3: Same month
	fmt.Println("\n3. Same month:")
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 14), littledate.DateRangeFormatOptions{}))

	// Example 4: Same year
	fmt.Println("\n4. Same year:")
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 2, 0), littledate.DateRangeFormatOptions{}))

	// Example 5: Different years
	fmt.Println("\n5. Different years:")
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(1, 2, 0), littledate.DateRangeFormatOptions{}))

	// Example 6: Using custom options
	fmt.Println("\n6. Custom options:")
	options := littledate.DateRangeFormatOptions{
		Today:       time.Now(),
		Locale:      "en_US",
		IncludeTime: true,
		Separator:   " → ",
	}
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 3), options))

	// Example 7: French locale (using simple language code)
	fmt.Println("\n7. French locale:")
	frOptions := littledate.DateRangeFormatOptions{
		Locale: "fr",
	}
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 3), frOptions))

	// Example 8: Spanish locale (using simple language code)
	fmt.Println("\n8. Spanish locale:")
	esOptions := littledate.DateRangeFormatOptions{
		Locale: "es",
	}
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 3), esOptions))

	// Example 9: Include time
	fmt.Println("\n9. Including time:")
	timeOptions := littledate.DateRangeFormatOptions{
		IncludeTime: true,
	}
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 1), timeOptions))

	// Example 10: Custom today string
	fmt.Println("\n10. Custom today reference date:")
	customToday := time.Now().AddDate(0, 0, -1) // yesterday
	todayOptions := littledate.DateRangeFormatOptions{
		Today: customToday,
	}
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 1), todayOptions))

	// Example 11: Custom separator
	fmt.Println("\n11. Custom separator:")
	sepOptions := littledate.DateRangeFormatOptions{
		Separator: " to ",
	}
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 3), sepOptions))

	// Example 12: Future date to today
	fmt.Println("\n12. Future date to today (should swap automatically):")
	future := time.Now().AddDate(0, 0, 5)
	fmt.Println(littledate.FormatDateRange(future, time.Now(), littledate.DateRangeFormatOptions{}))

	// Example 13: Full options customization with Japanese locale
	fmt.Println("\n13. Full options customization (Japanese):")
	fullOptions := littledate.DateRangeFormatOptions{
		Today:       time.Now(),
		Locale:      "ja",
		IncludeTime: true,
		Separator:   " から ",
	}
	fmt.Println(littledate.FormatDateRange(time.Now(), time.Now().AddDate(0, 0, 10), fullOptions))
}
