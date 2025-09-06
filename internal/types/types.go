package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Makennsky/fuzzy-eureka/pkg/interfaces"
)

type UserInfoImpl struct {
	IP              string            `json:"ip"`
	RealIP          string            `json:"real_ip,omitempty"`
	ForwardedFor    string            `json:"forwarded_for,omitempty"`
	UserAgent       string            `json:"user_agent"`
	Browser         BrowserInfoImpl   `json:"browser"`
	OS              OSInfoImpl        `json:"os"`
	Device          DeviceInfoImpl    `json:"device"`
	Headers         map[string]string `json:"headers"`
	Cookies         map[string]string `json:"cookies"`
	AcceptLanguage  []string          `json:"accept_language"`
	AcceptEncoding  []string          `json:"accept_encoding"`
	ContentType     string            `json:"content_type,omitempty"`
	Referer         string            `json:"referer,omitempty"`
	Origin          string            `json:"origin,omitempty"`
	Host            string            `json:"host"`
	Method          string            `json:"method"`
	URL             string            `json:"url"`
	Protocol        string            `json:"protocol"`
	Secure          bool              `json:"is_secure"`
	RequestTime     time.Time         `json:"request_time"`
	ContentLength   int64             `json:"content_length"`
	DNT             bool              `json:"do_not_track"`
	Connection      string            `json:"connection,omitempty"`
	CacheControl    string            `json:"cache_control,omitempty"`
	UpgradeInsecure bool              `json:"upgrade_insecure_requests"`
}

type BrowserInfoImpl struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Engine  string `json:"engine"`
}

type OSInfoImpl struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type DeviceInfoImpl struct {
	Type   string `json:"type"`
	Brand  string `json:"brand,omitempty"`
	Model  string `json:"model,omitempty"`
	Mobile bool   `json:"is_mobile"`
	Tablet bool   `json:"is_tablet"`
	Bot    bool   `json:"is_bot"`
}

func (u *UserInfoImpl) GetIP() string                 { return u.IP }
func (u *UserInfoImpl) GetRealIP() string             { return u.RealIP }
func (u *UserInfoImpl) GetForwardedFor() string       { return u.ForwardedFor }
func (u *UserInfoImpl) GetUserAgent() string          { return u.UserAgent }
func (u *UserInfoImpl) GetBrowser() interfaces.BrowserInfo   { return &u.Browser }
func (u *UserInfoImpl) GetOS() interfaces.OSInfo             { return &u.OS }
func (u *UserInfoImpl) GetDevice() interfaces.DeviceInfo     { return &u.Device }
func (u *UserInfoImpl) GetHeaders() map[string]string { return u.Headers }
func (u *UserInfoImpl) GetCookies() map[string]string { return u.Cookies }
func (u *UserInfoImpl) GetAcceptLanguage() []string   { return u.AcceptLanguage }
func (u *UserInfoImpl) GetAcceptEncoding() []string   { return u.AcceptEncoding }
func (u *UserInfoImpl) GetContentType() string        { return u.ContentType }
func (u *UserInfoImpl) GetReferer() string            { return u.Referer }
func (u *UserInfoImpl) GetOrigin() string             { return u.Origin }
func (u *UserInfoImpl) GetHost() string               { return u.Host }
func (u *UserInfoImpl) GetMethod() string             { return u.Method }
func (u *UserInfoImpl) GetURL() string                { return u.URL }
func (u *UserInfoImpl) GetProtocol() string           { return u.Protocol }
func (u *UserInfoImpl) IsSecure() bool                { return u.Secure }
func (u *UserInfoImpl) GetRequestTime() time.Time     { return u.RequestTime }
func (u *UserInfoImpl) GetContentLength() int64       { return u.ContentLength }
func (u *UserInfoImpl) IsDoNotTrack() bool            { return u.DNT }
func (u *UserInfoImpl) GetConnection() string         { return u.Connection }
func (u *UserInfoImpl) GetCacheControl() string       { return u.CacheControl }
func (u *UserInfoImpl) IsUpgradeInsecure() bool       { return u.UpgradeInsecure }

func (u *UserInfoImpl) ToJSON() (string, error) {
	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (u *UserInfoImpl) GetClientFingerprint() string {
	return fmt.Sprintf("%s|%s|%s|%s|%s",
		u.IP,
		u.UserAgent,
		strings.Join(u.AcceptLanguage, ","),
		strings.Join(u.AcceptEncoding, ","),
		u.Host,
	)
}

func (b *BrowserInfoImpl) GetName() string    { return b.Name }
func (b *BrowserInfoImpl) GetVersion() string { return b.Version }
func (b *BrowserInfoImpl) GetEngine() string  { return b.Engine }

func (o *OSInfoImpl) GetName() string    { return o.Name }
func (o *OSInfoImpl) GetVersion() string { return o.Version }

func (d *DeviceInfoImpl) GetType() string  { return d.Type }
func (d *DeviceInfoImpl) GetBrand() string { return d.Brand }
func (d *DeviceInfoImpl) GetModel() string { return d.Model }
func (d *DeviceInfoImpl) IsMobile() bool   { return d.Mobile }
func (d *DeviceInfoImpl) IsTablet() bool   { return d.Tablet }
func (d *DeviceInfoImpl) IsBot() bool      { return d.Bot }
