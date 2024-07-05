CREATE TABLE users
(
    id serial not null unique,
    name VARCHAR(255) not NULL,
    pasportNumber varchar(255) not null
);

CREATE TABLE timeTrackerItems
(
    id serial not null unique,
    name varchar(255) not null,
    timeStart TIMESTAMP not null,
    timeStop TIMESTAMP not null,
    timeUsed VARCHAR(255) not null,
    userId int references users (id) on delete cascade not null
);

INSERT INTO users (name, pasportnumber) VALUES
('Иван', '1234 123456'),
('Петр', '2345 234567');

INSERT INTO timeTrackerItems (name, userid, timestart, timestop, timeused) VALUES
('Завтрак', 1, '2024-07-04T07:30:00+03:00', '2024-07-04T09:00:00+03:00', 
    TRIM(BOTH ' ' FROM 
        CONCAT(
            CASE 
                WHEN EXTRACT(HOUR FROM justify_hours('2024-07-04T09:00:00+03:00'::timestamp - '2024-07-04T07:30:00+03:00'::timestamp)) > 0 
                THEN EXTRACT(HOUR FROM justify_hours('2024-07-04T09:00:00+03:00'::timestamp - '2024-07-04T07:30:00+03:00'::timestamp)) || ' часов ' 
                ELSE '' 
            END,
            CASE 
                WHEN EXTRACT(MINUTE FROM justify_hours('2024-07-04T09:00:00+03:00'::timestamp - '2024-07-04T07:30:00+03:00'::timestamp)) > 0 
                THEN EXTRACT(MINUTE FROM justify_hours('2024-07-04T09:00:00+03:00'::timestamp - '2024-07-04T07:30:00+03:00'::timestamp)) || ' минут' 
                ELSE '' 
            END
        )
    )),
('Обед', 1, '2024-07-04T13:00:00+03:00', '2024-07-04T15:40:00+03:00', 
    TRIM(BOTH ' ' FROM 
        CONCAT(
            CASE 
                WHEN EXTRACT(HOUR FROM justify_hours('2024-07-04T15:40:00+03:00'::timestamp - '2024-07-04T13:00:00+03:00'::timestamp)) > 0 
                THEN EXTRACT(HOUR FROM justify_hours('2024-07-04T15:40:00+03:00'::timestamp - '2024-07-04T13:00:00+03:00'::timestamp)) || ' часов ' 
                ELSE '' 
            END,
            CASE 
                WHEN EXTRACT(MINUTE FROM justify_hours('2024-07-04T15:40:00+03:00'::timestamp - '2024-07-04T13:00:00+03:00'::timestamp)) > 0 
                THEN EXTRACT(MINUTE FROM justify_hours('2024-07-04T15:40:00+03:00'::timestamp - '2024-07-04T13:00:00+03:00'::timestamp)) || ' минут' 
                ELSE '' 
            END
        )
    )),
('Ужин', 1, '2024-07-04T18:00:00+03:00', '2024-07-04T21:50:00+03:00', 
    TRIM(BOTH ' ' FROM 
        CONCAT(
            CASE 
                WHEN EXTRACT(HOUR FROM justify_hours('2024-07-04T21:50:00+03:00'::timestamp - '2024-07-04T18:00:00+03:00'::timestamp)) > IS NOT NULL
                THEN EXTRACT(HOUR FROM justify_hours('2024-07-04T21:50:00+03:00'::timestamp - '2024-07-04T18:00:00+03:00'::timestamp)) || ' часов ' 
                ELSE '' 
            END,
            CASE 
                WHEN EXTRACT(MINUTE FROM justify_hours('2024-07-04T21:50:00+03:00'::timestamp - '2024-07-04T18:00:00+03:00'::timestamp)) > IS NOT NULL
                THEN EXTRACT(MINUTE FROM justify_hours('2024-07-04T21:50:00+03:00'::timestamp - '2024-07-04T18:00:00+03:00'::timestamp)) || ' минут' 
                ELSE '' 
            END
        )
    )),
('Обед', 2, '2024-07-04T13:00:00+03:00', '2024-07-04T15:40:00+03:00', 
    TRIM(BOTH ' ' FROM 
        CONCAT(
            CASE 
                WHEN EXTRACT(HOUR FROM justify_hours('2024-07-04T15:40:00+03:00'::timestamp - '2024-07-04T13:00:00+03:00'::timestamp)) > IS NOT NULL 
                THEN EXTRACT(HOUR FROM justify_hours('2024-07-04T15:40:00+03:00'::timestamp - '2024-07-04T13:00:00+03:00'::timestamp)) || ' часов ' 
                ELSE '' 
            END,
            CASE 
                WHEN EXTRACT(MINUTE FROM justify_hours('2024-07-04T15:40:00+03:00'::timestamp - '2024-07-04T13:00:00+03:00'::timestamp)) > IS NOT NULL 
                THEN EXTRACT(MINUTE FROM justify_hours('2024-07-04T15:40:00+03:00'::timestamp - '2024-07-04T13:00:00+03:00'::timestamp)) || ' минут' 
                ELSE '' 
            END
        )
    ));
