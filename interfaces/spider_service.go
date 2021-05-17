package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpiderService interface {
	Run(id primitive.ObjectID, opts *RunOptions) (err error)
	Clone(id primitive.ObjectID, opts *CloneOptions) (err error)
	Delete(id primitive.ObjectID) (err error)
	Sync(id primitive.ObjectID) (err error)
}
