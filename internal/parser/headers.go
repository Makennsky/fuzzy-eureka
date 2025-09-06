package parser

import (
	"strings"
)

func ParseAcceptLanguage(header string) []string {
	if header == "" {
		return []string{}
	}

	var languages []string
	parts := strings.Split(header, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if idx := strings.Index(part, ";"); idx != -1 {
			part = part[:idx]
		}
		languages = append(languages, strings.TrimSpace(part))
	}

	return languages
}

func ParseAcceptEncoding(header string) []string {
	if header == "" {
		return []string{}
	}

	var encodings []string
	parts := strings.Split(header, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if idx := strings.Index(part, ";"); idx != -1 {
			part = part[:idx]
		}
		encodings = append(encodings, strings.TrimSpace(part))
	}

	return encodings
}
