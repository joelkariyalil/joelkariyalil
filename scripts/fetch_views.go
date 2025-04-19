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

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(string(resp.Status))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	re := regexp.MustCompile(`<text x="90.2" y="14">(\d+)</text>`)
	match := re.FindStringSubmatch(string(body))
	if len(match) < 2 {
		fmt.Println("View count not found.")
		return
	}
	count := match[1]
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Create or append to views.csv
	file, err := os.OpenFile("views.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Add header if new
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
