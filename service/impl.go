package service

import (
	"fmt"
	"net/http"

	"github.com/hiphopskynew/mock-api-service/repo"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddSetting(c echo.Context) error {
	data := new(repo.USetting)
	if e := c.Bind(data); e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}
	setting := repo.NewSetting()
	setting.Bind(data)

	// check existing the setting
	es, e := repo.DBSetting.GetSettingByMethodAndURI(c.Request().Context(), setting.Method, setting.URI)
	switch e {
	case mongo.ErrNoDocuments:
	case nil:
		return c.String(http.StatusBadRequest, fmt.Sprintf("existing uri, see in /admin/api/manage/settings/%s", es.ID.Hex()))
	default:
		return c.String(http.StatusBadRequest, e.Error())
	}

	// store a new setting
	result, e := repo.DBSetting.AddSetting(c.Request().Context(), *setting)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, result)
}
func GetSetting(c echo.Context) error {
	id := c.Param("id")
	oid, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}
	result, e := repo.DBSetting.GetSetting(c.Request().Context(), oid)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, result)
}
func GetSettingList(c echo.Context) error {
	result, e := repo.DBSetting.GetSettingList(c.Request().Context())
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, result)
}
func UpdateSetting(c echo.Context) error {
	id := c.Param("id")
	oid, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}
	data := new(repo.USetting)
	if e := c.Bind(data); e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}
	setting, e := repo.DBSetting.UpdateSetting(c.Request().Context(), oid, *data)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, setting)
}
func DeleteSetting(c echo.Context) error {
	id := c.Param("id")
	oid, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}
	if e := repo.DBSetting.DeleteSetting(c.Request().Context(), oid); e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, "OK")
}
func InvokeConfig(c echo.Context) error {
	method := c.Request().Method
	uri := c.Request().RequestURI
	setting, e := repo.DBSetting.GetSettingByMethodAndURI(c.Request().Context(), method, uri)
	if e != nil {
		return c.String(http.StatusBadRequest, "routing not match")
	}
	var (
		cfgBody   = setting.USetting.USettingResponse.Body
		cfgHeader = setting.USetting.USettingResponse.Header
		cfgCode   = setting.USetting.USettingResponse.Code
	)

	for k, v := range cfgHeader {
		c.Response().Header().Set(k, v)
	}
	c.Response().WriteHeader(cfgCode)
	return c.JSON(cfgCode, cfgBody)
}
