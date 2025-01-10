package storage

import (
	"sync"
)

var (
	urlMap       = make(map[string]string)
	domainCounts = make(map[string]int)
	mutex        = &sync.Mutex{}
)

func SaveURLMapping(shortURL, originalURL string) {
	mutex.Lock()
	defer mutex.Unlock()

	urlMap[shortURL] = originalURL
	domain := extractDomain(originalURL)
	domainCounts[domain]++
}

func GetOriginalURL(shortURL string) string {
	mutex.Lock()
	defer mutex.Unlock()

	return urlMap[shortURL]
}

func GetTopDomains() map[string]int {
	mutex.Lock()
	defer mutex.Unlock()

	result := make(map[string]int)
	count := 0
	for domain, total := range domainCounts {
		if count >= 3 {
			break
		}
		result[domain] = total
		count++
	}
	return result
}

func extractDomain(url string) string {
	// Simplified domain extraction
	if len(url) > 0 {
		return url
	}
	return "unknown"
}
