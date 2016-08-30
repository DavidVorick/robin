// Package robin aims to be a simple, complete, secure backup and
// synchronization tool for personal data. Robin will take a single directory
// (and all sub-folders) and securely synchronize that directory between
// multiple remote servers.
//
// When a file is added to the directory structure, an encrypted log entry is
// made to the server detailing the change. The message is in the form 'add the
// data with hash X to file location 'Y'. The data is padded such that the
// encrypted filesize rounded up to the nearest power of 2 (in bytes). The log
// entry itself is also padded and rounded up to the nearest power of 2. Though
// there is some overhead to this method, though as a result the server knows
// very little about the files being stored. The encrypted log entry contains a
// checksum of the data, preventing the server from invisibly modifying data
// that it has been given. The log entries contain an encrypted hash of their
// contents, and an encrypted hash of the previous log entry. As a result, the
// server has limited ability to omit or modify log entries as well. The
// smallest log entry allowed is 2^10 bytes (1 KiB), preventing the server from
// learning anything from tiny log entries. Ultimately, there's very little
// that the server can learn or do with the data that it's protecting.
//
// The server also operates with security in mind. Before a log entry can be
// added, the client must sign a cryptographic challenge using an approved
// public key. The server will never delete data, meaning that a compromised
// host has no power to destroy data on an uncompromised backup server. Log
// entries can specify that data has been deleted, which means it will be
// removed from the primary directory, but the data + log entry that originally
// added the file are both kept, meaning they can be recovered so long as the
// server is never compromised.
//
// Having multiple remote servers is both supported and recommended. Changes
// will be pushed to all servers, and they will be kept in synchronization. In
// the event that the servers partition or desynchronize, a merge process can
// be used which has no risk of destroying data. A file may be written and then
// over-written, but previous versions of the file will appear in the changelog
// and can always be recovered. Robin also guarantees data durability, meaning
// that unexpected network failures or power outages mid-sync will not cause
// data loss or corruption.
//
// Note that due to limitations with Windows, Windows machines cannot provide
// the same level of disk-efficiency, nor can Windows guarantee that files will
// not be corrupted in the event of power failure or network interruption. Full
// security (in terms of encryption and privacy) can still be achieved on
// Windows servers. An open source operating system such as Linux is highly
// recommended, especially for clients.
package main

const (
	dataDir = "robin-data"
)
