package repo

import (
	"context"

	"github.com/Vantuan1606/app-test/config"

	"github.com/Vantuan1606/app-test/domain"
	"github.com/Vantuan1606/app-test/service/database"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoCommentRepo struct {
	c *mongo.Collection
}

func NewMongoCommentRepo() domain.ICommentRepo {
	c := config.GetConfig()
	mongo := database.NewMongoService()

	return &mongoCommentRepo{
		c: mongo.Client.Database(c.Mongo.Database).Collection("Comment"),
	}
}

func (m *mongoCommentRepo) Find(ctx context.Context, conditions map[string]interface{}, options ...*options.FindOptions) ([]*domain.Comment, error) {
	cursor, err := m.c.Find(ctx, conditions, options...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comments []*domain.Comment
	for cursor.Next(ctx) {
		var comment *domain.Comment
		if err := cursor.Decode(&comment); err != nil {
			logrus.WithError(err).Error("[DECODE] fail at ", comment.ID)
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
