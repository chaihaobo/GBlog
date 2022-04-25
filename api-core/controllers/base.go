package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/thoas/go-funk"
	"strings"
)
import "api-core/models/response"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) successData(data interface{}) {
	c.Data["json"] = response.Success(data)
	c.ServeJSON()
}

func (c *BaseController) success() {
	c.successData(nil)
}

func (c *BaseController) error(err error) {
	c.Data["json"] = response.FromError(err)
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) validation(data interface{}) {
	v := validation.Validation{}
	valid, _ := v.Valid(data)
	if !valid {
		errorList := funk.Map(v.Errors, func(error *validation.Error) string {
			return error.Key + "-" + error.Message
		}).([]string)
		errorMsg := strings.Join(errorList, ",")
		paramValidException := response.ParamValidException
		paramValidException.Message = errorMsg
		c.Data["json"] = response.FromException(&paramValidException)
		c.ServeJSON()
		c.StopRun()

	}

}

func (c *BaseController) parseBody(target interface{}) {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, target)
	if err != nil {
		c.error(err)
	}
}
