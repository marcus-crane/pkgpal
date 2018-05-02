package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/marcus-crane/pkgpal/search"
	"github.com/spf13/cobra"
)

var language string

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "Language of the package you're searching (required)")
	rootCmd.MarkFlagRequired("language")
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a package",
	Long:  `Enter a package name accompanied by a language flag`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch language {
			case "node":
				req := search.NpmPackage(args[0])
				color.Green("↳  " + req.Name + " => " + req.Description)
			case "python":
				req := search.PypiPackage(args[0])
				color.Blue("↳  " + req.Name + " => " + req.Description)
			default:
				fmt.Println("Please enter a valid language")
				os.Exit(1)
			}
		} else {
			fmt.Println("Please enter a package to search")
			os.Exit(1)
		}
	},
}
