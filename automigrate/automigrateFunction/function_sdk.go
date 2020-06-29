package automigrateFunction

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/conexionBD/Function/structFunction"
	"unicode"
)

func CrearFunctionEnMicroservicioFormulas(db *gorm.DB, funcion string) error {

	var funcionData structFunction.Function

	err := json.Unmarshal([]byte(funcion), &funcionData)

	if err != nil {
		return err
	}

	err = SavePublicFunction(db, funcionData)

	return err
}

func SavePublicFunction(db *gorm.DB, functionData structFunction.Function) error {

	var err error
	runeName := []rune(functionData.Name)

	if !unicode.IsLetter(runeName[0]) {
		return errors.New("Las formulas deben empezar con una letra")
	}
	if unicode.IsLower(runeName[0]) {
		runeName[0] = unicode.ToUpper(runeName[0])
		functionData.Name = string(runeName)
	}

	var function structFunction.Function
	existeFuncion := true
	//Busco si existe
	if err := db.Set("gorm:auto_preload", true).First(&function, "name = ?", functionData.Name).Error; gorm.IsRecordNotFoundError(err) {
		existeFuncion = false;
	}

	if existeFuncion {
		//--Borrado Fisico
		if err := db.Unscoped().Where("name = ?", functionData.Name).Delete(structFunction.Function{}).Error; err != nil {
			return err
		}

		err := deleteValue(function.Value, db)
		if err != nil {
			return err
		}
	}

	if err := db.Create(&functionData).Error; err != nil {
		return err
	}

	return err
}

func deleteValue(value *structFunction.Value, db *gorm.DB) error {
	if value.Valueinvoke != nil {
		for i := 0; i < len(value.Valueinvoke.Args); i++ {
			err := deleteValue(&value.Valueinvoke.Args[i], db)
			if err != nil {
				return err
			}
		}

		if err := db.Unscoped().Where("id = ?", value.Valueinvokeid).Delete(structFunction.Invoke{}).Error; err != nil {
			return err
		}
	}

	if err := db.Unscoped().Where("id = ?", value.ID).Delete(structFunction.Value{}).Error; err != nil {
		return err
	}

	return nil
}