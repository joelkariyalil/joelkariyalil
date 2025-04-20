package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	username := "joelkariyalil"
	url := fmt.Sprintf("https://komarev.com/ghpvc/?username=%s&style=flat-square", username)

	// Step 1: Fetch SVG
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	re := regexp.MustCompile(`<text[^>]*>(\d+)</text>`)
	match := re.FindAllStringSubmatch(string(body), -1)

	if len(match) == 0 {
		fmt.Println("View count not found.")
		return
	}

	count := match[len(match)-1][1]
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	file, err := os.OpenFile("assets/views.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	info, _ := file.Stat()
	if info.Size() == 0 {
		writer := csv.NewWriter(file)
		writer.Write([]string{"timestamp", "views"})
		writer.Flush()
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{timestamp, count})
	writer.Flush()

	fmt.Println("Logged:", timestamp, count)
}
