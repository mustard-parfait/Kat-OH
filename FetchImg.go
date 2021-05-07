package main

import (
		"github.com/charmbracelet/lipgloss"
)

func RenderFetch() string {
		image := "                          ,\n" +
				 "   ,-.       _,---._ __  / \\\n" +
				 "  /  )    .-'       `./ /   \\\n" +
				 " (  (   ,'            `/    /|\n" +
				 "  \\  `-\"             \\'\\   / |\n" +
				 "   `.              ,  \\ \\ /  |\n" +
				 "    /`.          ,'-`----Y   |\n" +
				 "   (            ;        |   '\n"+
				 "   |  ,-.    ,-'         |  /\n"+
				 "   |  | (   |        hjw | /\n"+
				 "   )  |  \\  `.___________|/\n" +
				 "   `--'   `--'\n"

		img := lipgloss.NewStyle().MarginRight(12).Foreground(lipgloss.AdaptiveColor{Light: "#DB6BF3", Dark: "#CB0FF3"}).Render(image)

		user := lipgloss.NewStyle().Align(lipgloss.Left).Foreground(lipgloss.Color("#9F03CF")).Render(GetUser())
		at := lipgloss.NewStyle().Align(lipgloss.Left).Bold(true).Render("@")
		host := lipgloss.NewStyle().Align(lipgloss.Left).Foreground(lipgloss.Color("#844E94")).Render(GetHost())

		OS := lipgloss.NewStyle().Align(lipgloss.Left).Width(9).Bold(true).Foreground(lipgloss.Color("#7633D7")).Render("os: ")
		osName := lipgloss.NewStyle().Align(lipgloss.Left).Foreground(lipgloss.Color("#4D425E")).Render(GetOS())

		ker := lipgloss.NewStyle().Align(lipgloss.Left).Width(9).Bold(true).Foreground(lipgloss.Color("#672DBD")).Render("ker: ")
		kerName := lipgloss.NewStyle().Align(lipgloss.Left).Foreground(lipgloss.Color("#5F5073")).Render(GetKer())

		shell := lipgloss.NewStyle().Align(lipgloss.Left).Width(9).Bold(true).Foreground(lipgloss.Color("#57269E")).Render("shell: ")
		shellName := lipgloss.NewStyle().Align(lipgloss.Left).Foreground(lipgloss.Color("#6A5A82")).Render(GetShell())

		pac := lipgloss.NewStyle().Align(lipgloss.Left).Width(9).Bold(true).Foreground(lipgloss.Color("#4A2188")).Render("pks: ")
		pacNum := lipgloss.NewStyle().Align(lipgloss.Left).Foreground(lipgloss.Color("#7B6896")).Render(GetPks())

		uhBox := lipgloss.JoinHorizontal(lipgloss.Left, user, at, host)

		osBox := lipgloss.JoinHorizontal(lipgloss.Left, OS, osName)

		kerBox := lipgloss.JoinHorizontal(lipgloss.Left, ker, kerName)

		shellBox := lipgloss.JoinHorizontal(lipgloss.Left, shell, shellName)

		pacBox := lipgloss.JoinHorizontal(lipgloss.Left, pac, pacNum)

		uhStyle := lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				Width(20).
				Align(lipgloss.Center).
				BorderBottom(true)

		restStyle := lipgloss.NewStyle().
				Width(20).
				Align(lipgloss.Center)

		fetch := lipgloss.JoinVertical(lipgloss.Left,
				uhStyle.Render(uhBox),
				restStyle.Render(osBox),
				restStyle.Render(kerBox),
				restStyle.Render(shellBox),
				restStyle.Render(pacBox),
		)

		fetchStyle := lipgloss.NewStyle().Align(lipgloss.Center).MarginTop(2)

		row := lipgloss.JoinHorizontal(lipgloss.Top, img, fetchStyle.Render(fetch))

		return row
}
