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
	RunE: func(cmd *cobra.Command, args []string) error {
		var selected string

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
					Value(&selected),
			),
		)
		err := form.Run()
		if err != nil {
			return err
		}

		kind, err := lottery.ParseKind(selected)
		if err != nil {
			return err
		}

		l, err := lottery.New(kind)
		if err != nil {
			return err
		}
		l.Draw()
		fmt.Println(renderDraw(l))

		return nil
	},
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
