insert into tipoconcepto(id, created_at, nombre, codigo, descripcion, activo) values
(-1, current_timestamp, 'Importe Remunerativo', 'IMPORTE_REMUNERATIVO', '', 1),
(-2, current_timestamp, 'Importe No Remunerativo', 'IMPORTE_NO_REMUNERATIVO', '', 1),
(-3, current_timestamp, 'Descuento', 'DESCUENTO', '', 1),
(-4, current_timestamp, 'Retención', 'RETENCION', '', 1),
(-5, current_timestamp, 'Aporte Patronal', 'APORTE_PATRONAL', '', 1);

insert into tipodecalculo(id, created_at, nombre, codigo, descripcion, activo) values
(-1, current_timestamp, 'Remunerativos', 'REMUNERATIVOS', '', 1),
(-2, current_timestamp, 'No Remunerativos', 'NO_REMUNERATIVOS', '', 1),
(-3, current_timestamp, 'Remunerativos - Descuentos', 'REMUNERATIVOS_DESCUENTOS', '', 1),
(-4, current_timestamp, 'Remunerativos + No Remunerativos', 'REMUNERATIVOS_NO_REMUNERATIVOS', '', 1),
(-5, current_timestamp, 'Aporte Remunerativos + No Remunerativos - Descuentos', 'REMUNERATIVOS_NO_REMUNERATIVOS_DESCUENTOS', '', 1);

INSERT INTO concepto(
    id, nombre, cuenta_contable ,created_at ,codigo, descripcion,
    activo, tipo,esimprimible )
VALUES
(-1,'Sueldo Básico',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-2,'Sueldo Anual Complementario',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-3,'Vacaciones',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-4,'Antiguedad',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-5,'Horas Extras 50%',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-6,'Horas Extras 100%',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-7,'Indemnización por Despido',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-8,'Preaviso',	-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-9,'SAC sobre Preaviso',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-10,'SAC sobre Vacaciones No Gozadas',	-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-11,'Vacaciones No Gozadas',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-12,'Gratificación',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-13,'Integración Mes de despido',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-14,'SAC sobre Int. Mes de despido',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-15,'Días Enfermedad',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-16,'Días Accidente',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-17,'Días Faltas injustificadas',-46,current_timestamp, '' , '' , 1 , '',true ) ,
(-18,'Jubilación',-48,current_timestamp, '' , '' , 1 , '',true ) ,
(-19,'Ley 19.032',-48,current_timestamp, '' , '' , 1 , '',true ) ,
(-20,'Obra Social',-48,current_timestamp, '' , '' , 1 , '',true ) ,
(-21,'Ley de Riesgos de Trabajo',-47,current_timestamp, '' , '' , 1 , '',false ) ,
(-22,'Seguro Colectivo de Vida',-47,current_timestamp, '' , '' , 1 , '',false ) ,
(-23,'Vales Alimentarios',-47,current_timestamp, '' , '' , 1 , '',false),
(-24,'Contribuciones RENATRE',-47,current_timestamp, '' , '' , 1 , '',false),
(-25,'Seguro Sepelio UATRE',-47,current_timestamp, '' , '' , 1 , '',false),
(-26,'Jubilación – Aportes Patronales',-47,current_timestamp, '' , '' , 1 , '',false),
(-27,'Ley 19,032 – Aportes Patronales',-47,current_timestamp, '' , '' , 1 , '',false),
(-28,'Obra Social – Aportes Patronales',-47,current_timestamp, '' , '' , 1 , '',false);

CREATE OR REPLACE FUNCTION ST_copy_concepto_public_privado() RETURNS void AS $$
BEGIN
    insert into concepto select * from public.concepto where public.concepto.id not in (select id from concepto);

    EXECUTE (
        SELECT
                            'UPDATE concepto as v SET   (' || string_agg(quote_ident(column_name), ',') || ') = (' || string_agg('public.concepto.' || quote_ident(column_name), ',') || ') FROM public.concepto WHERE v.id = public.concepto.id'
        FROM   information_schema.columns
        WHERE  table_name   = 'concepto'       -- table name, case sensitive
          AND    table_schema = 'public'  -- schema name, case sensitive
          AND    column_name <> 'id'      -- all columns except id
    );

END
$$ LANGUAGE plpgsql;