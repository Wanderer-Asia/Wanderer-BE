package files

import (
	"context"
	"io"
	"wanderer/config"

	cld "github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func NewCloudinary(config config.Cloudinary) (Cloud, error) {
	client, err := cld.NewFromParams(config.CloudName, config.ApiKey, config.ApiSecret)
	if err != nil {
		return nil, err
	}

	return &cloudinary{
		config: config,
		client: client,
	}, nil
}

type cloudinary struct {
	config config.Cloudinary
	client *cld.Cloudinary
}

func (cloud *cloudinary) Upload(ctx context.Context, folder string, Raw io.Reader) (*string, error) {
	UniqueFilename := true
	res, err := cloud.client.Upload.Upload(ctx, Raw, uploader.UploadParams{
		UniqueFilename: &UniqueFilename,
		Folder:         folder,
	})

	if err != nil {
		return nil, err
	}

	return &res.SecureURL, nil
}
