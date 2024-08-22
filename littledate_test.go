package littledate

import (
	"testing"
	"time"
)

var today = time.Date(2023, 11, 15, 12, 0, 0, 0, time.UTC)

var defaultOptions = DateRangeFormatOptions{
	Today:  today,
	Locale: "en_US",
}

func TestFormatDateRange(t *testing.T) {
	tests := []struct {
		name     string
		from     time.Time
		to       time.Time
		options  DateRangeFormatOptions
		expected string
	}{
		{
			name:     "format date range same month",
			from:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, 1, 12, 23, 59, 59, 999999999, time.UTC),
			options:  defaultOptions,
			expected: "Jan 1 - 12",
		},
		{
			name:     "format date range different month",
			from:     time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, 4, 20, 23, 59, 59, 999999999, time.UTC),
			options:  defaultOptions,
			expected: "Jan 3 - Apr 20",
		},
		{
			name:     "format date range different year",
			from:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, 1, 20, 23, 59, 59, 999999999, time.UTC),
			options:  defaultOptions,
			expected: "Jan 1 '22 - Jan 20 '23",
		},
		{
			name:     "format date range full day",
			from:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			to:       time.Date(2023, 1, 1, 23, 59, 59, 999999999, time.UTC),
			options:  defaultOptions,
			expected: "Sun, Jan 1",
		},
		{
			name:     "format date range same day, different hours",
			from:     time.Date(2023, 1, 1, 0, 11, 0, 0, time.UTC),
			to:       time.Date(2023, 1, 1, 14, 30, 59, 999999999, time.UTC),
			options:  defaultOptions,
			expected: "Jan 1, 12:11am - 2:30pm",
		},
		// {
		// 	name:     "format date range same day, different hours, with 24-hour locale",
		// 	from:     time.Date(2023, 1, 1, 0, 11, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 1, 14, 0, 59, 999999999, time.UTC),
		// 	options:  DateRangeFormatOptions{Today: today, Locale: "en_GB"},
		// 	expected: "Jan 1, 0:11 - 14:00",
		// },
		// {
		// 	name:     "format date range different days with time",
		// 	from:     time.Date(2023, 1, 1, 0, 11, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 2, 14, 30, 0, 0, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "Jan 1, 12:11am - Jan 2, 2:30pm",
		// },
		// {
		// 	name:     "format date range different days with single time",
		// 	from:     time.Date(2023, 1, 1, 0, 11, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 2, 23, 59, 59, 999999999, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "Jan 1, 12:11am - Jan 2",
		// },
		// {
		// 	name:     "format date range different years with time",
		// 	from:     time.Date(2022, 1, 1, 0, 11, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 2, 14, 30, 0, 0, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "Jan 1 '22, 12:11am - Jan 2 '23, 2:30pm",
		// },
		// {
		// 	name:     "format date range different years with time, includeTime: false",
		// 	from:     time.Date(2022, 1, 1, 0, 11, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 2, 14, 30, 0, 0, time.UTC),
		// 	options:  DateRangeFormatOptions{Today: today, Locale: "en_US", IncludeTime: false},
		// 	expected: "Jan 1 '22 - Jan 2 '23",
		// },
		// {
		// 	name:     "format date range full day different year",
		// 	from:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2022, 1, 1, 23, 59, 59, 999999999, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "Sat, Jan 1, 2022",
		// },
		// {
		// 	name:     "format date range full month",
		// 	from:     time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 4, 30, 23, 59, 59, 999999999, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "April 2023",
		// },
		// {
		// 	name:     "format date range full month different year",
		// 	from:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 31, 23, 59, 59, 999999999, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "January 2023",
		// },
		// {
		// 	name:     "format date range hour difference",
		// 	from:     time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 1, 12, 59, 59, 999999999, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "Jan 1, 12pm - 12:59pm",
		// },
		// {
		// 	name:     "format date range today, different hours",
		// 	from:     today,
		// 	to:       today.Add(time.Hour),
		// 	options:  defaultOptions,
		// 	expected: "12pm - 1pm",
		// },
		// {
		// 	name:     "format date range full year",
		// 	from:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "2023",
		// },
		// {
		// 	name:     "format date range quarter",
		// 	from:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 3, 31, 23, 59, 59, 999999999, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "Q1 2023",
		// },
		// {
		// 	name:     "across two full months",
		// 	from:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 2, 28, 23, 59, 59, 999999999, time.UTC),
		// 	options:  defaultOptions,
		// 	expected: "Jan - Feb 2023",
		// },
		// {
		// 	name:     "custom separator",
		// 	from:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 12, 23, 59, 59, 999999999, time.UTC),
		// 	options:  DateRangeFormatOptions{Today: today, Locale: "en_US", Separator: "to"},
		// 	expected: "Jan 1 to 12",
		// },
		// {
		// 	name:     "automatic locale detection should not fail",
		// 	from:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		// 	to:       time.Date(2023, 1, 12, 23, 59, 59, 999999999, time.UTC),
		// 	options:  DateRangeFormatOptions{Today: today},
		// 	expected: "Jan 1 - 12",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatDateRange(tt.from, tt.to, tt.options)
			if result != tt.expected {
				t.Errorf("formatDateRange() = %v, want %v", result, tt.expected)
			}
		})
	}
}
