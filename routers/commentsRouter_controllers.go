package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ActividadCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ActividadCalendarioController"],
        beego.ControllerComments{
            Method: "PostActividadCalendario",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ActividadCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ActividadCalendarioController"],
        beego.ControllerComments{
            Method: "UpdateActividadResponsables",
            Router: "/update/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ClonarCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ClonarCalendarioController"],
        beego.ControllerComments{
            Method: "PostCalendario",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ClonarCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ClonarCalendarioController"],
        beego.ControllerComments{
            Method: "PostCalendarioExtension",
            Router: "/calendario_extension",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ClonarCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ClonarCalendarioController"],
        beego.ControllerComments{
            Method: "PostCalendarioPadre",
            Router: "/calendario_padre",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"],
        beego.ControllerComments{
            Method: "GetOnePorId",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"],
        beego.ControllerComments{
            Method: "PostCalendarioHijo",
            Router: "/calendario_padre",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"],
        beego.ControllerComments{
            Method: "PutInhabilitarCalendario",
            Router: "/inhabilitar_calendario/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioAcademicoController"],
        beego.ControllerComments{
            Method: "GetCalendarInfo",
            Router: "/v2/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioProyectoController"],
        beego.ControllerComments{
            Method: "GetCalendarByProjectId",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:ConsultaCalendarioProyectoController"],
        beego.ControllerComments{
            Method: "GetCalendarProject",
            Router: "/nivel/:idNiv/periodo/:idPer",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:EventoController"],
        beego.ControllerComments{
            Method: "PostEvento",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:EventoController"],
        beego.ControllerComments{
            Method: "PutEvento",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:EventoController"],
        beego.ControllerComments{
            Method: "DeleteEvento",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_calendario_academico/controllers:EventoController"],
        beego.ControllerComments{
            Method: "GetEvento",
            Router: "/:persona",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
