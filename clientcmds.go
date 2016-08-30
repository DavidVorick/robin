package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// robinSetupHelpString returns helpful information about the robin setup commands.
func robinSetupHelpString() string {
	return fmt.Sprintf(`Perform setup on a robin client or server.

To see more information about a command, run 'robin setup [command] --help'`)
}

// robinSetupString returns helpful information about the robin setup commands.
func robinSetupString() string {
	return fmt.Sprintf(`Perform setup on a robin client or server.

Available options:
	robin setup init-client
	robin setup add-remote    [address:port]
	robin setup remove-remote [address:port]

	robin setup init-server [port]
	robin setup add-key    	[key]
	robin setup revoke-key 	[key]

To see more information about a command, run 'robin setup [command] --help'`)
}

// robinSetupCmd returns helpful information about the robin setup commands.
func robinSetupCmd(cmd *cobra.Command, args []string) {
	fmt.Println(robinSetupString())
}

// robinSetupInitclientHelpString returns helpful information about
// initializing a robin client.
func robinSetupInitclientHelpString() string {
	return fmt.Sprint("" +
		`Initialize a robin client. A key will be created for the client which must be
given to all remote servers so that the client can make changes to the
directory. Local databases will also be established for the client.`)
}

// robinSetupInitclientCmd will initialize a robin client.
func robinSetupInitclientCmd(cmd *cobra.Command, args []string) {
	fmt.Println("client initialization not yet implemented")
}

// robinSetupAddremoteHelpString() returns helpful information about adding a
// remote server to a robin client.
func robinSetupAddremoteHelpString() string {
	return fmt.Sprint("" +
		`Add a remote server to the client. The client will contact the server and pull
changes every time the 'robin sync' command is issued. The client will also
upload any changes that the server is missing. The client will also upload new
changes to the server every time the 'robin commit' command is called.`)
}

// robinSetupAddremoteCmd adds a remote server to the robin client.
func robinSetupAddremoteCmd(cmd *cobra.Command, args []string) {
	fmt.Println("adding a remote to a client is not yet implemented")
}

// robinSetupRemoveremoteHelpString() returns helpful information about
// removing a remote server from the robin client.
func robinSetupRemoveremoteHelpString() string {
	return fmt.Sprint("" +
		`Remove a remote server from the client. The client will no longer talk to this
server when synchronizing the directory and committing changes.`)
}

func robinSetupRemoveremoteCmd(cmd *cobra.Command, args []string) {
	fmt.Println("removing a remote from a client is not yet implemented")
}
