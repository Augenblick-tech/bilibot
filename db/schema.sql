PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `users` (`id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` varchar(255) NOT NULL,`password` varchar(255) NOT NULL,PRIMARY KEY (`id`));
CREATE TABLE `bots` (`uid` integer NOT NULL,`name` varchar(255) NOT NULL,`face` text,`cookie` text NOT NULL,`user_id` integer NOT NULL,PRIMARY KEY (`uid`),CONSTRAINT `fk_bots_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE SET NULL ON UPDATE CASCADE);
CREATE TABLE `authors` (`uid` integer NOT NULL,`name` varchar(255) NOT NULL,`face` text,`bot_id` integer NOT NULL,PRIMARY KEY (`uid`),CONSTRAINT `fk_authors_bot` FOREIGN KEY (`bot_id`) REFERENCES `bots`(`uid`) ON DELETE SET NULL ON UPDATE CASCADE);
CREATE TABLE `dynamics` (`dynamic_id` text NOT NULL,`pub_ts` integer NOT NULL,`content` text NOT NULL,`author_id` integer NOT NULL,PRIMARY KEY (`dynamic_id`),CONSTRAINT `fk_dynamics_author` FOREIGN KEY (`author_id`) REFERENCES `authors`(`uid`) ON DELETE SET NULL ON UPDATE CASCADE);
CREATE INDEX `idx_users_deleted_at` ON `users`(`deleted_at`);
COMMIT;
