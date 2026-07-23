package main

import (
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/onyx-and-iris/lottery-cli"
)

func formatNumberList(numbers []int) string {
	parts := make([]string, len(numbers))
	for i, n := range numbers {
		parts[i] = strconv.Itoa(n)
	}
	return strings.Join(parts, "  ")
}

// nolint:misspell
func renderDraw(l lottery.Lottery) string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	numbersStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))
	specialStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("11"))

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

	body := strings.Join(lines, "\n")
	card := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")).
		Padding(0, 1).
		MarginTop(1)

	return card.Render(titleStyle.Render(title+" draw") + "\n" + body)
}
