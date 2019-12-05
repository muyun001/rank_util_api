package utils

import (
	"regexp"
	"strings"
)

func TopDomain(domain string) string {
	re := regexp.MustCompile(`^(?:[a-z][a-z0-9\-_]*\.)*?([a-z0-9][a-z0-9\-]*?\.(?:com\.cn|net\.cn|org\.cn|com|net|cn|org|cc|top|vip))$`)

	subMatch := re.FindStringSubmatch(domain)
	if len(subMatch) == 2 {
		return subMatch[1]
	}

	return domain
}

// FormatUrl 检查scheme，没有则加上http://
func FormatUrl(url string) string {
	url = strings.Trim(url, " ")
	schemeRegx := regexp.MustCompile(`^https?://.*`)
	if false == schemeRegx.MatchString(url) {
		url = "http://" + url
	}

	return url
}
