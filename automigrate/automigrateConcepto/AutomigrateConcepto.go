package automigrateConcepto

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Concepto/structConcepto"
	"github.com/xubiosueldos/framework/configuracion"
)

type MicroservicioConcepto struct {
}

func (*MicroservicioConcepto) GetNombre() string {
	return "concepto"
}

func (*MicroservicioConcepto) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionconcepto
}


func (*MicroservicioConcepto) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionConceptoConfiguracion(), ObtenerVersionConceptoDB(db))
}

func (*MicroservicioConcepto) BeforeAutomigrarPublic(db *gorm.DB) error {
	err := db.AutoMigrate(&structConcepto.Tipocalculoautomatico{}, &structConcepto.Tipoconcepto{}, &structConcepto.Tipodecalculo{}, &structConcepto.Tipoimpuestoganancias{}, &structConcepto.Conceptoafip{}, &structConcepto.Concepto{}).Error
	return err
}

func (*MicroservicioConcepto) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*MicroservicioConcepto) BeforeAutomigrarPrivate(db *gorm.DB) error {
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

func (*MicroservicioConcepto) AfterAutomigrarPrivate(db *gorm.DB) error {
	err := ObtenerConceptosPublicos(db)
	return err
}

func (*MicroservicioConcepto) ActualizarVersion(db *gorm.DB) {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionConceptoConfiguracion(), Concepto, db)
}

func ObtenerConceptosPublicos(db *gorm.DB) error {

	versionConceptoDB := ObtenerVersionConceptoDB(db)

	db.Exec("select ST_copy_concepto_public_privado()")

	if versionConceptoDB < 16 {
		db.Exec("ALTER SEQUENCE concepto_codigointerno_seq RESTART with 1000")
		db.Exec("UPDATE CONCEPTO SET codigointerno = nextval('concepto_codigointerno_seq') WHERE id > 0")
	}

	return nil
}
