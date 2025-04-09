package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Vantuan1606/app-test/comment"
	"github.com/Vantuan1606/app-test/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type commentUsecase struct {
	commentRepo    domain.ICommentRepo
	contextTimeout time.Duration
}

func NewCommentUsecase(pr domain.ICommentRepo, timeOut time.Duration) domain.ICommentUsecase {
	return &commentUsecase{
		commentRepo:    pr,
		contextTimeout: timeOut,
	}
}
func (cms *commentUsecase) GetComment(c context.Context, cmtID string) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, cms.contextTimeout)
	defer cancel()

	cmtObjectID, err := primitive.ObjectIDFromHex(cmtID)
	log.Println("hashtagObjectID", cmtObjectID)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("Invalid ID")
	}

	conditions := bson.M{"_id": cmtObjectID}
	comment, err := cms.commentRepo.FindOne(ctx, conditions)

	if err != nil {
		logrus.WithField("Data", fmt.Sprintf("%v", conditions)).WithError(err).Error("Comment not found")
		return nil, err
	}

	return comment, nil
}
func (cms *commentUsecase) List(c context.Context, input *comment.ListCommentInput) ([]*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(c, cms.contextTimeout)
	defer cancel()

	conditions := bson.M{}

	options := options.Find()
	options.SetLimit(int64(*input.Limit)) // Giới hạn số lượng user trả về

	comments, err := cms.commentRepo.Find(ctx, conditions, options)
	if err != nil {
		logrus.WithError(err).Error("Get list Comment failed")
		return nil, err
	}

	return comments, nil
}
