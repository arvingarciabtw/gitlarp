package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	flag.StringVar(&start, "s", today, "Start date (YYYY-MM-DD)")
	flag.StringVar(&start, "start", today, "Start date (YYYY-MM-DD)")
	flag.StringVar(&end, "e", today, "End date (YYYY-MM-DD)")
	flag.StringVar(&end, "end", today, "End date (YYYY-MM-DD)")

	flag.Parse()

	// time parsing
	startDate, err := time.Parse(layout, start)
	if err != nil {
		log.Fatalf("Invalid start date: %v", err)
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		log.Fatalf("Invalid end date: %v", err)
	}

	if startDate.After(endDate) {
		fmt.Fprintf(os.Stderr, "Error: invalid date range\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if count < 1 {
		fmt.Fprintf(os.Stderr, "Error: count should at least be 1\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// commits for each date
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		for i := 0; i < count; i++ {
			gitDate := d.Format(time.RFC3339)
			cmd := exec.Command("git", "commit", "--allow-empty", "-m", message, "--date", gitDate)

			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to commit for date %s: %v\n", gitDate, err)
				os.Exit(1)
			}
		}
	}

	// push to remote
	push := exec.Command("git", "push", "origin", "main")
	err = push.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to push to remote: %v", err)
		os.Exit(1)
	}

	fmt.Println("Successfully larped!")
}
