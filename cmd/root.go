package cmd

import (
	"fmt"
	"os"
	"xiongba/shell-ai/config"

	"github.com/spf13/cobra"
)

var (
	Conf       *config.Config
	ConfigPath string
)
var rootCmd = &cobra.Command{
	Short: "AI 命令行工具",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(aiCmd)

}
