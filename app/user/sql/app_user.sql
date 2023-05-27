-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- 主机： localhost:3306
-- 生成日期： 2023-04-08 10:03:31
-- 服务器版本： 8.0.29
-- PHP 版本： 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `gva_xiangxiang`
--

-- --------------------------------------------------------

--
-- 表的结构 `app_user`
--

CREATE TABLE `app_user` (
  `id` bigint NOT NULL,
  `mobile` longtext,
  `password` longtext,
  `nick_name` longtext,
  `sign` longtext,
  `avatar` longtext,
  `province` longtext,
  `city` longtext,
  `district` longtext,
  `birthday` timestamp NOT NULL,
  `register_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `is_male` bigint DEFAULT '0' COMMENT '0女 1男',
  `real_name` varchar(64) NOT NULL,
  `identity` varchar(64) NOT NULL,
  `authtime` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


--
-- 转储表的索引
--

--
-- 表的索引 `app_user`
--
ALTER TABLE `app_user`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `app_user`
--
ALTER TABLE `app_user`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
