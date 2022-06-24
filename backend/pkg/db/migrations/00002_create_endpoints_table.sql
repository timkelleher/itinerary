-- +goose Up
CREATE TABLE `endpoints` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `collection_id` int(11) unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `method` varchar(10) NOT NULL,
  `response_status_code` smallint(6) NOT NULL,
  `path` varchar(255) NOT NULL,
  `computed_path` varchar(255) NOT NULL,
  `output` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `collection_id` (`collection_id`),
  CONSTRAINT `endpoints_ibfk_3` FOREIGN KEY (`collection_id`) REFERENCES `collections` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS `endpoints`;