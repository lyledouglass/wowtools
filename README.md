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
  * The curseforge_args will need to be filled out for Curseforge to open. If they are not, only the full full Overwolf app will open. 
  * These can be grabbed from the Curseforge exe by right-clicking and selecting Properties. 
  * Copy the target -launchapp xxxxxxxxxxxxxxxxx -from-startmenu section and insert it into the yml.
* Run wowtools.exe 
  
## Usage
Update the wowtools-cli.yml file if you have WoW installed in a custom location. This application relies on the yml for file paths

## Current Functionality
1.  Creates `_retail_\Backups`, `_retail_\Backups\WTF` and `_retail_\Backups\ElvUI` directories if they don't exist
2.  Zips up the WTF directory (`C:\Program Files (x86)\World of Warcraft\_retail_\WTF`) and backs it up to `C:\Program Files (x86)\World of Warcraft\_retail_\Backups\WTF`, with the format of YYYY-MM-DD.zip
3.  Checks for any updates to ElvUI
    1.  If a newer version is found via the API, you will be asked if you want to update
    2.  If Yes
        1.  Zips current ElvUI directories and moves them to `_retail_\Backups\ElvUI`
        2.  Deletes current ElvUI install in the Addons directory
        3.  Downloads latest zip of ElvUI
        4.  Unzips and moves folder to your Addons directory
4. Asks user if they want to open the Curseforge application

## Planned enhancements 
* Implement CLI commands to allow users to perform specific actions on demand. 
