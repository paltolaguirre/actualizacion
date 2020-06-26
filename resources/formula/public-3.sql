DO $$
BEGIN
    IF ((select count(*) from param where functionname = 'And') = 0) THEN
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-30,current_timestamp,'val1','number','LessEqual');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-31,current_timestamp,'val2','number','LessEqual');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-32,current_timestamp,'val1','boolean','Not');
        INSERT INTO param(id,created_at, name, type, functionname) VALUES(-33,current_timestamp,'val1','boolean','And');
        delete from function where name in ('HoraExtra50', 'HoraExtra100', 'FechadeIngreso', 'FechadeLiquidacion');
        delete from value where id in (-51, -52, -24, -23);
    END IF;
END $$;