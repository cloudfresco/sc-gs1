INSERT INTO `inventory_reports`(
  uuid4,
  inventory_report_type_code,
  structure_type_code,
  inventory_report_identification,
  inventory_reporting_party,
  inventory_report_to_party,
  reporting_period_begin,
  reporting_period_end,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
  UNHEX(REPLACE('2037d2fa-3d30-43f5-b399-b85b24e8df29','-','')),
  'INVENTORY_STATUS',
  'LOCATION_BY_ITEM',
  0,
  8,
  9,
  '2005-02-09 08:00:00',
  '2005-03-13 08:00:00',
  'active',
  'auth0|673ee1a719dd4000cd5a3832',
  'auth0|673ee1a719dd4000cd5a3832',
  '2005-02-09 08:00:00',
  '2012-07-11 08:00:00'
);

INSERT INTO `inventory_reports`(
  uuid4,
  inventory_report_type_code,
  structure_type_code,
  inventory_report_identification,
  inventory_reporting_party,
  inventory_report_to_party,
  reporting_period_begin,
  reporting_period_end,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
  UNHEX(REPLACE('4a7a05ed-f632-457a-b51c-7631a99ac079','-','')),
  'INVENTORY_ACTIVITY',
  'ITEM_BY_LOCATION',
  0,
  8,
  9,
  '2005-02-07 08:00:00',
  '2005-02-08 08:00:00',
  'active',
  'auth0|673ee1a719dd4000cd5a3832',
  'auth0|673ee1a719dd4000cd5a3832',
  '2005-02-09 08:00:00',
  '2012-07-11 08:00:00'
);
