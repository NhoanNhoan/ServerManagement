-- MySQL dump 10.17  Distrib 10.3.25-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: ServerManagement
-- ------------------------------------------------------
-- Server version	10.3.25-MariaDB-0ubuntu0.20.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `CABLE_TYPE`
--

DROP TABLE IF EXISTS `CABLE_TYPE`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `CABLE_TYPE` (
  `id` varchar(6) NOT NULL,
  `name` varchar(12) CHARACTER SET utf8 DEFAULT NULL,
  `sign_port` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CABLE_TYPE`
--

LOCK TABLES `CABLE_TYPE` WRITE;
/*!40000 ALTER TABLE `CABLE_TYPE` DISABLE KEYS */;
INSERT INTO `CABLE_TYPE` VALUES ('CT0001','type 1','xe'),('CT0002','type 2','xe');
/*!40000 ALTER TABLE `CABLE_TYPE` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `CURATOR`
--

DROP TABLE IF EXISTS `CURATOR`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `CURATOR` (
  `id_ERROR` varchar(6) NOT NULL,
  `id_PERSON` varchar(6) NOT NULL,
  PRIMARY KEY (`id_ERROR`,`id_PERSON`),
  KEY `id_PERSON` (`id_PERSON`),
  CONSTRAINT `CURATOR_ibfk_1` FOREIGN KEY (`id_ERROR`) REFERENCES `ERROR` (`id`),
  CONSTRAINT `CURATOR_ibfk_2` FOREIGN KEY (`id_PERSON`) REFERENCES `PERSON` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CURATOR`
--

LOCK TABLES `CURATOR` WRITE;
/*!40000 ALTER TABLE `CURATOR` DISABLE KEYS */;
INSERT INTO `CURATOR` VALUES ('ERR001','PS0001'),('ERR002','PS0001');
/*!40000 ALTER TABLE `CURATOR` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DC`
--

DROP TABLE IF EXISTS `DC`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `DC` (
  `id` varchar(6) NOT NULL,
  `description` varchar(20) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DC`
--

LOCK TABLES `DC` WRITE;
/*!40000 ALTER TABLE `DC` DISABLE KEYS */;
INSERT INTO `DC` VALUES ('DC0001','DC 9'),('DC0002','DC 08');
/*!40000 ALTER TABLE `DC` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ERROR`
--

DROP TABLE IF EXISTS `ERROR`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ERROR` (
  `id` varchar(6) NOT NULL,
  `summary` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `description` varchar(500) CHARACTER SET utf8 DEFAULT NULL,
  `solution` varchar(1000) CHARACTER SET utf8 DEFAULT NULL,
  `occurs` date DEFAULT NULL,
  `id_SERVER` varchar(12) DEFAULT NULL,
  `ID_ERROR_STATE` varchar(6) DEFAULT NULL,
  `id_STATUS_ROW` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_SERVER` (`id_SERVER`),
  KEY `id_ERROR_STATUS` (`ID_ERROR_STATE`),
  KEY `id_STATUS_ROW` (`id_STATUS_ROW`),
  CONSTRAINT `ERROR_ibfk_1` FOREIGN KEY (`id_SERVER`) REFERENCES `SERVER` (`id`),
  CONSTRAINT `ERROR_ibfk_2` FOREIGN KEY (`ID_ERROR_STATE`) REFERENCES `ERROR_STATE` (`id`),
  CONSTRAINT `ERROR_ibfk_3` FOREIGN KEY (`id_STATUS_ROW`) REFERENCES `STATUS_ROW` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ERROR`
--

LOCK TABLES `ERROR` WRITE;
/*!40000 ALTER TABLE `ERROR` DISABLE KEYS */;
INSERT INTO `ERROR` VALUES ('ERR001','abcd','hello world\r\n                    \r\n                    \r\n                    \r\n                    \r\n                    \r\n                    \r\n                    ','hello world\r\n                      \r\n                      \r\n                      \r\n                      \r\n                      \r\n                      \r\n                      ','2021-01-20','SV0000000001','ES0002',1),('ERR002','error','error','error','2021-01-20','SV0000000002','ES0001',1);
/*!40000 ALTER TABLE `ERROR` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ERROR_STATE`
--

DROP TABLE IF EXISTS `ERROR_STATE`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ERROR_STATE` (
  `id` varchar(6) NOT NULL,
  `description` varchar(20) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ERROR_STATE`
--

LOCK TABLES `ERROR_STATE` WRITE;
/*!40000 ALTER TABLE `ERROR_STATE` DISABLE KEYS */;
INSERT INTO `ERROR_STATE` VALUES ('ES0001','unresolved'),('ES0002','RESOLVED');
/*!40000 ALTER TABLE `ERROR_STATE` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `IP_NET`
--

DROP TABLE IF EXISTS `IP_NET`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `IP_NET` (
  `id` varchar(6) NOT NULL,
  `value` varchar(12) DEFAULT NULL,
  `id_STATUS_ROW` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_STATUS_ROW` (`id_STATUS_ROW`),
  CONSTRAINT `IP_NET_ibfk_1` FOREIGN KEY (`id_STATUS_ROW`) REFERENCES `STATUS_ROW` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `IP_NET`
--

LOCK TABLES `IP_NET` WRITE;
/*!40000 ALTER TABLE `IP_NET` DISABLE KEYS */;
INSERT INTO `IP_NET` VALUES ('IP0001','42.118.242.',NULL),('IP0002','192.168.1.',NULL);
/*!40000 ALTER TABLE `IP_NET` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `IP_SERVER`
--

DROP TABLE IF EXISTS `IP_SERVER`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `IP_SERVER` (
  `id_SERVER` varchar(12) NOT NULL,
  `id_IP_NET` varchar(6) NOT NULL,
  `ip_host` int(11) NOT NULL,
  `id_STATUS_ROW` int(11) DEFAULT NULL,
  PRIMARY KEY (`id_SERVER`,`id_IP_NET`,`ip_host`),
  KEY `id_IP_NET` (`id_IP_NET`),
  CONSTRAINT `IP_SERVER_ibfk_1` FOREIGN KEY (`id_SERVER`) REFERENCES `SERVER` (`id`),
  CONSTRAINT `IP_SERVER_ibfk_2` FOREIGN KEY (`id_IP_NET`) REFERENCES `IP_NET` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `IP_SERVER`
--

LOCK TABLES `IP_SERVER` WRITE;
/*!40000 ALTER TABLE `IP_SERVER` DISABLE KEYS */;
INSERT INTO `IP_SERVER` VALUES ('SV0000000001','IP0001',5,1),('SV0000000001','IP0001',9,1),('SV0000000002','IP0002',125,1);
/*!40000 ALTER TABLE `IP_SERVER` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `IP_SWITCH`
--

DROP TABLE IF EXISTS `IP_SWITCH`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `IP_SWITCH` (
  `id_SWITCH` varchar(10) NOT NULL,
  `id_IP_NET` varchar(6) NOT NULL,
  `ip_host` int(11) NOT NULL,
  PRIMARY KEY (`id_SWITCH`,`id_IP_NET`,`ip_host`),
  KEY `id_IP_NET` (`id_IP_NET`),
  CONSTRAINT `IP_SWITCH_ibfk_1` FOREIGN KEY (`id_SWITCH`) REFERENCES `SWITCH` (`id`),
  CONSTRAINT `IP_SWITCH_ibfk_2` FOREIGN KEY (`id_IP_NET`) REFERENCES `IP_NET` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `IP_SWITCH`
--

LOCK TABLES `IP_SWITCH` WRITE;
/*!40000 ALTER TABLE `IP_SWITCH` DISABLE KEYS */;
INSERT INTO `IP_SWITCH` VALUES ('SW00000001','IP0001',51),('SW00000002','IP0002',52);
/*!40000 ALTER TABLE `IP_SWITCH` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PERSON`
--

DROP TABLE IF EXISTS `PERSON`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `PERSON` (
  `id` varchar(6) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PERSON`
--

LOCK TABLES `PERSON` WRITE;
/*!40000 ALTER TABLE `PERSON` DISABLE KEYS */;
INSERT INTO `PERSON` VALUES ('PS0001','SYSADMIN 1'),('PS0002','SYSADMIN 2');
/*!40000 ALTER TABLE `PERSON` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PORT_TYPE`
--

DROP TABLE IF EXISTS `PORT_TYPE`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `PORT_TYPE` (
  `id` varchar(4) NOT NULL,
  `description` varchar(10) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PORT_TYPE`
--

LOCK TABLES `PORT_TYPE` WRITE;
/*!40000 ALTER TABLE `PORT_TYPE` DISABLE KEYS */;
INSERT INTO `PORT_TYPE` VALUES ('PT01','IDRAC'),('PT02','ILO');
/*!40000 ALTER TABLE `PORT_TYPE` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RACK`
--

DROP TABLE IF EXISTS `RACK`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RACK` (
  `id` varchar(6) NOT NULL,
  `description` varchar(20) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RACK`
--

LOCK TABLES `RACK` WRITE;
/*!40000 ALTER TABLE `RACK` DISABLE KEYS */;
INSERT INTO `RACK` VALUES ('RK0001','RACK 01'),('RK0002','RACK 02');
/*!40000 ALTER TABLE `RACK` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RACK_UNIT`
--

DROP TABLE IF EXISTS `RACK_UNIT`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RACK_UNIT` (
  `id` varchar(6) NOT NULL,
  `description` varchar(20) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RACK_UNIT`
--

LOCK TABLES `RACK_UNIT` WRITE;
/*!40000 ALTER TABLE `RACK_UNIT` DISABLE KEYS */;
INSERT INTO `RACK_UNIT` VALUES ('RU0001','U 30'),('RU0002','U31');
/*!40000 ALTER TABLE `RACK_UNIT` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SERVER`
--

DROP TABLE IF EXISTS `SERVER`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SERVER` (
  `id` varchar(12) NOT NULL,
  `id_DC` varchar(6) DEFAULT NULL,
  `id_RACK` varchar(6) DEFAULT NULL,
  `id_U_start` varchar(6) DEFAULT NULL,
  `id_U_end` varchar(6) DEFAULT NULL,
  `num_disks` int(11) DEFAULT NULL,
  `maker` varchar(10) DEFAULT NULL,
  `id_PORT_TYPE` varchar(4) DEFAULT NULL,
  `serial_number` varchar(20) DEFAULT NULL,
  `id_STATUS_ROW` int(11) DEFAULT NULL,
  `id_SERVER_STATUS` varchar(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_DC` (`id_DC`),
  KEY `id_RACK` (`id_RACK`),
  KEY `id_U_start` (`id_U_start`),
  KEY `id_U_end` (`id_U_end`),
  KEY `id_PORT_TYPE` (`id_PORT_TYPE`),
  KEY `id_STATUS_ROW` (`id_STATUS_ROW`),
  CONSTRAINT `SERVER_ibfk_1` FOREIGN KEY (`id_DC`) REFERENCES `DC` (`id`),
  CONSTRAINT `SERVER_ibfk_2` FOREIGN KEY (`id_RACK`) REFERENCES `RACK` (`id`),
  CONSTRAINT `SERVER_ibfk_3` FOREIGN KEY (`id_U_start`) REFERENCES `RACK_UNIT` (`id`),
  CONSTRAINT `SERVER_ibfk_4` FOREIGN KEY (`id_U_end`) REFERENCES `RACK_UNIT` (`id`),
  CONSTRAINT `SERVER_ibfk_5` FOREIGN KEY (`id_PORT_TYPE`) REFERENCES `PORT_TYPE` (`id`),
  CONSTRAINT `SERVER_ibfk_6` FOREIGN KEY (`id_STATUS_ROW`) REFERENCES `STATUS_ROW` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SERVER`
--

LOCK TABLES `SERVER` WRITE;
/*!40000 ALTER TABLE `SERVER` DISABLE KEYS */;
INSERT INTO `SERVER` VALUES ('SV0000000001','DC0001','RK0002','RU0001','RU0001',12,'','PT02','fewoigewf',1,'SS0002'),('SV0000000002','DC0002','RK0002','RU0001','RU0001',12,'','PT01','12345abcdefgh',1,'SS0001');
/*!40000 ALTER TABLE `SERVER` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SERVER_EVENT`
--

DROP TABLE IF EXISTS `SERVER_EVENT`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SERVER_EVENT` (
  `id` varchar(6) NOT NULL,
  `id_SERVER` varchar(12) DEFAULT NULL,
  `event` varchar(500) CHARACTER SET utf8 DEFAULT NULL,
  `occur_at` date DEFAULT NULL,
  `id_STATUS_ROW` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_SERVER` (`id_SERVER`),
  KEY `id_STATUS_ROW` (`id_STATUS_ROW`),
  CONSTRAINT `SERVER_EVENT_ibfk_1` FOREIGN KEY (`id_SERVER`) REFERENCES `SERVER` (`id`),
  CONSTRAINT `SERVER_EVENT_ibfk_2` FOREIGN KEY (`id_STATUS_ROW`) REFERENCES `STATUS_ROW` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SERVER_EVENT`
--

LOCK TABLES `SERVER_EVENT` WRITE;
/*!40000 ALTER TABLE `SERVER_EVENT` DISABLE KEYS */;
/*!40000 ALTER TABLE `SERVER_EVENT` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SERVER_STATUS`
--

DROP TABLE IF EXISTS `SERVER_STATUS`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SERVER_STATUS` (
  `id` varchar(6) NOT NULL,
  `description` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SERVER_STATUS`
--

LOCK TABLES `SERVER_STATUS` WRITE;
/*!40000 ALTER TABLE `SERVER_STATUS` DISABLE KEYS */;
INSERT INTO `SERVER_STATUS` VALUES ('SS0001','ACTIVE'),('SS0002','INTERACTIVE');
/*!40000 ALTER TABLE `SERVER_STATUS` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SERVICES`
--

DROP TABLE IF EXISTS `SERVICES`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SERVICES` (
  `id` varchar(6) NOT NULL,
  `id_SERVER` varchar(12) DEFAULT NULL,
  `service` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_SERVER` (`id_SERVER`),
  CONSTRAINT `SERVICES_ibfk_1` FOREIGN KEY (`id_SERVER`) REFERENCES `SERVER` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SERVICES`
--

LOCK TABLES `SERVICES` WRITE;
/*!40000 ALTER TABLE `SERVICES` DISABLE KEYS */;
INSERT INTO `SERVICES` VALUES ('SE0001','SV0000000001','SERVICE 01'),('SE0002','SV0000000001','SERVICE 02'),('SE0003','SV0000000001','SERVICE 03'),('SE0004','SV0000000001','SERVICE 04'),('SE0005','SV0000000002','SERVICE 01'),('SE0006','SV0000000002','SERVICE 02'),('SE0007','SV0000000002','SERVICE 03');
/*!40000 ALTER TABLE `SERVICES` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `STATUS_ROW`
--

DROP TABLE IF EXISTS `STATUS_ROW`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `STATUS_ROW` (
  `id` int(11) NOT NULL,
  `description` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `STATUS_ROW`
--

LOCK TABLES `STATUS_ROW` WRITE;
/*!40000 ALTER TABLE `STATUS_ROW` DISABLE KEYS */;
INSERT INTO `STATUS_ROW` VALUES (1,'available'),(2,'stop');
/*!40000 ALTER TABLE `STATUS_ROW` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SWITCH`
--

DROP TABLE IF EXISTS `SWITCH`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SWITCH` (
  `id` varchar(10) NOT NULL,
  `name` varchar(20) DEFAULT NULL,
  `id_DC` varchar(6) DEFAULT NULL,
  `id_RACK` varchar(6) DEFAULT NULL,
  `id_U_start` varchar(6) DEFAULT NULL,
  `id_U_end` varchar(6) DEFAULT NULL,
  `maximum_port` int(11) DEFAULT NULL,
  `id_STATUS_ROW` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_DC` (`id_DC`),
  KEY `id_RACK` (`id_RACK`),
  KEY `id_U_start` (`id_U_start`),
  KEY `id_U_end` (`id_U_end`),
  KEY `id_STATUS_ROW` (`id_STATUS_ROW`),
  CONSTRAINT `SWITCH_ibfk_1` FOREIGN KEY (`id_DC`) REFERENCES `DC` (`id`),
  CONSTRAINT `SWITCH_ibfk_2` FOREIGN KEY (`id_RACK`) REFERENCES `RACK` (`id`),
  CONSTRAINT `SWITCH_ibfk_3` FOREIGN KEY (`id_U_start`) REFERENCES `RACK_UNIT` (`id`),
  CONSTRAINT `SWITCH_ibfk_4` FOREIGN KEY (`id_U_end`) REFERENCES `RACK_UNIT` (`id`),
  CONSTRAINT `SWITCH_ibfk_5` FOREIGN KEY (`id_STATUS_ROW`) REFERENCES `STATUS_ROW` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SWITCH`
--

LOCK TABLES `SWITCH` WRITE;
/*!40000 ALTER TABLE `SWITCH` DISABLE KEYS */;
INSERT INTO `SWITCH` VALUES ('SW00000001','switch 01','DC0001','RK0001','RU0001','RU0002',90,1),('SW00000002','switch 01','DC0001','RK0001','RU0001','RU0002',90,1);
/*!40000 ALTER TABLE `SWITCH` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SWITCH_CONNECTION`
--

DROP TABLE IF EXISTS `SWITCH_CONNECTION`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SWITCH_CONNECTION` (
  `id` varchar(6) NOT NULL,
  `id_SERVER` varchar(12) DEFAULT NULL,
  `id_SWITCH` varchar(10) DEFAULT NULL,
  `id_CABLE_TYPE` varchar(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_SERVER` (`id_SERVER`),
  KEY `id_SWITCH` (`id_SWITCH`),
  KEY `id_CABLE_TYPE` (`id_CABLE_TYPE`),
  CONSTRAINT `SWITCH_CONNECTION_ibfk_1` FOREIGN KEY (`id_SERVER`) REFERENCES `SERVER` (`id`),
  CONSTRAINT `SWITCH_CONNECTION_ibfk_2` FOREIGN KEY (`id_SWITCH`) REFERENCES `SWITCH` (`id`),
  CONSTRAINT `SWITCH_CONNECTION_ibfk_3` FOREIGN KEY (`id_CABLE_TYPE`) REFERENCES `CABLE_TYPE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SWITCH_CONNECTION`
--

LOCK TABLES `SWITCH_CONNECTION` WRITE;
/*!40000 ALTER TABLE `SWITCH_CONNECTION` DISABLE KEYS */;
INSERT INTO `SWITCH_CONNECTION` VALUES ('SC0001','SV0000000002','SW00000001','CT0001'),('SC0002','SV0000000002','SW00000001','CT0002');
/*!40000 ALTER TABLE `SWITCH_CONNECTION` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SWITCH_CONNECTION_PORT`
--

DROP TABLE IF EXISTS `SWITCH_CONNECTION_PORT`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `SWITCH_CONNECTION_PORT` (
  `id_SWITCH` varchar(10) NOT NULL,
  `sv_port` int(11) NOT NULL,
  `switch_port` int(11) NOT NULL,
  PRIMARY KEY (`id_SWITCH`,`sv_port`,`switch_port`),
  CONSTRAINT `SWITCH_CONNECTION_PORT_ibfk_1` FOREIGN KEY (`id_SWITCH`) REFERENCES `SWITCH` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SWITCH_CONNECTION_PORT`
--

LOCK TABLES `SWITCH_CONNECTION_PORT` WRITE;
/*!40000 ALTER TABLE `SWITCH_CONNECTION_PORT` DISABLE KEYS */;
INSERT INTO `SWITCH_CONNECTION_PORT` VALUES ('SW00000001',7877,1221),('SW00000002',7877,1221);
/*!40000 ALTER TABLE `SWITCH_CONNECTION_PORT` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `USER`
--

DROP TABLE IF EXISTS `USER`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `USER` (
  `id` varchar(10) NOT NULL,
  `username` varchar(20) DEFAULT NULL,
  `pass` varchar(50) DEFAULT NULL,
  `id_STATUS_ROW` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_STATUS_ROW` (`id_STATUS_ROW`),
  CONSTRAINT `USER_ibfk_1` FOREIGN KEY (`id_STATUS_ROW`) REFERENCES `STATUS_ROW` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `USER`
--

LOCK TABLES `USER` WRITE;
/*!40000 ALTER TABLE `USER` DISABLE KEYS */;
/*!40000 ALTER TABLE `USER` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-02-02 17:26:19
