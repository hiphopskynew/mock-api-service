package repo

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type settingInSvcMemory map[string]Setting

func NewSettingInServiceMemory() settingInSvcMemory {
	return settingInSvcMemory{}
}
func (s settingInSvcMemory) SetA(settings ...Setting) {
	for _, setting := range settings {
		s.Set(setting.ID.Hex(), setting)
	}
}
func (s settingInSvcMemory) Set(k string, settings ...Setting) {
	for _, setting := range settings {
		s[k] = setting
	}
}
func (s settingInSvcMemory) Unset(k string) {
	delete(s, k)
}
func (s settingInSvcMemory) Get(k string) *Setting {
	if v, ok := s[k]; ok {
		return &v
	}
	return nil
}
func (s settingInSvcMemory) GetAll() *[]Setting {
	all := []Setting{}
	for _, setting := range s {
		all = append(all, setting)
	}
	if len(all) > 0 {
		return &all
	} else {
		return nil
	}
}

type Setting struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	USetting  `bson:"config"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type (
	USettingResponse struct {
		Format int               `json:"format" bson:"format"`
		Header map[string]string `json:"header" bson:"header"`
		Body   primitive.M       `json:"body" bson:"body"`
		Code   int               `json:"code" bson:"code"`
	}
	USetting struct {
		Method            string             `json:"method" bson:"method"`
		URI               string             `json:"uri" bson:"uri"`
		ResponseFormat    int                `json:"response_format" bson:"response_format"`
		USettingResponses []USettingResponse `json:"responses" bson:"responses"`
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
		s.ResponseFormat = d.ResponseFormat
		s.USettingResponses = d.USettingResponses
	}
}
func (s *Setting) SetModifyTimestamp() {
	s.UpdatedAt = time.Now()
}

func NewUSetting() *USetting {
	ue := new(USetting)
	return ue
}
