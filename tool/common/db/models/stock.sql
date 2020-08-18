/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : localhost:3306
 Source Schema         : stock

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 15/08/2020 10:40:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for history
-- ----------------------------
DROP TABLE IF EXISTS `history`;
CREATE TABLE `history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` int(11) DEFAULT '600036',
  `name` varchar(50) DEFAULT '',
  `day_time` int(11) DEFAULT '0' COMMENT '20200715',
  `typ` int(11) DEFAULT '0' COMMENT '0-分红或购买；1-增发配股',
  `price` double DEFAULT '0' COMMENT '当天收盘价',
  `dividend` varchar(50) DEFAULT '' COMMENT '分红说明',
  `share_per` double DEFAULT '0' COMMENT '每股赠送股数',
  `money_per` double DEFAULT '0' COMMENT '每股分红金额',
  `buy_price` double DEFAULT '0' COMMENT '增发配售价格',
  PRIMARY KEY (`id`),
  KEY `dayTime` (`day_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8 COMMENT='回溯历史上优秀的股票';

-- ----------------------------
-- Records of history
-- ----------------------------
BEGIN;
INSERT INTO `history` VALUES (1, 600519, '贵州茅台', 20010827, 0, 35.55, '无', 0, 0, 0);
INSERT INTO `history` VALUES (2, 600519, '贵州茅台', 20020725, 0, 32.31, '10转1股派6元', 0.1, 0.6, 0);
INSERT INTO `history` VALUES (3, 600519, '贵州茅台', 20030714, 0, 23.78, '10送1股派2元', 0.1, 0.2, 0);
INSERT INTO `history` VALUES (4, 600519, '贵州茅台', 20040701, 0, 26.73, '10转3股派3元', 0.3, 0.3, 0);
INSERT INTO `history` VALUES (5, 600519, '贵州茅台', 20050805, 0, 46.65, '10转2股派5元', 0.2, 0.5, 0);
INSERT INTO `history` VALUES (6, 600519, '贵州茅台', 20060519, 0, 39.59, '10转10股派3元', 1, 0.3, 0);
INSERT INTO `history` VALUES (7, 600519, '贵州茅台', 20060524, 0, 39.59, '10派5.91元', 0, 0.591, 0);
INSERT INTO `history` VALUES (8, 600519, '贵州茅台', 20070713, 0, 114.49, '10派7元', 0, 0.7, 0);
INSERT INTO `history` VALUES (9, 600519, '贵州茅台', 20080616, 0, 144.5, '10派8.36元', 0, 0.836, 0);
INSERT INTO `history` VALUES (10, 600519, '贵州茅台', 20090701, 0, 145.65, '10派11.56元', 0, 1.156, 0);
INSERT INTO `history` VALUES (11, 600519, '贵州茅台', 20100705, 0, 125.81, '10派11.85元', 0, 1.185, 0);
INSERT INTO `history` VALUES (12, 600519, '贵州茅台', 20110701, 0, 191.63, '10送1股派23元', 0.1, 2.3, 0);
INSERT INTO `history` VALUES (13, 600519, '贵州茅台', 20120705, 0, 251.66, '10派39.97元', 0, 3.997, 0);
INSERT INTO `history` VALUES (14, 600519, '贵州茅台', 20130607, 0, 196.26, '10派64.19元', 0, 6.419, 0);
INSERT INTO `history` VALUES (15, 600519, '贵州茅台', 20140625, 0, 142.23, '10送1派43.74元', 0.1, 4.374, 0);
INSERT INTO `history` VALUES (16, 600519, '贵州茅台', 20150717, 0, 228.29, '10送1派43.74元', 0.1, 4.374, 0);
INSERT INTO `history` VALUES (17, 600519, '贵州茅台', 20160701, 0, 286.17, '10派61.71元', 0, 6.171, 0);
INSERT INTO `history` VALUES (18, 600519, '贵州茅台', 20170707, 0, 445.97, '10派67.87元', 0, 6.787, 0);
INSERT INTO `history` VALUES (19, 600519, '贵州茅台', 20180615, 0, 773.33, '10派109.99元', 0, 10.999, 0);
INSERT INTO `history` VALUES (20, 600519, '贵州茅台', 20190628, 0, 984, '10派145.39元', 0, 14.539, 0);
INSERT INTO `history` VALUES (21, 600519, '贵州茅台', 20200624, 0, 1460.01, '10派170.25元', 0, 17.025, 0);
INSERT INTO `history` VALUES (22, 601318, '中国平安', 20070301, 0, 46.79, '无', 0, 0, 0);
INSERT INTO `history` VALUES (23, 601318, '中国平安', 20070622, 0, 74.79, '10派2.2元', 0, 0.22, 0);
INSERT INTO `history` VALUES (24, 601318, '中国平安', 20070903, 0, 103.07, '10派2元', 0, 0.2, 0);
INSERT INTO `history` VALUES (25, 601318, '中国平安', 20080523, 0, 55.79, '10派5元', 0, 0.5, 0);
INSERT INTO `history` VALUES (26, 601318, '中国平安', 20081006, 0, 33.4, '10派2元', 0, 0.2, 0);
INSERT INTO `history` VALUES (27, 601318, '中国平安', 20090901, 0, 44.92, '10派1.5元', 0, 0.15, 0);
INSERT INTO `history` VALUES (28, 601318, '中国平安', 20100713, 0, 48.61, '10派3元', 0, 0.3, 0);
INSERT INTO `history` VALUES (29, 601318, '中国平安', 20100909, 0, 48.49, '10派1.5元', 0, 0.15, 0);
INSERT INTO `history` VALUES (30, 601318, '中国平安', 20110721, 0, 45.98, '10派4元', 0, 0.4, 0);
INSERT INTO `history` VALUES (31, 601318, '中国平安', 20110902, 0, 41.12, '10派1.5元', 0, 0.15, 0);
INSERT INTO `history` VALUES (32, 601318, '中国平安', 20120716, 0, 43.57, '10派2.5元', 0, 0.25, 0);
INSERT INTO `history` VALUES (33, 601318, '中国平安', 20120926, 0, 40.35, '10派1.5元', 0, 0.15, 0);
INSERT INTO `history` VALUES (34, 601318, '中国平安', 20130520, 0, 40.86, '10派3元', 0, 0.3, 0);
INSERT INTO `history` VALUES (35, 601318, '中国平安', 20130910, 0, 38.16, '10派2元', 0, 0.2, 0);
INSERT INTO `history` VALUES (36, 601318, '中国平安', 20140627, 0, 38.87, '10派4.5元', 0, 0.45, 0);
INSERT INTO `history` VALUES (37, 601318, '中国平安', 20140912, 0, 42.88, '10派2.5元', 0, 0.25, 0);
INSERT INTO `history` VALUES (38, 601318, '中国平安', 20150727, 0, 36.56, '10转10派5元', 1, 0.5, 0);
INSERT INTO `history` VALUES (39, 601318, '中国平安', 20150909, 0, 30.22, '10派1.8元', 0, 0.18, 0);
INSERT INTO `history` VALUES (40, 601318, '中国平安', 20160705, 0, 32.07, '10派3.5元', 0, 0.35, 0);
INSERT INTO `history` VALUES (41, 601318, '中国平安', 20160905, 0, 34.76, '10派2元', 0, 0.2, 0);
INSERT INTO `history` VALUES (42, 601318, '中国平安', 20170711, 0, 51.23, '10派5.5元', 0, 0.55, 0);
INSERT INTO `history` VALUES (43, 601318, '中国平安', 20170904, 0, 55.15, '10派5元', 0, 0.5, 0);
INSERT INTO `history` VALUES (44, 601318, '中国平安', 20180607, 0, 63.24, '10派12元', 0, 1.2, 0);
INSERT INTO `history` VALUES (45, 601318, '中国平安', 20180906, 0, 61.32, '10派6.2元', 0, 0.62, 0);
INSERT INTO `history` VALUES (46, 601318, '中国平安', 20190523, 0, 75.44, '10派11元', 0, 1.1, 0);
INSERT INTO `history` VALUES (47, 601318, '中国平安', 20190904, 0, 88.3, '10派7.5元', 0, 0.75, 0);
INSERT INTO `history` VALUES (48, 601318, '中国平安', 20200508, 0, 72.97, '10派13元', 0, 1.3, 0);
INSERT INTO `history` VALUES (49, 2, '万科A', 19910129, 0, 14.58, '无', 0, 0, 0);
INSERT INTO `history` VALUES (50, 2, '万科A', 19910608, 0, 7.95, '10送2股', 0.2, 0, 0);
INSERT INTO `history` VALUES (51, 2, '万科A', 19920320, 0, 11.1, '10送2股', 0.2, 0, 0);
INSERT INTO `history` VALUES (52, 2, '万科A', 19930405, 0, 20.5, '10送5股派0.6元', 0.5, 0.06, 0);
INSERT INTO `history` VALUES (53, 2, '万科A', 19940621, 0, 5.72, '10送3.5股派1.5元', 0.35, 0.15, 0);
INSERT INTO `history` VALUES (54, 2, '万科A', 19950704, 0, 3.97, '10送1.5股派1.5元', 0.15, 0.15, 0);
INSERT INTO `history` VALUES (55, 2, '万科A', 19960806, 0, 7.64, '10送1股派1.4元', 0.1, 0.14, 0);
INSERT INTO `history` VALUES (56, 2, '万科A', 19970627, 0, 21.09, '10送1.5股派1元', 0.15, 0.1, 0);
INSERT INTO `history` VALUES (57, 2, '万科A', 19980710, 0, 11.6, '10送1股派1.5元', 0.1, 0.15, 0);
INSERT INTO `history` VALUES (58, 2, '万科A', 19990806, 0, 11.9, '10送1股派1元', 0.1, 0.1, 0);
INSERT INTO `history` VALUES (59, 2, '万科A', 20000817, 0, 14.42, '10派1.5元', 0, 0.15, 0);
INSERT INTO `history` VALUES (60, 2, '万科A', 20010821, 0, 14.39, '10派1.8元', 0, 0.18, 0);
INSERT INTO `history` VALUES (61, 2, '万科A', 20020717, 0, 12.38, '10派2元', 0, 0.2, 0);
INSERT INTO `history` VALUES (62, 2, '万科A', 20030523, 0, 6.79, '10转10股派2元', 1, 0.2, 0);
INSERT INTO `history` VALUES (63, 2, '万科A', 20040526, 0, 4.94, '10送1股转4股派0.5元', 0.5, 0.05, 0);
INSERT INTO `history` VALUES (64, 2, '万科A', 20050629, 0, 3.24, '10转5股派1.5元', 0.5, 0.15, 0);
INSERT INTO `history` VALUES (65, 2, '万科A', 20060721, 0, 6.04, '10派1.5元', 0, 0.15, 0);
INSERT INTO `history` VALUES (66, 2, '万科A', 20070516, 0, 15.72, '10转5股派1.5元', 0.5, 0.15, 0);
INSERT INTO `history` VALUES (67, 2, '万科A', 20080616, 0, 10.6, '10转6股派1元', 0.6, 0.1, 0);
INSERT INTO `history` VALUES (68, 2, '万科A', 20090608, 0, 11.14, '10派0.5元', 0, 0.05, 0);
INSERT INTO `history` VALUES (69, 2, '万科A', 20100518, 0, 7.34, '10派0.7元', 0, 0.07, 0);
INSERT INTO `history` VALUES (70, 2, '万科A', 20110527, 0, 7.9, '10派1元', 0, 0.1, 0);
INSERT INTO `history` VALUES (71, 2, '万科A', 20120705, 0, 9.29, '10派1.3元', 0, 0.13, 0);
INSERT INTO `history` VALUES (72, 2, '万科A', 20130516, 0, 11.7, '10派1.8元', 0, 0.18, 0);
INSERT INTO `history` VALUES (73, 2, '万科A', 20140508, 0, 7.35, '10派4.1元', 0, 0.41, 0);
INSERT INTO `history` VALUES (74, 2, '万科A', 20150721, 0, 14.51, '10派5元', 0, 0.5, 0);
INSERT INTO `history` VALUES (75, 2, '万科A', 20160729, 0, 17.14, '10派7.2元', 0, 0.72, 0);
INSERT INTO `history` VALUES (76, 2, '万科A', 20170829, 0, 23.36, '10派7.9元', 0, 0.79, 0);
INSERT INTO `history` VALUES (77, 2, '万科A', 20180823, 0, 22.82, '10派9元', 0, 0.9, 0);
INSERT INTO `history` VALUES (78, 2, '万科A', 20190815, 0, 26.34, '10派10.45102元', 0, 1.045102, 0);
INSERT INTO `history` VALUES (79, 600036, '招商银行', 20020409, 0, 10.66, '无', 0, 0, 0);
INSERT INTO `history` VALUES (80, 600036, '招商银行', 20030716, 0, 12.1, '10派1.2元', 0, 0.12, 0);
INSERT INTO `history` VALUES (81, 600036, '招商银行', 20040511, 0, 8.94, '10转2股派0.92元', 0.2, 0.092, 0);
INSERT INTO `history` VALUES (82, 600036, '招商银行', 20050620, 0, 6.19, '10转5派1.1元', 0.5, 0.11, 0);
INSERT INTO `history` VALUES (83, 600036, '招商银行', 20060224, 0, 6.66, '10转0.8589股', 0.08589, 0, 0);
INSERT INTO `history` VALUES (84, 600036, '招商银行', 20060616, 0, 7.16, '10派0.8元', 0, 0.08, 0);
INSERT INTO `history` VALUES (85, 600036, '招商银行', 20060921, 0, 9.63, '10派1.8元', 0, 0.18, 0);
INSERT INTO `history` VALUES (86, 600036, '招商银行', 20070704, 0, 23.93, '10派1.2元', 0, 0.12, 0);
INSERT INTO `history` VALUES (87, 600036, '招商银行', 20080728, 0, 24.74, '10派2.8元', 0, 0.28, 0);
INSERT INTO `history` VALUES (88, 600036, '招商银行', 20090703, 0, 18.59, '10送3股派1元', 0.3, 0.1, 0);
INSERT INTO `history` VALUES (89, 600036, '招商银行', 20100701, 0, 12.81, '10派2.1元', 0, 0.21, 0);
INSERT INTO `history` VALUES (90, 600036, '招商银行', 20110610, 0, 12.78, '10派2.9元', 0, 0.29, 0);
INSERT INTO `history` VALUES (91, 600036, '招商银行', 20120607, 0, 11.25, '10派4.2元', 0, 0.42, 0);
INSERT INTO `history` VALUES (92, 600036, '招商银行', 20130613, 0, 12.32, '10派6.3元', 0, 0.63, 0);
INSERT INTO `history` VALUES (93, 600036, '招商银行', 20140711, 0, 9.69, '10派6.2元', 0, 0.62, 0);
INSERT INTO `history` VALUES (94, 600036, '招商银行', 20150703, 0, 17.61, '10派6.7元', 0, 0.67, 0);
INSERT INTO `history` VALUES (95, 600036, '招商银行', 20160713, 0, 17.05, '10派6.9元', 0, 0.69, 0);
INSERT INTO `history` VALUES (96, 600036, '招商银行', 20170614, 0, 21.3, '10派7.4元', 0, 0.74, 0);
INSERT INTO `history` VALUES (97, 600036, '招商银行', 20180712, 0, 26.04, '10派8.4元', 0, 0.84, 0);
INSERT INTO `history` VALUES (98, 600036, '招商银行', 20190712, 0, 35.4, '10派9.4元', 0, 0.94, 0);
INSERT INTO `history` VALUES (99, 600036, '招商银行', 20200710, 0, 37.25, '10派12元', 0, 1.2, 0);
INSERT INTO `history` VALUES (100, 600036, '招商银行', 20130905, 1, 10.66, '10配1.74股', 0.174, 0, 9.29);
INSERT INTO `history` VALUES (101, 600036, '招商银行', 20100315, 1, 14.89, '10配1.3股', 0.13, 0, 8.85);
INSERT INTO `history` VALUES (102, 2, '万科A', 20000110, 1, 11.44, '10配2.727股', 0.2727, 0, 7.5);
INSERT INTO `history` VALUES (103, 2, '万科A', 19970714, 1, 15.6, '10配2.37股', 0.237, 0, 4.5);
INSERT INTO `history` VALUES (104, 2, '万科A', 19910601, 1, 11.29, '10配5股', 0.5, 0, 4.4);
COMMIT;

-- ----------------------------
-- Table structure for invest
-- ----------------------------
DROP TABLE IF EXISTS `invest`;
CREATE TABLE `invest` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `month_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '年月',
  `capital` double unsigned NOT NULL DEFAULT '0' COMMENT '当月本金',
  `profit` double NOT NULL DEFAULT '0',
  `uid` int(11) NOT NULL DEFAULT '1314',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=95 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of invest
-- ----------------------------
BEGIN;
INSERT INTO `invest` VALUES (1, 201608, 4288, 638, 1);
INSERT INTO `invest` VALUES (2, 201609, 17088, 1020, 1);
INSERT INTO `invest` VALUES (3, 201610, 32338, -1893, 1);
INSERT INTO `invest` VALUES (4, 201611, 38881, -2393, 1);
INSERT INTO `invest` VALUES (5, 201612, 78121, -4108, 1);
INSERT INTO `invest` VALUES (6, 201701, 88227, 3515.88, 1);
INSERT INTO `invest` VALUES (7, 201702, 98984, 1515, 1);
INSERT INTO `invest` VALUES (8, 201703, 81999, 1304, 1);
INSERT INTO `invest` VALUES (9, 201704, 96303, -2117.64, 1);
INSERT INTO `invest` VALUES (10, 201705, 133686, 3926, 1);
INSERT INTO `invest` VALUES (11, 201706, 134212, 3776.92, 1);
INSERT INTO `invest` VALUES (12, 201707, 145988, 10620, 1);
INSERT INTO `invest` VALUES (13, 201708, 169710, 3144.16, 1);
INSERT INTO `invest` VALUES (14, 201709, 193708, -550, 1);
INSERT INTO `invest` VALUES (15, 201710, 207695, 4862, 1);
INSERT INTO `invest` VALUES (16, 201711, 211575, 3799, 1);
INSERT INTO `invest` VALUES (17, 201712, 256017, 236, 1);
INSERT INTO `invest` VALUES (18, 201801, 284512, 13924, 1);
INSERT INTO `invest` VALUES (19, 201802, 300005, -23307, 1);
INSERT INTO `invest` VALUES (20, 201803, 338834, -16802, 1);
INSERT INTO `invest` VALUES (21, 201804, 317781, -15739.52, 1);
INSERT INTO `invest` VALUES (22, 201805, 322593, -165.3, 1);
INSERT INTO `invest` VALUES (23, 201806, 320717, -15392.25, 1);
INSERT INTO `invest` VALUES (24, 201807, 315315, 11179.29, 1);
INSERT INTO `invest` VALUES (25, 201808, 303018, 2000.51, 1);
INSERT INTO `invest` VALUES (26, 201809, 304963, 17281.5, 1);
INSERT INTO `invest` VALUES (27, 201810, 304653, -5712.8, 1);
INSERT INTO `invest` VALUES (28, 201811, 290520, -195, 1);
INSERT INTO `invest` VALUES (29, 201812, 279852, -23994, 1);
INSERT INTO `invest` VALUES (30, 201901, 279056, 30708.85, 1);
INSERT INTO `invest` VALUES (31, 201902, 304874, 29341, 1);
INSERT INTO `invest` VALUES (32, 201903, 311967, 20242, 1);
INSERT INTO `invest` VALUES (33, 201904, 331306, 31364, 1);
INSERT INTO `invest` VALUES (34, 201905, 372794, -26018, 1);
INSERT INTO `invest` VALUES (35, 201906, 332894, 28797, 1);
INSERT INTO `invest` VALUES (36, 201907, 359431, 14606, 1);
INSERT INTO `invest` VALUES (37, 201908, 383956, -8875.18, 1);
INSERT INTO `invest` VALUES (38, 201909, 408794, 6368.39, 1);
INSERT INTO `invest` VALUES (39, 201910, 439362, 14370.81, 1);
INSERT INTO `invest` VALUES (40, 201911, 454984, -3990.08, 1);
INSERT INTO `invest` VALUES (41, 201912, 446994, 22224.37, 1);
INSERT INTO `invest` VALUES (42, 202001, 475256, -17398.28, 1);
INSERT INTO `invest` VALUES (43, 202002, 479741.63, -20391.67, 1);
INSERT INTO `invest` VALUES (44, 202003, 470015.96, -46833.24, 1);
INSERT INTO `invest` VALUES (45, 202004, 468666.72, 31113.64, 1);
INSERT INTO `invest` VALUES (46, 202005, 499780.36, -16467.06, 1);
INSERT INTO `invest` VALUES (47, 202006, 477840.3, 7743.92, 1);
INSERT INTO `invest` VALUES (48, 202007, 476305.04, 25766.63, 1);
INSERT INTO `invest` VALUES (49, 202008, 499666, 0, 1);
INSERT INTO `invest` VALUES (50, 202002, 10000, -623.03, 2);
INSERT INTO `invest` VALUES (51, 202003, 29376.97, -2516.2, 2);
INSERT INTO `invest` VALUES (52, 202004, 26860.77, 1984.3, 2);
INSERT INTO `invest` VALUES (53, 202005, 28845.07, -1083.15, 2);
INSERT INTO `invest` VALUES (54, 202006, 27761.92, 131.87, 2);
INSERT INTO `invest` VALUES (55, 202007, 27893.79, 2545, 2);
INSERT INTO `invest` VALUES (56, 202008, 30439, 0, 2);
INSERT INTO `invest` VALUES (57, 202003, 50000, -117.28, 3);
INSERT INTO `invest` VALUES (58, 202004, 49882.72, 2870.2, 3);
INSERT INTO `invest` VALUES (59, 202005, 52752.92, -1832, 3);
INSERT INTO `invest` VALUES (60, 202006, 50920.92, 1056.8, 3);
INSERT INTO `invest` VALUES (61, 202007, 51977.72, 2936, 3);
INSERT INTO `invest` VALUES (62, 202008, 54913, 0, 3);
INSERT INTO `invest` VALUES (63, 201801, 2000, 0, 1314);
INSERT INTO `invest` VALUES (64, 201802, 4000, 0, 1314);
INSERT INTO `invest` VALUES (65, 201803, 6000, 0, 1314);
INSERT INTO `invest` VALUES (66, 201804, 8000, 0, 1314);
INSERT INTO `invest` VALUES (67, 201805, 10000, 0, 1314);
INSERT INTO `invest` VALUES (68, 201806, 12000, 0, 1314);
INSERT INTO `invest` VALUES (69, 201807, 14000, 0, 1314);
INSERT INTO `invest` VALUES (70, 201808, 16000, 0, 1314);
INSERT INTO `invest` VALUES (71, 201809, 18000, 0, 1314);
INSERT INTO `invest` VALUES (72, 201810, 20000, 0, 1314);
INSERT INTO `invest` VALUES (73, 201811, 22000, 0, 1314);
INSERT INTO `invest` VALUES (74, 201812, 24000, -2623, 1314);
INSERT INTO `invest` VALUES (75, 201901, 23477, 0, 1314);
INSERT INTO `invest` VALUES (76, 201902, 25577, 0, 1314);
INSERT INTO `invest` VALUES (77, 201903, 27677, 0, 1314);
INSERT INTO `invest` VALUES (78, 201904, 29777, 0, 1314);
INSERT INTO `invest` VALUES (79, 201905, 31877, 0, 1314);
INSERT INTO `invest` VALUES (80, 201906, 33977, 0, 1314);
INSERT INTO `invest` VALUES (81, 201907, 36077, 0, 1314);
INSERT INTO `invest` VALUES (82, 201908, 38177, 0, 1314);
INSERT INTO `invest` VALUES (83, 201909, 40277, 6000, 1314);
INSERT INTO `invest` VALUES (84, 201910, 48377, 1030.46, 1314);
INSERT INTO `invest` VALUES (85, 201911, 51507.46, -748.87, 1314);
INSERT INTO `invest` VALUES (86, 201912, 52858.59, 3692.41, 1314);
INSERT INTO `invest` VALUES (87, 202001, 58756, -3231.07, 1314);
INSERT INTO `invest` VALUES (88, 202002, 57729.93, -2603.41, 1314);
INSERT INTO `invest` VALUES (89, 202003, 57331.52, -5826.12, 1314);
INSERT INTO `invest` VALUES (90, 202004, 53710.4, 3453.64, 1314);
INSERT INTO `invest` VALUES (91, 202005, 59369.04, -2006.84, 1314);
INSERT INTO `invest` VALUES (92, 202006, 59567.2, 598.48, 1314);
INSERT INTO `invest` VALUES (93, 202007, 62370.68, 4203.62, 1314);
INSERT INTO `invest` VALUES (94, 202008, 68779.3, 0, 1314);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `real_name` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1315 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'clearluo', '');
INSERT INTO `user` VALUES (2, 'luosg', '');
INSERT INTO `user` VALUES (3, 'zhuangjx', '');
INSERT INTO `user` VALUES (1314, 'clearluo_retire', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
