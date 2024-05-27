package cmd

import (
	"os"
	"strconv"

	"github.com/semihbkgr/sprite-animator-cli/model"
	"github.com/semihbkgr/sprite-animator-cli/sprite"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sprite-animator-cli <file> <col> <row>",
	Short: "animate the sprites in terminal",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		frame, err := sprite.LoadPNG(args[0])
		if err != nil {
			return err
		}

		col, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		row, err := strconv.Atoi(args[2])
		if err != nil {
			return err
		}

		s := sprite.NewSprite(frame, col, row)
		return model.Start(s)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
