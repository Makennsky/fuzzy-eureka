package parser

import (
	"regexp"
	"strings"

	"github.com/Makennsky/fuzzy-eureka/internal/types"
)

// BrowserRule defines a browser detection rule
type BrowserRule struct {
	Name           string
	Pattern        string
	VersionPattern string
	Engine         string
}

// OSRule defines an OS detection rule
type OSRule struct {
	Name           string
	Pattern        string
	VersionPattern string
	VersionMap     map[string]string
}

// DeviceRule defines a device detection rule
type DeviceRule struct {
	Type         string
	Pattern      string
	Brand        string
	ModelPattern string
	IsMobile     bool
	IsTablet     bool
	IsBot        bool
}

// Browser detection rules in priority order
var browserRules = []BrowserRule{
	{"Edge", EdgePattern, EdgeVersionPattern, BlinkEngine},
	{"Opera", OperaPattern, OperaVersionPattern, BlinkEngine},
	{"Chrome", ChromePattern, ChromeVersionPattern, BlinkEngine},
	{"Firefox", FirefoxPattern, FirefoxVersionPattern, GeckoEngine},
	{"Safari", SafariPattern, SafariVersionPattern, WebKitEngine},
	{"Internet Explorer", IEPattern, IEVersionPattern, TridentEngine},
}

// OS detection rules in priority order
var osRules = []OSRule{
	{"iOS", IOSPattern, IOSVersionPattern, nil},
	{"Windows", WindowsPattern, WindowsVersionPattern, WindowsVersionMap},
	{"macOS", MacOSPattern, MacOSVersionPattern, nil},
	{"Android", AndroidPattern, AndroidVersionPattern, nil},
	{"Chrome OS", ChromeOSPattern, ChromeOSVersionPattern, nil},
	{"Ubuntu", UbuntuPattern, "", nil},
	{"Linux", LinuxPattern, "", nil},
}

// Device detection rules in priority order
var deviceRules = []DeviceRule{
	{DeviceTypeBot, "", "", "", false, false, true}, // Special case for bots
	{DeviceTypeTablet, iPadPattern, BrandApple, "", false, true, false},
	{DeviceTypeMobile, iPhonePattern, BrandApple, "", true, false, false},
	{DeviceTypeMobile, SamsungPattern, BrandSamsung, SamsungModelPattern, true, false, false},
	{DeviceTypeTablet, TabletPattern, "", "", false, true, false},
	{DeviceTypeMobile, MobilePattern, "", "", true, false, false},
}

// ParseBrowser parses browser information from User-Agent string
func ParseBrowser(userAgent string) types.BrowserInfoImpl {
	for _, rule := range browserRules {
		if matched, _ := regexp.MatchString(rule.Pattern, userAgent); matched {
			browser := types.BrowserInfoImpl{
				Name:   rule.Name,
				Engine: rule.Engine,
			}
			
			if rule.VersionPattern != "" {
				if re := regexp.MustCompile(rule.VersionPattern); re.MatchString(userAgent) {
					if matches := re.FindStringSubmatch(userAgent); len(matches) > 1 {
						// Handle IE special case with multiple capture groups
						if rule.Name == "Internet Explorer" {
							if matches[1] != "" {
								browser.Version = matches[1]
							} else if len(matches) > 2 && matches[2] != "" {
								browser.Version = matches[2]
							}
						} else {
							browser.Version = matches[1]
						}
					}
				}
			}
			return browser
		}
	}
	return types.BrowserInfoImpl{}
}

// ParseOS parses operating system information from User-Agent string
func ParseOS(userAgent string) types.OSInfoImpl {
	for _, rule := range osRules {
		if matched, _ := regexp.MatchString(rule.Pattern, userAgent); matched {
			os := types.OSInfoImpl{Name: rule.Name}
			
			if rule.VersionPattern != "" {
				if re := regexp.MustCompile(rule.VersionPattern); re.MatchString(userAgent) {
					if matches := re.FindStringSubmatch(userAgent); len(matches) > 1 {
						version := matches[1]
						
						// Apply version mapping if available
						if rule.VersionMap != nil {
							if mappedVersion, exists := rule.VersionMap[version]; exists {
								version = mappedVersion
							}
						}
						
						// Handle special cases for underscores in iOS/macOS versions
						if rule.Name == "iOS" || rule.Name == "macOS" {
							version = strings.ReplaceAll(version, "_", ".")
						}
						
						os.Version = version
					}
				}
			}
			return os
		}
	}
	return types.OSInfoImpl{}
}

// ParseDevice parses device information from User-Agent string
func ParseDevice(userAgent string) types.DeviceInfoImpl {
	// Check for bots first (highest priority)
	for _, pattern := range BotPatterns {
		if matched, _ := regexp.MatchString(pattern, userAgent); matched {
			return types.DeviceInfoImpl{
				Type: DeviceTypeBot,
				Bot:  true,
			}
		}
	}

	// Check other device rules
	for _, rule := range deviceRules {
		if rule.IsBot {
			continue // Skip bot rule as it's handled above
		}
		
		if matched, _ := regexp.MatchString(rule.Pattern, userAgent); matched {
			device := types.DeviceInfoImpl{
				Type:   rule.Type,
				Brand:  rule.Brand,
				Mobile: rule.IsMobile,
				Tablet: rule.IsTablet,
				Bot:    rule.IsBot,
			}
			
			// Extract model if pattern is provided
			if rule.ModelPattern != "" {
				if re := regexp.MustCompile(rule.ModelPattern); re.MatchString(userAgent) {
					if matches := re.FindStringSubmatch(userAgent); len(matches) > 1 {
						device.Model = matches[1]
					}
				}
			} else if rule.Brand == BrandApple {
				// Set default models for Apple devices
				if rule.Type == DeviceTypeTablet {
					device.Model = "iPad"
				} else if rule.Type == DeviceTypeMobile {
					device.Model = "iPhone"
				}
			}
			
			return device
		}
	}

	// Default to desktop if nothing else matched
	return types.DeviceInfoImpl{Type: DeviceTypeDesktop}
}
