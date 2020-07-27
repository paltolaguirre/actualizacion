package automigrateTarea

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Tarea/structTarea"
	"github.com/xubiosueldos/framework/configuracion"
)

type AutomigrateTarea struct{
}

func (*AutomigrateTarea) GetNombre() string {
	return "tarea"
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

	err := db.AutoMigrate(&structTarea.Tareapendiente{}, &structTarea.Estadotareaitem{}, &structTarea.Estadotarea{}).Error

	return err
}

func (*AutomigrateTarea) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*AutomigrateTarea) BeforeAutomigrarPrivate(tenant string) error {
	db := conexionBD.ConnectBD(tenant)
	defer conexionBD.CerrarDB(db)

	err := db.AutoMigrate(&structTarea.Tarea{}, &structTarea.Tareaitem{}, &structTarea.Cancelacion{}).Error
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

