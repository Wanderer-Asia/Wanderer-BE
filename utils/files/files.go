package files

import (
	"context"
	"io"
)

type Cloud interface {
	Upload(ctx context.Context, folder string, Raw io.Reader) (*string, error)
}
