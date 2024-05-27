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

		start, err := cmd.Flags().GetUint("start")
		if err != nil {
			return err
		}
		end, err := cmd.Flags().GetUint("end")
		if err != nil {
			return err
		}
		fps, err := cmd.Flags().GetInt("fps")
		if err != nil {
			return err
		}
		if fps < 0 {
			fps = int(end) - int(start) + 1
		}

		// todo: parameter validation
		s := sprite.NewSprite(frame, col, row)
		return model.Start(s, int(start), int(end), fps)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().UintP("start", "s", 0, "animation start index")
	rootCmd.Flags().UintP("end", "e", 0, "animation end index")
	rootCmd.Flags().IntP("fps", "f", -1, "frame per second")
}
