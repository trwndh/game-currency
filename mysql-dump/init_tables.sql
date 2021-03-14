DROP TABLE IF EXISTS `currency`;

CREATE TABLE currency (
id int(5) NOT NULL AUTO_INCREMENT,
name varchar(25) NOT NULL,
created_at timestamp NULL DEFAULT current_timestamp(),
updated_at timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
PRIMARY KEY (id),
UNIQUE KEY name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `conversion_rate`;

CREATE TABLE conversion_rate (
id int(5) NOT NULL AUTO_INCREMENT,
currency_id_from int(5) NOT NULL,
currency_id_to int(5) NOT NULL,
rate int(5) NOT NULL,
created_at timestamp NOT NULL DEFAULT current_timestamp(),
updated_at timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
PRIMARY KEY (id),
KEY id_currency_from (currency_id_from),
KEY id_currency_to (currency_id_to),
CONSTRAINT conversion_rate_ibfk_1 FOREIGN KEY (currency_id_from) REFERENCES currency (id),
CONSTRAINT conversion_rate_ibfk_2 FOREIGN KEY (currency_id_to) REFERENCES currency (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;