package busFactor

import (
	"fmt"
	"math"
	"os"

	"github.com/josa42/git-bus-factor/githubApi"
	"github.com/justincampbell/timeago"
)

const forksThreshold = 5
const watcherThreshold = 5
const starsThreshold = 5
const contributionThreshold = 0.7

// Print :
func Print(owner string, name string) {
	repo, err := githubApi.RepoInfo(owner, name)
	if err != nil {
		fmt.Printf("RepoInfo: %s\n", err)
		os.Exit(1)
	}

	// Forks
	if *repo.ForksCount > forksThreshold {
		fmt.Printf("🍴  %d forks.\n", *repo.ForksCount)
	} else {
		fmt.Printf("🍴  Few forks (%d).\n", *repo.ForksCount)
	}

	// Watchers
	if *repo.SubscribersCount > watcherThreshold {
		fmt.Printf("🔭  %d watchers.\n", *repo.SubscribersCount)
	} else {
		fmt.Printf("🔭  Few watchers (%d).\n", *repo.SubscribersCount)
	}

	// Stars
	if *repo.StargazersCount > watcherThreshold {
		fmt.Printf("🌟  %d stars.\n", *repo.StargazersCount)
	} else {
		fmt.Printf("🌟  Few stars (%d).\n", *repo.StargazersCount)
	}

	// Age
	created := timeago.FromTime(repo.CreatedAt.Time)
	pushed := timeago.FromTime(repo.PushedAt.Time)
	fmt.Printf("📆  Created about %s; last push %s.\n", created, pushed)

	// PRs
	openPRsCount, err2 := githubApi.OpenRepoPRsCount(owner, name)
	if err2 != nil {
		fmt.Println(err2)
	}

	closedPRsCount, err2 := githubApi.ClosedRepoPRsCount(owner, name)
	if err2 != nil {
		fmt.Println(err2)
	}

	totalPRsCount := openPRsCount + closedPRsCount
	if totalPRsCount > 0 {
		prsRatio := (float64(openPRsCount) / float64(totalPRsCount)) * 100
		fmt.Printf("🍻  %d PRs: %d closed; %d open; %.2f%% are closed.\n", totalPRsCount, closedPRsCount, openPRsCount, prsRatio)
	} else {
		fmt.Printf("🍻  No PRs opened yet for this repository.\n")
	}

	// Refactoring
	stats, err3 := githubApi.CodeFrequency(owner, name)
	if err3 != nil {
		fmt.Println(err3)
	}

	additions := 0
	deletions := 0
	for _, stat := range stats {
		additions += *stat.Additions
		deletions += *stat.Deletions
	}

	refactingRatio := (math.Abs(float64(deletions)) / float64(additions)) * 100

	fmt.Printf("🛠️  Deletions to additions ratio: %.2f%% (%d/%d).\n", refactingRatio, deletions, additions)

	// Releases

	releases, err4 := githubApi.Releases(owner, name)
	if err4 != nil {
		fmt.Println(err4)
	}

	releasesCount := len(releases)
	if releasesCount == 0 {
		fmt.Printf("📦  No releases.\n")
	} else {
		latesRelease := releases[0]
		published := timeago.FromTime(latesRelease.PublishedAt.Time)
		fmt.Printf("📦  %d releases; latest release \"%s\": %s.\n", releasesCount, *latesRelease.Name, published)
	}

	// Bus factor
	// Contributions
	contributions, err5 := githubApi.Contributions(owner, name)
	if err5 != nil {
		fmt.Println(err5)
	}

	totalContributions := 0
	maxContributions := 0
	minContributions := math.MaxInt8

	for _, contribution := range contributions {
		if *contribution.Total > maxContributions {
			maxContributions = *contribution.Total
		}

		if *contribution.Total < minContributions {
			minContributions = *contribution.Total
		}

		totalContributions += *contribution.Total
	}

	delta := maxContributions - minContributions

	meaningfulCount := 0
	for _, contribution := range contributions {
		if float64(maxContributions-*contribution.Total) < float64(delta)*contributionThreshold {
			meaningfulCount++
		}
	}

	busFactor := 0.0
	if meaningfulCount == 0 {
		busFactor = 100
	} else {
		averageContributions := 0.0
		for _, contribution := range contributions {
			averageContributions += float64(*contribution.Total) / float64(totalContributions)
		}

		busFactor = (averageContributions / float64(meaningfulCount) * 100.0)
	}

	if busFactor > 90 {
		fmt.Printf("🚌  Bus factor: %2.f%%. Most likely one core contributor.\n", busFactor)
	} else {
		fmt.Printf("🚌  Bus factor: %2.f%% (%d impactful contributors out of %d).\n", busFactor, meaningfulCount, len(contributions))
	}
}
