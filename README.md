# pkgpal

pkgpal (name pending) is a CLI tool written in Go. It's designed to be the successor to my side project [pkgparse-node](https://github.com/marcus-crane/pkgparse-node) but where that was Node specific, this version is language agnostic.

# What is this?

The main functionality is that you can easily query packages from various package managers and also automatically parse, and better understand package manager files.

For example, with pkgparse-node, you could execute `pkgparse feast` in a directory with a `package.json` file and you would be returned a list of dependencies and their descriptions. It was useful for figuring out what dependenices a project has, and also usually what functionality you could expect from it.

This project strives to do the same but in a way that you could run eg; `pkgpal feast` in any directory and it'll automatically detect what files are present. Let's say it finds a requirements.txt file then it would parse that and query pypi by way of the [pkgparse server](https://github.com/marcus-crane/pkgparse).

# Development

The project currently has a functional CLI although it only has searching implemented with 2 registries. It still needs some cleaning up but for one of searching for either NPM or PyPi, it's usable.

You can install everything needed by cloning the project into your Go path and running `make setup`.

To build a binary for your system, run `make {platform}` where platform is either `windows`, `mac` or `linux`.

That will generate an executable binary. If you're on `mac` or `linux`, running `make bin` will attempt to move the binary to your `/usr/local/bin` folder. Once there, the binary will be in your `$PATH` meaning you can execute `pkgpal` from anywhere on your system.

If you're on Windows, *shrug* I suppose you just execute it from the command prompt? I'll have to get back to you on that. Sorry!

# Tests

 I'm trying to do this project using TDD which is interesting and hard because I'm still fairly new to go.

 You can run the current tests with `make test`.
