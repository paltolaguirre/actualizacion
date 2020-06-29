package automigrateNovedad

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Novedad/structNovedad"
	"github.com/xubiosueldos/framework/configuracion"
)

type AutomigrateNovedad struct{
}

func (*AutomigrateNovedad) GetNombre() string {
	return "novedad"
}

func (am *AutomigrateNovedad) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(am.GetVersionConfiguracion(), am.GetVersionDB(db))
}

func (*AutomigrateNovedad) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionnovedad
}

func (*AutomigrateNovedad) BeforeAutomigrarPublic(db *gorm.DB) error {
	return nil
}
func (*AutomigrateNovedad) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (am *AutomigrateNovedad) BeforeAutomigrarPrivate(db *gorm.DB) error {
	err := db.AutoMigrate(&structNovedad.Novedad{}).Error
	db.Model(&structNovedad.Novedad{}).AddForeignKey("conceptoid", "concepto(id)", "RESTRICT", "RESTRICT")

	return err
}

func (*AutomigrateNovedad) AfterAutomigrarPrivate(db *gorm.DB) error {
	return nil
}

func (am *AutomigrateNovedad) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(am.GetVersionConfiguracion(), am.GetNombre(), db)
}

func (am *AutomigrateNovedad) GetVersionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(am.GetNombre(), db)
}
