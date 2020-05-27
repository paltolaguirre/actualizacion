package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
	"github.com/xubiosueldos/conexionBD/apiclientconexionbd"
	"github.com/xubiosueldos/framework/configuracion"
	"log"
	"net/http"
)

func main() {
	var err error
	configuracion := configuracion.GetInstance()

	dbPublic := conexionBD.ObtenerDB("public")
	defer func() {
		if r := recover(); r != nil {
			conexionBD.CerrarDB(dbPublic)
		}
	}()

	err, actualizoMicro := apiclientconexionbd.AutomigrateTablasPublicas(dbPublic)
	if err != nil {
		fmt.Println("Error Public Automigrate: ", err)
		return
	}
	conexionBD.CerrarDB(dbPublic)

	dbSecurity := conexionBD.ObtenerDB("security")
	txSecurity := dbSecurity.Begin()
	defer func() {
		if r := recover(); r != nil {
			txSecurity.Rollback()
			conexionBD.CerrarDB(dbSecurity)
		}
	}()

	err, actualizoSecurity := apiclientconexionbd.AutomigrateTablaSecurity(txSecurity)
	if err != nil {
		txSecurity.Rollback()
		fmt.Println("Error Security Automigrate: ", err)
		return
	}

	if actualizoMicro || actualizoSecurity {
		cleanConnections(txSecurity)
	}

	txSecurity.Commit()
	conexionBD.CerrarDB(dbSecurity)

	router := newRouter()

	fmt.Printf("Iniciando servidor escuchando puerto %s", configuracion.Puertomicroservicioactualizacion)
	server := http.ListenAndServe(":"+configuracion.Puertomicroservicioactualizacion, router)

	log.Fatal(server)

}

func cleanConnections(db *gorm.DB)  {
	db.Model(&structAutenticacion.Security{}).Update("necesitaupdate", true)
}