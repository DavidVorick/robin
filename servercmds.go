package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// robinSetupInitserverString returns helpful information about initializing a
// robin server.
func robinSetupInitserverHelpString() string {
	return fmt.Sprint("" +
		`Initialize a robin server. The robin server can safely run on untrusted
hardware. If no server yet exists in the directory, an empty logfile will be
created, a database will be created, and each will be populated the first time
that an authorized client connects to the server.`)
}

// robinSetupInitserverCmd will initialize and begin running a robin server.
func robinSetupInitserverCmd(cmd *cobra.Command, args []string) {
	err = 
}

// robinSetupAddkeyHelpString() returns helpful information about adding a
// client key to a server.
func robinSetupAddkeyHelpString() string {
	return fmt.Sprint("" +
		`Add a client key to the server. The server will now accept sync and commit
requests from that client.`)
}

// robinSetupAddkeyCmd will add a client key to the robin server.
func robinSetupAddkeyCmd(cmd *cobra.Command, args []string) {
	fmt.Println("add-key is not yet implemented")
}

// robinSetupRevokekeyHelpString returns helpful information about revoking a
// client key from a robin server.
func robinSetupRevokekeyHelpString() string {
	return fmt.Sprint("" +
		`Revoke a client key from the server. The server will no longer accept sync and
commit requests from that client.`)
}

// robinSetupRevokekeyCmd will revoke a client key from the server.
func robinSetupRevokekeyCmd(cmd *cobra.Command, args []string) {
	fmt.Println("revoke-key is not yet implemented")
}
