# Incantations

Silly one off tasks that are nice to have automated.

## Setup
Create `~/.cast.yaml` with your own [personal access token](https://github.com/settings/tokens).
```
github_personal_access_token: asdfqwerty1234
```

## Github

### Cast Commands
```
(cast(master)) $ cast -h
These incantations help automate the creation of a Deft workflow:

Cast is a CLI library for Go that quickly automates the
setup of a Deft workflow. This application is a tool to
update the necessary configuration of repositories and
other tools to prepare for SaaS application development.

Usage:
  cast [command]

Available Commands:
  help        Help about any command
  labels      Manipulate labels on a repository
  repos       Actions to be taken on github repositories

Flags:
      --config string   config file (default is $HOME/.cast.yaml)
  -h, --help            help for cast
  -t, --toggle          Help message for toggle

```

### List Repos
```
(cast(master)) $ cast repos list -h
List repositories

Usage:
  cast repos list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --config string   config file (default is $HOME/.cast.yaml)
```

### List labels on a repo
```
(cast(master)) $ cast labels list -h | pbcopy
lists the labels on the repository

Usage:
  cast labels list [org-name/repo-name] [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --config string   config file (default is $HOME/.cast.yaml)

```

### Delete all labels on a repo
```
Deletes all labels on the repositories

Usage:
  cast labels delete [org-name/repo-name ...] [flags]

Flags:
  -h, --help   help for delete

Global Flags:
      --config string   config file (default is $HOME/.cast.yaml)
```

### Create deft's standard labels on a repo
```
Creates the Deft labels on the repository

Usage:
  cast labels create [org-name/repo-name ...] [flags]

Flags:
  -h, --help   help for create

Global Flags:
      --config string   config file (default is $HOME/.cast.yaml)
```

