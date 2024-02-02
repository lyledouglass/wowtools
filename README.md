# wowtools

## Announcment

Please see the latest announcements under the [Discussion](https://github.com/lyledouglass/wowtools/discussions) tab

## About

This project is based off other applications that manage certain functions with WoW, specifically api scrapping, addon management, and backups. With the Curseforge API changes, many developers are deprecating their applications. While this application will _not_ manage addons found on Curseforge, my intent is to offer some of the other features that other applications had, such as WTF folder backups, PTR profiles copies, new character setups, etc.

## Requirements

- This application is written and compiled in Go. All dependencies are compiled in the application executable.
- You will need to have the `config.yaml` in the same directory as the exe

## Functionality

This is the standard operation of the app if you don't specify any CLI flags. If CLI flags are detected, this full process will not run.

1. Creates `_retail_\Backups` and `_retail_\Backups\WTF` directories if they don't exist

### CLI Flags

wowtools allows you to specify specific flags when calling the app from the CLI which will perform specific tasks outside the functionality listed above

`backup`

- Zips up the WTF directory (`C:\Program Files (x86)\World of Warcraft\_retail_\WTF`) and backs it up to `C:\Program Files (x86)\World of Warcraft\_retail_\Backups\WTF`, with the format of YYYY-MM-DD.zip
- Reads the `retention_rate` from the yml and removes the oldest zip file if the count in the folder is higher than it.

`restore`

- Performs a destructive action to remove the current WTF folder and replace it with a specific backup from the `C:\Program Files (x86)\World of Warcraft\_retail_\Backups\WTF` directory

`copy-ptr`

- This flag will run wowtools and _only_ do the following
- Remove the WTF and Interface folder from your PTR install folder
- Use Windows Robocopy to copy the WTF and Interface folder from your retail folder to the PTR folder.

## Download/Installation

The latest release can be found on the [releases](https://github.com/ldougbmx/wowtools/releases) page

- Download the wowtools.exe and config.yaml
  - The README and LICENSE files are also included in the release
- Place the files anywhere on your system (but both in the same directory)
- Verify/Update the config.yml to set custom paths if needed

## Planned enhancements

For planned enhancements, please review the [issues](https://github.com/lyledouglass/wowtools/issues) tab in the github repository
