package main

import (
	"fmt"
	"log"
	"os"

	git "github.com/libgit2/git2go/v33"
)

func calculatePrompt(config PromptConfiguration) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get the current directory: %w", err)
	}

	repo, err := git.OpenRepositoryExtended(dir, 0, "")
	if err != nil {
		if git.IsErrorCode(err, git.ErrorCodeNotFound) {
			if promptConfig.Debug {
				log.Println("The current directory is not a git repository")
			}
			return "", nil
		}

		return "", fmt.Errorf("failed to open a repository at %s: %w", dir, err)
	}

	stashCount, err := getStashCount(repo)
	if err != nil {
		return "", fmt.Errorf("failed to read the status of the stash: %w", err)
	}
	branchStatus, err := getBranchStatus(repo)
	if err != nil {
		return "", fmt.Errorf("failed to read information about the branch of the repository: %w", err)
	}
	treeStatus, err := getTreeStatus(repo)
	if err != nil {
		return "", fmt.Errorf("failed to read information about the tree of the repository: %w", err)
	}

	var (
		formatString string
		ahead        string
		behind       string
		stash        string
		staged       string
		unstaged     string
		untracked    string
		conflict     string
		clean        string
	)

	if config.ZshMode {
		formatString = "%%F{white}(git:%s%s%s%%B%%F{red}%s%%b%%f|%%F{cyan}%s%%F{blue}%s%%F{yellow}%s%%F{white}%s%%F{red}%s%%F{green}%s%%f)"
	} else {
		formatString = "\033[37m(git:%s%s%s\033[31;1m%s\033[22;39m|\033[22;36m%s\033[34m%s\033[33m%s\033[37m%s\033[31m%s\033[32m%s\033[39m)"
	}

	if branchStatus.AheadCount > 0 {
		ahead = config.AheadSigil
	}
	if branchStatus.BehindCount > 0 {
		behind = config.BehindSigil
	}
	if stashCount > 0 {
		stash = config.StashedSigil
	}
	if treeStatus.StagedCount > 0 {
		staged = config.StagedSigil
	}
	if treeStatus.UnstagedCount > 0 {
		unstaged = config.UnstagedSigil
	}
	if treeStatus.UntrackedCount > 0 {
		untracked = config.UntrackedSigil
	}
	if treeStatus.ConflictCount > 0 {
		conflict = config.ConflictsSigil
	}
	// (!tstatus.staged_count && !tstatus.unstaged_count && !tstatus.untracked_count && !tstatus.conflict_count) ? options.sigil_clean : ""
	if treeStatus.StagedCount == 0 && treeStatus.UnstagedCount == 0 && treeStatus.UntrackedCount == 0 && treeStatus.ConflictCount == 0 {
		clean = config.CleanSigil
	}

	output := fmt.Sprintf(
		formatString,
		branchStatus.HeadName,
		ahead,
		behind,
		branchStatus.CurrentOperation,
		stash,
		staged,
		unstaged,
		untracked,
		conflict,
		clean,
	)

	return output, nil
}
