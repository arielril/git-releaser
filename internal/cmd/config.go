package cmd

import (
	"fmt"
	"git-releaser/internal/integrations/gitlab"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Gitlab struct {
		Url string `yaml:"url"`
	} `yaml:"gitlab"`
	User struct {
		PersonalToken string `yaml:"personalToken"`
	} `yaml:"user"`
	Projects []*gitlab.Project `yaml:"projects"`
	Telegram struct {
		Url      string `yaml:"url"`
		BotToken string `yaml:"botToken"`
		GroupId  int    `yaml:"groupId"`
	} `yaml:"telegram"`
}

var showConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Shows the parsed config file",
	RunE:  parseConfig,
}

var globalConfig *Config

func parseConfig(cmd *cobra.Command, args []string) error {
	if globalConfig != nil {
		return nil
	}

	err := viper.Unmarshal(&globalConfig)

	if err != nil {
		return fmt.Errorf("Failed to parse config: %v", err)
	}
	fmt.Printf("Config !@#@!\n")
	fmt.Printf("  Gitlab: %#v\n", globalConfig.Gitlab)
	fmt.Printf("  User: %#v\n", globalConfig.User)
	fmt.Printf("  Projects: %#v\n", globalConfig.Projects)
	fmt.Printf("  Telegram: %#v\n", globalConfig.Telegram)
	return nil
}

func init() {
	rootCmd.AddCommand(showConfigCmd)
}
