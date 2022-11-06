/*
Copyright © 2022 @ffelipelimao
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobuilder",
	Short: "gobuilder is a file generator",
	Long: `gobuilder is a CLI library for Go that empowers applications.
This application is a tool to generate files quickly.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {

		os.Exit(1)
	}
}
