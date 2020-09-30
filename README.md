# Commit Colors

See a lovely color swatch in your terminal every time you author a commit. Here's what it looks like:

![animated gif demonstrating commit colors](https://raw.githubusercontent.com/sparkbox/commit-colors/master/demo.gif)

The hexadecimal color comes from the first six characters in your commit hash.

*Note: commit-colors 1.0.0 was written in JavaScript and [distributed on npm](https://www.npmjs.com/package/@sparkbox/commit-colors). For instructions on that version of the project, see [the v1 README](https://github.com/sparkbox/commit-colors/tree/76a6b46fed76aeb5e1c813d86ead5185ee1e5cc1).*

## Installing

1. Install the package with homebrew:

    ```bash
    brew tap sparkbox/commit-colors
    brew install commit-colors
    ```

2. Set up a post-commit hook. Options include:

    **The manual way**

    Copy/paste the following text into a post-commit hook:

    ```bash
    #!/bin/bash
    commit-colors $(git rev-parse HEAD)
    ```

    In other words, put the above code in a file named `post-commit` at the location `.git/hooks/post-commit` in your git project of choice. [Make sure this file is executable](https://stackoverflow.com/a/14208849/1154642). If you want this hook to run on all your repos, [see how to do that here](https://stackoverflow.com/q/2293498/1154642).

    **Using a Git Hooks manager**

    See [this list of tools for managing git hooks](https://github.com/aitemr/awesome-git-hooks#tools).

