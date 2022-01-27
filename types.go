package main

type PromptConfiguration struct {
	AheadSigil     string
	BehindSigil    string
	StagedSigil    string
	ConflictsSigil string
	UnstagedSigil  string
	UntrackedSigil string
	StashedSigil   string
	CleanSigil     string
	ZshMode        bool
}

type BranchStatus struct {
	HeadName         string
	AheadCount       int
	BehindCount      int
	CurrentOperation string
}

type TreeStatus struct {
	StagedCount    int
	UnstagedCount  int
	UntrackedCount int
	ConflictCount  int
}
