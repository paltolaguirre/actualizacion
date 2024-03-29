DO $$
    DECLARE
        _kind "char";
    BEGIN

        SELECT relkind
        FROM   pg_class
        WHERE  relname = 'deduccionespersonales_id_seq'  -- sequence name
        INTO  _kind;
        IF NOT FOUND THEN       -- name is free
            CREATE SEQUENCE deduccionespersonales_id_seq
                INCREMENT 1
                MINVALUE 1
                MAXVALUE 9223372036854775807
                START 1
                CACHE 1;
            ALTER TABLE deduccionespersonales_id_seq
                OWNER TO postgres;
            GRANT ALL ON SEQUENCE deduccionespersonales_id_seq TO postgres;
        ELSIF _kind = 'S' THEN  -- sequence exists
        -- nada
        ELSE                    -- object name exists for different kind
        -- nada
        END IF;

        CREATE TABLE deduccionespersonales(
                                              id INTEGER NOT NULL DEFAULT nextval('deduccionespersonales_id_seq'::regclass),
                                              created_at TIMESTAMP,
                                              updated_at TIMESTAMP,
                                              deleted_at TIMESTAMP,
                                              nombre VARCHAR,
                                              codigo VARCHAR,
                                              descripcion VARCHAR,
                                              activo int,
                                              valorfijoconyuge NUMERIC(19,4) NOT NULL,
                                              valorfijohijo NUMERIC(19,4) NOT NULL,
                                              valorfijomni NUMERIC(19,4) NOT NULL,
                                              valorfijoddei NUMERIC(19,4) NOT NULL,
                                              anio INT UNIQUE,
                                              CONSTRAINT deduccionespersonales_seq PRIMARY KEY (id)
        )

            WITH (
                OIDS=FALSE
            );

        INSERT INTO deduccionespersonales(created_at,activo,valorfijoconyuge,valorfijohijo,valorfijomni,valorfijoddei,anio)
        VALUES
        (current_timestamp, 1, 80033.90, 40361.43, 103018.79, 494490.18, 2019),
        (current_timestamp, 1, 115471.38, 58232.65, 123861.17, 594533.62, 2020);

        UPDATE public.deduccionespersonales SET id = -id;

        SELECT relkind
        FROM   pg_class
        WHERE  relname = 'escalaimpuestoaplicable_id_seq'  -- sequence name
        INTO  _kind;
        IF NOT FOUND THEN       -- name is free
            CREATE SEQUENCE escalaimpuestoaplicable_id_seq
                INCREMENT 1
                MINVALUE 1
                MAXVALUE 9223372036854775807
                START 1
                CACHE 1;
            ALTER TABLE escalaimpuestoaplicable_id_seq
                OWNER TO postgres;
            GRANT ALL ON SEQUENCE escalaimpuestoaplicable_id_seq TO postgres;
        ELSIF _kind = 'S' THEN  -- sequence exists
        -- nada
        ELSE                    -- object name exists for different kind
        -- nada
        END IF;

        CREATE TABLE escalaimpuestoaplicable(
                                                id INTEGER NOT NULL DEFAULT nextval('escalaimpuestoaplicable_id_seq'::regclass),
                                                created_at TIMESTAMP,
                                                updated_at TIMESTAMP,
                                                deleted_at TIMESTAMP,
                                                nombre VARCHAR,
                                                codigo VARCHAR,
                                                descripcion VARCHAR,
                                                activo int,
                                                limiteinferior NUMERIC(19,4) NOT NULL,
                                                limitesuperior NUMERIC(19,4) NOT NULL,
                                                valorfijo NUMERIC(19,4) NOT NULL,
                                                valorvariable NUMERIC(19,4) NOT NULL,
                                                mesanio VARCHAR NOT NULL,
                                                CONSTRAINT escalaimpuestoaplicableid__seq PRIMARY KEY (id)
        )

            WITH (
                OIDS=FALSE
            );

        INSERT INTO escalaimpuestoaplicable(created_at,activo,limiteinferior,limitesuperior,valorfijo,valorvariable,mesanio)
        VALUES
        (current_timestamp, 1, 0.00, 2753.32, 0.00, 0.05, '01/2019'),
        (current_timestamp, 1, 2753.32, 5506.63, 137.67, 0.09, '01/2019'),
        (current_timestamp, 1, 5506.63, 8259.95, 385.46, 0.12, '01/2019'),
        (current_timestamp, 1, 8259.95, 11013.27, 715.86, 0.15, '01/2019'),
        (current_timestamp, 1, 11013.27, 16519.90, 1128.86, 0.19, '01/2019'),
        (current_timestamp, 1, 16519.90, 22026.54, 2175.12, 0.23, '01/2019'),
        (current_timestamp, 1, 22026.54, 33039.81, 3441.65, 0.27, '01/2019'),
        (current_timestamp, 1, 33039.81, 44053.08, 6415.23, 0.31, '01/2019'),
        (current_timestamp, 1, 44053.08, 9999999999.00, 9829.34, 0.35, '01/2019'),
        (current_timestamp, 1, 0.00, 5506.64, 0.00, 0.05, '02/2019'),
        (current_timestamp, 1, 5506.64, 11013.27, 275.33, 0.09, '02/2019'),
        (current_timestamp, 1, 11013.27, 16519.90, 770.93, 0.12, '02/2019'),
        (current_timestamp, 1, 16519.90, 22026.54, 1431.73, 0.15, '02/2019'),
        (current_timestamp, 1, 22026.54, 33039.81, 2257.72, 0.19, '02/2019'),
        (current_timestamp, 1, 33039.81, 44053.08, 4350.24, 0.23, '02/2019'),
        (current_timestamp, 1, 44053.08, 66079.61, 6883.29, 0.27, '02/2019'),
        (current_timestamp, 1, 66079.61, 88106.15, 12830.46, 0.31, '02/2019'),
        (current_timestamp, 1, 88106.15, 9999999999.00, 19658.69, 0.35, '02/2019'),
        (current_timestamp, 1, 0.00, 8259.95, 0.00, 0.05, '03/2019'),
        (current_timestamp, 1,8259.95, 16519.90, 413.00, 0.09, '03/2019'),
        (current_timestamp, 1, 16519.90, 24779.86, 1156.39, 0.12, '03/2019'),
        (current_timestamp, 1, 24779.86, 33039.81, 2147.59, 0.15, '03/2019'),
        (current_timestamp, 1, 33039.81, 49559.71, 3386.58, 0.19, '03/2019'),
        (current_timestamp, 1, 49559.71, 66079.61, 6525.36, 0.23, '03/2019'),
        (current_timestamp, 1, 66079.61, 99119.42, 10324.94, 0.27, '03/2019'),
        (current_timestamp, 1, 99119.42, 132159.23, 19245.69, 0.31, '03/2019'),
        (current_timestamp, 1, 132159.23, 9999999999.00, 29488.03, 0.35, '03/2019'),
        (current_timestamp, 1, 0.00, 11013.27, 0.00, 0.05, '04/2019'),
        (current_timestamp, 1, 11013.27, 22026.54, 550.66, 0.09, '04/2019'),
        (current_timestamp, 1, 22026.54, 33039.81, 1541.86, 0.12, '04/2019'),
        (current_timestamp, 1, 33039.81, 44053.08, 2863.45, 0.15, '04/2019'),
        (current_timestamp, 1, 44053.08, 66079.61, 4515.44, 0.19, '04/2019'),
        (current_timestamp, 1, 66079.61, 88106.15, 8700.48, 0.23, '04/2019'),
        (current_timestamp, 1, 88106.15, 132159.23, 13766.59, 0.27, '04/2019'),
        (current_timestamp, 1, 132159.23, 176212.30, 25660.92, 0.31, '04/2019'),
        (current_timestamp, 1, 176212.30, 9999999999.00, 39317.37, 0.35, '04/2019'),
        (current_timestamp, 1, 0.00, 13766.59, 0.00, 0.05, '05/2019'),
        (current_timestamp, 1, 13766.59, 27533.17, 688.33, 0.09, '05/2019'),
        (current_timestamp, 1, 27533.17, 41299.76, 1927.32, 0.12, '05/2019'),
        (current_timestamp, 1, 41299.76, 55066.35, 3579.31, 0.15, '05/2019'),
        (current_timestamp, 1, 55066.35, 82599.52, 5644.30, 0.19, '05/2019'),
        (current_timestamp, 1, 82599.52, 110132.69, 10875.60, 0.23, '05/2019'),
        (current_timestamp, 1, 110132.69, 165199.03, 17208.23, 0.27, '05/2019'),
        (current_timestamp, 1, 165199.03, 220265.38, 32076.15, 0.31, '05/2019'),
        (current_timestamp, 1, 220265.38, 9999999999.00, 49146.71, 0.35, '05/2019'),
        (current_timestamp, 1, 0.00, 16519.91, 0.00, 0.05, '06/2019'),
        (current_timestamp, 1, 16519.91, 33039.81, 826.00, 0.09, '06/2019'),
        (current_timestamp, 1, 33039.81, 49559.71, 2312.79, 0.12, '06/2019'),
        (current_timestamp, 1, 49559.71, 66079.62, 4295.18, 0.15, '06/2019'),
        (current_timestamp, 1, 66079.62, 99119.42, 6773.16, 0.19, '06/2019'),
        (current_timestamp, 1, 99119.42, 132159.23, 13050.73, 0.23, '06/2019'),
        (current_timestamp, 1, 132159.23, 198238.84, 20649.88, 0.27, '06/2019'),
        (current_timestamp, 1, 198238.84, 264318.46, 38491.38, 0.31, '06/2019'),
        (current_timestamp, 1, 264318.46, 9999999999.00, 58976.06, 0.35, '06/2019'),
        (current_timestamp, 1, 0.00, 19273.22, 0.00, 0.05, '07/2019'),
        (current_timestamp, 1, 19273.22, 38546.44, 963.66, 0.09, '07/2019'),
        (current_timestamp, 1, 38546.44, 57819.66, 2698.25, 0.12, '07/2019'),
        (current_timestamp, 1, 57819.66, 77092.88, 5011.04, 0.15, '07/2019'),
        (current_timestamp, 1, 77092.88, 115639.32, 7902.02, 0.19, '07/2019'),
        (current_timestamp, 1, 115639.32, 154185.76, 15225.85, 0.23, '07/2019'),
        (current_timestamp, 1, 154185.76, 231278.65, 24091.53, 0.27, '07/2019'),
        (current_timestamp, 1, 231278.65, 308371.53, 44906.60, 0.31, '07/2019'),
        (current_timestamp, 1, 308371.53, 9999999999.00, 68805.40, 0.35, '07/2019'),
        (current_timestamp, 1, 0.00, 22026.54, 0.00, 0.05, '08/2019'),
        (current_timestamp, 1, 22026.54, 44053.07, 1101.33, 0.09, '08/2019'),
        (current_timestamp, 1, 44053.07, 66079.61, 3083.71, 0.12, '08/2019'),
        (current_timestamp, 1, 66079.61, 88106.15, 5726.90, 0.15, '08/2019'),
        (current_timestamp, 1, 88106.15, 132159.23, 9030.88, 0.19, '08/2019'),
        (current_timestamp, 1, 132159.23, 176212.30, 17400.97, 0.23, '08/2019'),
        (current_timestamp, 1, 176212.30, 264318.45, 27533.17, 0.27, '08/2019'),
        (current_timestamp, 1, 264318.45, 352424.61, 51321.83, 0.31, '08/2019'),
        (current_timestamp, 1, 352424.61, 9999999999.00, 78634.74, 0.35, '08/2019'),
        (current_timestamp, 1, 0.00, 24779.86, 0.00, 0.05, '09/2019'),
        (current_timestamp, 1, 24779.86, 49559.71, 1238.99, 0.09, '09/2019'),
        (current_timestamp, 1, 49559.71, 74339.57, 3469.18, 0.12, '09/2019'),
        (current_timestamp, 1, 74339.57, 99119.42, 6442.76, 0.15, '09/2019'),
        (current_timestamp, 1, 99119.42, 148679.13, 10159.74, 0.19, '09/2019'),
        (current_timestamp, 1, 148679.13, 198238.84, 19576.09, 0.23, '09/2019'),
        (current_timestamp, 1, 198238.84, 297358.26, 30974.82, 0.27, '09/2019'),
        (current_timestamp, 1, 297358.26, 396477.68, 57737.06, 0.31, '09/2019'),
        (current_timestamp, 1, 396477.68, 9999999999.00, 88464.08, 0.35, '09/2019'),
        (current_timestamp, 1, 0.00, 27533.18, 0.00, 0.05, '10/2019'),
        (current_timestamp, 1, 27533.18, 55066.34, 1376.66, 0.09, '10/2019'),
        (current_timestamp, 1, 55066.34, 82599.52, 3854.64, 0.12, '10/2019'),
        (current_timestamp, 1, 82599.52, 110132.69, 7158.63, 0.15, '10/2019'),
        (current_timestamp, 1, 110132.69, 165199.03, 11288.60, 0.19, '10/2019'),
        (current_timestamp, 1, 165199.03, 220265.38, 21751.21, 0.23, '10/2019'),
        (current_timestamp, 1, 220265.38, 330398.07, 34416.47, 0.27, '10/2019'),
        (current_timestamp, 1, 330398.07, 440530.76, 64152.29, 0.31, '10/2019'),
        (current_timestamp, 1, 440530.76, 9999999999.00, 98293.42, 0.35, '10/2019'),
        (current_timestamp, 1, 0.00, 30286.49, 0.00, 0.05, '11/2019'),
        (current_timestamp, 1, 30286.49, 60572.98, 1514.32, 0.09, '11/2019'),
        (current_timestamp, 1, 60572.98, 90859.47, 4240.11, 0.12, '11/2019'),
        (current_timestamp, 1, 90859.47, 121145.96, 7874.49, 0.15, '11/2019'),
        (current_timestamp, 1, 121145.96, 181718.94, 12417.46, 0.19, '11/2019'),
        (current_timestamp, 1, 181718.94, 242291.91, 23926.33, 0.23, '11/2019'),
        (current_timestamp, 1, 242291.91, 363437.87, 37858.11, 0.27, '11/2019'),
        (current_timestamp, 1, 363437.87, 484583.83, 70567.52, 0.31, '11/2019'),
        (current_timestamp, 1, 484583.83, 9999999999.00, 108122.77, 0.35, '11/2019'),
        (current_timestamp, 1, 0.00, 33039.81, 0.00, 0.05, '12/2019'),
        (current_timestamp, 1, 33039.81, 66079.61, 1651.99, 0.09, '12/2019'),
        (current_timestamp, 1, 66079.61, 99119.42, 4625.57, 0.12, '12/2019'),
        (current_timestamp, 1, 99119.42, 132159.23, 8590.35, 0.15, '12/2019'),
        (current_timestamp, 1, 132159.23, 198238.84, 13546.32, 0.19, '12/2019'),
        (current_timestamp, 1, 198238.84, 264318.45, 26101.45, 0.23, '12/2019'),
        (current_timestamp, 1, 264318.45, 396477.68, 41299.76, 0.27, '12/2019'),
        (current_timestamp, 1, 396477.68, 528636.91, 76982.75, 0.31, '12/2019'),
        (current_timestamp, 1, 528636.91, 9999999999.00, 117952.11, 0.35, '12/2019'),
        (current_timestamp, 1, 0.00, 3972.43, 0.00, 0.05, '01/2020'),
        (current_timestamp, 1, 3972.43, 7944.86, 198.62, 0.09, '01/2020'),
        (current_timestamp, 1, 7944.86, 11917.29, 556.14, 0.12, '01/2020'),
        (current_timestamp, 1, 11917.29, 15889.72, 1032.83, 0.15, '01/2020'),
        (current_timestamp, 1, 15889.72, 23834.58, 1628.70, 0.19, '01/2020'),
        (current_timestamp, 1, 23834.58, 31779.44, 3138.22, 0.23, '01/2020'),
        (current_timestamp, 1, 31779.44, 47669.16, 4965.54, 0.27, '01/2020'),
        (current_timestamp, 1, 47669.16, 63558.88, 9255.76, 0.31, '01/2020'),
        (current_timestamp, 1, 63558.88, 9999999999.00, 14181.58, 0.35, '01/2020'),
        (current_timestamp, 1, 0.00, 7944.86, 0.00, 0.05, '02/2020'),
        (current_timestamp, 1, 7944.86,	15889.72, 397.24, 0.09, '02/2020'),
        (current_timestamp, 1, 15889.72, 23834.58, 1112.28, 0.12, '02/2020'),
        (current_timestamp, 1, 23834.58, 31779.44, 2065.66, 0.15, '02/2020'),
        (current_timestamp, 1, 31779.44, 47669.16, 3257.39, 0.19, '02/2020'),
        (current_timestamp, 1, 47669.16, 63558.88, 6276.44, 0.23, '02/2020'),
        (current_timestamp, 1, 63558.88, 95338.32, 9931.07, 0.27, '02/2020'),
        (current_timestamp, 1, 95338.32, 127117.76, 18511.52, 0.31, '02/2020'),
        (current_timestamp, 1, 127117.76, 9999999999.00, 28363.15, 0.35, '02/2020'),
        (current_timestamp, 1, 0.00, 11917.29, 0.00, 0.05, '03/2020'),
        (current_timestamp, 1, 11917.29, 23834.58, 595.86, 0.09, '03/2020'),
        (current_timestamp, 1, 23834.58, 35751.87, 1668.42, 0.12, '03/2020'),
        (current_timestamp, 1, 35751.87, 47669.16, 3098.49, 0.15, '03/2020'),
        (current_timestamp, 1, 47669.16, 71503.74, 4886.09, 0.19, '03/2020'),
        (current_timestamp, 1, 71503.74, 95338.32, 9414.66, 0.23, '03/2020'),
        (current_timestamp, 1, 95338.32, 143007.48, 14896.61, 0.27, '03/2020'),
        (current_timestamp, 1, 143007.48, 190676.64, 27767.29, 0.31, '03/2020'),
        (current_timestamp, 1, 190676.64, 9999999999.00, 42544.73, 0.35, '03/2020'),
        (current_timestamp, 1, 0.00, 15889.72, 0.00, 0.05, '04/2020'),
        (current_timestamp, 1, 15889.72, 31779.44, 794.49, 0.09, '04/2020'),
        (current_timestamp, 1, 31779.44, 47669.16, 2224.56, 0.12, '04/2020'),
        (current_timestamp, 1, 47669.16, 63558.88, 4131.33, 0.15, '04/2020'),
        (current_timestamp, 1, 63558.88, 95338.32, 6514.79, 0.19, '04/2020'),
        (current_timestamp, 1, 95338.32, 127117.76, 12552.88, 0.23, '04/2020'),
        (current_timestamp, 1, 127117.76, 190676.64, 19862.15, 0.27, '04/2020'),
        (current_timestamp, 1, 190676.64, 254235.52, 37023.05, 0.31, '04/2020'),
        (current_timestamp, 1, 254235.52, 9999999999.00, 56726.30, 0.35, '04/2020'),
        (current_timestamp, 1, 0.00, 19862.15, 0.00, 0.05, '05/2020'),
        (current_timestamp, 1, 19862.15, 39724.30, 993.11, 0.09, '05/2020'),
        (current_timestamp, 1, 39724.30, 59586.45, 2780.70, 0.12, '05/2020'),
        (current_timestamp, 1, 59586.45, 79448.60, 5164.16, 0.15, '05/2020'),
        (current_timestamp, 1, 79448.60, 119172.90, 8143.48, 0.19, '05/2020'),
        (current_timestamp, 1, 119172.90, 158897.20, 15691.10, 0.23, '05/2020'),
        (current_timestamp, 1, 158897.20, 238345.80, 24827.69, 0.27, '05/2020'),
        (current_timestamp, 1, 238345.80, 317794.40, 46278.81, 0.31, '05/2020'),
        (current_timestamp, 1, 317794.40, 9999999999.00, 70907.88, 0.35, '05/2020'),
        (current_timestamp, 1, 0.00, 23834.58, 0.00, 0.05, '06/2020'),
        (current_timestamp, 1, 23834.58, 47669.16, 1191.73, 0.09, '06/2020'),
        (current_timestamp, 1, 47669.16, 71503.74, 3336.84, 0.12, '06/2020'),
        (current_timestamp, 1, 71503.74, 95338.33, 6196.99, 0.15, '06/2020'),
        (current_timestamp, 1, 95338.33, 143007.48, 9772.18, 0.19, '06/2020'),
        (current_timestamp, 1, 143007.48, 190676.64, 18829.32, 0.23, '06/2020'),
        (current_timestamp, 1, 190676.64, 286014.96, 29793.22, 0.27, '06/2020'),
        (current_timestamp, 1, 286014.96, 381353.29, 55534.57, 0.31, '06/2020'),
        (current_timestamp, 1, 381353.29, 9999999999.00, 85089.45, 0.35, '06/2020'),
        (current_timestamp, 1, 0.00, 27807.01, 0.00, 0.05, '07/2020'),
        (current_timestamp, 1, 27807.01, 55614.02, 1390.35, 0.09, '07/2020'),
        (current_timestamp, 1, 55614.02, 83421.03, 3892.98, 0.12, '07/2020'),
        (current_timestamp, 1, 83421.03, 111228.05, 7229.82, 0.15, '07/2020'),
        (current_timestamp, 1, 111228.05, 166842.06, 11400.87, 0.19, '07/2020'),
        (current_timestamp, 1, 166842.06, 222456.08, 21967.54, 0.23, '07/2020'),
        (current_timestamp, 1, 222456.08, 333684.12, 34758.76, 0.27, '07/2020'),
        (current_timestamp, 1, 333684.12, 444912.17, 64790.33, 0.31, '07/2020'),
        (current_timestamp, 1, 444912.17, 9999999999.00, 99271.03, 0.35, '07/2020'),
        (current_timestamp, 1, 0.00, 31779.44, 0.00, 0.05, '08/2020'),
        (current_timestamp, 1, 31779.44, 63558.88, 1588.97, 0.09, '08/2020'),
        (current_timestamp, 1, 63558.88, 95338.32, 4449.12, 0.12, '08/2020'),
        (current_timestamp, 1, 95338.32, 127117.77, 8262.65, 0.15, '08/2020'),
        (current_timestamp, 1, 127117.77, 190676.64, 13029.57, 0.19, '08/2020'),
        (current_timestamp, 1, 190676.64, 254235.52, 25105.76, 0.23, '08/2020'),
        (current_timestamp, 1, 254235.52, 381353.28, 39724.30, 0.27, '08/2020'),
        (current_timestamp, 1, 381353.28, 508471.05, 74046.09, 0.31, '08/2020'),
        (current_timestamp, 1, 508471.05, 9999999999.00, 113452.60, 0.35, '08/2020'),
        (current_timestamp, 1, 0.00, 35751.87, 0.00, 0.05, '09/2020'),
        (current_timestamp, 1, 35751.87, 71503.74, 1787.59, 0.09, '09/2020'),
        (current_timestamp, 1, 71503.74, 107255.61, 5005.26, 0.12, '09/2020'),
        (current_timestamp, 1, 107255.61, 143007.49, 9295.49, 0.15, '09/2020'),
        (current_timestamp, 1, 143007.49, 214511.22, 14658.27, 0.19, '09/2020'),
        (current_timestamp, 1, 214511.22, 286014.96, 28243.98, 0.23, '09/2020'),
        (current_timestamp, 1, 286014.96, 429022.44, 44689.84, 0.27, '09/2020'),
        (current_timestamp, 1, 429022.44, 572029.93, 83301.86, 0.31, '09/2020'),
        (current_timestamp, 1, 572029.93, 9999999999.00, 127634.18, 0.35, '09/2020'),
        (current_timestamp, 1, 0.00, 39724.30, 0.00, 0.05, '10/2020'),
        (current_timestamp, 1, 39724.30, 79448.60, 1986.21, 0.09, '10/2020'),
        (current_timestamp, 1, 79448.60, 119172.90, 5561.40, 0.12, '10/2020'),
        (current_timestamp, 1, 119172.90, 158897.21, 10328.32, 0.15, '10/2020'),
        (current_timestamp, 1, 158897.21, 238345.80, 16286.96, 0.19, '10/2020'),
        (current_timestamp, 1, 238345.80, 317794.40, 31382.20, 0.23, '10/2020'),
        (current_timestamp, 1, 317794.40, 476691.60, 49655.37, 0.27, '10/2020'),
        (current_timestamp, 1, 476691.60, 635588.81, 92557.62, 0.31, '10/2020'),
        (current_timestamp, 1, 635588.81, 9999999999.00, 141815.75, 0.35, '10/2020'),
        (current_timestamp, 1, 0.00, 43696.73, 0.00, 0.05, '11/2020'),
        (current_timestamp, 1, 43696.73, 87393.46, 2184.84, 0.09, '11/2020'),
        (current_timestamp, 1, 87393.46, 131090.19, 6117.54, 0.12, '11/2020'),
        (current_timestamp, 1, 131090.19, 174786.93, 11361.15, 0.15, '11/2020'),
        (current_timestamp, 1, 174786.93, 262180.38, 17915.66, 0.19, '11/2020'),
        (current_timestamp, 1, 262180.38, 349573.84, 34520.42, 0.23, '11/2020'),
        (current_timestamp, 1, 349573.84, 524360.76, 54620.91, 0.27, '11/2020'),
        (current_timestamp, 1, 524360.76, 699147.69, 101813.38, 0.31, '11/2020'),
        (current_timestamp, 1, 699147.69, 9999999999.00, 155997.33, 0.35, '11/2020'),
        (current_timestamp, 1, 0.00, 47669.16, 0.00, 0.05, '12/2020'),
        (current_timestamp, 1, 47669.16, 95338.32, 2383.46,	0.09, '12/2020'),
        (current_timestamp, 1, 95338.32, 143007.48,	6673.68, 0.12, '12/2020'),
        (current_timestamp, 1, 143007.48, 190676.65, 12393.98, 0.15, '12/2020'),
        (current_timestamp, 1, 190676.65, 286014.96, 19544.36, 0.19, '12/2020'),
        (current_timestamp, 1, 286014.96, 381353.28, 37658.64, 0.23, '12/2020'),
        (current_timestamp, 1, 381353.28, 572029.92, 59586.45, 0.27, '12/2020'),
        (current_timestamp, 1, 572029.92, 762706.57, 111069.14,	0.31, '12/2020'),
        (current_timestamp, 1, 762706.57, 9999999999.00, 170178.90, 0.35, '12/2020');

        UPDATE public.escalaimpuestoaplicable SET id = -id;

        SELECT relkind
        FROM   pg_class
        WHERE  relname = 'topemaximodescuento_id_seq'  -- sequence name
        INTO  _kind;
        IF NOT FOUND THEN       -- name is free
            CREATE SEQUENCE topemaximodescuento_id_seq
                INCREMENT 1
                MINVALUE 1
                MAXVALUE 9223372036854775807
                START 1
                CACHE 1;
            ALTER TABLE topemaximodescuento_id_seq
                OWNER TO postgres;
            GRANT ALL ON SEQUENCE topemaximodescuento_id_seq TO postgres;
        ELSIF _kind = 'S' THEN  -- sequence exists
        -- nada
        ELSE                    -- object name exists for different kind
        -- nada
        END IF;

        CREATE TABLE topemaximodescuento(
                                            id INTEGER NOT NULL DEFAULT nextval('topemaximodescuento_id_seq'::regclass),
                                            created_at TIMESTAMP,
                                            updated_at TIMESTAMP,
                                            deleted_at TIMESTAMP,
                                            nombre VARCHAR,
                                            codigo VARCHAR,
                                            descripcion VARCHAR,
                                            activo int,
                                            topecasomuerte NUMERIC(19,4) NOT NULL,
                                            topeseguroahorro NUMERIC(19,4) NOT NULL,
                                            toperetiroprivado NUMERIC(19,4) NOT NULL,
                                            topesepelio NUMERIC(19,4) NOT NULL,
                                            topehipotecarios NUMERIC(19,4) NOT NULL,
                                            anio INT UNIQUE,
                                            CONSTRAINT topemaximodescuento_seq PRIMARY KEY (id)
        )

            WITH (
                OIDS=FALSE
            );

        INSERT INTO topemaximodescuento(created_at,activo,topecasomuerte,topeseguroahorro,toperetiroprivado,topesepelio,topehipotecarios,anio)
        VALUES
        (current_timestamp, 1, 12000.00, 12000.00, 12000.00, 993.26, 20000.00, 2019),
        (current_timestamp, 1, 18000.00, 18000.00, 18000.00, 993.26, 20000.00, 2020);

        UPDATE public.topemaximodescuento SET id = -id;

        SELECT relkind
        FROM   pg_class
        WHERE  relname = 'tipoliquidacion_id_seq'  -- sequence name
        INTO  _kind;
        IF NOT FOUND THEN       -- name is free
            CREATE SEQUENCE tipoliquidacion_id_seq
                INCREMENT 1
                MINVALUE 1
                MAXVALUE 9223372036854775807
                START 1
                CACHE 1;
            ALTER TABLE tipoliquidacion_id_seq
                OWNER TO postgres;
            GRANT ALL ON SEQUENCE tipoliquidacion_id_seq TO postgres;
        ELSIF _kind = 'S' THEN  -- sequence exists
        -- nada
        ELSE                    -- object name exists for different kind
        -- nada
        END IF;


        CREATE TABLE public.tipoliquidacion (
                                                id INTEGER NOT NULL DEFAULT nextval('tipoliquidacion_id_seq'::regclass),
                                                created_at timestamp with time zone,
                                                updated_at timestamp with time zone,
                                                deleted_at timestamp with time zone,
                                                nombre text,
                                                codigo text,
                                                descripcion text,
                                                activo integer
        );

        INSERT INTO liquidaciontipo(id,nombre, codigo, activo,created_at)
        VALUES
        (-1,'Mensual','MENSUAL',1,current_timestamp),
        (-2,'Primer Quincena','PRIMER_QUINCENA',1,current_timestamp),
        (-3,'Segunda Quincena','SEGUNDA_QUINCENA',1,current_timestamp),
        (-4,'Vacaciones','VACACIONES',1,current_timestamp),
        (-5,'SAC','SAC',1,current_timestamp),
        (-6,'Liquidación Final','LIQUIDACION_FINAL',1,current_timestamp);

        UPDATE public.tipoliquidacion SET id = -id;

        INSERT INTO liquidacioncondicionpago(id,nombre,codigo,activo,created_at)
        VALUES
        (-1,'Cuenta Corriente','CUENTA_CORRIENTE',1,current_timestamp),
        (-2,'Contado','CONTADO',1,current_timestamp);

        IF NOT EXISTS (SELECT ID
                       FROM CONCEPTOAFIP
                       LIMIT 1)
        THEN
            INSERT INTO CONCEPTOAFIP(id, created_at, nombre, codigo, descripcion, activo, tipoconceptoid)
            VALUES(-1, current_timestamp,'Sueldo', '110000', 'Sueldo', 1, -1),
                  (-2, current_timestamp,'Preaviso', '110001', 'Preaviso', 1, -1),
                  (-3, current_timestamp,'Remuneraciones en especie', '110002', 'Remuneraciones en especie', 1, -1),
                  (-4, current_timestamp,'Comida', '110003', 'Comida', 1, -1),
                  (-5, current_timestamp,'Habitacion', '110004', 'Habitacion', 1, -1),
                  (-6, current_timestamp,'Licencias por estudio', '110005', 'Licencias por estudio', 1, -1),
                  (-7, current_timestamp,'Donacion de sangre', '110006', 'Donacion de sangre', 1, -1),
                  (-8, current_timestamp,'Feriado', '110007', 'Feriado', 1, -1),
                  (-9, current_timestamp,'Prest. Dineraria Ley 24577 (primeros 10d)', '110008', 'Prest. Dineraria Ley 24577 (primeros 10d)', 1, -1),
                  (-10, current_timestamp,'Prest. Dineraria Ley 24577 (a cargo de ART)', '110009', 'Prest. Dineraria Ley 24577 (a cargo de ART)', 1, -1),
                  (-11, current_timestamp,'Sueldo anual complementario', '120000', 'Sueldo anual complementario', 1, -1),
                  (-12, current_timestamp,'SAC 1er semestre', '120001', 'SAC 1er semestre', 1, -1),
                  (-13, current_timestamp,'SAC 2do semestre', '120002', 'SAC 2do semestre', 1, -1),
                  (-14, current_timestamp,'SAC proporcional', '120003', 'SAC proporcional', 1, -1),
                  (-15, current_timestamp,'Horas extras', '130000', 'Horas extras', 1, -1),
                  (-16, current_timestamp,'Horas extras 50%', '130001', 'Horas extras 50%', 1, -1),
                  (-17, current_timestamp,'Horas extras 100%', '130002', 'Horas extras 100%', 1, -1),
                  (-18, current_timestamp,'Horas extras 200%', '130003', 'Horas extras 200%', 1, -1),
                  (-19, current_timestamp,'Zona desfavorable', '140000', 'Zona desfavorable', 1, -1),
                  (-20, current_timestamp,'Adelanto vacacional', '150000', 'Adelanto vacacional', 1, -1),
                  (-21, current_timestamp,'Adicionales', '160000', 'Adicionales', 1, -1),
                  (-22, current_timestamp,'Adicional por antigüedad', '160001', 'Adicional por antigüedad', 1, -1),
                  (-23, current_timestamp,'Adicional por titulo', '160002', 'Adicional por titulo', 1, -1),
                  (-24, current_timestamp,'Adicional por tarea', '160003', 'Adicional por tarea', 1, -1),
                  (-25, current_timestamp,'Adicional por desarraigo', '160004', 'Adicional por desarraigo', 1, -1),
                  (-26, current_timestamp,'Gratificaciones y/o Premios', '170000', 'Gratificaciones y/o Premios', 1, -1),
                  (-27, current_timestamp,'Premio por presentismo', '170001', 'Premio por presentismo', 1, -1),
                  (-28, current_timestamp,'Premio por produccion', '170002', 'Premio por produccion', 1, -1),
                  (-29, current_timestamp,'Comisiones', '170003', 'Comisiones', 1, -1),
                  (-30, current_timestamp,'Accesorios', '170004', 'Accesorios', 1, -1),
                  (-31, current_timestamp,'Viáticos sin comprobante', '170005', 'Viáticos sin comprobante', 1, -1),
                  (-32, current_timestamp,'Propinas habituales no prohibidas', '170006', 'Propinas habituales no prohibidas', 1, -1),
                  (-33, current_timestamp,'Redondeo (Remunerativo)', '499999', 'Redondeo (Remunerativo)', 1, -1),
                  (-34, current_timestamp,'Asignaciones Familiares', '510000', 'Asignaciones Familiares', 1, -2),
                  (-35, current_timestamp,'Ayuda escolar', '510001', 'Ayuda escolar', 1, -2),
                  (-36, current_timestamp,'Asignacion por hijo/hijo con discapacidad', '510002', 'Asignacion por hijo/hijo con discapacidad', 1, -2),
                  (-37, current_timestamp,'Asignacion por maternidad', '510003', 'Asignacion por maternidad', 1, -2),
                  (-38, current_timestamp,'Asignacion por maternidad down', '510004', 'Asignacion por maternidad down', 1, -2),
                  (-39, current_timestamp,'Asignacion por matrimonio', '510005', 'Asignacion por matrimonio', 1, -2),
                  (-40, current_timestamp,'Asignacion por nacimiento / adopcion', '510006', 'Asignacion por nacimiento / adopcion', 1, -2),
                  (-41, current_timestamp,'Asignacion por prenatal', '510007', 'Asignacion por prenatal', 1, -2),
                  (-42, current_timestamp,'Beneficios sociales', '520000', 'Beneficios sociales', 1, -2),
                  (-43, current_timestamp,'Servicio de comedor', '520001', 'Servicio de comedor', 1, -2),
                  (-44, current_timestamp,'Gastos medicos', '520002', 'Gastos medicos', 1, -2),
                  (-45, current_timestamp,'Provision de ropa de trabajo', '520003', 'Provision de ropa de trabajo', 1, -2),
                  (-46, current_timestamp,'Guarderia', '520004', 'Guarderia', 1, -2),
                  (-47, current_timestamp,'Provision de utiles escolares', '520005', 'Provision de utiles escolares', 1, -2),
                  (-48, current_timestamp,'Gastos de sepelio', '520006', 'Gastos de sepelio', 1, -2),
                  (-49, current_timestamp,'Cursos de capacitacion', '520007', 'Cursos de capacitacion', 1, -2),
                  (-50, current_timestamp,'Becas (art. 7 Ley 24.241 y modif.)', '520008', 'Becas (art. 7 Ley 24.241 y modif.)', 1, -2),
                  (-51, current_timestamp,'Desempleo (art. 7 Ley 24.241 y modif.)', '520009', 'Desempleo (art. 7 Ley 24.241 y modif.)', 1, -2),
                  (-52, current_timestamp,'Gratificacion por cese laboral (art. 7 Ley 24.241 y modif.)', '520010', 'Gratificacion por cese laboral (art. 7 Ley 24.241 y modif.)', 1, -2),
                  (-53, current_timestamp,'Indemnizacion por extincion del contrato de trabajo (art. 7 Ley 24.241 y modif.)', '520011', 'Indemnizacion por extincion del contrato de trabajo (art. 7 Ley 24.241 y modif.)', 1, -2),
                  (-54, current_timestamp,'Vacaciones no gozadas (art. 7 Ley 24.241 y modif.)', '520012', 'Vacaciones no gozadas (art. 7 Ley 24.241 y modif.)', 1, -2),
                  (-55, current_timestamp,'Incapacidad permanente (art. 7 Ley 24.241 y modif.)', '520013', 'Incapacidad permanente (art. 7 Ley 24.241 y modif.)', 1, -2),
                  (-56, current_timestamp,'Indemizacion por despido', '520014', 'Indemizacion por despido', 1, -2),
                  (-57, current_timestamp,'Indemizacion sustitutiva del preaviso', '520015', 'Indemizacion sustitutiva del preaviso', 1, -2),
                  (-58, current_timestamp,'Integracion mes de despido', '520016', 'Integracion mes de despido', 1, -2),
                  (-59, current_timestamp,'SAC sobre integracion o preaviso', '520017', 'SAC sobre integracion o preaviso', 1, -2),
                  (-60, current_timestamp,'SAC sobre vacaciones no gozadas', '520018', 'SAC sobre vacaciones no gozadas', 1, -2),
                  (-61, current_timestamp,'Incrementos no remunerativos (con aportes OS)', '530000', 'Incrementos no remunerativos (con aportes OS)', 1, -2),
                  (-62, current_timestamp,'Incrementos no remunerativos (con aportes y contribuciones OS)', '540000', 'Incrementos no remunerativos (con aportes y contribuciones OS)', 1, -2),
                  (-63, current_timestamp,'Importes no remunerativos especiales', '550000', 'Importes no remunerativos especiales', 1, -2),
                  (-64, current_timestamp,'Redondeo (No Remunerativo)', '799999', 'Redondeo (No Remunerativo)', 1, -2),
                  (-65, current_timestamp,'Sistema previsional', '810000', 'Sistema previsional', 1, -4),
                  (-66, current_timestamp,'INSSJyP', '810001', 'INSSJyP', 1, -4),
                  (-67, current_timestamp,'Obra Social', '810002', 'Obra Social', 1, -4),
                  (-68, current_timestamp,'Fondo Solidario de Redistribucion (ex ANSSAL)', '810003', 'Fondo Solidario de Redistribucion (ex ANSSAL)', 1, -3),
                  (-69, current_timestamp,'Cuota Sindical', '810004', 'Cuota Sindical', 1, -4),
                  (-70, current_timestamp,'Seguro de Vida', '810005', 'Seguro de Vida', 1, -4),
                  (-71, current_timestamp,'RENATEA (ex RENATRE)', '810006', 'RENATEA (ex RENATRE)', 1, -4),
                  (-72, current_timestamp,'Prestamos', '810007', 'Prestamos', 1, -4),
                  (-73, current_timestamp,'Impuesto a las Ganancias', '810008', 'Impuesto a las Ganancias', 1, -4),
                  (-74, current_timestamp,'Obra Social - Adherentes', '810009', 'Obra Social - Adherentes', 1, -4),
                  (-75, current_timestamp,'Fondo Solidario de Redistribucion (ex ANSSAL) - Adherentes', '810010', 'Fondo Solidario de Redistribucion (ex ANSSAL) - Adherentes', 1, -4),
                  (-76, current_timestamp,'Otros descuentos', '820000', 'Otros descuentos', 1, -4);
        END IF;

    END
$$;




DROP FUNCTION IF EXISTS ST_eliminar_liquidaciones_sin_legajo();

CREATE OR REPLACE FUNCTION ST_eliminar_liquidaciones_sin_legajo() RETURNS void AS $$
DECLARE
    schema RECORD;
BEGIN
    FOR schema IN SELECT schema_name FROM information_schema.schemata INNER JOIN information_schema.tables ON table_schema = schema_name WHERE left(schema_name, 3) = 'tnt' AND table_name = 'liquidacion'
        LOOP
            EXECUTE format('delete from %I.liquidacion where legajoid is null', schema.schema_name);
        END LOOP;
END;
$$ LANGUAGE plpgsql;