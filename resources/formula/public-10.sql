UPDATE deduccionespersonales SET valorfijomni = 85848.99 WHERE anio = 2019;
UPDATE deduccionespersonales SET valorfijoddei = 412075.12 WHERE anio = 2019;
UPDATE deduccionespersonales SET valorfijoconyuge = 85848.99 WHERE anio = 2019;
UPDATE deduccionespersonales SET valorfijohijo = 40361.43 WHERE anio = 2019;

ALTER TABLE deduccionespersonales ADD COLUMN ley27541gni NUMERIC(19,4) NOT NULL DEFAULT 0;
ALTER TABLE deduccionespersonales ADD COLUMN ley27541de NUMERIC(19,4) NOT NULL DEFAULT 0;

UPDATE deduccionespersonales SET ley27541gni = 17169.80 WHERE anio = 2019;
UPDATE deduccionespersonales SET ley27541de = 82415.03 WHERE anio = 2019;
