package logwriter

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	backupTimeFormat        = "2006-01-02T15-04-05.000"
	compressSuffix          = ".gz"
	defaultMaxSize    int64 = 100 // MB
	defaultRotateTime       = time.Hour * 24
)

// FileFormat return file path
type FileFormat interface {
	String() string
}

// FileWriter is an io.WriteCloser that writes to the specified filename
//
type FileWriter struct {
	FileFmt FileFormat
	MaxSize int64
	MaxAge  time.Duration

	size int64
	file *os.File
	mu   sync.Mutex

	millCh    chan bool
	startMill sync.Once
}

var (
	// currentTime exists so it can be mocked by tests
	currentTime = time.Now
	// os_Stat exists so it can be mocked out by tests
	osStat = os.Stat
	// megabyte is the conversion factor between MaxSize and bytes
	megabyte int64 = 1024 * 1024
)

// Write implements io.Writer
func (f *FileWriter) Write(p []byte) (n int, err error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	writeLen := int64(len(p))
	if writeLen > f.max() {
		return 0, fmt.Errorf("write length %d exceeds maximum file size %d", writeLen, f.max())
	}

	if f.file == nil {

	}

	return 1, nil
}

func (f *FileWriter) openExistingOrNew(writeLen int) error {
	return nil
}

func (f *FileWriter) max() int64 {
	if f.MaxSize == 0 {
		return defaultMaxSize * megabyte
	}
	return f.MaxSize * megabyte
}
