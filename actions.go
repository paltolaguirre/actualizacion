package main

import (
	"encoding/json"
	"fmt"
	"github.com/xubiosueldos/actualizacion/automigrate"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
	"github.com/xubiosueldos/framework"
	"net/http"
)

// Sirve para controlar si el server esta OK
func Healthy(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Healthy."))
}

func Actualizar(writer http.ResponseWriter, request *http.Request) {
	var err error
	decoder := json.NewDecoder(request.Body)

	var security structAutenticacion.Security

	if err := decoder.Decode(&security); err != nil {
		framework.RespondError(writer, http.StatusBadRequest, err.Error())
		return
	}

	tenant := security.Tenant
	db := conexionBD.ObtenerDB(tenant)
	defer conexionBD.CerrarDB(db)
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err= automigrate.AutomigrateTablasPrivadas(tx)

	dbSecurity := conexionBD.ObtenerDB("security")
	defer conexionBD.CerrarDB(dbSecurity)
	if err != nil {
		db.Rollback()
		fmt.Println("Error Automigrate Tablas Privadas: ", err)
		if err := dbSecurity.Unscoped().Where("token = ?", security.Token).Delete(structAutenticacion.Security{}).Error; err != nil {
			framework.RespondError(writer, http.StatusInternalServerError, err.Error())
		}
		framework.RespondError(writer, http.StatusInternalServerError, fmt.Sprintf("%s: %s", framework.ErrorAutomigrate, err.Error()))
		return
	} else {
		tx.Commit()
		security.Necesitaupdate = false
		dbSecurity.Model(&security).Where("tenant = ?", security.Tenant).Update("necesitaupdate", false)
	}
	framework.RespondJSON(writer, http.StatusOK, &security)
	return
}
