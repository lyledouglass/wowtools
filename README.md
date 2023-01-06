# wowtools

## Announcment:

Please see the latest announcements under the [Discussion](https://github.com/lyledouglass/wowtools/discussions) tab

---

## About:

This project is based off other applications that manage certain functions with WoW, specifically api scrapping, addon management, and backups. With the Curseforge API changes, many developers are deprecating their applications. While this application will _not_ manage addons found on Curseforge, my intent is to offer some of the other features that other applications had, such as WTF folder backups, PTR profiles copies, new character setups, etc.

---

## Windows

<details>
<summary>Expand</summary>

### Requirements

- This application is written and compiled in Go. All dependencies are compiled in the application executable.
- You will need to have the `config.yaml` in the same directory as the exe

### Download

The latest release can be found here: https://github.com/ldougbmx/wowtools/releases

- Download the wowtools.exe and config.yaml
  - The README and LICENSE files are also included in the release
- Place the files anywhere on your system (but both in the same directory)
- Verify/Update the config.yml to set custom paths if needed

### Usage

Update the config.yaml file if you have WoW installed in a custom location. This application relies on the yaml for file paths. You can leave the Linux options blank

### Functionality

This is the standard operation of the app if you don't specify any CLI flags. If CLI flags are detected, this full process will not run.

1.  Creates `_retail_\Backups` and `_retail_\Backups\WTF` directories if they don't exist

### CLI Flags

wowtools allows you to specify specific flags when calling the app from the CLI which will perform specific tasks outside the functionality listed above

###

`backup`

- Zips up the WTF directory (`C:\Program Files (x86)\World of Warcraft\_retail_\WTF`) and backs it up to `C:\Program Files (x86)\World of Warcraft\_retail_\Backups\WTF`, with the format of YYYY-MM-DD.zip
- Reads the `retention_rate` from the yml and removes the oldest zip file if the count in the folder is higher than it.

###

`copy-ptr`

- This flag will run wowtools and _only_ do the following
- Remove the WTF and Interface folder from your PTR install folder
- Use Windows Robocopy to copy the WTF and Interface folder from your retail folder to the PTR folder.

### Adding an icon to the exe

There are a few options that can be used to add an icon to the executable so it doesn't look as 'ugly' in the Start Menu or a folder you store commonly used apps.

1. Create a shortcut to the application and add an icon to the shortcut.
   1. I store my wowtools in my WoW retail folder and create a shortcut to it on my Start menu
   2. The icon I use can be found in the images folder in this repo
2. Another more complex process is using an application like 'Resource Hacker', or even using the [go-winres](https://github.com/tc-hib/go-winres) package and manually build the app from source.
</details>

---

## Docker

<details>
    <summary>Expand</summary>
        The `wowtools_server` application is compiled for Linux and published to dockerhub for use in docker.

### Requirements

- Docker

## Download/Installation

The latest release can be found here: LINK_TO_DOCKER_IMAGE
<br>
Docker-Compose is the suggested method for creating and maintaining the deployment

```yaml
version: "3"
services:
  wowtools:
    container_name: wowtools
    image: "image_name"
    env:
      BlizzardAccessToken: ""
      WebhookUri: ""
    restart: unless-stopped
```

### Functionality

- Scrapes the WoW Token Price via the Blizzard and reports it if the sale price is above a set value
- More coming soon &trade;
</details>

---

## Planned enhancements

- Implement more CLI commands to allow users to perform specific actions on demand.
- Implement better method of managing secrets inside the docker container
