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
		log.Fatalf("Error: Invalid start date: %v", err)
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		log.Fatalf("Error: Invalid end date: %v", err)
	}

	// input validation
	if startDate.After(endDate) {
		fmt.Fprintf(os.Stderr, "Error: Invalid date range\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if startDate.After(time.Now()) || endDate.After(time.Now()) {
		fmt.Fprintf(os.Stderr, "Error: Cannot make commits for future dates\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if count < 1 {
		fmt.Fprintf(os.Stderr, "Error: Commit count should at least be 1\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if count > 50 {
		fmt.Fprintf(os.Stderr, "Error: Commit count is capped at 50\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// check for git
	checkRepo := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err = checkRepo.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: You are not in a Git repository.")
		os.Exit(1)
	}

	// commits for each date
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		for i := 0; i < count; i++ {
			gitDate := d.Format(time.RFC3339)
			cmd := exec.Command("git", "commit", "--allow-empty", "-m", message, "--date", gitDate)

			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Failed to commit for date %s: %v\n", gitDate, err)
				os.Exit(1)
			}
		}
	}

	// push to remote
	push := exec.Command("git", "push", "origin", "main")
	err = push.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to push to remote: %v", err)
		os.Exit(1)
	}

	fmt.Println("Successfully larped!")
}
