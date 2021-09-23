package countwi

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil ||
		n.Data == "script" ||
		n.Data == "style" {
		return
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
		sp := strings.Split(n.Data, " ")
		words = words + len(sp)
	}

	w, i := countWordsAndImages(n.FirstChild)
	words = words + w
	images = images + i

	w, i = countWordsAndImages(n.NextSibling)
	words = words + w
	images = images + i
	return
}
