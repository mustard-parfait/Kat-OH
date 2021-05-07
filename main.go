package main

import (
		"fmt"
		"strings"
		"os"

		"github.com/charmbracelet/lipgloss"
		"golang.org/x/term"
)

func main() {
		_, _, _ = term.GetSize(int(os.Stdout.Fd()))
		ketch := strings.Builder{}
		styledFetch := lipgloss.NewStyle().Padding(1, 2, 1, 2)

		// ketch.WriteString(RenderBar() + "\n\n")
		ketch.WriteString(RenderAWBar() + "\n\n")
		// ketch.WriteString(RenderVIMBar() + "\n\n")
		ketch.WriteString(RenderFetch() + "\n\n")
		ketch.WriteString(RenderQuote() + "\n\n")
		// ketch.WriteString(RenderStatusBar())
		ketch.WriteString(RenderSpotBar())

		fmt.Println(styledFetch.Render(ketch.String()))
}
