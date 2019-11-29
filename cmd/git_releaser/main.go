package main

import (
	"git-releaser/internal/integrations/gitlab"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

/*
 * Steps:
 * 1. Get the project url
 * 2. If the GitLab API allows
 * 2.1. Send an API request to create the merge request from develop -> master
 * 3. Send some notification
 * 4. Run for the hug :D
 */
func main() {
	pvtToken := os.Getenv("PERSONAL_TOKEN")
	gitUrl := os.Getenv("GITLAB_URL")

	gl := gitlab.New(gitUrl, pvtToken)
	projId, _ := strconv.ParseInt(os.Getenv("PROJECT_ID"), 10, 0)
	proj := gitlab.NewProject(int(projId), "", "", "", gl)

	proj.SubmitMergeRequest(&gitlab.MergeRequestOptions{
		SourceBranch: "develop",
		TargetBranch: "master",
		Title:        "Automatic release of 1.0.0",
	})
}
