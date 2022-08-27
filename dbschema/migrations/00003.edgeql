CREATE MIGRATION m1j2tcgnrsl5satp7n3sfg3jsqekkcmhvtuswwkjajtu334fbzus4a
    ONTO m1ajja2pdgzlgrtmfip7pd62u6lkw2hjd52mkeaofzwspjhada5yiq
{
  ALTER TYPE default::Person {
      CREATE PROPERTY full_name := ((((.first_name ++ ' ') ++ .last_name) IF EXISTS (.last_name) ELSE .first_name));
  };
};
