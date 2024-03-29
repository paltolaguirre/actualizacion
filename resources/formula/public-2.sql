DO $$
BEGIN
    IF ((select count(*) from function where name = 'CantidadSueldos') = 0) THEN
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-6,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Sueldo', current_timestamp, 'Sueldo ingresado en el campo Remuneración del Legajo', 'primitive', 'helper', 'public', 'number', -6);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-7,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('HorasMensuales', current_timestamp, 'Horas ingresadas en el campo Horas Mensuales Normales del legajo', 'primitive', 'helper', 'public', 'number', -7);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-8,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasMesTrabajadosFechaLiquidacion', current_timestamp, 'Cantidad de dias desde el primero del mes hasta la fecha de la liquidación', 'primitive', 'helper', 'public', 'number', -8);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-9,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasMesTrabajadosFechaPeriodo', current_timestamp, 'Cantidad de dias desde el primero del mes hasta el último día del período liquidado', 'primitive', 'helper', 'public', 'number', -9);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-10,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemRemunerativaSemestre', current_timestamp, 'Mejor Remuneración del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre), comparando grillas de (Remunerativo - Desacuentos)', 'primitive', 'helper', 'public', 'number', -10);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-11,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasSemTrabajados', current_timestamp, 'Dias Trabajados en el Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)', 'primitive', 'helper', 'public', 'number', -11);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-12,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemNoRemunerativaSemestre', current_timestamp, 'Mejor Remuneración No Remunerativa del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)', 'primitive', 'helper', 'public', 'number', -12);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-13,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('ValorDiasVacaciones', current_timestamp, 'Valor de los días correspondientes a las Vacaciones', 'primitive', 'helper', 'public', 'number', -13);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-14, current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalHaberesNoRemunerativosMensual', current_timestamp, 'Total de conceptos no remunerativos de la liquidación', 'primitive', 'helper', 'public', 'number', -14);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-15,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalAportesPatronalesMensual', current_timestamp, 'Total de conceptos de aportes patronales de la liquidación', 'primitive', 'helper', 'public', 'number', -15);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-16,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalRetencionesMensual', current_timestamp, 'Total de conceptos de retenciones de la liquidación', 'primitive', 'helper', 'public', 'number', -16);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-17,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('CantidadMesesTrabajados', current_timestamp, 'La cantidad de meses trabajados desde  Legajo > Fecha de Ingreso  hasta la Fecha de la liquidación', 'primitive', 'helper', 'public', 'number', -17);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-18,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasLicenciaMes', current_timestamp, 'Cantidad de días de licencias del mes de liquidación', 'primitive', 'helper', 'public', 'number', -18);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-19,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasLicenciaSemestre', current_timestamp, 'Cantidad de días de licencias del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)', 'primitive', 'helper', 'public', 'number', -19);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-20,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalDescuentosMensual', current_timestamp, 'Total de conceptos de descuentos de la liquidación', 'primitive', 'helper', 'public', 'number', -20);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-21,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemNormalYHabitualSemestre', current_timestamp, 'Mejor remuneración normal y habitual del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre)', 'primitive', 'helper', 'public', 'number', -21);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-22,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasEfectivamenteTrabajadosSemestre', current_timestamp, 'Cantidad de días efectivamente trabajados en el Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre). Días hábiles - Días Faltas injustificadas - Días enfermedad - Días accidente - Días de Licencia', 'primitive', 'helper', 'public', 'number', -22);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-23,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('FechadeIngreso', current_timestamp, 'Fecha de ingreso del Legajo', 'primitive', 'helper', 'public', 'Time', -23);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-24,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('FechadeLiquidacion', current_timestamp, 'Fecha de la liquidación donde se esta utilizando el concepto', 'primitive', 'helper', 'public', 'Time', -24);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-25,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('FecIngHASTAFecLiq', current_timestamp, 'Fecha de Ingreso hasta Fecha de Liquidacion expresada en años', 'primitive', 'helper', 'public', 'number', -25);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-26,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('AntiguedadResto', current_timestamp, ' ( (Dias desde la Fecha Ingreso hasta la Fecha Liquidación) / 365 ) - antiguedad', 'primitive', 'helper', 'public', 'number', -26);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-27,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemTotalSemestre', current_timestamp, 'Mejor remuneración total del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre). Comparando las grillas de Remunerativo - Descuentos + No Remunerativo', 'primitive', 'helper', 'public', 'number', -27);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-28,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemTotalAnual', current_timestamp, 'Mejor remuneración total del año. Comparando las grillas de Remunerativo - Descuentos + No Remunerativo', 'primitive', 'helper', 'public', 'number', -28);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-29,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemTotalSinRemVarSemestre', current_timestamp, 'Mejor remuneración total, sin incluir Remuneraciones Variables del Semestre. Comparando las grillas de Remunerativo - Descuentos + No Remunerativo', 'primitive', 'helper', 'public', 'number', -29);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-30,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemTotalSinRemVarAnual', current_timestamp, 'Mejor remuneración total, sin remuneraciones variables del año liquidado. Comparando las grillas de Remunerativo - Descuentos + No Remunerativo', 'primitive', 'helper', 'public', 'number', -30);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-31,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('PromRemVariablesSemestre', current_timestamp, 'Promedio remuneraciones variables del Semestre', 'primitive', 'helper', 'public', 'number', -31);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-32,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('PromRemVariablesAnual', current_timestamp, 'Promedio remuneraciones variables del Año Liquidado', 'primitive', 'helper', 'public', 'number', -32);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-33,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('MejorRemRemunerativaBaseSACSemestre', current_timestamp, 'Mejor Remuneración del Semestre (Semestre 1: Enero - Junio | Semestre 2: Julio - Diciembre), comparando grillas de Remunerativo - Descuentos solo de los conceptos que tienen configurado BASE_SAC = SI', 'primitive', 'helper', 'public', 'number', -33);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-34,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('GetConceptValue', current_timestamp, 'Obtiene el valor de un concepto', 'primitive', 'internal', 'public', 'number', -34);
        INSERT INTO param(id, created_at, name, type, functionname) VALUES(-4, current_timestamp, 'conceptid', 'number', 'GetConceptValue');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-35,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('If', current_timestamp, 'Retorna el primer valor si la condicion es verdadera y el segundo valor si la condicion es falsa', 'primitive', 'operator', 'public', 'number', -35);
        INSERT INTO param(id, created_at, name, type, functionname) VALUES(-7, current_timestamp, 'condicion', 'boolean', 'If');
        INSERT INTO param(id, created_at, name, type, functionname) VALUES(-6, current_timestamp, 'valueTrue', 'number', 'If');
        INSERT INTO param(id, created_at, name, type, functionname) VALUES(-5, current_timestamp, 'valueFalse', 'number', 'If');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-36,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Equality', current_timestamp, 'Dado dos valores retorna si son iguales', 'primitive', 'operator', 'public', 'boolean', -36);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-8,current_timestamp,'val1','number','Equality');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-9,current_timestamp,'val2','number','Equality');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-37,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Inequality', current_timestamp, 'Dado dos valores retorna si son distintos', 'primitive', 'operator', 'public', 'boolean', -37);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-10,current_timestamp,'val1','number','Inequality');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-11,current_timestamp,'val2','number','Inequality');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-38,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Greater', current_timestamp, 'Dado dos valores retorna si el primero es mayor al segundo', 'primitive', 'operator', 'public', 'boolean', -38);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-12,current_timestamp,'val1','number','Greater');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-13,current_timestamp,'val2','number','Greater');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-39,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('GreaterEqual', current_timestamp, 'Dado dos valores retorna si el primero es mayor o igual al segundo', 'primitive', 'operator', 'public', 'boolean', -39);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-14,current_timestamp,'val1','number','GreaterEqual');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-15,current_timestamp,'val2','number','GreaterEqual');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-40,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Less', current_timestamp, 'Dado dos valores retorna si el primero es menor al segundo', 'primitive', 'operator', 'public', 'boolean', -40);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-16,current_timestamp,'val1','number','Less');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-17,current_timestamp,'val2','number','Less');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-41,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('LessEqual', current_timestamp, 'Dado dos valores retorna si el primero es menor o igual al segundo', 'primitive', 'operator', 'public', 'boolean', -41);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-42,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Not', current_timestamp, 'Dada un valor de verdad, devuelve el contrario', 'primitive', 'operator', 'public', 'boolean', -42);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-43,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('And', current_timestamp, 'Dados dos valores de verdad, devuelve el valor de verdad de la conjuncion (Solo es verdadero si ambos son verdaderos)', 'primitive', 'operator', 'public', 'boolean', -43);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-18,current_timestamp,'val2','boolean','And');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-44,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Or', current_timestamp, 'Dados dos valores de verdad, devuelve el valor de verdad de la disyuncion no excluyente (Solo es verdadero si alguno es verdadero)', 'primitive', 'operator', 'public', 'boolean', -44);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-19,current_timestamp,'val1','boolean','Or');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-20,current_timestamp,'val2','boolean','Or');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-45,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Percent', current_timestamp, 'Dado un numero y un porcentaje, devuelve el numero que representa ese porcentaje', 'primitive', 'operator', 'public', 'number', -45);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-21,current_timestamp,'val','number','Percent');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-22,current_timestamp,'percent','number','Percent');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-46,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Diff', current_timestamp, 'Devuelve el primer numero menos el segundo numero', 'primitive', 'operator', 'public', 'number', -46);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-23,current_timestamp,'val1','number','Diff');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-24,current_timestamp,'val2','number','Diff');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-47,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Div', current_timestamp, 'Devuelve el primer numero dividido el segundo numero', 'primitive', 'operator', 'public', 'number', -47);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-25,current_timestamp,'val1','number','Div');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-26,current_timestamp,'val2','number','Div');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-48,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Multi', current_timestamp, 'Devuelve el primer numero multiplicado por el segundo numero', 'primitive', 'operator', 'public', 'number', -48);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-27,current_timestamp,'val','number','Multi');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-28,current_timestamp,'percent','number','Multi');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-49,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Jornal', current_timestamp, 'Variable Sueldo / 30', 'primitive', 'generic', 'public', 'number', -49);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-50,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('ValorHora', current_timestamp, 'Variables Sueldo / HorasMensuales', 'primitive', 'generic', 'public', 'number', -50);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-51,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('HoraExtra50', current_timestamp, 'Variable ValorHora * 1,5 * Cantidad (será el campo Cantidad Ingresado en una novedad)', 'primitive', 'generic', 'public', 'number', -51);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-52,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('HoraExtra100', current_timestamp, 'Variable ValorHora * 2 * Cantidad (será el campo Cantidad Ingresado en una novedad)', 'primitive', 'generic', 'public', 'number', -52);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-53,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Sac', current_timestamp, 'Variable (MejorRemRemunerativaBaseSACSemestre / 2) * DiasSemTrabajados / 180', 'primitive', 'generic', 'public', 'number', -53);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-54,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('SacNoRemunerativo', current_timestamp, 'Variable (MejorRemNoRemunerativaSemestre / 2) * DiasSemTrabajados / 180', 'primitive', 'generic', 'public', 'number', -54);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-55,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Antiguedad', current_timestamp, 'Variable (MejorRemNoRemunerativaSemestre / 2) * DiasSemTrabajados / 180', 'primitive', 'generic', 'public', 'number', -55);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-56,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasFaltasInjustificadas', current_timestamp, '', 'primitive', 'generic', 'public', 'number', -56);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-57,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Vacaciones', current_timestamp, 'Si la antiguedad es menor a 5 entonces le corresponden 14 dias de vacaciones. Entre 5 y 10 le corresponden 21 dias de vacaciones. Entre 10 y 15 le corresponden 28 dias de vacaciones. Mayor a 15 le corresponden 35 dias de vacaciones.', 'primitive', 'generic', 'public', 'number', -57);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-58,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Preaviso', current_timestamp, '', 'primitive', 'generic', 'public', 'number', -58);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-59,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('SacSinPreaviso', current_timestamp, 'SAC sobre Preaviso', 'primitive', 'generic', 'public', 'number', -59);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-60,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('IntegracionMesDespido', current_timestamp, '( 30 - DiasMesTrabajados ) * Jornal', 'primitive', 'generic', 'public', 'number', -60);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-61,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('SacSinImd', current_timestamp, 'Integración mes Despido / 12', 'primitive', 'generic', 'public', 'number', -61);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-62,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('IndemnizacionPorDespido', current_timestamp, '', 'primitive', 'generic', 'public', 'number', -62);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-29,current_timestamp,'val','number','IndemnizacionPorDespido');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-63,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('CantidadSueldos', current_timestamp, '', 'primitive', 'generic', 'public', 'number', -63);
    END IF;
END $$;
