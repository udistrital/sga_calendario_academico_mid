package main

import (
	"github.com/go-co-op/gocron"
	_ "github.com/udistrital/sga_calendario_mid/routers"
	"github.com/udistrital/sga_calendario_mid/services"

	apistatus "github.com/udistrital/utils_oas/apiStatusLib"

	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/udistrital/auditoria"
	"github.com/udistrital/utils_oas/xray"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	xray.InitXRay()
	apistatus.Init()
	auditoria.InitMiddleware()

	// Inicializar el scheduler
	scheduler := gocron.NewScheduler(time.UTC)

	// Programar la verificación de actividades diarias cada hora
	scheduler.Every(1).Days().Do(services.VerificacionActividadesDiarias)

	// Programar las actividades de ejecución única
	services.VerificacionActividadesEjecucionUnica(scheduler)

	// Iniciar el scheduler de manera asíncrona
	scheduler.StartAsync()

	beego.Run()

	// Mantener el scheduler corriendo
	select {}
}
