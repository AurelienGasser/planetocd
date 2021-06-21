package server

// DefaultPort ...
const DefaultPort = 4242

// Host ...
const Host = "planetocd.org"

// DonateURL ...
const DonateURL = "https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=23LG7JTZSCA54"

// Email ...
const Email = "aurelien.ocd@gmail.com"

// FacebookURL ...
var FacebookURL = map[string]string{"fr": "https://www.facebook.com/PlanetOCDfr",
	"es": "https://www.facebook.com/Planet-OCD-Trastorno-Obsesivo-Compulsivo-TOC-106315328115819",
	"zh": "https://www.facebook.com/Planet-OCD-%E5%85%B3%E4%BA%8E%E5%BC%BA%E8%BF%AB%E7%97%87%E7%9A%84%E6%96%87%E7%AB%A0-103706645049739"}

// GithubURL ...
const GithubURL = "https://github.com/momow/planetocd"

// GithubContributeURL ...
const GithubContributeURL = GithubURL + "/blob/master/README.md#contribute"

// GoogleAnalyticsID ...
const GoogleAnalyticsID = "UA-137169000-1"

// SiteName ...
const SiteName = "Planet OCD"

// TwitterHandle ...
var TwitterHandle = map[string]string{"fr": "PlanetOCD_fr",
	"es": "PlanetOCD_es",
	"zh": "PlanetOCD_zh"}

// TwitterHost ...
var TwitterHost = "https://twitter.com/"

// Constants ...
var Constants = map[string]interface{}{
	"DefaultPort":         DefaultPort,
	"Host":                Host,
	"DonateURL":           DonateURL,
	"Email":               Email,
	"FacebookURL":         FacebookURL,
	"GithubURL":           GithubURL,
	"GithubContributeURL": GithubContributeURL,
	"GoogleAnalyticsID":   GoogleAnalyticsID,
	"SiteName":            SiteName,
	"TwitterHandle":       TwitterHandle,
	"TwitterHost":         TwitterHost,
}
