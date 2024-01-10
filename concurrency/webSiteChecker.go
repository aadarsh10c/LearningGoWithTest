package concurrency

type WebsiteChecker interface {
	CallBack(url string) bool
}

func CheckWebsites(c WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)
	for _, url := range urls {
		result[url] = c.CallBack(url)
	}
	return result
}
