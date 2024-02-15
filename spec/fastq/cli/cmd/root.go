/*
Copyright © 2024 utubun <utubun@icloud.com>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/utubun/ngs/spec/fastq"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qsc",
	Short: "Quality score calls for fastq data",
	Long:  `Quality score blah TODO:Add documentation`,

	Run: func(cmd *cobra.Command, args []string) {
		countFlag, _ := cmd.Flags().GetBool("count")
		pathFlag, _ := cmd.Flags().GetString("path")

		if countFlag {
			fmt.Println("Count flag found")
		} else {
			fmt.Println("Not counting on you")
		}

		if pathFlag != "" {
			fmt.Printf("Reading the file: %s\n", pathFlag)
			start := time.Now()
			r, err := fastq.ReadLocalFile(pathFlag)
			t := time.Now()
			elapsed := t.Sub(start)
			if err != nil {
				fmt.Printf("Error: %+s", err)
				return
			}
			fmt.Printf("Time elapsed: %+v\n", elapsed)

			/*fmt.Println("Extracting archive")
			start = time.Now()
			gz, err := fastq.ReadGz(r)
			t = time.Now()
			elapsed = t.Sub(start)
			if err != nil {
				fmt.Printf("Unable to read archive: %s", err)
				return
			}
			fmt.Printf("Read gz archive in %+v\n", elapsed)*/

			fmt.Println("Reading fastq file")
			start = time.Now()
			res, _ := fastq.Read(r)
			t = time.Now()
			elapsed = t.Sub(start)
			fmt.Printf("First read: \n%#v\nQuality scores:%#v\nDone in %+v\n", res.Sequence[11], res.QScores[11], elapsed)

			start = time.Now()
			count := len(res.Sequence)
			t = time.Now()
			elapsed = t.Sub(start)
			fmt.Printf("Quality control count: %#v\nDone in %+v\n", count, elapsed)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("count", "c", false, "returns the count of reads in fastq file")
	rootCmd.Flags().StringP("path", "p", ".", "Path to the fastq file")
}
