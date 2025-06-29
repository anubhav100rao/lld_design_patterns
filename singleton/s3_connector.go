package singleton

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Client *s3.Client
	once     sync.Once
	initErr  error
)

// S3 returns a singleton S3 client, loading config only once.
func S3(ctx context.Context) (*s3.Client, error) {
	once.Do(func() {
		cfg, initErr := config.LoadDefaultConfig(ctx)
		if initErr != nil {
			return
		}
		s3Client = s3.NewFromConfig(cfg)
	})
	return s3Client, initErr
}
