// Package scanner is a CLI to parse files for a list or to get metadata.
//
// Examples/readme can be found on the GitHub page at https://github.com/ewilan-riviere/scanner
//
// If you want to use it as CLI, you can install it with:
//
//	go install github.com/ewilan-riviere/scanner
//
// All commands have output option to print the result into a JSON file.
//
// Parse files from a path.
//
//	files parse /path/to/directory
//
// Parse file metadata
//
//	files metadata /path/to/file
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ewilan-riviere/scanner/pkg/files"
	"github.com/ewilan-riviere/scanner/pkg/mediainfo"
	"github.com/ewilan-riviere/scanner/pkg/printer"
	"github.com/spf13/cobra"
)

func main() {
	var cmdParse = &cobra.Command{
		Use:   "parse [path]",
		Short: "Parse files from a path",
		Long:  `Parse files from a path`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			output, _ := cmd.Flags().GetString("output")
			path := args[0]

			fmt.Println("Parse files from", path)
			if output != "" {
				fmt.Println("output:", output)
			}

			executionTime(func() {
				items := files.Make(path)
				printer.Print(items)

				if output != "" {
					printer.ToJSON(printer.Params{
						V:    items,
						Path: output,
					})
				}
			})
		},
	}

	var cmdMetadata = &cobra.Command{
		Use:   "metadata [path]",
		Short: "Parse metadata of file from a path",
		Long:  `Parse metadata of file from a path`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			output, _ := cmd.Flags().GetString("output")
			path := args[0]

			if !isFile(path) {
				fmt.Println("Error: path is not a file")
				return
			}

			fmt.Println("Parse metadata of file", path)
			if output != "" {
				fmt.Println("output:", output)
			}

			metadata := mediainfo.Make(path)
			printer.Print(metadata)

			printer.ToJSON(printer.Params{
				V:    metadata,
				Path: output,
			})
		},
	}

	cmdParse.Flags().StringP("output", "o", "", "Print files into JSON output file")
	cmdMetadata.Flags().StringP("output", "o", "", "Print files into JSON output file")

	var rootCmd = &cobra.Command{Use: "files"}
	rootCmd.AddCommand(cmdParse)
	rootCmd.AddCommand(cmdMetadata)
	rootCmd.Execute()
}

type Callback func()

func executionTime(callback Callback) {
	startTime := time.Now()
	callback()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	return !fileInfo.IsDir()
}
