// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/sga_mid_calendario_academico/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/clonar-calendario",
			beego.NSInclude(
				&controllers.ClonarCalendarioController{},
			),
		),
		beego.NSNamespace("/calendario_academico",
			beego.NSInclude(
				&controllers.ConsultaCalendarioAcademicoController{},
			),
		),
		beego.NSNamespace("/calendario_proyecto",
			beego.NSInclude(
				&controllers.ConsultaCalendarioProyectoController{},
			),
		),
		beego.NSNamespace("/actividad_calendario",
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
