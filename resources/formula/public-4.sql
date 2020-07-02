DO $$
BEGIN
    IF ((select count(*) from function where name = 'DiasDelSemestre') = 0) THEN
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-64,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('DiasDelSemestre', current_timestamp, 'Cantidad de dias corridos del semestre donde se encuentra utilizada la fórmula. Por ejemplo si la variable se utiliza en una liquidación del primer semestre del 2020 el resultado será 182. (utilizada para el cálculo automático de SAC)', 'primitive', 'helper', 'public', 'number', -64);
    END IF;
END $$;