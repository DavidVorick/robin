package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// robinSyncHelpString returns helpful information about synchronizing a
// client.
func robinSyncHelpString() string {
	return fmt.Sprint("" +
		`Synchronize a client with all of the remote servers. Uncommitted changes will
be ignored. All changes are pulled from the remote servers into the client
directory. If the servers are not synchronized, the client will synchronize
them. Conflicts are handled such that there is no risk of data loss.`)
}

// robinSyncCmd will synchronize a client to all of the remote servers, and
// synchronize all of the remote servers to eachother.
func robinSyncCmd(cmd *cobra.Command, args []string) {
	fmt.Println("syncing is not implemented")
}

// robinCommitExampleUsageString returns a comprehensive example for using the
// 'robin commit' command.
func robinCommitExampleUsageString() string {
	return fmt.Sprint(`
Usage Examples:
	robin commit file1.jpg file2.jpg -m "pictures of the broken tiles"
	robin commit folder1 -m "add backup linux configuration files"
	robin commit . -m "some directory restructuring"`)
}

// robinCommitHelpString returns helpful information about committing files
// from a client.
func robinCommitHelpString() string {
	return fmt.Sprint("" +
		`Commit a file or folder to the directory. Upon commit, the change will be
synchronized to all of the remote servers. All commits must have a message.
All files will be padded to the nearest blocksize, where the block size is a
power of two and at least 5% of the filesize. Groups of files will be batched
and uploaded as a set, such that the server only sees a single file.
` + robinCommitExampleUsageString())
}

// robinCommitCmd will commit a change to the client and synchronize the change
// to all of the remote servers.
func robinCommitCmd(cmd *cobra.Command, args []string) {
	fmt.Println("committing is not yet implemented")
}
