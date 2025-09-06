package parser

// Browser detection patterns
const (
	// Chrome patterns
	ChromePattern        = `Chrome/`
	ChromeVersionPattern = `Chrome/(\d+\.\d+)`

	// Firefox patterns
	FirefoxPattern        = `Firefox/`
	FirefoxVersionPattern = `Firefox/(\d+\.\d+)`

	// Safari patterns
	SafariPattern        = `Safari/`
	SafariVersionPattern = `Version/(\d+\.\d+)`

	// Edge patterns
	EdgePattern        = `Edg/`
	EdgeVersionPattern = `Edg/(\d+\.\d+)`

	// OperaPattern Opera patterns
	OperaPattern        = `OPR/`
	OperaVersionPattern = `OPR/(\d+\.\d+)`

	// Internet Explorer patterns
	IEPattern        = `(?:MSIE|Trident)`
	IEVersionPattern = `(?:MSIE\s(\d+\.\d+)|rv:(\d+\.\d+))`
)

// Operating System patterns
const (
	// Windows patterns
	WindowsPattern        = `Windows`
	WindowsVersionPattern = `Windows NT (\d+\.\d+)`

	// macOS patterns
	MacOSPattern        = `Mac OS X`
	MacOSVersionPattern = `Mac OS X (\d+[._]\d+)`

	// iOS patterns
	IOSPattern        = `OS \d+_\d+`
	IOSVersionPattern = `OS (\d+_\d+)`

	// Android patterns
	AndroidPattern        = `Android`
	AndroidVersionPattern = `Android (\d+\.\d+)`

	// Linux patterns
	LinuxPattern  = `Linux`
	UbuntuPattern = `Ubuntu`

	// Chrome OS patterns
	ChromeOSPattern        = `CrOS`
	ChromeOSVersionPattern = `CrOS [a-zA-Z0-9_]+ (\d+\.\d+)`
)

// Device type patterns
const (
	// Mobile patterns
	MobilePattern = `Mobile`
	iPhonePattern = `iPhone`

	// Tablet patterns
	TabletPattern = `Tablet|iPad`
	iPadPattern   = `iPad`

	// Samsung patterns
	SamsungPattern      = `SM-`
	SamsungModelPattern = `SM-([A-Z0-9]+)`
)

// Bot detection patterns
const (
	BotPattern1  = `(?i)bot`
	BotPattern2  = `(?i)crawler`
	BotPattern3  = `(?i)spider`
	BotPattern4  = `(?i)scraper`
	BotPattern5  = `(?i)slurp`
	BotPattern6  = `(?i)googlebot`
	BotPattern7  = `(?i)bingbot`
	BotPattern8  = `(?i)yandexbot`
	BotPattern9  = `(?i)facebookexternalhit`
	BotPattern10 = `(?i)twitterbot`
	BotPattern11 = `(?i)linkedinbot`
	BotPattern12 = `(?i)whatsapp`
	BotPattern13 = `(?i)telegrambot`
)

// Browser engines
const (
	BlinkEngine   = "Blink"
	GeckoEngine   = "Gecko"
	WebKitEngine  = "WebKit"
	TridentEngine = "Trident"
)

// Device types
const (
	DeviceTypeDesktop = "desktop"
	DeviceTypeMobile  = "mobile"
	DeviceTypeTablet  = "tablet"
	DeviceTypeBot     = "bot"
)

// Brand names
const (
	BrandApple   = "Apple"
	BrandSamsung = "Samsung"
)

// Windows version mapping
var WindowsVersionMap = map[string]string{
	"10.0": "10/11",
	"6.3":  "8.1",
	"6.2":  "8",
	"6.1":  "7",
	"6.0":  "Vista",
	"5.2":  "XP x64",
	"5.1":  "XP",
	"5.0":  "2000",
}

// All bot patterns for easy iteration
var BotPatterns = []string{
	BotPattern1, BotPattern2, BotPattern3, BotPattern4, BotPattern5,
	BotPattern6, BotPattern7, BotPattern8, BotPattern9, BotPattern10,
	BotPattern11, BotPattern12, BotPattern13,
}
