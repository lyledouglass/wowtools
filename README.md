# wowtools
* **
## About:
This project is based off other applications that manage certain functions with WoW, specifically addon management and backups. With the Curseforge API changes, many developers are deprecating their applications. While this application will *not* manage addons found on Curseforge, my intent is to offer some of the other features that other applications had, such as WTF folder backups, ElvUI updates, and updates of addons from GitHub projects that are manually managed.

## Requirements
This app now uses the _standalone_ CurseForge launcher as of **v1.4.0**. If you are using the Overwolf client alongside CurseForge, you will need to use v1.3.1 or below

ElvUI is also currently required. This will be addressed in https://github.com/lyledouglass/wowtools/issues/32

## Download
The latest release can be found here: https://github.com/ldougbmx/wowtools/releases
* Download the wowtools.exe and wowtools-cli.yml
  * The README and LICENSE files are also included in the release
* Place the files anywhere on your system (but both in the same directory)
* Verify/Update the wowtools-cli.yml to set custom paths if needed
  * *Depreciated as of v1.4.0*
    * ~~The curseforge_args will need to be filled out for Curseforge to open. If they are not, only the full full Overwolf app will open.~~ 
    * ~~These can be grabbed from the Curseforge exe by right-clicking and selecting Properties.~~ 
    * ~~Copy the target -launchapp xxxxxxxxxxxxxxxxx -from-startmenu section and insert it into the yml.~~
* Run wowtools.exe 
  
## Usage
Update the wowtools-cli.yml file if you have WoW installed in a custom location. This application relies on the yml for file paths

## Example Output
![Alt text](https://github.com/ldougbmx/wowtools/blob/main/images/example-output.png)

## Functionality
This is the standard operation of the app if you don't specify any CLI flags. If CLI flags are detected, this full process will not run. 
1.  Creates `_retail_\Backups`, `_retail_\Backups\WTF` and `_retail_\Backups\ElvUI` directories if they don't exist
2.  Zips up the WTF directory (`C:\Program Files (x86)\World of Warcraft\_retail_\WTF`) and backs it up to `C:\Program Files (x86)\World of Warcraft\_retail_\Backups\WTF`, with the format of YYYY-MM-DD.zip
    1.  Reads the `retention_rate` from the yml and removes the oldest zip file if the count in the folder is higher than it.
3.  Checks for any updates to ElvUI
    1.  If a newer version is found via the API, you will be asked if you want to update
    2.  If Yes
        1.  Zips current ElvUI directories and moves them to `_retail_\Backups\ElvUI`
        2.  Deletes current ElvUI install in the Addons directory
        3.  Downloads latest zip of ElvUI
        4.  Unzips and moves folder to your Addons directory
4. Asks user if they want to open the Curseforge application

## CLI Flags
wowtools allows you to specify specific flags when calling the app from the CLI which will perform specific tasks outside of the functionality listed above

###
`--copy-ptr`
<br>
This flag will run wowtools and *only* do the following
1. Remove the WTF and Interface folder from your PTR install folder
2. Use Windows Robocopy to copy the WTF and Interface folder from your retail folder to the PTR folder.

`--backup-only`
<br>
This flag will run a backup of your WTF folder. Useful for running automation on a scheduled basis.
## Planned enhancements 
* Implement more CLI commands to allow users to perform specific actions on demand.

## Adding an icon to the exe
There are a few options that can be used to add an icon to the executable so it doesn't look as 'ugly' in the Start Menu or a folder you store commonly used apps. 

1. Create a shortcut to the application and add an icon to the shortcut. 
   1. I store my wowtools in my WoW retail folder and create a shortcut to it on my Start menu
   2. The icon I use can be found in the images folder in this repo
2. Another more complex process is using an application like 'Resource Hacker', or even using the [go-winres](https://github.com/tc-hib/go-winres) package and manually build the app from source.
