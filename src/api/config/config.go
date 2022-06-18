package config

import "os"

const (
	secretApiGitAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(secretApiGitAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}
