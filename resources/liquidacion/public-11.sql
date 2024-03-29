--funcion sacada de aqui:https://wiki.pstgresql.org/wiki/Numeros_a_letras

CREATE OR REPLACE FUNCTION fu_numero_letras(numero numeric) RETURNS text AS
$body$
DECLARE
    lnEntero bigint;
    lcRetorno TEXT;
    lnTerna bigint;
    lcMiles TEXT;
    lcCadena TEXT;
    lnUnidades bigint;
    lnDecenas bigint;
    lnCentenas bigint;
    lnFraccion bigint;
    lnSw bigint;
BEGIN
    lnEntero := FLOOR(numero)::bigint;--Obtenemos la parte Entera
    lnFraccion := FLOOR(((numero - lnEntero) * 100))::bigint;--Obtenemos la Fraccion del Monto
    lcRetorno := '';
    lnTerna := 1;
    IF lnEntero > 0 THEN
        lnSw := LENGTH(lnEntero::VARCHAR);
        WHILE lnTerna <= lnSw LOOP
                -- Recorro terna por terna
                lcCadena = '';
                lnUnidades = lnEntero % 10;
                lnEntero = CAST(lnEntero/10 AS bigint);
                lnDecenas = lnEntero % 10;
                lnEntero = CAST(lnEntero/10 AS bigint);
                lnCentenas = lnEntero % 10;
                lnEntero = CAST(lnEntero/10 AS bigint);
                -- Analizo las unidades
                SELECT
                    CASE /* UNIDADES */
                        WHEN lnUnidades = 1 AND lnTerna = 1 THEN 'UNO ' || lcCadena
                        WHEN lnUnidades = 1 AND lnTerna <> 1 THEN 'UN ' || lcCadena
                        WHEN lnUnidades = 2 THEN 'DOS ' || lcCadena
                        WHEN lnUnidades = 3 THEN 'TRES ' || lcCadena
                        WHEN lnUnidades = 4 THEN 'CUATRO ' || lcCadena
                        WHEN lnUnidades = 5 THEN 'CINCO ' || lcCadena
                        WHEN lnUnidades = 6 THEN 'SEIS ' || lcCadena
                        WHEN lnUnidades = 7 THEN 'SIETE ' || lcCadena
                        WHEN lnUnidades = 8 THEN 'OCHO ' || lcCadena
                        WHEN lnUnidades = 9 THEN 'NUEVE ' || lcCadena
                        ELSE lcCadena
                        END INTO lcCadena;
                /* UNIDADES */
                -- Analizo las decenas
                SELECT
                    CASE /* DECENAS */
                        WHEN lnDecenas = 1 THEN
                            CASE lnUnidades
                                WHEN 0 THEN 'DIEZ '
                                WHEN 1 THEN 'ONCE '
                                WHEN 2 THEN 'DOCE '
                                WHEN 3 THEN 'TRECE '
                                WHEN 4 THEN 'CATORCE '
                                WHEN 5 THEN 'QUINCE '
                                ELSE 'DIECI' || lcCadena
                                END
                        WHEN lnDecenas = 2 AND lnUnidades = 0 THEN 'VEINTE ' || lcCadena
                        WHEN lnDecenas = 2 AND lnUnidades <> 0 THEN 'VEINTI' || lcCadena
                        WHEN lnDecenas = 3 AND lnUnidades = 0 THEN 'TREINTA ' || lcCadena
                        WHEN lnDecenas = 3 AND lnUnidades <> 0 THEN 'TREINTA Y ' || lcCadena
                        WHEN lnDecenas = 4 AND lnUnidades = 0 THEN 'CUARENTA ' || lcCadena
                        WHEN lnDecenas = 4 AND lnUnidades <> 0 THEN 'CUARENTA Y ' || lcCadena
                        WHEN lnDecenas = 5 AND lnUnidades = 0 THEN 'CINCUENTA ' || lcCadena
                        WHEN lnDecenas = 5 AND lnUnidades <> 0 THEN 'CINCUENTA Y ' || lcCadena
                        WHEN lnDecenas = 6 AND lnUnidades = 0 THEN 'SESENTA ' || lcCadena
                        WHEN lnDecenas = 6 AND lnUnidades <> 0 THEN 'SESENTA Y ' || lcCadena
                        WHEN lnDecenas = 7 AND lnUnidades = 0 THEN 'SETENTA ' || lcCadena
                        WHEN lnDecenas = 7 AND lnUnidades <> 0 THEN 'SETENTA Y ' || lcCadena
                        WHEN lnDecenas = 8 AND lnUnidades = 0 THEN 'OCHENTA ' || lcCadena
                        WHEN lnDecenas = 8 AND lnUnidades <> 0 THEN 'OCHENTA Y ' || lcCadena
                        WHEN lnDecenas = 9 AND lnUnidades = 0 THEN 'NOVENTA ' || lcCadena
                        WHEN lnDecenas = 9 AND lnUnidades <> 0 THEN 'NOVENTA Y ' || lcCadena
                        ELSE lcCadena
                        END INTO lcCadena; /* DECENAS */
                -- Analizo las centenas
                SELECT
                    CASE /* CENTENAS */
                        WHEN lnCentenas = 1 AND lnUnidades = 0 AND lnDecenas = 0 THEN 'CIEN ' || lcCadena
                        WHEN lnCentenas = 1 AND NOT(lnUnidades = 0 AND lnDecenas = 0) THEN 'CIENTO ' || lcCadena
                        WHEN lnCentenas = 2 THEN 'DOSCIENTOS ' || lcCadena
                        WHEN lnCentenas = 3 THEN 'TRESCIENTOS ' || lcCadena
                        WHEN lnCentenas = 4 THEN 'CUATROCIENTOS ' || lcCadena
                        WHEN lnCentenas = 5 THEN 'QUINIENTOS ' || lcCadena
                        WHEN lnCentenas = 6 THEN 'SEISCIENTOS ' || lcCadena
                        WHEN lnCentenas = 7 THEN 'SETECIENTOS ' || lcCadena
                        WHEN lnCentenas = 8 THEN 'OCHOCIENTOS ' || lcCadena
                        WHEN lnCentenas = 9 THEN 'NOVECIENTOS ' || lcCadena
                        ELSE lcCadena
                        END INTO lcCadena;/* CENTENAS */
                -- Analizo la terna
                SELECT
                    CASE /* TERNA */
                        WHEN lnTerna = 1 THEN lcCadena
                        WHEN lnTerna = 2 AND (lnUnidades + lnDecenas + lnCentenas <> 0) THEN lcCadena || ' MIL '
                        WHEN lnTerna = 3 AND (lnUnidades + lnDecenas + lnCentenas <> 0) AND
                             lnUnidades = 1 AND lnDecenas = 0 AND lnCentenas = 0 THEN lcCadena || ' MILLON '
                        WHEN lnTerna = 3 AND (lnUnidades + lnDecenas + lnCentenas <> 0) AND
                             NOT (lnUnidades = 1 AND lnDecenas = 0 AND lnCentenas = 0) THEN lcCadena || ' MILLONES '
                        WHEN lnTerna = 4 AND (lnUnidades + lnDecenas + lnCentenas <> 0) THEN lcCadena || ' MIL MILLONES '
                        ELSE ''
                        END INTO lcCadena;/* TERNA */

                --Retornamos los Valores Obtenidos
                lcRetorno = lcCadena  || lcRetorno;
                lnTerna = lnTerna + 1;
            END LOOP;
    END IF;
    IF lnTerna = 1 THEN
        lcRetorno := 'CERO';
    END IF;
    lcRetorno := RTRIM(lcRetorno::VARCHAR);
    RETURN lcRetorno;
END;
$body$
    LANGUAGE 'plpgsql' VOLATILE CALLED ON NULL INPUT SECURITY INVOKER;

CREATE OR REPLACE FUNCTION getImporteEnLetras(importe numeric)
    RETURNS character varying AS
$BODY$

BEGIN

    importe = importe::numeric(16,2);


    IF split_part(importe::VARCHAR,'.',2) = '00' THEN
        IF split_part(importe::VARCHAR, '.',1) = '1' THEN
            RETURN concat_ws(' ','Pesos uno')::VARCHAR;
        ELSE
            RETURN concat_ws(' ','Pesos', lower(replace(fu_numero_letras(coalesce(split_part(importe::VARCHAR, '.',1),'0')::NUMERIC ),'  ',' ')))::VARCHAR;
        END IF;
    ELSE
        IF split_part(importe::VARCHAR, '.',1) = '1' THEN
            RETURN concat_ws(' ','Pesos uno con', lower(replace(fu_numero_letras(coalesce(split_part(importe::VARCHAR, '.',2),'0')::NUMERIC),'  ',' ')) , 'centavos')::VARCHAR;
        ELSE
            RETURN concat_ws(' ','Pesos', lower(replace(fu_numero_letras(coalesce(split_part(importe::VARCHAR, '.',1),'0')::NUMERIC ),'  ',' ')), 'con' , lower(replace(fu_numero_letras(coalesce(split_part(importe::VARCHAR, '.',2),'0')::NUMERIC),'  ',' ')) , 'centavos')::VARCHAR;
        END IF;
    END IF;

END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;
ALTER FUNCTION getImporteEnLetras(numeric)
    OWNER TO postgres;
