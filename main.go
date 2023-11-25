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
// Parse files from a path (options: -o to print into a JSON file)
//
//	files parse /path/to/directory
//
// Parse file information from a filepath (options: -o to print into a JSON file)
//
//	files info /path/to/file
//
// Parse file metadata (options: -j to output as JSON, -o to print into a JSON file)
//
//	files metadata /path/to/file
//
// Parse files from a path with info, output is raw JSON (options: -o to print into a JSON file)
//
//	files periscope /path/to/directory
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ewilan-riviere/scanner/pkg/files"
	"github.com/ewilan-riviere/scanner/pkg/info"
	"github.com/ewilan-riviere/scanner/pkg/mediainfo"
	"github.com/ewilan-riviere/scanner/pkg/periscope"
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
			json, _ := cmd.Flags().GetBool("json")
			path := args[0]

			if isFile(path) {
				fmt.Println("Error: path is file")
				return
			}

			if !json {
				fmt.Println("Parse files from", path)
			}

			if output != "" {
				fmt.Println("output:", output)
			}

			if !json {
				executionTime(func() {
					items := files.Make(path)
					printer.Print(items)

					if output != "" {
						printer.ToJSON(printer.Params{
							Data: items,
							Path: output,
						})
					}
				})
			} else {
				items := files.Make(path)
				fmt.Println(printer.ToJSONString(items))
			}
		},
	}

	var cmdInfo = &cobra.Command{
		Use:   "info [filepath]",
		Short: "Parse file information from a filepath",
		Long:  `Parse file information from a filepath`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			output, _ := cmd.Flags().GetString("output")
			json, _ := cmd.Flags().GetBool("json")
			path := args[0]

			if !isFile(path) {
				fmt.Println("Error: path is not a file")
				return
			}

			if !json {
				fmt.Println("Parse info from", path)
			}
			if output != "" {
				fmt.Println("output:", output)
			}

			if !json {
				executionTime(func() {
					file := info.Make(path)
					printer.Print(file)

					if output != "" {
						printer.ToJSON(printer.Params{
							Data: file,
							Path: output,
						})
					}
				})
			} else {
				file := info.Make(path)
				fmt.Println(printer.ToJSONString(file))
			}
		},
	}

	var cmdMetadata = &cobra.Command{
		Use:   "metadata [filepath]",
		Short: "Parse metadata of file from a filepath",
		Long:  `Parse metadata of file from a filepath`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			output, _ := cmd.Flags().GetString("output")
			json, _ := cmd.Flags().GetBool("json")
			path := args[0]

			if !isFile(path) {
				fmt.Println("Error: path is not a file")
				return
			}

			if !json {
				fmt.Println("Parse metadata of file", path)
			}
			if output != "" {
				fmt.Println("output:", output)
			}

			metadata := mediainfo.Make(path)
			if !json {
				printer.Print(metadata)
			}

			if json {
				fmt.Println(printer.ToJSONString(metadata))
			}

			if output != "" {
				printer.ToJSON(printer.Params{
					Data: metadata,
					Path: output,
				})
			}
		},
	}

	var cmdPeriscope = &cobra.Command{
		Use:   "periscope [path]",
		Short: "Parse files from a path with info",
		Long:  `Parse files from a path with info`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			output, _ := cmd.Flags().GetString("output")
			path := args[0]

			if isFile(path) {
				fmt.Println("Error: path is file")
				return
			}

			if output != "" {
				fmt.Println("output:", output)
			}

			periscope := periscope.Make(path)
			fmt.Println(printer.ToJSONString(periscope))

			if output != "" {
				printer.ToJSON(printer.Params{
					Data: periscope,
					Path: output,
				})
			}
		},
	}

	cmdParse.Flags().StringP("output", "o", "", "Print files into JSON output file")
	cmdParse.Flags().BoolP("json", "j", false, "Output as JSON")

	cmdInfo.Flags().StringP("output", "o", "", "Print files into JSON output file")
	cmdInfo.Flags().BoolP("json", "j", false, "Output as JSON")

	cmdMetadata.Flags().StringP("output", "o", "", "Print files into JSON output file")
	cmdMetadata.Flags().BoolP("json", "j", false, "Output as JSON")

	cmdPeriscope.Flags().StringP("output", "o", "", "Print files into JSON output file")

	var rootCmd = &cobra.Command{Use: "scanner"}
	rootCmd.AddCommand(cmdParse)
	rootCmd.AddCommand(cmdInfo)
	rootCmd.AddCommand(cmdMetadata)
	rootCmd.AddCommand(cmdPeriscope)
	rootCmd.Version = "0.0.4"
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
