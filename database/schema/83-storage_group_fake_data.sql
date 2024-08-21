INSERT INTO storage_group (
        id,
        owner_address,
        group_id,
        group_name,
        source_type,
        extra,
        account_id,
        operator_address,
        expiration_time,
        create_at,
        create_time,
        update_at,
        update_time,
        removed,
        tags
    )
VALUES (
        1,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef',
        'group1',
        'typeA',
        'extra info 1',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2025-08-18 12:00:00+00',
        1620000000,
        '2023-08-18 12:00:00+00',
        1620000000,
        '2023-08-18 12:30:00+00',
        FALSE,
        '{"key": "value1"}'
    ),
    (
        2,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0x2234567890abcdef2234567890abcdef2234567890abcdef2234567890abcdef',
        'group2',
        'typeB',
        'extra info 2',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2025-08-18 13:00:00+00',
        1620000001,
        '2023-08-18 13:00:00+00',
        1620000001,
        '2023-08-18 13:30:00+00',
        FALSE,
        '{"key": "value2"}'
    ),
    (
        3,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0x3234567890abcdef3234567890abcdef3234567890abcdef3234567890abcdef',
        'group3',
        'typeC',
        'extra info 3',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2025-08-18 14:00:00+00',
        1620000002,
        '2023-08-18 14:00:00+00',
        1620000002,
        '2023-08-18 14:30:00+00',
        FALSE,
        '{"key": "value3"}'
    ),
    (
        4,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0x4234567890abcdef4234567890abcdef4234567890abcdef4234567890abcdef',
        'group4',
        'typeD',
        'extra info 4',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2025-08-18 15:00:00+00',
        1620000003,
        '2023-08-18 15:00:00+00',
        1620000003,
        '2023-08-18 15:30:00+00',
        FALSE,
        '{"key": "value4"}'
    ),
    (
        5,
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0x5234567890abcdef5234567890abcdef5234567890abcdef5234567890abcdef',
        'group5',
        'typeE',
        'extra info 5',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2025-08-18 16:00:00+00',
        1620000004,
        '2023-08-18 16:00:00+00',
        1620000004,
        '2023-08-18 16:30:00+00',
        FALSE,
        '{"key": "value5"}'
    );
INSERT INTO storage_group_member (id, group_id, member, expiration_time)
VALUES (
        1,
        '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-08-18 12:00:00+00'
    ),
    (
        2,
        '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-09-18 12:00:00+00'
    ),
    (
        3,
        '0x2234567890abcdef2234567890abcdef2234567890abcdef2234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-08-18 13:00:00+00'
    ),
    (
        4,
        '0x2234567890abcdef2234567890abcdef2234567890abcdef2234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-09-18 13:00:00+00'
    ),
    (
        5,
        '0x3234567890abcdef3234567890abcdef3234567890abcdef3234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-08-18 14:00:00+00'
    ),
    (
        6,
        '0x3234567890abcdef3234567890abcdef3234567890abcdef3234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-09-18 14:00:00+00'
    ),
    (
        7,
        '0x4234567890abcdef4234567890abcdef4234567890abcdef4234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-08-18 15:00:00+00'
    ),
    (
        8,
        '0x4234567890abcdef4234567890abcdef4234567890abcdef4234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-09-18 15:00:00+00'
    ),
    (
        9,
        '0x5234567890abcdef5234567890abcdef5234567890abcdef5234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-08-18 16:00:00+00'
    ),
    (
        10,
        '0x5234567890abcdef5234567890abcdef5234567890abcdef5234567890abcdef',
        '0xabcdefabcdefabcdefabcdefabcdefabcdefabcdef',
        '2024-09-18 16:00:00+00'
    );