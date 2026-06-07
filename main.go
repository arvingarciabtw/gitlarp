package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("=== GITLARP ===")

	// time setup
	layout := "2006-01-02"
	today := time.Now().Format(layout)

	// flags
	var message string
	var count int
	var start string
	var end string

	flag.StringVar(&message, "m", "larping...", "Commit message")
	flag.StringVar(&message, "message", "larping...", "Commit message")
	flag.IntVar(&count, "c", 1, "Commit count")
	flag.IntVar(&count, "count", 1, "Commit count")
	flag.StringVar(&start, "s", today, "Start date (YYYY/MM/DD)")
	flag.StringVar(&start, "start", today, "Start date (YYYY/MM/DD)")
	flag.StringVar(&end, "e", today, "End date (YYYY/MM/DD)")
	flag.StringVar(&end, "end", today, "End date (YYYY/MM/DD)")

	flag.Parse()

	// TODO: remove logging once done with program
	fmt.Println("\nmessage:", message)
	fmt.Println("count:", count)
	fmt.Println("start:", start)
	fmt.Println("end:", end)

	// time parsing
	startDate, err := time.Parse(layout, start)
	if err != nil {
		log.Fatalf("Invalid start date: %v", err)
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		log.Fatalf("Invalid end date: %v", err)
	}

	// TODO: remove logging once done with program
	fmt.Println("startDate:", startDate)
	fmt.Println("endDate:", endDate)

	if startDate.After(endDate) {
		fmt.Fprintf(os.Stderr, "Error: invalid date range\n\n")
		flag.Usage()
		os.Exit(1)
	}
}
