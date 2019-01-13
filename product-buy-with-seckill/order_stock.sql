/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : localhost:3306
 Source Schema         : order_stock

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 16/10/2018 20:11:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_detail
-- ----------------------------
DROP TABLE IF EXISTS `order_detail`;
CREATE TABLE `order_detail` (
  `order_detail_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) NOT NULL,
  `product_id` bigint(20) DEFAULT NULL,
  `product_buy_number` int(11) DEFAULT NULL,
  PRIMARY KEY (`order_detail_id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of order_detail
-- ----------------------------
BEGIN;
INSERT INTO `order_detail` VALUES (24, 28, 54, 0);
INSERT INTO `order_detail` VALUES (25, 29, 30, 4);
INSERT INTO `order_detail` VALUES (26, 30, 28, 7);
INSERT INTO `order_detail` VALUES (27, 31, 64, 5);
COMMIT;

-- ----------------------------
-- Table structure for order_master
-- ----------------------------
DROP TABLE IF EXISTS `order_master`;
CREATE TABLE `order_master` (
  `order_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `order_payment` decimal(10,0) DEFAULT NULL,
  `order_status` int(2) DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of order_master
-- ----------------------------
BEGIN;
INSERT INTO `order_master` VALUES (28, 4632647315599027822, 0, 0);
INSERT INTO `order_master` VALUES (29, 4996006649777336045, 3, 0);
INSERT INTO `order_master` VALUES (30, 2426836667417963642, 5, 0);
COMMIT;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `product_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `product_price` decimal(7,5) DEFAULT NULL,
  `product_stock` int(7) DEFAULT NULL,
  `product_level` int(1) DEFAULT '0' COMMENT '0 普通商品 1秒杀商品',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` VALUES (1, 'product_0', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (2, 'product_1', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (3, 'product_2', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (4, 'product_3', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (5, 'product_4', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (6, 'product_5', 0.72712, 992, 0);
INSERT INTO `product` VALUES (7, 'product_6', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (8, 'product_7', 0.72712, 988, 0);
INSERT INTO `product` VALUES (9, 'product_8', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (10, 'product_9', 0.72712, 992, 0);
INSERT INTO `product` VALUES (11, 'product_10', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (12, 'product_11', 0.72712, 991, 0);
INSERT INTO `product` VALUES (13, 'product_12', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (14, 'product_13', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (15, 'product_14', 0.72712, 996, 0);
INSERT INTO `product` VALUES (16, 'product_15', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (17, 'product_16', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (18, 'product_17', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (19, 'product_18', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (20, 'product_19', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (21, 'product_20', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (22, 'product_21', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (23, 'product_22', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (24, 'product_23', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (25, 'product_24', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (26, 'product_25', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (27, 'product_26', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (28, 'product_27', 0.72712, 984, 0);
INSERT INTO `product` VALUES (29, 'product_28', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (30, 'product_29', 0.72712, 996, 0);
INSERT INTO `product` VALUES (31, 'product_30', 0.72712, 992, 0);
INSERT INTO `product` VALUES (32, 'product_31', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (33, 'product_32', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (34, 'product_33', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (35, 'product_34', 0.72712, 992, 0);
INSERT INTO `product` VALUES (36, 'product_35', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (37, 'product_36', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (38, 'product_37', 0.72712, 992, 0);
INSERT INTO `product` VALUES (39, 'product_38', 0.72712, 995, 0);
INSERT INTO `product` VALUES (40, 'product_39', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (41, 'product_40', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (42, 'product_41', 0.72712, 997, 0);
INSERT INTO `product` VALUES (43, 'product_42', 0.72712, 994, 0);
INSERT INTO `product` VALUES (44, 'product_43', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (45, 'product_44', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (46, 'product_45', 0.72712, 998, 0);
INSERT INTO `product` VALUES (47, 'product_46', 0.72712, 995, 0);
INSERT INTO `product` VALUES (48, 'product_47', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (49, 'product_48', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (50, 'product_49', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (51, 'product_50', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (52, 'product_51', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (53, 'product_52', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (54, 'product_53', 0.72712, 998, 0);
INSERT INTO `product` VALUES (55, 'product_54', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (56, 'product_55', 0.72712, 994, 0);
INSERT INTO `product` VALUES (57, 'product_56', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (58, 'product_57', 0.72712, 994, 0);
INSERT INTO `product` VALUES (59, 'product_58', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (60, 'product_59', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (61, 'product_60', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (62, 'product_61', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (63, 'product_62', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (64, 'product_63', 0.72712, 995, 0);
INSERT INTO `product` VALUES (65, 'product_64', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (66, 'product_65', 0.72712, 998, 0);
INSERT INTO `product` VALUES (67, 'product_66', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (68, 'product_67', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (69, 'product_68', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (70, 'product_69', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (71, 'product_70', 0.72712, 994, 0);
INSERT INTO `product` VALUES (72, 'product_71', 0.72712, 995, 0);
INSERT INTO `product` VALUES (73, 'product_72', 0.72712, 996, 0);
INSERT INTO `product` VALUES (74, 'product_73', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (75, 'product_74', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (76, 'product_75', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (77, 'product_76', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (78, 'product_77', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (79, 'product_78', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (80, 'product_79', 0.72712, 986, 0);
INSERT INTO `product` VALUES (81, 'product_80', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (82, 'product_81', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (83, 'product_82', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (84, 'product_83', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (85, 'product_84', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (86, 'product_85', 0.72712, 997, 0);
INSERT INTO `product` VALUES (87, 'product_86', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (88, 'product_87', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (89, 'product_88', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (90, 'product_89', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (91, 'product_90', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (92, 'product_91', 0.72712, 997, 0);
INSERT INTO `product` VALUES (93, 'product_92', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (94, 'product_93', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (95, 'product_94', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (96, 'product_95', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (97, 'product_96', 0.72712, 990, 0);
INSERT INTO `product` VALUES (98, 'product_97', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (99, 'product_98', 0.72712, 1000, 0);
INSERT INTO `product` VALUES (100, 'product_99', 0.72712, 1000, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
