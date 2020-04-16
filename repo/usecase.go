package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	AddSetting(context.Context, Setting) (Setting, error)
	GetSetting(context.Context, primitive.ObjectID) (Setting, error)
	GetSettingList(context.Context) ([]Setting, error)
	DeleteSetting(context.Context, primitive.ObjectID) error
	UpdateSetting(context.Context, primitive.ObjectID, USetting) (Setting, error)
	GetSettingByMethodAndURI(context.Context, string, string) (Setting, error)
}
