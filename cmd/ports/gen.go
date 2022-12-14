/*
Copyright © 2022 @ffelipelimao
*/
package ports

import (
	"github.com/ffelipelimao/gobuilder/cmd/gen"
	"github.com/spf13/cobra"
)

var name string
var fields string
var destination string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen can create your mockbuilder files",
	Long: "gen create your mockbuilder file.\n" +
		"You can pass -n to create with te name of your struct and -f with fields and types.\n" +
		"If you have any doubt see README.md to see examples",

	Run: func(cmd *cobra.Command, args []string) {
		gen.Start(name, fields, destination)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&name, "name", "n", "", "mockbuilder name")
	genCmd.Flags().StringVarP(&fields, "fields", "f", "", "fields separated by comma with types separated by -")
	genCmd.Flags().StringVarP(&destination, "destination", "d", "", "the destination folder")
}
