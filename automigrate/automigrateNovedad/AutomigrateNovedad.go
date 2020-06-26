package automigrateNovedad

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Novedad/structNovedad"
	"github.com/xubiosueldos/framework/configuracion"
)

type MicroservicioNovedad struct{
}

func (*MicroservicioNovedad) GetNombre() string {
	return "novedad"
}

func (*MicroservicioNovedad) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionNovedadConfiguracion(), ObtenerVersionNovedadDB(db))
}

func (*MicroservicioNovedad) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionnovedad
}

func (*MicroservicioNovedad) BeforeAutomigrarPublic(db *gorm.DB) error {
	return nil
}
func (*MicroservicioNovedad) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*MicroservicioNovedad) BeforeAutomigrarPrivate(db *gorm.DB) error {
	err := db.AutoMigrate(&structNovedad.Novedad{}).Error
	db.Model(&structNovedad.Novedad{}).AddForeignKey("conceptoid", "concepto(id)", "RESTRICT", "RESTRICT")

	err = versiondbmicroservicio.ActualizarVersionesScript(ObtenerVersionNovedadDB(db), ObtenerVersionNovedadConfiguracion(), "novedad", "private", db)

	return err
}

func (*MicroservicioNovedad) AfterAutomigrarPrivate(db *gorm.DB) error {
	return nil
}

func (*MicroservicioNovedad) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionNovedadConfiguracion(), Novedad, db)
}

