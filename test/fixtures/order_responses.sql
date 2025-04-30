  INSERT INTO `order_responses`(
    uuid4,
    order_response_reason_code,
    response_status_code,
    order_response_identification,
    seller,
    buyer,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
  ) VALUES (
    UNHEX(REPLACE('51340d6c-f416-4026-90e5-e1298a514808','-','')),
    '', 
    'ACCEPTED',
    3, 
    1, 
    2, 
    'active',
    '2019-03-11 10:04:26',
    '2019-03-11 10:04:26',
    'auth0|673ee1a719dd4000cd5a3832',
    'auth0|673ee1a719dd4000cd5a3832'
  );

  INSERT INTO `order_responses`(
    uuid4,
    order_response_reason_code,
    response_status_code,
    order_response_identification,
    seller,
    buyer,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
  ) VALUES (
    UNHEX(REPLACE('0542056b-7fb4-4a79-9e0e-e65bdb07a2a2','-','')),
    'DISCONTINUED_LINE', 
    'REJECTED',
    4, 
    1, 
    2, 
    'active',
    '2019-03-11 10:04:26',
    '2019-03-11 10:04:26',
    'auth0|673ee1a719dd4000cd5a3832',
    'auth0|673ee1a719dd4000cd5a3832'
  );

  INSERT INTO `order_responses`(
    uuid4,
    order_response_reason_code,
    response_status_code,
    order_response_identification,
    seller,
    buyer,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
  ) VALUES (
    UNHEX(REPLACE('a9cf4bc4-bf4b-4013-95d3-d3bb4c7af952','-','')),
    '', 
    'MODIFIED',
    5, 
    1, 
    2, 
    'active',
    '2019-03-11 10:04:26',
    '2019-03-11 10:04:26',
    'auth0|673ee1a719dd4000cd5a3832',
    'auth0|673ee1a719dd4000cd5a3832'
  );
