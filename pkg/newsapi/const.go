package newsapi

type Article struct {
	source      string
	author      string
	title       string
	description string
	url         string
	urlToImage  string
	publishedAt string
	content     string
}

type Source struct {
	id          string
	name        string
	description string
	url         string
	category    string
	language    string
	country     string
}

type Language string
type Country string
type Category string
type SortMethod string

const (
	CAT_business      Category = "business"
	CAT_entertainment Category = "entertainment"
	CAT_general       Category = "general"
	CAT_health        Category = "health"
	CAT_science       Category = "science"
	CAT_sports        Category = "sports"
	CAT_technology    Category = "technology"
)

const (
	SORT_relevancy   SortMethod = "relevancy"
	SORT_popularity  SortMethod = "popularity"
	SORT_publishedAt SortMethod = "publishedAt"
)

const (
	LANG_AR Language = "ar"
	LANG_DE Language = "de"
	LANG_EN Language = "en"
	LANG_ES Language = "es"
	LANG_FR Language = "fr"
	LANG_HE Language = "he"
	LANG_IT Language = "it"
	LANG_NL Language = "nl"
	LANG_NO Language = "no"
	LANG_PT Language = "pt"
	LANG_RU Language = "ru"
	LANG_SV Language = "sv"
	LANG_UD Language = "ud"
	LANG_ZH Language = "zh"
)

const (
	COUNTRY_AE Country = "ae"
	COUNTRY_AR Country = "ar"
	COUNTRY_AT Country = "at"
	COUNTRY_AU Country = "au"
	COUNTRY_BE Country = "be"
	COUNTRY_BG Country = "bg"
	COUNTRY_BR Country = "br"
	COUNTRY_CA Country = "ca"
	COUNTRY_CH Country = "ch"
	COUNTRY_CN Country = "cn"
	COUNTRY_CO Country = "co"
	COUNTRY_CU Country = "cu"
	COUNTRY_CZ Country = "cz"
	COUNTRY_DE Country = "de"
	COUNTRY_EG Country = "eg"
	COUNTRY_FR Country = "fr"
	COUNTRY_GB Country = "gb"
	COUNTRY_GR Country = "gr"
	COUNTRY_HK Country = "hk"
	COUNTRY_HU Country = "hu"
	COUNTRY_ID Country = "id"
	COUNTRY_IE Country = "ie"
	COUNTRY_IL Country = "il"
	COUNTRY_IN Country = "in"
	COUNTRY_IT Country = "it"
	COUNTRY_JP Country = "jp"
	COUNTRY_KR Country = "kr"
	COUNTRY_LT Country = "lt"
	COUNTRY_LV Country = "lv"
	COUNTRY_MA Country = "ma"
	COUNTRY_MX Country = "mx"
	COUNTRY_MY Country = "my"
	COUNTRY_NG Country = "ng"
	COUNTRY_NL Country = "nl"
	COUNTRY_NO Country = "no"
	COUNTRY_NZ Country = "nz"
	COUNTRY_PH Country = "ph"
	COUNTRY_PL Country = "pl"
	COUNTRY_PT Country = "pt"
	COUNTRY_RO Country = "ro"
	COUNTRY_RS Country = "rs"
	COUNTRY_RU Country = "ru"
	COUNTRY_SA Country = "sa"
	COUNTRY_SE Country = "se"
	COUNTRY_SG Country = "sg"
	COUNTRY_SI Country = "si"
	COUNTRY_SK Country = "sk"
	COUNTRY_TH Country = "th"
	COUNTRY_TR Country = "tr"
	COUNTRY_TW Country = "tw"
	COUNTRY_UA Country = "ua"
	COUNTRY_US Country = "us"
	COUNTRY_VE Country = "ve"
	COUNTRY_ZA Country = "za"
)
