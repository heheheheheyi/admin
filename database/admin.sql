-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2021-12-24 07:11:59
-- 服务器版本： 8.0.17
-- PHP 版本： 7.3.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `admin`
--

-- --------------------------------------------------------

--
-- 表的结构 `borrow_record`
--

CREATE TABLE `borrow_record` (
  `id` int(11) NOT NULL,
  `rid` int(11) NOT NULL,
  `cid` int(11) NOT NULL,
  `uid` int(11) NOT NULL,
  `mid` int(11) NOT NULL,
  `time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `borrow_record`
--

INSERT INTO `borrow_record` (`id`, `rid`, `cid`, `uid`, `mid`, `time`) VALUES
(1, 2, 2, 2, 1, '2021-12-19 13:27:33'),
(2, 3, 3, 1, 2, '2021-12-20 13:30:09'),
(3, 5, 1, 1, 2, '2021-12-20 13:30:55'),
(4, 4, 4, 1, 1, '2021-12-20 21:05:50');

-- --------------------------------------------------------

--
-- 表的结构 `cate`
--

CREATE TABLE `cate` (
  `id` int(11) NOT NULL,
  `catename` varchar(32) NOT NULL,
  `num` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `cate`
--

INSERT INTO `cate` (`id`, `catename`, `num`) VALUES
(1, '桌子', 3),
(2, '电脑', 3),
(3, '椅子', 3),
(4, '教室', 3);

-- --------------------------------------------------------

--
-- 表的结构 `chat`
--

CREATE TABLE `chat` (
  `id` int(11) NOT NULL,
  `uid` int(11) NOT NULL,
  `fid` int(11) NOT NULL,
  `content` varchar(100) NOT NULL,
  `time` datetime NOT NULL,
  `isread` int(11) NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `chat`
--

INSERT INTO `chat` (`id`, `uid`, `fid`, `content`, `time`, `isread`) VALUES
(1, 1, 2, '你申请的电脑1通过申请。', '2021-12-20 13:27:33', 2),
(2, 2, 1, '你申请的椅子1通过申请。', '2021-12-20 13:30:09', 2),
(3, 2, 1, '你申请的桌子2通过申请。', '2021-12-20 13:30:55', 2),
(4, 2, 1, '你借用的桌子2归还成功。', '2021-12-20 13:31:06', 2),
(5, 1, 1, '你申请的教室1通过申请。', '2021-12-20 21:05:50', 2),
(6, 1, 2, '你借用的电脑1归还成功。', '2021-12-20 21:06:15', 1),
(7, 1, 1, '你借用的椅子1未按期归还。', '2021-12-20 21:06:59', 2);

-- --------------------------------------------------------

--
-- 表的结构 `res`
--

CREATE TABLE `res` (
  `id` int(11) NOT NULL,
  `resname` varchar(32) NOT NULL,
  `cid` int(11) NOT NULL,
  `img` varchar(100) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `uid` int(11) DEFAULT NULL,
  `borrow_time` datetime DEFAULT NULL,
  `return_time` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `res`
--

INSERT INTO `res` (`id`, `resname`, `cid`, `img`, `status`, `uid`, `borrow_time`, `return_time`) VALUES
(1, '桌子1', 1, 'http://r3kwcz679.hn-bkt.clouddn.com/FqH16RdEphFWcZAzV5YOZwca-2JR', 1, NULL, NULL, NULL),
(2, '电脑1', 2, 'http://r3kwcz679.hn-bkt.clouddn.com/FnRPmW_Hei6YlKm6agTSplFMgb5_', 1, NULL, NULL, NULL),
(3, '椅子1', 3, 'http://r3kwcz679.hn-bkt.clouddn.com/FqqdGn2lDe61jOubhcsnpp7HOqqm', 3, 1, '2021-12-21 00:00:00', '2021-12-19 00:00:00'),
(4, '教室1', 4, 'http://r3kwcz679.hn-bkt.clouddn.com/Fh8x9D-bTkkopVKOJqOKVHt1O3xr', 3, 1, '2021-12-21 00:00:00', '2021-12-24 00:00:00'),
(5, '桌子2', 1, 'http://r3kwcz679.hn-bkt.clouddn.com/FqH16RdEphFWcZAzV5YOZwca-2JR', 1, NULL, NULL, NULL),
(6, '桌子3', 1, 'http://r3kwcz679.hn-bkt.clouddn.com/FqH16RdEphFWcZAzV5YOZwca-2JR', 1, NULL, NULL, NULL),
(7, '椅子2', 3, 'http://r3kwcz679.hn-bkt.clouddn.com/FqqdGn2lDe61jOubhcsnpp7HOqqm', 1, NULL, NULL, NULL),
(8, '椅子3', 3, 'http://r3kwcz679.hn-bkt.clouddn.com/FqqdGn2lDe61jOubhcsnpp7HOqqm', 1, NULL, NULL, NULL),
(9, '教室2', 4, 'http://r3kwcz679.hn-bkt.clouddn.com/Fh8x9D-bTkkopVKOJqOKVHt1O3xr', 1, NULL, NULL, NULL),
(10, '电脑2', 2, 'http://r3kwcz679.hn-bkt.clouddn.com/FnRPmW_Hei6YlKm6agTSplFMgb5_', 1, NULL, NULL, NULL),
(11, '教室3', 4, 'http://r3kwcz679.hn-bkt.clouddn.com/Fh8x9D-bTkkopVKOJqOKVHt1O3xr', 1, NULL, NULL, NULL),
(12, '电脑3', 2, 'http://r3kwcz679.hn-bkt.clouddn.com/FnRPmW_Hei6YlKm6agTSplFMgb5_', 1, NULL, NULL, NULL);

-- --------------------------------------------------------

--
-- 表的结构 `return_record`
--

CREATE TABLE `return_record` (
  `id` int(11) NOT NULL,
  `rid` int(11) NOT NULL,
  `cid` int(11) NOT NULL,
  `uid` int(11) NOT NULL,
  `mid` int(11) NOT NULL,
  `time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `return_record`
--

INSERT INTO `return_record` (`id`, `rid`, `cid`, `uid`, `mid`, `time`) VALUES
(1, 5, 1, 1, 2, '2021-12-20 13:31:06'),
(2, 2, 2, 2, 1, '2021-12-20 21:06:15');

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(32) NOT NULL,
  `num` int(11) NOT NULL,
  `password` varchar(32) NOT NULL,
  `role` int(11) NOT NULL DEFAULT '2'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `users`
--

INSERT INTO `users` (`id`, `username`, `num`, `password`, `role`) VALUES
(1, 'admin', 2, 'X5/L/sWah1CB2A==', 1),
(2, 'admin1', 0, 'X5/L/sWah1CB2A==', 1),
(3, 'user', 0, 'X5/L/sWah1CB2A==', 2),
(4, 'user1', 0, 'X5/L/sWah1CB2A==', 2);

--
-- 转储表的索引
--

--
-- 表的索引 `borrow_record`
--
ALTER TABLE `borrow_record`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `borrow_record_id_uindex` (`id`);

--
-- 表的索引 `cate`
--
ALTER TABLE `cate`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `cate_catename_uindex` (`catename`),
  ADD UNIQUE KEY `cate_id_uindex` (`id`);

--
-- 表的索引 `chat`
--
ALTER TABLE `chat`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `chat_id_uindex` (`id`);

--
-- 表的索引 `res`
--
ALTER TABLE `res`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `res_id_uindex` (`id`);

--
-- 表的索引 `return_record`
--
ALTER TABLE `return_record`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `return_record_id_uindex` (`id`);

--
-- 表的索引 `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_username_uindex` (`username`),
  ADD UNIQUE KEY `users_id_uindex` (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `borrow_record`
--
ALTER TABLE `borrow_record`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `cate`
--
ALTER TABLE `cate`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `chat`
--
ALTER TABLE `chat`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- 使用表AUTO_INCREMENT `res`
--
ALTER TABLE `res`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- 使用表AUTO_INCREMENT `return_record`
--
ALTER TABLE `return_record`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
