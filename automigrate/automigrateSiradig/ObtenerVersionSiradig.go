package automigrateSiradig

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Siradig = "siradig"
)

func ObtenerVersionSiradigDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Siradig, db)
}

func ObtenerVersionSiradigConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionsiradig
}
