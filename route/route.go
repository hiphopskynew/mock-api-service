package route

import (
	"net/http"
	"sort"

	"github.com/hiphopskynew/mock-api-service/config"
	"github.com/hiphopskynew/mock-api-service/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type routes []*echo.Route

func (r routes) Less(i, j int) bool {
	return len(r[i].Path)+len(r[i].Method) < len(r[j].Path)+len(r[j].Method)
}
func (r routes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r routes) Len() int {
	return len(r)
}

func RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		rs := routes(e.Routes())
		sort.Sort(rs)
		return c.JSON(http.StatusOK, rs)
	}).Name = "inquiry available apis"

	for _, r := range e.Any("/*", service.InvokeConfig) {
		r.Name = "all supported path see in /admin/api/manage/settings"
	}

	ge := e.Group("/admin/api/manage")
	ge.Use(middleware.JWT([]byte(config.Config.JWT.Secret)))

	ge.POST("/settings", service.AddSetting).Name = `create a new setting, example request body {'uri':'/api/v1/mock/inquiry','method':'GET','response':{'header':{'hmock':'Wuttikrai Limsakul','content-type':'application/json'},'body':{'a':'','b':0,'c':[],'d':{}},'code':200}}`
	ge.GET("/settings/:id", service.GetSetting).Name = "inquiry a setting"
	ge.GET("/settings", service.GetSettingList).Name = "inquiry setting list"
	ge.PUT("/settings/:id", service.UpdateSetting).Name = `update a setting, example request body {'uri':'/api/v1/mock/inquiry','method':'GET','response':{'header':{'hmock':'Wuttikrai Limsakul','content-type':'application/json'},'body':{'a':'','b':0,'c':[],'d':{}},'code':200}}`
	ge.DELETE("/settings/:id", service.DeleteSetting).Name = "delete a setting"
}
