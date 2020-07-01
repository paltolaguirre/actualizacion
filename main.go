package main

import (
	"fmt"
	"github.com/xubiosueldos/actualizacion/automigrate"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/framework/configuracion"
	"log"
	"net/http"
)

func main() {

	configuracion := configuracion.GetInstance()

	err, actualizoMicro := actualizarTablasPublicas()

	if err != nil {
		fmt.Println("Error Public Automigrate: ", err)
		return
	}

	err = actualizarSecurity(actualizoMicro)

	if err != nil {
		fmt.Println("Error Security Automigrate: ", err)
		return
	}

	router := newRouter()

	fmt.Printf("Iniciando servidor escuchando puerto %s\n", configuracion.Puertomicroservicioactualizacion)
	server := http.ListenAndServe(":"+configuracion.Puertomicroservicioactualizacion, router)

	log.Fatal(server)

}

func actualizarTablasPublicas() (error, bool) {
	dbPublic := conexionBD.ObtenerDB("public")
	defer conexionBD.CerrarDB(dbPublic)

	err, actualizoMicro := automigrate.AutomigrateTablasPublicas(dbPublic)
	if err != nil {
		return err, actualizoMicro
	}

	return nil, actualizoMicro
}

func actualizarSecurity(actualizoMicro bool) error {
	dbSecurity := conexionBD.ObtenerDB("security")
	defer conexionBD.CerrarDB(dbSecurity)

	err := automigrate.AutomigrateTablaSecurity(dbSecurity, actualizoMicro)
	if err != nil {
		return err
	}

	return nil
}