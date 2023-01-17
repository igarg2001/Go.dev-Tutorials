DROP TABLE IF EXISTS album;
CREATE TABLE album(
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(128) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    PRIMARY KEY(`id`)
);

insert into ALBUM (title, artist, price)
values (
    'Blue Train', 'John Coltrane', 56.99
),
(
    'Jeru', 'Gerry Mulligan', 17.99
),
(
    'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99
),
(
    'Giant Steps', 'John Coltrane', 63.99
);

