# Functionality

`wowtools.exe <flag> <parameter>`

## CLI Flags

wowtools allows you to specify specific flags when calling the app from the CLI which will perform specific tasks outside the functionality listed above

`backup`

- Zips up the WTF directory (`C:\Program Files (x86)\World of Warcraft\_retail_\WTF`) and backs it up to `C:\Program Files (x86)\World of Warcraft\_retail_\Backups\WTF`, with the format of YYYY-MM-DD.zip
- Reads the `retention_rate` from the yml and removes the oldest zip file if the count in the folder is higher than it.

`wtfrestore`

- Performs a destructive action to remove the current WTF folder and replace it with a specific backup from the `C:\Program Files (x86)\World of Warcraft\_retail_\Backups\WTF` directory

`ptrcopy`

- This flag will run wowtools and _only_ do the following
  - Remove the WTF and Interface folder from your PTR install folder
  - Use Windows Robocopy to copy the WTF and Interface folder from your retail folder to the PTR folder.

`addoncopy`

- This flag will copy all or a specific addon from the retail folder to the folder of your choice
  - If no `addonname` parameter is passed, it will copy all addons, first removing the destination folder
  - If the parameter is passed, it will copy that specific addon, first deleting that folder in the destination

`update`

- This flag will run wowtools and _only_ do the following
  - Check the version of the application against the latest release on github
  - If the version is different, it will download the latest release to your downloads folder so you can replace the current executable

`completion`

- This flag will generate a completion script for the application based on the shell you select

`help`

- Shows help text for the application
- This flag is not complete and there is more help text to be added