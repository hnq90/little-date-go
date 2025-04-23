//go:build embed_resources
// +build embed_resources

package littledate

import (
	"bytes"
)

// Init function will be called after package-level variable initialization
// It will load the embedded translations
func init() {
	// Add these translators after the normal init() has run
	// The order is important because we need the bundle to be initialized first
	registerEmbeddedTranslations()
}

// Register embedded translations
func registerEmbeddedTranslations() {
	// English
	if err := loadEmbeddedTranslation("en", englishTranslation); err == nil {
		// Successfully loaded
	}

	// French
	if err := loadEmbeddedTranslation("fr", frenchTranslation); err == nil {
		// Successfully loaded
	}

	// Spanish
	if err := loadEmbeddedTranslation("es", spanishTranslation); err == nil {
		// Successfully loaded
	}

	// German
	if err := loadEmbeddedTranslation("de", germanTranslation); err == nil {
		// Successfully loaded
	}

	// Japanese
	if err := loadEmbeddedTranslation("ja", japaneseTranslation); err == nil {
		// Successfully loaded
	}

	// Korean
	if err := loadEmbeddedTranslation("ko", koreanTranslation); err == nil {
		// Successfully loaded
	}

	// Chinese Simplified
	if err := loadEmbeddedTranslation("zh-CN", chineseSimplifiedTranslation); err == nil {
		// Successfully loaded
	}

	// Chinese Traditional
	if err := loadEmbeddedTranslation("zh-TW", chineseTraditionalTranslation); err == nil {
		// Successfully loaded
	}

	// Vietnamese
	if err := loadEmbeddedTranslation("vi", vietnameseTranslation); err == nil {
		// Successfully loaded
	}
}

// Load a translation from an embedded string
func loadEmbeddedTranslation(locale string, translation string) error {
	// Only load if the bundle exists
	if bundle == nil {
		return nil
	}

	// Create a reader from the string
	reader := bytes.NewReader([]byte(translation))

	// Parse the JSON
	_, err := bundle.ParseMessageFileBytes(reader.Bytes(), locale+".json")
	return err
}

// Embedded translations as strings
var englishTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "January"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "February"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "March"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "April"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "May"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "June"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "July"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "August"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "September"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "October"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "November"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "December"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "Jan"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "Feb"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "Mar"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "Apr"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "May"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "Jun"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "Jul"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "Aug"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "Sep"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "Oct"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "Nov"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "Dec"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "Sunday"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "Monday"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "Tuesday"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "Wednesday"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "Thursday"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "Friday"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "Saturday"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "Sun"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "Mon"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "Tue"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "Wed"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "Thu"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "Fri"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "Sat"
  }
}`

var frenchTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "Janvier"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "Février"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "Mars"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "Avril"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "Mai"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "Juin"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "Juillet"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "Août"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "Septembre"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "Octobre"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "Novembre"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "Décembre"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "Jan"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "Fév"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "Mar"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "Avr"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "Mai"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "Jun"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "Jul"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "Aoû"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "Sep"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "Oct"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "Nov"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "Déc"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "Dimanche"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "Lundi"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "Mardi"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "Mercredi"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "Jeudi"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "Vendredi"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "Samedi"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "Dim"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "Lun"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "Mar"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "Mer"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "Jeu"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "Ven"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "Sam"
  }
}`

var spanishTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "Enero"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "Febrero"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "Marzo"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "Abril"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "Mayo"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "Junio"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "Julio"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "Agosto"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "Septiembre"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "Octubre"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "Noviembre"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "Diciembre"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "Ene"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "Feb"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "Mar"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "Abr"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "May"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "Jun"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "Jul"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "Ago"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "Sep"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "Oct"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "Nov"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "Dic"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "Domingo"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "Lunes"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "Martes"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "Miércoles"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "Jueves"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "Viernes"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "Sábado"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "Dom"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "Lun"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "Mar"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "Mié"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "Jue"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "Vie"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "Sáb"
  }
}`

var germanTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "Januar"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "Februar"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "März"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "April"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "Mai"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "Juni"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "Juli"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "August"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "September"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "Oktober"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "November"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "Dezember"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "Jan"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "Feb"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "Mär"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "Apr"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "Mai"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "Jun"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "Jul"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "Aug"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "Sep"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "Okt"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "Nov"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "Dez"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "Sonntag"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "Montag"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "Dienstag"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "Mittwoch"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "Donnerstag"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "Freitag"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "Samstag"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "So"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "Mo"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "Di"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "Mi"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "Do"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "Fr"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "Sa"
  }
}`

var japaneseTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "1月"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "2月"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "3月"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "4月"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "5月"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "6月"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "7月"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "8月"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "9月"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "10月"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "11月"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "12月"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "1月"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "2月"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "3月"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "4月"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "5月"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "6月"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "7月"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "8月"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "9月"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "10月"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "11月"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "12月"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "日曜日"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "月曜日"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "火曜日"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "水曜日"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "木曜日"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "金曜日"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "土曜日"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "日"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "月"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "火"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "水"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "木"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "金"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "土"
  }
}`

var koreanTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "1월"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "2월"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "3월"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "4월"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "5월"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "6월"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "7월"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "8월"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "9월"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "10월"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "11월"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "12월"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "1월"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "2월"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "3월"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "4월"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "5월"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "6월"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "7월"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "8월"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "9월"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "10월"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "11월"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "12월"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "일요일"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "월요일"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "화요일"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "수요일"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "목요일"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "금요일"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "토요일"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "일"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "월"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "화"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "수"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "목"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "금"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "토"
  }
}`

var chineseSimplifiedTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "一月"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "二月"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "三月"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "四月"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "五月"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "六月"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "七月"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "八月"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "九月"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "十月"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "十一月"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "十二月"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "1月"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "2月"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "3月"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "4月"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "5月"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "6月"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "7月"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "8月"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "9月"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "10月"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "11月"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "12月"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "星期日"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "星期一"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "星期二"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "星期三"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "星期四"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "星期五"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "星期六"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "日"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "一"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "二"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "三"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "四"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "五"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "六"
  }
}`

var chineseTraditionalTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "一月"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "二月"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "三月"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "四月"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "五月"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "六月"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "七月"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "八月"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "九月"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "十月"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "十一月"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "十二月"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "1月"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "2月"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "3月"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "4月"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "5月"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "6月"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "7月"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "8月"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "9月"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "10月"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "11月"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "12月"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "星期日"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "星期一"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "星期二"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "星期三"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "星期四"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "星期五"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "星期六"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "日"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "一"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "二"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "三"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "四"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "五"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "六"
  }
}`

var vietnameseTranslation = `{
  "month.long.1": {
    "description": "Full name of January",
    "other": "Tháng Một"
  },
  "month.long.2": {
    "description": "Full name of February",
    "other": "Tháng Hai"
  },
  "month.long.3": {
    "description": "Full name of March",
    "other": "Tháng Ba"
  },
  "month.long.4": {
    "description": "Full name of April",
    "other": "Tháng Tư"
  },
  "month.long.5": {
    "description": "Full name of May",
    "other": "Tháng Năm"
  },
  "month.long.6": {
    "description": "Full name of June",
    "other": "Tháng Sáu"
  },
  "month.long.7": {
    "description": "Full name of July",
    "other": "Tháng Bảy"
  },
  "month.long.8": {
    "description": "Full name of August",
    "other": "Tháng Tám"
  },
  "month.long.9": {
    "description": "Full name of September",
    "other": "Tháng Chín"
  },
  "month.long.10": {
    "description": "Full name of October",
    "other": "Tháng Mười"
  },
  "month.long.11": {
    "description": "Full name of November",
    "other": "Tháng Mười Một"
  },
  "month.long.12": {
    "description": "Full name of December",
    "other": "Tháng Mười Hai"
  },
  "month.short.1": {
    "description": "Short name of January",
    "other": "Th1"
  },
  "month.short.2": {
    "description": "Short name of February",
    "other": "Th2"
  },
  "month.short.3": {
    "description": "Short name of March",
    "other": "Th3"
  },
  "month.short.4": {
    "description": "Short name of April",
    "other": "Th4"
  },
  "month.short.5": {
    "description": "Short name of May",
    "other": "Th5"
  },
  "month.short.6": {
    "description": "Short name of June",
    "other": "Th6"
  },
  "month.short.7": {
    "description": "Short name of July",
    "other": "Th7"
  },
  "month.short.8": {
    "description": "Short name of August",
    "other": "Th8"
  },
  "month.short.9": {
    "description": "Short name of September",
    "other": "Th9"
  },
  "month.short.10": {
    "description": "Short name of October",
    "other": "Th10"
  },
  "month.short.11": {
    "description": "Short name of November",
    "other": "Th11"
  },
  "month.short.12": {
    "description": "Short name of December",
    "other": "Th12"
  },
  "weekday.long.0": {
    "description": "Full name of Sunday",
    "other": "Chủ Nhật"
  },
  "weekday.long.1": {
    "description": "Full name of Monday",
    "other": "Thứ Hai"
  },
  "weekday.long.2": {
    "description": "Full name of Tuesday",
    "other": "Thứ Ba"
  },
  "weekday.long.3": {
    "description": "Full name of Wednesday",
    "other": "Thứ Tư"
  },
  "weekday.long.4": {
    "description": "Full name of Thursday",
    "other": "Thứ Năm"
  },
  "weekday.long.5": {
    "description": "Full name of Friday",
    "other": "Thứ Sáu"
  },
  "weekday.long.6": {
    "description": "Full name of Saturday",
    "other": "Thứ Bảy"
  },
  "weekday.short.0": {
    "description": "Short name of Sunday",
    "other": "CN"
  },
  "weekday.short.1": {
    "description": "Short name of Monday",
    "other": "T2"
  },
  "weekday.short.2": {
    "description": "Short name of Tuesday",
    "other": "T3"
  },
  "weekday.short.3": {
    "description": "Short name of Wednesday",
    "other": "T4"
  },
  "weekday.short.4": {
    "description": "Short name of Thursday",
    "other": "T5"
  },
  "weekday.short.5": {
    "description": "Short name of Friday",
    "other": "T6"
  },
  "weekday.short.6": {
    "description": "Short name of Saturday",
    "other": "T7"
  }
}`
