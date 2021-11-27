# wowtools
* **
## About:
This project is based off other applications that manage certain functions with WoW, specifically addon management and backups. With the Curseforge API changes, many developers are deprecating their applications. While this application will *not* manage addons found on Curseforge, my intent is to offer some of the other features that other applications had, such as WTF folder backups, ElvUI updates, and updates of addons from GitHub projects that are manually managed.

## Download
The latest release can be found here: https://github.com/ldougbmx/wowtools/releases
* Download the wowtools.exe and wowtools-cli.yml
  * The README and LICENSE files are also included in the release
* Place the files anywhere on your system (but both in the same directory)
* Verify/Update the wowtools-cli.yml to set custom paths if needed
* Run wowtools.exe 
  
## Usage
Update the wowtools-cli.yml file if you have WoW installed in a custom location. This application relies on the yml for file paths

## Current Functionality
Currently, this tool is only zipping up the WTF directory (*C:\Program Files (x86)\World of Warcraft\_retail_\WTF*) and backing it up to *C:\Program Files (x86)\World of Warcraft\_retail_\WTF-Backup*, with the format of YYYY-MM-DD.zip

## Planned enhancements 
* Ability to install/update ElvUI and other Github based projects
* Implement CLI commands to allow users to perform specific actions
