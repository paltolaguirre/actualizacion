package automigrateConcepto

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
	"github.com/xubiosueldos/conexionBD/Concepto/structConcepto"
	"github.com/xubiosueldos/framework/configuracion"
)

type AutomigrateConcepto struct {
	security structAutenticacion.Security
}

func (*AutomigrateConcepto) GetNombre() string {
	return "concepto"
}

func (am *AutomigrateConcepto) GetSecurity() structAutenticacion.Security {
	return am.security
}

func (am *AutomigrateConcepto) SetSecurity(security structAutenticacion.Security) {
	am.security = security
}

func (*AutomigrateConcepto) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionconcepto
}
func (am *AutomigrateConcepto) GetVersionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(am.GetNombre(), db)
}


func (am *AutomigrateConcepto) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(am.GetVersionConfiguracion(), am.GetVersionDB(db))
}

func (*AutomigrateConcepto) BeforeAutomigrarPublic() error {
	db := conexionBD.ObtenerDB("public")
	defer conexionBD.CerrarDB(db)
	err := db.AutoMigrate(&structConcepto.Tipocalculoautomatico{}, &structConcepto.Tipoconcepto{}, &structConcepto.Tipodecalculo{}, &structConcepto.Tipoimpuestoganancias{}, &structConcepto.Conceptoafip{}, &structConcepto.Concepto{}).Error
	return err
}

func (*AutomigrateConcepto) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*AutomigrateConcepto) BeforeAutomigrarPrivate(tenant string) error {
	db := conexionBD.ConnectBD(tenant)
	defer conexionBD.CerrarDB(db)

	err := db.AutoMigrate(&structConcepto.Concepto{}).Error
	if err == nil {
		db.Model(&structConcepto.Concepto{}).AddForeignKey("tipoconceptoid", "tipoconcepto(id)", "RESTRICT", "RESTRICT")
		db.Model(&structConcepto.Concepto{}).AddForeignKey("tipodecalculoid", "tipodecalculo(id)", "RESTRICT", "RESTRICT")
		db.Model(&structConcepto.Concepto{}).AddForeignKey("tipoimpuestogananciasid", "tipoimpuestoganancias(id)", "RESTRICT", "RESTRICT")
		db.Model(&structConcepto.Concepto{}).AddForeignKey("tipocalculoautomaticoid", "tipocalculoautomatico(id)", "RESTRICT", "RESTRICT")
		db.Model(&structConcepto.Concepto{}).AddForeignKey("formulanombre", "function(name)", "RESTRICT", "RESTRICT")
	}
	return err
}

func (am *AutomigrateConcepto) AfterAutomigrarPrivate(db *gorm.DB) error {
	err := am.ObtenerConceptosPublicos(db)
	return err
}

func (am *AutomigrateConcepto) ActualizarVersion(db *gorm.DB) {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(am.GetVersionConfiguracion(), am.GetNombre(), db)
}

func (am *AutomigrateConcepto) ObtenerConceptosPublicos(db *gorm.DB) error {

	versionConceptoDB := am.GetVersionDB(db)

	db.Exec("select ST_copy_concepto_public_privado()")

	if versionConceptoDB < 16 {
		db.Exec("ALTER SEQUENCE concepto_codigointerno_seq RESTART with 1000")
		db.Exec("UPDATE CONCEPTO SET codigointerno = nextval('concepto_codigointerno_seq') WHERE id > 0")
	}

	return nil
}
