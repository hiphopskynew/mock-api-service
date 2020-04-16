package route

import (
	"net/http"
	"sort"
	"strings"

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
		filtered := routes{}
		for _, r := range rs {
			if strings.HasPrefix(r.Name, "github.com/labstack/echo.(*Group)") || r.Path == "/*" {
				continue
			}
			filtered = append(filtered, r)
		}
		return c.JSON(http.StatusOK, filtered)
	}).Name = "inquiry available apis"

	ers := e.Any("/*", service.InvokeConfig)
	for _, r := range ers {
		r.Name = "all supported path see in /admin/api/manage/settings"
	}

	registerAdminManagement(e)
}

func registerAdminManagement(e *echo.Echo) {
	e.GET("/admin/api/manage/settings/:id", service.GetSetting).Name = "inquiry a setting"
	e.GET("/admin/api/manage/settings", service.GetSettingList).Name = "inquiry setting list"

	ge := e.Group("/admin/api/manage")
	ge.Use(middleware.JWT([]byte(config.Config.JWT.Secret)))

	ge.POST("/settings", service.AddSetting).Name = `create a new setting, example request body {'uri':'/api/v1/mock/inquiry','method':'GET','response':{'header':{'hmock':'Wuttikrai Limsakul','content-type':'application/json'},'body':{'a':'','b':0,'c':[],'d':{}},'code':200}}`
	ge.PUT("/settings/:id", service.UpdateSetting).Name = `update a setting, example request body {'uri':'/api/v1/mock/inquiry','method':'GET','response':{'header':{'hmock':'Wuttikrai Limsakul','content-type':'application/json'},'body':{'a':'','b':0,'c':[],'d':{}},'code':200}}`
	ge.DELETE("/settings/:id", service.DeleteSetting).Name = "delete a setting"
}
