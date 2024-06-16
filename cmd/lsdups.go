package cmd

import (
	"github.com/rajeshamdev/lsdups/lsdups"
	"github.com/spf13/cobra"
)

// lsdupsCmd represents the lsdups command
var ls = &cobra.Command{
	Use:   "ls",
	Short: "List duplicate files",
	Long:  "Iterates recursively through a dir finding checksum of files and list duplicates.",
	Run: func(cmd *cobra.Command, args []string) {
		lsdups.Lsdups()
	},
}

func init() {
	rootCmd.AddCommand(ls)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsdupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsdupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
