DO $$
BEGIN
    IF ((select count(*) from function where name = 'ImpuestoALasGananciasDevolucion') = 0) THEN
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-1,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('GetParamValue', current_timestamp, 'Obtiene el valor de un parametro de la formula', 'primitive', 'internal', 'public', 'number', -1);
        INSERT INTO param(id, created_at, name, type, functionname) VALUES(-1, current_timestamp, 'paramName', 'string', 'GetParamValue');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-2,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('Sum', current_timestamp, 'Dado dos valores obtiene la suma de ambos', 'primitive', 'operator', 'public', 'number', -2);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-2,current_timestamp,'val1','number','Sum');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-3,current_timestamp,'val2','number','Sum');
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-3,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('TotalImporteRemunerativo', current_timestamp, 'Dada una liquidacion obtiene la suma total de importes remunerativos de la misma', 'primitive', 'helper', 'public', 'number', -3);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-4,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('ImpuestoALasGanancias', current_timestamp, '', 'primitive', 'helper', 'public', 'number', -4);
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-5,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('ImpuestoALasGananciasDevolucion', current_timestamp, '', 'primitive', 'helper', 'public', 'number', -5);
    END IF;
END $$;