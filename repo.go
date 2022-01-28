package main

import (
	git "github.com/libgit2/git2go/v33"
)

func getStashCount(repo *git.Repository) (int, error) {
	stashCount := 0
	_ = repo.Stashes.Foreach(func(index int, message string, id *git.Oid) error {
		stashCount++
		return nil
	})

	return stashCount, nil
}

func getBranchStatus(repo *git.Repository) (*BranchStatus, error) {
	status := &BranchStatus{}

	// Get current operation
	switch repo.State() {
	case git.RepositoryStateNone:
	case git.RepositoryStateMerge:
		status.CurrentOperation = "[merge]"
	case git.RepositoryStateRevert:
		status.CurrentOperation = "[revert]"
	case git.RepositoryStateCherrypick:
		status.CurrentOperation = "[cherry-pick]"
	case git.RepositoryStateBisect:
		status.CurrentOperation = "[bisect]"
	case git.RepositoryStateRebase:
		status.CurrentOperation = "[rebase]"
	case git.RepositoryStateRebaseInteractive:
		status.CurrentOperation = "[rebase-i]"
	case git.RepositoryStateRebaseMerge:
		status.CurrentOperation = "[rebase-merge]"
	case git.RepositoryStateApplyMailbox:
		status.CurrentOperation = "[mailbox]"
	case git.RepositoryStateApplyMailboxOrRebase:
		status.CurrentOperation = "[mailbox-rebase]"
	}

	// Get repository head
	head, err := repo.Head()
	if err != nil {
		// The repository HEAD does not exist, maybe the repo is empty and has no commit yet
		status.HeadName = "<none>"
		return status, nil
	}
	if head.IsBranch() {
		status.HeadName = head.Shorthand()
		upstream, err := head.Branch().Upstream()
		if err == nil {
			ahead, behind, err := repo.AheadBehind(head.Target(), upstream.Target())
			if err == nil {
				if ahead > 0 {
					status.AheadCount = ahead
				}
				if behind > 0 {
					status.BehindCount = behind
				}
			}
		}
	} else if head.IsTag() {
		status.HeadName = head.Shorthand()
	} else {
		// TODO: check if this works, or is short enough
		commit, err := head.Peel(git.ObjectCommit)
		if err == nil {
			status.HeadName, _ = commit.ShortId()
			// TODO handle this error
		}
	}

	return status, nil
}

func getTreeStatus(repo *git.Repository) (*TreeStatus, error) {
	status := &TreeStatus{}

	options := &git.StatusOptions{
		Show:     git.StatusShowIndexAndWorkdir,
		Flags:    git.StatusOptIncludeUntracked & git.StatusOptRecurseUntrackedDirs,
		Pathspec: nil,
	}
	statusList, _ := repo.StatusList(options)
	statusCount, _ := statusList.EntryCount()
	for i := 0; i < statusCount; i++ {
		statusEntry, err := statusList.ByIndex(i)
		if err == nil {
			switch {
			case statusEntry.Status&git.StatusCurrent != 0:
				continue
			case statusEntry.Status&git.StatusWtNew != 0:
				status.UntrackedCount++
			case statusEntry.Status&(git.StatusWtDeleted|git.StatusWtModified|git.StatusWtRenamed|git.StatusWtTypeChange) != 0:
				status.UnstagedCount++
			case statusEntry.Status&(git.StatusIndexNew|git.StatusIndexModified|git.StatusIndexDeleted|git.StatusIndexRenamed|git.StatusIndexTypeChange) != 0:
				status.StagedCount++
			case statusEntry.Status&(git.StatusConflicted) != 0:
				status.ConflictCount++
			}
		}
	}

	return status, nil
}
