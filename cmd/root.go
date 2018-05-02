package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkgpal",
	Short: "Pkgpal tells you what packages in your project do!",
	Long: `When you start a new project, you often don't know what most of the packages
		   used actually do. pkgpal tries to help you figure that out at a glance.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please enter a command")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
