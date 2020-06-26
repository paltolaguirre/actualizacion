package automigrate

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/automigrateAutenticacion"
	"github.com/xubiosueldos/actualizacion/automigrate/automigrateConcepto"
	"github.com/xubiosueldos/actualizacion/automigrate/automigrateFunction"
	"github.com/xubiosueldos/actualizacion/automigrate/automigrateLegajo"
	"github.com/xubiosueldos/actualizacion/automigrate/automigrateLiquidacion"
	"github.com/xubiosueldos/actualizacion/automigrate/automigrateNovedad"
	"github.com/xubiosueldos/actualizacion/automigrate/automigrateSiradig"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
)



type Microservicio interface {
	NecesitaActualizar(*gorm.DB) bool
	BeforeAutomigrarPublic(*gorm.DB) error
	AfterAutomigrarPublic(*gorm.DB) error
	BeforeAutomigrarPrivate(*gorm.DB) error
	AfterAutomigrarPrivate(*gorm.DB) error
	ActualizarVersion(*gorm.DB)
	GetNombre() string
	GetVersionConfiguracion() int
}

var automigratePublicArray = []Microservicio{&automigrateLegajo.MicroservicioLegajo{}, &automigrateConcepto.MicroservicioConcepto{}, &automigrateLiquidacion.MicroservicioLiquidacion{}, &automigrateSiradig.MicroservicioSiradig{}, &automigrateFunction.MicroservicioFunction{}, &automigrateNovedad.MicroservicioNovedad{}}
var automigratePrivateArray = []Microservicio{&automigrateFunction.MicroservicioFunction{}, &automigrateLegajo.MicroservicioLegajo{}, &automigrateConcepto.MicroservicioConcepto{}, &automigrateNovedad.MicroservicioNovedad{}, &automigrateLiquidacion.MicroservicioLiquidacion{}, &automigrateSiradig.MicroservicioSiradig{}}

func AutomigrateTablaSecurity(db *gorm.DB) (error,bool) {

	actualizo := false
	var err error
	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.ObtenerVersionAutenticacionDB(db)) {
		if err = automigrateAutenticacion.AutomigrateAutenticacionTablaSecurity(db); err != nil {
			return err, actualizo
		} else {
			actualizo = true
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.Security, db)
		}
	}
	return err, actualizo

}

func AutomigrateTablasPublicas(db *gorm.DB) (error, bool) {

	actualizo := false
	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	for _, microservicio := range automigratePublicArray {
		if microservicio.NecesitaActualizar(db) {
			err := microservicio.BeforeAutomigrarPublic(db)

			if err != nil {
				return err, false
			}

			err = versiondbmicroservicio.ActualizarVersionesScript(versiondbmicroservicio.UltimaVersion(microservicio.GetNombre(), db), microservicio.GetVersionConfiguracion(), microservicio.GetNombre(), "public", db)

			if err != nil {
				return err, false
			}

			err = microservicio.AfterAutomigrarPublic(db)

			if err != nil {
				return err, false
			}

			actualizo = true
			microservicio.ActualizarVersion(db)
		}
	}

	return nil, actualizo
}


func AutomigrateTablasPrivadas(db *gorm.DB) error {

	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	for _, microservicio := range automigratePublicArray {
		if microservicio.NecesitaActualizar(db) {
			err := microservicio.BeforeAutomigrarPrivate(db)

			if err != nil {
				return err
			}

			err = versiondbmicroservicio.ActualizarVersionesScript(versiondbmicroservicio.UltimaVersion(microservicio.GetNombre(), db), microservicio.GetVersionConfiguracion(), microservicio.GetNombre(), "private", db)

			if err != nil {
				return err
			}

			err = microservicio.AfterAutomigrarPrivate(db)

			if err != nil {
				return err
			}

			microservicio.ActualizarVersion(db)
		}
	}

	return nil
}

