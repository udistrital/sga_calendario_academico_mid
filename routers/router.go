// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/sga_calendario_academico_mid/controllers"
	"github.com/udistrital/utils_oas/errorhandler"
)

func init() {

	beego.ErrorController(&errorhandler.ErrorHandlerController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/clonar-calendario",
			beego.NSInclude(
				&controllers.ClonarCalendarioController{},
			),
		),
		beego.NSNamespace("/calendario-academico",
			beego.NSInclude(
				&controllers.ConsultaCalendarioAcademicoController{},
			),
		),
		beego.NSNamespace("/calendario-proyecto",
			beego.NSInclude(
				&controllers.ConsultaCalendarioProyectoController{},
			),
		),
		beego.NSNamespace("/actividad-calendario",
			beego.NSInclude(
				&controllers.ActividadCalendarioController{},
			),
		),
		beego.NSNamespace("/evento",
			beego.NSInclude(
				&controllers.EventoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
