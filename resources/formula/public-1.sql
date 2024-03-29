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


CREATE OR REPLACE FUNCTION ST_copy_formulas_public_privado() RETURNS void AS $$
BEGIN

    insert into invoke (id) select id from public.invoke where public.invoke.id not in (select id from invoke);
    insert into value select * from public.value where public.value.id not in (select id from value);
    insert into function select * from public.function where public.function.name not in (select name from function);
    insert into param select * from public.param where public.param.id not in (select id from param);



    EXECUTE (
        SELECT
                            'UPDATE value as v SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.value.' || quote_ident(column_name), ',') || ') FROM public.value WHERE v.id = public.value.id'
        FROM   information_schema.columns
        WHERE  table_name   = 'value'       -- table name, case sensitive
          AND    table_schema = 'public'  -- schema name, case sensitive
          AND    column_name <> 'id'      -- all columns except id
    );

    EXECUTE (
        SELECT
                            'UPDATE function as fun SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.function.' || quote_ident(column_name), ',') || ') FROM public.function WHERE fun.name = public.function.name'
        FROM   information_schema.columns
        WHERE  table_name   = 'function'       -- table name, case sensitive
          AND    table_schema = 'public'  -- schema name, case sensitive
          AND    column_name <> 'name'      -- all columns except name
    );

    EXECUTE (
        SELECT
                            'UPDATE param as par SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.param.' || quote_ident(column_name), ',') || ') FROM public.param WHERE par.id = public.param.id'
        FROM   information_schema.columns
        WHERE  table_name   = 'param'       -- table name, case sensitive
          AND    table_schema = 'public'  -- schema name, case sensitive
          AND    column_name <> 'id'      -- all columns except id
    );

    EXECUTE (
        SELECT
                            'UPDATE invoke as v SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.invoke.' || quote_ident(column_name), ',') || ') FROM public.invoke WHERE v.id = public.invoke.id'
        FROM   information_schema.columns
        WHERE  table_name   = 'invoke'       -- table name, case sensitive
          AND    table_schema = 'public'  -- schema name, case sensitive
          AND    column_name <> 'id'      -- all columns except id
    );
END
$$ LANGUAGE plpgsql;