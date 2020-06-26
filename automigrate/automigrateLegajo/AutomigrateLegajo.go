package automigrateLegajo

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/framework/configuracion"
)

type MicroservicioLegajo struct{
}

func (*MicroservicioLegajo) GetNombre() string {
	return "legajo"
}

func (*MicroservicioLegajo) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionlegajo
}

func (*MicroservicioLegajo) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionLegajoConfiguracion(), ObtenerVersionLegajoDB(db))
}

func (*MicroservicioLegajo) BeforeAutomigrarPublic(db *gorm.DB) error {
	err := db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Obrasocial{}, &structLegajo.Estadocivil{}).Error
	return err
}
func (*MicroservicioLegajo) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*MicroservicioLegajo) BeforeAutomigrarPrivate(db *gorm.DB) error {
	err := db.AutoMigrate(&structLegajo.Conyuge{}, &structLegajo.Hijo{}, &structLegajo.Legajo{}).Error
	if err == nil {
		db.Model(&structLegajo.Hijo{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
		db.Model(&structLegajo.Conyuge{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
	}
	return err
}

func (*MicroservicioLegajo) AfterAutomigrarPrivate(db *gorm.DB) error {
	return nil
}

func (*MicroservicioLegajo) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionLegajoConfiguracion(), Legajo, db)
}
