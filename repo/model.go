package repo

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Setting struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	USetting  `bson:"config"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type (
	USettingResponse struct {
		Header map[string]string `json:"header" bson:"header"`
		Body   interface{}       `json:"body" bson:"body"`
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
	e := new(Setting)
	e.ID = primitive.NewObjectID()
	e.CreatedAt = now
	e.UpdatedAt = now
	e.Bind(NewUSetting())
	return e
}
func (e *Setting) Bind(ue interface{}) {

	switch ue.(type) {
	case *USetting:
		d := ue.(*USetting)
		e.URI = strings.ReplaceAll(strings.TrimSpace(d.URI), " ", "%20")
		e.Method = strings.ToUpper(strings.TrimSpace(d.Method))
		e.USettingResponse.Header = d.USettingResponse.Header
		e.USettingResponse.Body = d.USettingResponse.Body
		e.USettingResponse.Code = d.USettingResponse.Code
	}
}
func (e *Setting) SetModifyTimestamp() {
	e.UpdatedAt = time.Now()
}

func NewUSetting() *USetting {
	ue := new(USetting)
	return ue
}
