/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"baguette/internal/config"
	"baguette/internal/loadbench"
	"fmt"

	"github.com/spf13/cobra"
)

// loadbenchCmd represents the loadbench command
var loadbenchCmd = &cobra.Command{
	Use:   "load",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("loadbench load")
		loadbench.Load()
	},
}

var wizard = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.ReadJson("./data/config.json")
		if err != nil {
			return
		}

		//loadbench.InitializeConfig()
	},
}

func init() {
	rootCmd.AddCommand(loadbenchCmd)
	rootCmd.AddCommand(wizard)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadbenchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadbenchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
