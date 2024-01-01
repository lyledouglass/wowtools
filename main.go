package main

import (
	"sync"
	"wowtools/internal"
	util "wowtools/pkg/utilities"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {

	const version = "5.0.2"

	var (
		logging     string
		help        bool
		versionFlag bool
		wtfzip      string
	)

	util.LoadConfig(".")

	var rootCmd = &cobra.Command{
		Use:  "wowtools",
		Long: "wowtools is a CLI tool for managing World of Warcraft file configurations.",
		Run: func(cmd *cobra.Command, args []string) {
			util.SetupLogger("info")
			if versionFlag {
				util.Log.Info("wowtools version: " + version)
			}
		},
	}

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update your erp-devops-workflow-assistant application",
		Run: func(cmd *cobra.Command, args []string) {
			util.SetupLogger(logging)
			internal.UpdateApp(version)
		},
	}

	var backupCmd = &cobra.Command{
		Use:   "backup",
		Short: "Backup your WTF folder",
		Run: func(cmd *cobra.Command, args []string) {
			util.SetupLogger(logging)
			// WaitGroup for creating missing folders.
			util.Log.Debug("Creating WaitGroup if folders are missing")
			var wg sync.WaitGroup
			wg.Add(2)
			go util.VerifyFolders(viper.GetString("backup_dir"), &wg)
			go util.VerifyFolders(viper.GetString("backup_dir")+"WTF", &wg)
			wg.Wait()

			util.Log.Debug("Beginning WtfBackup function")
			internal.WtfBackup()
		},
	}

	var wtfRestoreCmd = &cobra.Command{
		Use:   "wtfrestore",
		Short: "Restore your WTF folder",
		Run: func(cmd *cobra.Command, args []string) {
			util.SetupLogger(logging)
			if wtfzip == "" {
				util.Log.Fatal("You must specify a WTF zip file to restore")
			}
			internal.WtfRestore(wtfzip)
		},
	}

	var ptrCopyCmd = &cobra.Command{
		Use:   "ptrcopy",
		Short: "Copy PTR data from Retail",
		Run: func(cmd *cobra.Command, args []string) {
			util.SetupLogger(logging)
			internal.CopyPtrData()
		},
	}

	// updateCmd Flags
	updateCmd.Flags().StringVarP(&logging, "logging", "l", "info", "Enables logging. Options are: trace, debug, info, warn, error, fatal, panic")

	// backupCmd Flags
	backupCmd.Flags().StringVarP(&logging, "logging", "l", "info", "Enables logging. Options are: trace, debug, info, warn, error, fatal, panic")
	backupCmd.Flags().BoolVarP(&help, "help", "h", false, "Displays useful information")

	// wtfRestore Flags
	wtfRestoreCmd.Flags().StringVarP(&logging, "logging", "l", "info", "Enables logging. Options are: trace, debug, info, warn, error, fatal, panic")
	wtfRestoreCmd.Flags().BoolVarP(&help, "help", "h", false, "Displays useful information")
	wtfRestoreCmd.Flags().StringVarP(&wtfzip, "wtfzip", "w", "", "File name of WTF Zip")

	// ptrCopyCmd Flags
	ptrCopyCmd.Flags().StringVarP(&logging, "logging", "l", "info", "Enables logging. Options are: trace, debug, info, warn, error, fatal, panic")
	ptrCopyCmd.Flags().BoolVarP(&help, "help", "h", false, "Displays useful information")

	// rootCmd Flags
	rootCmd.Flags().BoolVarP(&help, "help", "h", false, "Displays useful information")
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Displays the current version of wowtools")

	// Add commands to rootCmd
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(wtfRestoreCmd)
	rootCmd.AddCommand(ptrCopyCmd)

	if err := rootCmd.Execute(); err != nil {
		util.Log.Fatal(err)
	}
}
