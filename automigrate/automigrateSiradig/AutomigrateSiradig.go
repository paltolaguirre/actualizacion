package automigrateSiradig

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Siradig/structSiradig"
	"github.com/xubiosueldos/framework/configuracion"
)


type AutomigrateSiradig struct{
}

func (*AutomigrateSiradig) GetNombre() string {
	return "siradig"
}

func (*AutomigrateSiradig) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionsiradig
}

func (am *AutomigrateSiradig) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(am.GetVersionConfiguracion(), am.GetVersionDB(db))
}

func (*AutomigrateSiradig) BeforeAutomigrarPublic(db *gorm.DB) error {
	err := db.AutoMigrate(&structSiradig.Siradigtipoimpuesto{}, &structSiradig.Siradigtipooperacion{}, &structSiradig.Siradigtipogrilla{}).Error

	return err
}

func (*AutomigrateSiradig) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*AutomigrateSiradig) BeforeAutomigrarPrivate(db *gorm.DB) error {
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

func (*AutomigrateSiradig) AfterAutomigrarPrivate(db *gorm.DB) error {
	return nil
}

func (am *AutomigrateSiradig) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(am.GetVersionConfiguracion(), am.GetNombre(), db)
}

func (am *AutomigrateSiradig) GetVersionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(am.GetNombre(), db)
}
