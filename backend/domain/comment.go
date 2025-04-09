package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Content   string             `json:"content" bson:"content"`
	HagstagId string             `json:"hagstagId" bson:"hagstagId"`
	Status    int                `json:"status" bson:"Status,omitempty"`
}

type CommentsRequest struct {
	ID string `json:"id"`
}

type ICommentUsecase interface {
	List(context.Context, *Comment.ListCommentInput) ([]*Comment, error)
}

type ICommentRepo interface {
	FindOne(context.Context, map[string]interface{}) (*Comment, error)
	Find(context.Context, map[string]interface{}, ...*options.FindOptions) ([]*Comment, error)
	// Count(ctx context.Context, conditions interface{}) (int64, error)
}
