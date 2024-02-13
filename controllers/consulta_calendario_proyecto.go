package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/sga_calendario_academico_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type ConsultaCalendarioProyectoController struct {
	beego.Controller
}

// URLMapping
func (c *ConsultaCalendarioProyectoController) URLMapping() {
	c.Mapping("GetCalendarByProjectId", c.GetCalendarByProjectId)
	c.Mapping("GetCalendarProject", c.GetCalendarProject)
}

// GetCalendarByProjectId ...
// @Title GetCalendarByProjectId
// @Description get ConsultaCalendarioAcademico by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ConsultaCalendarioProyectoController) GetCalendarByProjectId() {
	defer errorhandler.HandlePanic(&c.Controller)

	idCalendario, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	resultado, err := services.GetCalendarByProjectId(idCalendario)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// GetCalendarProject ...
// @Title GetCalendarProject
// @Description get ConsultaCalendarioAcademico & id y Project By Id
// @Param	idNiv	path	int	true		"Id nivel"
// @Param	:idNiv	path	int	true	"Id periodo"
// @Success 200
// @Failure 403 :id is empty
// @router /nivel/:idNiv/periodo/:idPer [get]
func (c *ConsultaCalendarioProyectoController) GetCalendarProject() {
	defer errorhandler.HandlePanic(&c.Controller)

	idNiv := c.Ctx.Input.Param(":idNiv")
	idPer := c.Ctx.Input.Param(":idPer")

	resultado, err := services.GetCalendarProject(idNiv, idPer)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}
