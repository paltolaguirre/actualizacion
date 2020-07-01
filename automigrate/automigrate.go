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
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
)



type Automigrate interface {
	NecesitaActualizar(*gorm.DB) bool
	BeforeAutomigrarPublic(*gorm.DB) error
	AfterAutomigrarPublic(*gorm.DB) error
	BeforeAutomigrarPrivate(*gorm.DB) error
	AfterAutomigrarPrivate(*gorm.DB) error
	ActualizarVersion(*gorm.DB)
	GetNombre() string
	GetVersionConfiguracion() int
	GetVersionDB(db *gorm.DB) int
}

var automigratePublicArray = []Automigrate{&automigrateLegajo.AutomigrateLegajo{}, &automigrateFunction.AutomigrateFunction{}, &automigrateConcepto.AutomigrateConcepto{}, &automigrateNovedad.AutomigrateNovedad{}, &automigrateSiradig.AutomigrateSiradig{}, &automigrateLiquidacion.AutomigrateLiquidacion{}}
var automigratePrivateArray = []Automigrate{&automigrateLegajo.AutomigrateLegajo{}, &automigrateFunction.AutomigrateFunction{}, &automigrateConcepto.AutomigrateConcepto{}, &automigrateNovedad.AutomigrateNovedad{}, &automigrateSiradig.AutomigrateSiradig{}, &automigrateLiquidacion.AutomigrateLiquidacion{}}

func AutomigrateTablaSecurity(db *gorm.DB, actualizoMicro bool) error {

	actualizo := false
	var err error
	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	txSecurity := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			txSecurity.Rollback()
		}
	}()

	if versiondbmicroservicio.ActualizarMicroservicio(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.ObtenerVersionAutenticacionDB(db)) {
		if err = automigrateAutenticacion.AutomigrateAutenticacionTablaSecurity(db); err != nil {
			txSecurity.Rollback()
			return err
		} else {
			actualizo = true
			versiondbmicroservicio.ActualizarVersionMicroservicioDB(automigrateAutenticacion.ObtenerVersionAutenticacionConfiguracion(), automigrateAutenticacion.Security, db)
		}
	}

	if actualizoMicro || actualizo {
		cleanConnections(txSecurity)
	}

	txSecurity.Commit()

	return err


}
func cleanConnections(db *gorm.DB)  {
	db.Model(&structAutenticacion.Security{}).Update("necesitaupdate", true)
}


func AutomigrateTablasPublicas(db *gorm.DB) (error, bool) {

	actualizo := false
	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	for _, microservicio := range automigratePublicArray {
		err := AutomigratePublico(microservicio, db, &actualizo)

		if err != nil {
			return err, actualizo
		}
	}

	return nil, actualizo
}

func AutomigratePublico(microservicio Automigrate, db *gorm.DB, actualizo *bool) error {

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	conexionBD.LockTable(tx, "versiondbmicroservicio")

	if microservicio.NecesitaActualizar(tx) {
		err := microservicio.BeforeAutomigrarPublic(tx)

		if err != nil {
			tx.Rollback()
			return err
		}

		err = versiondbmicroservicio.ActualizarVersionesScript(microservicio.GetVersionDB(tx), microservicio.GetVersionConfiguracion(), microservicio.GetNombre(), "public", tx)

		if err != nil {
			tx.Rollback()
			return err
		}

		err = microservicio.AfterAutomigrarPublic(tx)

		if err != nil {
			tx.Rollback()
			return err
		}

		*actualizo = true
		microservicio.ActualizarVersion(tx)
	}

	tx.Commit()

	return nil
}

func AutomigrateTablasPrivadas(db *gorm.DB) error {

	versiondbmicroservicio.CrearTablaVersionDBMicroservicio(db)

	for _, microservicio := range automigratePrivateArray {

		err := AutomigratePrivado(microservicio, db)

		if err != nil {
			return err
		}
	}

	return nil
}

func AutomigratePrivado(microservicio Automigrate, db *gorm.DB) error {

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	conexionBD.LockTable(tx, "versiondbmicroservicio")

	if microservicio.NecesitaActualizar(tx) {
		err := microservicio.BeforeAutomigrarPrivate(tx)

		if err != nil {
			tx.Rollback()
			return err
		}

		err = versiondbmicroservicio.ActualizarVersionesScript(microservicio.GetVersionDB(tx), microservicio.GetVersionConfiguracion(), microservicio.GetNombre(), "private", tx)

		if err != nil {
			tx.Rollback()
			return err
		}

		err = microservicio.AfterAutomigrarPrivate(tx)

		if err != nil {
			tx.Rollback()
			return err
		}

		microservicio.ActualizarVersion(tx)

	}

	tx.Commit()

	return nil
}

