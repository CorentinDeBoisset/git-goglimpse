# Git GoGlimpse

`git-glimpse` is a command-line tool that is aimed at generating a git prompt like the one from [zsh-vcs-prompt](https://github.com/yonchu/zsh-vcs-prompt).

The particularity of this tool is that it is aimed at maximizing the performances using the native Go interface for git (vs. Python that can slow down a terminal)

## Requirements

* libgit2-dev
* make
* golang

## Installation

You can build the program by running:

    make all

Then install it with

    make install


## Configure for ZSH

You can add a right prompt in ZSH, by using the following line in your `.zshrc`:

    RPROMPT='$(git-glimpse shell-prompt --zsh-mode)'

The exact path to git-glimpse can be found using `which git-glimpse`


## Using the CLI

You can configure the output of this tool with the following arguments:

    Print the shell prompt content and exit

    Usage:
    git-glimpse shell-prompt [flags]

    Flags:
    -a, --ahead-sigil string       Sigil to signal the branch is ahead of the remote (default "↑")
    -b, --behind-sigil string      Sigil to signal the branch is behind of the remote (default "↓")
    -C, --clean-sigil string       Sigil to signal the working tree is clean (default "✔")
    -c, --conflicts-sigil string   Sigil to signal there are conflicts to resolve (default "✖")
    -h, --help                     help for shell-prompt
    -s, --staged-sigil string      Sigil to signal there are staged edits (default "●")
    -S, --stashed-sigil string     Sigil to signal there are stashed edits (default "⚑")
    -u, --unstaged-sigil string    Sigil to signal there are unstaged edits (default "✚")
    -U, --untracked-sigil string   Sigil to signal there are untracked files (default "…")
        --zsh-mode                 Print the output using color tags in the zsh standard instead of AINSI

## Uninstalling git-glimpse

TODO
