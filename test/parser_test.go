package test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/Makennsky/fuzzy-eureka/pkg/requestparser"
)

func TestRequestParser_Parse(t *testing.T) {
	parser := requestparser.NewRequestParser()

	req := &http.Request{
		Method:     "GET",
		Header:     make(http.Header),
		Host:       "example.com",
		RemoteAddr: "203.0.113.45:54321",
	}

	testURL, _ := url.Parse("https://example.com/test")
	req.URL = testURL

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,ru;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("X-Forwarded-For", "192.168.1.100")
	req.Header.Set("DNT", "1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	userInfo := parser.Parse(req)

	if userInfo.GetIP() != "192.168.1.100" {
		t.Errorf("Ожидался IP 192.168.1.100, получен %s", userInfo.GetIP())
	}

	if userInfo.GetHost() != "example.com" {
		t.Errorf("Ожидался хост example.com, получен %s", userInfo.GetHost())
	}

	if userInfo.GetMethod() != "GET" {
		t.Errorf("Ожидался метод GET, получен %s", userInfo.GetMethod())
	}

	if userInfo.GetUserAgent() == "" {
		t.Error("User-Agent не должен быть пустым")
	}

	browser := userInfo.GetBrowser()
	if browser.GetName() != "Chrome" {
		t.Errorf("Ожидался браузер Chrome, получен %s", browser.GetName())
	}

	os := userInfo.GetOS()
	if os.GetName() != "Windows" {
		t.Errorf("Ожидалась ОС Windows, получена %s", os.GetName())
	}

	device := userInfo.GetDevice()
	if device.GetType() != "desktop" {
		t.Errorf("Ожидался тип устройства desktop, получен %s", device.GetType())
	}

	languages := userInfo.GetAcceptLanguage()
	if len(languages) == 0 {
		t.Error("Список языков не должен быть пустым")
	}
	if languages[0] != "en-US" {
		t.Errorf("Первый язык должен быть en-US, получен %s", languages[0])
	}

	if !userInfo.IsDoNotTrack() {
		t.Error("Флаг Do Not Track должен быть true")
	}

	if !userInfo.IsUpgradeInsecure() {
		t.Error("Флаг Upgrade Insecure Requests должен быть true")
	}

	json, err := userInfo.ToJSON()
	if err != nil {
		t.Errorf("Ошибка при создании JSON: %v", err)
	}
	if json == "" {
		t.Error("JSON не должен быть пустым")
	}

	fingerprint := userInfo.GetClientFingerprint()
	if fingerprint == "" {
		t.Error("Отпечаток клиента не должен быть пустым")
	}
}

func TestMobileUserAgent(t *testing.T) {
	parser := requestparser.NewRequestParser()

	req := &http.Request{
		Method: "GET",
		Header: make(http.Header),
		Host:   "example.com",
	}

	testURL, _ := url.Parse("https://example.com/test")
	req.URL = testURL

	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1")

	userInfo := parser.Parse(req)

	device := userInfo.GetDevice()
	if !device.IsMobile() {
		t.Error("Устройство должно быть мобильным")
	}

	if device.GetBrand() != "Apple" {
		t.Errorf("Ожидался бренд Apple, получен %s", device.GetBrand())
	}

	os := userInfo.GetOS()
	if os.GetName() != "iOS" {
		t.Errorf("Ожидалась ОС iOS, получена %s", os.GetName())
	}
}

func TestBotUserAgent(t *testing.T) {
	parser := requestparser.NewRequestParser()

	req := &http.Request{
		Method: "GET",
		Header: make(http.Header),
		Host:   "example.com",
	}

	testURL, _ := url.Parse("https://example.com/test")
	req.URL = testURL

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	userInfo := parser.Parse(req)

	device := userInfo.GetDevice()
	if !device.IsBot() {
		t.Error("Устройство должно быть ботом")
	}

	if device.GetType() != "bot" {
		t.Errorf("Ожидался тип bot, получен %s", device.GetType())
	}
}
