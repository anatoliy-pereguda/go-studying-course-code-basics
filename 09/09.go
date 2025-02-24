package main

import (
	"fmt"
)

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}
	return fmt.Sprintf("%s.%s", locale, domain)
}

func main() {
	tests := []struct {
		domain string
		locale string
		result string
	}{
		{"site.com", "", "en.site.com"},
		{"site.com", "ru", "ru.site.com"},
		{"example.org", "", "en.example.org"},
		{"example.org", "fr", "fr.example.org"},
		{"mydomain.net", "de", "de.mydomain.net"},
	}

	for _, test := range tests {
		output := DomainForLocale(test.domain, test.locale)
		fmt.Printf("DomainForLocale(%q, %q) = %q (expected: %q)\n",
			test.domain, test.locale, output, test.result)
	}
}
