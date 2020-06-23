package automigrateNovedad

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Novedad/structNovedad"
)

type MicroservicioNovedad struct{
}

func (*MicroservicioNovedad) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(ObtenerVersionNovedadConfiguracion(), ObtenerVersionNovedadDB(db))
}

func (*MicroservicioNovedad) AutomigrarPublic(db *gorm.DB) error {
	return nil
}

func (*MicroservicioNovedad) AutomigrarPrivate(db *gorm.DB) error {
	return AutomigrateNovedadTablasPrivadas(db)
}

func (*MicroservicioNovedad) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(ObtenerVersionNovedadConfiguracion(), Novedad, db)
}

func AutomigrateNovedadTablasPrivadas(db *gorm.DB) error {

	//para actualizar tablas...agrega columnas e indices, pero no elimina
	err := db.AutoMigrate(&structNovedad.Novedad{}).Error
	db.Model(&structNovedad.Novedad{}).AddForeignKey("conceptoid", "concepto(id)", "RESTRICT", "RESTRICT")
	versionNovedadDB := ObtenerVersionNovedadDB(db)
	if versionNovedadDB < 2 {
		// err = db.Exec("alter table novedad alter column cantidad type numeric(19,4);").Error
		db.Exec("alter table novedad alter column cantidad type numeric(19,4);")
	}

	if versionNovedadDB < 4 {
		db.Exec("alter table novedad alter column importe drop not null;")
	}

	if versionNovedadDB < 5 {
		db.Exec("update novedad set fecha = fecha - interval '21 hours';")
	}
	return err
}
