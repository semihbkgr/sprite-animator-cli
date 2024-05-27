package cmd

import (
	"os"

	"github.com/semihbkgr/sprite-animator-cli/model"
	"github.com/semihbkgr/sprite-animator-cli/sprite"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sprite-animator-cli",
	Short: "animate the sprites in terminal",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return model.Start(sprite.TestSprite)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
