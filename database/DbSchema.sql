CREATE DATABASE IF NOT EXISTS restapi;

USE restapi;

CREATE TABLE IF NOT EXISTS `products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` int NOT NULL,
  `vat` varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `user` (
  `user_id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255)
);

CREATE TABLE IF NOT EXISTS `shopping_carts` (
  `product_id` int,
  `user_id` int,
  `quantity` int DEFAULT 1
);

CREATE TABLE IF NOT EXISTS `placed_orders` (
  `user_id` int,
  `total_price` int,
  `created_at` timestamp
);

CREATE TABLE IF NOT EXISTS `given_amounts` (
  `given_amount` int
);

ALTER TABLE `shopping_carts` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `shopping_carts` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);

ALTER TABLE `placed_orders` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);

Insert Into products ( name,price,vat)
values ("airpods", 2000, 1),("watch",4000,8),("iphone",8000,18),("macbook",14000,18);

Insert Into user (name)
values ("user1"),("user2");


/* Optinal
SET SQL_SAFE_UPDATES = 0;

SET FOREIGN_KEY_CHECKS = 0;
*/
