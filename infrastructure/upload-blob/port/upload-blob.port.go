package port

import "os"

type HandleBlobPort interface {
	UploadBlobStorage(file *os.File, blobName string) error
}
