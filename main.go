package main

// SETUP FLOW:
//
// 1. Init robin server
// 2. Init robin client
// 3. Configure server to recognize client
// 4. Configure client to recognize server
// 5. (add additional servers and clients as needed)
// 6. Commit a file from one client.
// 7. Get the file by syncing the other clients.

// TODO: If a server rejects a client's sync request, print the client key to
// the terminal so that the user can easily add the client key to the server.

// TODO: If a server cannot be reached, 'robin setup init-server --help' for
// information about setting up a remote server.

// TODO: The help command should more fully explain what happens, and what the
// user needs to do, including explaining port-forwarding.

// TODO: Provide examples if the command is mis-used.

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	// The version of the program. There are no guarantees on the security,
	// compatibilty, etc. until version 1.0.0 is announced.
	Version = "0.0.1"
)

var (
	// Flags that can be set. These flags are only active for certain commands,
	// though they are created as global flags.
	directory string
	message   string
)

func main() {
	// Establish the root command.
	robin := &cobra.Command{
		Use:   os.Args[0],
		Short: "Robin - secure data synchronization and backup.",
		Long:  robinHelpString(),
		Run:   robinCmd,
	}

	// Establish the sync command, a child of the root command.
	robinSync := &cobra.Command{
		Use:   "sync",
		Short: "Synchronize the directory to all of the remote servers.",
		Long:  robinSyncHelpString(),
		Run:   robinSyncCmd,
	}
	robin.AddCommand(robinSync)

	// Establish the commit command, a child of the root command.
	robinCommit := &cobra.Command{
		Use:   "commit [files]",
		Short: "Commit a change to the directory.",
		Long:  robinCommitHelpString(),
		Run:   robinCommitCmd,
	}
	robin.AddCommand(robinCommit)

	// Establish the setup command, a child of the root command.
	robinSetup := &cobra.Command{
		Use:   "setup",
		Short: "Perform setup tasks related to a client or server.",
		Long:  robinSetupHelpString(),
		Run:   robinSetupCmd,
	}
	robin.AddCommand(robinSetup)

	// Establish client initialization, a child of the setup command.
	robinSetupInitclient := &cobra.Command{
		Use:   "init-client",
		Short: "Initialize a robin client.",
		Long:  robinSetupInitclientHelpString(),
		Run:   robinSetupInitclientCmd,
	}
	robinSetup.AddCommand(robinSetupInitclient)

	// Establish the add-remote command, a child of the setup command. Applies
	// primarily to clients.
	robinSetupAddremote := &cobra.Command{
		Use:   "add-remote [address:port]",
		Short: "Add a remote server to a robin client.",
		Long:  robinSetupAddremoteHelpString(),
		Run:   robinSetupAddremoteCmd,
	}
	robinSetup.AddCommand(robinSetupAddremote)

	// Establish the remove-remote command, a child of the setup command.
	// Applies primarily to clients.
	robinSetupRemoveremote := &cobra.Command{
		Use:   "remove-remote [address:port]",
		Short: "Remove a remote server from a robin client.",
		Long:  robinSetupRemoveremoteHelpString(),
		Run:   robinSetupRemoveremoteCmd,
	}
	robinSetup.AddCommand(robinSetupRemoveremote)

	// Establish server initialization, a child of the setup command.
	robinSetupInitserver := &cobra.Command{
		Use:   "init-server [port]",
		Short: "Initialize a robin server.",
		Long:  robinSetupInitserverHelpString(),
		Run:   robinSetupInitserverCmd,
	}
	robinSetup.AddCommand(robinSetupInitserver)

	// Establish the add-key command, a child of the setup command. Applies
	// primarily to servers.
	robinSetupAddkey := &cobra.Command{
		Use:   "add-key [key]",
		Short: "Add a client key to a robin server.",
		Long:  robinSetupAddkeyHelpString(),
		Run:   robinSetupAddkeyCmd,
	}
	robinSetup.AddCommand(robinSetupAddkey)

	// Establish the revoke-key command, a child of the setup command. Applies
	// primarily to clients.
	robinSetupRevokekey := &cobra.Command{
		Use:   "revoke-key [key]",
		Short: "Remove a client key from a robin server.",
		Long:  robinSetupRevokekeyHelpString(),
		Run:   robinSetupRevokekeyCmd,
	}
	robinSetup.AddCommand(robinSetupRevokekey)

	// Add a flag to set the directory.
	root.Flags().StringVarP(&directory, "directory", "d", ".", "The filesystem directory to be used as the root robin directory.")
	// Add a flag to set the commit message.
	root.Flags().StringVarP(&message, "message", "m", "", "The message to accompany an action.")

	// Start the party.
	err := robin.Execute()
	if err != nil {
		os.Exit(64) // EX_USAGE in sysexits.h
	}
}
