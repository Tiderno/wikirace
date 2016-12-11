package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// arguments
	fromTitle string
	toTitle   string
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s from_title to_title\n", os.Args[0])
}

func init() {
	flag.Usage = usage
	flag.Parse()

	fromTitle = flag.Arg(0)
	toTitle = flag.Arg(1)

	if len(fromTitle) == 0 || len(toTitle) == 0 {
		usage()
		os.Exit(1)
	}
}

func main() {
	forwardGraph := NewPageGraph(fromTitle)
	forwardQueue := NewStringQueue()
	forwardQueue.Enqueue(fromTitle)

	for !forwardQueue.IsEmpty() {
		title := forwardQueue.Dequeue()
		log.Printf("VISIT %#v", title)
		for _, link := range forwardGraph.UnvisitedLinks(title) {
			if link == toTitle {
				for _, title := range forwardGraph.Path(link) {
					fmt.Println(title)
				}
				os.Exit(1)
			}
			forwardQueue.Enqueue(link)
		}
	}
}
