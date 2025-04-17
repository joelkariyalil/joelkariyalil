// scripts/fetch_views.go
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	re := regexp.MustCompile(`\d+`)
	match := re.Find(body)
	if match == nil {
		fmt.Println("No number found in badge.")
		return
	}

	count := string(match)
	date := time.Now().Format("2006-01-02")

	// Create or append to views.csv
	file, err := os.OpenFile("views.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Check if file is new → add header
	info, _ := file.Stat()
	if info.Size() == 0 {
		writer := csv.NewWriter(file)
		writer.Write([]string{"date", "views"})
		writer.Flush()
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{date, count})
	writer.Flush()

	fmt.Println("Logged:", date, count)
}
