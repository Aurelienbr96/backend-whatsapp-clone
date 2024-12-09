package mock

import (
	"fmt"
	"os"
)

type MockBlobAdapter struct {
}

func NewMockBlobAdapter() *MockBlobAdapter {
	return &MockBlobAdapter{}
}

func (a *MockBlobAdapter) UploadBlobStorage(file *os.File, blobName string) error {
	fmt.Printf("Uploading file %s to blob %s\n", file.Name(), blobName)
	return nil
}
