package podimator

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

type Update struct {
    // Podcast name to update
    PodcastName string
    // Number of episodes to update
    Episodes string
    // Date range for update
    DateRange string
}

func (podi *Podimator) Update(up Update) {
    if len(up.PodcastName) > 0 {
        // Apply filter to retrieve single podcast
        err := podi.filter(up.PodcastName)
        if err != nil {
            fmt.Fprintf(os.Stderr, "unable to filter podcast %s: %v\n", up.PodcastName, err)
            return
        }
    }

    for _, p := range podi.Config.Podcasts {
        feed, err := podi.FeedParser.ParseURL(p.URL)
        if err != nil {
            fmt.Fprintf(os.Stderr, "unable to parse %s: %w\n", p, err)
            continue
        }
        requests := buildRequests([]*gofeed.Item{feed.Items[0]}, podi.Config.Location + "/" + p.Name)
        fmt.Printf("[\033[0;35mupdating\033[0m] %s\n", p.Name)
        podi.download(requests)
    }
}