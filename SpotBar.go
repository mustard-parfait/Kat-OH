package main

import (
		"github.com/charmbracelet/lipgloss"
)

func RenderSpotBar() string {
		width := 72

		barStyle := lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
				Background(lipgloss.AdaptiveColor{Light: "#4B4B7E", Dark: "#202034"})

		itemStyle := lipgloss.NewStyle().
				Inherit(barStyle).
				Padding(0, 1)

		songStyle := lipgloss.NewStyle().
				Inherit(barStyle).
				Foreground(lipgloss.AdaptiveColor{Light: "#DB6BF3", Dark: "#CB0FF3"}).
				Padding(0, 1)

		_ = songStyle.Copy().Render("")
		songInfo := songStyle.Copy().Render(GetSpot())

		w := lipgloss.Width
		albumInfo := songStyle.Copy().
				Width(width - w(songInfo)).
				Align(lipgloss.Right).
				Render(GetSpotAlb())

		song := lipgloss.JoinHorizontal(lipgloss.Top, songInfo, albumInfo)

		return itemStyle.Width(72).Render(song)
}
