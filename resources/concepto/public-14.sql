DO $$
BEGIN
    IF ((select count(*) from CONCEPTO where id = -34) = 0) THEN
        INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid, eseditable, tipocalculoautomaticoid, formulanombre, esremvariable) VALUES(-33, current_timestamp,'Medicina Prepaga', 'MEDICINA_PREPAGA',  '', 1, '',-46, false, -2,false, null, null, false, true, -2, true, -1, null, false);
        INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid, eseditable, tipocalculoautomaticoid, formulanombre, esremvariable) VALUES(-34, current_timestamp,'Dias De Licencia', 'DIAS_DE_LICENCIA',  '', 1, '',-46, true, -3,true, null, null, false, true, -1, true, -1, null, false);
    END IF;
END $$;