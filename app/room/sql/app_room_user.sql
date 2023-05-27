-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- 主机： localhost:3306
-- 生成日期： 2023-04-08 09:59:56
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
-- 数据库： `room`
--

-- --------------------------------------------------------

--
-- 表的结构 `app_room_user`
--

CREATE TABLE `app_room_user` (
  `id` bigint NOT NULL,
  `party_id` int NOT NULL COMMENT '房间id',
  `user_id` int NOT NULL COMMENT '用户id',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `leave_at` timestamp NOT NULL COMMENT '离开时间',
  `cool_at` timestamp NOT NULL COMMENT '冷却时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3
PARTITION BY HASH (`id`)
PARTITIONS 100;

--
-- 转储表的索引
--

--
-- 表的索引 `app_room_user`
--
ALTER TABLE `app_room_user`
  ADD PRIMARY KEY (`id`),
  ADD KEY `party_id` (`party_id`,`user_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `app_room_user`
--
ALTER TABLE `app_room_user`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
