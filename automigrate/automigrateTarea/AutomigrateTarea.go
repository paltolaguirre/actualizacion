package automigrateTarea

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
	"github.com/xubiosueldos/conexionBD/Tarea/structTarea"
	"github.com/xubiosueldos/framework/configuracion"
)

type AutomigrateTarea struct {
	security structAutenticacion.Security
}

func (*AutomigrateTarea) GetNombre() string {
	return "tarea"
}

func (am *AutomigrateTarea) GetSecurity() structAutenticacion.Security {
	return am.security
}

func (am *AutomigrateTarea) SetSecurity(security structAutenticacion.Security) {
	am.security = security
}

func (*AutomigrateTarea) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versiontarea
}

func (am *AutomigrateTarea) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(am.GetVersionConfiguracion(), am.GetVersionDB(db))
}

func (*AutomigrateTarea) BeforeAutomigrarPublic() error {
	db := conexionBD.ObtenerDB("public")
	defer conexionBD.CerrarDB(db)

	err := db.AutoMigrate(&structTarea.Tareapendiente{}).Error

	return err
}

func (*AutomigrateTarea) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*AutomigrateTarea) BeforeAutomigrarPrivate(tenant string) error {
	db := conexionBD.ConnectBD(tenant)
	defer conexionBD.CerrarDB(db)

	err := db.AutoMigrate(&structTarea.Tarea{}, &structTarea.Tareaitem{}).Error
	return err
}

func (*AutomigrateTarea) AfterAutomigrarPrivate(db *gorm.DB) error {
	return nil
}

func (am *AutomigrateTarea) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(am.GetVersionConfiguracion(), am.GetNombre(), db)
}

func (am *AutomigrateTarea) GetVersionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(am.GetNombre(), db)
}

