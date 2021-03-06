package yelp

import "fmt"

// ValidateLocale checks if the provided locale is one of Yelp's supported locales.
func ValidateLocale(locale string) error {
	if _, ok := validLocales[locale]; !ok {
		return fmt.Errorf("Invalid locale provided: %s", locale)
	}
	return nil
}

// validLocales are the valid locales from the Yelp Fusion API. This list was pulled
// from https://www.yelp.com/developers/documentation/v3/supported_locales.
var validLocales = map[string]struct{}{
	"cs_CZ":  struct{}{},
	"da_DK":  struct{}{},
	"de_AT":  struct{}{},
	"de_CH":  struct{}{},
	"de_DE":  struct{}{},
	"en_AU":  struct{}{},
	"en_BE":  struct{}{},
	"en_CA":  struct{}{},
	"en_CH":  struct{}{},
	"en_GB":  struct{}{},
	"en_HK":  struct{}{},
	"en_IE":  struct{}{},
	"en_MY":  struct{}{},
	"en_NZ":  struct{}{},
	"en_PH":  struct{}{},
	"en_SG":  struct{}{},
	"en_US":  struct{}{},
	"es_AR":  struct{}{},
	"es_CL":  struct{}{},
	"es_ES":  struct{}{},
	"es_MX":  struct{}{},
	"fi_FI":  struct{}{},
	"fil_PH": struct{}{},
	"fr_BE":  struct{}{},
	"fr_CA":  struct{}{},
	"fr_CH":  struct{}{},
	"fr_FR":  struct{}{},
	"it_CH":  struct{}{},
	"it_IT":  struct{}{},
	"ja_JP":  struct{}{},
	"ms_MY":  struct{}{},
	"nb_NO":  struct{}{},
	"nl_BE":  struct{}{},
	"nl_NL":  struct{}{},
	"pl_PL":  struct{}{},
	"pt_BR":  struct{}{},
	"pt_PT":  struct{}{},
	"sv_FI":  struct{}{},
	"sv_SE":  struct{}{},
	"tr_TR":  struct{}{},
	"zh_HK":  struct{}{},
	"zh_TW":  struct{}{},
}
