DO $$
    BEGIN
     INSERT INTO TIPOIMPUESTOGANANCIAS(id,created_at, nombre, codigo, descripcion, activo, aplicaimporteremunerativo, aplicaimportenoremunerativo,aplicadescuento, aplicaretencion,aplicaaportepatronal) 
     VALUES (-17,current_timestamp,'SAC Primer Cuota Exentas/No Alcanzadas','SAC_PRIMER_CUOTA_EXENTAS_NO_ALCANZADAS','',1, true, true, true, true, false),
		    (-18,current_timestamp,'SAC Segunda Cuota Exentas/No Alcanzadas','SAC_SEGUNDA_CUOTA_EXENTAS_NO_ALCANZADAS','',1, true, true, true, true, false),
		    (-19,current_timestamp,'Retribuciones No Habituales Exentas/No Alcanzadas','RETRIBUCIONES_NO_HABITUALES_EXENTAS_NO_ALCANZADAS','',1, true, true, true, false, false),
		    (-20,current_timestamp,'Ajustes Períodos Anteriores - Remuneraciones Gravadas','AJUSTES_PERÍODOS_ANTERIORES_REMUNERACIONES_GRAVADAS','',1, true, true, true, false, false),
		    (-21,current_timestamp,'Ajustes Períodos Anteriores - Remuneraciones Exentas/No Alcanzadas','AJUSTES_PERÍODOS_ANTERIORES_REMUNERACIONES_EXENTAS_NO_ALCANZADAS','',1, true, true, true, false, false);
END $$;