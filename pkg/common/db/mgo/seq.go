package mgo

import (
	"context"
	"errors"
	"github.com/OpenIMSDK/tools/errs"
	"github.com/OpenIMSDK/tools/mgoutil"
	"github.com/openimsdk/open-im-server/v3/pkg/common/db/table/relation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewSeq(db *mongo.Database) (*SeqMongo, error) {
	coll := db.Collection("seq")

	return &SeqMongo{coll: coll}, nil
}

type SeqMongo struct {
	coll *mongo.Collection
}

func (s *SeqMongo) MallocSeq(ctx context.Context, conversationID string, size int64) (int64, error) {
	if size <= 0 {
		return 0, errors.New("size must be greater than 0")
	}
	filter := map[string]any{"conversation_id": conversationID}
	update := map[string]any{
		"$inc": map[string]any{"max_seq": size},
		"$set": map[string]any{"min_seq": 1},
	}
	opt := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After).SetProjection(map[string]any{"_id": 0, "max_seq": 1})
	return mgoutil.FindOneAndUpdate[int64](ctx, s.coll, filter, update, opt)
}

func (s *SeqMongo) Malloc(ctx context.Context, conversationID string, size int64) ([]int64, error) {
	seq, err := s.MallocSeq(ctx, conversationID, size)
	if err != nil {
		return nil, err
	}
	seqs := make([]int64, 0, size)
	for i := seq - size + 1; i <= seq; i++ {
		seqs = append(seqs, i)
	}
	return seqs, nil
}

func (s *SeqMongo) GetMaxSeq(ctx context.Context, conversationID string) (int64, error) {
	seq, err := mgoutil.FindOne[int64](ctx, s.coll, bson.M{"conversation_id": conversationID}, options.FindOne().SetProjection(map[string]any{"_id": 0, "max_seq": 1}))
	if err == nil {
		return seq, nil
	} else if errs.Unwrap(err) == mongo.ErrNoDocuments {
		return 0, nil
	} else {
		return 0, err
	}
}

func (s *SeqMongo) GetMinSeq(ctx context.Context, conversationID string) (int64, error) {
	seq, err := mgoutil.FindOne[int64](ctx, s.coll, bson.M{"conversation_id": conversationID}, options.FindOne().SetProjection(map[string]any{"_id": 0, "min_seq": 1}))
	if err == nil {
		return seq, nil
	} else if errs.Unwrap(err) == mongo.ErrNoDocuments {
		return 0, nil
	} else {
		return 0, err
	}
}

func (s *SeqMongo) SetMinSeq(ctx context.Context, conversationID string, seq int64) error {
	return mgoutil.UpdateOne(ctx, s.coll, bson.M{"conversation_id": conversationID}, bson.M{"$set": bson.M{"min_seq": seq}}, false)
}

func (s *SeqMongo) GetConversation(ctx context.Context, conversationID string) (*relation.Seq, error) {
	return mgoutil.FindOne[*relation.Seq](ctx, s.coll, bson.M{"conversation_id": conversationID})
}
