package tld

import (
	"golang.org/x/net/publicsuffix"
)

func GetTLD(domain string) string {
	tld, _ := publicsuffix.PublicSuffix(domain)
	return tld
}
