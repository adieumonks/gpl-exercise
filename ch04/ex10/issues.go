package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adieumonks/gpl-exercise/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var issuesLessThanAMonth []*github.Issue
	var issuesLessThanAYear []*github.Issue
	var others []*github.Issue

	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthAgo) {
			issuesLessThanAMonth = append(issuesLessThanAMonth, item)
			continue
		}

		if item.CreatedAt.After(oneYearAgo) {
			issuesLessThanAYear = append(issuesLessThanAYear, item)
			continue
		}

		others = append(others, item)
	}

	showIssues(issuesLessThanAMonth, "===== less than a month =====")
	showIssues(issuesLessThanAYear, "===== less than a year =====")
	showIssues(others, "===== more than a year =====")
}

func showIssues(issues []*github.Issue, header string) {
	if len(issues) > 0 {
		fmt.Printf("%s\n", header)
		for _, item := range issues {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
