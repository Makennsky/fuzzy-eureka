package interfaces

import (
	"net/http"
	"time"
)

// RequestParser interface for parsing HTTP requests
type RequestParser interface {
	Parse(r *http.Request) UserInfo
}

// UserInfo represents user information extracted from HTTP request
type UserInfo interface {
	GetIP() string
	GetRealIP() string
	GetForwardedFor() string
	GetUserAgent() string
	GetBrowser() BrowserInfo
	GetOS() OSInfo
	GetDevice() DeviceInfo
	GetHeaders() map[string]string
	GetCookies() map[string]string
	GetAcceptLanguage() []string
	GetAcceptEncoding() []string
	GetContentType() string
	GetReferer() string
	GetOrigin() string
	GetHost() string
	GetMethod() string
	GetURL() string
	GetProtocol() string
	IsSecure() bool
	GetRequestTime() time.Time
	GetContentLength() int64
	IsDoNotTrack() bool
	GetConnection() string
	GetCacheControl() string
	IsUpgradeInsecure() bool
	ToJSON() (string, error)
	GetClientFingerprint() string
}

// BrowserInfo represents browser information
type BrowserInfo interface {
	GetName() string
	GetVersion() string
	GetEngine() string
}

// OSInfo represents operating system information
type OSInfo interface {
	GetName() string
	GetVersion() string
}

// DeviceInfo represents device information
type DeviceInfo interface {
	GetType() string
	GetBrand() string
	GetModel() string
	IsMobile() bool
	IsTablet() bool
	IsBot() bool
}