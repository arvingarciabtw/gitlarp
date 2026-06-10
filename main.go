package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Config struct {
	Message   string
	Count     int
	StartDate time.Time
	EndDate   time.Time
}

func main() {
	config := parse()
	verify()
	larp(config)
	push()
}

func parse() Config {
	layout := "2006-01-02"
	today := time.Now().Format(layout)

	var message string
	var count int
	var start string
	var end string

	messageDefault := "larping..."
	countDefault := 1
	startDefault := today
	endDefault := today

	flag.StringVar(&message, "m", messageDefault, "Commit message")
	flag.StringVar(&message, "message", messageDefault, "Commit message")
	flag.IntVar(&count, "c", countDefault, "Commit count")
	flag.IntVar(&count, "count", countDefault, "Commit count")
	flag.StringVar(&start, "s", startDefault, "Start date (YYYY-MM-DD)")
	flag.StringVar(&start, "start", startDefault, "Start date (YYYY-MM-DD)")
	flag.StringVar(&end, "e", endDefault, "End date (YYYY-MM-DD)")
	flag.StringVar(&end, "end", endDefault, "End date (YYYY-MM-DD)")

	flag.Parse()

	start, end = trimDates(start, end)
	startDate, endDate := parseDates(start, end, layout)
	validate(startDate, endDate, count)

	return Config{
		Message:   message,
		Count:     count,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func trimDates(start string, end string) (string, string) {
	return strings.TrimSpace(start), strings.TrimSpace(end)
}

func parseDates(start string, end string, layout string) (time.Time, time.Time) {
	startDate, err := time.Parse(layout, start)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid start date. Format for dates are in YYYY-MM-DD.\n\n")
		flag.Usage()
		os.Exit(1)
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid end date. Format for dates are in YYYY-MM-DD.\n\n")
		flag.Usage()
		os.Exit(1)
	}
	return startDate, endDate
}

func validate(startDate time.Time, endDate time.Time, count int) {
	if startDate.After(endDate) {
		fmt.Fprintf(os.Stderr, "Error: Invalid date range.\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if startDate.After(time.Now()) || endDate.After(time.Now()) {
		fmt.Fprintf(os.Stderr, "Error: Cannot make commits for future dates.\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if count < 1 {
		fmt.Fprintf(os.Stderr, "Error: Commit count should at least be 1.\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if count > 50 {
		fmt.Fprintf(os.Stderr, "Error: Commit count is capped at 50.\n\n")
		flag.Usage()
		os.Exit(1)
	}
}

func verify() {
	checkRepo := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := checkRepo.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: You are not in a Git repository.")
		os.Exit(1)
	}
}

func larp(cfg Config) {
	for d := cfg.StartDate; !d.After(cfg.EndDate); d = d.AddDate(0, 0, 1) {
		for i := 0; i < cfg.Count; i++ {
			gitDate := d.Format(time.RFC3339)
			cmd := exec.Command("git", "commit", "--allow-empty", "-m", cfg.Message, "--date", gitDate)
			fmt.Printf("\rLarping a commit for date: %s (%d/%d)\n", gitDate, i+1, cfg.Count)
			time.Sleep(100 * time.Millisecond)

			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Failed to commit for date %s: %v\n", gitDate, err)
				os.Exit(1)
			}
		}
	}
}

func push() {
	push := exec.Command("git", "push", "origin", "main")
	err := push.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to push to remote: %v\n", err)
		fmt.Fprintln(os.Stderr, "You can undo local commits with 'git reset --hard origin/main'")
		os.Exit(1)
	}

	fmt.Println("Successfully larped!")
}
