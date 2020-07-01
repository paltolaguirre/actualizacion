DO $$
BEGIN
    IF ((select count(*) from function where name = 'BooleanInequality') = 0) THEN
        INSERT INTO value(id, created_at, name, valuenumber, valuestring, valueboolean, valueinvokeid, arginvokeid) VALUES(-65,current_timestamp,'return',0,'',false,null,0);
        INSERT INTO function(name, created_at, description, origin, type, scope, result, valueid) VALUES('BooleanInequality', current_timestamp, 'Dado dos valores booleanos retorna si son distintos', 'primitive', 'operator', 'public', 'boolean', -65);
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-34,current_timestamp,'val1','boolean','BooleanInequality');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-35,current_timestamp,'val2','boolean','BooleanInequality');
    END IF;
END $$;