package main

import (
		"strings"

		"github.com/charmbracelet/lipgloss"
)

func max(a, b int) int {
		if a > b {
				return a
		} else {
				return b
		}
}

func RenderBar() string {
		highlight := lipgloss.AdaptiveColor{Light: "#8E3ED7", Dark: "#7F21D7"}
		width := 72

		activeTabBorder := lipgloss.Border {
				Top:         "─",
				Bottom:      " ",
				Left:        "│",
				Right:       "│",
				TopLeft:     "╭",
				TopRight:    "╮",
				BottomLeft:  "┘",
				BottomRight: "└",
		}

		tabBorder := lipgloss.Border {
				Top:         "─",
				Bottom:      "─",
				Left:        "│",
				Right:       "│",
				TopLeft:     "╭",
				TopRight:    "╮",
				BottomLeft:  "┴",
				BottomRight: "┴",
		}

		tab := lipgloss.NewStyle().
				Border(tabBorder, true).
				BorderForeground(highlight).
				Padding(0, 1)

		activeTab := tab.Copy().Border(activeTabBorder, true)

		tabGap := tab.Copy().
				BorderTop(false).
				BorderLeft(false).
				BorderRight(false)

		row := lipgloss.JoinHorizontal(
				lipgloss.Top,
				activeTab.Render("Fetch"),
				tab.Render("Cats"),
				tab.Render("Are"),
				tab.Render("Cute"),
		)

		gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

		return row
}
