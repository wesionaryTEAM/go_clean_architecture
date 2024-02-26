
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint unsigned AUTO_INCREMENT PRIMARY KEY,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NULL,
  `uuid` BINARY(16) NOT NULL,
  `cognito_uid` VARCHAR(50) NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `first_name_ja` VARCHAR(255) NOT NULL,
  `last_name_ja` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL UNIQUE,
  `role` VARCHAR(25) NOT NULL,
  `is_active` BOOLEAN NOT NULL DEFAULT false,
  `is_email_verified` BOOLEAN NOT NULL DEFAULT false,
  `is_admin` BOOLEAN NOT NULL DEFAULT false,
  `profile_pic` text NOT NULL,
  INDEX `idx_users_uuid` (`uuid`),
  INDEX `idx_users_cognito_uid` (`cognito_uid`),
  INDEX `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- +migrate Down
DROP TABLE IF EXISTS `users`;