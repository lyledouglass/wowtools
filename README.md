# wowtools

## Announcments

Please see the latest announcements under the [Discussion](https://github.com/lyledouglass/wowtools/discussions) tab

## About

This project is based off other applications that manage certain functions with WoW, specifically api scrapping, addon management, and backups. With the Curseforge API changes, many developers are deprecating their applications. While this application will _not_ manage addons found on Curseforge, my intent is to offer some of the other features that other applications had, such as WTF folder backups, PTR profiles copies, new character setups, etc.

## Requirements

- This application is written and compiled in Go. All dependencies are compiled in the application executable.
- You will need to have the `config.yaml` in the same directory as the exe

## Download/Installation

The latest release can be found on the [releases](https://github.com/lyledouglass/wowtools/releases) page

- Download the wowtools.exe and config.yaml
  - The README and LICENSE files are also included in the release
- Place the files anywhere on your system (but both in the same directory)
- Verify/Update the config.yml to set custom paths if needed

## Functionality

Functionality of the app can be found in the [functionality.md](https://github.com/lyledouglass/wowtools/blob/main/functionality.md) file

## Planned enhancements

For planned enhancements, please review the [issues](https://github.com/lyledouglass/wowtools/issues) tab in the github repository

## Contributing

### Pull Requests

Every change in the codebase should be done via pull requests. This
applies to everyone, including the project maintainers. This is to
ensure that every change is reviewed and discussed before it goes into
the codebase, and builds are succeeding

### Conventional Commits

In order to have a clean and readable commit history, this repository
requires the use of [ConventionalCommits](https://www.conventionalcommits.org/en/v1.0.0/).
Using conventional commits will allow for automatic versioning and
changelog generation as well
