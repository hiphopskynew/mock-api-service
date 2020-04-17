package repo

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Settings []Setting

type Setting struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	USetting  `bson:"config"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type (
	USettingResponse struct {
		Header map[string]string `json:"header" bson:"header"`
		Body   primitive.M       `json:"body" bson:"body"`
		Code   int               `json:"code" bson:"code"`
	}
	USetting struct {
		Method           string           `json:"method" bson:"method"`
		URI              string           `json:"uri" bson:"uri"`
		USettingResponse USettingResponse `json:"response" bson:"response"`
	}
)

func NewSetting() *Setting {
	now := time.Now()
	s := new(Setting)
	s.ID = primitive.NewObjectID()
	s.CreatedAt = now
	s.UpdatedAt = now
	s.Bind(NewUSetting())
	return s
}
func (s *Setting) Bind(ue interface{}) {
	switch ue.(type) {
	case *USetting:
		d := ue.(*USetting)
		s.URI = strings.ReplaceAll(strings.TrimSpace(d.URI), " ", "%20")
		s.Method = strings.ToUpper(strings.TrimSpace(d.Method))
		s.USettingResponse.Header = d.USettingResponse.Header
		s.USettingResponse.Body = d.USettingResponse.Body
		s.USettingResponse.Code = d.USettingResponse.Code
	}
}
func (s *Setting) SetModifyTimestamp() {
	s.UpdatedAt = time.Now()
}

func NewUSetting() *USetting {
	ue := new(USetting)
	return ue
}
