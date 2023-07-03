import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getAllURLs(urlString string) ([]string, error) {
	resp, err := http.Get(urlString)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	urls := make([]string, 0)
	baseURL, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}

		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					absoluteURL := resolveURL(baseURL, attr.Val)
					urls = append(urls, absoluteURL)
					break
				}
			}
		}
	}

	return urls, nil
}

func resolveURL(baseURL *url.URL, href string) string {
	relativeURL, err := url.Parse(href)
	if err != nil {
		return ""
	}

	resolvedURL := baseURL.ResolveReference(relativeURL)
	return resolvedURL.String()
}

func main() {
	urls, err := getAllURLs("https://stackoverflow.com/")
	if err != nil {
		log.Fatal(err)
	}

	for _, url := range urls {
		fmt.Println(url)
	}
}