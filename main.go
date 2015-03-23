package main

//import "net/http"
import (
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"os"
)

type itemHandler int

func (h *itemHandler) ProcessItems(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	for i := range newitems {
		fmt.Println("Title: %s", newitems[i].Title)
	}
}

type chanHandler func(f *rss.Feed, newchannels []*rss.Channel)

func (h *chanHandler) ProcessChannels(feed *rss.Feed, newchannels []*rss.Channel) {
	// Empty channel handler
}

func main() {

	// Pre-downloaded test
	xmlFile, _ := os.Open("rss.xml")

	defer xmlFile.Close()

	feed := rss.NewWithHandlers(5, false, chanHandler, itemHandler)

	feed.Fetch("http://www.ubuntu.com/usn/rss.xml", nil)
}
