package constants

// CountryInfo : CountryInfo
type CountryInfo struct {
	Country    string `json:"country"`
	MerchGroup string `json:"merchGroup"`
	Language   string `json:"language"`
	Currency   string `json:"currency"`
	Locale     string `json:"locale"`
}

// GetCountryInfo : Get country info
func GetCountryInfo(country string) *CountryInfo {
	switch country {
	// US Merchgroup
	case "US":
		return &CountryInfo{
			Country:    "US",
			Currency:   "USD",
			Language:   "en",
			Locale:     "en_US",
			MerchGroup: "US",
		}

	// JP Merchgroup
	case "JP":
		return &CountryInfo{
			Country:    "JP",
			Currency:   "JPY",
			Language:   "en",
			Locale:     "en_US",
			MerchGroup: "JP",
		}

	// CN Merchgroup
	case "CN":
		return &CountryInfo{
			Country:    "CN",
			Currency:   "CNY",
			Locale:     "en_US",
			Language:   "zh-Hans",
			MerchGroup: "CN",
		}

	// EU Merchgroup
	case "AT":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "BE":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "BG":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "BGN",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "HR":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "HRK",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "CZ":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "DK":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "DKK",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "FI":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "FR":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "fr_FR",
			MerchGroup: "EU",
		}
	case "DE":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "de_DE",
			MerchGroup: "EU",
		}
	case "GR":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "el_GR",
			MerchGroup: "EU",
		}
	case "HU":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "IE":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "IL":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "ILS",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "IT":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "it_IT",
			MerchGroup: "EU",
		}
	case "LU":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "NL":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "NO":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "NOK",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "PL":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "PLN",
			Locale:     "pl_PL",
			MerchGroup: "EU",
		}
	case "PT":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "RO":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "RON",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "RU":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "RUB",
			Locale:     "ru_RU",
			MerchGroup: "EU",
		}
	case "SK":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "ES":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "EUR",
			Locale:     "es_ES",
			MerchGroup: "EU",
		}
	case "SE":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "SEK",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "CH":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "CHF",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}
	case "TR":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "TRY",
			Locale:     "tr_TR",
			MerchGroup: "EU",
		}
	case "GB":
		return &CountryInfo{
			Country:    "GB",
			Language:   "en-GB",
			Currency:   "GBP",
			Locale:     "en_GB",
			MerchGroup: "EU",
		}

	// XA Merchgroup
	case "ID":
		return &CountryInfo{
			Country:    "ID",
			Language:   "en-GB",
			Currency:   "IDR",
			Locale:     "en_GB",
			MerchGroup: "XA",
		}
	case "IN":
		return &CountryInfo{
			Country:    "IN",
			Language:   "en-GB",
			Currency:   "INR",
			Locale:     "en_GB",
			MerchGroup: "XA",
		}
	case "MY":
		return &CountryInfo{
			Country:    "MY",
			Language:   "en-GB",
			Currency:   "MYR",
			Locale:     "en_GB",
			MerchGroup: "XA",
		}
	case "PH":
		return &CountryInfo{
			Country:    "PH",
			Language:   "en-GB",
			Currency:   "PHP",
			Locale:     "en_GB",
			MerchGroup: "XA",
		}
	case "SG":
		return &CountryInfo{
			Country:    "SG",
			Language:   "en-GB",
			Currency:   "SGD",
			Locale:     "en_GB",
			MerchGroup: "XA",
		}
	case "TH":
		return &CountryInfo{
			Country:    "TH",
			Language:   "en-GB",
			Currency:   "THB",
			Locale:     "th-TH",
			MerchGroup: "XA",
		}
	case "VN":
		return &CountryInfo{
			Country:    "VN",
			Language:   "en-GB",
			Currency:   "VND",
			Locale:     "vi-VN",
			MerchGroup: "XA",
		}
	case "TW":
		return &CountryInfo{
			Country:    "TW",
			Currency:   "TWD",
			Language:   "zh-Hant",
			Locale:     "zh-TW",
			MerchGroup: "XA",
		}

	// XP Merchgroup
	case "AU":
		return &CountryInfo{
			Country:    "AU",
			Language:   "en-GB",
			Currency:   "AUD",
			Locale:     "en_GB",
			MerchGroup: "XP",
		}
	case "CA":
		return &CountryInfo{
			Country:    "CA",
			Language:   "en-GB",
			Currency:   "CAD",
			Locale:     "en_GB",
			MerchGroup: "XP",
		}
	case "NZ":
		return &CountryInfo{
			Country:    "NZ",
			Language:   "en-GB",
			Currency:   "NZD",
			Locale:     "en_GB",
			MerchGroup: "XP",
		}

	// MX Merchgroup
	case "MX":
		return &CountryInfo{
			Country:    "MX",
			Language:   "es-419",
			Currency:   "MXN",
			Locale:     "es-LA",
			MerchGroup: "MX",
		}

	// Default Value
	default:
		return &CountryInfo{
			Country:  country,
			Language: "en-GB",
			Locale:   "en_GB",
		}
	}
}
