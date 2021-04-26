package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"os"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	segID := r.URL.Path[1:]

	if segID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dom, err := parseSegment(segID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Fprint(w, dom)
}

const urlFormat = "https://www.strava.com/segments/%s"

func parseSegment(segmentID string) (dom string, err error) {
	c := colly.NewCollector()

	c.OnHTML("div#segment-leaderboard", func(e *colly.HTMLElement) {
		dom, err = e.DOM.Html()
	})

	c.Visit(fmt.Sprintf(urlFormat, segmentID))

	return
}
