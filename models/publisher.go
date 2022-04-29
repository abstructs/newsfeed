package models

import "context"

type IPublisher interface {
	Publish(ctx context.Context, key string, value string) error
}
