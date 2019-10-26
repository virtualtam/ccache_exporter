package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	ccache "github.com/virtualtam/ccacheparser"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Mode()&os.ModeNamedPipe == 0 {
		// TODO add flags, read from stdin / file(s)
		// TODO add help
		panic("No data piped to stdin")
	}

	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}
	stats := ccache.Statistics{}
	stats.Parse(text)
	statsJson, _ := json.Marshal(stats)
	fmt.Println(string(statsJson))
}
