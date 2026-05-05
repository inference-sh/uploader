package reader

import (
	"fmt"
	"io"
	"time"
)

// Operation represents the type of operation being tracked
type Operation string

const (
	OperationEncryption Operation = "encrypting"
	OperationUpload     Operation = "uploading"
	OperationDownload   Operation = "downloading"
)

// Progress wraps an io.Reader with progress reporting
type Progress struct {
	reader     io.Reader
	total      int64
	processed  int64
	lastUpdate time.Time
	operation  Operation
}

// NewProgress creates a new progress reader
func NewProgress(reader io.Reader, total int64, operation Operation) *Progress {
	return &Progress{
		reader:     reader,
		total:      total,
		lastUpdate: time.Now(),
		operation:  operation,
	}
}

// Read implements io.Reader
func (pr *Progress) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	if n > 0 {
		pr.processed += int64(n)
		if time.Since(pr.lastUpdate) > 100*time.Millisecond {
			percentage := float64(pr.processed) / float64(pr.total) * 100
			fmt.Printf("\r%s... %.1f%% (%d/%d bytes)", pr.operation, percentage, pr.processed, pr.total)
			pr.lastUpdate = time.Now()
		}
	}
	return n, err
}
