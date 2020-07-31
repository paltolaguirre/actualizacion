package automigrateLiquidacion

import (
	"net/http"

	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD"
	"github.com/xubiosueldos/conexionBD/Autenticacion/structAutenticacion"
	"github.com/xubiosueldos/conexionBD/Liquidacion/structLiquidacion"
	"github.com/xubiosueldos/framework/configuracion"
	"github.com/xubiosueldos/monoliticComunication"
)

type AutomigrateLiquidacion struct {
	security structAutenticacion.Security
}

func (*AutomigrateLiquidacion) GetNombre() string {
	return "liquidacion"
}

func (am *AutomigrateLiquidacion) GetSecurity() structAutenticacion.Security {
	return am.security
}

func (am *AutomigrateLiquidacion) SetSecurity(security structAutenticacion.Security) {
	am.security = security
}

func (*AutomigrateLiquidacion) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionliquidacion
}

func (am *AutomigrateLiquidacion) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(am.GetVersionConfiguracion(), am.GetVersionDB(db))
}

func (*AutomigrateLiquidacion) BeforeAutomigrarPublic() error {
	db := conexionBD.ObtenerDB("public")
	defer conexionBD.CerrarDB(db)
	err := db.AutoMigrate(&structLiquidacion.Liquidacioncondicionpago{}, &structLiquidacion.Liquidaciontipo{}, &structLiquidacion.Zona{}).Error
	return err
}

func (*AutomigrateLiquidacion) AfterAutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (am *AutomigrateLiquidacion) BeforeAutomigrarPrivate(tenant string) error {
	db := conexionBD.ConnectBD(tenant)
	defer conexionBD.CerrarDB(db)

	err := db.AutoMigrate(&structLiquidacion.Acumulador{}, &structLiquidacion.Liquidacionitem{}, &structLiquidacion.Liquidacion{}).Error
	if err == nil {

		versionLiquidacionDB := am.GetVersionDB(db)

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

func (am *AutomigrateLiquidacion) AfterAutomigrarPrivate(db *gorm.DB) error {
	var err error
	versionLiquidacionDB := am.GetVersionDB(db)
	if versionLiquidacionDB < 12 {
		var w http.ResponseWriter
		var r *http.Request
		security := am.GetSecurity()
		strempresa := monoliticComunication.Obtenerdatosempresa(w, r, &security, false)

		zonaid := strempresa.Zonaid

		db.Exec("UPDATE LIQUIDACION SET zonaid = " + strconv.Itoa(zonaid))
	}
	return err
}

func (am *AutomigrateLiquidacion) GetVersionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(am.GetNombre(), db)
}

func (am *AutomigrateLiquidacion) ActualizarVersion(db *gorm.DB) {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(am.GetVersionConfiguracion(), am.GetNombre(), db)
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
