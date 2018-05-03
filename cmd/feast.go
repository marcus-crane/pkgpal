package cmd

import (
	"fmt"

	"github.com/marcus-crane/pkgpal/parsers"
	"github.com/marcus-crane/pkgpal/search"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(feastCmd)
}

var feastCmd = &cobra.Command{
	Use:   "feast",
	Short: "Detect package manager files",
	Long:  `Automatically find package manager files and see what their dependencies do`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd := parsers.GetCurrentDirectory()
		file := parsers.DetectPackageFile(cwd)
		if file != "" {
			switch file {
			case "package.json":
				packages := search.ParseNpmPackage(cwd + "/" + file)
				for _, dependency := range packages {
					pkg := search.NpmPackage(dependency)
					fmt.Println(pkg.Name + " => " + pkg.Description)
				}
			case "requirements.txt":
				packages := search.ParseRequirements(cwd + "/" + file)
				fmt.Println(packages)
				for _, dependency := range packages {
					pkg := search.PypiPackage(dependency)
					fmt.Println(pkg.Name + " => " + pkg.Description)
				}
			default:
				return
			}
		}
	},
}
