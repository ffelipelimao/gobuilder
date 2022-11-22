/*
Copyright © 2022 @ffelipelimao
*/
package ports

import (
	"os"

	"github.com/spf13/cobra"
)

const version = "0.0.8"

var rootCmd = &cobra.Command{
	Use:     "gobuilder",
	Short:   "gobuilder is a file generator",
	Version: version,
	Long: `gobuilder is a CLI library for Go that empowers applications.
This application is a tool to generate files quickly.`,
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
