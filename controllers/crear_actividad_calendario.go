package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/sga_calendario_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type ActividadCalendarioController struct {
	beego.Controller
}

func (c *ActividadCalendarioController) URLMapping() {
	c.Mapping("PostActividadCalendario", c.PostActividadCalendario)
	c.Mapping("UpdateActividadResponsables", c.UpdateActividadResponsables)
}

// PostActividadCalendario ...
// @Title PostActividadCalendario
// @Description Agregar actividad calendario, tipo_publico y tabla de rompimiento calendario_evento_tipo_publico
// @Param	body		body 	{}	true		"body Agregar Actividad calendario content"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *ActividadCalendarioController) PostActividadCalendario() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.PostActividadCalendario(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// UpdateActividadResponsables ...
// @Title UpdateActividadResponsables
// @Description Actualiza tabla de rompimiento calendario_evento_tipo_publico segun los responsables de una Actividad
// @Param	body		body 	{}	true		"body Actualizar responsables de una Actividad content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /calendario/actividad/:id [put]
func (c *ActividadCalendarioController) UpdateActividadResponsables() {
	defer errorhandler.HandlePanic(&c.Controller)

	idStr := c.Ctx.Input.Param(":id")
	data := c.Ctx.Input.RequestBody

	resultado, err := services.UpdateActividadResponsables(idStr, data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}
