package automigrateSiradig

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Siradig/structSiradig"
	"github.com/xubiosueldos/framework/configuracion"
)


type MicroservicioSiradig struct{
}

func (*MicroservicioSiradig) GetNombre() string {
	return "siradig"
}

func (*MicroservicioSiradig) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionsiradig
}

func (*MicroservicioSiradig) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionSiradigConfiguracion(), ObtenerVersionSiradigDB(db))
}

func (*MicroservicioSiradig) BeforeAutomigrarPublic(db *gorm.DB) error {
	err := db.AutoMigrate(&structSiradig.Siradigtipoimpuesto{}, &structSiradig.Siradigtipooperacion{}, &structSiradig.Siradigtipogrilla{}).Error
	if err == nil {
		versiondbmicroservicio.ActualizarVersionesScript(ObtenerVersionSiradigDB(db), ObtenerVersionSiradigConfiguracion(), "siradig", "public", db)
	}

	return err
}

func (*MicroservicioSiradig) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*MicroservicioSiradig) BeforeAutomigrarPrivate(db *gorm.DB) error {
	err := db.AutoMigrate(&structSiradig.Detallecargofamiliarsiradig{}, &structSiradig.Importegananciasotroempleosiradig{}, &structSiradig.Deducciondesgravacionsiradig{}, &structSiradig.Retencionpercepcionsiradig{}, &structSiradig.Beneficiosiradig{}, &structSiradig.Ajustesiradig{}, &structSiradig.Siradig{}).Error
	if err == nil {
		db.Model(&structSiradig.Detallecargofamiliarsiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Detallecargofamiliarsiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Importegananciasotroempleosiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Deducciondesgravacionsiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Deducciondesgravacionsiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Retencionpercepcionsiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Retencionpercepcionsiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Beneficiosiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Beneficiosiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

		db.Model(&structSiradig.Ajustesiradig{}).AddForeignKey("siradigid", "siradig(id)", "CASCADE", "CASCADE")
		db.Model(&structSiradig.Ajustesiradig{}).AddForeignKey("siradigtipogrillaid", "siradigtipogrilla(id)", "CASCADE", "CASCADE")

	}
	return err
}

func (*MicroservicioSiradig) AfterAutomigrarPrivate(db *gorm.DB) error {
	return nil
}

func (*MicroservicioSiradig) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionSiradigConfiguracion(), Siradig, db)
}
