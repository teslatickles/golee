package main

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var FullPullCmd cli.Command = cli.Command{
	Name:        "fullpull",
	Aliases:     []string{"fp"},
	Usage:       "safely pull down all changes for all projects in a given folder",
	UsageText:   "gg fp",
	Description: "automates pull on all repos in directory",
	ArgsUsage:   "",
	Category:    "",
	BashComplete: func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "--better\n")
	},
	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
	Action:                 Fullpull,
	Subcommands:            []*cli.Command{},
	Flags:                  []cli.Flag{},
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}

var AddCpCmd cli.Command = cli.Command{
	Name:        "addCommitPush",
	Aliases:     []string{"acp"},
	Usage:       "add -> commit -> push in one command",
	UsageText:   "gg acp",
	Description: "automates remote push of changes in current branch",
	ArgsUsage:   "",
	Category:    "",
	BashComplete: func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "--better\n")
	},
	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
	Action:                 AddCommitPush,
	Subcommands:            []*cli.Command{},
	Flags:                  []cli.Flag{},
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}

var StashPpCmd cli.Command = cli.Command{
	Name:        "stashPullPop",
	Aliases:     []string{"spp"},
	Usage:       "add -> commit -> push in one command",
	UsageText:   "gg spp",
	Description: "automates remote push of changes in current branch",
	ArgsUsage:   "",
	Category:    "",
	BashComplete: func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "--better\n")
	},
	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
	Action:                 StashPullPop,
	Subcommands:            []*cli.Command{},
	Flags:                  []cli.Flag{},
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}

var SoftResetCmd cli.Command = cli.Command{
	Name:        "softReset",
	Aliases:     []string{"sr", "soft"},
	Usage:       "git soft reset commit in one command",
	UsageText:   "gg soft",
	Description: "uncommits most recent commit, keeps changes intact",
	ArgsUsage:   "",
	Category:    "",
	BashComplete: func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "--better\n")
	},
	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
	Action:                 SoftReset,
	Subcommands:            []*cli.Command{},
	Flags:                  []cli.Flag{},
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}

var HardResetCmd = cli.Command{
	Name:        "hardReset",
	Aliases:     []string{"hr", "hard"},
	Usage:       "git hard reset commit in one command",
	UsageText:   "gg hard",
	Description: "uncommits most recent commit, destroys changes from commit as well!",
	ArgsUsage:   "",
	Category:    "",
	BashComplete: func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "--better\n")
	},
	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
	Action:                 HardReset,
	Subcommands:            []*cli.Command{},
	Flags:                  []cli.Flag{},
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}

var UndoMergeCmd = cli.Command{
	Name:        "undoMerge",
	Aliases:     []string{"um", "undomerge"},
	Usage:       "undo last merge in one command",
	UsageText:   "gg um",
	Description: "reverts master branch to state before selected/last merge",
	ArgsUsage:   "",
	Category:    "",
	BashComplete: func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "--better\n")
	},
	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
	Action:                 UndoMerge,
	Subcommands:            []*cli.Command{},
	Flags:                  []cli.Flag{},
	SkipFlagParsing:        false,
	HideHelp:               false,
	HideHelpCommand:        false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}