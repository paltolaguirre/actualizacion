SELECT ST_LLENARCONCEPTOAFIP();
UPDATE CONCEPTO SET conceptoafipid = -1 WHERE id IN (-1,-3,-4,-15,-16,-17);
UPDATE CONCEPTO SET conceptoafipid = -11 WHERE id = -2;
UPDATE CONCEPTO SET conceptoafipid = -16 WHERE id = -5;
UPDATE CONCEPTO SET conceptoafipid = -17 WHERE id = -6;
UPDATE CONCEPTO SET conceptoafipid = -56 WHERE id = -7;
UPDATE CONCEPTO SET conceptoafipid = -57 WHERE id = -8;
UPDATE CONCEPTO SET conceptoafipid = -59 WHERE id IN (-9,-14);
UPDATE CONCEPTO SET conceptoafipid = -60 WHERE id = -10;
UPDATE CONCEPTO SET conceptoafipid = -54 WHERE id = -11;
UPDATE CONCEPTO SET conceptoafipid = -52 WHERE id = -12;
UPDATE CONCEPTO SET conceptoafipid = -58 WHERE id = -13;
UPDATE CONCEPTO SET conceptoafipid = -65 WHERE id = -18;
UPDATE CONCEPTO SET conceptoafipid = -66 WHERE id = -19;
UPDATE CONCEPTO SET conceptoafipid = -67 WHERE id = -20;
UPDATE CONCEPTO SET conceptoafipid = -73 WHERE id = -29;
UPDATE CONCEPTO SET conceptoafipid = -69 WHERE id = -31;
UPDATE CONCEPTO SET conceptoafipid = -63 WHERE id = -30;
UPDATE CONCEPTO SET conceptoafipid = -21 WHERE id = -32;
UPDATE CONCEPTO SET conceptoafipid = -6 WHERE id = -34;
UPDATE CONCEPTO SET marcarepeticion = true, aportesipa = true, contribucionsipa = true, aportesinssjyp = true, contribucionesinssjyp = true, aportesobrasocial = true, contribucionesobrasocial = true, aportesfondosolidario = true, contribucionesfondosolidario = true, aportesrenatea = true, contribucionesrenatea = true, asignacionesfamiliares = true, contribucionesfondonacional = true, contribucionesleyriesgo = true, aportesregimenesdiferenciales = false, aportesregimenesespeciales = false WHERE ID IN (-1,-2,-3,-4,-5,-6,-15,-16,-17,-32,-34);
UPDATE CONCEPTO SET marcarepeticion = true, aportesipa = false, contribucionsipa = false, aportesinssjyp = false, contribucionesinssjyp = false, aportesobrasocial = false, contribucionesobrasocial = false, aportesfondosolidario = false, contribucionesfondosolidario = false, aportesrenatea = false, contribucionesrenatea = false, asignacionesfamiliares = false, contribucionesfondonacional = false, contribucionesleyriesgo = true, aportesregimenesdiferenciales = false, aportesregimenesespeciales = false WHERE ID IN (-7,-8,-9,-10,-11,-12,-13,-14);
UPDATE CONCEPTO SET marcarepeticion = true, aportesipa = false, contribucionsipa = false, aportesinssjyp = false, contribucionesinssjyp = false, aportesobrasocial = false, contribucionesobrasocial = false, aportesfondosolidario = false, contribucionesfondosolidario = false, aportesrenatea = false, contribucionesrenatea = false, asignacionesfamiliares = false, contribucionesfondonacional = false, contribucionesleyriesgo = false, aportesregimenesdiferenciales = false, aportesregimenesespeciales = false WHERE ID IN (-18,-19,-20,-29,-30,-31);