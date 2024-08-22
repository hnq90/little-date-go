package littledate

import (
	"fmt"
	"strings"
	"time"

	"github.com/goodsign/monday"
)

func shortenAmPm(text string) string {
	shortened := strings.ReplaceAll(strings.ReplaceAll(text, " AM", "am"), " PM", "pm")
	if strings.Contains(shortened, "m") {
		return strings.ReplaceAll(shortened, ":00", "")
	}
	return shortened
}

func removeLeadingZero(text string) string {
	return strings.TrimPrefix(text, "0")
}

func formatTime(date time.Time, locale string) string {
	loc := monday.Locale(locale)
	format := "3:04 PM"
	if is24Hour(locale) {
		format = "15:04"
	}
	return removeLeadingZero(shortenAmPm(monday.Format(date, format, loc)))
}

func createFormatTime(locale string) func(time.Time) string {
	return func(date time.Time) string {
		return formatTime(date, locale)
	}
}

type DateRangeFormatOptions struct {
	Today       time.Time
	Locale      string
	IncludeTime bool
	Separator   string
}

func is24Hour(locale string) bool {
	loc := monday.Locale(locale)
	testTime := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	formattedTime := monday.Format(testTime, "3:04 PM", loc)
	return !strings.Contains(formattedTime, "PM") && !strings.Contains(formattedTime, "AM")
}

func formatDateRange(from, to time.Time, options DateRangeFormatOptions) string {
	if options.Today.IsZero() {
		options.Today = time.Now()
	}
	if options.Separator == "" {
		options.Separator = "-"
	}

	loc := monday.Locale(options.Locale)

	sameYear := from.Year() == to.Year()
	sameMonth := from.Month() == to.Month() && sameYear
	sameDay := from.Day() == to.Day() && sameMonth
	thisYear := from.Year() == options.Today.Year()
	thisDay := from.YearDay() == options.Today.YearDay() && thisYear

	var yearSuffix string
	if !thisYear {
		yearSuffix = fmt.Sprintf(", %d", to.Year())
	}

	formatTime := createFormatTime(options.Locale)

	var startTimeSuffix, endTimeSuffix string
	if options.IncludeTime {
		if !from.Truncate(24 * time.Hour).Equal(from) {
			startTimeSuffix = ", " + formatTime(from)
		}
		if !to.Truncate(24 * time.Hour).Add(23*time.Hour + 59*time.Minute + 59*time.Second).Equal(to) {
			endTimeSuffix = ", " + formatTime(to)
		}
	}

	if from.Equal(time.Date(from.Year(), 1, 1, 0, 0, 0, 0, from.Location())) &&
		to.Equal(time.Date(to.Year(), 12, 31, 23, 59, 59, 999999999, to.Location())) {
		return fmt.Sprintf("%d", from.Year())
	}

	if from.Equal(startOfQuarter(from)) && to.Equal(endOfQuarter(to)) && getQuarter(from) == getQuarter(to) {
		return fmt.Sprintf("Q%d %d", getQuarter(from), from.Year())
	}

	if from.Equal(time.Date(from.Year(), from.Month(), 1, 0, 0, 0, 0, from.Location())) &&
		to.Equal(time.Date(to.Year(), to.Month()+1, 0, 23, 59, 59, 999999999, to.Location())) {
		if sameMonth && sameYear {
			return monday.Format(from, "January 2006", loc)
		}
		return fmt.Sprintf("%s %s %s",
			monday.Format(from, "Jan", loc),
			options.Separator,
			monday.Format(to, "Jan 2006", loc))
	}

	if from.Year() == to.Year() && from.Month() == to.Month() && from.Day() == to.Day() {
		// Same day, different hours
		layout := "Jan 2, 3:04pm"
		if is24Hour(options.Locale) {
			layout = "Jan 2, 15:04"
		}
		fromStr := monday.Format(from, layout, loc)
		toStr := monday.Format(to, layout, loc)
		return shortenAmPm(removeLeadingZero(fromStr + " - " + toStr))
	}

	if !sameYear {
		return fmt.Sprintf("%s%s %s %s%s",
			monday.Format(from, "Jan 2 '06", loc), startTimeSuffix,
			options.Separator,
			monday.Format(to, "Jan 2 '06", loc), endTimeSuffix)
	}

	if !sameMonth {
		return fmt.Sprintf("%s%s %s %s%s%s",
			monday.Format(from, "Jan 2", loc), startTimeSuffix,
			options.Separator,
			monday.Format(to, "Jan 2", loc), endTimeSuffix, yearSuffix)
	}

	if !sameDay {
		if startTimeSuffix != "" || endTimeSuffix != "" {
			return fmt.Sprintf("%s%s %s %s%s%s",
				monday.Format(from, "Jan 2", loc), startTimeSuffix,
				options.Separator,
				monday.Format(to, "Jan 2", loc), endTimeSuffix, yearSuffix)
		}
		return fmt.Sprintf("%s %s %s%s",
			monday.Format(from, "Jan 2", loc),
			options.Separator,
			monday.Format(to, "2", loc), yearSuffix)
	}

	if startTimeSuffix != "" || endTimeSuffix != "" {
		if thisDay {
			return fmt.Sprintf("%s %s %s", formatTime(from), options.Separator, formatTime(to))
		}
		return fmt.Sprintf("%s%s %s %s%s",
			monday.Format(from, "Jan 2", loc), startTimeSuffix,
			options.Separator, formatTime(to), yearSuffix)
	}

	return fmt.Sprintf("%s%s", monday.Format(from, "Mon, Jan 2", loc), yearSuffix)
}

func startOfQuarter(t time.Time) time.Time {
	month := t.Month()
	quarter := (month-1)/3 + 1
	return time.Date(t.Year(), time.Month((quarter-1)*3+1), 1, 0, 0, 0, 0, t.Location())
}

func endOfQuarter(t time.Time) time.Time {
	month := t.Month()
	quarter := (month-1)/3 + 1
	return time.Date(t.Year(), time.Month(quarter*3+1), 0, 23, 59, 59, 999999999, t.Location())
}

func getQuarter(t time.Time) int {
	return int((t.Month()-1)/3 + 1)
}
