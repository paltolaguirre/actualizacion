package automigrateLegajo

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/framework/configuracion"
)

const (
	Legajo = "legajo"
)

func ObtenerVersionLegajoDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(Legajo, db)
}

func ObtenerVersionLegajoConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionlegajo
}
