package automigrateLiquidacion

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Liquidacion/structLiquidacion"
	"github.com/xubiosueldos/framework/configuracion"
)

type MicroservicioLiquidacion struct {
}

func (*MicroservicioLiquidacion) GetNombre() string {
	return "liquidacion"
}

func (*MicroservicioLiquidacion) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionliquidacion
}

func (*MicroservicioLiquidacion) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionLiquidacionConfiguracion(), ObtenerVersionLiquidacionDB(db))
}

func (*MicroservicioLiquidacion) BeforeAutomigrarPublic(db *gorm.DB) error {
	err := db.AutoMigrate(&structLiquidacion.Liquidacioncondicionpago{}, &structLiquidacion.Liquidaciontipo{}).Error
	return err
}

func (*MicroservicioLiquidacion) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*MicroservicioLiquidacion) BeforeAutomigrarPrivate(db *gorm.DB) error {
	err := db.AutoMigrate(&structLiquidacion.Acumulador{}, &structLiquidacion.Liquidacionitem{}, &structLiquidacion.Liquidacion{}).Error
	if err == nil {

		versionLiquidacionDB := ObtenerVersionLiquidacionDB(db)

		if versionLiquidacionDB < 7 {
			err = db.Exec("DELETE FROM liquidacionitem WHERE id IN (SELECT li.id FROM liquidacionitem li LEFT JOIN concepto c ON li.conceptoid = c.id WHERE c.id IS NULL)").Error
		}
		db.Model(&structLiquidacion.Liquidacionitem{}).AddForeignKey("liquidacionid", "liquidacion(id)", "CASCADE", "CASCADE")
		db.Model(&structLiquidacion.Liquidacionitem{}).AddForeignKey("conceptoid", "concepto(id)", "RESTRICT", "RESTRICT")
		db.Model(&structLiquidacion.Acumulador{}).AddForeignKey("liquidacionitemid", "liquidacionitem(id)", "CASCADE", "CASCADE")

		if versionLiquidacionDB < 4 {
			err = unificarDatosEnLaTablaLiquidacionItem(db)
		}
	}
	return err
}

func (*MicroservicioLiquidacion) AfterAutomigrarPrivate(db *gorm.DB) error {
	return nil
}

func (*MicroservicioLiquidacion) ActualizarVersion(db *gorm.DB) {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionLiquidacionConfiguracion(), Liquidacion, db)
}

func unificarDatosEnLaTablaLiquidacionItem(db *gorm.DB) error {
	//abro una transacción para que si hay un error no persista en la DB
	var err error

	if err = insertTablaLiquidacionTipo(db); err != nil {
		return err
	}
	return err
}

func insertTablaLiquidacionTipo(tx *gorm.DB) error {
	var err error

	//Necesito comparar porque las empresas nuevas no tienen las cinco tablas (importeremunerativo,descuento,retencion...)
	if tx.HasTable("importeremunerativo") {

		if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM importeremunerativo)").Error; err != nil {
			return err
		} else {
			tx.Exec("DELETE FROM importeremunerativo")
		}

		if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM importenoremunerativo)").Error; err != nil {
			return err
		} else {
			tx.Exec("DELETE FROM importenoremunerativo")

		}

		if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM descuento)").Error; err != nil {
			return err
		} else {
			tx.Exec("DELETE FROM descuento")
		}

		if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM retencion)").Error; err != nil {
			return err
		} else {
			tx.Exec("DELETE FROM retencion")
		}

		if err = tx.Exec("INSERT INTO liquidacionitem(created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid) (SELECT created_at,updated_at,deleted_at,conceptoid,importeunitario,liquidacionid FROM aportepatronal)").Error; err != nil {
			return err
		} else {
			tx.Exec("DELETE FROM aportepatronal")
		}
	}
	return err
}
