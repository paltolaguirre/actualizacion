DO $$
    BEGIN
        IF ((select count(*) from CONCEPTO where id = -35) = 0) THEN
            update concepto set formulanombre = 'Sac', tipocalculoautomaticoid = -3 where nombre = 'Sueldo Anual Complementario';
            INSERT INTO CONCEPTO(id, created_at, nombre, codigo, descripcion, activo, tipo, cuenta_contable, esimprimible, tipoconceptoid, esnovedad, porcentaje, tipodecalculoid, prorrateo, basesac, tipoimpuestogananciasid, eseditable, tipocalculoautomaticoid, formulanombre, esremvariable, conceptoafipid, marcarepeticion, aportesipa, contribucionsipa, aportesinssjyp, contribucionesinssjyp, aportesobrasocial, contribucionesobrasocial, aportesfondosolidario, contribucionesfondosolidario, aportesrenatea, contribucionesrenatea, asignacionesfamiliares, contribucionesfondonacional, contribucionesleyriesgo, aportesregimenesdiferenciales, aportesregimenesespeciales, codigointerno) VALUES(-35, current_timestamp,'Sueldo Anual Complementario No Remunerativo', 'SUELDO_ANUAL_COMPLEMENTARIO_NO_REMUNERATIVO',  '', 1, '',-46, true, -2,false, null, null, null, null, -16, true, -3, 'SacNoRemunerativo', false, -11, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, 18);
            update concepto set formulanombre = 'Vacaciones', tipocalculoautomaticoid = -3 where nombre = 'Vacaciones';
            update concepto set formulanombre = 'Preaviso', tipocalculoautomaticoid = -3 where nombre = 'Preaviso';
            update concepto set formulanombre = 'SacSinPreaviso', tipocalculoautomaticoid = -3 where nombre = 'SAC sobre Preaviso';
            update concepto set formulanombre = 'IntegracionMesDespido', tipocalculoautomaticoid = -3 where nombre = 'Integración Mes de despido';
        END IF;
END $$;