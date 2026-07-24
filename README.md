![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![macOS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=macos&logoColor=F0F0F0)

# lottery

Play National Lottery games from your terminal.

![Selection Prompt](./img/selectionprompt.png)

![Draw](./img/draw.png)

## Install

```console
go install github.com/onyx-and-iris/lottery-cli/cmd/lottery@latest
```

## Configuration

*flags*

-   --count/-c: Number of draws to generate.
-   --count-prompt/-C: Prompt for the number of draws to generate.
> Note. If both --count and --count-prompt are passed the count prompt will win.

*environment variables*

```bash
#!/usr/bin/env bash

export LOTTERY_COUNT=3
export LOTTERY_COUNT_PROMPT=false
```

## Use

There are no subcommands, just run the CLI directly passing any desired flags:

```console
lottery --count=3
```

You will then be entered into the selection prompt.

## Special Thanks

-   [spf13](https://github.com/spf13) for the [cobra](https://github.com/spf13/cobra) package.
-   [Charm](https://github.com/charmbracelet) developers for the [fang](https://github.com/charmbracelet/fang), [lipgloss](https://github.com/charmbracelet/lipgloss) and [huh](https://github.com/charmbracelet/huh) packages.