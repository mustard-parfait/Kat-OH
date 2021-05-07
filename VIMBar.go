package main

import (
		"github.com/charmbracelet/lipgloss"
)

func RenderVIMBar() string {
		barStyle := lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
				Background(lipgloss.AdaptiveColor{Light: "#4B4B7E", Dark: "#202034"})

		itemStyle := lipgloss.NewStyle().
				Inherit(barStyle).
				Padding(0, 1)

		activeTabStyle := lipgloss.NewStyle().
				Inherit(barStyle).
				Foreground(lipgloss.Color("#37355E")).
				Background(lipgloss.Color("#A599E9")).
				Padding(0, 1).
				Render("1 Fetch")

		nonActiveTab := lipgloss.NewStyle().
				Inherit(barStyle).
				Foreground(lipgloss.Color("#9B90DC")).
				Background(lipgloss.Color("#2B2B5A")).
				Padding(0, 1)

		eBox := lipgloss.NewStyle().
				Inherit(barStyle).
				Foreground(lipgloss.Color("#9B90DC")).
				Background(lipgloss.Color("#2B2B5A"))

		emptyTab1 := nonActiveTab.Copy().Render("2 Cats")
		emptyTab2 := nonActiveTab.Copy().Render("3 Are")
		emptyTab3 := nonActiveTab.Copy().Render("4 Cute")

		pipe := eBox.Copy().Render("|")

		closeButton := lipgloss.NewStyle().
				Inherit(barStyle).
				Foreground(lipgloss.Color("#37355E")).
				Background(lipgloss.Color("#A599E9")).
				Padding(0, 1).
				Render("X")

		w := lipgloss.Width
		emptySpace := lipgloss.NewStyle().
				Inherit(barStyle).
				Width(72 - w(activeTabStyle) - w(emptyTab1) - w(emptyTab2) - w(emptyTab3) - w(pipe) * 3 - w(closeButton)).
				Align(lipgloss.Right).
				Foreground(lipgloss.Color("#A599E9")).
				Padding(0, 1).
				Render("")

		tabs := lipgloss.JoinHorizontal(lipgloss.Top, activeTabStyle, emptyTab1, pipe, emptyTab2, pipe, emptyTab3, emptySpace, closeButton)

		return itemStyle.Width(72).Render(tabs)
}
