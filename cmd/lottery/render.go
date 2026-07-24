package main

import (
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/onyx-and-iris/lottery-cli"
)

// nolint:misspell
var (
	titleStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	drawHeadingStyle = lipgloss.NewStyle().
				Bold(true).
				Underline(true).
				Foreground(lipgloss.Color("14"))
	labelStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	numbersStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))
	specialStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("11"))
	separatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	cardStyle      = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("8")).
			Padding(0, 1).
			MarginTop(1)
)

func formatNumberList(numbers []int) string {
	parts := make([]string, len(numbers))
	for i, n := range numbers {
		parts[i] = strconv.Itoa(n)
	}
	return strings.Join(parts, "  ")
}

func drawTitleAndLines(l lottery.Lottery) (string, []string) {
	var title string
	var lines []string

	switch game := l.(type) {
	case *lottery.Lotto:
		title = "Lotto"
		lines = append(
			lines,
			labelStyle.Render(
				"Numbers:",
			)+" "+numbersStyle.Render(
				formatNumberList(game.Numbers[:]),
			),
		)
	case *lottery.EuroMillions:
		title = "EuroMillions"
		lines = append(
			lines,
			labelStyle.Render("Main:")+" "+numbersStyle.Render(formatNumberList(game.Numbers[:])),
		)
		lines = append(
			lines,
			labelStyle.Render(
				"Lucky Stars:",
			)+" "+specialStyle.Render(
				formatNumberList(game.LuckyStars[:]),
			),
		)
	case *lottery.SetForLife:
		title = "Set For Life"
		lines = append(
			lines,
			labelStyle.Render("Main:")+" "+numbersStyle.Render(formatNumberList(game.Numbers[:])),
		)
		lines = append(
			lines,
			labelStyle.Render("Life Ball:")+" "+specialStyle.Render(strconv.Itoa(game.LifeBall)),
		)
	case *lottery.Thunderball:
		title = "Thunderball"
		lines = append(
			lines,
			labelStyle.Render("Main:")+" "+numbersStyle.Render(formatNumberList(game.Numbers[:])),
		)
		lines = append(
			lines,
			labelStyle.Render(
				"Thunderball:",
			)+" "+specialStyle.Render(
				strconv.Itoa(game.Thunderball),
			),
		)
	case *lottery.Powerball:
		title = "Powerball"
		lines = append(
			lines,
			labelStyle.Render("Main:")+" "+numbersStyle.Render(formatNumberList(game.Numbers[:])),
		)
		lines = append(
			lines,
			labelStyle.Render("Powerball:")+" "+specialStyle.Render(strconv.Itoa(game.Powerball)),
		)
	default:
		title = "Lottery"
		lines = append(lines, "Unknown lottery type")
	}

	return title, lines
}

func renderDrawEntry(l lottery.Lottery, drawNumber int, includeHeading bool) (string, string) {
	title, lines := drawTitleAndLines(l)

	entry := make([]string, 0, len(lines)+1)
	if includeHeading {
		entry = append(entry, drawHeadingStyle.Render("Draw "+strconv.Itoa(drawNumber)))
	}
	entry = append(entry, lines...)

	return title, strings.Join(entry, "\n")
}

func maxLineWidth(blocks []string) int {
	maxWidth := 0
	for _, block := range blocks {
		for line := range strings.SplitSeq(block, "\n") {
			width := lipgloss.Width(line)
			if width > maxWidth {
				maxWidth = width
			}
		}
	}
	if maxWidth < 12 {
		return 12
	}
	return maxWidth
}

func renderDrawCollection(title string, entries []string) string {
	if len(entries) == 0 {
		return ""
	}

	separator := separatorStyle.Render(strings.Repeat("─", maxLineWidth(entries)))
	body := strings.Join(entries, "\n"+separator+"\n")

	return cardStyle.Render(titleStyle.Render(title+" draws") + "\n" + body)
}
