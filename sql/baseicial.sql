DROP TABLE IF EXISTS `eventos`;
CREATE TABLE `eventos` (
    `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `titulo` varchar(255) NOT NULL,
    `descricao` text NOT NULL,
    `inicio` datetime NOT NULL,
    `fim` datetime NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

GRANT ALL PRIVILEGES ON eventos.* TO 'eventos'@'%' WITH GRANT OPTION;

FLUSH PRIVILEGES;