INSERT INTO bucket (
        id,
        bucket_id,
        bucket_name,
        owner_address,
        payment_address,
        global_virtual_group_family_id,
        operator_address,
        source_type,
        charged_read_quota,
        visibility,
        status,
        delete_at,
        delete_reason,
        storage_size,
        charge_size,
        create_at,
        create_tx_hash,
        create_time,
        update_at,
        update_tx_hash,
        update_time,
        removed,
        off_chain_status,
        tags
    )
VALUES (
        1,
        1,
        'bucket1',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        4,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeA',
        10000,
        'public',
        'active',
        NULL,
        NULL,
        1000000,
        1000000,
        1620000000,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 12:00:00+00',
        1620000000,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 12:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        2,
        2,
        'bucket2',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        1,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeB',
        20000,
        'private',
        'inactive',
        NULL,
        NULL,
        2000000,
        2000000,
        1620000001,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 13:00:00+00',
        1620000001,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 13:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        3,
        3,
        'bucket3',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        5,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeC',
        30000,
        'public',
        'active',
        NULL,
        NULL,
        3000000,
        3000000,
        1620000002,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 14:00:00+00',
        1620000002,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 14:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        4,
        4,
        'bucket4',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        1,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeD',
        40000,
        'private',
        'inactive',
        NULL,
        NULL,
        4000000,
        4000000,
        1620000003,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 15:00:00+00',
        1620000003,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 15:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        5,
        5,
        'bucket5',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        2,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeE',
        50000,
        'public',
        'active',
        NULL,
        NULL,
        5000000,
        5000000,
        1620000004,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 16:00:00+00',
        1620000004,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 16:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        6,
        6,
        'bucket6',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        3,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeF',
        60000,
        'private',
        'inactive',
        NULL,
        NULL,
        6000000,
        6000000,
        1620000005,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 17:00:00+00',
        1620000005,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 17:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        7,
        7,
        'bucket7',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        5,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeG',
        70000,
        'public',
        'active',
        NULL,
        NULL,
        7000000,
        7000000,
        1620000006,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 18:00:00+00',
        1620000006,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 18:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        8,
        8,
        'bucket8',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        4,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeH',
        80000,
        'private',
        'inactive',
        NULL,
        NULL,
        8000000,
        8000000,
        1620000007,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 19:00:00+00',
        1620000007,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 19:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        9,
        9,
        'bucket9',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        2,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeI',
        90000,
        'public',
        'active',
        NULL,
        NULL,
        9000000,
        9000000,
        1620000008,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 20:00:00+00',
        1620000008,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 20:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    ),
    (
        10,
        10,
        'bucket10',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdef',
        1,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        'typeJ',
        100000,
        'private',
        'inactive',
        NULL,
        NULL,
        10000000,
        10000000,
        1620000009,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-17 21:00:00+00',
        1620000009,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2023-08-18 21:00:00+00',
        FALSE,
        0,
        '{"key": "value"}'
    );