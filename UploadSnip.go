package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

var languageMap = map[string]string{
	".sh":     "shell",
	".c":      "c",
	".cs":     "csharp",
	".css":    "css",
	".elm":    "elm",
	".erl":    "erlang",
	".go":     "go",
	".hs":     "haskell",
	".html":   "html",
	".java":   "java",
	".js":     "javascript",
	".json":   "json",
	".jsx":    "jsx",
	".kt":     "kotlin",
	".lua":    "lua",
	".md":     "markdown",
	".ps1":    "powershell",
	".php":    "php",
	".py":     "python",
	".r":      "r",
	".rb":     "ruby",
	".rs":     "rust",
	".scala":  "scala",
	".sol":    "solidity",
	".sql":    "sql",
	".swift":  "swift",
	".svelte": "svelte",
	".toml":   "toml",
	".ts":     "typescript",
	".tsx":    "tsx",
	".vue":    "vue",
	".xml":    "xml",
	".yaml":   "yaml",
	".yml":    "yaml",
}

func UploadSnip(file string, name string) (ResponseData, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return ResponseData{}, fmt.Errorf("failed to read file: %w", err)
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	var fileName string

	if name == "nil" {
		fileName = filepath.Base(file)
	} else {
		fileName = name
	}

	extension := strings.ToLower(filepath.Ext(file))

	lang, ok := languageMap[extension]
	if !ok {
		lang = "plaintext"
	}

	payload := UploadPayload{
		Content: string(data),
		Name:    fileName,
		Lang:    lang,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return ResponseData{}, fmt.Errorf("failed to marshal the payload: %w", err)
	}

	req, err := http.NewRequest("POST", "https://www.snippets.so/api/upload", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return ResponseData{}, fmt.Errorf("failed to create the request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseData{}, fmt.Errorf("failed to send the request: %w", err)
	}
	defer resp.Body.Close()

	var response ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return ResponseData{}, fmt.Errorf("failed to decode the response: %w", err)
	}
	s.Stop()

	fmt.Println("https://snippets.so/snip/" + response.Slug)

	return response, nil
}
