package database

import "context"

type Seq interface {
	Malloc(ctx context.Context, conversationID string, size int64) ([]int64, error)
	GetMaxSeq(ctx context.Context, conversationID string) (int64, error)
	GetMinSeq(ctx context.Context, conversationID string) (int64, error)
	SetMinSeq(ctx context.Context, conversationID string, seq int64) error
}
