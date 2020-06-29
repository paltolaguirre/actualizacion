package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
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

func cleanConnections(db *gorm.DB)  {
	db.Model(&structAutenticacion.Security{}).Update("necesitaupdate", true)
}


func actualizarTablasPublicas() (error, bool) {
	dbPublic := conexionBD.ObtenerDB("public")
	defer conexionBD.CerrarDB(dbPublic)

	txPublic := dbPublic.Begin()
	defer func() {
		if r := recover(); r != nil {
			txPublic.Rollback()
		}
	}()

	err, actualizoMicro := automigrate.AutomigrateTablasPublicas(txPublic)
	if err != nil {
		txPublic.Rollback()
		return err, actualizoMicro
	}
	txPublic.Commit()

	return nil, actualizoMicro
}

func actualizarSecurity(actualizoMicro bool) error {
	dbSecurity := conexionBD.ObtenerDB("security")
	defer conexionBD.CerrarDB(dbSecurity)
	txSecurity := dbSecurity.Begin()
	defer func() {
		if r := recover(); r != nil {
			txSecurity.Rollback()
		}
	}()

	err, actualizoSecurity := automigrate.AutomigrateTablaSecurity(txSecurity)
	if err != nil {
		txSecurity.Rollback()
		return err
	}

	if actualizoMicro || actualizoSecurity {
		cleanConnections(txSecurity)
	}

	txSecurity.Commit()

	return nil
}