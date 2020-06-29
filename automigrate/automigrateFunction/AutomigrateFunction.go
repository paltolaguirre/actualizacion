package automigrateFunction

import (
	"github.com/jinzhu/gorm"
	"github.com/xubiosueldos/actualizacion/automigrate/versiondbmicroservicio"
	"github.com/xubiosueldos/conexionBD/Function/structFunction"
	"github.com/xubiosueldos/framework/configuracion"
)

type AutomigrateFunction struct{
}

func (*AutomigrateFunction) GetNombre() string {
	return "formula"
}

func (*AutomigrateFunction) GetVersionConfiguracion() int {
	configuracion := configuracion.GetInstance()

	return configuracion.Versionformula
}

func (am *AutomigrateFunction) GetVersionDB(db *gorm.DB) int {
	return versiondbmicroservicio.UltimaVersion(am.GetNombre(), db)
}

func (am *AutomigrateFunction) NecesitaActualizar(db *gorm.DB) bool {
	return versiondbmicroservicio.ActualizarMicroservicio(am.GetVersionConfiguracion(), am.GetVersionDB(db))
}

func (*AutomigrateFunction) BeforeAutomigrarPublic(db *gorm.DB) error {
	err := db.AutoMigrate(&structFunction.Value{}, &structFunction.Invoke{}, &structFunction.Param{}, &structFunction.Function{}).Error

	return err
}

func (am *AutomigrateFunction) AfterAutomigrarPublic(db *gorm.DB) error {
	versionFunctionDB := am.GetVersionDB(db)

	if versionFunctionDB < 7 {
		err := CrearFunctionEnMicroservicioFormulas(db, "{ \"name\": \"Vacaciones\", \"CreatedAt\": \"2020-05-04T13:37:55.763839Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.84028Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Si la antiguedad es menor a 5 entonces le corresponden 14 dias de vacaciones.Entre 5 y 10 le corresponden 21 dias de vacacionesEntre 10 y 15 le corresponden 28 dias de vacacionesMayor a 15 le corresponden 35 dias de vacaciones\", \"origin\": \"custom\", \"type\": \"generic\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.762509Z\", \"UpdatedAt\": \"2020-05-04T13:37:55.762509Z\", \"DeletedAt\": null, \"name\": \"\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.69295Z\", \"UpdatedAt\": \"2020-05-04T13:37:55.69295Z\", \"DeletedAt\": null, \"function\": { \"name\": \"If\", \"CreatedAt\": \"2020-04-24T21:13:54.24135Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -7, \"CreatedAt\": \"2020-04-24T21:13:54.242805Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"condicion\", \"type\": \"boolean\", \"functionname\": \"If\" }, { \"ID\": -6, \"CreatedAt\": \"2020-04-24T21:13:54.244226Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"valueTrue\", \"type\": \"number\", \"functionname\": \"If\" }, { \"ID\": -5, \"CreatedAt\": \"2020-04-24T21:13:54.245668Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"valueFalse\", \"type\": \"number\", \"functionname\": \"If\" } ], \"description\": \"Retorna el primer valor si la condicion es verdadera y el segundo valor si la condicion es falsa\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -35, \"CreatedAt\": \"2020-04-24T21:13:54.239828Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -35 }, \"functionname\": \"If\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.714484Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.761506Z\", \"DeletedAt\": null, \"name\": \"condicion\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.695073Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.755597Z\", \"DeletedAt\": null, \"function\": { \"name\": \"LessEqual\", \"CreatedAt\": \"2020-04-24T21:13:54.282005Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -30, \"CreatedAt\": \"2020-04-29T15:33:53.033408Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val1\", \"type\": \"number\", \"functionname\": \"LessEqual\" }, { \"ID\": -31, \"CreatedAt\": \"2020-04-29T15:33:53.037409Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val2\", \"type\": \"number\", \"functionname\": \"LessEqual\" } ], \"description\": \"Dado dos valores retorna si el primero es menor o igual al segundo\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"boolean\", \"value\": { \"ID\": -41, \"CreatedAt\": \"2020-04-24T21:13:54.277353Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -41 }, \"functionname\": \"LessEqual\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.697676Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.758454Z\", \"DeletedAt\": null, \"name\": \"val1\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.696397Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.757062Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Antiguedad\", \"CreatedAt\": \"2020-04-24T21:13:54.349568Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Variable (MejorRemNoRemunerativaSemestre / 2) * DiasSemTrabajados / 180\", \"origin\": \"primitive\", \"type\": \"generic\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -55, \"CreatedAt\": \"2020-04-24T21:13:54.347864Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -55 }, \"functionname\": \"Antiguedad\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 181 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.713107Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.759987Z\", \"DeletedAt\": null, \"name\": \"val2\", \"valuenumber\": 5, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 181 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 196 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.721461Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.769113Z\", \"DeletedAt\": null, \"name\": \"valueTrue\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.715959Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.763077Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Multi\", \"CreatedAt\": \"2020-04-24T21:13:54.324499Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -27, \"CreatedAt\": \"2020-04-24T21:13:54.325985Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val\", \"type\": \"number\", \"functionname\": \"Multi\" }, { \"ID\": -28, \"CreatedAt\": \"2020-04-24T21:13:54.327421Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"percent\", \"type\": \"number\", \"functionname\": \"Multi\" } ], \"description\": \"Devuelve el primer numero multiplicado por el segundo numero\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -48, \"CreatedAt\": \"2020-04-24T21:13:54.323073Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -48 }, \"functionname\": \"Multi\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.718678Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.765914Z\", \"DeletedAt\": null, \"name\": \"val\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.717343Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.76453Z\", \"DeletedAt\": null, \"function\": { \"name\": \"ValorDiasVacaciones\", \"CreatedAt\": \"2020-04-24T21:13:54.14565Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Valor de los días correspondientes a las Vacaciones\", \"origin\": \"primitive\", \"type\": \"helper\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -13, \"CreatedAt\": \"2020-04-24T21:13:54.142806Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -13 }, \"functionname\": \"ValorDiasVacaciones\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 183 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.720101Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.767502Z\", \"DeletedAt\": null, \"name\": \"percent\", \"valuenumber\": 14, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 183 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 196 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.761228Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.83745Z\", \"DeletedAt\": null, \"name\": \"valueFalse\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.722771Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.77066Z\", \"DeletedAt\": null, \"function\": { \"name\": \"If\", \"CreatedAt\": \"2020-04-24T21:13:54.24135Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -7, \"CreatedAt\": \"2020-04-24T21:13:54.242805Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"condicion\", \"type\": \"boolean\", \"functionname\": \"If\" }, { \"ID\": -6, \"CreatedAt\": \"2020-04-24T21:13:54.244226Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"valueTrue\", \"type\": \"number\", \"functionname\": \"If\" }, { \"ID\": -5, \"CreatedAt\": \"2020-04-24T21:13:54.245668Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"valueFalse\", \"type\": \"number\", \"functionname\": \"If\" } ], \"description\": \"Retorna el primer valor si la condicion es verdadera y el segundo valor si la condicion es falsa\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -35, \"CreatedAt\": \"2020-04-24T21:13:54.239828Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -35 }, \"functionname\": \"If\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.729313Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.778358Z\", \"DeletedAt\": null, \"name\": \"condicion\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.724023Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.77224Z\", \"DeletedAt\": null, \"function\": { \"name\": \"LessEqual\", \"CreatedAt\": \"2020-04-24T21:13:54.282005Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -30, \"CreatedAt\": \"2020-04-29T15:33:53.033408Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val1\", \"type\": \"number\", \"functionname\": \"LessEqual\" }, { \"ID\": -31, \"CreatedAt\": \"2020-04-29T15:33:53.037409Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val2\", \"type\": \"number\", \"functionname\": \"LessEqual\" } ], \"description\": \"Dado dos valores retorna si el primero es menor o igual al segundo\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"boolean\", \"value\": { \"ID\": -41, \"CreatedAt\": \"2020-04-24T21:13:54.277353Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -41 }, \"functionname\": \"LessEqual\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.726609Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.77528Z\", \"DeletedAt\": null, \"name\": \"val1\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.72533Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.773716Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Antiguedad\", \"CreatedAt\": \"2020-04-24T21:13:54.349568Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Variable (MejorRemNoRemunerativaSemestre / 2) * DiasSemTrabajados / 180\", \"origin\": \"primitive\", \"type\": \"generic\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -55, \"CreatedAt\": \"2020-04-24T21:13:54.347864Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -55 }, \"functionname\": \"Antiguedad\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 185 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.727903Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.776851Z\", \"DeletedAt\": null, \"name\": \"val2\", \"valuenumber\": 10, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 185 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 195 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.736406Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.786142Z\", \"DeletedAt\": null, \"name\": \"valueTrue\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.730731Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.779965Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Multi\", \"CreatedAt\": \"2020-04-24T21:13:54.324499Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -27, \"CreatedAt\": \"2020-04-24T21:13:54.325985Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val\", \"type\": \"number\", \"functionname\": \"Multi\" }, { \"ID\": -28, \"CreatedAt\": \"2020-04-24T21:13:54.327421Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"percent\", \"type\": \"number\", \"functionname\": \"Multi\" } ], \"description\": \"Devuelve el primer numero multiplicado por el segundo numero\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -48, \"CreatedAt\": \"2020-04-24T21:13:54.323073Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -48 }, \"functionname\": \"Multi\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.733533Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.783006Z\", \"DeletedAt\": null, \"name\": \"val\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.732137Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.78152Z\", \"DeletedAt\": null, \"function\": { \"name\": \"ValorDiasVacaciones\", \"CreatedAt\": \"2020-04-24T21:13:54.14565Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Valor de los días correspondientes a las Vacaciones\", \"origin\": \"primitive\", \"type\": \"helper\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -13, \"CreatedAt\": \"2020-04-24T21:13:54.142806Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -13 }, \"functionname\": \"ValorDiasVacaciones\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 187 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.734963Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.784618Z\", \"DeletedAt\": null, \"name\": \"percent\", \"valuenumber\": 21, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 187 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 195 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.759885Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.835895Z\", \"DeletedAt\": null, \"name\": \"valueFalse\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.737867Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.81102Z\", \"DeletedAt\": null, \"function\": { \"name\": \"If\", \"CreatedAt\": \"2020-04-24T21:13:54.24135Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -7, \"CreatedAt\": \"2020-04-24T21:13:54.242805Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"condicion\", \"type\": \"boolean\", \"functionname\": \"If\" }, { \"ID\": -6, \"CreatedAt\": \"2020-04-24T21:13:54.244226Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"valueTrue\", \"type\": \"number\", \"functionname\": \"If\" }, { \"ID\": -5, \"CreatedAt\": \"2020-04-24T21:13:54.245668Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"valueFalse\", \"type\": \"number\", \"functionname\": \"If\" } ], \"description\": \"Retorna el primer valor si la condicion es verdadera y el segundo valor si la condicion es falsa\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -35, \"CreatedAt\": \"2020-04-24T21:13:54.239828Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -35 }, \"functionname\": \"If\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.744627Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.819108Z\", \"DeletedAt\": null, \"name\": \"condicion\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.739265Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.812829Z\", \"DeletedAt\": null, \"function\": { \"name\": \"LessEqual\", \"CreatedAt\": \"2020-04-24T21:13:54.282005Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -30, \"CreatedAt\": \"2020-04-29T15:33:53.033408Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val1\", \"type\": \"number\", \"functionname\": \"LessEqual\" }, { \"ID\": -31, \"CreatedAt\": \"2020-04-29T15:33:53.037409Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val2\", \"type\": \"number\", \"functionname\": \"LessEqual\" } ], \"description\": \"Dado dos valores retorna si el primero es menor o igual al segundo\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"boolean\", \"value\": { \"ID\": -41, \"CreatedAt\": \"2020-04-24T21:13:54.277353Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -41 }, \"functionname\": \"LessEqual\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.742083Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.816061Z\", \"DeletedAt\": null, \"name\": \"val1\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.740784Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.814445Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Antiguedad\", \"CreatedAt\": \"2020-04-24T21:13:54.349568Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Variable (MejorRemNoRemunerativaSemestre / 2) * DiasSemTrabajados / 180\", \"origin\": \"primitive\", \"type\": \"generic\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -55, \"CreatedAt\": \"2020-04-24T21:13:54.347864Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -55 }, \"functionname\": \"Antiguedad\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 189 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.743374Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.817695Z\", \"DeletedAt\": null, \"name\": \"val2\", \"valuenumber\": 15, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 189 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 194 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.752199Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.826883Z\", \"DeletedAt\": null, \"name\": \"valueTrue\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.746028Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.820612Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Multi\", \"CreatedAt\": \"2020-04-24T21:13:54.324499Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -27, \"CreatedAt\": \"2020-04-24T21:13:54.325985Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val\", \"type\": \"number\", \"functionname\": \"Multi\" }, { \"ID\": -28, \"CreatedAt\": \"2020-04-24T21:13:54.327421Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"percent\", \"type\": \"number\", \"functionname\": \"Multi\" } ], \"description\": \"Devuelve el primer numero multiplicado por el segundo numero\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -48, \"CreatedAt\": \"2020-04-24T21:13:54.323073Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -48 }, \"functionname\": \"Multi\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.749173Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.823581Z\", \"DeletedAt\": null, \"name\": \"val\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.747878Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.822187Z\", \"DeletedAt\": null, \"function\": { \"name\": \"ValorDiasVacaciones\", \"CreatedAt\": \"2020-04-24T21:13:54.14565Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Valor de los días correspondientes a las Vacaciones\", \"origin\": \"primitive\", \"type\": \"helper\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -13, \"CreatedAt\": \"2020-04-24T21:13:54.142806Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -13 }, \"functionname\": \"ValorDiasVacaciones\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 191 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.750503Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.825352Z\", \"DeletedAt\": null, \"name\": \"percent\", \"valuenumber\": 28, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 191 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 194 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.758608Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.834433Z\", \"DeletedAt\": null, \"name\": \"valueFalse\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.753583Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.828323Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Multi\", \"CreatedAt\": \"2020-04-24T21:13:54.324499Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -27, \"CreatedAt\": \"2020-04-24T21:13:54.325985Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val\", \"type\": \"number\", \"functionname\": \"Multi\" }, { \"ID\": -28, \"CreatedAt\": \"2020-04-24T21:13:54.327421Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"percent\", \"type\": \"number\", \"functionname\": \"Multi\" } ], \"description\": \"Devuelve el primer numero multiplicado por el segundo numero\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -48, \"CreatedAt\": \"2020-04-24T21:13:54.323073Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -48 }, \"functionname\": \"Multi\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.756049Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.831427Z\", \"DeletedAt\": null, \"name\": \"val\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.754855Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.829803Z\", \"DeletedAt\": null, \"function\": { \"name\": \"ValorDiasVacaciones\", \"CreatedAt\": \"2020-04-24T21:13:54.14565Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Valor de los días correspondientes a las Vacaciones\", \"origin\": \"primitive\", \"type\": \"helper\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -13, \"CreatedAt\": \"2020-04-24T21:13:54.142806Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -13 }, \"functionname\": \"ValorDiasVacaciones\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 193 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-04T13:37:55.757327Z\", \"UpdatedAt\": \"2020-05-04T21:00:44.832947Z\", \"DeletedAt\": null, \"name\": \"percent\", \"valuenumber\": 35, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 193 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 194 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 195 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 196 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 0 }, \"valueid\": 0, \"valueinvokeid\": 0 }")
		if err != nil {
			return err
		}

		err = CrearFunctionEnMicroservicioFormulas(db, "{ \"name\": \"Sac\", \"CreatedAt\": \"2020-05-07T14:26:30.133562Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.133562Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"(MejorRemRemunerativaBaseSACSemestre/2)*(DiasSemTrabajados/DiasDelSemestre)\", \"origin\": \"custom\", \"type\": \"generic\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.132532Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.132532Z\", \"DeletedAt\": null, \"name\": \"\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.119491Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.119491Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Multi\", \"CreatedAt\": \"2020-04-24T21:13:54.324499Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -27, \"CreatedAt\": \"2020-04-24T21:13:54.325985Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val\", \"type\": \"number\", \"functionname\": \"Multi\" }, { \"ID\": -28, \"CreatedAt\": \"2020-04-24T21:13:54.327421Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"percent\", \"type\": \"number\", \"functionname\": \"Multi\" } ], \"description\": \"Devuelve el primer numero multiplicado por el segundo numero\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -48, \"CreatedAt\": \"2020-04-24T21:13:54.323073Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -48 }, \"functionname\": \"Multi\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.125457Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.125457Z\", \"DeletedAt\": null, \"name\": \"val\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.121238Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.121238Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Div\", \"CreatedAt\": \"2020-04-24T21:13:54.318823Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -25, \"CreatedAt\": \"2020-04-24T21:13:54.320286Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val1\", \"type\": \"number\", \"functionname\": \"Div\" }, { \"ID\": -26, \"CreatedAt\": \"2020-04-24T21:13:54.321675Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val2\", \"type\": \"number\", \"functionname\": \"Div\" } ], \"description\": \"Devuelve el primer numero dividido el segundo numero\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -47, \"CreatedAt\": \"2020-04-24T21:13:54.317357Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -47 }, \"functionname\": \"Div\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.123242Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.123242Z\", \"DeletedAt\": null, \"name\": \"val1\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.122267Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.122267Z\", \"DeletedAt\": null, \"function\": { \"name\": \"MejorRemRemunerativaBaseSACSemestre\", \"CreatedAt\": \"2020-04-24T21:13:54.23007Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Mejor Remuneración del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre), comparando grillas de Remunerativo - Descuentos solo de los conceptos que tienen configurado BASE_SAC = SI\", \"origin\": \"primitive\", \"type\": \"helper\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -33, \"CreatedAt\": \"2020-04-24T21:13:54.228538Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -33 }, \"functionname\": \"MejorRemRemunerativaBaseSACSemestre\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 259 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.124525Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.124525Z\", \"DeletedAt\": null, \"name\": \"val2\", \"valuenumber\": 2, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 259 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 258 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.131428Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.131428Z\", \"DeletedAt\": null, \"name\": \"percent\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.12654Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.12654Z\", \"DeletedAt\": null, \"function\": { \"name\": \"Div\", \"CreatedAt\": \"2020-04-24T21:13:54.318823Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [ { \"ID\": -25, \"CreatedAt\": \"2020-04-24T21:13:54.320286Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val1\", \"type\": \"number\", \"functionname\": \"Div\" }, { \"ID\": -26, \"CreatedAt\": \"2020-04-24T21:13:54.321675Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"val2\", \"type\": \"number\", \"functionname\": \"Div\" } ], \"description\": \"Devuelve el primer numero dividido el segundo numero\", \"origin\": \"primitive\", \"type\": \"operator\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -47, \"CreatedAt\": \"2020-04-24T21:13:54.317357Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -47 }, \"functionname\": \"Div\", \"args\": [ { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.128409Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.128409Z\", \"DeletedAt\": null, \"name\": \"val1\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.127481Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.127481Z\", \"DeletedAt\": null, \"function\": { \"name\": \"DiasSemTrabajados\", \"CreatedAt\": \"2020-04-24T21:13:54.138341Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Dias Trabajados en el Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)\", \"origin\": \"primitive\", \"type\": \"helper\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -11, \"CreatedAt\": \"2020-04-24T21:13:54.136906Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -11 }, \"functionname\": \"DiasSemTrabajados\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 261 }, { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.130401Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.130401Z\", \"DeletedAt\": null, \"name\": \"val2\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": { \"ID\": 0, \"CreatedAt\": \"2020-05-07T14:26:30.129402Z\", \"UpdatedAt\": \"2020-05-07T14:26:30.129402Z\", \"DeletedAt\": null, \"function\": { \"name\": \"DiasDelSemestre\", \"CreatedAt\": \"2020-05-06T20:00:34.804702Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"params\": [], \"description\": \"Cantidad de dias corridos del semestre donde se encuentra utilizada la fórmula. Por ejemplo si la variable se utiliza en una liquidación del primer semestre del 2020 el resultado será 182. (utilizada para el cálculo automático de SAC)\", \"origin\": \"primitive\", \"type\": \"helper\", \"scope\": \"public\", \"result\": \"number\", \"value\": { \"ID\": -64, \"CreatedAt\": \"2020-05-06T20:00:34.801406Z\", \"UpdatedAt\": \"0001-01-01T00:00:00Z\", \"DeletedAt\": null, \"name\": \"return\", \"valuenumber\": 0, \"valuestring\": \"\", \"Valueboolean\": false, \"valueinvoke\": null, \"valueinvokeid\": null, \"arginvokeid\": 0 }, \"valueid\": -64 }, \"functionname\": \"DiasDelSemestre\", \"args\": [] }, \"valueinvokeid\": 0, \"arginvokeid\": 261 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 258 } ] }, \"valueinvokeid\": 0, \"arginvokeid\": 0 }, \"valueid\": 0, \"valueinvokeid\": 0 }")
		if err != nil {
			return err
		}

	}

	return nil
}

func (*AutomigrateFunction) BeforeAutomigrarPrivate(db *gorm.DB) error {
	err := db.AutoMigrate(&structFunction.Invoke{}, &structFunction.Value{}, &structFunction.Param{}, &structFunction.Function{}).Error
	if err == nil {
		db.Model(&structFunction.Param{}).AddForeignKey("functionname", "function(name)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Function{}).AddForeignKey("valueid", "value(id)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Invoke{}).AddForeignKey("functionname", "function(name)", "CASCADE", "CASCADE")
		db.Model(&structFunction.Value{}).AddForeignKey("valueinvokeid", "invoke(id)", "CASCADE", "CASCADE")
	}

	return err
}

func (*AutomigrateFunction) AfterAutomigrarPrivate(db *gorm.DB) error {
	return ObtenerFormulasPublicas(db)
}

func (am *AutomigrateFunction) ActualizarVersion(db *gorm.DB)  {
	versiondbmicroservicio.ActualizarVersionMicroservicioDB(am.GetVersionConfiguracion(), am.GetNombre(), db)
}

func ObtenerFormulasPublicas(db *gorm.DB) error {

	db.Exec("select ST_copy_formulas_public_privado()")

	return nil
}
