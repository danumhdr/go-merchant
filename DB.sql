CREATE DATABASE  IF NOT EXISTS `merchant` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `merchant`;
-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: localhost    Database: merchant
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `merchants`
--

DROP TABLE IF EXISTS `merchants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `merchants` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `merchant_name` varchar(40) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` bigint NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `merchants`
--

LOCK TABLES `merchants` WRITE;
/*!40000 ALTER TABLE `merchants` DISABLE KEYS */;
INSERT INTO `merchants` VALUES (1,1,'merchant 1','2022-09-16 15:27:22',1,'2022-09-16 15:27:22',1),(2,2,'Merchant 2','2022-09-16 15:27:22',2,'2022-09-16 15:27:22',2);
/*!40000 ALTER TABLE `merchants` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `outlets`
--

DROP TABLE IF EXISTS `outlets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `outlets` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `merchant_id` bigint NOT NULL,
  `outlet_name` varchar(40) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` bigint NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `outlets`
--

LOCK TABLES `outlets` WRITE;
/*!40000 ALTER TABLE `outlets` DISABLE KEYS */;
INSERT INTO `outlets` VALUES (1,1,'Outlet 1','2022-09-16 15:27:22',1,'2022-09-16 15:27:22',1),(2,2,'Outlet 1','2022-09-16 15:27:22',2,'2022-09-16 15:27:22',2),(3,1,'Outlet 2','2022-09-16 15:27:22',1,'2022-09-16 15:27:22',1);
/*!40000 ALTER TABLE `outlets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `merchant_id` bigint NOT NULL,
  `outlet_id` bigint NOT NULL,
  `bill_total` double NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` bigint NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (1,1,1,2000,'2021-11-01 05:30:04',1,'2021-11-01 05:30:04',1),(2,1,1,2500,'2021-11-01 10:20:14',1,'2021-11-01 10:20:14',1),(3,1,1,4000,'2021-11-02 05:30:04',1,'2021-11-02 05:30:04',1),(4,1,1,1000,'2021-11-04 05:30:04',1,'2021-11-04 05:30:04',1),(5,1,1,7000,'2021-11-05 09:59:30',1,'2021-11-05 09:59:30',1),(6,1,3,2000,'2021-11-02 11:30:04',1,'2021-11-02 11:30:04',1),(7,1,3,2500,'2021-11-03 10:20:14',1,'2021-11-03 10:20:14',1),(8,1,3,4000,'2021-11-04 05:30:04',1,'2021-11-04 05:30:04',1),(9,1,3,1000,'2021-11-04 05:31:04',1,'2021-11-04 05:31:04',1),(10,1,3,7000,'2021-11-05 09:59:30',1,'2021-11-05 09:59:30',1),(11,2,2,2000,'2021-11-01 11:30:04',2,'2021-11-01 11:30:04',2),(12,2,2,2500,'2021-11-02 10:20:14',2,'2021-11-02 10:20:14',2),(13,2,2,4000,'2021-11-03 05:30:04',2,'2021-11-03 05:30:04',2),(14,2,2,1000,'2021-11-04 05:31:04',2,'2021-11-04 05:31:04',2),(15,2,2,7000,'2021-11-05 09:59:30',2,'2021-11-05 09:59:30',2),(16,2,2,2000,'2021-11-05 11:30:04',2,'2021-11-05 11:30:04',2),(17,2,2,2500,'2021-11-06 10:20:14',2,'2021-11-06 10:20:14',2),(18,2,2,4000,'2021-11-07 05:30:04',2,'2021-11-07 05:30:04',2),(19,2,2,1000,'2021-11-08 05:31:04',2,'2021-11-08 05:31:04',2),(20,2,2,7000,'2021-11-09 09:59:30',2,'2021-11-09 09:59:30',2),(21,2,2,1000,'2021-11-10 05:31:04',2,'2021-11-10 05:31:04',2),(22,2,2,7000,'2021-11-11 09:59:30',2,'2021-11-11 09:59:30',2);
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `user_name` varchar(45) DEFAULT NULL,
  `password` varchar(225) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` bigint NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Admin 1','admin1','21232f297a57a5a743894a0e4a801fc3','2022-09-16 15:27:22',1,'2022-09-16 15:27:22',1),(2,'Admin 2','admin2','21232f297a57a5a743894a0e4a801fc3','2022-09-16 15:27:22',2,'2022-09-16 15:27:22',2);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-18  8:04:31
