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
	projectList         gitlab.IProject
)

func init() {
	cobra.OnInitialize(configMergeRequest)

	rootCmd.AddCommand(submitMergeRequest)
}

func configMergeRequest() {
	parseConfig(submitMergeRequest, submitMergeRequest.ArgAliases)

	if gitlabIntegration == nil {
		gitlabIntegration = gitlab.New(
			globalConfig.Gitlab.Url,
			globalConfig.User.PersonalToken,
		)

		for _, proj := range globalConfig.Projects {
			proj.SetGitlab(gitlabIntegration)
		}
	}

	if telegramIntegration == nil {
		telegramIntegration = t.New(viper.GetString("telegram.url"))
		telegramIntegration.SetBotToken(viper.GetString("telegram.botToken"))
	}
}

func notify(res gitlab.Response, err error) {
	if err != nil {
		return
	}

	if err = telegramIntegration.SendMessage(
		viper.GetString("telegram.groupId"),
		res.String(),
	); err != nil {
		fmt.Printf("Failed to send notification. %#v\n", err)
	}
	return
}

func run(cmd *cobra.Command, args []string) error {
	for _, project := range globalConfig.Projects {
		mrRes, err := project.SubmitMergeRequest(&gitlab.MergeRequestOptions{
			SourceBranch: project.Branches.Develop,
			TargetBranch: project.Branches.Master,
			Title:        "Automatic release of 1.0.0",
		})
		defer notify(mrRes, err)

		if err != nil {
			return err
		}
	}

	return nil
}
