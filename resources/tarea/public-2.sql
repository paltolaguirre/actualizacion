DO $$
BEGIN
    IF ((select count(*) from estadotareaitem where codigo = 'SIN_PROCESAR') = 0) THEN
        insert into estadotareaitem(id, created_at, nombre, codigo, activo) values (-1, current_timestamp, 'Sin procesar', 'SIN_PROCESAR', 1), (-2, current_timestamp, 'Procesando', 'PROCESANDO', 1), (-3, current_timestamp, 'Error', 'ERROR', 1), (-4, current_timestamp, 'Exito', 'EXITO', 1);
        insert into estadotarea(id, created_at, nombre, codigo, activo) values (-1, current_timestamp, 'Procesando', 'PROCESANDO', 1), (-2, current_timestamp, 'Finalizada', 'FINALIZADA', 1), (-3, current_timestamp, 'Cancelada', 'CANCELADA', 1);
    END IF;
END;
$$;