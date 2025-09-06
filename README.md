# Request Parser Library

A comprehensive Go library for parsing HTTP requests and extracting detailed user information.

## Features

- 📱 Device type detection (desktop, mobile, tablet, bot)
- 🌐 Browser and operating system parsing
- 🔍 Real IP address extraction (proxy-aware)
- 🍪 Cookie and header processing
- 🌍 Language preference parsing
- 🔒 Security flag detection (DNT, Upgrade-Insecure-Requests)
- 📊 Unique client fingerprinting
- 🧪 Full test coverage
- 🔌 Interface-based design for easy testing and extension

## Project Structure

```
├── pkg/requestparser/       # Public API
│   ├── interfaces.go        # Interfaces
│   └── parser.go           # Main implementation
├── internal/               # Internal packages
│   ├── types/             # Data types
│   └── parser/            # Parsers
│       ├── ip.go          # IP extraction
│       ├── useragent.go   # User-Agent parsing
│       └── headers.go     # Header processing
├── test/                  # Tests
├── examples/              # Usage examples
└── README.md
```

## Installation

```bash
go get github.com/Makennsky/fuzzy-eureka
```

## Quick Start

```go
package main

import (
    "fmt"
    "net/http"
    
    "github.com/Makennsky/fuzzy-eureka/pkg/requestparser"
)

func main() {
    parser := requestparser.NewRequestParser()
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        userInfo := parser.Parse(r)
        
        fmt.Printf("IP: %s\n", userInfo.GetIP())
        fmt.Printf("Browser: %s %s\n", userInfo.GetBrowser().GetName(), userInfo.GetBrowser().GetVersion())
        fmt.Printf("OS: %s\n", userInfo.GetOS().GetName())
        fmt.Printf("Device: %s\n", userInfo.GetDevice().GetType())
        
        // JSON representation
        json, _ := userInfo.ToJSON()
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(json))
    })
    
    http.ListenAndServe(":8080", nil)
}
```

## API

### Core Interfaces

#### RequestParser
```go
type RequestParser interface {
    Parse(r *http.Request) UserInfo
}
```

#### UserInfo
```go
type UserInfo interface {
    GetIP() string
    GetRealIP() string
    GetUserAgent() string
    GetBrowser() BrowserInfo
    GetOS() OSInfo
    GetDevice() DeviceInfo
    GetHeaders() map[string]string
    GetCookies() map[string]string
    // ... and many other methods
    ToJSON() (string, error)
    GetClientFingerprint() string
}
```

### Creating Parser

```go
parser := requestparser.NewRequestParser()
userInfo := parser.Parse(httpRequest)
```

### Getting Information

```go
// Basic information
ip := userInfo.GetIP()
method := userInfo.GetMethod()
url := userInfo.GetURL()

// Browser
browser := userInfo.GetBrowser()
browserName := browser.GetName()        // Chrome, Firefox, Safari, etc.
browserVersion := browser.GetVersion()   // 91.0
browserEngine := browser.GetEngine()     // Blink, Gecko, WebKit

// Operating system
os := userInfo.GetOS()
osName := os.GetName()        // Windows, macOS, Linux, Android, iOS
osVersion := os.GetVersion()   // 10, 11.2, etc.

// Device
device := userInfo.GetDevice()
deviceType := device.GetType()    // desktop, mobile, tablet, bot
isMobile := device.IsMobile()
isBot := device.IsBot()

// Additional information
languages := userInfo.GetAcceptLanguage()  // ["en-US", "en", "ru"]
encodings := userInfo.GetAcceptEncoding()  // ["gzip", "deflate", "br"]
cookies := userInfo.GetCookies()           // map[string]string
headers := userInfo.GetHeaders()           // map[string]string
```

## Examples

### Basic Usage
See `examples/basic_usage/main.go`

### Web Server Middleware
See `examples/middleware/main.go`

## Testing

```bash
# Run all tests
go test ./test/...

# Run with verbose output
go test -v ./test/...

# Check coverage
go test -cover ./test/...
```

## Supported Browsers

- ✅ Chrome/Chromium
- ✅ Firefox
- ✅ Safari
- ✅ Edge
- ✅ Opera
- ✅ Others based on WebKit/Blink/Gecko

## Supported Operating Systems

- ✅ Windows (all versions)
- ✅ macOS
- ✅ Linux/Ubuntu
- ✅ Android
- ✅ iOS

## IP Extraction

The library correctly handles IP addresses behind proxies/load balancers:

1. `X-Forwarded-For` (first IP from list)
2. `X-Real-IP`
3. `X-Client-IP`
4. `RemoteAddr` (fallback)

## Bot Detection

Automatically detects popular bots:
- Googlebot
- Bingbot
- YandexBot
- Facebook crawler
- And others

## License

MIT License