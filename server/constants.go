package server

// ...
const DefaultPort = 4242

// ...
const Host = "planetocd.org"

// ...
const DonateURL = "https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=23LG7JTZSCA54"

// ...
const Email = "aurelien.ocd@gmail.com"

// ...
var FacebookURL = map[string]string{"fr": "https://www.facebook.com/PlanetOCDfr",
	"es": "https://www.facebook.com/Planet-OCD-Trastorno-Obsesivo-Compulsivo-TOC-106315328115819",
	"zh": "https://www.facebook.com/Planet-OCD-%E5%85%B3%E4%BA%8E%E5%BC%BA%E8%BF%AB%E7%97%87%E7%9A%84%E6%96%87%E7%AB%A0-103706645049739"}

// ...
const GithubURL = "https://github.com/momow/planetocd"

// ...
const GithubContributeURL = GithubURL + "/blob/master/README.md#contribute"

// ...
const GoogleAnalyticsID = "UA-137169000-1"

// ...
const SiteName = "Planet OCD"

// ...
var TwitterHandle = map[string]string{"fr": "PlanetOCD_fr",
	"es": "PlanetOCD_es",
	"zh": "PlanetOCD_zh"}

// ...
var TwitterHost = "https://twitter.com/"

// ...
var PetitionURL = "https://www.change.org/p/1er-ministre-jean-castex-et-ministre-de-la-sant%C3%A9-olivier-v%C3%A9ran-toc-am%C3%A9liorer-la-prise-en-charge-par-les-services-de-sant%C3%A9"

// ...
var DismissBannerCookieName = "dismiss-banner"

// ...
var Constants = map[string]interface{}{
	"DefaultPort":             DefaultPort,
	"Host":                    Host,
	"DonateURL":               DonateURL,
	"Email":                   Email,
	"FacebookURL":             FacebookURL,
	"GithubURL":               GithubURL,
	"GithubContributeURL":     GithubContributeURL,
	"GoogleAnalyticsID":       GoogleAnalyticsID,
	"SiteName":                SiteName,
	"TwitterHandle":           TwitterHandle,
	"TwitterHost":             TwitterHost,
	"PetitionURL":             PetitionURL,
	"DismissBannerCookieName": DismissBannerCookieName,
}
