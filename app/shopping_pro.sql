/*
 Navicat Premium Data Transfer

 Source Server         : gmx
 Source Server Type    : MySQL
 Source Server Version : 80036 (8.0.36)
 Source Host           : localhost:3306
 Source Schema         : shopping_pro

 Target Server Type    : MySQL
 Target Server Version : 80036 (8.0.36)
 File Encoding         : 65001

 Date: 11/03/2024 01:02:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart`  (
  `ID` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `UserID` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `idx_Cart_DeletedAt`(`DeletedAt` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cart
-- ----------------------------
INSERT INTO `cart` VALUES (1, '2024-03-07 20:02:39.734', '2024-03-07 20:02:39.734', NULL, 18);
INSERT INTO `cart` VALUES (10, '2024-03-10 16:19:02.501', '2024-03-10 16:19:02.501', NULL, 19);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `ID` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `Name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `IsActive` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  UNIQUE INDEX `uni_Category_Name`(`Name` ASC) USING BTREE,
  INDEX `idx_Category_DeletedAt`(`DeletedAt` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 96 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (84, '2024-03-06 22:55:15.787', '2024-03-06 22:55:15.787', NULL, '家電', 'テレビ・電気冷蔵庫・洗濯機・自動炊飯器など、更にはケータイも含め、家庭用電気（電子）器具', 1);
INSERT INTO `category` VALUES (85, '2024-03-06 22:56:05.461', '2024-03-06 22:56:05.461', NULL, 'ドリンク', '液体を口に入れて飲むこと', 1);
INSERT INTO `category` VALUES (86, '2024-03-06 22:56:36.612', '2024-03-06 22:56:36.612', NULL, '車', '心棒を中心に回る仕組みの輪', 1);
INSERT INTO `category` VALUES (87, '2024-03-06 22:57:07.885', '2024-03-06 22:57:07.885', NULL, 'ペット', 'かわいがって飼う動物。', 1);
INSERT INTO `category` VALUES (88, '2024-03-06 22:57:35.702', '2024-03-06 22:57:35.702', NULL, '本', '本', 1);
INSERT INTO `category` VALUES (89, '2024-03-06 22:58:02.974', '2024-03-06 22:58:02.974', NULL, 'ゲーム', '遊び', 1);
INSERT INTO `category` VALUES (90, '2024-03-06 22:58:27.999', '2024-03-06 22:58:27.999', NULL, '日用雑貨', '生活していくために必要な物', 1);
INSERT INTO `category` VALUES (91, '2024-03-06 22:58:56.560', '2024-03-06 22:58:56.560', NULL, '食品', 'すべての飲食物(薬事法(昭和三十五年法律第百四十五 号)に規定する医薬品及び医薬部外品を除く', 1);
INSERT INTO `category` VALUES (92, '2024-03-06 22:59:54.393', '2024-03-06 22:59:54.393', NULL, '健康', '心身健全に関わるもの', 1);
INSERT INTO `category` VALUES (93, '2024-03-06 23:00:21.681', '2024-03-06 23:00:21.681', NULL, ' スポーツ', '運動', 1);
INSERT INTO `category` VALUES (94, '2024-03-06 23:01:11.122', '2024-03-06 23:01:11.122', NULL, 'PC・スマホ・通信', '電子用品', 1);
INSERT INTO `category` VALUES (95, '2024-03-06 23:01:51.764', '2024-03-06 23:01:51.764', NULL, 'インテリア', '室内、つまり家の中をを彩るもの', 1);

-- ----------------------------
-- Table structure for item
-- ----------------------------
DROP TABLE IF EXISTS `item`;
CREATE TABLE `item`  (
  `ID` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `ProductID` bigint UNSIGNED NULL DEFAULT NULL,
  `Count` bigint NULL DEFAULT NULL,
  `CartID` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `idx_Item_DeletedAt`(`DeletedAt` ASC) USING BTREE,
  INDEX `ProductID`(`ProductID` ASC) USING BTREE,
  INDEX `CartID`(`CartID` ASC) USING BTREE,
  CONSTRAINT `CartID` FOREIGN KEY (`CartID`) REFERENCES `cart` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ProductID` FOREIGN KEY (`ProductID`) REFERENCES `product` (`ID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of item
-- ----------------------------

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `ID` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `UserID` bigint UNSIGNED NULL DEFAULT NULL,
  `TotalPrice` double NULL DEFAULT NULL,
  `IsCanceled` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `idx_Order_DeletedAt`(`DeletedAt` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (9, '2024-03-10 16:41:46.967', '2024-03-10 16:41:46.967', NULL, 19, 660000, 0);
INSERT INTO `order` VALUES (10, '2024-03-10 16:43:31.153', '2024-03-10 16:43:31.153', NULL, 19, 404482, 0);
INSERT INTO `order` VALUES (11, '2024-03-10 23:11:39.370', '2024-03-10 23:11:39.370', NULL, 19, 200000, 0);

-- ----------------------------
-- Table structure for ordereditem
-- ----------------------------
DROP TABLE IF EXISTS `ordereditem`;
CREATE TABLE `ordereditem`  (
  `ID` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `ProductID` bigint UNSIGNED NULL DEFAULT NULL,
  `Count` bigint NULL DEFAULT NULL,
  `OrderID` bigint UNSIGNED NULL DEFAULT NULL,
  `IsCanceled` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `idx_OrderedItem_DeletedAt`(`DeletedAt` ASC) USING BTREE,
  INDEX `fk_OrderedItem_Product`(`ProductID` ASC) USING BTREE,
  INDEX `fk_Order_OrderedItems`(`OrderID` ASC) USING BTREE,
  CONSTRAINT `fk_Order_OrderedItems` FOREIGN KEY (`OrderID`) REFERENCES `order` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_OrderedItem_Product` FOREIGN KEY (`ProductID`) REFERENCES `product` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ordereditem
-- ----------------------------
INSERT INTO `ordereditem` VALUES (14, '2024-03-10 16:41:46.971', '2024-03-10 16:41:46.971', NULL, 24, 3, 9, 0);
INSERT INTO `ordereditem` VALUES (15, '2024-03-10 16:41:46.971', '2024-03-10 16:41:46.971', NULL, 28, 3, 9, 0);
INSERT INTO `ordereditem` VALUES (16, '2024-03-10 16:43:31.156', '2024-03-10 16:43:31.156', NULL, 24, 2, 10, 0);
INSERT INTO `ordereditem` VALUES (17, '2024-03-10 16:43:31.156', '2024-03-10 16:43:31.156', NULL, 33, 2, 10, 0);
INSERT INTO `ordereditem` VALUES (18, '2024-03-10 23:11:39.374', '2024-03-10 23:11:39.374', NULL, 27, 2, 11, 0);

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `ID` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `Name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `SKU` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `StockCount` bigint NULL DEFAULT NULL,
  `Price` double NULL DEFAULT NULL,
  `CategoryName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `idx_Product_DeletedAt`(`DeletedAt` ASC) USING BTREE,
  INDEX `fk_Product_Category`(`CategoryName` ASC) USING BTREE,
  CONSTRAINT `fk_Product_Category` FOREIGN KEY (`CategoryName`) REFERENCES `category` (`Name`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (24, '2024-03-06 23:06:04.477', '2024-03-10 16:43:31.154', NULL, 'Microsoft QZI-00020 Surface Laptop 5', 'ec12b6bd-95d1-4520-a282-7470c60be1ff', 'http://localhost:8080/static/Microsoft QZI-00020 Surface Laptop 5.jpg', 'OS：Windows 11 Home CPU ：第12世代 Intel Core i5-1235U Intel Evo プラットフォーム準拠 メインメモリ：8GB グラフィック：Intel Iris Xe グラフィックス ストレージ容量：256GB ディスプレイサイズ：13.5 インチ PixelSense ディスプレイ Dolby Vision IQ 対応*1 ディスプレイ解像度：2256 x 1504 (201 PPI) バッテリー持続時間：最大約18時間*2 約1時間で80%充電 無線LAN ：Wi-Fi 6: 802.11a/b/g/ac/ax Bluetooth ：Bluetooth 5.1', 90, 200000, 'PC・スマホ・通信', 0);
INSERT INTO `product` VALUES (25, '2024-03-07 15:10:00.458', '2024-03-10 16:06:33.119', NULL, 'iphone15ケース', '89da94a0-d9f1-4dc6-81ed-638d319dcdd9', 'http://localhost:8080/static/iphone15ケース.jpg', '手帳型 透明 15Pro 15Plus 15Promax iPhone14 ケース iPhone13 iphone15 iphone 12 SE 11pro XR XS スマホケース カバー クリア シリコン', 87, 1999, 'PC・スマホ・通信', 0);
INSERT INTO `product` VALUES (26, '2024-03-07 15:25:53.607', '2024-03-10 16:18:59.974', NULL, '柴犬', '0fb6f4d9-be57-409e-9f9e-43bc7bf6579b', 'http://localhost:8080/static/柴犬.jpg', '日本原産の日本犬の一種', 19, 200000, 'ペット', 0);
INSERT INTO `product` VALUES (27, '2024-03-07 15:37:02.996', '2024-03-10 23:11:39.372', NULL, '冷蔵庫', '2437ab7a-1cd9-4ca8-b6e6-007989d7ba80', 'http://localhost:8080/static/冷蔵庫.webp', '冷蔵のための箱型の装置', 98, 100000, '家電', 0);
INSERT INTO `product` VALUES (28, '2024-03-07 15:43:25.668', '2024-03-10 16:41:46.970', NULL, '柴犬test', '04c64fac-3720-449e-8aaf-fe98bf2a301d', 'http://localhost:8080/static/柴犬test.jpg', '柴犬test柴犬test柴犬test', 997, 20000, 'ペット', 0);
INSERT INTO `product` VALUES (29, '2024-03-07 15:47:11.564', '2024-03-07 15:47:38.787', NULL, '冷蔵庫test0', 'fd262814-0d67-4876-a6fd-edaaeceabede', 'http://localhost:8080/static/冷蔵庫test0.webp', '冷蔵庫test0冷蔵庫test0冷蔵庫test0', 2353, 324132, '家電', 0);
INSERT INTO `product` VALUES (30, '2024-03-07 15:52:05.160', '2024-03-10 14:55:24.087', NULL, 'iphone15ケースwafwaf', '051b5bd4-5f22-45df-a285-e96a82995745', 'http://localhost:8080/static/iphone15ケースwafwaf.jpg', 'iphone15ケースwafwafwqfwq', 2137, 2132, 'PC・スマホ・通信', 0);
INSERT INTO `product` VALUES (31, '2024-03-07 15:54:03.569', '2024-03-10 16:13:08.846', NULL, '辛いもの', '7a47cdfc-5f90-4599-b0ab-dc774f0455eb', 'http://localhost:8080/static/辛いもの.jpg', '辛い辛い辛い辛い', 19995, 2000, '食品', 0);
INSERT INTO `product` VALUES (32, '2024-03-07 15:55:53.220', '2024-03-10 15:24:52.557', NULL, '辛いもの２', '49ced63b-b090-48ec-9e3d-152bc5b95d66', 'http://localhost:8080/static/辛いもの２.jpg', '辛いもの２辛いもの２', 235230, 2130, '食品', 0);
INSERT INTO `product` VALUES (33, '2024-03-07 15:58:01.431', '2024-03-10 16:43:31.155', NULL, 'ぱいぱい茶ぱいぱい', 'b154a6bf-805c-45a6-b967-4cc9d5d2e357', 'http://localhost:8080/static/ぱいぱい茶ぱいぱい.jpg', 'ぱいぱい茶ぱいぱいぱいぱい茶ぱいぱいぱいぱい茶ぱいぱい', 14237, 2241, '食品', 0);
INSERT INTO `product` VALUES (34, '2024-03-07 15:58:49.712', '2024-03-07 15:58:58.753', NULL, '茶', '4dfd9f01-0e42-47e1-94dd-d227c4793ef4', 'http://localhost:8080/static/茶.jpg', 'ぱいぱい茶ぱいぱいぱいぱい茶ぱいぱいぱいぱい茶ぱいぱい', 142443, 2000, '食品', 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `ID` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `Username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Salt` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Token` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Gender` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Birth` datetime NULL DEFAULT NULL,
  `Url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Address` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `IsDeleted` tinyint(1) NULL DEFAULT NULL,
  `IsAdmin` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `idx_User_DeletedAt`(`DeletedAt` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (18, '2024-03-05 17:22:50.318', '2024-03-10 13:20:47.268', NULL, 'gmx', '$2a$10$q6HKeuaOZvqxVKy/qwTgI.FL7J6cV1nq4CKy/JzBIMbZ6xRUFK0Jm', 'bvlf5OvJFHiQcZg7VReyO5KDaOPXcZo', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAxMzA4NDcsImlhdCI6MTcxMDA0NDQ0NywiaXNBZG1pbiI6dHJ1ZSwiaXNzIjoiIiwidXNlcklkIjoiMTgiLCJ1c2VybmFtZSI6ImdteCJ9.IpgMAgoGWY4eZJk-qH6UcJqZf3HKm8FVmT116P13_AY', '', '2000-03-30 00:00:00', '', '', '', 0, 1);
INSERT INTO `user` VALUES (19, '2024-03-07 20:57:18.191', '2024-03-10 23:11:31.552', NULL, 'hhh', '$2a$10$N55dDfplAodLHB7NNGlpiueX6zYlCKYSMHoL0BDtHfXuvEaSamlBK', '9QQ893OZagHrHMk9dxALudilfTdOOHR', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAxNjYyOTEsImlhdCI6MTcxMDA3OTg5MSwiaXNBZG1pbiI6ZmFsc2UsImlzcyI6IiIsInVzZXJJZCI6IjE5IiwidXNlcm5hbWUiOiJoaGgifQ.Y0XZzRTaUInG3D9oaAv7FQjcR335QI1VmDi23T1KwlQ', '', '2000-03-17 00:00:00', '', '', '', 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
