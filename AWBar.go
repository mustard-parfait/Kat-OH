package main

import (
		"time"

		"github.com/charmbracelet/lipgloss"
)

func RenderAWBar() string {
		width := 72
		barStyle := lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
				Background(lipgloss.AdaptiveColor{Light: "#4B4B7E", Dark: "#202034"})

		tagStyle := lipgloss.NewStyle().
				Inherit(barStyle).
				Padding(0, 1)

		activeIcon := lipgloss.NewStyle().Inherit(barStyle).Width(2).Foreground(lipgloss.Color("#EE08A1")).Render("")

		nonActiveIcons := lipgloss.NewStyle().Inherit(barStyle).Width(2).Foreground(lipgloss.Color("#4F4F49")).Render("")

		icons := lipgloss.JoinHorizontal(lipgloss.Top, activeIcon, nonActiveIcons, nonActiveIcons, nonActiveIcons)

		currentTime := time.Now().UTC()
		timeText := lipgloss.NewStyle().Inherit(barStyle).Margin(0, 1).Foreground(lipgloss.Color("#45ADB1")).Render(currentTime.Format("Mon Jan 02"))

		batText := lipgloss.NewStyle().Inherit(barStyle).Margin(0, 1).Foreground(lipgloss.Color("#4574B1")).Render(GetBat())

		w := lipgloss.Width

		windowText := lipgloss.NewStyle().Inherit(barStyle).
				Foreground(lipgloss.Color("#5045B1")).
				Align(lipgloss.Center).
				Width(width - w(icons) - w(timeText) - w(batText) - 3).
				Render("Fetch")

		barItems := lipgloss.JoinHorizontal(lipgloss.Top, tagStyle.Padding(0, 1).Render(icons), windowText, batText, timeText)

		return tagStyle.Width(width).Render(barItems)
}
