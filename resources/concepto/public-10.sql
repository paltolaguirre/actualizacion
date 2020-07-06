update concepto set tipocalculoautomaticoid = -1 where tipodecalculoid is null;
update concepto set tipocalculoautomaticoid = -2 where tipodecalculoid is not null;
update concepto set formulanombre = 'ImpuestoALasGanancias', tipocalculoautomaticoid = -3 where id = -29;
update concepto set formulanombre = 'ImpuestoALasGananciasDevolucion', tipocalculoautomaticoid = -3 where id = -30;
update concepto set eseditable = false where tipocalculoautomaticoid != -1;
update concepto set eseditable = true where tipocalculoautomaticoid = -1;