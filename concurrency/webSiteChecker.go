package concurrency

// import "time"

type WebsiteChecker func(string) bool
type result struct {
	url string
	is  bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(url string) {
			resultChannel <- result{url: url, is: wc(url)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <- resultChannel
		results[result.url] = result.is
	}
	// time.Sleep(2 * time.Second)
	return results
}
