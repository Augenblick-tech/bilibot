CREATE TABLE `users` (
  `id` INTEGER NOT NULL PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL UNIQUE,
  `password` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME
);

CREATE TABLE `bots` (
  `uid` INTEGER NOT NULL PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `face` TEXT,
  `user_id` INTEGER NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES users (`id`) ON DELETE CASCADE
);

CREATE TABLE `uploaders` (
  `uid` INTEGER NOT NULL PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `face` TEXT,
  `bot_id` INTEGER NOT NULL,
  FOREIGN KEY (`bot_id`) REFERENCES bots (`uid`) ON DELETE CASCADE
);

CREATE TABLE `dynamics` (
  `dynamic_id` TEXT NOT NULL PRIMARY KEY,
  `pub_ts` INTEGER NOT NULL,
  `content` TEXT NOT NULL,
  `up_id` INTEGER NOT NULL,
  FOREIGN KEY (`up_id`) REFERENCES uploaders (`uid`) ON DELETE CASCADE
);