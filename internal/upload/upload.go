package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/rhnvrm/simples3"
	"inference.sh/uploader/internal/reader"
)

// Config holds the configuration for uploading a file
type Config struct {
	SourcePath      string
	DestPath        string
	Endpoint        string
	Region          string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	URL             string
}

// Upload uploads a file to S3-compatible storage using a presigned PUT URL.
func Upload(cfg *Config) error {
	client := simples3.New(cfg.Region, cfg.AccessKeyID, cfg.AccessKeySecret)
	if cfg.Endpoint != "" {
		client.SetEndpoint(cfg.Endpoint)
	}

	file, err := os.Open(cfg.SourcePath)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting file info: %v", err)
	}

	presignedURL := client.GeneratePresignedURL(simples3.PresignedInput{
		Bucket:        cfg.Bucket,
		ObjectKey:     cfg.DestPath,
		Method:        "PUT",
		ExpirySeconds: int((10 * time.Minute).Seconds()),
	})

	progress := reader.NewProgress(file, fileInfo.Size(), reader.OperationUpload)

	req, err := http.NewRequest(http.MethodPut, presignedURL, progress)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.ContentLength = fileInfo.Size()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error uploading file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status %s: %s", resp.Status, string(body))
	}

	dest := cfg.DestPath
	if cfg.URL != "" {
		dest = cfg.URL + "/" + cfg.DestPath
	}
	fmt.Printf("\nuploaded %s to %s\n", cfg.SourcePath, dest)
	return nil
}
