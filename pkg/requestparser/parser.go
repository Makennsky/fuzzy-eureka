package requestparser

import (
	"net/http"
	"strings"
	"time"

	"github.com/Makennsky/fuzzy-eureka/internal/parser"
	"github.com/Makennsky/fuzzy-eureka/internal/types"
	"github.com/Makennsky/fuzzy-eureka/pkg/interfaces"
)

type RequestParserImpl struct{}

func NewRequestParser() RequestParser {
	return &RequestParserImpl{}
}

func (p *RequestParserImpl) Parse(r *http.Request) interfaces.UserInfo {
	userInfo := &types.UserInfoImpl{
		Headers:     make(map[string]string),
		Cookies:     make(map[string]string),
		RequestTime: time.Now(),
		Method:      r.Method,
		Protocol:    r.Proto,
		Host:        r.Host,
		Secure:      r.TLS != nil,
	}

	if r.URL != nil {
		userInfo.URL = r.URL.String()
	}

	userInfo.IP = parser.ExtractIP(r)
	userInfo.RealIP = r.Header.Get("X-Real-IP")
	userInfo.ForwardedFor = r.Header.Get("X-Forwarded-For")

	for name, values := range r.Header {
		userInfo.Headers[name] = strings.Join(values, ", ")
	}

	for _, cookie := range r.Cookies() {
		userInfo.Cookies[cookie.Name] = cookie.Value
	}

	userAgent := r.Header.Get("User-Agent")
	userInfo.UserAgent = userAgent
	userInfo.Browser = parser.ParseBrowser(userAgent)
	userInfo.OS = parser.ParseOS(userAgent)
	userInfo.Device = parser.ParseDevice(userAgent)

	userInfo.AcceptLanguage = parser.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	userInfo.AcceptEncoding = parser.ParseAcceptEncoding(r.Header.Get("Accept-Encoding"))
	userInfo.ContentType = r.Header.Get("Content-Type")
	userInfo.Referer = r.Header.Get("Referer")
	userInfo.Origin = r.Header.Get("Origin")
	userInfo.Connection = r.Header.Get("Connection")
	userInfo.CacheControl = r.Header.Get("Cache-Control")

	if r.ContentLength > 0 {
		userInfo.ContentLength = r.ContentLength
	}

	if dnt := r.Header.Get("DNT"); dnt == "1" {
		userInfo.DNT = true
	}

	if uir := r.Header.Get("Upgrade-Insecure-Requests"); uir == "1" {
		userInfo.UpgradeInsecure = true
	}

	return userInfo
}
