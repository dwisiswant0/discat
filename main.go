package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/acarl005/stripansi"
)

var wg sync.WaitGroup

func main() {
	var oneLine, verboseMode bool
	var webhookURL, lines string

	flag.StringVar(&webhookURL, "u", "", "Discord Webhook URL")
	flag.BoolVar(&oneLine, "1", false, "Send message line-by-line")
	flag.BoolVar(&verboseMode, "v", false, "Verbose mode")
	flag.Parse()

	webhookEnv := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookEnv != "" {
		webhookURL = webhookEnv
	} else {
		if webhookURL == "" {
			if verboseMode {
				fmt.Println("Discord Webhook URL not set!")
			}
		}
	}

	if !isStdin() {
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()

		fmt.Println(line)
		if oneLine {
			if webhookURL != "" {
				wg.Add(1)
				go disCat(webhookURL, line)
			}
		} else {
			lines += line
			lines += "\n"
		}
	}

	if !oneLine {
		wg.Add(1)
		go disCat(webhookURL, lines)
	}
	wg.Wait()
}

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

type data struct {
	Content string `json:"content"`
}

func disCat(url string, line string) {
	data, _ := json.Marshal(data{Content: stripansi.Strip(line)})
	http.Post(url, "application/json", strings.NewReader(string(data)))
	wg.Done()
}
