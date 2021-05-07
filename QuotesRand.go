package main

import (
		"github.com/charmbracelet/lipgloss"
)

func RenderQuote() string {
		width := 72

		dialogBoxStyle := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground((lipgloss.Color("#B071CC"))).
				Padding(1, 0).
				BorderTop(true).
				BorderBottom(true).
				BorderLeft(true).
				BorderRight(true)

		quote := GetQuotes()

		cnt := lipgloss.NewStyle().Width(60).Align(lipgloss.Center).Render(quote.Content)
		atr := lipgloss.NewStyle().Width(20).Align(lipgloss.Center).Render(quote.Author)

		ui := lipgloss.JoinVertical(lipgloss.Center, cnt, atr)

		dialog := lipgloss.Place(width, 10,
				lipgloss.Center, lipgloss.Center,
				dialogBoxStyle.Render(ui),
				lipgloss.WithWhitespaceChars(""),
				lipgloss.WithWhitespaceForeground(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}),
		)

		return dialog
}
