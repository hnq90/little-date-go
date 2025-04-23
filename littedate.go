package littledate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// i18n bundle for localization
var bundle *i18n.Bundle

// Supported languages with their translation files
var supportedLocales = []string{"en", "fr", "es", "de", "ja", "ko", "zh-CN", "zh-TW", "vi"}

// Initialize the i18n bundle
func init() {
	// Initialize a new bundle with English as the default language
	bundle = i18n.NewBundle(language.English)

	// Set the JSON unmarshal function
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load all translation files
	loadTranslationFiles()
}

// loadTranslationFiles loads all JSON translation files from the i18n/locales directory
func loadTranslationFiles() {
	// Try different paths to find the translation files
	// This allows the library to work both when used as a dependency and when run from the repo root
	possiblePaths := []string{
		"i18n/locales",          // From repo root
		"../i18n/locales",       // From example directory
		"../../../i18n/locales", // When used as a Go module
	}

	// Determine the path to the translations
	var basePath string
	for _, path := range possiblePaths {
		if dirExists(path) {
			basePath = path
			break
		}
	}

	// If no translation directory was found, use embedded translations instead
	if basePath == "" {
		// Fallback to hard-coded translations
		addHardCodedTranslations()
		return
	}

	// Load translation files for all supported locales
	for _, locale := range supportedLocales {
		filePath := filepath.Join(basePath, locale+".json")
		if fileExists(filePath) {
			// Load the translation file
			bundle.MustLoadMessageFile(filePath)
		} else {
			// File not found, use embedded translations for this locale
			addHardCodedTranslationsForLocale(locale)
		}
	}
}

// dirExists checks if a directory exists
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// fileExists checks if a file exists
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// addHardCodedTranslations adds built-in translations as a fallback
func addHardCodedTranslations() {
	// Add translations for all supported locales
	for _, locale := range supportedLocales {
		addHardCodedTranslationsForLocale(locale)
	}
}

// addHardCodedTranslationsForLocale adds built-in translations for a specific locale
func addHardCodedTranslationsForLocale(locale string) {
	switch locale {
	case "en":
		// English translations
		addMonthTranslations("en", []string{
			"January", "February", "March", "April", "May", "June",
			"July", "August", "September", "October", "November", "December",
		}, []string{
			"Jan", "Feb", "Mar", "Apr", "May", "Jun",
			"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
		})

		addWeekdayTranslations("en", []string{
			"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
		}, []string{
			"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat",
		})

	case "fr":
		// French translations
		addMonthTranslations("fr", []string{
			"Janvier", "Février", "Mars", "Avril", "Mai", "Juin",
			"Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre",
		}, []string{
			"Jan", "Fév", "Mar", "Avr", "Mai", "Jun",
			"Jul", "Aoû", "Sep", "Oct", "Nov", "Déc",
		})

		addWeekdayTranslations("fr", []string{
			"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi",
		}, []string{
			"Dim", "Lun", "Mar", "Mer", "Jeu", "Ven", "Sam",
		})

	case "es":
		// Spanish translations
		addMonthTranslations("es", []string{
			"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
			"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre",
		}, []string{
			"Ene", "Feb", "Mar", "Abr", "May", "Jun",
			"Jul", "Ago", "Sep", "Oct", "Nov", "Dic",
		})

		addWeekdayTranslations("es", []string{
			"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado",
		}, []string{
			"Dom", "Lun", "Mar", "Mié", "Jue", "Vie", "Sáb",
		})

	case "de":
		// German translations
		addMonthTranslations("de", []string{
			"Januar", "Februar", "März", "April", "Mai", "Juni",
			"Juli", "August", "September", "Oktober", "November", "Dezember",
		}, []string{
			"Jan", "Feb", "Mär", "Apr", "Mai", "Jun",
			"Jul", "Aug", "Sep", "Okt", "Nov", "Dez",
		})

		addWeekdayTranslations("de", []string{
			"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag",
		}, []string{
			"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa",
		})

	case "ja":
		// Japanese translations
		addMonthTranslations("ja", []string{
			"1月", "2月", "3月", "4月", "5月", "6月",
			"7月", "8月", "9月", "10月", "11月", "12月",
		}, []string{
			"1月", "2月", "3月", "4月", "5月", "6月",
			"7月", "8月", "9月", "10月", "11月", "12月",
		})

		addWeekdayTranslations("ja", []string{
			"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日",
		}, []string{
			"日", "月", "火", "水", "木", "金", "土",
		})

	case "ko":
		// Korean translations
		addMonthTranslations("ko", []string{
			"1월", "2월", "3월", "4월", "5월", "6월",
			"7월", "8월", "9월", "10월", "11월", "12월",
		}, []string{
			"1월", "2월", "3월", "4월", "5월", "6월",
			"7월", "8월", "9월", "10월", "11월", "12월",
		})

		addWeekdayTranslations("ko", []string{
			"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일",
		}, []string{
			"일", "월", "화", "수", "목", "금", "토",
		})

	case "zh-CN", "zh":
		// Chinese Simplified translations
		addMonthTranslations("zh-CN", []string{
			"一月", "二月", "三月", "四月", "五月", "六月",
			"七月", "八月", "九月", "十月", "十一月", "十二月",
		}, []string{
			"1月", "2月", "3月", "4月", "5月", "6月",
			"7月", "8月", "9月", "10月", "11月", "12月",
		})

		addWeekdayTranslations("zh-CN", []string{
			"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六",
		}, []string{
			"日", "一", "二", "三", "四", "五", "六",
		})

	case "zh-TW":
		// Chinese Traditional translations
		addMonthTranslations("zh-TW", []string{
			"一月", "二月", "三月", "四月", "五月", "六月",
			"七月", "八月", "九月", "十月", "十一月", "十二月",
		}, []string{
			"1月", "2月", "3月", "4月", "5月", "6月",
			"7月", "8月", "9月", "10月", "11月", "12月",
		})

		addWeekdayTranslations("zh-TW", []string{
			"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六",
		}, []string{
			"日", "一", "二", "三", "四", "五", "六",
		})

	case "vi":
		// Vietnamese translations
		addMonthTranslations("vi", []string{
			"Tháng Một", "Tháng Hai", "Tháng Ba", "Tháng Tư", "Tháng Năm", "Tháng Sáu",
			"Tháng Bảy", "Tháng Tám", "Tháng Chín", "Tháng Mười", "Tháng Mười Một", "Tháng Mười Hai",
		}, []string{
			"Th1", "Th2", "Th3", "Th4", "Th5", "Th6",
			"Th7", "Th8", "Th9", "Th10", "Th11", "Th12",
		})

		addWeekdayTranslations("vi", []string{
			"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy",
		}, []string{
			"CN", "T2", "T3", "T4", "T5", "T6", "T7",
		})
	}
}

// Helper function to add month translations to the bundle
func addMonthTranslations(lang string, longNames, shortNames []string) {
	for i, name := range longNames {
		bundle.AddMessages(language.MustParse(lang), &i18n.Message{
			ID:    fmt.Sprintf("month.long.%d", i+1),
			Other: name,
		})
	}

	for i, name := range shortNames {
		bundle.AddMessages(language.MustParse(lang), &i18n.Message{
			ID:    fmt.Sprintf("month.short.%d", i+1),
			Other: name,
		})
	}
}

// Helper function to add weekday translations to the bundle
func addWeekdayTranslations(lang string, longNames, shortNames []string) {
	for i, name := range longNames {
		bundle.AddMessages(language.MustParse(lang), &i18n.Message{
			ID:    fmt.Sprintf("weekday.long.%d", i),
			Other: name,
		})
	}

	for i, name := range shortNames {
		bundle.AddMessages(language.MustParse(lang), &i18n.Message{
			ID:    fmt.Sprintf("weekday.short.%d", i),
			Other: name,
		})
	}
}

// Get a localizer for the specified locale
func getLocalizer(locale string) *i18n.Localizer {
	// Convert locale format from "en_US" to "en-US" if needed
	var tag language.Tag
	localeParts := strings.Split(locale, "_")
	if len(localeParts) == 2 {
		tag = language.MustParse(localeParts[0] + "-" + localeParts[1])
	} else if locale != "" {
		tag = language.MustParse(locale)
	} else {
		tag = language.English
	}

	return i18n.NewLocalizer(bundle, tag.String())
}

// Get localized month name (short or long)
func getMonthName(month time.Month, locale string, short bool) string {
	localizer := getLocalizer(locale)

	var format string
	if short {
		format = "month.short." + strconv.Itoa(int(month))
	} else {
		format = "month.long." + strconv.Itoa(int(month))
	}

	localized, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: format,
	})

	if err != nil {
		// Fallback to English format if translation not found
		if short {
			return time.Month(month).String()[:3]
		}
		return time.Month(month).String()
	}

	return localized
}

// Get localized weekday name (short or long)
func getWeekdayName(weekday time.Weekday, locale string, short bool) string {
	localizer := getLocalizer(locale)

	var format string
	if short {
		format = "weekday.short." + strconv.Itoa(int(weekday))
	} else {
		format = "weekday.long." + strconv.Itoa(int(weekday))
	}

	localized, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: format,
	})

	if err != nil {
		// Fallback to English format if translation not found
		if short {
			return time.Weekday(weekday).String()[:3]
		}
		return time.Weekday(weekday).String()
	}

	return localized
}

// Format utility functions
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

// Determine if the locale typically uses 24-hour time format
// This is based on common international standards
func is24Hour(locale string) bool {
	// Extract language and country if available
	var language, country string

	parts := strings.Split(locale, "_")
	if len(parts) >= 1 {
		language = strings.ToLower(parts[0])
	}
	if len(parts) >= 2 {
		country = strings.ToUpper(parts[1])
	}

	// United States is known to use 12-hour format
	if country == "US" {
		return false
	}

	// English speakers in countries other than the US typically use 12-hour format
	// (with some exceptions like the UK)
	if language == "en" && country != "GB" {
		return false
	}

	// Most countries use 24-hour format
	// Specific exceptions for countries/languages that commonly use 12-hour format
	switch language {
	case "en":
		// English-speaking countries vary, but we already handled the main cases above
		return country == "GB" // UK uses 24-hour format
	case "es":
		// Some Spanish-speaking regions use 12-hour
		if country == "MX" || country == "CO" || country == "AR" || country == "CL" {
			return false
		}
	case "zh", "zh-cn", "zh-tw":
		// Chinese typically uses 12-hour in informal contexts
		return false
	case "ja":
		// Japanese typically uses 12-hour format in casual contexts
		return false
	case "ko":
		// Korean typically uses 12-hour format in casual contexts
		return false
	case "vi":
		// Vietnamese typically uses 12-hour format in casual contexts
		return false
	case "fil", "tl":
		// Filipino typically uses 12-hour
		return false
	case "hi", "ur", "bn", "pa", "gu", "mr", "ta", "te", "kn", "ml":
		// Languages of the Indian subcontinent often use 12-hour
		return false
	}

	// Default to 24-hour format for most other locales
	return true
}

func FormatTime(date time.Time, locale string) string {
	if is24Hour(locale) {
		hour := date.Hour()
		minute := date.Minute()
		formatted := fmt.Sprintf("%02d:%02d", hour, minute)

		// Remove leading zero from hours for 24-hour format
		if hour < 10 {
			return formatted[1:]
		}
		return formatted
	} else {
		hour := date.Hour()
		minute := date.Minute()
		period := "AM"
		if hour >= 12 {
			period = "PM"
			if hour > 12 {
				hour -= 12
			}
		}
		if hour == 0 {
			hour = 12
		}
		formatted := fmt.Sprintf("%d:%02d %s", hour, minute, period)
		return removeLeadingZero(shortenAmPm(formatted))
	}
}

func createFormatTime(locale string) func(time.Time) string {
	return func(date time.Time) string {
		return FormatTime(date, locale)
	}
}

// DateRangeFormatOptions specifies configuration options for formatting a date range.
type DateRangeFormatOptions struct {
	// Today is the reference date for determining relative dates.
	// If not specified, time.Now() will be used.
	Today time.Time

	// Locale determines the language and regional formatting to use.
	// Supported values include "en_US", "en_GB", "fr", "es", "de", etc.
	// If not specified, "en_US" will be used.
	Locale string

	// IncludeTime determines whether to include time information in the output.
	// Default is false.
	IncludeTime bool

	// Separator is the string used to separate date ranges.
	// If not specified, "-" will be used.
	Separator string
}

// Helper time functions
func isSameMinute(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day() &&
		t1.Hour() == t2.Hour() &&
		t1.Minute() == t2.Minute()
}

func startOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func endOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func startOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func endOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, t.Location())
}

func startOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func endOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
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

// FormatDateRange formats a date range in a human-readable way.
// It intelligently shortens the format based on the range.
//
// Examples:
// - Jan 1 - 12
// - Jan 3 - Apr 20
// - January 2023
// - Q1 2023
// - Jan 1 '22 - Jan 20 '23
func FormatDateRange(from, to time.Time, options DateRangeFormatOptions) string {
	// Set default values if not provided
	if options.Today.IsZero() {
		options.Today = time.Now()
	}
	if options.Locale == "" {
		options.Locale = "en_US"
	}
	if options.Separator == "" {
		options.Separator = "-"
	}

	// Extract language code from locale if needed (e.g. "en_US" -> "en")
	lang := options.Locale
	if parts := strings.Split(options.Locale, "_"); len(parts) > 0 {
		lang = parts[0]
	}

	sameYear := from.Year() == to.Year()
	sameMonth := from.Month() == to.Month() && sameYear
	sameDay := from.Day() == to.Day() && sameMonth
	thisYear := from.Year() == options.Today.Year()
	thisDay := from.Day() == options.Today.Day() &&
		from.Month() == options.Today.Month() &&
		from.Year() == options.Today.Year()

	var yearSuffix string
	if !thisYear {
		yearSuffix = fmt.Sprintf(", %d", from.Year())
	}

	formatTimeFunc := createFormatTime(options.Locale)

	var startTimeSuffix, endTimeSuffix string
	if options.IncludeTime {
		if !isSameMinute(startOfDay(from), from) {
			startTimeSuffix = ", " + formatTimeFunc(from)
		}
		if !isSameMinute(endOfDay(to), to) {
			endTimeSuffix = ", " + formatTimeFunc(to)
		}
	}

	// Check if the range is the entire year
	if isSameMinute(startOfYear(from), from) && isSameMinute(endOfYear(to), to) {
		return fmt.Sprintf("%d", from.Year())
	}

	// Check if the range is an entire quarter
	if isSameMinute(startOfQuarter(from), from) &&
		isSameMinute(endOfQuarter(to), to) &&
		getQuarter(from) == getQuarter(to) {
		return fmt.Sprintf("Q%d %d", getQuarter(from), from.Year())
	}

	// Check if the range is across entire month
	if isSameMinute(startOfMonth(from), from) && isSameMinute(endOfMonth(to), to) {
		if sameMonth && sameYear {
			// Example: January 2023
			return getMonthName(from.Month(), lang, false) + " " + strconv.Itoa(from.Year())
		}
		// Example: Jan - Feb 2023
		return fmt.Sprintf("%s %s %s %d",
			getMonthName(from.Month(), lang, true),
			options.Separator,
			getMonthName(to.Month(), lang, true),
			to.Year())
	}

	// Range across years
	if !sameYear {
		fromYear := " '" + fmt.Sprintf("%02d", from.Year()%100)
		toYear := " '" + fmt.Sprintf("%02d", to.Year()%100)

		return fmt.Sprintf("%s %d%s %s %s %d%s",
			getMonthName(from.Month(), lang, true),
			from.Day(),
			fromYear+startTimeSuffix,
			options.Separator,
			getMonthName(to.Month(), lang, true),
			to.Day(),
			toYear+endTimeSuffix)
	}

	// Range across months
	if !sameMonth {
		return fmt.Sprintf("%s %d%s %s %s %d%s%s",
			getMonthName(from.Month(), lang, true),
			from.Day(),
			startTimeSuffix,
			options.Separator,
			getMonthName(to.Month(), lang, true),
			to.Day(),
			endTimeSuffix,
			yearSuffix)
	}

	// Range across days
	if !sameDay {
		// Check for a time suffix, if so print the month twice
		if startTimeSuffix != "" || endTimeSuffix != "" {
			return fmt.Sprintf("%s %d%s %s %s %d%s%s",
				getMonthName(from.Month(), lang, true),
				from.Day(),
				startTimeSuffix,
				options.Separator,
				getMonthName(to.Month(), lang, true),
				to.Day(),
				endTimeSuffix,
				yearSuffix)
		}

		// Example: Jan 1 - 12[, 2023]
		return fmt.Sprintf("%s %d %s %d%s",
			getMonthName(from.Month(), lang, true),
			from.Day(),
			options.Separator,
			to.Day(),
			yearSuffix)
	}

	// Same day, different times
	if startTimeSuffix != "" || endTimeSuffix != "" {
		// If it's today, don't include the date
		if thisDay {
			return fmt.Sprintf("%s %s %s",
				formatTimeFunc(from),
				options.Separator,
				formatTimeFunc(to))
		}

		// Example: Jan 1, 12pm - 1pm[, 2023]
		return fmt.Sprintf("%s %d%s %s %s%s",
			getMonthName(from.Month(), lang, true),
			from.Day(),
			startTimeSuffix,
			options.Separator,
			formatTimeFunc(to),
			yearSuffix)
	}

	// Full day
	// Example: Fri, Jan 1[, 2023]
	return fmt.Sprintf("%s, %s %d%s",
		getWeekdayName(from.Weekday(), lang, true),
		getMonthName(from.Month(), lang, true),
		from.Day(),
		yearSuffix)
}
