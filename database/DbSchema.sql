CREATE TABLE `products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` int NOT NULL,
  `vat` varchar(255) NOT NULL
);

CREATE TABLE `user` (
  `user_id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255)
);

CREATE TABLE `shopping_cart` (
  `cart_product_id` int,
  `user_id` int,
  `quantity` int DEFAULT 1
);

CREATE TABLE `placed_orders` (
  `user_id` int,
  `total_price` int,
  `created_at` timestamp
);

CREATE TABLE `given_amounts` (
  `given_amount` int
);

ALTER TABLE `products` ADD FOREIGN KEY (`id`) REFERENCES `shopping_cart` (`cart_product_id`);

ALTER TABLE `shopping_cart` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);

ALTER TABLE `placed_orders` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);
