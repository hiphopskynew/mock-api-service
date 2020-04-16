package route

import (
	"net/http"
	"sort"

	"github.com/hiphopskynew/mock-api-service/service"

	"github.com/labstack/echo"
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
	r := e.GET("/", func(c echo.Context) error {
		rs := routes(e.Routes())
		sort.Sort(rs)
		return c.JSON(http.StatusOK, rs)
	})
	r.Name = "inquiry available apis"

	r = e.POST("/admin/api/manage/settings", service.AddSetting)
	r.Name = `create a new setting, example request body {'uri':'/api/v1/mock/inquiry','method':'GET','response':{'header':{'hmock':'Wuttikrai Limsakul','content-type':'application/json'},'body':{'a':'','b':0,'c':[],'d':{}},'code':200}}`
	r = e.GET("/admin/api/manage/settings/:id", service.GetSetting)
	r.Name = "inquiry a setting"

	r = e.GET("/admin/api/manage/settings", service.GetSettingList)
	r.Name = "inquiry setting list"

	r = e.PUT("/admin/api/manage/settings/:id", service.UpdateSetting)
	r.Name = `update a setting, example request body {'uri':'/api/v1/mock/inquiry','method':'GET','response':{'header':{'hmock':'Wuttikrai Limsakul','content-type':'application/json'},'body':{'a':'','b':0,'c':[],'d':{}},'code':200}}`

	r = e.DELETE("/admin/api/manage/settings/:id", service.DeleteSetting)
	r.Name = "delete a setting"

	rs := e.Any("/*", service.InvokeConfig)
	for _, r := range rs {
		r.Name = "all supported path see in /admin/api/manage/settings"
	}
}
