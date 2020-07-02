update concepto set esremvariable = false where esremvariable is NULL;
update concepto set esremvariable = true where id in (-5, -6);