-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.26 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.0.0.6029
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 stock 的数据库结构
DROP DATABASE IF EXISTS `stock`;
CREATE DATABASE IF NOT EXISTS `stock` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `stock`;

-- 导出  表 stock.history 结构
DROP TABLE IF EXISTS `history`;
CREATE TABLE IF NOT EXISTS `history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` int(11) DEFAULT '600519',
  `name` varchar(50) DEFAULT '',
  `day_time` int(11) DEFAULT '0' COMMENT '20200715',
  `price` double DEFAULT '0' COMMENT '当天收盘价',
  `dividend` varchar(50) DEFAULT '' COMMENT '分红说明',
  `share_per` double DEFAULT '0' COMMENT '每股赠送股数',
  `money_per` double DEFAULT '0' COMMENT '每股分红金额',
  PRIMARY KEY (`id`),
  KEY `dayTime` (`day_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

-- 正在导出表  stock.history 的数据：~21 rows (大约)
/*!40000 ALTER TABLE `history` DISABLE KEYS */;
INSERT INTO `history` (`id`, `code`, `name`, `day_time`, `price`, `dividend`, `share_per`, `money_per`) VALUES
	(1, 600519, '贵州茅台', 20010827, 35.55, '', 0, 0),
	(2, 600519, '贵州茅台', 20020725, 32.31, '10转1股派6元', 0.1, 0.6),
	(3, 600519, '贵州茅台', 20030714, 23.78, '10送1股派2元', 0.1, 0.2),
	(4, 600519, '贵州茅台', 20040701, 26.73, '10转3股派3元', 0.3, 0.3),
	(5, 600519, '贵州茅台', 20050805, 46.65, '10转2股派5元', 0.2, 0.5),
	(6, 600519, '贵州茅台', 20060519, 39.59, '10转10股派3元', 1, 0.3),
	(7, 600519, '贵州茅台', 20060524, 39.59, '10派5.91元', 0, 0.591),
	(8, 600519, '贵州茅台', 20070713, 114.49, '10派7元', 0, 0.7),
	(9, 600519, '贵州茅台', 20080616, 144.5, '10派8.36', 0, 0.836),
	(10, 600519, '贵州茅台', 20090701, 145.65, '10派11.56', 0, 1.156),
	(11, 600519, '贵州茅台', 20100705, 125.81, '10派11.85', 0, 1.185),
	(12, 600519, '贵州茅台', 20110701, 191.63, '10送1股派23元', 0.1, 2.3),
	(13, 600519, '贵州茅台', 20120705, 251.66, '10派39.97', 0, 3.997),
	(14, 600519, '贵州茅台', 20130607, 196.26, '10派64.19', 0, 6.419),
	(15, 600519, '贵州茅台', 20140625, 142.23, '10送1派43.74', 0.1, 4.374),
	(16, 600519, '贵州茅台', 20150717, 228.29, '10送1派43.74', 0.1, 4.374),
	(17, 600519, '贵州茅台', 20160701, 286.17, '10派61.71', 0, 6.171),
	(18, 600519, '贵州茅台', 20170707, 445.97, '10派67.87', 0, 6.787),
	(19, 600519, '贵州茅台', 20180615, 773.33, '10派109.99', 0, 10.999),
	(20, 600519, '贵州茅台', 20190628, 984, '10派145.39', 0, 14.539),
	(21, 600519, '贵州茅台', 20200624, 1460.01, '10派170.25', 0, 17.025);
/*!40000 ALTER TABLE `history` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
