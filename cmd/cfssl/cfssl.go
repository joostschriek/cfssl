/*
cfssl is the command line tool to issue/sign/bundle client certificate. It's
also a tool to start a HTTP server to handle web requests for signing, bundling
and verification.

Usage:
	cfssl command [-flags] arguments

	The commands are

	bundle	 create a certificate bundle
	sign	 signs a certificate signing request (CSR)
	serve	 starts a HTTP server handling sign and bundle requests
	version	 prints the current cfssl version
	genkey   generates a key and an associated CSR
	gencert  generates a key and a signed certificate
	gencsr   generates a certificate request
	selfsign generates a self-signed certificate

Use "cfssl [command] -help" to find out more about a command.
*/
package main

import (
	"flag"
	"os"

	"github.com/joostschriek/cfssl/cli"
	"github.com/joostschriek/cfssl/cli/bundle"
	"github.com/joostschriek/cfssl/cli/certinfo"
	"github.com/joostschriek/cfssl/cli/crl"
	"github.com/joostschriek/cfssl/cli/gencert"
	"github.com/joostschriek/cfssl/cli/gencrl"
	"github.com/joostschriek/cfssl/cli/gencsr"
	"github.com/joostschriek/cfssl/cli/genkey"
	"github.com/joostschriek/cfssl/cli/info"
	"github.com/joostschriek/cfssl/cli/ocspdump"
	"github.com/joostschriek/cfssl/cli/ocsprefresh"
	"github.com/joostschriek/cfssl/cli/ocspserve"
	"github.com/joostschriek/cfssl/cli/ocspsign"
	printdefaults "github.com/joostschriek/cfssl/cli/printdefault"
	"github.com/joostschriek/cfssl/cli/revoke"
	"github.com/joostschriek/cfssl/cli/scan"
	"github.com/joostschriek/cfssl/cli/selfsign"
	"github.com/joostschriek/cfssl/cli/serve"
	"github.com/joostschriek/cfssl/cli/sign"
	"github.com/joostschriek/cfssl/cli/version"

	_ "github.com/go-sql-driver/mysql" // import to support MySQL
	_ "github.com/lib/pq"              // import to support Postgres
	_ "github.com/mattn/go-sqlite3"    // import to support SQLite3
)

// main defines the cfssl usage and registers all defined commands and flags.
func main() {
	// Add command names to cfssl usage
	flag.Usage = nil // this is set to nil for testability
	// Register commands.
	cmds := map[string]*cli.Command{
		"bundle":         bundle.Command,
		"certinfo":       certinfo.Command,
		"crl":            crl.Command,
		"sign":           sign.Command,
		"serve":          serve.Command,
		"version":        version.Command,
		"genkey":         genkey.Command,
		"gencert":        gencert.Command,
		"gencsr":         gencsr.Command,
		"gencrl":         gencrl.Command,
		"ocspdump":       ocspdump.Command,
		"ocsprefresh":    ocsprefresh.Command,
		"ocspsign":       ocspsign.Command,
		"ocspserve":      ocspserve.Command,
		"selfsign":       selfsign.Command,
		"scan":           scan.Command,
		"info":           info.Command,
		"print-defaults": printdefaults.Command,
		"revoke":         revoke.Command,
	}

	// If the CLI returns an error, exit with an appropriate status
	// code.
	err := cli.Start(cmds)
	if err != nil {
		os.Exit(1)
	}
}
