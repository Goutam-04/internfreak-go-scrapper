package main

import (
	"fmt"
	// "time"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	l := launcher.New().Headless(true)
	url := l.MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()

	page := browser.MustPage("https://internfreak.co/jobs-and-internship-opportunities?page=1&limit=10")

	page.MustWaitLoad()

	posts := page.MustElements("h2.heading.mb-3 a")

	for _, post := range posts {
		title := post.MustText()
		link := post.MustProperty("href").String()

		fmt.Println("Title:", title)
		fmt.Println("Link:", link)

		postPage := browser.MustPage(link)
		postPage.MustWaitLoad()

		applyLinkElement, err := postPage.Element("#applylink")
		if err != nil {
			fmt.Println("Error finding apply link:", err)
			continue
		}

		if applyLinkElement != nil {
			applyLink := applyLinkElement.MustProperty("href").String()
			fmt.Println("Apply Link:", applyLink)
		} else {
			fmt.Println("Apply Link not found on this page.")
		}

		postPage.MustClose()

		// time.Sleep(2 * time.Second)
	}

	defer browser.MustClose()
}
