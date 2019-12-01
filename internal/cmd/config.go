package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Shows the parsed config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		var config map[string]interface{}
		err := viper.Unmarshal(&config)

		if err != nil {
			return fmt.Errorf("Failed to parse config: %v", err)
		}

		fmt.Printf("Configs %#v\n\n", config)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showConfigCmd)
}
