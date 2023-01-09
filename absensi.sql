# ************************************************************
# Sequel Ace SQL dump
# Version 20044
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 8.0.31)
# Database: attendance
# Generation Time: 2023-01-09 00:42:39 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table activities
# ------------------------------------------------------------

DROP TABLE IF EXISTS `activities`;

CREATE TABLE `activities` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `activity` longtext,
  `date` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_activities_deleted_at` (`deleted_at`),
  KEY `fk_activities_user` (`user_id`),
  CONSTRAINT `fk_activities_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `activities` WRITE;
/*!40000 ALTER TABLE `activities` DISABLE KEYS */;

INSERT INTO `activities` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `activity`, `date`)
VALUES
	(1,'2023-01-08 13:10:58.276','2023-01-09 05:26:47.827','2023-01-09 05:27:12.045',1,'belajar edit','2022-01-31 00:00:50'),
	(2,'2023-01-08 13:11:11.058','2023-01-08 13:11:11.058',NULL,1,'belajar','2022-01-01 00:00:50'),
	(3,'2023-01-08 13:12:55.285','2023-01-08 13:12:55.285',NULL,1,'belajar','2022-01-31 00:00:50'),
	(4,'2023-01-09 05:24:43.178','2023-01-09 05:24:43.178',NULL,1,'belajar','2023-01-31 00:00:50');

/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table attendances
# ------------------------------------------------------------

DROP TABLE IF EXISTS `attendances`;

CREATE TABLE `attendances` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `type_attendance_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_attendances_deleted_at` (`deleted_at`),
  KEY `fk_attendances_user` (`user_id`),
  KEY `fk_attendances_type_attendance` (`type_attendance_id`),
  CONSTRAINT `fk_attendances_type_attendance` FOREIGN KEY (`type_attendance_id`) REFERENCES `type_attendances` (`id`),
  CONSTRAINT `fk_attendances_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `attendances` WRITE;
/*!40000 ALTER TABLE `attendances` DISABLE KEYS */;

INSERT INTO `attendances` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `type_attendance_id`)
VALUES
	(1,'2023-01-09 05:17:41.654','2023-01-09 05:17:41.654',NULL,1,1);

/*!40000 ALTER TABLE `attendances` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table type_attendances
# ------------------------------------------------------------

DROP TABLE IF EXISTS `type_attendances`;

CREATE TABLE `type_attendances` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `type` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_type_attendances_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `type_attendances` WRITE;
/*!40000 ALTER TABLE `type_attendances` DISABLE KEYS */;

INSERT INTO `type_attendances` (`id`, `created_at`, `updated_at`, `deleted_at`, `type`)
VALUES
	(1,NULL,NULL,NULL,'Check In'),
	(2,NULL,NULL,NULL,'Check Out');

/*!40000 ALTER TABLE `type_attendances` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(256) NOT NULL,
  `email` varchar(256) NOT NULL,
  `password` longtext NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `email`, `password`)
VALUES
	(1,'2023-01-08 08:19:34.324','2023-01-08 08:19:34.324',NULL,'Findryan','findryankurnia@gmail.com','$2a$14$NQu2wLiTDJI2FBDo.gGnwe89fguguREKvpqBn/pUlxMlhBsgdy9a6');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
