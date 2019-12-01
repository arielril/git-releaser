package main

import (
	"git-releaser/internal/cmd"
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
	cmd.Execute()
}
