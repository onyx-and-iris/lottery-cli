package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"charm.land/huh/v2"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onyx-and-iris/lottery-cli"
)

var version string

func versionFromBuild() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(unable to read version)"
	}
	return strings.Split(info.Main.Version, "-")[0]
}

var rootCmd = &cobra.Command{
	Use:   "lottery",
	Short: "A CLI for National Lottery games.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Fail fast if the count is invalid when the count-prompt flag is not set.
		return validateNonPromptCount()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		kindStr := viper.GetString("kind")
		if kindStr == "" {
			if err := huh.NewSelect[string]().
				Title("Pick a lottery.").
				Options(kindPromptOptions()...).
				Value(&kindStr).Run(); err != nil {
				return err
			}
		}

		kind, err := lottery.ParseKind(kindStr)
		if err != nil {
			return err
		}

		count, err := resolveCount()
		if err != nil {
			return err
		}

		return runLottery(kind, count)
	},
}

// resolveCount resolves the count of draws to generate, either from the command line flag or by prompting the user.
func resolveCount() (int, error) {
	if viper.GetBool("count-prompt") {
		var count int
		if err := huh.NewInput().
			Title("How many draws would you like to generate?").
			Validate(func(s string) error {
				parsedCount, err := parseCount(s)
				if err != nil {
					return err
				}
				count = parsedCount
				return nil
			}).
			Run(); err != nil {
			return 0, err
		}

		return count, nil
	}

	return viper.GetInt("count"), nil
}

// runLottery runs the lottery draw for the specified kind and count.
func runLottery(kind lottery.Kind, count int) error {
	selectedLottery, err := lottery.New(kind)
	if err != nil {
		return err
	}

	includeDrawHeading := count > 1
	renders := make([]string, 0, count)
	drawTitle := "Lottery"

	for i := range count {
		selectedLottery.Draw()
		title, entry := renderDrawEntry(selectedLottery, i+1, includeDrawHeading)
		drawTitle = title
		renders = append(renders, entry)
	}

	if len(renders) > 0 {
		fmt.Println(renderDrawCollection(drawTitle, renders))
	}

	return nil
}

// kindPromptLabel returns a user-friendly label for the given lottery kind.
func kindPromptLabel(kind lottery.Kind) string {
	switch kind {
	case lottery.KindLotto:
		return "Lotto"
	case lottery.KindEuroMillions:
		return "EuroMillions"
	case lottery.KindSetForLife:
		return "Set For Life"
	case lottery.KindThunderball:
		return "Thunderball"
	case lottery.KindPowerball:
		return "Powerball"
	default:
		return string(kind)
	}
}

// kindPromptOptions returns a slice of options for the lottery kind prompt.
func kindPromptOptions() []huh.Option[string] {
	kinds := lottery.AllKinds()
	options := make([]huh.Option[string], 0, len(kinds))
	for _, kind := range kinds {
		options = append(options, huh.NewOption(kindPromptLabel(kind), string(kind)))
	}
	return options
}

func init() {
	rootCmd.Flags().StringP("kind", "k", "", "Lottery kind to generate draws for.")
	rootCmd.Flags().IntP("count", "c", 1, "Number of draws to generate.")
	rootCmd.Flags().BoolP("count-prompt", "C", false, "Prompt for the number of draws to generate.")
	rootCmd.MarkFlagsMutuallyExclusive("count", "count-prompt")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("LOTTERY")
	viper.AutomaticEnv()
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		panic(err)
	}
}

func main() {
	if err := fang.Execute(
		context.Background(),
		rootCmd,
		fang.WithVersion(versionFromBuild()),
	); err != nil {
		os.Exit(1)
	}
}
