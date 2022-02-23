CREATE DATABASE IF NOT EXISTS madre;

-- utf8mb4_0900_ai_ci;
-- utf8mb4_unicode_ci
ALTER DATABASE madre CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- user
CREATE TABLE IF NOT EXISTS madre.user (
  `id` varchar(36) COLLATE utf8mb4_general_ci NOT NULL DEFAULT ( uuid() ),
  `email` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `username` varchar(16) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `display_name` varchar(48) COLLATE utf8mb4_general_ci NOT NULL,
  `photo_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- data
CREATE TABLE IF NOT EXISTS madre.data (
  `id` varchar(36) COLLATE utf8mb4_general_ci NOT NULL DEFAULT ( uuid() ),
  `user_id` varchar(36) COLLATE utf8mb4_general_ci NOT NULL,
  `file_url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_public` tinyint NOT NULL DEFAULT 0,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`),
  KEY `ix_created_at` (`created_at`),
  KEY `ix_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;