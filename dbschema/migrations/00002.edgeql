CREATE MIGRATION m1ajja2pdgzlgrtmfip7pd62u6lkw2hjd52mkeaofzwspjhada5yiq
    ONTO m13nabibq23bzg35acwxj5vc22thafyniyoahtgq2vlzie6oafpl7a
{
  ALTER TYPE default::Person {
      ALTER PROPERTY last_name {
          RESET OPTIONALITY;
      };
  };
};
