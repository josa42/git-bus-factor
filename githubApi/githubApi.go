package githubApi

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/oauth2"

	"gopkg.in/AlecAivazis/survey.v1"
	"github.com/google/go-github/github"
	keychain "github.com/lunixbochs/go-keychain"
)

const keychainService = "github.com/josa42/git-bus-factor"
const maxRetries = 3

// ParseURL :
func ParseURL(url string) (string, string, error) {
	if owner, name := matchURL(url, `^([^/:.]+)/([^/:.]+)$`); owner != "" {
		return owner, name, nil
	}

	if owner, name := matchURL(url, `github\.com/([^/:.]+)/([^/:.]+)(\.git)?$`); owner != "" {
		return owner, name, nil
	}

	if owner, name := matchURL(url, `git@github\.com:([^/:..]+)/([^/:..]+)(\.git)?$`); owner != "" {
		return owner, name, nil
	}

	return "", "", errors.New("Repo URL cannot be matched")
}

func matchURL(url string, pattern string) (string, string) {
	re1, err := regexp.Compile(pattern)
	result := re1.FindStringSubmatch(url)

	if err != nil || len(result) == 0 {
		return "", ""
	}

	return result[1], result[2]
}

// Login :
func Login() {

	if getToken() != "" {
		replace := false
		prompt := &survey.Confirm{
			Message: "Replace current token?",
		}
		survey.AskOne(prompt, &replace, nil)
		if !replace {
			return
		}
	}

	token := ""
	prompt := &survey.Password{
		Message: "Token:",
		Help:    "Create a GitHub access token at https://github.com/settings/tokens - Check: public_repo",
	}
	survey.AskOne(prompt, &token, nil)

	if token == "" {
		return
	}

	setToken(token)
	fmt.Println("Added token to keychain")
}

// Logout :
func Logout() {
	removeToken()
	fmt.Println("Removed token from keychain")
}

// HasToken :
func HasToken() bool {
	return getToken() != ""
}

// RepoInfo :
func RepoInfo(owner string, name string) (*github.Repository, error) {
	ctx := context.Background()
	client := client(ctx)

	repo, _, err := client.Repositories.Get(ctx, owner, name)
	if err != nil {
		fmt.Printf("%s\n", err)

		return nil, err
	}

	return repo, nil
}

// OpenRepoPRsCount :
func OpenRepoPRsCount(owner string, name string) (int, error) {
	ctx := context.Background()
	client := client(ctx)

	prs, _, err := client.PullRequests.List(ctx, owner, name, &github.PullRequestListOptions{ListOptions: github.ListOptions{PerPage: 100}})
	if err != nil {
		return 0, err
	}

	return len(prs), nil
}

// ClosedRepoPRsCount :
func ClosedRepoPRsCount(owner string, name string) (int, error) {
	ctx := context.Background()
	client := client(ctx)

	prs, _, err := client.PullRequests.List(ctx, owner, name, &github.PullRequestListOptions{
		State:       "Closed",
		ListOptions: github.ListOptions{PerPage: 100},
	})
	if err != nil {
		return 0, err
	}

	return len(prs), nil
}

// CodeFrequency :
func CodeFrequency(owner string, name string) ([]*github.WeeklyStats, error) {
	return loadCodeFrequency(owner, name, 1)
}

func loadCodeFrequency(owner string, name string, count int) ([]*github.WeeklyStats, error) {
	if count > maxRetries {
		return nil, errors.New("Exceeded max retries for loadCodeFrequency")
	}

	ctx := context.Background()
	client := client(ctx)

	stats, response, err := client.Repositories.ListCodeFrequency(ctx, owner, name)

	if response.StatusCode == 202 {
		time.Sleep(3 * time.Second)
		return loadCodeFrequency(owner, name, count+1)
	}

	if err != nil {
		return nil, err
	}

	return stats, nil
}

// Releases :
func Releases(owner string, name string) ([]*github.RepositoryRelease, error) {
	ctx := context.Background()
	client := client(ctx)

	releases, _, err := client.Repositories.ListReleases(ctx, owner, name, &github.ListOptions{PerPage: 100})

	return releases, err
}

// Contributions :
func Contributions(owner string, name string) ([]*github.ContributorStats, error) {
	return loadContributions(owner, name, 1)
}

func loadContributions(owner string, name string, count int) ([]*github.ContributorStats, error) {
	if count > maxRetries {
		return nil, errors.New("Exceeded max retries for loadCodeFrequency")
	}

	ctx := context.Background()
	client := client(ctx)

	contributions, response, err := client.Repositories.ListContributorsStats(ctx, owner, name)

	if response.StatusCode == 202 {
		time.Sleep(3 * time.Second)
		return loadContributions(owner, name, count+1)
	}

	if err != nil {
		return nil, err
	}

	return contributions, err
}

func client(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: getToken()},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func getToken() string {
	token, error := keychain.Find(keychainService, "token")
	if error == nil && token != "" {
		return token
	}

	return ""
}

func setToken(token string) bool {
	error := keychain.Add(keychainService, "token", token)
	return error == nil
}

func removeToken() bool {
	error := keychain.Remove(keychainService, "token")
	return error == nil
}
