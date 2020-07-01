UPDATE CONCEPTO SET prorrateo = false, basesac = true, tipoimpuestogananciasid = -1 WHERE ID IN (-1,-3,-4,-15,-16,-17);
UPDATE CONCEPTO SET prorrateo = false, basesac = true, tipoimpuestogananciasid = -3 WHERE ID IN (-5,-6);
UPDATE CONCEPTO SET prorrateo = false, basesac = false, tipoimpuestogananciasid = -2 WHERE ID IN (-8,-9,-10,-11,-13,-14);
UPDATE CONCEPTO SET prorrateo = true, basesac = true, tipoimpuestogananciasid = -2 WHERE ID = -12;
UPDATE CONCEPTO SET prorrateo = false, tipoimpuestogananciasid = -10 WHERE ID IN (-18,-19);
UPDATE CONCEPTO SET prorrateo = false, tipoimpuestogananciasid = -11 WHERE ID = -20;
