  INSERT INTO `order_response_line_items`(
    uuid4,
    line_item_number,
    confirmed_quantity,
    line_item_change_indicator,
    original_order_line_item_number,
    order_response_reason_code,
    order_response_id,
    delivery_date_time,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
  ) VALUES (
    UNHEX(REPLACE('34390ee7-ca05-47dc-9740-3929511dbbef','-','')),
    1, 
    48, 
    'ACCEPTED',
    1,
    '',
    3,
    '2019-03-11 10:04:26',
    'active',
    'auth0|673ee1a719dd4000cd5a3832',
    'auth0|673ee1a719dd4000cd5a3832',
    '2019-03-11 10:04:26',
    '2019-03-11 10:04:26'
  );

  INSERT INTO `order_response_line_items`(
    uuid4,
    line_item_number,
    confirmed_quantity,
    line_item_change_indicator,
    original_order_line_item_number,
    order_response_reason_code,
    order_response_id,
    delivery_date_time,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
  ) VALUES (
    UNHEX(REPLACE('372e1778-e853-4745-9470-169d00935315','-','')),
    2, 
    24, 
    'MODIFIED',
    2,
    'PRODUCT_OUT_OF_STOCK',
    3,
    '2019-03-11 10:04:26',
    'active',
    'auth0|673ee1a719dd4000cd5a3832',
    'auth0|673ee1a719dd4000cd5a3832',
    '2019-03-11 10:04:26',
    '2019-03-11 10:04:26'
  );
