package adapter

import (
	"context"
	"example.com/boiletplate/config"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type AzureBlobAdapter struct {
	client        *azblob.Client
	containerName string
}

func NewAzureBlobAdapter(config *config.Config) (*AzureBlobAdapter, error) {
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", config.AzureBlobStorage.AccountName)

	credential, err := azblob.NewSharedKeyCredential(config.AzureBlobStorage.AccountName, config.AzureBlobStorage.AccountKey)
	if err != nil {
		return nil, err
	}

	client, err := azblob.NewClientWithSharedKeyCredential(serviceURL, credential, nil)
	if err != nil {
		return nil, err
	}
	return &AzureBlobAdapter{client, config.AzureBlobStorage.ContainerName}, nil
}

func (a *AzureBlobAdapter) UploadBlobStorage(file *os.File, blobName string) error {
	_, err := a.client.UploadFile(context.Background(), a.containerName, blobName, file, nil)
	if err != nil {
		return err
	}

	return nil
}
