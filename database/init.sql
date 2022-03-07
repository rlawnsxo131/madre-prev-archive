CREATE DATABASE IF NOT EXISTS madre;

-- utf8mb4_0900_ai_ci;
-- utf8mb4_unicode_ci

-- ALTER DATABASE madre CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- user
CREATE TABLE IF NOT EXISTS madre.user (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `username` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `display_name` varchar(48) COLLATE utf8mb4_general_ci NOT NULL,
  `photo_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_uuid` (`uuid`),
  UNIQUE KEY `ix_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- auth
CREATE TABLE IF NOT EXISTS madre.social_account (
  `id` int(10) unsigned COLLATE utf8mb4_general_ci NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` int(10) unsigned COLLATE utf8mb4_general_ci NOT NULL,
  `provider` enum("GOOGLE") COLLATE utf8mb4_general_ci NOT NULL DEFAULT "GOOGLE",
  `social_id` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_uuid` (`uuid`),
  UNIQUE KEY `ix_provider_social_id` (`provider`, `social_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- data
CREATE TABLE IF NOT EXISTS madre.data (
  `id` int(10) unsigned COLLATE utf8mb4_general_ci NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` int(10) unsigned COLLATE utf8mb4_general_ci NOT NULL,
  `file_url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_public` tinyint NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_uuid` (`uuid`),
  KEY `ix_created_at` (`created_at`),
  KEY `ix_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;