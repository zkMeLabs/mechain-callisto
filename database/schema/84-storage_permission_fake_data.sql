INSERT INTO permission (
        id,
        principal_type,
        principal_value,
        resource_type,
        resource_id,
        policy_id,
        create_time,
        update_time,
        expiration_time,
        removed
    )
VALUES (
        1,
        1,
        'principal_value_1',
        'resource_type_1',
        1,
        1,
        '2022-01-20 12:45:19',
        '2022-10-06 12:47:57',
        '2025-10-15 12:40:57',
        false
    ),
    (
        2,
        1,
        'principal_value_2',
        'resource_type_2',
        2,
        2,
        '2021-04-23 12:41:21',
        '2020-11-08 12:31:23',
        '2024-09-13 12:44:07',
        false
    ),
    (
        3,
        1,
        'principal_value_3',
        'resource_type_3',
        3,
        3,
        '2020-12-01 12:47:08',
        '2021-05-10 12:42:44',
        '2025-10-31 12:53:24',
        true
    ),
    (
        4,
        3,
        'principal_value_4',
        'resource_type_4',
        4,
        4,
        '2021-05-11 12:33:35',
        '2022-09-12 12:55:23',
        '2025-09-30 12:45:29',
        false
    ),
    (
        5,
        2,
        'principal_value_5',
        'resource_type_5',
        5,
        5,
        '2023-02-26 12:28:52',
        '2020-09-17 12:28:11',
        '2026-08-14 12:34:52',
        true
    );