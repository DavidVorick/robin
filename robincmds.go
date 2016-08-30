package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// robinHelpString returns helpful information about the robin program.
func robinHelpString() string {
	return fmt.Sprintf(`Robin v%v		

Robin is a backup and synchronization tool for personal data. Data can be
stored on untrusted remote servers such that the remote server never learns
directory structure, file names, file contents, or file sizes.`, Version)
}

// robinString returns helpful information about the robin program.
func robinString() string {
	return fmt.Sprintf(`Robin v%v		

Robin is a backup and synchronization tool for personal data. Data can be
stored on untrusted remote servers such that the remote server never learns
directory structure, file names, file contents, or file sizes.

Available Commands:
	robin sync
	robin commit
	robin setup

To see more information about a command, run 'robin [command] --help'`, Version)

}

// robinCmd returns helpful information about the robin program.
func robinCmd(cmd *cobra.Command, args []string) {
	fmt.Println(robinString())
}
