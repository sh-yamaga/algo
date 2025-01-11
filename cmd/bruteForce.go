/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// bruteForceCmd represents the bruteForce command
var bruteForceCmd = &cobra.Command{
	Use:     "bruteForce",
	Aliases: []string{"bf"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bruteForce called")
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(bruteForceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bruteForceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bruteForceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
