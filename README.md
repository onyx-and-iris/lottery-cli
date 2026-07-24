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

-   --kind/-k: The kind of lottery.
-   --count/-c: Number of draws to generate.
-   --count-prompt/-C: Prompt for the number of draws to generate.
> Note. If both --count and --count-prompt are passed the count prompt will win.

*environment variables*

```bash
#!/usr/bin/env bash

export LOTTERY_KIND=lotto
export LOTTERY_COUNT=3
export LOTTERY_COUNT_PROMPT=false
```

## Use

Run with the selection prompt without prompting for a count:

```console
lottery
```

Run with the selection prompt but also prompt for a count:

```console
lottery --count-prompt
```

Run with the selection prompt but pass in the count directly:

```console
lottery --count=4
```

Run a single draw directly:

```console
lottery -k=euromillions 
```

Run multiple draws directly:

```console
lottery -k=euromillions -c=3
```

## Special Thanks

-   [spf13](https://github.com/spf13) for the [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper) packages.
-   [Charm](https://github.com/charmbracelet) developers for the [fang](https://github.com/charmbracelet/fang), [lipgloss](https://github.com/charmbracelet/lipgloss) and [huh](https://github.com/charmbracelet/huh) packages.