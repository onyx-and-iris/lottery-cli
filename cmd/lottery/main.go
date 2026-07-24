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

var cmd = &cobra.Command{
	Use:   "lottery",
	Short: "A CLI for National Lottery games.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		kindStr := viper.GetString("kind")
		if kindStr != "" {
			_, err := lottery.ParseKind(kindStr)
			if err != nil {
				return err
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		kindStr := viper.GetString("kind")
		if kindStr == "" {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Pick a lottery.").
						Options(
							huh.NewOption("Lotto", "lotto"),
							huh.NewOption("EuroMillions", "euromillions"),
							huh.NewOption("Set For Life", "setforlife"),
							huh.NewOption("Thunderball", "thunderball"),
							huh.NewOption("Powerball", "powerball"),
						).
						Value(&kindStr),
				),
			)
			err := form.Run()
			if err != nil {
				return err
			}
		}

		kind, err := lottery.ParseKind(kindStr)
		if err != nil {
			return err
		}

		l, err := lottery.New(kind)
		if err != nil {
			return err
		}

		if countPrompt := viper.GetBool("count-prompt"); countPrompt {
			var count string
			countPrompt := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("How many draws would you like to generate?").
						Value(&count),
				),
			)
			err := countPrompt.Run()
			if err != nil {
				return err
			}
			viper.Set("count", count)
		}

		count := viper.GetInt("count")
		includeDrawHeading := count > 1
		renders := make([]string, 0, count)
		drawTitle := "Lottery"

		for i := range count {
			l.Draw()
			title, entry := renderDrawEntry(l, i+1, includeDrawHeading)
			drawTitle = title
			renders = append(renders, entry)
		}

		if len(renders) > 0 {
			fmt.Println(renderDrawCollection(drawTitle, renders))
		}

		return nil
	},
}

func init() {
	cmd.Flags().StringP("kind", "k", "", "Lottery kind to generate draws for.")
	cmd.Flags().IntP("count", "c", 1, "Number of draws to generate.")
	cmd.Flags().BoolP("count-prompt", "C", false, "Prompt for the number of draws to generate.")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("LOTTERY")
	viper.AutomaticEnv()
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
}

func main() {
	if err := fang.Execute(
		context.Background(),
		cmd,
		fang.WithVersion(versionFromBuild()),
	); err != nil {
		os.Exit(1)
	}
}
