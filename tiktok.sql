-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: tiktok
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论id，自增主键',
  `user_id` bigint NOT NULL COMMENT '评论发布用户id',
  `video_id` bigint NOT NULL COMMENT '评论视频id',
  `comment_text` varchar(255) NOT NULL COMMENT '评论内容',
  `create_date` datetime NOT NULL COMMENT '评论发布时间',
  `cancel` tinyint NOT NULL DEFAULT '0' COMMENT '默认评论发布为0，取消后为1',
  PRIMARY KEY (`id`),
  KEY `videoIdIdx` (`video_id`) USING BTREE COMMENT '评论列表使用视频id作为索引-方便查看视频下的评论列表'
) ENGINE=InnoDB AUTO_INCREMENT=1224 DEFAULT CHARSET=utf8mb3 COMMENT='评论表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1206,2,135,'加油','2023-02-20 21:54:52',0),(1207,2,135,'冲呀!!!','2023-02-20 21:55:10',0),(1208,2,134,'真悠闲','2023-02-20 21:55:20',0),(1209,2,134,'祁门红茶','2023-02-20 21:55:35',0),(1210,3,135,'少年郎','2023-02-20 21:56:41',0),(1211,2,134,'大家好','2023-02-20 22:45:07',0),(1212,3,134,'楼上真帅','2023-02-20 22:45:37',0),(1213,4,133,'一楼','2023-02-20 23:46:46',0),(1214,4,136,'可爱的大熊猫','2023-02-21 16:43:36',0),(1215,3,137,'热烈的青春呐','2023-02-21 16:47:43',0),(1216,3,134,'测试下','2023-02-21 22:15:08',0),(1217,3,138,'测试评**能','2023-02-22 15:44:48',0),(1218,3,138,'删除评论操作','2023-02-22 16:15:19',0),(1219,6,139,'天赋不能给你刀刻般的肌肉','2023-02-22 21:47:09',0),(1220,6,139,'发表评论','2023-02-22 21:47:36',1),(1221,6,139,'发表评论','2023-02-22 21:47:52',1),(1222,7,141,'人的梦想是不会死去的','2023-02-23 17:31:49',0),(1223,7,141,'测试删除评**能','2023-02-23 17:32:01',1);
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `follows`
--

DROP TABLE IF EXISTS `follows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `follows` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `follower_id` bigint NOT NULL COMMENT '关注的用户',
  `cancel` tinyint NOT NULL DEFAULT '0' COMMENT '默认关注为0，取消关注为1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userIdToFollowerIdIdx` (`user_id`,`follower_id`) USING BTREE,
  KEY `FollowerIdIdx` (`follower_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1118 DEFAULT CHARSET=utf8mb3 COMMENT='关注表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `follows`
--

LOCK TABLES `follows` WRITE;
/*!40000 ALTER TABLE `follows` DISABLE KEYS */;
INSERT INTO `follows` VALUES (1096,2,4,1),(1103,3,4,2),(1106,4,4,1),(1111,3,2,1),(1112,2,3,1),(1113,3,3,1),(1114,3,6,1),(1115,6,3,1),(1116,3,7,1),(1117,7,3,1);
/*!40000 ALTER TABLE `follows` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `likes`
--

DROP TABLE IF EXISTS `likes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `likes` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '点赞用户id',
  `video_id` bigint NOT NULL COMMENT '被点赞的视频id',
  `cancel` tinyint NOT NULL DEFAULT '0' COMMENT '默认点赞为0，取消赞为1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userIdtoVideoIdIdx` (`user_id`,`video_id`) USING BTREE,
  KEY `userIdIdx` (`user_id`) USING BTREE,
  KEY `videoIdx` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1244 DEFAULT CHARSET=utf8mb3 COMMENT='点赞表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `likes`
--

LOCK TABLES `likes` WRITE;
/*!40000 ALTER TABLE `likes` DISABLE KEYS */;
INSERT INTO `likes` VALUES (1231,2,132,1),(1232,2,134,2),(1233,3,132,1),(1234,4,136,1),(1235,3,134,1),(1236,2,136,1),(1237,3,138,2),(1238,3,135,2),(1239,6,139,1),(1240,6,140,2),(1241,6,137,2),(1242,7,139,2),(1243,7,140,2);
/*!40000 ALTER TABLE `likes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `messages`
--

DROP TABLE IF EXISTS `messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `messages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `receiver_id` int DEFAULT NULL,
  `msg_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `have_get` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `messages`
--

LOCK TABLES `messages` WRITE;
/*!40000 ALTER TABLE `messages` DISABLE KEYS */;
INSERT INTO `messages` VALUES (10,2,20050,'你好呀',1676972289,0),(11,2,20050,'你在吗(⊙o⊙)！',1676972315,0),(12,3,2,'你好呀',1676985736,1),(13,2,3,'你也好呀',1676985764,1),(14,6,3,'在吗(⊙o⊙)！',1677074748,1),(15,3,6,'我在呀！',1677074771,1),(16,7,3,'在吗？',1677144820,1),(17,3,7,'我在的,,Ծ^Ծ,,',1677144849,1);
/*!40000 ALTER TABLE `messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `username` varchar(30) NOT NULL COMMENT '账号',
  `password` varchar(100) NOT NULL COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'test1','b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441'),(2,'vale','b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441'),(3,'辛弃疾','b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441'),(4,'曾经高三斜刘海','b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441'),(5,'橙子','b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441'),(6,'青山','b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441'),(7,'七缕晨阳','b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `videos`
--

DROP TABLE IF EXISTS `videos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键，视频唯一id',
  `author_id` bigint NOT NULL COMMENT '视频作者id',
  `play_url` varchar(255) NOT NULL COMMENT '播放url',
  `cover_url` varchar(255) NOT NULL COMMENT '封面url',
  `publish_time` datetime NOT NULL COMMENT '发布时间戳',
  `title` varchar(255) DEFAULT NULL COMMENT '视频名称',
  PRIMARY KEY (`id`),
  KEY `time` (`publish_time`) USING BTREE,
  KEY `author` (`author_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=142 DEFAULT CHARSET=utf8mb3 COMMENT='\r\n视频表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `videos`
--

LOCK TABLES `videos` WRITE;
/*!40000 ALTER TABLE `videos` DISABLE KEYS */;
INSERT INTO `videos` VALUES (132,2,'http://192.168.31.7:3000/static/videos/21676817886.mp4','http://192.168.31.7:3000/static/images/21676817886.jpg','2023-02-19 22:44:47','重庆'),(133,2,'http://192.168.31.7:3000/static/videos/21676818644.mp4','http://192.168.31.7:3000/static/images/21676818644.jpg','2023-02-19 22:57:24','开车'),(134,2,'http://192.168.31.7:3000/static/videos/21676818658.mp4','http://192.168.31.7:3000/static/images/21676818658.jpg','2023-02-19 22:57:39','喝茶'),(135,3,'http://192.168.31.7:3000/static/videos/31676891654.mp4','http://192.168.31.7:3000/static/images/31676891654.jpg','2023-02-20 19:14:15','张宇'),(136,2,'http://192.168.31.7:3000/static/videos/21676959959.mp4','http://192.168.31.7:3000/static/images/21676959959.jpg','2023-02-21 14:12:40','熊猫'),(137,4,'http://192.168.31.7:3000/static/videos/41676964184.mp4','http://192.168.31.7:3000/static/images/41676964184.jpg','2023-02-21 15:23:05','云轨电车'),(138,3,'http://192.168.31.7:3000/static/videos/31676981679.mp4','http://192.168.31.7:3000/static/images/31676981679.jpg','2023-02-21 20:14:40','写代码'),(139,3,'http://192.168.31.7:3000/static/videos/31677055810.mp4','http://192.168.31.7:3000/static/images/31677055810.jpg','2023-02-22 16:50:11','彭于晏'),(140,6,'http://192.168.31.7:3000/static/videos/61677074282.mp4','http://192.168.31.7:3000/static/images/61677074282.jpg','2023-02-22 21:58:03','天赋不会给你刀刻般的肌肉'),(141,7,'http://192.168.31.7:3000/static/videos/71677144636.mp4','http://192.168.31.7:3000/static/images/71677144636.jpg','2023-02-23 17:30:37','海贼王');
/*!40000 ALTER TABLE `videos` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-15 21:34:26
