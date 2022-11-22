# wowtools

## Announcment:
With WowUp officially releasing a new client with Curseforge support, the ElvUI and CF open functionality are now deprecated, as Wowup can now do CF and Elvui updates. This app will still get updates to other utilities like the backups, PTR copies, new character setups, etc.

*If you still want to use ElvUI and CF Open functionality, you will need to use v2.4.1 or below*

* **
## About:
This project is based off other applications that manage certain functions with WoW, specifically addon management and backups. With the Curseforge API changes, many developers are deprecating their applications. While this application will *not* manage addons found on Curseforge, my intent is to offer some of the other features that other applications had, such as WTF folder backups, PTR profiles copies, new character setups, etc.

## Requirements
* This application is written and compiled in Go. All dependencies are compiled in the application executable.
* You will need to have the wowtools-cli.yml in the same directory as the exe

## Download
The latest release can be found here: https://github.com/ldougbmx/wowtools/releases
* Download the wowtools.exe and wowtools-cli.yml
  * The README and LICENSE files are also included in the release
* Place the files anywhere on your system (but both in the same directory)
* Verify/Update the wowtools-cli.yml to set custom paths if needed
* Run wowtools.exe 
  
## Usage
Update the wowtools-cli.yml file if you have WoW installed in a custom location. This application relies on the yml for file paths.
<br>
## Example Output
![Alt text](https://github.com/ldougbmx/wowtools/blob/main/images/example-output.png)

## Functionality
This is the standard operation of the app if you don't specify any CLI flags. If CLI flags are detected, this full process will not run. 
1.  Creates `_retail_\Backups` and `_retail_\Backups\WTF` directories if they don't exist

## CLI Flags
wowtools allows you to specify specific flags when calling the app from the CLI which will perform specific tasks outside of the functionality listed above
###
`backup`
* Zips up the WTF directory (`C:\Program Files (x86)\World of Warcraft\_retail_\WTF`) and backs it up to `C:\Program Files (x86)\World of Warcraft\_retail_\Backups\WTF`, with the format of YYYY-MM-DD.zip
* Reads the `retention_rate` from the yml and removes the oldest zip file if the count in the folder is higher than it.
###
`copy-ptr`
* This flag will run wowtools and *only* do the following
* Remove the WTF and Interface folder from your PTR install folder
* Use Windows Robocopy to copy the WTF and Interface folder from your retail folder to the PTR folder.

## Planned enhancements 
* Implement more CLI commands to allow users to perform specific actions on demand.

## Adding an icon to the exe
There are a few options that can be used to add an icon to the executable so it doesn't look as 'ugly' in the Start Menu or a folder you store commonly used apps. 

1. Create a shortcut to the application and add an icon to the shortcut. 
   1. I store my wowtools in my WoW retail folder and create a shortcut to it on my Start menu
   2. The icon I use can be found in the images folder in this repo
2. Another more complex process is using an application like 'Resource Hacker', or even using the [go-winres](https://github.com/tc-hib/go-winres) package and manually build the app from source.
