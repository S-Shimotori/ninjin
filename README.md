# ninjin

Why don't we have :ninjin: !!?!?!?!?!?!??!?

## Description

```ninjin``` calls ```xcode-select -s``` .

```xcode-select -h``` :  
> Print or change the path to the active developer directory. This directory controls which tools are used for the Xcode command line tools (for example, xcodebuild) as well as the BSD development commands (such as cc and make).

## Usage

```sh
$ ninjin list
  Xcode.app (7.1.1 7B1005)
* Xcode-beta.app (7.2 7C62b)
```

```sh
$ ninjin switch 7.1.1
Password:
succeed in switching to Xcode(version 7.1.1 7B1005)
$ ninjin switch 7C62b
succeed in switching to Xcode(version 7.2 7C62b)
```

```sh
$ ninjin switch-at-least 7.1
succeed in switching to Xcode(version 7.2 7C62b)
```

```sh
$ ninjin switch-compatible 7.1.0
succeed in switching to Xcode(version 7.1.1 7B1005)
```

```sh
$ ninjin switch-latest
succeed in switching to Xcode(version 7.2 7C62b)
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/S-Shimotori/ninjin
```

## Contribution

1. Fork ([https://github.com/S-Shimotori/ninjin/fork](https://github.com/S-Shimotori/ninjin/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[S-Shimotori](https://github.com/S-Shimotori)
