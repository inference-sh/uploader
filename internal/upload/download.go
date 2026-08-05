package upload

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/inference-sh/uploader/internal/reader"
)

// DownloadConfig holds the configuration for downloading a file
type DownloadConfig struct {
	URL      string
	DestPath string
}

// Download downloads a file from a URL to a local path.
func Download(cfg *DownloadConfig) error {
	destDir := filepath.Dir(cfg.DestPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("error creating destination directory: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, cfg.URL, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Accept-Encoding", "identity")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error downloading file: %v", err)
	}
	if resp == nil {
		return fmt.Errorf("nil response from download request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("download failed with status %s: %s", resp.Status, string(body))
	}

	tempFile, err := os.CreateTemp(destDir, "download-*.tmp")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %v", err)
	}
	tempPath := tempFile.Name()
	defer os.Remove(tempPath)

	progress := reader.NewProgress(resp.Body, resp.ContentLength, reader.OperationDownload)
	bufWriter := bufio.NewWriter(tempFile)

	if _, err := io.Copy(bufWriter, progress); err != nil {
		tempFile.Close()
		return fmt.Errorf("error writing file: %v", err)
	}

	if err := bufWriter.Flush(); err != nil {
		tempFile.Close()
		return fmt.Errorf("error flushing data: %v", err)
	}
	tempFile.Close()

	if err := os.Rename(tempPath, cfg.DestPath); err != nil {
		return fmt.Errorf("error moving file to destination: %v", err)
	}

	fmt.Printf("\ndownloaded %s to %s\n", cfg.URL, cfg.DestPath)
	return nil
}
