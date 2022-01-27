package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ExecutableName string
	Version        string

	promptConfig PromptConfiguration

	rootCmd *cobra.Command
)

func init() {
	rootCmd = &cobra.Command{
		Use:   ExecutableName,
		Short: fmt.Sprintf("%s is a tool to watch updates on a registry and customize the update using hooks", ExecutableName),
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version and exit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s: %s\n", ExecutableName, Version)
		},
	}
	rootCmd.AddCommand(versionCmd)

	shellPromptCmd := &cobra.Command{
		Use:   "shell-prompt",
		Short: "Print the shell prompt content and exit",
		Run: func(cmd *cobra.Command, args []string) {
			prompt, err := calculatePrompt(promptConfig)
			if err != nil {
				fmt.Printf("An error occured: %v\n", err)
				os.Exit(1)
				return
			}

			fmt.Printf("%s\n", prompt)
		},
	}

	shellPromptCmd.Flags().StringVarP(&promptConfig.AheadSigil, "ahead-sigil", "a", "↑", "Sigil to signal the branch is ahead of the remote")
	shellPromptCmd.Flags().StringVarP(&promptConfig.BehindSigil, "behind-sigil", "b", "↓", "Sigil to signal the branch is behind of the remote")
	shellPromptCmd.Flags().StringVarP(&promptConfig.StagedSigil, "staged-sigil", "s", "●", "Sigil to signal there are staged edits")
	shellPromptCmd.Flags().StringVarP(&promptConfig.ConflictsSigil, "conflicts-sigil", "c", "✖", "Sigil to signal there are conflicts to resolve")
	shellPromptCmd.Flags().StringVarP(&promptConfig.UnstagedSigil, "unstaged-sigil", "u", "✚", "Sigil to signal there are unstaged edits")
	shellPromptCmd.Flags().StringVarP(&promptConfig.UntrackedSigil, "untracked-sigil", "U", "…", "Sigil to signal there are untracked files")
	shellPromptCmd.Flags().StringVarP(&promptConfig.StashedSigil, "stashed-sigil", "S", "⚑", "Sigil to signal there are stashed edits")
	shellPromptCmd.Flags().StringVarP(&promptConfig.CleanSigil, "clean-sigil", "C", "✔", "Sigil to signal the working tree is clean")
	shellPromptCmd.Flags().BoolVar(&promptConfig.ZshMode, "zsh-mode", false, "Print the output using color tags in the zsh standard instead of AINSI")
	rootCmd.AddCommand(shellPromptCmd)
}

func main() {
	rootCmd.Execute()
}
