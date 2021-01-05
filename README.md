# Commit Colors [![Test Status](https://github.com/sparkbox/commit-colors/workflows/Tests/badge.svg)](https://github.com/sparkbox/commit-colors/actions?workflow=Tests)

See a lovely color swatch in your terminal every time you author a commit. The hexadecimal color comes from the first six characters in your commit ID. Here's what it looks like:

![animated gif demonstrating commit colors](https://raw.githubusercontent.com/sparkbox/commit-colors/master/demo.gif)

Commit colors is written in Go, and is usable on Mac, Windows, and Linux.

*Note: this readme contains instructions for commit-colors 2.0+. commit-colors 1.0 was written in JavaScript and [distributed on npm](https://www.npmjs.com/package/@sparkbox/commit-colors). For instructions on that version of the project, see [the v1 README](https://github.com/sparkbox/commit-colors/tree/76a6b46fed76aeb5e1c813d86ead5185ee1e5cc1).*

## Installation

**Via Homebrew**

```bash
brew tap sparkbox/commit-colors
brew install commit-colors
```

**Via Direct Download**

1. Browse to [the list of releases](https://github.com/sparkbox/commit-colors/releases), and download the package that corresponds to your system (Mac, Windows or Linux).
2. Move the executable to a location on [your PATH](https://superuser.com/a/284351/193516) ([like `/usr/local/bin`](https://superuser.com/q/7150/193516), for example).

<div align="center">· · ·</div></br>

Once installed, you can do a quick test in the terminal:

```bash
commit-colors 123ADD5db8d67ba9621eb5d6765ffdef4c24077a
```

## Commit Hooks Setup

**Option 1: Create the hook by hand**

Copy/paste the following text into a post-commit hook:

```bash
#!/bin/bash
commit-colors $(git rev-parse HEAD)
```

In other words, put the above code in a file named `post-commit` at the location `.git/hooks/post-commit` in your git project of choice. [Make sure this file is executable](https://stackoverflow.com/a/14208849/1154642). If you want this hook to run on all your repos, [see how to do that here](https://stackoverflow.com/q/2293498/1154642).

**Option 2: Use a Git Hooks manager**

See [this list of tools for managing git hooks](https://github.com/aitemr/awesome-git-hooks#tools).

Whatever tool you use, this is the command you'll want it to run: `commit-colors $(git rev-parse HEAD)`

## For Maintainers

### Publishing

1. Create a new tag, and push it up to Github:

```bash
git tag -a 2.0.0 -m "Version 2 release"
git push origin 2.0.0
```

2. The Github Action handles the rest!
  - The action publishes [a Github release](https://github.com/sparkbox/commit-colors/releases) and a [homebrew update](https://github.com/sparkbox/homebrew-commit-colors).
  - Additional publishing settings can be found in `goreleaser.yml`


