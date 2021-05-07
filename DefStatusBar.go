package main

import (
		"github.com/charmbracelet/lipgloss"
)

func RenderStatusBar() string {
		width := 72

		statusNugget := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFDF5")).
				Padding(0, 1)

		statusBarStyle := lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
				Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

		statusStyle := lipgloss.NewStyle().
				Inherit(statusBarStyle).
				Foreground(lipgloss.Color("#FFFDF5")).
				Background(lipgloss.Color("#9F03CF")).
				Padding(0, 1).
				MarginRight(1)

		encodingStyle := statusNugget.Copy().
				Background(lipgloss.Color("#d35d6e")).
				Align(lipgloss.Right)

		statusText := lipgloss.NewStyle().Inherit(statusBarStyle)

		catStyle := statusNugget.Copy().Background(lipgloss.Color("#6124DF"))

		statusKey := statusStyle.Render("STATUS")
		encoding := encodingStyle.Render("UTF-8")
		fishCake := catStyle.Render(" Cat Poop")

		w := lipgloss.Width
		statusVal := statusText.Copy().
				Foreground(lipgloss.Color("#A599E9")).
				Width(width - w(statusKey) - w(encoding) - w(fishCake)).
				Render("Fetching")

		bar := lipgloss.JoinHorizontal(lipgloss.Top,
				statusKey,
				statusVal,
				encoding,
				fishCake,
		)

		return statusBarStyle.Width(width).Render(bar)
}
