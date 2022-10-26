//go:generate goversioninfo

package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"sync"
	"wowtools/internal"
	"wowtools/pkg/utilities"
)

func main() {
	var (
		copyPtr           bool
		backupOnly        bool
		noUpdates         bool
		standardFunctions bool
		characterCopy     bool
		serverName        string
		characterName     string
		accountName       string
	)

	// Flags
	flag.BoolVar(&copyPtr, "copy-ptr", false, "only performs copy of PTR folders from Retail")
	flag.BoolVar(&backupOnly, "backup-only", false, "perfomrs only backup of wtf folder")
	flag.BoolVar(&noUpdates, "no-updates", false, "skips checking updates for wowtools")
	// Default this to true, as most users will want this at runtime
	flag.BoolVar(&standardFunctions, "standard-functions", true, "performs standard functions. `"+
		"Defaults to true")

	// Character Copy Flags
	flag.BoolVar(&characterCopy, "character-copy", false, "copy a template folder to `"+
		"account for a new character")
	flag.StringVar(&serverName, "server-name", "", "server the new character will exist on")
	flag.StringVar(&characterName, "character-name", "", "name of character")
	flag.StringVar(&accountName, "account-name", "", "account name you want this character to live in")

	flag.Parse()

	internal.InitConfig()
	// Check for updates to the application
	if !noUpdates {
		internal.UpdateWowtools()
	}

	// WaitGroup for creating missing folders.
	var wg sync.WaitGroup
	wg.Add(3)
	go utilities.VerifyFolders(viper.GetString("backup_dir"), &wg)
	go utilities.VerifyFolders(viper.GetString("backup_dir")+"ElvUI", &wg)
	go utilities.VerifyFolders(viper.GetString("backup_dir")+"WTF", &wg)
	wg.Wait()

	if copyPtr {
		internal.CopyPtrData()
	}

	if backupOnly {
		internal.WtfBackup()
	}

	if characterCopy {
		if serverName == "" || characterName == "" || accountName == "" {
			log.Fatal("Missing server name, character name, or account name")
		}
		internal.NewCharacter(serverName, characterName, accountName)
	}

	if standardFunctions {
		internal.WtfBackup()
		if viper.GetString("elvui_dir") != "" {
			internal.UpdateElvUI()
		}
		internal.OpenCurseforge()
	}
}
