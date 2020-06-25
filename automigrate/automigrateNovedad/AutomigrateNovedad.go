package automigrateNovedad

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Novedad/structNovedad"
	"io/ioutil"
	"strconv"
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

	err = actualizarVersionesScript(ObtenerVersionNovedadDB(db), ObtenerVersionNovedadConfiguracion(), "novedad", "private", db)

	return err
}

func RunVersion(microservicio string, tipo string, version string, db *gorm.DB) error{
	path := "resources/" + microservicio + "/" + tipo + "-" + version + ".sql"

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		fmt.Printf("No existe el archivo o directorio, o no se puede acceder: %s\n", path)
		return nil
	}
	sql := string(c)
	err := db.Exec(sql).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Error al ejecutar el script %s/%s-%s: %s\n", microservicio, tipo, version, err.Error()))
	}
	return nil
}

func actualizarVersionesScript(versionDB int, versionConfig int, microservicio string, tipo string, db *gorm.DB) error{
	for versionDB++ ; versionDB <= versionConfig; versionDB++ {
		err := RunVersion(microservicio, tipo, strconv.Itoa(versionDB), db)
		if err != nil {
			return err
		}
	}
	return nil
}