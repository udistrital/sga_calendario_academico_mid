package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/sga_calendario_mid/services"
	"github.com/udistrital/utils_oas/requestresponse"

	"github.com/udistrital/utils_oas/errorhandler"
)

type ClonarCalendarioController struct {
	beego.Controller
}

func (c *ClonarCalendarioController) URLMapping() {
	c.Mapping("PostCalendario", c.PostCalendario)
	c.Mapping("PostCalendarioPadre", c.PostCalendarioPadre)
	c.Mapping("PostCalendarioExtension", c.PostCalendarioExtension)
}

// PostCalendario ...
// @Title PostCalendario
// @Description Clona calendario, crea tipo_evento si lo tiene, crea calendario_evento si tiene, crea calendario_evento_tipo_publico si tiene, crea tipo_publico si lo tiene
// @Param	body		body 	{}	true		"body id calendario content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *ClonarCalendarioController) PostCalendario() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.PostCalendario(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()

}

// PostCalendarioPadre ...
// @Title PostCalendarioPadre
// @Description Clona calendario padre, crea tipo_evento si lo tiene, crea calendario_evento si tiene, crea calendario_evento_tipo_publico si tiene, crea tipo_publico si lo tiene
// @Param	body		body 	{}	true		"body id calendario content"
// @Success 200 {}
// @Failure 400 the request contains incorrect syntax
// @router /padre [post]
func (c *ClonarCalendarioController) PostCalendarioPadre() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.PostCalendarioPadre(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// PostCalendarioExtension ...
// @Title PostCalendarioExtension
// @Description Clona calendario para extension, crea tipo_evento si lo tiene, crea calendario_evento si tiene, crea calendario_evento_tipo_publico si tiene, crea tipo_publico si lo tiene
// @Param	body		body 	{}	true		"body id calendario content"
// @Success 200 {}
// @Failure 400 the request contains incorrect syntax
// @router /extension [post]
func (c *ClonarCalendarioController) PostCalendarioExtension() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.PostCalendarioExtension(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resultado
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()

}
