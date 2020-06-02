CREATE DATABASE IF NOT EXISTS `system`;
USE `system`;

DROP TABLE IF EXISTS `md_members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `md_members` (
  `member_id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(30) NOT NULL DEFAULT '',
  `nickname` varchar(30) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `description` varchar(640) NOT NULL DEFAULT '',
  `email` varchar(100) NOT NULL DEFAULT '',
  `phone` varchar(20) DEFAULT 'null',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `role` int(11) NOT NULL DEFAULT '1',
  `status` int(11) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL,
  `create_at` int(11) NOT NULL DEFAULT '0',
  `last_login_time` datetime DEFAULT NULL,
  PRIMARY KEY (`member_id`),
  UNIQUE KEY `account` (`account`),
  UNIQUE KEY `nickname` (`nickname`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `md_members`
--

LOCK TABLES `md_members` WRITE;
/*!40000 ALTER TABLE `md_members` DISABLE KEYS */;
INSERT INTO `md_members` VALUES (1,'admin','admin','123qweasd','','admin@ziyoubiancheng.com','','/static/images/avatar.png',0,0,'2019-12-16 06:13:31',0,'2019-12-16 14:13:31'),(2,'user1','user1','123456','','user1@ziyoubiancheng.com','','/static/images/avatar.png',2,0,'2019-12-19 17:04:26',0,'2019-12-20 01:04:26');
/*!40000 ALTER TABLE `md_members` ENABLE KEYS */;
UNLOCK TABLES;