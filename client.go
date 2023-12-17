package b2fs

import (
	"context"

	"github.com/Backblaze/blazer/b2"
)

type Client struct {
	client *b2.Client
}

func NewFromB2(client *b2.Client) *Client {
	return &Client{client: client}
}

// NewFS creates a new filesystem for the given bucket.
func (c *Client) NewFS(ctx context.Context, bucket string) (*FS, error) {
	b, err := c.client.Bucket(ctx, bucket)
	if err != nil {
		return nil, err
	}

	return &FS{
		ctx: ctx,
		b:   b,
	}, nil
}
