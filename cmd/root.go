package cmd

import (
	"fmt"
	"os"

	"github.com/cbess/go-textwrap"
	"github.com/graytonio/meet-summary/internal/summarizer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "zoltan-bot",
	Short: "A discord bot for generating AI DND resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath, err := cmd.PersistentFlags().GetString("transcript")
		if err != nil {
			return err
		}

		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		chunks, err := textwrap.WordWrap(string(data), 3500, -1)
		if err != nil {
			return err
		}

		var fullSummary string

		for _, chunk := range chunks.TextGroups {
			summary, err := summarizer.SummarizeMeeting(chunk, viper.GetString("OPENAI_API_KEY"))
			if err != nil {
				return err
			}
			fullSummary += "\n" + summary
		}

		

		fmt.Println(fullSummary)
		return nil
	},
}

var transcriptPath string

func init() {
	cobra.OnInitialize(loadConfig)
	rootCmd.PersistentFlags().StringVarP(&transcriptPath, "transcript", "t", "./transcript.txt", "Transcript file for zoom meeting")
}

func loadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("gpt")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Error loading config file: %v", err)
	}
}

func Execute() error {
	return rootCmd.Execute()
}