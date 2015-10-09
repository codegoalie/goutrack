# GouTrack

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

GouTrack (pronounced _goo-trak_) is a command line interface for the YouTrack
project management system.

## Install

Download one of the [pre-built
binaries](https://github.com/chrismar035/goutrack/releases/latest) for your
operating system and place it in your PATH. If your OS does not have a
pre-built binary, please open an issue and I will make a binary for you, or
follow the steps below to install from source.

Alternatively to the pre-built binaries, you can install from source by
go-getting:

```
$ go get github.com/chrismar035/goutrack
```

## Configuration

GouTrack requires three pieces of information in order to locate your YouTrack
installation and authenticate with your account. These are your YouTrack
installation host, or URL, your username, and your password. There are two ways
to specify this information.

The first and easiest way to specify your YouTrack information is with command
line flags.

```
$ goutrack -h <base YouTrack URL> -u <YouTrack username> -p <YouTrack Password>
```

The second way to specify your login info is to add your information to a config
file. This file lives at `~/.goutrack` and is a YAML file. An example would be:

```
# ~/.goutrack

host: https://youtrack.com
username: best.youtrack.user
password: s00$up3rs3kr37
```

When you have a config file, you do not need to specify your credentials with
command line flags. However, command line flags take prescedence over the
config file.

## Usage

GouTrack uses a sub-command interface (think git). The basic format of commands
are:

```
$ goutrack <command> <story ID> [optional params]
```

### Get a story

The `g` sub-command fetches basic details about at story. There are no optional
params for this command.

```
$ goutrack g <story ID>
```

This command prints the raw response to the console.

### Command a story

The `c` sub-command sends a YouTrack command string to be applied to the given
story. The `c` command takes a command string parameter and an optional comment
string second parameter

```
$ goutrack c <story> "<command string>" ["<comment string>"]
```

The `c` command prints a verification string to the console and then the raw
response from YouTrack. Successful commands return `<nil>`.

#### Examples

```
$ goutrack c yt-1234 "Assignee me" "I'll take this story."
```

```
$ goutrack c yt-4321 "Subsystem server-side"
```

## Aliases

GouTrack has aliases which can shorten frequently used commands. Simply define
a set of aliases in your config file and use the `a` sub-command to expand the
alias.

```
# ~/.goutrack

...
aliases:
  - name: complete
    command: "State: Complete"
  - name: mine
    command: "Assigned to : me"
```

Then, you can simply specify the alias as the command (with an optional comment
string).

```
$ goutrack a c-1234 complete
Applying State: Complete to c-1234
```

## Contributing

When contributing please

1. [Fork](https://github.com/chrismar035/goutrack/fork) the repo
1. Start a topic branch.
1. Write awesome code!
   1. Please break your commits into logical units.
   1. Please add specs when necessary.
1. Open a [Pull
   Request](https://github.com/chrismar035/goutrack/pulls)
1. See your awesomeness merged in!

# Thanks and Happy GouTracking!
