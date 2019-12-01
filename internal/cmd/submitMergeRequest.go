package cmd

import (
	"fmt"
	"git-releaser/internal/integrations/gitlab"
	t "git-releaser/internal/integrations/telegram"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	submitMergeRequest = &cobra.Command{
		Use:   "submit-mr",
		Short: "Create the releases merge requests based on the config file",
		RunE:  run,
	}
	telegramIntegration t.Telegram
	gitlabIntegration   gitlab.Gitlab
	project             gitlab.Project
)

func init() {
	cobra.OnInitialize(configMergeRequest)

	rootCmd.AddCommand(submitMergeRequest)
}

func configMergeRequest() {
	if gitlabIntegration == nil {
		pvtToken := viper.GetString("user.personalToken")
		gitUrl := viper.GetString("gitlab.url")
		gitlabIntegration = gitlab.New(gitUrl, pvtToken)

		project = gitlab.NewProject(viper.GetInt("project.id"), viper.GetString("project.name"), "", "", gitlabIntegration)
	}

	if telegramIntegration == nil {
		telegramIntegration = t.New(viper.GetString("telegram.url"))
		telegramIntegration.SetBotToken(viper.GetString("telegram.botToken"))
	}
}

func run(cmd *cobra.Command, args []string) error {
	mrRes, err := project.SubmitMergeRequest(&gitlab.MergeRequestOptions{
		SourceBranch: viper.GetString("project.branches.develop"),
		TargetBranch: viper.GetString("project.branches.master"),
		Title:        "Automatic release of 1.0.0",
	})

	if err != nil {
		return err
	}

	fmt.Printf("Merge request result: %#v\n\n", mrRes)
	// Notify in the group
	err = telegramIntegration.SendMessage(
		viper.GetString("telegram.groupId"),
		mrRes.String(),
	)

	if err != nil {
		return err
	}

	return nil
}
