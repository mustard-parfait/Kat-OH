package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func GetBat() string {
		per, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")

		if err != nil {
				panic(err)
		}

		stat, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/status")

		if err != nil {
				panic(err)
		}

		if string(stat) == "Not charging\n" || string(stat) == "Charging\n" {
				return strings.TrimSuffix(string(per), "\n") + "%" + " plug"
		} else {
				return strings.TrimSuffix(string(per), "\n") + "%" + " batr"
		}
}

func GetSpotAlb() string {
		cmdArr := []string {"--player=spotify", "metadata", "--format", "{{ album }}"}
		cmd := exec.Command("playerctl", cmdArr...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		output, err := cmd.Output()

		if err != nil {
			output = []byte("Unknown")
		}

		info := strings.TrimSuffix(string(output), "\n")
		
		if len(info) >= 20 {
			info = info[:18]
		}

		spotInfo := lipgloss.NewStyle().Height(1).Width(20).Render(info)

		return spotInfo
}

func GetSpot() string {
		cmdArr := []string {"--player=spotify", "metadata", "--format", "{{ title }} - {{ artist }}"}
		cmd := exec.Command("playerctl", cmdArr...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		output, err := cmd.Output()

		if err != nil {
			output = []byte("Spotify - Not Available")
		}

		info := strings.TrimSuffix(string(output), "\n")
		
		if len(info) >= 40 {
			info = info[:38]
		}

		spotInfo := lipgloss.NewStyle().Height(1).Width(40).Render(info)

		return spotInfo
}

type Quotes struct {
		id string
		tags []string
		Content string
		Author string
		authorSlug string
		length int
}

func GetQuotes() Quotes {
		resp, err := http.Get("https://api.quotable.io/random?maxLength=50")

		if err != nil {
				panic(err)
		}

		quotesJson, err := ioutil.ReadAll(resp.Body)

		if err != nil {
				panic(err)
		}

		var quote Quotes
		err = json.Unmarshal((quotesJson), &quote)

		if err != nil {
				panic(err)
		}

		return quote
}

func GetUser() string {
		return os.Getenv("USER")
}

func GetHost() string {
		host, _ := os.Hostname()
		return host
}

func GetOS() string {
		if runtime.GOOS == "linux" {
			cmd := exec.Command("lsb_release", "-sd")
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr

			os, err := cmd.Output()

			if err != nil {
					panic(err)
			}

			if strings.Contains(string(os), "Void") {
					return "void"
			} else if strings.Contains(string(os), "Arch") {
					return "arch"
			} else if strings.Contains(string(os), "Ubuntu") {
					return "ubuntu"
			} else {
				return "unknown"
			}
		} else if runtime.GOOS == "darwin" {
			cmd := exec.Command("sw_vers", "-productVersion")
			wc := exec.Command("awk", "-F", "[.]", "{print $2}")

			read, write := io.Pipe()
			var buf bytes.Buffer

			cmd.Stdout = write
			wc.Stdin = read
			wc.Stdout = &buf

			cmd.Start()
			wc.Start()

			cmd.Wait()
			write.Close()

			wc.Wait()
			read.Close()

			if buf.String() == "13\n" || buf.String() == "12\n" {
				return "sierra"
			} else if buf.String() == "14\n" {
				return "mojave"
			} else if buf.String() == "15\n" {
				return "big sur"
			}
			
			return "macos"
		} else {
				return "unknown"
		}
}

func GetKer() string {
		cmd := exec.Command("uname", "-r")
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		ker, err := cmd.Output()

		if err != nil {
				panic(err)
		}

		kernel := strings.Split(strings.TrimSuffix(string(ker), "\n"), ".")

		return kernel[0] + "." + kernel[1]
}

func GetShell() string {
		return path.Base(string(os.Getenv("SHELL")))
}

func GetPks() string {
		if GetOS() == "void" {
				cmd := exec.Command("xbps-query", "-l")
				wc := exec.Command("wc", "-l")

				read, write := io.Pipe()
				var buf bytes.Buffer

				cmd.Stdout = write
				wc.Stdin = read
				wc.Stdout = &buf

				cmd.Start()
				wc.Start()

				cmd.Wait()
				write.Close()

				wc.Wait()
				read.Close()

				return strings.TrimSuffix(buf.String(), "\n")
		} else if runtime.GOOS == "darwin" {
				cmd := exec.Command("brew", "list")
				wc := exec.Command("wc", "-l")

				read, write := io.Pipe()
				var buf bytes.Buffer

				cmd.Stdout = write
				wc.Stdin = read
				wc.Stdout = &buf

				cmd.Start()
				wc.Start()

				cmd.Wait()
				write.Close()

				wc.Wait()
				read.Close()

				return strings.TrimSpace(buf.String())
		}

		return "calc"
}
