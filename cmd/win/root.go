package win

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wingoutil",
	Short: "DeveloperDurp's Windows Utility - Install Programs, Tweaks, Fixes, and Updates",
	Long:  `DeveloperDurp's Windows Utility - Install Programs, Tweaks, Fixes, and Updates (Inspired by Chris Titus Tech's winutil)`,
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
