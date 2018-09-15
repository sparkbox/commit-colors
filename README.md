# Commit Colors

See a lovely color swatch in your terminal every time you author a commit. Here's what it looks like:

![animated gif demonstrating commit colors](https://raw.githubusercontent.com/sparkbox/commit-colors/master/demo.gif)

The hexadecimal color comes from the first six characters in your commit hash.

## Installing

1. Install the package globally

    ```bash
    npm install -g @sparkbox/commit-colors
    ```

2. Copy/paste the following text into a post-commit hook:

    ```
    #!/bin/bash
    SHA=$(git rev-parse HEAD)
    SHA6=${SHA:0:6}
    commit-colors $SHA6
    ```

    In other words, put the above code in a file named `post-commit` at the location `.git/hooks/post-commit` in your git project of choice. [Make sure this file is executable](https://stackoverflow.com/a/14208849/1154642). If you want this hook to run an all your repos, [see how to do that here](https://stackoverflow.com/q/2293498/1154642).



## Todo

- Make it easier to install the commit hook.
- Ensure it is cross-platform.
- Support more color names.
- ???
