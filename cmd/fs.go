/*
Copyright © 2026 Jordan BOUTROIS
*/
package cmd

import (
	"baguette/internal/filesearch"
	"flag"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// fsCmd represents the fs command
var fsCmd = &cobra.Command{
	Use:        "fs",
	Short:      "Search string in file",
	Long:       `Search string in file`,
	flag.Usage: "baguette fs -f \"./file.txt\" -s \"hello world\"",
	Run: func(cmd *cobra.Command, args []string) {

		filepath, err := cmd.Flags().GetString("filepath")
		if err != nil {
			log.Fatal("Error while getting args: ", err)
		}
		search, err := cmd.Flags().GetString("search")
		if err != nil {
			log.Fatal("Error while getting args: ", err)
		}

		occurrences, err := filesearch.Search(filepath, search)
		if err != nil {
			log.Fatal("Error while searching files: ", err)
		}

		for _, occurrence := range occurrences {
			fmt.Printf("%s", occurrence)
		}
	},
}

func init() {
	rootCmd.AddCommand(fsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	fsCmd.Flags().StringP("filepath", "f", "", "File to search")
	fsCmd.Flags().StringP("search", "s", "", "string to search")

	if err := fsCmd.MarkFlagRequired("filepath"); err != nil {
		log.Fatal(err)
	}

	if err := fsCmd.MarkFlagRequired("search"); err != nil {
		log.Fatal(err)
	}

}

//func checkArgs(cmd *cobra.Command, args []string) error {
//
//	for _, arg := range args {
//		_, err := cmd.Flags().GetStrings(arg)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//
//	return fmt.Errorf("Invalid arguments")
//}
