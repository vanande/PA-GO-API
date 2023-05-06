-- MySQL dump 10.13  Distrib 8.0.33, for Linux (x86_64)
--
-- Host: localhost    Database: tas
-- ------------------------------------------------------
-- Server version	8.0.33

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
-- Table structure for table `ActiviteSalle`
--

DROP TABLE IF EXISTS `ActiviteSalle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ActiviteSalle` (
  `idActivite` int NOT NULL,
  `idSALLE` int NOT NULL,
  `idLIEU` int NOT NULL,
  PRIMARY KEY (`idActivite`,`idSALLE`,`idLIEU`),
  KEY `fk_activite_has_salle_salle1_idx` (`idSALLE`,`idLIEU`),
  KEY `fk_activite_has_salle_activite1_idx` (`idActivite`),
  CONSTRAINT `fk_activite_has_salle_activite1` FOREIGN KEY (`idActivite`) REFERENCES `activite` (`idActivite`),
  CONSTRAINT `fk_activite_has_salle_salle1` FOREIGN KEY (`idSALLE`, `idLIEU`) REFERENCES `salle` (`idSALLE`, `idLIEU`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ActiviteSalle`
--

LOCK TABLES `ActiviteSalle` WRITE;
/*!40000 ALTER TABLE `ActiviteSalle` DISABLE KEYS */;
/*!40000 ALTER TABLE `ActiviteSalle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `activite`
--

DROP TABLE IF EXISTS `activite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `activite` (
  `idActivite` int NOT NULL,
  `idlist_activite` int NOT NULL,
  `prix` int DEFAULT NULL,
  PRIMARY KEY (`idActivite`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activite`
--

LOCK TABLES `activite` WRITE;
/*!40000 ALTER TABLE `activite` DISABLE KEYS */;
INSERT INTO `activite` VALUES (1,1,1982),(2,20,2643),(3,9,1700),(4,16,1442),(5,17,1328),(6,2,2771),(7,10,2768),(8,8,1165),(9,14,2117),(10,6,257),(11,14,2538),(12,1,2237),(13,23,507),(14,15,1606),(15,12,2062),(16,9,947),(17,30,2327),(18,23,1656),(19,22,752),(20,20,401),(21,15,967),(22,22,275),(23,20,1038),(24,14,2399),(25,13,720),(26,1,1379),(27,19,1804),(28,28,1964),(29,6,2614),(30,11,201),(31,27,499),(32,24,492),(33,1,1831),(34,21,2962),(35,23,1404),(36,16,1897),(37,1,1408),(38,30,1454),(39,22,1064),(40,25,535),(41,24,2376),(42,26,2223),(43,7,922),(44,14,2982),(45,3,2536),(46,2,586),(47,1,1782),(48,10,1132),(49,11,1947),(50,6,2172),(51,14,2228),(52,1,2210),(53,2,1837),(54,15,2919),(55,1,251),(56,23,859),(57,25,2897),(58,2,2274),(59,4,1148),(60,14,2059),(61,23,1861),(62,2,1860),(63,20,845),(64,1,2368),(65,15,241),(66,3,1518),(67,24,652),(68,1,1351),(69,13,2454),(70,9,810),(71,10,2184),(72,8,1144),(73,13,584),(74,19,381),(75,15,1574),(76,24,1327),(77,30,2792),(78,24,810),(79,21,2642),(80,4,514),(81,4,1611),(82,9,2941),(83,3,487),(84,23,313),(85,24,1796),(86,14,1296),(87,13,207),(88,9,1017),(89,30,980),(90,22,2409),(91,13,605),(92,25,1546),(93,14,622),(94,28,955),(95,7,2767),(96,5,2969),(97,21,2651),(98,29,1762),(99,7,2080),(100,5,676),(101,20,2158),(102,1,649),(103,25,1900),(104,20,2844),(105,8,735),(106,20,1728),(107,3,2100),(108,20,2600),(109,1,2010),(110,19,970),(111,17,369),(112,26,2747),(113,17,2378),(114,12,2468),(115,29,2136),(116,21,1762),(117,16,1399),(118,13,2230),(119,14,2147),(120,15,2659),(121,28,2587),(122,30,2009),(123,26,383),(124,17,1172),(125,16,838),(126,13,2275),(127,2,1684),(128,2,2199),(129,19,2294),(130,11,2791),(131,7,1187),(132,27,2931),(133,20,1028),(134,25,2306),(135,28,473),(136,9,1254),(137,17,878),(138,18,631),(139,20,753),(140,18,2979),(141,29,1520),(142,3,334),(143,12,2863),(144,10,2116),(145,17,601),(146,24,282),(147,11,1014),(148,22,284),(149,21,2815),(150,11,1532),(151,4,468),(152,11,2021),(153,9,1585),(154,19,1156),(155,24,1926),(156,4,321),(157,13,2433),(158,6,738),(159,22,2797),(160,22,514),(161,9,1954),(162,1,2875),(163,26,2948),(164,3,535),(165,16,505),(166,18,2940),(167,17,1398),(168,19,2630),(169,14,557),(170,4,1460),(171,14,1947),(172,21,564),(173,11,2167),(174,12,834),(175,30,2453),(176,12,2016),(177,6,2830),(178,4,1326),(179,4,680),(180,6,2509),(181,6,989),(182,19,2219),(183,26,880),(184,16,222),(185,8,823),(186,25,2122),(187,28,2081),(188,28,2903),(189,19,1868),(190,28,2976),(191,13,1698),(192,11,2726),(193,3,2986),(194,20,495),(195,6,227),(196,20,220),(197,26,1644),(198,22,1462),(199,6,1237),(200,6,2318),(201,5,1040);
/*!40000 ALTER TABLE `activite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `activite_possible`
--

DROP TABLE IF EXISTS `activite_possible`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `activite_possible` (
  `idlist_activite` int NOT NULL,
  `idSALLE` int NOT NULL,
  PRIMARY KEY (`idlist_activite`,`idSALLE`),
  KEY `idSALLE` (`idSALLE`),
  CONSTRAINT `activite_possible_ibfk_1` FOREIGN KEY (`idlist_activite`) REFERENCES `list_activite` (`idlist_activite`),
  CONSTRAINT `activite_possible_ibfk_2` FOREIGN KEY (`idSALLE`) REFERENCES `salle` (`idSALLE`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activite_possible`
--

LOCK TABLES `activite_possible` WRITE;
/*!40000 ALTER TABLE `activite_possible` DISABLE KEYS */;
/*!40000 ALTER TABLE `activite_possible` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `anime`
--

DROP TABLE IF EXISTS `anime`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `anime` (
  `idlist_activite` int NOT NULL,
  `idPRESTATAIRE` int NOT NULL,
  PRIMARY KEY (`idlist_activite`,`idPRESTATAIRE`),
  KEY `fk_list_activite_has_prestataire_prestataire1_idx` (`idPRESTATAIRE`),
  KEY `fk_list_activite_has_prestataire_list_activite1_idx` (`idlist_activite`),
  CONSTRAINT `fk_list_activite_has_prestataire_list_activite1` FOREIGN KEY (`idlist_activite`) REFERENCES `list_activite` (`idlist_activite`),
  CONSTRAINT `fk_list_activite_has_prestataire_prestataire1` FOREIGN KEY (`idPRESTATAIRE`) REFERENCES `prestataire` (`idPRESTATAIRE`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `anime`
--

LOCK TABLES `anime` WRITE;
/*!40000 ALTER TABLE `anime` DISABLE KEYS */;
/*!40000 ALTER TABLE `anime` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `idCategory` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(45) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`idCategory`)
) ENGINE=MyISAM AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'Art','Activités artistiques pour développer la créativité'),(2,'Sport','Activités sportives pour encourager l\'esprit d\'équipe'),(3,'Cuisine','Activités culinaires pour favoriser la coopération'),(4,'Musique','Activités musicales pour stimuler la communication'),(5,'Théâtre','Activités théâtrales pour encourager la confiance en soi'),(6,'Jeux de société','Activités ludiques pour favoriser la détente'),(7,'Jeux de rôle','Activités de simulation pour encourager la prise de décision'),(8,'Développement personnel','Activités pour favoriser la croissance personnelle'),(9,'Communication','Activités pour améliorer les compétences de communication'),(10,'Leadership','Activités pour développer les compétences de leadership'),(11,'Marketing','Activités pour améliorer les compétences en marketing'),(12,'Innovation','Activités pour stimuler la créativité et l\'innovation'),(13,'Réflexion','Activités pour encourager la réflexion et la résolution de problèmes'),(14,'Découverte','Activités pour découvrir de nouvelles choses'),(15,'Nature','Activités en pleine nature pour encourager la prise de conscience environnementale'),(16,'Aventure','Activités d\'aventure pour stimuler l\'adrénaline'),(17,'Défi physique','Activités physiques pour se surpasser'),(18,'Défi intellectuel','Activités intellectuelles pour se dépasser'),(19,'Coopération','Activités pour encourager la coopération et la solidarité'),(20,'Connaissance de soi','Activités pour mieux se connaître soi-même'),(21,'Connaissance des autres','Activités pour mieux connaître les autres'),(22,'Confiance','Activités pour développer la confiance en soi et la confiance en l\'équipe'),(23,'Créativité','Activités pour stimuler la créativité');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category_activite`
--

DROP TABLE IF EXISTS `category_activite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category_activite` (
  `idCategory` int NOT NULL,
  `idlist_activite` int NOT NULL,
  PRIMARY KEY (`idCategory`,`idlist_activite`),
  KEY `idlist_activite` (`idlist_activite`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category_activite`
--

LOCK TABLES `category_activite` WRITE;
/*!40000 ALTER TABLE `category_activite` DISABLE KEYS */;
INSERT INTO `category_activite` VALUES (1,6),(1,25),(2,1),(2,15),(3,1),(4,7),(4,26),(5,11),(6,13),(7,2),(7,4),(8,5),(9,8),(10,16),(11,20),(12,2),(13,2),(13,22),(14,2),(14,3),(14,12),(15,14),(16,18),(17,1),(17,10),(18,2),(18,21),(19,3),(20,1),(20,3),(20,24),(21,1),(21,17),(22,1),(22,9),(23,19),(23,23);
/*!40000 ALTER TABLE `category_activite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `client`
--

DROP TABLE IF EXISTS `client`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `client` (
  `idCLIENT` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `token` varchar(64) DEFAULT NULL,
  `tel` char(10) DEFAULT NULL,
  `adresse` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `raison_sociale` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `statut` varchar(45) DEFAULT NULL,
  `creation_du_compte` date DEFAULT NULL,
  `pts_fidelite` int DEFAULT NULL,
  PRIMARY KEY (`idCLIENT`)
) ENGINE=InnoDB AUTO_INCREMENT=132 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `client`
--

LOCK TABLES `client` WRITE;
/*!40000 ALTER TABLE `client` DISABLE KEYS */;
INSERT INTO `client` VALUES (26,'ABC Company','contact@abccompany.com',NULL,'26','0123456789','123 Main St, Anytown, USA','ABC Company','Entreprise','2022-04-01',0),(27,'XYZ Corporation','info@xyzcorp.com',NULL,'27','0987654321','456 Oak St, Anytown, USA','XYZ Corporation','Entreprise','2022-04-01',0),(28,'Acme Inc.','sales@acmeinc.com',NULL,'28','555-1234','789 Maple Ave, Anytown, USA','Acme Inc.','Entreprise','2022-04-01',0),(29,'Globex Corporation','info@globexcorp.com',NULL,'29','555-4321','321 Pine St, Anytown, USA','Globex Corporation','Entreprise','2022-04-01',0),(30,'Initech','info@initech.com',NULL,'30','555-5678','654 Elm St, Anytown, USA','Initech','Entreprise','2022-04-01',0),(31,'Monarch Solutions','contact@monarchsolutions.com',NULL,'31','555-9876','987 Broadway, Anytown, USA','Monarch Solutions','Entreprise','2022-04-01',0),(32,'Oscorp Industries','info@oscorp.com',NULL,'32','555-2468','1358 Market St, Anytown, USA','Oscorp Industries','Entreprise','2022-04-01',0),(33,'Pied Piper','info@piedpiper.com',NULL,'33','555-3698','2468 Chestnut St, Anytown, USA','Pied Piper','Entreprise','2022-04-01',0),(34,'Stark Industries','info@starkindustries.com',NULL,'34','555-1357','2468 Park Ave, Anytown, USA','Stark Industries','Entreprise','2022-04-01',0),(35,'Wayne Enterprises','info@wayneent.com',NULL,'35','555-2468','123 Gotham Dr, Anytown, USA','Wayne Enterprises','Entreprise','2022-04-01',0),(46,'Air France','contact@airfrance.com',NULL,'46','0142771300','45 Rue de Paris, 95747 Roissy CDG Cedex','Air France-KLM','Société anonyme','2021-05-10',100),(47,'Orange','contact@orange.com',NULL,'47','0177370780','78 Rue Olivier de Serres, 75015 Paris','Orange S.A.','Société anonyme','2022-01-05',50),(48,'Peugeot','contact@peugeot.com',NULL,'48','0149822121','75 Avenue de la Grande Armée, 75116 Paris','Peugeot S.A.','Société anonyme','2021-09-20',200),(49,'Renault','contact@renault.com',NULL,'49','0141484343','13-15 Quai le Gallo, 92100 Boulogne-Billancourt','Renault S.A.','Société anonyme','2022-03-15',150),(50,'Société Générale','contact@socgen.com',NULL,'50','0144206060','29 Boulevard Haussmann, 75009 Paris','Société Générale S.A.','Société anonyme','2021-12-01',75),(51,'Total','contact@total.com',NULL,'51','0173574567','2 Place Jean Millier, La Défense 6, 92078 Paris La Défense','TotalEnergies SE','Société anonyme','2021-06-17',300),(52,'Thales','contact@thalesgroup.com',NULL,'52','0170733600','Tour Carpe Diem, 31 Place des Corolles, 92400 Courbevoie','Thales S.A.','Société anonyme','2022-02-14',80),(53,'Veolia','contact@veolia.com',NULL,'53','0175858100','30 Rue Madeleine Vionnet, 93300 Aubervilliers','Veolia Environnement S.A.','Société anonyme','2021-11-05',120),(54,'EDF','contact@edf.fr',NULL,'54','0158105050','22-30 Avenue de Wagram, 75008 Paris','Électricité de France S.A.','Société anonyme','2022-04-10',250),(55,'Capgemini','contact@capgemini.com',NULL,'55','0169294747','11 Rue de Tilsitt, 75017 Paris','Capgemini SE','Société européenne','2021-08-12',90),(56,'Air France','contact@airfrance.com',NULL,'56','0145678901','45 Rue de Paris, 75001 Paris','Air France-KLM','Entreprise','2022-01-01',0),(57,'Orange','contact@orange.com',NULL,'57','0145678902','78 Rue Olivier de Serres, 75015 Paris','Orange SA','Entreprise','2022-01-01',0),(58,'BNP Paribas','contact@bnpparibas.com',NULL,'58','0145678903','16 Boulevard des Italiens, 75009 Paris','BNP Paribas SA','Entreprise','2022-01-01',0),(59,'Total','contact@total.com',NULL,'59','0145678904','2 Place Jean Millier, 92400 Courbevoie','TotalEnergies SE','Entreprise','2022-01-01',0),(60,'Société Générale','contact@societegenerale.com',NULL,'60','0145678905','17 Cours Valmy, 92987 Paris La Défense','Société Générale S.A.','Entreprise','2022-01-01',0),(61,'Orange','contact@orange.com',NULL,'61','0145678902','78 Rue Olivier de Serres, 75015 Paris','Orange SA','Entreprise','2022-01-01',0),(62,'BNP Paribas','contact@bnpparibas.com',NULL,'62','0145678903','16 Boulevard des Italiens, 75009 Paris','BNP Paribas SA','Entreprise','2022-01-01',0),(63,'Total','contact@total.com',NULL,'63','0145678904','2 Place Jean Millier, 92400 Courbevoie','TotalEnergies SE','Entreprise','2022-01-01',0),(64,'Société Générale','contact@societegenerale.com',NULL,'64','0145678905','17 Cours Valmy, 92987 Paris La Défense','Société Générale S.A.','Entreprise','2022-01-01',0),(75,'LVMH','contact@lvmh.com',NULL,'75','0148411111','22 Avenue Montaigne, 75008 Paris','LVMH Moët Hennessy - Louis Vuitton SE','Société européenne','2021-07-05',150),(76,'BNP Paribas','contact@bnpparibas.com',NULL,'76','0144222222','16 Boulevard des Italiens, 75009 Paris','BNP Paribas S.A.','Société anonyme','2022-03-01',100),(77,'Air Liquide','contact@airliquide.com',NULL,'77','0169082121','75 Quai d\'Orsay, 75007 Paris','Air Liquide S.A.','Société anonyme','2021-10-18',200),(78,'Michelin','contact@michelin.com',NULL,'78','0473752424','23 Place des Carmes Déchaux, 63000 Clermont-Ferrand','Compagnie Générale des Établissements Michelin SCA','Société en commandite par actions','2022-01-02',75),(79,'Sodexo','contact@sodexo.com',NULL,'79','0155225200','255 Quai de la Bataille de Stalingrad, 92130 Issy-les-Moulineaux','Sodexo S.A.','Société anonyme','2021-12-15',120),(80,'Vinci','contact@vinci.com',NULL,'80','0148438000','1 Cours Ferdinand de Lesseps, 92851 Rueil-Malmaison','Vinci S.A.','Société anonyme','2022-02-28',90),(81,'Danone','contact@danone.com',NULL,'81','0153484646','17 Boulevard Haussmann, 75009 Paris','Danone S.A.','Société anonyme','2021-09-10',300),(82,'Sanofi','contact@sanofi.com',NULL,'82','0178745000','54 Rue La Boétie, 75008 Paris','Sanofi S.A.','Société anonyme','2022-04-05',150),(83,'Carrefour','contact@carrefour.com',NULL,'83','0156806060','33 Avenue Émile Zola, 92100 Boulogne-Billancourt','Carrefour S.A.','Société anonyme','2021-06-28',250),(84,'Saint-Gobain','contact@saint-gobain.com',NULL,'84','0145433333','18 Avenue d\'Alsace, 92400 Courbevoie','Compagnie de Saint-Gobain S.A.','Société anonyme','2022-01-20',80),(85,'Danone','contact@danone.com',NULL,'85','0155592000','17 Boulevard Haussmann, 75009 Paris','Danone S.A.','Société anonyme','2022-01-01',110),(86,'L\'Oréal','contact@loreal.com',NULL,'86','0156331800','41 Rue Martre, 92110 Clichy','L\'Oréal S.A.','Société anonyme','2021-11-15',220),(87,'Carrefour','contact@carrefour.com',NULL,'87','0169712100','33 Avenue Émile Zola, 92100 Boulogne-Billancourt','Carrefour S.A.','Société anonyme','2021-09-05',180),(88,'BNP Paribas','contact@bnpparibas.com',NULL,'88','0144488888','16 Boulevard des Italiens, 75009 Paris','BNP Paribas S.A.','Société anonyme','2022-03-30',90),(89,'Air Liquide','contact@airliquide.com',NULL,'89','0148607474','75 Quai d\'Orsay, 75007 Paris','Air Liquide S.A.','Société anonyme','2021-12-20',130),(90,'Michelin','contact@michelin.com',NULL,'90','0148357900','23 Place des Carmes-Dechaux, 63000 Clermont-Ferrand','Compagnie Générale des Établissements Michelin SCA','Société en commandite par actions','2021-06-30',250),(91,'Airbus','contact@airbus.com',NULL,'91','0157226200','2 Rond-Point Emile Dewoitine, 31700 Blagnac','Airbus S.A.S.','Société par actions simplifiée','2022-02-05',70),(92,'Sanofi','contact@sanofi.com',NULL,'92','0155446200','54 Rue La Boétie, 75008 Paris','Sanofi S.A.','Société anonyme','2021-07-22',190),(93,'Bouygues','contact@bouygues.com',NULL,'93','0155855000','32 Avenue Hoche, 75008 Paris','Bouygues S.A.','Société anonyme','2021-10-10',140),(94,'Accor','contact@accor.com',NULL,'94','0180052000','82 Rue Henri Farman, 92130 Issy-les-Moulineaux','Accor S.A.','Société anonyme','2022-04-25',100),(95,'Wayne Enterprises','contact@wayneenterprises.com',NULL,'95','5558675309','1007 Mountain Drive, Gotham City, USA','Wayne Enterprises','Corporation','2022-05-01',100),(96,'Stark Industries','contact@starkindustries.com',NULL,'96','5558675310','10880 Malibu Point, Malibu, USA','Stark Industries','Corporation','2022-05-01',100),(97,'Acme Corporation','contact@acme.com',NULL,'97','5558675311','123 Acme Street, Looney Tunes, USA','Acme Corporation','Corporation','2022-05-01',100),(98,'Oscorp Industries','contact@oscorp.com',NULL,'98','5558675312','5th Avenue, New York, USA','Oscorp Industries','Corporation','2022-05-01',100),(99,'Virtucon Industries','contact@virtucon.com',NULL,'99','5558675313','1 Evil Way, Austin Powers, USA','Virtucon Industries','Corporation','2022-05-01',100),(100,'Wonka Industries','contact@wonka.com',NULL,'100','5558675314','1 Wonka Way, Charlie and the Chocolate Factory, USA','Wonka Industries','Corporation','2022-05-01',100),(101,'Weyland-Yutani Corporation','contact@weyland-yutani.com',NULL,'101','5558675315','20th Floor, Tower B, The Blade Runner Building, Los Angeles, USA','Weyland-Yutani Corporation','Corporation','2022-05-01',100),(102,'Dunder Mifflin Paper Company','contact@dundermifflin.com',NULL,'102','5558675316','1725 Slough Avenue, Scranton, USA','Dunder Mifflin Paper Company','Corporation','2022-05-01',100),(103,'Globex Corporation','contact@globex.com',NULL,'103','5558675317','5th Avenue, New York, USA','Globex Corporation','Corporation','2022-05-01',100),(104,'BlackRock','contact@blackrock.com',NULL,'104','2128105300','55 East 52nd Street, New York, USA','BlackRock, Inc.','Corporation','2022-05-01',100),(125,'FantasticCompany','contact@fantcomp.com','notmymdp',NULL,'0612121212','3 rue de la Charette','FantasticCompany Corp','Societe anonyme',NULL,0),(126,'FantasticCompany','contact@fantcomp.com','notmymdp',NULL,'0612121212','3 rue de la Charette','FantasticCompany Corp','Societe anonyme',NULL,0),(127,'FantasticCompany','contact@fantcomp.com','notmymdp',NULL,'0612121212','3 rue de la Charette','FantasticCompany Corp','Societe anonyme',NULL,0),(128,'FantasticCompany','contact@fantcomp.com','notmymdp',NULL,'0612121212','3 rue de la Charette','FantasticCompany Corp','Societe anonyme','2023-04-24',0),(129,'FantasticCompany','contact@fantcomp.com','notmymdp',NULL,'0612121212','3 rue de la Charette','FantasticCompany Corp','Societe anonyme','2023-04-24',0),(130,'FantasticCompany','contact@fantcomp.com','notmymdp',NULL,'0612121212','3 rue de la Charette','FantasticCompany Corp','Societe anonyme','2023-04-24',0),(131,'FantasticCompany','contact@fantcomp.com','notmymdp',NULL,'0612121212','3 rue de la Charette','FantasticCompany Corp','Societe anonyme','2023-04-30',0);
/*!40000 ALTER TABLE `client` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `devis`
--

DROP TABLE IF EXISTS `devis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `devis` (
  `idDEVIS` int NOT NULL AUTO_INCREMENT,
  `titre` varchar(45) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `num_sirene` varchar(45) DEFAULT NULL,
  `somme_totale_a_payer` int DEFAULT NULL,
  `CLIENT_idCLIENT` int NOT NULL,
  PRIMARY KEY (`idDEVIS`,`CLIENT_idCLIENT`),
  KEY `fk_DEVIS_CLIENT` (`CLIENT_idCLIENT`),
  CONSTRAINT `fk_DEVIS_CLIENT` FOREIGN KEY (`CLIENT_idCLIENT`) REFERENCES `client` (`idCLIENT`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `devis`
--

LOCK TABLES `devis` WRITE;
/*!40000 ALTER TABLE `devis` DISABLE KEYS */;
INSERT INTO `devis` VALUES (1,'Devis Travaux Toiture','2020-05-20','123456789',5000,1),(2,'Devis Peinture Intérieure','2020-06-15','987654321',3000,2),(3,'Devis Électricité','2020-07-05','24681012',8000,3),(4,'Devis Plomberie','2020-08-10','369121518',4000,4),(5,'Devis Rénovation Salle de Bain','2020-09-12','1516171819',6000,5),(6,'Devis Aménagement Combles','2020-10-01','2021222324',7000,6),(7,'Devis Menuiserie','2020-11-15','2526272829',9000,7),(8,'Devis Chauffage','2020-12-05','3031323334',10000,8),(9,'Devis Isolation Thermique','2021-01-10','3536373839',4500,9),(10,'Devis Pose de Carrelage','2021-02-14','4041424344',5500,10),(11,'Devis Menuiserie Extérieure','2021-03-02','4546474849',8000,11),(12,'Devis Terrassement','2021-04-20','5051525354',12000,12),(13,'Devis Décoration Intérieure','2021-05-12','5556575859',5000,13),(14,'Devis Aménagement Jardin','2021-06-25','6061626364',7000,14),(15,'Devis Installation Alarme','2021-07-11','6566676869',2000,15),(16,'Devis Installation de Piscine','2021-08-05','7071727374',15000,16),(17,'Devis Réparation de Toiture','2021-09-17','7576777879',4000,17),(18,'Devis Installation de VMC','2021-10-30','8081828384',3000,18),(19,'Devis Construction de Garage','2021-11-20','8586878889',10000,19),(20,'Devis Travaux de Maçonnerie','2021-12-15','9091929394',6000,20),(21,'Devis Challenge sportif en équipe','2020-05-20','123456789',7000,1),(22,'Devis Atelier culinaire en équipe','2020-06-15','987654321',5500,2),(23,'Devis Escape game en entreprise','2020-07-05','24681012',8000,3),(24,'Devis Randonnée en équipe','2020-08-10','369121518',6500,4),(25,'Devis Course d\'orientation en entreprise','2020-09-12','1516171819',7500,5),(26,'Devis Jeux de rôle en entreprise','2020-10-01','2021222324',9000,6),(27,'Devis Activités créatives en équipe','2020-11-15','2526272829',6500,7),(28,'Devis Activités musicales en entreprise','2020-12-05','3031323334',8000,8),(29,'Devis Tournoi de jeux vidéo en équipe','2021-01-10','3536373839',6000,9),(30,'Devis Activités de construction en entreprise','2021-02-14','4041424344',7000,10),(31,'Jeu de piste sur Mars','2022-06-30','0987654321',12000,12),(32,'Safari en poney','2022-07-10','1234567890',8000,13),(33,'Lancer de chat','2022-08-01','1357924680',15000,14),(34,'Course de lits médicalisés','2022-09-15','2468135790',10000,15),(35,'Jeu de piste sur Mars','2022-06-30','0987654321',12000,12),(36,'Safari en poney','2022-07-10','1234567890',8000,13),(37,'Lancer de chat','2022-08-01','1357924680',15000,14),(38,'Course de lits médicalisés','2022-09-15','2468135790',10000,15),(39,'Escape game sous l\'eau','2022-10-05','3692581470',18000,16),(40,'Saut en parachute sans parachute 2.0','2022-11-01','4815162342',20000,17),(41,'Chasse aux papillons de nuit','2022-12-10','5142325262',9000,18),(42,'Concours de pets','2023-01-15','5768901234',6000,19),(43,'Tournoi de fléchettes avec des bananes','2023-02-14','6196374857',7000,20),(44,'Joutes en bateaux sur la Seine','2023-03-20','7283946571',11000,21);
/*!40000 ALTER TABLE `devis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `engage`
--

DROP TABLE IF EXISTS `engage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `engage` (
  `idTEAM_BUILDING` int NOT NULL,
  `idPRESTATAIRE` int NOT NULL,
  `prix_a_payer` int DEFAULT NULL,
  PRIMARY KEY (`idTEAM_BUILDING`,`idPRESTATAIRE`),
  KEY `fk_TEAM_BUILDING_has_PRESTATAIRE_PRESTATAIRE1` (`idPRESTATAIRE`),
  CONSTRAINT `fk_TEAM_BUILDING_has_PRESTATAIRE_PRESTATAIRE1` FOREIGN KEY (`idPRESTATAIRE`) REFERENCES `prestataire` (`idPRESTATAIRE`),
  CONSTRAINT `fk_TEAM_BUILDING_has_PRESTATAIRE_TEAM_BUILDING1` FOREIGN KEY (`idTEAM_BUILDING`) REFERENCES `team_building` (`idTEAM_BUILDING`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `engage`
--

LOCK TABLES `engage` WRITE;
/*!40000 ALTER TABLE `engage` DISABLE KEYS */;
INSERT INTO `engage` VALUES (1,1,500),(1,3,750),(1,8,1000),(1,9,13304),(1,20,3769),(1,31,6153),(1,39,19855),(2,2,600),(2,7,900),(2,10,1200),(2,11,1500),(2,14,4315),(2,19,18181),(2,26,3693),(2,45,18447),(3,4,800),(3,6,1200),(3,9,1500),(3,38,11261),(3,39,3809),(3,42,18241),(4,5,1000),(4,10,1500),(4,37,3922),(4,42,8483),(4,57,9551),(5,3,750),(5,6,1200),(5,7,900),(5,9,1500),(5,11,17844),(5,29,14633),(6,1,500),(6,8,1000),(6,41,3488),(6,45,17117),(7,2,600),(7,4,800),(7,5,1000),(7,7,741),(7,8,6011),(7,11,1500),(8,3,750),(8,6,1200),(8,7,900),(8,10,1200),(8,15,3953),(9,1,500),(9,4,800),(9,5,1000),(9,17,10012),(9,35,3677),(9,48,16430),(10,2,600),(10,6,1200),(10,8,1000),(10,14,15606),(10,22,5331),(10,31,1822),(10,43,14811),(10,49,10884),(11,3,750),(11,7,900),(11,8,9049),(11,9,1500),(11,33,764),(11,35,8187),(11,43,14769),(12,1,500),(12,2,600),(12,4,800),(12,6,1200),(12,36,14928),(12,39,8806),(13,3,750),(13,5,1000),(13,7,900),(13,10,1200),(13,40,11524),(13,58,15257),(14,1,500),(14,8,1000),(14,11,1500),(14,23,2480),(14,32,19391),(14,45,8094),(15,2,600),(15,3,750),(15,6,1200),(15,9,1500),(15,48,7431),(16,4,800),(16,5,1000),(16,7,900),(16,10,1200),(16,24,10903),(16,40,6011),(16,43,2234),(17,1,500),(17,2,600),(17,4,800),(17,6,1200),(17,9,15043),(17,51,12425),(17,54,17195),(18,3,750),(18,4,13662),(18,5,1000),(18,7,900),(18,10,19381),(18,11,1500),(18,24,4361),(18,36,10049),(19,1,500),(19,8,1000),(19,10,1200),(19,38,12008),(20,2,600),(20,3,750),(20,6,1200),(20,9,1500),(20,20,16376),(20,38,16330),(20,45,19542),(20,56,7305),(21,4,800),(21,5,1000),(21,7,900),(21,9,8138),(21,11,1500),(21,34,12199),(21,54,2543),(21,55,1177),(21,58,9011),(22,1,500),(22,6,1428),(22,20,19684),(22,33,17719),(22,38,12285),(22,46,3153),(22,49,691),(23,44,8204),(24,1,18562),(24,43,16111),(25,6,2314),(25,10,14894),(25,33,3335),(25,50,15741),(28,4,5860),(28,56,17320),(28,60,10401),(29,7,7229),(29,37,15955),(29,58,11851),(30,32,5436),(31,49,3226),(32,6,15852),(32,17,4246),(32,29,13758),(33,16,780),(34,22,1760),(34,62,18744),(35,6,4030),(36,12,18191),(37,9,16247),(37,41,7405),(38,2,12619),(38,13,3564),(38,25,3780),(38,34,14138),(38,39,8746),(38,41,8009),(38,42,14611),(38,59,821),(39,6,10964),(39,42,2971),(39,57,6401),(41,1,16885),(41,20,3744),(41,48,12515),(42,6,15903),(42,18,618),(42,36,7020),(42,57,11020),(43,5,6712),(43,16,18558),(43,30,6687),(43,60,6714),(44,54,14161);
/*!40000 ALTER TABLE `engage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `equipe`
--

DROP TABLE IF EXISTS `equipe`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `equipe` (
  `idEQUIPE` int NOT NULL AUTO_INCREMENT,
  `idPARTICIPANT` int NOT NULL,
  `idCLIENT` int NOT NULL,
  `idTEAM_BUILDING` int NOT NULL,
  PRIMARY KEY (`idEQUIPE`,`idPARTICIPANT`,`idCLIENT`,`idTEAM_BUILDING`),
  KEY `idPARTICIPANT` (`idPARTICIPANT`,`idCLIENT`),
  KEY `idTEAM_BUILDING` (`idTEAM_BUILDING`),
  CONSTRAINT `equipe_ibfk_1` FOREIGN KEY (`idPARTICIPANT`, `idCLIENT`) REFERENCES `participant` (`idPARTICIPANT`, `idCLIENT`),
  CONSTRAINT `equipe_ibfk_2` FOREIGN KEY (`idTEAM_BUILDING`) REFERENCES `team_building` (`idTEAM_BUILDING`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `equipe`
--

LOCK TABLES `equipe` WRITE;
/*!40000 ALTER TABLE `equipe` DISABLE KEYS */;
/*!40000 ALTER TABLE `equipe` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `facture`
--

DROP TABLE IF EXISTS `facture`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `facture` (
  `idFACTURE` int NOT NULL AUTO_INCREMENT,
  `date` date DEFAULT NULL,
  `somme_totale_a_payer` int DEFAULT NULL,
  `addresse_de_facturation` varchar(100) DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `CLIENT_idCLIENT` int NOT NULL,
  PRIMARY KEY (`idFACTURE`,`CLIENT_idCLIENT`),
  KEY `fk_FACTURE_CLIENT1` (`CLIENT_idCLIENT`),
  CONSTRAINT `fk_FACTURE_CLIENT1` FOREIGN KEY (`CLIENT_idCLIENT`) REFERENCES `client` (`idCLIENT`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `facture`
--

LOCK TABLES `facture` WRITE;
/*!40000 ALTER TABLE `facture` DISABLE KEYS */;
INSERT INTO `facture` VALUES (1,'2022-01-10',1500,'12 rue de la Paix, Paris','Prestations pour événement d\'entreprise',24),(1,'2023-04-15',5000,'123 Rue de la Paix, Paris','Prestation de service',25),(2,'2022-02-12',800,'15 rue de Rivoli, Lyon','Location de matériel pour soirée privée',57),(3,'2023-04-15',5000,'123 Rue de la Paix, Paris','Prestation de service',25),(3,'2022-03-20',3000,'10 rue des Bouchers, Lille','Organisation d\'un séminaire',89),(4,'2022-04-15',500,'8 avenue des Champs-Elysées, Paris','Fourniture de matériel audiovisuel',12),(4,'2023-04-10',7500,'1600 Pennsylvania Avenue NW, Washington, D.C.','Software development',81),(5,'2023-04-08',2500,'Rosenstrasse 2-4, 10178 Berlin','Graphic design',16),(5,'2022-05-25',2000,'25 avenue du Général de Gaulle, Toulouse','Location d\'un lieu pour un mariage',102),(6,'2023-04-03',18000,'2 Chome-2-1 Nishi-Shinjuku, Shinjuku City, Tokyo','Marketing strategy consulting',98),(7,'2023-04-12',12000,'10 Downing Street, London','Consulting services',15),(9,'2023-04-12',12000,'10 Downing Street, London','Consulting services',15),(11,'2022-05-02',12000,'22 rue de la Paix, Paris','Paiement pour service de consulting',22),(12,'2022-06-17',4000,'10 rue des Lilas, Lille','Paiement pour service de nettoyage',58),(13,'2022-01-15',8000,'5 avenue des Champs-Elysées, Paris','Paiement pour service de maintenance',75),(14,'2022-02-25',3500,'30 rue de la République, Lyon','Paiement pour service de dépannage',96),(15,'2022-04-19',6000,'14 avenue de la Gare, Marseille','Paiement pour service de conseil en marketing',14),(16,'2022-03-06',2500,'8 rue des Roses, Nantes','Paiement pour service de formation en communication',83),(17,'2022-07-08',5000,'12 boulevard des Dames, Nice','Paiement pour service de traduction',49),(18,'2022-06-05',9000,'3 rue du Faubourg Saint-Honoré, Paris','Paiement pour service de conseil en gestion',21),(20,'2022-05-02',12000,'22 rue de la Paix, Paris','Paiement pour service de consulting',22),(21,'2022-06-17',4000,'10 rue des Lilas, Lille','Paiement pour service de nettoyage',58),(22,'2022-01-15',8000,'5 avenue des Champs-Elysées, Paris','Paiement pour service de maintenance',75),(23,'2022-02-25',3500,'30 rue de la République, Lyon','Paiement pour service de dépannage',96),(24,'2022-04-19',6000,'14 avenue de la Gare, Marseille','Paiement pour service de conseil en marketing',14),(25,'2022-03-06',2500,'8 rue des Roses, Nantes','Paiement pour service de formation en communication',83),(26,'2022-07-08',5000,'12 boulevard des Dames, Nice','Paiement pour service de traduction',49),(27,'2022-06-05',9000,'3 rue du Faubourg Saint-Honoré, Paris','Paiement pour service de conseil en gestion',21),(28,'2022-08-23',7000,'16 rue du Temple, Strasbourg','Paiement pour service de maintenance',77),(29,'2022-05-01',10856,'Rue de la Paix, Paris','Facture pour prestation de team building',26),(30,'2022-05-01',10856,'Rue de la Paix, Paris','Facture pour prestation de team building',26),(31,'2022-06-18',5600,'Avenue de l\'Opéra, Lyon','Facture pour prestation de team building',83),(32,'2022-08-03',7200,'Boulevard des Lices, Marseille','Facture pour prestation de team building',51),(33,'2022-07-12',4300,'Rue Sainte-Catherine, Bordeaux','Facture pour prestation de team building',98),(34,'2022-06-30',9800,'Rue du Faubourg Saint-Honoré, Paris','Facture pour prestation de team building',17);
/*!40000 ALTER TABLE `facture` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `infos`
--

DROP TABLE IF EXISTS `infos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `infos` (
  `idINFO` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(45) NOT NULL,
  PRIMARY KEY (`idINFO`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `infos`
--

LOCK TABLES `infos` WRITE;
/*!40000 ALTER TABLE `infos` DISABLE KEYS */;
INSERT INTO `infos` VALUES (1,'Sportif'),(2,'Enceinte'),(3,'Diabètique'),(4,'Obèse'),(5,'Handicapé'),(6,'Fumeur'),(7,'Allergies'),(8,'Végétarien');
/*!40000 ALTER TABLE `infos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lieu`
--

DROP TABLE IF EXISTS `lieu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lieu` (
  `idLIEU` int NOT NULL,
  `nom` varchar(45) DEFAULT NULL,
  `nb_salle` int DEFAULT NULL,
  `addresse` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`idLIEU`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lieu`
--

LOCK TABLES `lieu` WRITE;
/*!40000 ALTER TABLE `lieu` DISABLE KEYS */;
INSERT INTO `lieu` VALUES (1,'Salle de conférence de l\'hôtel Hilton',3,'1 Avenue des Champs-Élysées, 75008 Paris'),(2,'Salle de réunion de l\'hôtel Sofitel',1,'25 Rue de la Paix, 75002 Paris'),(3,'Espace de coworking WeWork',2,'33 Rue La Fayette, 75009 Paris'),(4,'Centre de congrès de la Villette',10,'211 Avenue Jean Jaurès, 75019 Paris'),(5,'Salle de réunion de l\'hôtel Novotel',1,'10 Place de Catalogne, 75014 Paris'),(6,'Salle de conférence de l\'hôtel Marriott',5,'17 Boulevard Saint-Jacques, 75014 Paris'),(7,'Espace de réunion de l\'hôtel Mercure',2,'2 Rue Simone Veil, 92400 Courbevoie'),(8,'Centre de congrès de Montpellier',12,'Route de La Foire, 34470 Pérols'),(9,'Salle de réunion de l\'hôtel Ibis',1,'20 Rue Madeleine Vionnet, 93300 Aubervilliers'),(10,'Salle de conférence de l\'hôtel Pullman',6,'65 Avenue de la Grande Armée, 75016 Paris'),(11,'Espace de coworking Station F',5,'55 Boulevard Vincent Auriol, 75013 Paris'),(12,'Salle de conférence de l\'hôtel Hyatt Regency',4,'3 Place du Général Kœnig, 75017 Paris'),(13,'Salle de réunion de l\'hôtel Radisson Blu',1,'2-4 Boulevard de Neuilly, 92200 Neuilly-sur-Seine'),(14,'Centre de congrès de Lyon',15,'50 Quai Charles de Gaulle, 69006 Lyon');
/*!40000 ALTER TABLE `lieu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `list_activite`
--

DROP TABLE IF EXISTS `list_activite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `list_activite` (
  `idlist_activite` int NOT NULL,
  `nom` varchar(45) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `nb_personne_min` int DEFAULT NULL,
  `nb_personne_max` int DEFAULT NULL,
  `prix_min` int DEFAULT NULL,
  `prix_max` int DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`idlist_activite`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `list_activite`
--

LOCK TABLES `list_activite` WRITE;
/*!40000 ALTER TABLE `list_activite` DISABLE KEYS */;
INSERT INTO `list_activite` VALUES (1,'Paintball','Jeu de stratégie en équipe',5,20,150,500,'default.png'),(2,'Escape Game','Jeu d\'évasion grandeur nature',4,10,200,800,'default.png'),(3,'Randonnée en montagne','Découverte de la nature et des paysages montagneux',2,15,100,400,'default.png'),(4,'Dégustation de vin','Initiation à la dégustation de vins locaux',4,20,50,300,'default.png'),(5,'Parcours accrobranche','Parcours aériens dans les arbres',3,10,120,400,'default.png'),(6,'Cours de cuisine','Atelier culinaire pour apprendre de nouvelles recettes',2,8,80,250,'default.png'),(7,'Chasse au trésor','Jeu de piste en équipe pour retrouver le trésor caché',5,20,150,500,'default.png'),(8,'Karting','Course de karting sur circuit',4,12,100,400,'default.png'),(9,'Tournoi de pétanque','Tournoi amical de pétanque',8,24,50,150,'default.png'),(11,'Visite guidée de la ville','Découverte de la ville et de ses monuments',2,15,30,150,'default.png'),(12,'Initiation à l\'équitation','Cours d\'équitation pour débutants',1,5,50,200,'default.png'),(13,'Balade en bateau','Promenade en mer pour admirer le paysage',2,8,80,250,'default.png'),(14,'Initiation au golf','Cours d\'initiation au golf',2,6,100,300,'default.png'),(15,'Séance de yoga','Séance de yoga en plein air',1,10,20,80,'default.png'),(16,'Atelier de poterie','Atelier de création de poterie',2,8,50,200,'default.png'),(17,'Initiation à la plongée','Cours d\'initiation à la plongée sous-marine',2,4,150,500,'default.png'),(18,'Séance de méditation','Séance de méditation en pleine nature',1,10,30,100,'default.png'),(19,'Tournoi de beach-volley','Tournoi amical de beach-volley',6,12,100,300,'default.png'),(20,'Initiation à la dégustation de chocolat','Atelier de dégustation de chocolat',2,10,500,1250,'default.png'),(21,'Atelier de création de cocktails','Atelier de création de cocktails',4,8,60,200,'default.png'),(22,'Rallye pédestre','Rallye pédestre dans la ville',4,20,80,300,'default.png'),(23,'Atelier de cuisine italienne','Atelier de cuisine italienne',2,6,50,200,'default.png'),(24,'Tour en montgolfière','Tour en montgolfière pour survoler la région',2,4,300,1000,'default.png'),(25,'Séance de sophrologie','Séance de sophrologie pour se détendre',1,15,40,120,'default.png'),(26,'Atelier de création de bijoux','Atelier de création de bijoux fantaisie',2,8,50,150,'default.png'),(27,'Atelier de création de parfum','Atelier de création de parfum',2,6,70,250,'default.png'),(28,'Escape game en extérieur','Escape game en extérieur',4,10,100,400,'default.png'),(29,'Séance de stretching','Séance de stretching pour se relaxer',1,20,20,80,'default.png'),(30,'Atelier de peinture','Atelier de peinture pour les débutants',2,10,50,150,'default.png');
/*!40000 ALTER TABLE `list_activite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `list_option`
--

DROP TABLE IF EXISTS `list_option`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `list_option` (
  `idlist_option` int NOT NULL,
  `idlist_activite` int NOT NULL,
  `nom` varchar(45) DEFAULT NULL,
  `prix` int DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`idlist_option`,`idlist_activite`),
  KEY `fk_list_option_list_activite1_idx` (`idlist_activite`),
  CONSTRAINT `fk_list_option_list_activite1` FOREIGN KEY (`idlist_activite`) REFERENCES `list_activite` (`idlist_activite`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `list_option`
--

LOCK TABLES `list_option` WRITE;
/*!40000 ALTER TABLE `list_option` DISABLE KEYS */;
INSERT INTO `list_option` VALUES (1,1,'Protection',79,'Cette option comprend un gilet de protection, un masque et des protège-mains pour votre sécurité pendant la partie de paintball.'),(1,2,'Indice supplémentaire',49,'Vous êtes bloqués sur une énigme ? Cet indice supplémentaire vous aidera à avancer.'),(1,3,'Guide de montagne',130,'Découvrez les plus beaux sentiers de randonnée avec un guide de montagne professionnel.'),(2,1,'Grenades fumigènes',220,'Ajoutez un effet de surprise à vos parties avec ces grenades fumigènes colorées.'),(2,2,'Extension de temps',120,'Vous êtes sur le point de résoudre l\'énigme, mais le temps est presque écoulé ? Ajoutez 30 minutes supplémentaires.'),(2,4,'Dégustation de vin prestige',390,'Dégustez les meilleurs crus de la région avec cette dégustation de vin prestige.'),(3,1,'Pistolet amélioré',78,'Passez au niveau supérieur avec ce pistolet amélioré avec une meilleure portée et une plus grande précision.'),(3,5,'Parcours extrême',75,'Testez votre courage avec ce parcours accrobranche extrême réservé aux plus téméraires.'),(4,1,'Tenue camouflage',234,'Immergez-vous dans le jeu avec cette tenue de camouflage pour vous fondre dans le décor.'),(4,2,'Mise en scène personnalisée',550,'Rendez l\'expérience encore plus immersive avec une mise en scène personnalisée pour votre groupe.'),(4,6,'Cours de cuisine moléculaire',289,'Découvrez les secrets de la cuisine moléculaire avec ce cours de cuisine original et surprenant.'),(5,1,'Pack sniper',387,'Devenez le roi du paintball avec ce pack comprenant un fusil sniper de précision et des billes supplémentaires.'),(5,2,'Défi supplémentaire',199,'Vous avez réussi l\'Escape Game en un temps record ? Pimentez l\'expérience avec un défi supplémentaire.'),(5,7,'Défi photo',59,'Partez à la chasse aux trésors avec ce défi photo qui vous fera découvrir la ville sous un nouvel angle.'),(6,8,'Course de karting en nocturne',199,'Vivez une expérience de pilotage unique avec cette course de karting en nocturne.'),(7,9,'Matériel de pétanque haut de gamme',120,'Faites la différence avec ce matériel de pétanque haut de gamme pour un tournoi réussi.'),(8,11,'Visite privée',299,'Profitez d\'une visite guidée de la ville en petit comité avec une visite privée.'),(9,3,'Baptême en parapente',179,'Envolez-vous au-dessus des montagnes avec ce baptême en parapente inoubliable.'),(10,4,'Découverte des caves à vin',49,'Partez à la découverte des caves à vin de la région avec cette visite guidée.'),(11,5,'Parcours enfant',25,'Initiez vos enfants à l\'accrobranche avec ce parcours spécialement conçu pour eux.'),(12,6,'Cours de cuisine italienne',199,'Apprenez à cuisiner comme un chef italien avec ce cours de cuisine spécialité italienne.'),(13,2,'Escape game en équipe',89,'Plongez dans une aventure immersive avec cet escape game en équipe.'),(14,8,'Tournoi de karting',349,'Affrontez vos amis avec ce tournoi de karting où la vitesse et l\'adresse seront de mise.'),(15,9,'Pétanque en duo',60,'Jouez en duo avec cette partie de pétanque pour une ambiance conviviale et détendue.'),(16,11,'Visite insolite',79,'Découvrez la ville autrement avec cette visite insolite qui vous fera voir les choses sous un angle original.'),(17,3,'Balade en raquettes',45,'Partez en balade en raquettes et découvrez des paysages enneigés à couper le souffle.'),(18,4,'Découverte des vignobles locaux',79,'Découvrez les vignobles de la région avec cette visite guidée.'),(19,5,'Parcours sportif',35,'Dépassez-vous avec ce parcours accrobranche réservé aux sportifs.'),(20,6,'Atelier cuisine du monde',129,'Voyagez en cuisine avec cet atelier qui vous fera découvrir des saveurs venues d\'ailleurs.'),(21,2,'Escape game en solo',59,'Testez votre esprit de déduction avec cet escape game en solo.'),(22,8,'Karting en groupe',299,'Faites la course entre amis avec ce karting en groupe.'),(23,9,'Tournoi de pétanque en équipe',120,'Affrontez d\'autres équipes lors de ce tournoi de pétanque en équipe.'),(24,11,'Visite nocturne',89,'Découvrez la ville sous un nouveau jour avec cette visite nocturne.'),(25,12,'Initiation à l\'équitation',65,'Découvrez les joies de l\'équitation lors de cette initiation'),(26,13,'Balade en bateau',99,'Profitez d\'une balade en bateau pour admirer la côte'),(27,14,'Initiation au golf',89,'Essayez-vous au golf avec cette initiation'),(28,15,'Séance de yoga',25,'Détendez-vous avec cette séance de yoga'),(29,16,'Atelier de poterie',79,'Créez vos propres oeuvres lors de cet atelier de poterie'),(30,17,'Initiation à la plongée',129,'Explorez les fonds marins lors de cette initiation à la plongée'),(31,18,'Séance de méditation',19,'Recentrez-vous grâce à cette séance de méditation'),(32,19,'Découverte de la nature',45,'Partez à la découverte de la nature lors de cette activité'),(33,20,'Cours de surf',79,'Initiez-vous au surf lors de ce cours'),(34,21,'Balade à cheval',59,'Profitez d\'une balade à cheval pour découvrir la région'),(35,2,'Escape Game',119,'Tentez de résoudre les énigmes pour sortir de la pièce lors de cet escape game'),(36,22,'Cours de danse',35,'Apprenez les bases de la danse lors de ce cours'),(37,23,'Atelier de création de bijoux',69,'Créez vos propres bijoux lors de cet atelier'),(38,24,'Visite de châteaux',89,'Découvrez l\'histoire des châteaux de la région lors de cette visite'),(39,25,'Séance de tir à l\'arc',49,'Essayez-vous au tir à l\'arc lors de cette séance'),(40,26,'Initiation à l\'escalade',79,'Initiez-vous à l\'escalade lors de cette séance d\'initiation'),(41,27,'Séance de sophrologie',25,'Détendez-vous et apprenez à gérer votre stress lors de cette séance de sophrologie'),(42,28,'Cours de piano',45,'Apprenez à jouer du piano lors de ce cours'),(43,29,'Atelier de dégustation de fromages',59,'Découvrez et dégustez différents fromages lors de cet atelier'),(44,30,'Atelier de création de parfums',89,'Créez votre propre parfum lors de cet atelier');
/*!40000 ALTER TABLE `list_option` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `loue`
--

DROP TABLE IF EXISTS `loue`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `loue` (
  `idTEAM_BUILDING` int NOT NULL,
  `idMATERIEL` int NOT NULL,
  `prix_a_payer` int DEFAULT NULL,
  `date_location` date DEFAULT NULL,
  `date_rendu` date DEFAULT NULL,
  PRIMARY KEY (`idTEAM_BUILDING`,`idMATERIEL`),
  KEY `fk_TEAM_BUILDING_has_MATERIEL_MATERIEL1` (`idMATERIEL`),
  CONSTRAINT `fk_TEAM_BUILDING_has_MATERIEL_MATERIEL1` FOREIGN KEY (`idMATERIEL`) REFERENCES `materiel` (`idMATERIEL`),
  CONSTRAINT `fk_TEAM_BUILDING_has_MATERIEL_TEAM_BUILDING1` FOREIGN KEY (`idTEAM_BUILDING`) REFERENCES `team_building` (`idTEAM_BUILDING`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `loue`
--

LOCK TABLES `loue` WRITE;
/*!40000 ALTER TABLE `loue` DISABLE KEYS */;
INSERT INTO `loue` VALUES (1,1,200,'2023-04-15','2023-04-17'),(2,2,100,'2023-04-15','2023-04-16'),(2,5,150,'2023-05-01','2023-05-03'),(3,3,150,'2023-04-15','2023-04-16'),(3,5,80,'2023-06-10','2023-06-12'),(4,4,200,'2023-04-15','2023-04-16'),(5,5,250,'2023-04-15','2023-04-16'),(5,14,150,'2023-09-15','2023-09-16'),(7,25,150,'2023-06-10','2023-06-12'),(8,13,100,'2023-07-02','2023-07-03'),(9,49,50,'2023-09-05','2023-09-06'),(12,23,100,'2023-06-15','2023-06-17'),(20,12,100,'2023-05-01','2023-05-02'),(20,44,75,'2023-07-01','2023-07-03'),(21,20,150,'2023-08-23','2023-08-25'),(30,15,50,'2023-08-15','2023-08-17'),(31,18,120,'2023-08-15','2023-08-16'),(33,45,300,'2023-09-12','2023-09-15'),(44,37,200,'2023-07-20','2023-07-23');
/*!40000 ALTER TABLE `loue` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `materiel`
--

DROP TABLE IF EXISTS `materiel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `materiel` (
  `idMATERIEL` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(45) DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `prix` int DEFAULT NULL,
  `quantite_disponible` int DEFAULT NULL,
  PRIMARY KEY (`idMATERIEL`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `materiel`
--

LOCK TABLES `materiel` WRITE;
/*!40000 ALTER TABLE `materiel` DISABLE KEYS */;
INSERT INTO `materiel` VALUES (1,'Ballons','Lot de 10 ballons pour les jeux en extérieur',15,50),(2,'Ballon de foot','Ballon officiel de la FFF',25,20),(3,'Chasuble','Chasuble de couleur bleue',15,30),(4,'Chronomètre','Chronomètre électronique de précision',40,10),(5,'Cône','Cône d\'entraînement en plastique',2,100),(6,'Sifflet','Sifflet d\'arbitre professionnel',5,50),(7,'Ballon de football','Ballon officiel en cuir pour une pratique du football de qualité',60,30),(8,'Raquettes de tennis','Raquettes légères et maniables pour jouer au tennis',100,20),(9,'Vélos tout-terrain','Vélos tout-terrain pour des randonnées sportives en pleine nature',250,15),(10,'Ballon de football','Ballon officiel en cuir pour une pratique du football de qualité',60,30),(11,'Raquettes de tennis','Raquettes légères et maniables pour jouer au tennis',100,20),(12,'Vélos tout-terrain','Vélos tout-terrain pour des randonnées sportives en pleine nature',250,15),(13,'Chasubles de football','Chasubles pour différencier les équipes lors d\'un match de football',15,50),(14,'Tee-shirts de sport','Tee-shirts en coton respirant pour une pratique sportive confortable',25,100),(15,'Ballon de foot','Ballon officiel de taille 5',20,30),(16,'Raquette de tennis','Raquette en graphite légère et résistante',80,25),(17,'Arc et flèches','Arc en bois et flèches en carbone',120,15),(18,'Tente de camping','Tente pour 2 personnes, imperméable',100,10),(19,'Casque de vélo','Casque de vélo de route en polycarbonate',50,40),(20,'Ballons','Lot de 10 ballons gonflables',15,50),(21,'Cônes','Lot de 20 cônes pour parcours d\'obstacles',30,30),(22,'Corde','Corde en chanvre de 10m pour activités en équipe',45,20),(23,'Sacs à dos','Lot de 15 sacs à dos pour randonnées',200,15),(24,'Talkie-walkies','Pack de 10 talkie-walkies pour communication en équipe',150,25),(25,'Ballons de football','Ensemble de ballons de football de taille standard',50,20),(26,'Ballons de basket-ball','Ensemble de ballons de basket-ball de taille standard',60,15),(27,'Cônes de signalisation','Ensemble de cônes de signalisation pour diverses activités de plein air',30,25),(28,'Sacs à dos','Sacs à dos de grande qualité pour les randonnées et les excursions',80,10),(29,'Lampes de poche','Lampes de poche puissantes pour les activités de plein air',25,30),(30,'Plaque de cuisson','Plaque de cuisson électrique',120,5),(31,'Robot culinaire','Robot multifonction pour la cuisine',250,3),(32,'Mixeur','Mixeur à main électrique',50,10),(33,'Couteau de chef','Couteau de cuisine polyvalent',80,7),(34,'Balance de cuisine','Balance électronique de précision',30,15),(35,'Couteau de chef','Couteau de cuisine indispensable pour toutes les tâches de découpe',40,10),(36,'Planche à découper','Planche en bois résistante pour découper les aliments',20,20),(37,'Poêle en acier inoxydable','Poêle pour cuire les aliments de façon uniforme',50,5),(38,'Casserole en cuivre','Casserole pour chauffer des liquides ou des sauces',80,8),(39,'Râpe à fromage','Râpe en acier inoxydable pour râper les fromages',30,15),(40,'Poêle','Poêle en inox',20,10),(41,'Casserole','Casserole en aluminium',15,20),(42,'Mixeur','Mixeur plongeant',30,5),(43,'Robot culinaire','Robot de cuisine multifonction',100,3),(44,'Moule à tarte','Moule à tarte en silicone',8,15),(45,'Couteau de chef','Couteau polyvalent pour couper les légumes et les viandes',50,20),(46,'Planche à découper','Planche en bois pour découper les ingrédients',25,30),(47,'Moule à tarte','Moule en métal pour la préparation des tartes',15,15),(48,'Poêle antiadhésive','Poêle pour la cuisson des aliments sans accrocher',35,25),(49,'Robot culinaire','Robot multifonction pour la préparation des aliments',200,10);
/*!40000 ALTER TABLE `materiel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `option`
--

DROP TABLE IF EXISTS `option`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `option` (
  `idlist_option` int NOT NULL,
  `idlist_activite` int NOT NULL,
  `idActivite` int NOT NULL,
  PRIMARY KEY (`idlist_option`,`idlist_activite`,`idActivite`),
  KEY `fk_list_option_has_activite_activite1_idx` (`idActivite`),
  KEY `fk_list_option_has_activite_list_option1_idx` (`idlist_option`,`idlist_activite`),
  CONSTRAINT `fk_list_option_has_activite_activite1` FOREIGN KEY (`idActivite`) REFERENCES `activite` (`idActivite`),
  CONSTRAINT `fk_list_option_has_activite_list_option1` FOREIGN KEY (`idlist_option`, `idlist_activite`) REFERENCES `list_option` (`idlist_option`, `idlist_activite`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `option`
--

LOCK TABLES `option` WRITE;
/*!40000 ALTER TABLE `option` DISABLE KEYS */;
INSERT INTO `option` VALUES (1,1,1),(33,20,2),(15,9,3),(23,9,3),(29,16,4),(30,17,5),(1,2,6),(2,2,6),(4,2,6),(13,2,6),(6,8,8),(22,8,8),(27,14,9),(4,6,10),(12,6,10),(20,6,10);
/*!40000 ALTER TABLE `option` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `participant`
--

DROP TABLE IF EXISTS `participant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `participant` (
  `idPARTICIPANT` int NOT NULL AUTO_INCREMENT,
  `token` varchar(64) DEFAULT NULL,
  `nom` varchar(45) DEFAULT NULL,
  `prenom` varchar(45) DEFAULT NULL,
  `telephone` char(10) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `sexe` char(1) DEFAULT NULL,
  `idCLIENT` int NOT NULL,
  PRIMARY KEY (`idPARTICIPANT`,`idCLIENT`),
  KEY `fk_PARTICIPANT_CLIENT1` (`idCLIENT`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `participant`
--

LOCK TABLES `participant` WRITE;
/*!40000 ALTER TABLE `participant` DISABLE KEYS */;
INSERT INTO `participant` VALUES (1,'111','Pitt','Brad','0612345678',37,'M',1),(1,'123','Dupont','Paul','0612345678',25,'H',2),(2,'217','Martin','Sophie','0678901234',32,'F',1),(2,'286','Roberts','Julia','0678901234',28,'F',8),(3,'345','Leblanc','Lucie','0698765432',28,'F',4),(3,'397','DiCaprio','Leonardo','0612345678',24,'M',9),(4,'439','Garcia','Antonio','0632109876',21,'H',3),(4,'507','Streep','Meryl','0678901234',42,'F',10),(5,'551','Rivière','Alexandre','0687654321',35,'H',5),(5,'618','Hanks','Tom','0612345678',29,'M',11),(6,'723','Winslet','Kate','0678901234',36,'F',12),(6,'756','Da Vinci','Leonardo','0612345678',35,'M',15),(7,'839','Depp','Johnny','0612345678',33,'M',13),(7,'938','Curie','Marie','0678453129',31,'F',23),(8,'949','Johansson','Scarlett','0678901234',27,'F',14),(8,'1338','Hawking','Stephen','0698765432',42,'M',52),(9,'1060','Cruise','Tom','0612345678',31,'M',15),(9,'1709','Roosevelt','Theodore','0676543210',45,'M',78),(10,'1170','Stone','Emma','0678901234',25,'F',16),(10,'2082','Keller','Helen','0611111111',27,'F',104),(11,'1202','Parker','Peter','0123456789',20,'M',10),(11,'1276','Ford','Harrison','0612345678',45,'M',17),(12,'1345','Simpson','Bart','0123456789',16,'M',14),(12,'1386','Lawrence','Jennifer','0678901234',29,'F',18),(13,'1497','Hathaway','Anne','0612345678',32,'F',19),(13,'1508','Everdeen','Katniss','0123456789',18,'F',20),(14,'1607','Bale','Christian','0678901234',40,'M',20),(14,'1651','Potter','Harry','0123456789',17,'M',24),(15,'1718','Portman','Natalie','0612345678',34,'F',21),(15,'1814','Stark','Arya','0123456789',16,'F',30),(16,'1828','McConaughey','Matthew','0678901234',38,'M',22),(17,'1939','Bullock','Sandra','0612345678',39,'F',23),(18,'2049','Clooney','George','0678901234',52,'M',24),(19,'2160','Law','Jude','0612345678',41,'M',25),(20,'2270','Jolie','Angelina','0678901234',36,'F',26),(38,NULL,'Zuck','Marc','0651422245',27,'M',35);
/*!40000 ALTER TABLE `participant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `participant_equipe`
--

DROP TABLE IF EXISTS `participant_equipe`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `participant_equipe` (
  `idPARTICIPANT` int NOT NULL,
  `idCLIENT` int NOT NULL,
  `idEQUIPE` int NOT NULL,
  `idTEAM_BUILDING` int NOT NULL,
  PRIMARY KEY (`idPARTICIPANT`,`idCLIENT`,`idEQUIPE`),
  KEY `idEQUIPE` (`idEQUIPE`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `participant_equipe`
--

LOCK TABLES `participant_equipe` WRITE;
/*!40000 ALTER TABLE `participant_equipe` DISABLE KEYS */;
INSERT INTO `participant_equipe` VALUES (1,1,1,1),(1,2,1,1),(2,1,1,1);
/*!40000 ALTER TABLE `participant_equipe` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `participant_info`
--

DROP TABLE IF EXISTS `participant_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `participant_info` (
  `idPARTICIPANT` int NOT NULL,
  `idCLIENT` int NOT NULL,
  `idINFO` int NOT NULL,
  PRIMARY KEY (`idPARTICIPANT`,`idCLIENT`,`idINFO`),
  KEY `idCLIENT` (`idCLIENT`),
  KEY `idINFO` (`idINFO`),
  CONSTRAINT `participant_info_ibfk_1` FOREIGN KEY (`idPARTICIPANT`) REFERENCES `participant` (`idPARTICIPANT`),
  CONSTRAINT `participant_info_ibfk_2` FOREIGN KEY (`idCLIENT`) REFERENCES `participant` (`idCLIENT`),
  CONSTRAINT `participant_info_ibfk_3` FOREIGN KEY (`idINFO`) REFERENCES `infos` (`idINFO`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `participant_info`
--

LOCK TABLES `participant_info` WRITE;
/*!40000 ALTER TABLE `participant_info` DISABLE KEYS */;
INSERT INTO `participant_info` VALUES (1,1,1),(1,1,2);
/*!40000 ALTER TABLE `participant_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `present`
--

DROP TABLE IF EXISTS `present`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `present` (
  `idPARTICIPANT` int NOT NULL,
  `idTEAM_BUILDING` int NOT NULL,
  `present` int DEFAULT NULL,
  `motif_si_non` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`idPARTICIPANT`,`idTEAM_BUILDING`),
  KEY `fk_PARTICIPANT_has_TEAM_BUILDING_TEAM_BUILDING1` (`idTEAM_BUILDING`),
  CONSTRAINT `fk_PARTICIPANT_has_TEAM_BUILDING_PARTICIPANT1` FOREIGN KEY (`idPARTICIPANT`) REFERENCES `participant` (`idPARTICIPANT`),
  CONSTRAINT `fk_PARTICIPANT_has_TEAM_BUILDING_TEAM_BUILDING1` FOREIGN KEY (`idTEAM_BUILDING`) REFERENCES `team_building` (`idTEAM_BUILDING`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `present`
--

LOCK TABLES `present` WRITE;
/*!40000 ALTER TABLE `present` DISABLE KEYS */;
INSERT INTO `present` VALUES (1,1,0,'malade'),(1,2,1,NULL),(1,4,1,NULL),(1,5,1,NULL),(2,1,1,NULL),(2,2,0,'autre engagement'),(2,5,0,'problème personnel'),(2,6,1,NULL),(3,2,1,NULL),(3,3,0,'retard'),(3,6,1,NULL),(3,7,1,NULL),(4,2,1,NULL),(4,3,1,NULL),(4,7,0,'congés'),(4,8,1,NULL),(5,3,1,NULL),(5,4,0,'fatigue'),(5,8,1,NULL),(5,9,0,'urgence familiale'),(6,9,1,NULL),(6,10,1,NULL),(7,10,1,NULL),(11,8,1,NULL);
/*!40000 ALTER TABLE `present` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `prestataire`
--

DROP TABLE IF EXISTS `prestataire`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `prestataire` (
  `idPRESTATAIRE` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(45) DEFAULT NULL,
  `prenom` varchar(45) DEFAULT NULL,
  `tel` char(10) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `metier` varchar(45) DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`idPRESTATAIRE`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prestataire`
--

LOCK TABLES `prestataire` WRITE;
/*!40000 ALTER TABLE `prestataire` DISABLE KEYS */;
INSERT INTO `prestataire` VALUES (1,'Durand','Jean','0601020304','jean.durand@gmail.com','1','Guide de montagne','Je suis un expert en randonnée en montagne et j\'adore partager ma passion avec les autres.'),(2,'Lefebvre','Sophie','0612345678','sophie.lefebvre@hotmail.com','2','Instructeur de pole dance','Je suis une danseuse de pole dance professionnelle et j\'enseigne cette discipline depuis plus de 10 ans.'),(3,'Bouchard','Marc','0607080910','marc.bouchard@yahoo.fr','3','Dresseur de serpents','Je suis un expert en reptiles et j\'ai une grande expérience dans la dressage de serpents.'),(4,'Petit','Elodie','0601020304','elodie.petit@gmail.com','4','Professeur de yodel','Je suis passionnée de musique folklorique et je suis spécialisée dans le yodel. Je propose des cours individuels et collectifs pour tous les niveaux.'),(5,'Dubois','Luc','0612345678','luc.dubois@hotmail.com','5','Magicien','Je suis un magicien professionnel et j\'ai participé à de nombreux spectacles à travers le monde.'),(6,'Moreau','Julie','0607080910','julie.moreau@yahoo.fr','6','Sculpteur de ballons','Je suis une artiste spécialisée dans la sculpture de ballons et je propose mes services pour des fêtes et des événements.'),(7,'Martin','Philippe','0601020304','philippe.martin@gmail.com','7','Professeur de jonglage','Je suis un expert en jonglage et j\'ai enseigné cette discipline pendant plusieurs années dans une école de cirque.'),(8,'Roux','Mireille','0612345678','mireille.roux@hotmail.com','8','Expert en déguisements','Je suis une créatrice de costumes professionnelle et j\'ai réalisé des déguisements pour des films et des théâtres.'),(9,'Girard','Franck','0607080910','franck.girard@yahoo.fr','9','Cuisinier exotique','Je suis un chef cuisinier spécialisé dans la cuisine exotique et j\'ai travaillé dans plusieurs restaurants renommés.'),(10,'Pitt','Brad','0123456789','brad.pitt@example.com','10','Organisateur de soirées','Spécialisé dans les soirées à thème'),(11,'Jolie','Angelina','0123456789','angelina.jolie@example.com','11','Spécialiste en décoration','Création d\'ambiances originales et atypiques'),(12,'Clooney','George','0123456789','george.clooney@example.com','12','Chef cuisinier','Cuisine inventive et raffinée'),(13,'Depp','Johnny','0123456789','johnny.depp@example.com','13','Animateur','Animations ludiques et originales'),(14,'Roberts','Julia','0123456789','julia.roberts@example.com','14','Coach sportif','Activités sportives variées et adaptées à tous les niveaux'),(15,'Hanks','Tom','0123456789','tom.hanks@example.com','15','Musicien','Ambiance musicale personnalisée et conviviale'),(16,'Winslet','Kate','0123456789','kate.winslet@example.com','16','Spécialiste en communication','Animation de groupes de travail et de formations'),(17,'Damon','Matt','0123456789','matt.damon@example.com','17','Photographe','Captation de moments uniques et originaux'),(18,'Lawrence','Jennifer','0123456789','jennifer.lawrence@example.com','18','Organisateur d\'activités en extérieur','Activités en pleine nature et sensations fortes'),(19,'Smith','Will','0123456789','will.smith@example.com','19','Artiste de spectacle','Spectacles de magie et d\'humour pour tous les publics'),(20,'Macron','Emmanuel','0612345678','emacron@gmail.com','20','Orateur','Ancien président de la République française'),(21,'Merkel','Angela','0498765432','amerkel@gmail.com','21','Chancelière','Ancienne chancelière allemande'),(22,'Obama','Barack','0123456789','bobama@gmail.com','22','Orateur','Ancien président des États-Unis'),(23,'Trudeau','Justin','0611111111','jtrudeau@gmail.com','23','Premier ministre','Premier ministre du Canada'),(24,'Mitterrand','François','0698765432','fmitterrand@gmail.com','24','Homme d\'État','Ancien président de la République française'),(25,'Thatcher','Margaret','0688888888','mthatcher@gmail.com','25','Premier ministre','Ancienne première ministre britannique'),(26,'Trump','Donald','0101010101','dtrump@gmail.com','26','Homme d\'affaires','Ancien président des États-Unis'),(27,'Mao','Zedong','0865432109','mzedong@gmail.com','27','Homme d\'État','Ancien dirigeant de la Chine'),(28,'Thiers','Adolphe','0644444444','athiers@gmail.com','28','Homme d\'État','Ancien président de la République française'),(29,'De Gaulle','Charles','0622222222','cdegaulle@gmail.com','29','Homme d\'État','Ancien président de la République française'),(30,'Robuchon','Joël','0123456789','joel.robuchon@gmail.com','30','Chef étoilé','Joël Robuchon, surnommé le \"Chef des Chefs\", a été récompensé de 31 étoiles au guide Michelin.'),(31,'Ducasse','Alain','0234567890','alain.ducasse@hotmail.com','31','Chef étoilé','Alain Ducasse, élu Meilleur Chef du Monde en 1987, a obtenu un total de 21 étoiles au guide Michelin.'),(33,'Guérard','Michel','0456789012','michel.guerard@gmail.com','33','Chef étoilé','Michel Guérard, pionnier de la cuisine minceur et fondateur de la cuisine santé, a obtenu un total de 8 étoiles au guide Michelin.'),(34,'Passard','Alain','0567890123','alain.passard@hotmail.com','34','Chef étoilé','Alain Passard, maître de la cuisine végétarienne et du jardinage, a obtenu un total de 11 étoiles au guide Michelin.'),(35,'Savoy','Guy','0678901234','guy.savoy@yahoo.com','35','Chef étoilé','Guy Savoy, chef emblématique de la gastronomie française, a obtenu un total de 12 étoiles au guide Michelin.'),(36,'Robuchon','Joël','0123456789','joel.robuchon@gmail.com','36','Chef étoilé','Joël Robuchon, surnommé le \"Chef des Chefs\", a été récompensé de 31 étoiles au guide Michelin.'),(37,'Ducasse','Alain','0234567890','alain.ducasse@hotmail.com','37','Chef étoilé','Alain Ducasse, élu Meilleur Chef du Monde en 1987, a obtenu un total de 21 étoiles au guide Michelin.'),(38,'Bocuse','Paul','0345678901','paul.bocuse@yahoo.com','38','Chef étoilé','Paul Bocuse, fondateur de la cuisine nouvelle et ambassadeur de la gastronomie française dans le monde entier, a obtenu 20 étoiles au guide Michelin.'),(41,'Savoy','Guy','0678901234','guy.savoy@yahoo.com','41','Chef étoilé','Guy Savoy, chef emblématique de la gastronomie française, a obtenu un total de 12 étoiles au guide Michelin.'),(42,'Gagnaire','Pierre','0789012345','pierre.gagnaire@gmail.com','42','Chef étoilé','Pierre Gagnaire, créateur d\'une cuisine fusion et expérimentale, a obtenu un total de 14 étoiles au guide Michelin.'),(43,'Roca','Joan','0890123456','joan.roca@hotmail.com','43','Chef étoilé','Joan Roca, chef espagnol de renommée internationale, a obtenu un total de 7 étoiles au guide Michelin.'),(46,'Napoléon','Bonaparte','0789564123','napoleon.bonaparte@empire.fr','46','Militaire','Empereur français, stratège militaire et homme d\'Etat'),(49,'Churchill','Winston','0712345678','winston.churchill@uk.gov','49','Politique','Homme d\'Etat britannique, Premier ministre pendant la Seconde Guerre mondiale'),(50,'Merkel','Angela','0687654321','angela.merkel@de.gov','50','Politique','Chancelière allemande de 2005 à 2021'),(51,'Bocuse','Paul','0632145879','paul.bocuse@gastronomie.fr','51','Cuisinier','Chef cuisinier français, fondateur de la nouvelle cuisine'),(52,'Ramsay','Gordon','0612345678','gordon.ramsay@culinary.com','52','Cuisinier','Chef cuisinier britannique, célèbre pour ses émissions de télévision'),(53,'Patton','George','0645896321','george.patton@army.com','53','Militaire','Général américain de la Seconde Guerre mondiale'),(59,'Général Tao','','0635487952','gen.tao@chine.mil','59','Militaire','Général fictif de l\'armée chinoise dans la bande dessinée Tintin'),(60,'Tintin','','0654789321','tintin@reporter.com','60','Bande dessinée','Jeune reporter belge, héros de la série de bande dessinée éponyme'),(61,'Asterix','','0659874123','asterix@gaulle.fr','61','Bande dessinée','Gaulois petit mais courageux, héros de la série de bande dessinée éponyme'),(62,'Captain America','','0687456321','captainamerica@avengers.com','62','Super-héros','Super-soldat américain, membre fondateur des Avengers dans l\'univers Marvel'),(63,'Ickbar','Jorj','0981732574','jorj@solutions.io','strong_pass','Boss','Jorj est le Boss Off'),(64,'Ickbar','Jorj','0981732574','jorj@solutions.io','strong_pass','Boss','Jorj est le Boss Off');
/*!40000 ALTER TABLE `prestataire` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reserve`
--

DROP TABLE IF EXISTS `reserve`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reserve` (
  `idTEAM_BUILDING` int NOT NULL,
  `idSALLE` int NOT NULL,
  `prix_a_payer` int DEFAULT NULL,
  `date_debut` date DEFAULT NULL,
  `date_fin` date DEFAULT NULL,
  PRIMARY KEY (`idTEAM_BUILDING`,`idSALLE`),
  KEY `fk_TEAM_BUILDING_has_SALLE_SALLE1` (`idSALLE`),
  CONSTRAINT `fk_TEAM_BUILDING_has_SALLE_SALLE1` FOREIGN KEY (`idSALLE`) REFERENCES `salle` (`idSALLE`),
  CONSTRAINT `fk_TEAM_BUILDING_has_SALLE_TEAM_BUILDING1` FOREIGN KEY (`idTEAM_BUILDING`) REFERENCES `team_building` (`idTEAM_BUILDING`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reserve`
--

LOCK TABLES `reserve` WRITE;
/*!40000 ALTER TABLE `reserve` DISABLE KEYS */;
/*!40000 ALTER TABLE `reserve` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `salle`
--

DROP TABLE IF EXISTS `salle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `salle` (
  `idSALLE` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(45) DEFAULT NULL,
  `num_salle` int DEFAULT NULL,
  `prix` int DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `disponibilite` int DEFAULT NULL,
  `idLIEU` int NOT NULL,
  PRIMARY KEY (`idSALLE`,`idLIEU`),
  KEY `fk_SALLE_LIEU1` (`idLIEU`),
  CONSTRAINT `fk_SALLE_LIEU1` FOREIGN KEY (`idLIEU`) REFERENCES `lieu` (`idLIEU`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `salle`
--

LOCK TABLES `salle` WRITE;
/*!40000 ALTER TABLE `salle` DISABLE KEYS */;
INSERT INTO `salle` VALUES (1,'Salle de conférence Hilton 1',1,100,'Grande salle de conférence avec équipements modernes',1,1),(2,'Salle de réunion Sofitel 1',1,50,'Salle de réunion cosy et équipée',1,2),(3,'Espace de coworking WeWork 1',1,20,'Espace de travail partagé avec bureaux modernes',1,3),(4,'Centre de congrès de la Villette 1',1,500,'Grand centre de congrès avec nombreuses salles équipées',1,4),(5,'Salle de réunion Novotel 1',1,60,'Salle de réunion avec vue panoramique',1,5),(6,'Salle de conférence Marriott 1',1,80,'Salle de conférence équipée et spacieuse',1,6),(7,'Espace de réunion Mercure 1',1,40,'Espace de réunion moderne et confortable',1,7),(8,'Centre de congrès Montpellier 1',1,300,'Centre de congrès moderne avec salles modulables',1,8),(9,'Salle de réunion Ibis 1',1,30,'Salle de réunion simple et fonctionnelle',1,9),(10,'Salle de conférence Pullman 1',1,150,'Salle de conférence haut de gamme avec équipements dernier cri',1,10),(11,'Espace de coworking Station F 1',1,25,'Espace de travail partagé dans un lieu innovant',1,11),(12,'Salle de conférence Hyatt Regency 1',1,120,'Salle de conférence avec vue panoramique sur la ville',1,12),(13,'Salle de réunion Radisson Blu 1',1,70,'Salle de réunion moderne et élégante',1,13),(14,'Centre de congrès de Lyon 1',1,400,'Grand centre de congrès avec nombreuses salles équipées',1,14),(15,'Salle A',101,200,'Grande salle de conférence avec vidéoprojecteur',1,1),(16,'Salle B',102,150,'Salle de réunion équipée avec tableau blanc',1,2),(17,'Salle C',103,250,'Salle polyvalente avec tables et chaises modulables',1,3),(18,'Salle D',201,300,'Grande salle de conférence avec vue sur la ville',1,4),(19,'Salle E',202,200,'Salle de réunion équipée avec écran plat',1,5),(20,'Salle F',203,350,'Salle de réception spacieuse avec bar',1,6),(21,'Salle G',301,400,'Salle de conférence équipée avec matériel audiovisuel',1,7),(22,'Salle H',302,250,'Salle de réunion lumineuse avec accès internet',1,8),(23,'Salle I',303,500,'Salle de conférence avec scène et rideau de fond',1,9),(24,'Salle J',401,600,'Grande salle de conférence avec sièges confortables',1,10),(25,'Salle K',402,200,'Salle de réunion équipée avec vidéoprojecteur',1,11),(26,'Salle L',403,350,'Salle de réception modulable avec bar et piste de danse',1,12),(27,'Salle M',501,700,'Salle de conférence équipée avec matériel de traduction simultanée',1,13),(28,'Salle N',502,300,'Salle de réunion lumineuse avec accès internet',1,14);
/*!40000 ALTER TABLE `salle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `team_building`
--

DROP TABLE IF EXISTS `team_building`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `team_building` (
  `idTEAM_BUILDING` int NOT NULL AUTO_INCREMENT,
  `idClient` int NOT NULL,
  `type` varchar(45) DEFAULT NULL,
  `titre` varchar(45) DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`idTEAM_BUILDING`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `team_building`
--

LOCK TABLES `team_building` WRITE;
/*!40000 ALTER TABLE `team_building` DISABLE KEYS */;
INSERT INTO `team_building` VALUES (1,60,'Sportif','Course d\'Orientation en Forêt','Travaillez ensemble pour retrouver tous les points de contrôle cachés dans la forêt !'),(2,56,'Créatif','Atelier de Graffiti','Exprimez votre créativité en réalisant une fresque murale collective !'),(3,34,'Culinaire','Atelier de Cuisine du Monde','Découvrez les saveurs des cuisines du monde et préparez un repas complet en équipe !'),(4,57,'Aventure','Raid en Rafting','Affrontez les rapides de la rivière en équipe et vivez des sensations fortes !'),(5,50,'Culturel','Visite de Musée Insolite','Explorez un musée insolite et résolvez des énigmes pour en apprendre davantage sur l\'histoire !'),(6,56,'Musical','Atelier de Percussions','Jouez ensemble pour créer une performance musicale rythmée !'),(7,35,'Nature','Journée de Plantation','Travaillez ensemble pour planter des arbres et créer un environnement plus vert !'),(8,35,'Défi','Chasse au Trésor Urbaine','Parcourez la ville en équipe pour résoudre des énigmes et retrouver le trésor !'),(9,35,'Innovant','Hackathon Créatif','Travaillez ensemble pour créer une innovation technologique et la présenter en équipe !'),(10,34,'Ludique','Escape Game en Equipe','Résolvez des énigmes en équipe pour sortir de la pièce avant la fin du temps imparti !'),(11,56,'Intérieur','Défi culinaire','Une compétition de cuisine dans laquelle les participants doivent utiliser des ingrédients surprenants pour préparer des plats délicieux !'),(12,34,'Extérieur','Course d\'orientation','Un défi sportif dans lequel les équipes doivent trouver leur chemin à travers la forêt en utilisant une carte et une boussole.'),(13,60,'Intérieur','Escape game','Un jeu d\'évasion dans lequel les participants doivent résoudre des énigmes pour sortir d\'une pièce en un temps limité.'),(14,50,'Extérieur','Rallye en 2CV','Une aventure dans laquelle les équipes conduisent des 2CV vintage à travers les routes de campagne pour atteindre des destinations spécifiques.'),(15,58,'Intérieur','Challenge artistique','Une compétition dans laquelle les participants doivent créer des œuvres d\'art uniques en utilisant des matériaux fournis.'),(16,57,'Extérieur','Chasse au trésor','Une aventure amusante dans laquelle les équipes cherchent des indices pour trouver un trésor caché.'),(17,48,'Intérieur','Quizz musical','Un quiz de musique dans lequel les participants doivent deviner les titres et les artistes des chansons jouées.'),(18,55,'Extérieur','Parcours d\'obstacles','Un défi physique dans lequel les équipes doivent surmonter des obstacles tels que des murs d\'escalade, des cordes et des filets.'),(19,34,'Intérieur','Improvisation théâtrale','Un défi créatif dans lequel les participants doivent improviser des scènes de théâtre avec des thèmes donnés.'),(20,50,'Extérieur','Randonnée à cheval','Une expérience équestre dans laquelle les participants font une randonnée à cheval dans la nature.'),(21,34,'Sportif','Olympiades Inter-Entreprises','Compétition sportive en équipe avec différentes épreuves : course de relais, lancer de javelot, saut en longueur, etc.'),(22,34,'Sportif','Olympiades Inter-Entreprises','Compétition sportive en équipe avec différentes épreuves : course de relais, lancer de javelot, saut en longueur, etc.'),(23,50,'Culturel','Chasse au Trésor','Découverte du patrimoine culturel local à travers une série d\'énigmes et de défis.'),(24,35,'Gastronomique','Atelier Cuisine','Cours de cuisine en équipe avec un chef étoilé, suivi d\'un repas dégustation.'),(25,46,'Artistique','Mur de Graffitis Collaboratif','Création collective d\'une œuvre d\'art urbain sur un mur de la ville.'),(26,34,'Aventure','Survie en Milieu Sauvage','Apprentissage des techniques de survie en milieu naturel : construction d\'un abri, allumage d\'un feu, etc.'),(27,35,'Technique','Escape Game High-Tech','Escape game en réalité virtuelle avec des énigmes techniques à résoudre.'),(28,34,'Musical','Composition de Chanson','Composition collective d\'une chanson avec l\'aide d\'un compositeur professionnel.'),(29,35,'Historique','Reconstitution de Bataille','Reconstitution d\'une bataille historique avec costumes, armes et tactiques de l\'époque.'),(30,53,'Cohésion','Course d\'Orientation','Course en équipe avec des épreuves d\'orientation et de coopération.'),(31,60,'Environnement','Nettoyage de Plage','Opération de nettoyage de plage en équipe, suivi d\'un pique-nique zéro déchet. '),(32,35,'Incentive','Cours de cuisine en équipe','Apprenez à cuisiner en équipe en suivant les instructions de notre chef renommé.'),(33,34,'Outdoor','Survie en milieu hostile','Affrontez la nature sauvage en équipe et apprenez à survivre dans des conditions extrêmes.'),(34,49,'Indoor','Escape game en entreprise','Déchiffrez des énigmes en équipe pour sortir de notre escape game unique en son genre.'),(35,34,'Incentive','Cours de cuisine en équipe','Apprenez à cuisiner en équipe en suivant les instructions de notre chef renommé.'),(36,47,'Outdoor','Survie en milieu hostile','Affrontez la nature sauvage en équipe et apprenez à survivre dans des conditions extrêmes.'),(37,52,'Indoor','Escape game en entreprise','Déchiffrez des énigmes en équipe pour sortir de notre escape game unique en son genre.'),(38,55,'Incentive','Atelier d\'improvisation','Apprenez à travailler ensemble en jouant des scènes de théâtre improvisées.'),(39,35,'Outdoor','Chasse au trésor en ville','Explorez la ville en résolvant des énigmes pour découvrir un trésor caché en équipe.'),(40,35,'Indoor','Tournoi de jeux de société','Affrontez-vous en équipe lors d\'un tournoi de jeux de société pour renforcer votre esprit d\'équipe.'),(41,52,'Incentive','Atelier de création de cocktails','Créez vos propres cocktails en équipe sous la supervision de notre mixologiste.'),(42,56,'Outdoor','Randonnée en montagne','Relevez le défi de notre randonnée en montagne en équipe pour renforcer votre cohésion.'),(43,34,'Indoor','Atelier de création de parfums','Créez votre propre parfum en équipe lors de cet atelier animé par notre expert en parfumerie.'),(44,60,'Incentive','Initiation à l\'art de la dégustation','Apprenez à déguster le vin, le fromage ou le chocolat en équipe lors de cette initiation animée par notre expert en dégustation.');
/*!40000 ALTER TABLE `team_building` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teambuilding_activite`
--

DROP TABLE IF EXISTS `teambuilding_activite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `teambuilding_activite` (
  `idTEAM_BUILDING` int NOT NULL,
  `idActivite` int NOT NULL,
  `prix_total` int DEFAULT NULL,
  `date_debut` date DEFAULT NULL,
  `date_fin` date DEFAULT NULL,
  `heure_debut` time DEFAULT NULL,
  `heure_fin` time DEFAULT NULL,
  PRIMARY KEY (`idTEAM_BUILDING`,`idActivite`),
  KEY `fk_TEAM_BUILDING_has_ACTIVITE_ACTIVITE1` (`idActivite`),
  CONSTRAINT `fk_TEAM_BUILDING_has_ACTIVITE_ACTIVITE1` FOREIGN KEY (`idActivite`) REFERENCES `activite` (`idActivite`),
  CONSTRAINT `fk_TEAM_BUILDING_has_ACTIVITE_TEAM_BUILDING1` FOREIGN KEY (`idTEAM_BUILDING`) REFERENCES `team_building` (`idTEAM_BUILDING`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teambuilding_activite`
--

LOCK TABLES `teambuilding_activite` WRITE;
/*!40000 ALTER TABLE `teambuilding_activite` DISABLE KEYS */;
INSERT INTO `teambuilding_activite` VALUES (1,1,2061,'2023-05-01','2023-05-01','09:00:00','12:00:00'),(1,2,2722,'2023-05-01','2023-05-01','14:00:00','18:00:00'),(2,3,1880,'2023-05-15','2023-05-16','10:00:00','16:00:00'),(2,4,1521,'2023-05-15','2023-05-16','10:00:00','16:00:00'),(3,1,2980,'2023-06-01','2023-06-01','09:00:00','12:00:00'),(3,3,500,'2023-06-01','2023-06-01','14:00:00','18:00:00'),(4,2,2722,'2023-06-15','2023-06-16','10:00:00','16:00:00'),(4,4,600,'2023-06-15','2023-06-16','10:00:00','16:00:00'),(5,1,2980,'2023-07-01','2023-07-01','09:00:00','12:00:00'),(5,3,500,'2023-07-01','2023-07-01','14:00:00','18:00:00'),(6,2,79,'2023-07-15','2023-07-16','10:00:00','16:00:00'),(6,4,600,'2023-07-15','2023-07-16','10:00:00','16:00:00'),(7,1,2980,'2023-08-01','2023-08-01','09:00:00','12:00:00'),(7,3,500,'2023-08-01','2023-08-01','14:00:00','18:00:00'),(8,2,79,'2023-08-15','2023-08-16','10:00:00','16:00:00'),(8,4,600,'2023-08-15','2023-08-16','10:00:00','16:00:00');
/*!40000 ALTER TABLE `teambuilding_activite` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `age` int NOT NULL,
  `city` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'John Doe','john@example.com',30,'New York'),(2,'Jane Doe','jane@example.com',25,'San Francisco'),(3,'Bob Smith','bob@example.com',40,'London');
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

-- Dump completed on 2023-05-05  7:51:57
