package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"

	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type EventoController struct {
	beego.Controller
}

// URLMapping ...
func (c *EventoController) URLMapping() {
	c.Mapping("PostEvento", c.PostEvento)
	c.Mapping("PutEvento", c.PutEvento)
	c.Mapping("GetEvento", c.GetEvento)
	c.Mapping("DeleteEvento", c.DeleteEvento)
}

// PostEvento ...
// @Title PostEvento
// @Description Agregar Evento
// @Param   body        body    {}  true        "body Agregar Evento content"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *EventoController) PostEvento() {
	defer errorhandler.HandlePanic(&c.Controller)

	var Evento map[string]interface{}
	var response interface{} = nil
	var success bool = false
	var message string = ""
	var statusCode int = 400
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &Evento); err == nil {

		EventoPost := make(map[string]interface{})
		/* EventoPost["Evento"] = map[string]interface{}{
			// "CalendarioEvento": Evento["Evento"],
			// "EncargadosEvento": Evento["EncargadosEvento"],
			// "TiposPublico": Evento["TiposPublico"]
		}*/

		Evento["Evento"].(map[string]interface{})["Activo"] = true
		EventoPost["CalendarioEvento"] = Evento["Evento"]

		encargadosEvento := make([]map[string]interface{}, 0)
		for _, encargadoTemp := range Evento["EncargadosEvento"].([]interface{}) {
			encargadoEvento := encargadoTemp.(map[string]interface{})
			encargadosEvento = append(encargadosEvento, map[string]interface{}{
				"RolEncargadoEventoId": encargadoEvento["RolEncargadoEventoId"],
				"EncargadoId":          encargadoEvento["EncargadoId"],
				"CalendarioEventoId":   map[string]interface{}{"Id": 0},
				"Activo":               true,
			})
		}
		EventoPost["EncargadosEvento"] = encargadosEvento

		tiposPublico := make([]map[string]interface{}, 0)
		for _, tipoPublicoTemp := range Evento["TiposPublico"].([]interface{}) {
			tipoPublico := tipoPublicoTemp.(map[string]interface{})
			tiposPublico = append(tiposPublico, map[string]interface{}{
				"Nombre":             tipoPublico["Nombre"],
				"CalendarioEventoId": map[string]interface{}{"Id": 0},
				"Activo":             true,
			})
		}
		EventoPost["TiposPublico"] = tiposPublico

		var resultadoEvento map[string]interface{}
		errProduccion := request.SendJson("http://"+beego.AppConfig.String("EventoService")+"tr_evento", "POST", &resultadoEvento, EventoPost)
		if resultadoEvento["Type"] == "error" || errProduccion != nil || resultadoEvento["Status"] == "404" || resultadoEvento["Message"] != nil {
			response = resultadoEvento
			statusCode = 400
			success = false
		} else {
			response = Evento
			statusCode = 200
			success = true
		}

	} else {
		statusCode = 400
		message += err.Error()
		success = false
	}

	if success {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, response)
	} else {
		c.Ctx.Output.SetStatus(statusCode)
		if message != "" {
			c.Data["json"] = requestresponse.APIResponseDTO(false, statusCode, nil, message)
		} else {
			c.Data["json"] = requestresponse.APIResponseDTO(false, statusCode, nil, message)
		}
	}
	c.ServeJSON()
}

// PutEvento ...
// @Title PutEvento
// @Description Modificar Evento
// @Param	id		path 	string	true		"el id del evento a modificar"
// @Param   body        body    {}  true        "body Modificar Evento content"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [put]
func (c *EventoController) PutEvento() {
	defer errorhandler.HandlePanic(&c.Controller)

	idStr := c.Ctx.Input.Param(":id")
	var response interface{} = nil
	var success bool = false
	var message string = ""
	var statusCode int = 400
	var Evento map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &Evento); err == nil {

		EventoPut := make(map[string]interface{})

		EventoPut["CalendarioEvento"] = Evento["Evento"]
		// EventoPut["EncargadosEvento"] = Evento["EncargadosEvento"];
		// EventoPut["TiposPublico"] = Evento["TiposPublico"];
		EventoPut["EncargadosEventoBorrados"] = Evento["EncargadosEventoBorrados"]
		EventoPut["TiposPublicoBorrados"] = Evento["TiposPublicoBorrados"]
		// Nuevos encargados de evento
		encargadosEvento := make([]map[string]interface{}, 0)
		for _, encargadoTemp := range Evento["EncargadosEvento"].([]interface{}) {
			encargadoEvento := encargadoTemp.(map[string]interface{})
			// solo se agregan los nuevos encargados
			fmt.Println("Encargado", encargadoEvento["Id"], encargadoEvento["EncargadoId"])
			if encargadoEvento["Id"].(float64) == 0 {
				fmt.Println("Agrega Encargado", encargadoEvento["Id"], encargadoEvento["EncargadoId"])
				encargadosEvento = append(encargadosEvento, map[string]interface{}{
					"RolEncargadoEventoId": encargadoEvento["RolEncargadoEventoId"],
					"EncargadoId":          encargadoEvento["EncargadoId"],
					"CalendarioEventoId":   map[string]interface{}{"Id": Evento["Evento"].(map[string]interface{})["Id"]},
					"Activo":               true,
				})
			}
		}
		EventoPut["EncargadosEvento"] = encargadosEvento

		tiposPublico := make([]map[string]interface{}, 0)
		for _, tipoPublicoTemp := range Evento["TiposPublico"].([]interface{}) {
			tipoPublico := tipoPublicoTemp.(map[string]interface{})
			if tipoPublico["Id"] != nil {
				tiposPublico = append(tiposPublico, map[string]interface{}{
					"Nombre":             tipoPublico["Nombre"],
					"CalendarioEventoId": map[string]interface{}{"Id": Evento["Evento"].(map[string]interface{})["Id"]},
					"Id":                 tipoPublico["Id"],
					"Activo":             true,
				})
			} else {
				tiposPublico = append(tiposPublico, map[string]interface{}{
					"Nombre":             tipoPublico["Nombre"],
					"CalendarioEventoId": map[string]interface{}{"Id": Evento["Evento"].(map[string]interface{})["Id"]},
					"Id":                 0,
					"Activo":             true,
				})
			}
		}
		EventoPut["TiposPublico"] = tiposPublico

		var resultadoEvento map[string]interface{}
		errProduccion := request.SendJson("http://"+beego.AppConfig.String("EventoService")+"/tr_evento/"+idStr, "PUT", &resultadoEvento, EventoPut)
		if resultadoEvento["Type"] == "error" || errProduccion != nil || resultadoEvento["Status"] == "404" || resultadoEvento["Message"] != nil {
			response = resultadoEvento
			statusCode = 400
			success = false
		} else {
			response = Evento
			statusCode = 200
			success = true
		}

	} else {
		statusCode = 400
		message += err.Error()
		success = false
	}

	if success {
		c.Ctx.Output.SetStatus(statusCode)
		c.Data["json"] = requestresponse.APIResponseDTO(success, statusCode, response)
	} else {
		c.Ctx.Output.SetStatus(statusCode)
		if message != "" {
			c.Data["json"] = requestresponse.APIResponseDTO(success, statusCode, response, message)
		} else {
			c.Data["json"] = requestresponse.APIResponseDTO(success, statusCode, response)
		}
	}
	c.ServeJSON()
}

// GetEvento ...
// @Title GetEvento
// @Description consultar Evento por persona
// @Param   persona      path    string  true        "Persona"
// @Success 200 {}
// @Failure 403 :persona is empty
// @router /:persona [get]
func (c *EventoController) GetEvento() {
	defer errorhandler.HandlePanic(&c.Controller)

	var response interface{} = nil
	var success bool = false
	var message string = ""
	var statusCode int = 400
	var eventos []map[string]interface{}

	persona := c.Ctx.Input.Param(":persona")
	fmt.Println("Get Evento")
	personaId, _ := strconv.ParseFloat(persona, 64)
	errEventos := request.GetJson("http://"+beego.AppConfig.String("EventoService")+"/tr_evento/"+persona, &eventos)
	if errEventos != nil || eventos[0]["CalendarioEvento"] == nil {
		success = false
		statusCode = 400
		message += "Error: errEventos es nil"

	} else {
		fmt.Println("paso")
		for _, evento := range eventos {

			if evento["CalendarioEvento"] != nil {

				encargados := evento["EncargadosEvento"].([]interface{})
				for _, encargadoTemp := range encargados {
					// seleccionar el rol de la persona
					encargado := encargadoTemp.(map[string]interface{})
					if encargado["EncargadoId"] == personaId {
						evento["RolPersona"] = encargado
					}
					// //cargar nombre del autor
					var encargadoEvento map[string]interface{}
					errEncargado := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"/tercero/"+fmt.Sprintf("%.f", encargado["EncargadoId"].(float64)), &encargadoEvento)
					if encargadoEvento["Type"] == "error" || errEncargado != nil {
						success = false
						statusCode = 400
						message += "Error: errEncargado es nil"
					} else {
						encargado["Nombre"] = encargadoEvento["PrimerNombre"].(string) + " " + encargadoEvento["SegundoNombre"].(string) + " " + encargadoEvento["PrimerApellido"].(string) + " " + encargadoEvento["SegundoApellido"].(string)
					}
				}
				// cargar nombre de la dependencia
				calendarioEvento := evento["CalendarioEvento"].(map[string]interface{})
				tipoEvento := calendarioEvento["TipoEventoId"].(map[string]interface{})
				var dependencia []map[string]interface{}
				errDependencia := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"dependencia_tipo_dependencia/?query=DependenciaId__Id:"+fmt.Sprintf("%.f", tipoEvento["DependenciaId"].(float64)), &dependencia)
				if dependencia == nil || errDependencia != nil {
					success = false
					message += "Error: errDependencia es nil"
					statusCode = 400
				} else {
					calendarioEvento["TipoDependenciaId"] = dependencia[0]["TipoDependenciaId"]
					calendarioEvento["DependenciaId"] = dependencia[0]["DependenciaId"]

				}

				// cargar periodo
				var periodo map[string]interface{}
				errPeriodo := request.GetJson("http://"+beego.AppConfig.String("CoreService")+"periodo/"+fmt.Sprintf("%.f", calendarioEvento["PeriodoId"].(float64)), &periodo)
				if periodo == nil || errPeriodo != nil {
					success = false
					message += "Error: errPeriodo es nil"
					statusCode = 400
				} else {
					evento["Periodo"] = periodo
				}
				evento["FechaInicio"] = calendarioEvento["FechaInicio"]
				evento["FechaFin"] = calendarioEvento["FechaFin"]
				evento["Descripcion"] = calendarioEvento["Descripcion"]
				evento["TipoEvento"] = tipoEvento["Nombre"]
				evento["Dependencia"] = calendarioEvento["DependenciaId"].(map[string]interface{})["Nombre"]

			}
		}
		response = eventos
		success = true
		statusCode = 200
	}

	c.Ctx.Output.SetStatus(statusCode)
	if success {
		c.Data["json"] = requestresponse.APIResponseDTO(success, statusCode, response)
	} else {
		if message != "" {
			c.Data["json"] = requestresponse.APIResponseDTO(success, statusCode, response, message)
		} else {
			c.Data["json"] = requestresponse.APIResponseDTO(success, statusCode, response)
		}
	}
	c.ServeJSON()
}

// DeleteEvento ...
// @Title DeleteEvento
// @Description eliminar Evento por id
// @Param   id      path    string  true        "Id del Evento"
// @Success 200 {}
// @Failure 403 :id is empty
// @router /:id [delete]
func (c *EventoController) DeleteEvento() {
	defer errorhandler.HandlePanic(&c.Controller)

	var response interface{} = nil
	var success bool = false
	var message string = ""
	var statusCode int = 400
	var eventoDeleted map[string]interface{}
	id := c.Ctx.Input.Param(":id")
	errEvento := request.SendJson(fmt.Sprintf("http://"+beego.AppConfig.String("EventoService")+"/tr_evento/"+id), "DELETE", &eventoDeleted, nil)
	if errEvento != nil || eventoDeleted["Message"] != nil {
		success = false
		message += "Error: errEvento"
		statusCode = 400
	} else {
		cadena, ok := eventoDeleted["Type"].(string)
		if ok {
			if cadena == "error" {
				success = false
				message += eventoDeleted["Body"].(string)
				statusCode = 400
			} else {
				success = true
				statusCode = 200
				response = eventoDeleted
			}
		} else {
			success = false
			statusCode = 400
			message += "Error: error al convertir la cadena type a string"
		}

	}
	c.Ctx.Output.SetStatus(statusCode)
	if message != "" {
		c.Data["json"] = requestresponse.APIResponseDTO(success, statusCode, response, message)
	} else {
		c.Data["json"] = requestresponse.APIResponseDTO(success, statusCode, response)
	}
	c.ServeJSON()
}
