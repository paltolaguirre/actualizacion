DO $$
BEGIN
INSERT INTO SIRADIGTIPOGRILLA(id, created_at, nombre, codigo, descripcion, activo) 
VALUES  (-31,current_timestamp,'Otras Deducciones - Actores - Retribución pagada a los representantes - RG 2442/08','OTRAS_DEDUCCIONES_ACTORES_RETRIBUCION_PAGADA_A_LOS_REPRESENTANTES_RG_2442_08','',1),
	    (-32,current_timestamp,'Otras Deducciones - Fondos Compensadores de Previsión','OTRAS_DEDUCCIONES_FONDOS_COMPENSADORES_DE_PREVISION','',1);

END 
$$;