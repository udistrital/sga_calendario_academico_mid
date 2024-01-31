package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/request"
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

	//Almacena el json que se trae desde el cliente
	var actividadCalendario map[string]interface{}
	//Almacena el resultado del json en algunas operaciones
	var actividadCalendarioPost map[string]interface{}
	var IdActividad interface{}
	var actividadPersonaPost map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &actividadCalendario); err == nil {
		Actividad := actividadCalendario["Actividad"]
		//Solicitid post a eventos service enviando el json recibido
		errActividad := request.SendJson("http://"+beego.AppConfig.String("EventoService")+"calendario_evento", "POST", &actividadCalendarioPost, Actividad)
		if errActividad == nil && fmt.Sprintf("%v", actividadCalendarioPost["System"]) != "map[]" && actividadCalendarioPost["Id"] != nil {
			if actividadCalendarioPost["Status"] != 400 {
				IdActividad = actividadCalendarioPost["Id"]
			} else {
				logs.Error(errActividad)
				c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, errActividad.Error())
				c.Data["system"] = actividadCalendarioPost
				c.Abort("400")
			}
		} else {
			logs.Error(errActividad)
			c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, errActividad.Error())
			c.Data["system"] = actividadCalendarioPost
			c.Abort("400")
		}

		var totalPublico []interface{}
		//Guarda el JSON de la tabla tipo publico
		totalPublico = actividadCalendario["responsable"].([]interface{})

		for _, publicoTemp := range totalPublico {
			CalendarioEventoTipoPersona := map[string]interface{}{
				"Activo":             true,
				"TipoPublicoId":      map[string]interface{}{"Id": publicoTemp.(map[string]interface{})["responsableID"].(float64)},
				"CalendarioEventoId": map[string]interface{}{"Id": IdActividad.(float64)},
			}

			errActividadPersona := request.SendJson("http://"+beego.AppConfig.String("EventoService")+"calendario_evento_tipo_publico", "POST", &actividadPersonaPost, CalendarioEventoTipoPersona)

			if errActividadPersona == nil && fmt.Sprintf("%v", actividadPersonaPost["System"]) != "map[]" && actividadPersonaPost["Id"] != nil {
				if actividadPersonaPost["Status"] != 400 {
					c.Ctx.Output.SetStatus(200)
					c.Data["json"] = requestresponse.APIResponseDTO(true, 200, actividadCalendarioPost)
				} else {
					var resultado2 map[string]interface{}
					request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("EventoService")+"/calendario_evento/%.f", actividadCalendarioPost["Id"]), "DELETE", &resultado2, nil)
					logs.Error(errActividadPersona)
					c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil)
					c.Data["system"] = actividadPersonaPost
					c.Abort("400")
				}
			} else {
				logs.Error(errActividadPersona)
				c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil)
				c.Data["system"] = actividadPersonaPost
				c.Abort("400")
			}
		}
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 404, nil, err.Error())
	}
	c.ServeJSON()
}

// UpdateActividadResponsables ...
// @Title UpdateActividadResponsables
// @Description Actualiza tabla de rompimiento calendario_evento_tipo_publico segun los responsables de una Actividad
// @Param	body		body 	{}	true		"body Actualizar responsables de una Actividad content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /update/:id [put]
func (c *ActividadCalendarioController) UpdateActividadResponsables() {
	defer errorhandler.HandlePanic(&c.Controller)

	var recibido map[string]interface{}
	var guardados []map[string]interface{}
	var actualizados []map[string]interface{}
	var auxDelete string
	var auxUpdate map[string]interface{}
	var errBorrado error

	idStr := c.Ctx.Input.Param(":id")
	actividadId, _ := strconv.Atoi(idStr)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &recibido); err == nil {
		datos := recibido["resp"].([]interface{})
		errConsulta := request.GetJson("http://"+beego.AppConfig.String("EventoService")+"calendario_evento_tipo_publico?query=CalendarioEventoId__Id:"+idStr, &guardados)
		if errConsulta == nil {
			if len(guardados) > 0 {
				for _, registro := range guardados {
					idRegistro := fmt.Sprintf("%.f", registro["Id"].(float64))
					errBorrado = request.SendJson("http://"+beego.AppConfig.String("EventoService")+"calendario_evento_tipo_publico/"+idRegistro, "DELETE", &auxDelete, nil)
					fmt.Println(errBorrado)
				}
			}
			if errBorrado == nil {
				for _, tipoPublico := range datos {
					nuevoPublico := map[string]interface{}{
						"Activo":             true,
						"TipoPublicoId":      map[string]interface{}{"Id": tipoPublico.(map[string]interface{})["responsableID"]},
						"CalendarioEventoId": map[string]interface{}{"Id": actividadId},
					}
					errPost := request.SendJson("http://"+beego.AppConfig.String("EventoService")+"calendario_evento_tipo_publico", "POST", &auxUpdate, nuevoPublico)
					if errPost == nil {
						actualizados = append(actualizados, auxUpdate)
					} else {
						logs.Error(errPost)
						c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, errPost.Error())
						c.Data["system"] = errPost
						c.Ctx.Output.SetStatus(400)
						c.ServeJSON()
						return
					}
				}
				c.Ctx.Output.SetStatus(200)
				c.Data["json"] = requestresponse.APIResponseDTO(false, 400, actualizados)
			} else {
				logs.Error(errBorrado)
				c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, errBorrado.Error())
				c.Data["system"] = errBorrado
				c.Ctx.Output.SetStatus(400)
			}
		} else {
			logs.Error(errConsulta)
			c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, errConsulta.Error())
			c.Data["system"] = errConsulta
			c.Ctx.Output.SetStatus(400)
		}
	} else {
		logs.Error(err)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
		c.Data["system"] = err
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
