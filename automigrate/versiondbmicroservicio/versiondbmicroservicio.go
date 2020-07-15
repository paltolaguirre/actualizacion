package versiondbmicroservicio

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/structGormModel"
	"io/ioutil"
	"strconv"
)

type Versiondbmicroservicio struct {
	structGormModel.GormModel
	Nombremicroservicio  string
	Versionmicroservicio int
}

var cacheDeScripts = make(map[string]string)

func CrearTablaVersionDBMicroservicio(db *gorm.DB) {

	db.AutoMigrate(&Versiondbmicroservicio{})

}

func UltimaVersion(nombre string, db *gorm.DB) int {

	var versiondbmicroservicio Versiondbmicroservicio

	db.First(&versiondbmicroservicio, "nombremicroservicio = ?", nombre)

	return versiondbmicroservicio.Versionmicroservicio

}

func ActualizarVersionMicroservicioDB(versionMicroservicioConfiguracion int, nombreMicroservicio string, db *gorm.DB) {

	var versiondbmicroservicio Versiondbmicroservicio

	db.First(&versiondbmicroservicio, "nombremicroservicio = ?", nombreMicroservicio)

	versiondbmicroservicio.Nombremicroservicio = nombreMicroservicio
	versiondbmicroservicio.Versionmicroservicio = versionMicroservicioConfiguracion

	db.Save(&versiondbmicroservicio)

}

func ActualizarMicroservicio(versionMicroservicioConfiguracion int, versionMicroservicioDB int) bool {
	return versionMicroservicioConfiguracion > versionMicroservicioDB
}

func RunVersion(microservicio string, tipo string, version string, db *gorm.DB) error {
	path := "resources/" + microservicio + "/" + tipo + "-" + version + ".sql"

	script := ""
	esta := false
	usacache := false
	if tipo == "private" {
		usacache = true
		script, esta = cacheDeScripts[path]
	}

	if !esta {
		c, ioErr := ioutil.ReadFile(path)
		if ioErr != nil {
			if usacache {
				cacheDeScripts[path] = ""
			}

		} else {
			script := string(c)
			if usacache {
				cacheDeScripts[path] = script
			}
		}
	}

	if script == "" {
		return nil
	}

	err := db.Exec(script).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Error al ejecutar el script %s/%s-%s: %s\n", microservicio, tipo, version, err.Error()))
	}
	fmt.Printf("Se ejecuto exitosamente el script: %s\n", path)
	return nil
}

func ActualizarVersionesScript(versionDB int, versionConfig int, microservicio string, tipo string, db *gorm.DB) error {
	for versionDB++; versionDB <= versionConfig; versionDB++ {
		err := RunVersion(microservicio, tipo, strconv.Itoa(versionDB), db)
		if err != nil {
			return err
		}
	}
	return nil
}
