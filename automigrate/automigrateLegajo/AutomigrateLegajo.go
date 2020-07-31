package automigrateLegajo

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
	"github.com/xubiosueldos/conexionBD/Legajo/structLegajo"
	"github.com/xubiosueldos/framework/configuracion"
)

type AutomigrateLegajo struct {
	security structAutenticacion.Security
}

func (*AutomigrateLegajo) GetNombre() string {
	return "legajo"
}

func (am *AutomigrateLegajo) GetSecurity() structAutenticacion.Security {
	return am.security
}

func (am *AutomigrateLegajo) SetSecurity(security structAutenticacion.Security) {
	am.security = security
}
func (*AutomigrateLegajo) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionlegajo
}

func (am *AutomigrateLegajo) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(am.GetVersionConfiguracion(), am.GetVersionDB(db))
}

func (*AutomigrateLegajo) BeforeAutomigrarPublic() error {
	db := conexionBD.ObtenerDB("public")
	defer conexionBD.CerrarDB(db)
	err := db.AutoMigrate(&structLegajo.Pais{}, &structLegajo.Provincia{}, &structLegajo.Localidad{}, &structLegajo.Modalidadcontratacion{}, &structLegajo.Situacion{}, &structLegajo.Condicion{}, &structLegajo.Condicionsiniestrado{}, &structLegajo.Obrasocial{}, &structLegajo.Estadocivil{}).Error
	return err
}
func (*AutomigrateLegajo) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*AutomigrateLegajo) BeforeAutomigrarPrivate(tenant string) error {
	db := conexionBD.ConnectBD(tenant)
	defer conexionBD.CerrarDB(db)

	err := db.AutoMigrate(&structLegajo.Conyuge{}, &structLegajo.Hijo{}, &structLegajo.Legajo{}).Error
	if err == nil {
		db.Model(&structLegajo.Hijo{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
		db.Model(&structLegajo.Conyuge{}).AddForeignKey("legajoid", "legajo(id)", "CASCADE", "CASCADE")
	}
	return err
}

func (*AutomigrateLegajo) AfterAutomigrarPrivate(db *gorm.DB) error {
	return nil
}

func (am *AutomigrateLegajo) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(am.GetVersionConfiguracion(), am.GetNombre(), db)
}

func (am *AutomigrateLegajo) GetVersionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(am.GetNombre(), db)
}