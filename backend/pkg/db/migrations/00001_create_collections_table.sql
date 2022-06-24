-- +goose Up
CREATE TABLE `collections` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `path_prefix` varchar(255) NOT NULL,
  `content_type` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uri` (`path_prefix`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `collections`;