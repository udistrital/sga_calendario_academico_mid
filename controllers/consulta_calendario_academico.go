package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/sga_calendario_academico_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

// ConsultaCalendarioAcademicoController operations for Consulta_calendario_academico
type ConsultaCalendarioAcademicoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ConsultaCalendarioAcademicoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOnePorId", c.GetOnePorId)
	c.Mapping("Put", c.PutInhabilitarCalendario)
	c.Mapping("PostCalendarioHijo", c.PostCalendarioHijo)
	c.Mapping("GetCalendarInfo", c.GetCalendarInfo)
}

// GetAll ...
// @Title GetAll
// @Description get todos los calendarios académicos junto a sus periodos correspondientes
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ConsultaCalendarioAcademico
// @Failure 404
// @router / [get]
func (c *ConsultaCalendarioAcademicoController) GetAll() {
	defer errorhandler.HandlePanic(&c.Controller)

	resultado, err := services.GetAll()

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// GetOnePorId ...
// @Title GetOnePorId
// @Description get obtener calendario académico por id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ConsultaCalendarioAcademicoController) GetOnePorId() {
	defer errorhandler.HandlePanic(&c.Controller)

	idCalendario := c.Ctx.Input.Param(":id")

	resultado, err := services.GetOnePorId(idCalendario)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// PutInhabilitarCalendario ...
// @Title PutInhabilitarCalendario
// @Description Inhabilitar Calendario
// @Param	id		path 	string	true		"el id del calendario a inhabilitar"
// @Param   body        body    {}  true        "body Inhabilitar calendario content"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /inhabilitar_calendario/:id [put]
func (c *ConsultaCalendarioAcademicoController) PutInhabilitarCalendario() {
	defer errorhandler.HandlePanic(&c.Controller)

	idCalendario := c.Ctx.Input.Param(":id")
	data := c.Ctx.Input.RequestBody

	resultado, err := services.PutInhabilitarCalendario(idCalendario, data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// PostCalendarioHijo ...
// @Title PostCalendarioHijo
// @Description  Proyecto obtener el Id de calendario padre, crear el nuevo calendario (hijo) e inactivar el calendario padre
// @Param   body        body    {}  true        "body crear calendario hijo content"
// @Success 200 {}
// @Failure 403 :body is empty
// @router /padre [post]
func (c *ConsultaCalendarioAcademicoController) PostCalendarioHijo() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.PostCalendarioHijo(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// GetCalendarInfo ...
// @Title GetCalendarInfo
// @Description get obtener información calendario académico por id
// @Param	id		path 	string	true		"Id de calendario"
// @Success 200 {}
// @Failure 404 not found resource
// @router /v2/:id [get]
func (c *ConsultaCalendarioAcademicoController) GetCalendarInfo() {
	defer errorhandler.HandlePanic(&c.Controller)

	idCalendario := c.Ctx.Input.Param(":id")

	resultado, err := services.GetCalendarInfo(idCalendario)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()

}
