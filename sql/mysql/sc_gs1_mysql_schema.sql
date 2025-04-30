/*!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19  Distrib 10.6.18-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: scfgs1db1
-- ------------------------------------------------------
-- Server version	10.6.18-MariaDB-0ubuntu0.22.04.1

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
-- Table structure for table `additional_party_identifications`
--

DROP TABLE IF EXISTS `additional_party_identifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `additional_party_identifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `additional_party_identification` varchar(80) DEFAULT '',
  `additional_party_identification_type_code` varchar(80) DEFAULT '',
  `code_list_version` varchar(80) DEFAULT '',
  `gln` varchar(80) DEFAULT '',
  `transactional_party_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `additional_party_identifications`
--

LOCK TABLES `additional_party_identifications` WRITE;
/*!40000 ALTER TABLE `additional_party_identifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `additional_party_identifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `additional_trade_item_identifications`
--

DROP TABLE IF EXISTS `additional_trade_item_identifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `additional_trade_item_identifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gtin` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `additional_trade_item_identifications`
--

LOCK TABLES `additional_trade_item_identifications` WRITE;
/*!40000 ALTER TABLE `additional_trade_item_identifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `additional_trade_item_identifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `addresses`
--

DROP TABLE IF EXISTS `addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `addresses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `city` varchar(80) DEFAULT '',
  `city_code` varchar(80) DEFAULT '',
  `country_code` varchar(80) DEFAULT '',
  `county_code` varchar(80) DEFAULT '',
  `cross_street` varchar(80) DEFAULT '',
  `currency_of_party_code` varchar(80) DEFAULT '',
  `language_of_the_party_code` varchar(80) DEFAULT '',
  `name` varchar(80) DEFAULT '',
  `po_box_number` varchar(80) DEFAULT '',
  `postal_code` varchar(80) DEFAULT '',
  `province_code` varchar(80) DEFAULT '',
  `state` varchar(80) DEFAULT '',
  `street_address_one` varchar(80) DEFAULT '',
  `street_address_three` varchar(80) DEFAULT '',
  `street_address_two` varchar(80) DEFAULT '',
  `latitude` varchar(80) DEFAULT '',
  `longitude` varchar(80) DEFAULT '',
  `transactional_party_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `addresses`
--

LOCK TABLES `addresses` WRITE;
/*!40000 ALTER TABLE `addresses` DISABLE KEYS */;
/*!40000 ALTER TABLE `addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `administrative_units`
--

DROP TABLE IF EXISTS `administrative_units`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `administrative_units` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `administrative_unit_type_code` varchar(50) DEFAULT NULL,
  `gln` varchar(50) DEFAULT NULL,
  `internal_administrative_unit_identification` varchar(50) DEFAULT NULL,
  `order_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `administrative_units`
--

LOCK TABLES `administrative_units` WRITE;
/*!40000 ALTER TABLE `administrative_units` DISABLE KEYS */;
/*!40000 ALTER TABLE `administrative_units` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `aggregation_events`
--

DROP TABLE IF EXISTS `aggregation_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `aggregation_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_id` varchar(120) DEFAULT '',
  `event_time_zone_offset` varchar(80) DEFAULT '',
  `certification` varchar(80) DEFAULT '',
  `event_time` datetime DEFAULT current_timestamp(),
  `reason` varchar(80) DEFAULT '',
  `declaration_time` datetime DEFAULT current_timestamp(),
  `parent_id` varchar(80) DEFAULT '',
  `action` varchar(80) DEFAULT '',
  `biz_step` varchar(80) DEFAULT '',
  `disposition` varchar(80) DEFAULT '',
  `read_point` varchar(80) DEFAULT '',
  `biz_location` varchar(80) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `aggregation_events`
--

LOCK TABLES `aggregation_events` WRITE;
/*!40000 ALTER TABLE `aggregation_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `aggregation_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `allowance_charges`
--

DROP TABLE IF EXISTS `allowance_charges`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `allowance_charges` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `allowance_charge_amount` double DEFAULT 0,
  `aca_code_list_version` varchar(35) DEFAULT '',
  `aca_currency_code` varchar(80) DEFAULT '',
  `allowance_charge_percentage` double DEFAULT 0,
  `allowance_charge_type` varchar(20) DEFAULT '',
  `allowance_or_charge_type` varchar(20) DEFAULT '',
  `amount_per_unit` double DEFAULT 0,
  `apu_code_list_version` varchar(35) DEFAULT '',
  `apu_currency_code` varchar(80) DEFAULT '',
  `base_amount` double DEFAULT 0,
  `ba_code_list_version` varchar(35) DEFAULT '',
  `ba_currency_code` varchar(80) DEFAULT '',
  `base_number_of_units` double DEFAULT 0,
  `bnou_code_list_version` varchar(35) DEFAULT '',
  `bnou_measurement_unit_code` varchar(10) DEFAULT NULL,
  `bracket_identifier` varchar(20) DEFAULT '',
  `effective_date_type` varchar(20) DEFAULT '',
  `sequence_number` int(10) unsigned DEFAULT 0,
  `settlement_type` varchar(20) DEFAULT '',
  `special_service_type` varchar(20) DEFAULT '',
  `order_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `allowance_charges`
--

LOCK TABLES `allowance_charges` WRITE;
/*!40000 ALTER TABLE `allowance_charges` DISABLE KEYS */;
/*!40000 ALTER TABLE `allowance_charges` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `animal_identifications`
--

DROP TABLE IF EXISTS `animal_identifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `animal_identifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `animal_identification_number` varchar(80) DEFAULT '',
  `animal_identification_type_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `animal_identifications`
--

LOCK TABLES `animal_identifications` WRITE;
/*!40000 ALTER TABLE `animal_identifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `animal_identifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `association_events`
--

DROP TABLE IF EXISTS `association_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `association_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_id` varchar(120) DEFAULT '',
  `event_time_zone_offset` varchar(80) DEFAULT '',
  `certification` varchar(80) DEFAULT '',
  `event_time` datetime DEFAULT current_timestamp(),
  `reason` varchar(80) DEFAULT '',
  `declaration_time` datetime DEFAULT current_timestamp(),
  `parent_id` varchar(80) DEFAULT '',
  `action` varchar(80) DEFAULT '',
  `biz_step` varchar(80) DEFAULT '',
  `disposition` varchar(80) DEFAULT '',
  `read_point` varchar(80) DEFAULT '',
  `biz_location` varchar(80) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `association_events`
--

LOCK TABLES `association_events` WRITE;
/*!40000 ALTER TABLE `association_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `association_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `biz_transactions`
--

DROP TABLE IF EXISTS `biz_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `biz_transactions` (
  `biz_transaction_type` varchar(80) DEFAULT '',
  `biz_transaction` varchar(80) DEFAULT '',
  `event_id` int(10) unsigned DEFAULT NULL,
  `type_of_event` varchar(80) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `biz_transactions`
--

LOCK TABLES `biz_transactions` WRITE;
/*!40000 ALTER TABLE `biz_transactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `biz_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `communication_channels`
--

DROP TABLE IF EXISTS `communication_channels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `communication_channels` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `communication_channel_code` varchar(80) DEFAULT '',
  `communication_channel_name` varchar(80) DEFAULT '',
  `communication_value` varchar(80) DEFAULT '',
  `contact_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `communication_channels`
--

LOCK TABLES `communication_channels` WRITE;
/*!40000 ALTER TABLE `communication_channels` DISABLE KEYS */;
/*!40000 ALTER TABLE `communication_channels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consumption_report_item_location_informations`
--

DROP TABLE IF EXISTS `consumption_report_item_location_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `consumption_report_item_location_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `total_consumption_amount` double DEFAULT 0,
  `tca_code_list_version` varchar(35) DEFAULT '',
  `tcac_currency_code` varchar(80) DEFAULT '',
  `inventory_location` int(10) unsigned DEFAULT 0,
  `ship_to` int(10) unsigned DEFAULT 0,
  `consumption_report_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consumption_report_item_location_informations`
--

LOCK TABLES `consumption_report_item_location_informations` WRITE;
/*!40000 ALTER TABLE `consumption_report_item_location_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `consumption_report_item_location_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consumption_report_line_items`
--

DROP TABLE IF EXISTS `consumption_report_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `consumption_report_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `consumed_quantity` double DEFAULT 0,
  `cq_measurement_unit_code` varchar(80) DEFAULT '',
  `cq_code_list_version` varchar(35) DEFAULT '',
  `line_item_number` int(10) unsigned DEFAULT 0,
  `net_consumption_amount` double DEFAULT 0,
  `ncac_code_list_version` varchar(35) DEFAULT '',
  `ncac_currency_code` varchar(80) DEFAULT '',
  `net_price` double DEFAULT 0,
  `np_code_list_version` varchar(35) DEFAULT '',
  `np_currency_code` varchar(80) DEFAULT '',
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `plan_bucket_size_code` varchar(35) DEFAULT '',
  `purchase_conditions` int(10) unsigned DEFAULT 0,
  `consumption_report_id` int(10) unsigned DEFAULT 0,
  `consumption_period_begin` datetime DEFAULT current_timestamp(),
  `consumption_period_end` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consumption_report_line_items`
--

LOCK TABLES `consumption_report_line_items` WRITE;
/*!40000 ALTER TABLE `consumption_report_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `consumption_report_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consumption_reports`
--

DROP TABLE IF EXISTS `consumption_reports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `consumption_reports` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `buyer` int(10) unsigned DEFAULT 0,
  `consumption_report_identification` int(10) unsigned DEFAULT 0,
  `seller` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consumption_reports`
--

LOCK TABLES `consumption_reports` WRITE;
/*!40000 ALTER TABLE `consumption_reports` DISABLE KEYS */;
/*!40000 ALTER TABLE `consumption_reports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `contacts`
--

DROP TABLE IF EXISTS `contacts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `contacts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `contact_type_code` varchar(80) DEFAULT '',
  `department_name` varchar(80) DEFAULT '',
  `job_title` varchar(80) DEFAULT '',
  `person_name` varchar(80) DEFAULT '',
  `transactional_party_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contacts`
--

LOCK TABLES `contacts` WRITE;
/*!40000 ALTER TABLE `contacts` DISABLE KEYS */;
/*!40000 ALTER TABLE `contacts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `currency_exchange_rate_informations`
--

DROP TABLE IF EXISTS `currency_exchange_rate_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `currency_exchange_rate_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `currency_conversion_from_code` varchar(80) DEFAULT '',
  `currency_conversion_to_code` varchar(80) DEFAULT '',
  `exchange_rate` double DEFAULT 0,
  `order_id` int(10) unsigned DEFAULT 0,
  `exchange_rate_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `currency_exchange_rate_informations`
--

LOCK TABLES `currency_exchange_rate_informations` WRITE;
/*!40000 ALTER TABLE `currency_exchange_rate_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `currency_exchange_rate_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `debit_credit_advice_line_item_details`
--

DROP TABLE IF EXISTS `debit_credit_advice_line_item_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `debit_credit_advice_line_item_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `aligned_price` double DEFAULT 0,
  `ap_code_list_version` varchar(35) DEFAULT '',
  `ap_currency_code` varchar(80) DEFAULT '',
  `invoiced_price` double DEFAULT 0,
  `ip_code_list_version` varchar(35) DEFAULT '',
  `ip_currency_code` varchar(80) DEFAULT '',
  `quantity` double DEFAULT 0,
  `q_measurement_unit_code` varchar(80) DEFAULT '',
  `q_code_list_version` varchar(35) DEFAULT '',
  `debit_credit_advice_id` int(10) unsigned DEFAULT 0,
  `debit_credit_advice_line_item_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `debit_credit_advice_line_item_details`
--

LOCK TABLES `debit_credit_advice_line_item_details` WRITE;
/*!40000 ALTER TABLE `debit_credit_advice_line_item_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `debit_credit_advice_line_item_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `debit_credit_advice_line_items`
--

DROP TABLE IF EXISTS `debit_credit_advice_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `debit_credit_advice_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `adjustment_amount` double DEFAULT 0,
  `aa_code_list_version` varchar(35) DEFAULT '',
  `aa_currency_code` varchar(80) DEFAULT '',
  `debit_credit_indicator_code` varchar(80) DEFAULT '',
  `financial_adjustment_reason_code` varchar(80) DEFAULT '',
  `line_item_number` int(10) unsigned DEFAULT 0,
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `debit_credit_advice_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `debit_credit_advice_line_items`
--

LOCK TABLES `debit_credit_advice_line_items` WRITE;
/*!40000 ALTER TABLE `debit_credit_advice_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `debit_credit_advice_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `debit_credit_advices`
--

DROP TABLE IF EXISTS `debit_credit_advices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `debit_credit_advices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `debit_credit_indicator_code` varchar(80) DEFAULT '',
  `total_amount` double DEFAULT 0,
  `ta_code_list_version` varchar(35) DEFAULT '',
  `ta_currency_code` varchar(80) DEFAULT '',
  `bill_to` int(10) unsigned DEFAULT 0,
  `buyer` int(10) unsigned DEFAULT 0,
  `carrier` int(10) unsigned DEFAULT 0,
  `debit_credit_advice_identification` int(10) unsigned DEFAULT 0,
  `seller` int(10) unsigned DEFAULT 0,
  `ship_from` int(10) unsigned DEFAULT 0,
  `ship_to` int(10) unsigned DEFAULT 0,
  `ultimate_consignee` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `debit_credit_advices`
--

LOCK TABLES `debit_credit_advices` WRITE;
/*!40000 ALTER TABLE `debit_credit_advices` DISABLE KEYS */;
/*!40000 ALTER TABLE `debit_credit_advices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `delivery_terms`
--

DROP TABLE IF EXISTS `delivery_terms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `delivery_terms` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `alternate_delivery_terms_code` varchar(50) DEFAULT '',
  `delivery_cost_payment` varchar(50) DEFAULT '',
  `incoterms_code` varchar(50) DEFAULT '',
  `is_signature_required` varchar(50) DEFAULT '',
  `order_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `delivery_terms`
--

LOCK TABLES `delivery_terms` WRITE;
/*!40000 ALTER TABLE `delivery_terms` DISABLE KEYS */;
/*!40000 ALTER TABLE `delivery_terms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `description1000`
--

DROP TABLE IF EXISTS `description1000`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `description1000` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `desc1000` varchar(1000) DEFAULT NULL,
  `code_list_version` varchar(10) DEFAULT NULL,
  `currency_code` varchar(10) DEFAULT NULL,
  `fk_id` int(10) unsigned DEFAULT 0,
  `fk_table` varchar(70) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `description1000`
--

LOCK TABLES `description1000` WRITE;
/*!40000 ALTER TABLE `description1000` DISABLE KEYS */;
/*!40000 ALTER TABLE `description1000` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `description200`
--

DROP TABLE IF EXISTS `description200`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `description200` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `desc200` varchar(200) DEFAULT NULL,
  `code_list_version` varchar(10) DEFAULT NULL,
  `currency_code` varchar(10) DEFAULT NULL,
  `fk_id` int(10) unsigned DEFAULT 0,
  `fk_table` varchar(70) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `description200`
--

LOCK TABLES `description200` WRITE;
/*!40000 ALTER TABLE `description200` DISABLE KEYS */;
/*!40000 ALTER TABLE `description200` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `description500`
--

DROP TABLE IF EXISTS `description500`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `description500` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `desc500` varchar(500) DEFAULT NULL,
  `code_list_version` varchar(10) DEFAULT NULL,
  `currency_code` varchar(10) DEFAULT NULL,
  `fk_id` int(10) unsigned DEFAULT 0,
  `fk_table` varchar(70) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `description500`
--

LOCK TABLES `description500` WRITE;
/*!40000 ALTER TABLE `description500` DISABLE KEYS */;
/*!40000 ALTER TABLE `description500` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `description70`
--

DROP TABLE IF EXISTS `description70`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `description70` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `desc70` varchar(70) DEFAULT NULL,
  `code_list_version` varchar(10) DEFAULT NULL,
  `currency_code` varchar(10) DEFAULT NULL,
  `fk_id` int(10) unsigned DEFAULT 0,
  `fk_table` varchar(70) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `description70`
--

LOCK TABLES `description70` WRITE;
/*!40000 ALTER TABLE `description70` DISABLE KEYS */;
/*!40000 ALTER TABLE `description70` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `despatch_advice_item_totals`
--

DROP TABLE IF EXISTS `despatch_advice_item_totals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `despatch_advice_item_totals` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `additional_trade_item_identification` varchar(80) DEFAULT '',
  `additional_trade_item_identification_type_code` varchar(80) DEFAULT '',
  `code_list_version` varchar(35) DEFAULT '',
  `gtin` int(10) unsigned DEFAULT 0,
  `trade_item_identification` int(10) unsigned DEFAULT 0,
  `despatch_advice_id` int(10) unsigned DEFAULT 0,
  `despatch_advice_line_item_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `despatch_advice_item_totals`
--

LOCK TABLES `despatch_advice_item_totals` WRITE;
/*!40000 ALTER TABLE `despatch_advice_item_totals` DISABLE KEYS */;
/*!40000 ALTER TABLE `despatch_advice_item_totals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `despatch_advice_line_items`
--

DROP TABLE IF EXISTS `despatch_advice_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `despatch_advice_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `actual_processed_quantity` double DEFAULT 0,
  `measurement_unit_code` varchar(80) DEFAULT '',
  `code_list_version` varchar(35) DEFAULT '',
  `country_of_last_processing` varchar(80) DEFAULT '',
  `country_of_origin` varchar(80) DEFAULT '',
  `despatched_quantity` double DEFAULT 0,
  `dq_measurement_unit_code` varchar(80) DEFAULT '',
  `dq_code_list_version` varchar(35) DEFAULT '',
  `duty_fee_tax_liability` varchar(80) DEFAULT '',
  `extension` varchar(80) DEFAULT '',
  `free_goods_quantity` double DEFAULT 0,
  `fgq_measurement_unit_code` varchar(80) DEFAULT '',
  `fgq_code_list_version` varchar(35) DEFAULT '',
  `handling_instruction_code` varchar(80) DEFAULT '',
  `has_item_been_scanned_at_pos` varchar(80) DEFAULT '',
  `inventory_status_type` varchar(80) DEFAULT '',
  `line_item_number` int(10) unsigned DEFAULT 0,
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `requested_quantity` double DEFAULT 0,
  `rq_measurement_unit_code` varchar(80) DEFAULT '',
  `rq_code_list_version` varchar(35) DEFAULT '',
  `contract` int(10) unsigned DEFAULT 0,
  `coupon_clearing_house` int(10) unsigned DEFAULT 0,
  `customer` int(10) unsigned DEFAULT 0,
  `customer_document_reference` int(10) unsigned DEFAULT 0,
  `customer_reference` int(10) unsigned DEFAULT 0,
  `delivery_note` int(10) unsigned DEFAULT 0,
  `item_owner` int(10) unsigned DEFAULT 0,
  `original_supplier` int(10) unsigned DEFAULT 0,
  `product_certification` int(10) unsigned DEFAULT 0,
  `promotional_deal` int(10) unsigned DEFAULT 0,
  `purchase_conditions` int(10) unsigned DEFAULT 0,
  `purchase_order` int(10) unsigned DEFAULT 0,
  `referenced_consignment` int(10) unsigned DEFAULT 0,
  `requested_item_identification` int(10) unsigned DEFAULT 0,
  `specification` int(10) unsigned DEFAULT 0,
  `despatch_advice_id` int(10) unsigned DEFAULT 0,
  `first_in_first_out_date_time` datetime DEFAULT current_timestamp(),
  `pick_up_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `despatch_advice_line_items`
--

LOCK TABLES `despatch_advice_line_items` WRITE;
/*!40000 ALTER TABLE `despatch_advice_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `despatch_advice_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `despatch_advice_logistic_units`
--

DROP TABLE IF EXISTS `despatch_advice_logistic_units`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `despatch_advice_logistic_units` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `additional_logisitic_unit_identification` varchar(80) DEFAULT '',
  `additional_logistic_unit_identification_type_code` varchar(80) DEFAULT '',
  `code_list_version` varchar(80) DEFAULT '',
  `sscc` varchar(80) DEFAULT '',
  `ultimate_consignee` int(10) unsigned DEFAULT 0,
  `despatch_advice_id` int(10) unsigned DEFAULT 0,
  `estimated_delivery_date_time_at_ultimate_consignee` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `despatch_advice_logistic_units`
--

LOCK TABLES `despatch_advice_logistic_units` WRITE;
/*!40000 ALTER TABLE `despatch_advice_logistic_units` DISABLE KEYS */;
/*!40000 ALTER TABLE `despatch_advice_logistic_units` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `despatch_advice_quantity_variances`
--

DROP TABLE IF EXISTS `despatch_advice_quantity_variances`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `despatch_advice_quantity_variances` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `remaining_quantity_status_code` varchar(80) DEFAULT '',
  `variance_quantity` double DEFAULT 0,
  `vq_measurement_unit_code` varchar(80) DEFAULT '',
  `vq_code_list_version` varchar(35) DEFAULT '',
  `variance_reason_code` varchar(80) DEFAULT '',
  `despatch_advice_id` int(10) unsigned DEFAULT 0,
  `delivery_date_variance` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `despatch_advice_quantity_variances`
--

LOCK TABLES `despatch_advice_quantity_variances` WRITE;
/*!40000 ALTER TABLE `despatch_advice_quantity_variances` DISABLE KEYS */;
/*!40000 ALTER TABLE `despatch_advice_quantity_variances` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `despatch_advice_totals`
--

DROP TABLE IF EXISTS `despatch_advice_totals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `despatch_advice_totals` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `measurement_type` varchar(80) DEFAULT '',
  `measurement_value` varchar(80) DEFAULT '',
  `package_total` int(10) unsigned DEFAULT 0,
  `despatch_advice_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `despatch_advice_totals`
--

LOCK TABLES `despatch_advice_totals` WRITE;
/*!40000 ALTER TABLE `despatch_advice_totals` DISABLE KEYS */;
/*!40000 ALTER TABLE `despatch_advice_totals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `despatch_advice_transport_informations`
--

DROP TABLE IF EXISTS `despatch_advice_transport_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `despatch_advice_transport_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `route_id` int(10) unsigned DEFAULT 0,
  `transport_means_id` int(10) unsigned DEFAULT 0,
  `transport_means_name` varchar(80) DEFAULT '',
  `transport_means_type` varchar(80) DEFAULT '',
  `transport_mode_code` varchar(80) DEFAULT '',
  `bill_of_lading_number` int(10) unsigned DEFAULT 0,
  `additional_consignment_identification` varchar(80) DEFAULT '',
  `additional_consignment_identification_type_code` varchar(80) DEFAULT '',
  `code_list_version` varchar(35) DEFAULT '',
  `ginc` varchar(80) DEFAULT '',
  `driver` int(10) unsigned DEFAULT 0,
  `driver_id` int(10) unsigned DEFAULT 0,
  `receiver` int(10) unsigned DEFAULT 0,
  `receiver_id` int(10) unsigned DEFAULT 0,
  `additional_shipment_identification` varchar(80) DEFAULT '',
  `additional_shipment_identification_type_code` varchar(80) DEFAULT '',
  `gsin` varchar(80) DEFAULT '',
  `despatch_advice_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `despatch_advice_transport_informations`
--

LOCK TABLES `despatch_advice_transport_informations` WRITE;
/*!40000 ALTER TABLE `despatch_advice_transport_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `despatch_advice_transport_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `despatch_advices`
--

DROP TABLE IF EXISTS `despatch_advices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `despatch_advices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `delivery_type_code` varchar(80) DEFAULT '',
  `rack_id_at_pick_up_location` varchar(80) DEFAULT '',
  `total_deposit_amount` double DEFAULT 0,
  `tda_code_list_version` varchar(35) DEFAULT '',
  `tda_currency_code` varchar(80) DEFAULT '',
  `total_number_of_lines` int(10) unsigned DEFAULT 0,
  `blanket_order` int(10) unsigned DEFAULT 0,
  `buyer` int(10) unsigned DEFAULT 0,
  `carrier` int(10) unsigned DEFAULT 0,
  `contract` int(10) unsigned DEFAULT 0,
  `customer_document_reference` int(10) unsigned DEFAULT 0,
  `declarants_customs_identity` int(10) unsigned DEFAULT 0,
  `delivery_note` int(10) unsigned DEFAULT 0,
  `delivery_schedule` int(10) unsigned DEFAULT 0,
  `despatch_advice_identification` int(10) unsigned DEFAULT 0,
  `freight_forwarder` int(10) unsigned DEFAULT 0,
  `inventory_location` int(10) unsigned DEFAULT 0,
  `invoice` int(10) unsigned DEFAULT 0,
  `invoicee` int(10) unsigned DEFAULT 0,
  `logistic_service_provider` int(10) unsigned DEFAULT 0,
  `order_response` int(10) unsigned DEFAULT 0,
  `pick_up_location` int(10) unsigned DEFAULT 0,
  `product_certification` int(10) unsigned DEFAULT 0,
  `promotional_deal` int(10) unsigned DEFAULT 0,
  `purchase_conditions` int(10) unsigned DEFAULT 0,
  `purchase_order` int(10) unsigned DEFAULT 0,
  `receiver` int(10) unsigned DEFAULT 0,
  `returns_instruction` int(10) unsigned DEFAULT 0,
  `seller` int(10) unsigned DEFAULT 0,
  `ship_from` int(10) unsigned DEFAULT 0,
  `shipper` int(10) unsigned DEFAULT 0,
  `ship_to` int(10) unsigned DEFAULT 0,
  `specification` int(10) unsigned DEFAULT 0,
  `transport_instruction` int(10) unsigned DEFAULT 0,
  `ultimate_consignee` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `despatch_advices`
--

LOCK TABLES `despatch_advices` WRITE;
/*!40000 ALTER TABLE `despatch_advices` DISABLE KEYS */;
/*!40000 ALTER TABLE `despatch_advices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `despatch_informations`
--

DROP TABLE IF EXISTS `despatch_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `despatch_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `despatch_advice_id` int(10) unsigned DEFAULT 0,
  `actual_ship_date_time` datetime DEFAULT current_timestamp(),
  `despatch_date_time` datetime DEFAULT current_timestamp(),
  `estimated_delivery_date_time` datetime DEFAULT current_timestamp(),
  `estimated_delivery_date_time_at_ultimate_consignee` datetime DEFAULT current_timestamp(),
  `loading_date_time` datetime DEFAULT current_timestamp(),
  `pick_up_date_time` datetime DEFAULT current_timestamp(),
  `release_date_time_of_supplier` datetime DEFAULT current_timestamp(),
  `estimated_delivery_period_begin` datetime DEFAULT current_timestamp(),
  `estimated_delivery_period_end` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `despatch_informations`
--

LOCK TABLES `despatch_informations` WRITE;
/*!40000 ALTER TABLE `despatch_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `despatch_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `destinations`
--

DROP TABLE IF EXISTS `destinations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `destinations` (
  `dest_type` varchar(80) DEFAULT '',
  `destination` varchar(80) DEFAULT '',
  `event_id` int(10) unsigned DEFAULT NULL,
  `type_of_event` varchar(80) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `destinations`
--

LOCK TABLES `destinations` WRITE;
/*!40000 ALTER TABLE `destinations` DISABLE KEYS */;
/*!40000 ALTER TABLE `destinations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `duty_fee_tax_registrations`
--

DROP TABLE IF EXISTS `duty_fee_tax_registrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `duty_fee_tax_registrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `duty_fee_tax_agency_name` varchar(80) DEFAULT '',
  `duty_fee_tax_description` varchar(80) DEFAULT '',
  `duty_fee_tax_registration_type` varchar(80) DEFAULT '',
  `duty_fee_tax_type_code` varchar(80) DEFAULT '',
  `transactional_party_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `duty_fee_tax_registrations`
--

LOCK TABLES `duty_fee_tax_registrations` WRITE;
/*!40000 ALTER TABLE `duty_fee_tax_registrations` DISABLE KEYS */;
/*!40000 ALTER TABLE `duty_fee_tax_registrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ecom_document_references`
--

DROP TABLE IF EXISTS `ecom_document_references`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ecom_document_references` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `line_item_number` int(10) unsigned DEFAULT 0,
  `referenced_document_url` varchar(50) DEFAULT '',
  `revision_number` int(10) unsigned DEFAULT 0,
  `order_id` int(10) unsigned DEFAULT 0,
  `order_response_id` int(10) unsigned DEFAULT 0,
  `creation_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ecom_document_references`
--

LOCK TABLES `ecom_document_references` WRITE;
/*!40000 ALTER TABLE `ecom_document_references` DISABLE KEYS */;
/*!40000 ALTER TABLE `ecom_document_references` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ecom_entity_identifications`
--

DROP TABLE IF EXISTS `ecom_entity_identifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ecom_entity_identifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `entity_identification` varchar(50) DEFAULT '',
  `gln` varchar(80) DEFAULT '',
  `order_id` int(10) unsigned DEFAULT 0,
  `order_response_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ecom_entity_identifications`
--

LOCK TABLES `ecom_entity_identifications` WRITE;
/*!40000 ALTER TABLE `ecom_entity_identifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `ecom_entity_identifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ecom_trade_item_identifications`
--

DROP TABLE IF EXISTS `ecom_trade_item_identifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ecom_trade_item_identifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `additional_trade_item_identification` varchar(80) DEFAULT '',
  `additional_trade_item_identification_type_code` varchar(80) DEFAULT '',
  `code_list_version` varchar(35) DEFAULT '',
  `gtin` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ecom_trade_item_identifications`
--

LOCK TABLES `ecom_trade_item_identifications` WRITE;
/*!40000 ALTER TABLE `ecom_trade_item_identifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `ecom_trade_item_identifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `epcs`
--

DROP TABLE IF EXISTS `epcs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `epcs` (
  `epc_value` varchar(80) DEFAULT '',
  `event_id` int(10) unsigned DEFAULT NULL,
  `type_of_event` varchar(80) DEFAULT '',
  `type_of_epc` varchar(80) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `epcs`
--

LOCK TABLES `epcs` WRITE;
/*!40000 ALTER TABLE `epcs` DISABLE KEYS */;
/*!40000 ALTER TABLE `epcs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `financial_accounts`
--

DROP TABLE IF EXISTS `financial_accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `financial_accounts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `financial_account_name` varchar(80) DEFAULT '',
  `financial_account_number` varchar(80) DEFAULT '',
  `financial_account_number_type_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `financial_accounts`
--

LOCK TABLES `financial_accounts` WRITE;
/*!40000 ALTER TABLE `financial_accounts` DISABLE KEYS */;
/*!40000 ALTER TABLE `financial_accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `financial_institution_informations`
--

DROP TABLE IF EXISTS `financial_institution_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `financial_institution_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `financial_institution_branch_name` varchar(80) DEFAULT '',
  `financial_institution_name` varchar(80) DEFAULT '',
  `address` int(10) unsigned DEFAULT 0,
  `financial_routing_number` varchar(80) DEFAULT '',
  `financial_routing_number_type_code` varchar(80) DEFAULT '',
  `financial_account_name` varchar(80) DEFAULT '',
  `financial_account_number` varchar(80) DEFAULT '',
  `financial_account_number_type_code` varchar(80) DEFAULT '',
  `transactional_party_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `financial_institution_informations`
--

LOCK TABLES `financial_institution_informations` WRITE;
/*!40000 ALTER TABLE `financial_institution_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `financial_institution_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `financial_routing_numbers`
--

DROP TABLE IF EXISTS `financial_routing_numbers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `financial_routing_numbers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `financial_routing_number` varchar(80) DEFAULT '',
  `financial_routing_number_type_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `financial_routing_numbers`
--

LOCK TABLES `financial_routing_numbers` WRITE;
/*!40000 ALTER TABLE `financial_routing_numbers` DISABLE KEYS */;
/*!40000 ALTER TABLE `financial_routing_numbers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fish_catch_or_production_dates`
--

DROP TABLE IF EXISTS `fish_catch_or_production_dates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fish_catch_or_production_dates` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `catch_end_date` datetime DEFAULT current_timestamp(),
  `catch_start_date` datetime DEFAULT current_timestamp(),
  `first_freeze_date` datetime DEFAULT current_timestamp(),
  `catch_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fish_catch_or_production_dates`
--

LOCK TABLES `fish_catch_or_production_dates` WRITE;
/*!40000 ALTER TABLE `fish_catch_or_production_dates` DISABLE KEYS */;
/*!40000 ALTER TABLE `fish_catch_or_production_dates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fish_catch_or_productions`
--

DROP TABLE IF EXISTS `fish_catch_or_productions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fish_catch_or_productions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `catch_area` varchar(80) DEFAULT '',
  `fishing_gear_type_code` varchar(80) DEFAULT '',
  `production_method_for_fish_and_seafood_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fish_catch_or_productions`
--

LOCK TABLES `fish_catch_or_productions` WRITE;
/*!40000 ALTER TABLE `fish_catch_or_productions` DISABLE KEYS */;
/*!40000 ALTER TABLE `fish_catch_or_productions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fish_despatch_advice_line_item_extensions`
--

DROP TABLE IF EXISTS `fish_despatch_advice_line_item_extensions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fish_despatch_advice_line_item_extensions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `aquatic_species_code` varchar(80) DEFAULT '',
  `aquatic_species_name` varchar(80) DEFAULT '',
  `fish_presentation_code` varchar(80) DEFAULT '',
  `fp_code_list_agency_name` varchar(35) DEFAULT '',
  `fp_code_list_version` varchar(35) DEFAULT '',
  `fp_code_list_name` varchar(35) DEFAULT '',
  `fish_size_code` varchar(80) DEFAULT '',
  `fs_code_list_agency_name` varchar(35) DEFAULT '',
  `fs_code_list_version` varchar(35) DEFAULT '',
  `fs_code_list_name` varchar(35) DEFAULT '',
  `quality_grade_code` varchar(80) DEFAULT '',
  `qg_code_list_agency_name` varchar(35) DEFAULT '',
  `qg_code_list_version` varchar(35) DEFAULT '',
  `qg_code_list_name` varchar(35) DEFAULT '',
  `storage_state_code` varchar(80) DEFAULT '',
  `aqua_culture_production_unit` int(10) unsigned DEFAULT 0,
  `fishing_vessel` int(10) unsigned DEFAULT 0,
  `place_of_slaughter` int(10) unsigned DEFAULT 0,
  `port_of_landing` int(10) unsigned DEFAULT 0,
  `fish_catch_or_production_date_id` int(10) unsigned DEFAULT 0,
  `fish_catch_or_production_id` int(10) unsigned DEFAULT 0,
  `date_of_landing` datetime DEFAULT current_timestamp(),
  `date_of_slaughter` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fish_despatch_advice_line_item_extensions`
--

LOCK TABLES `fish_despatch_advice_line_item_extensions` WRITE;
/*!40000 ALTER TABLE `fish_despatch_advice_line_item_extensions` DISABLE KEYS */;
/*!40000 ALTER TABLE `fish_despatch_advice_line_item_extensions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `identifiers`
--

DROP TABLE IF EXISTS `identifiers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `identifiers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `identifier` varchar(80) DEFAULT NULL,
  `identification_scheme_agency_code` varchar(20) DEFAULT '',
  `identification_scheme_agency_code_code_list_version` varchar(20) DEFAULT '',
  `identification_scheme_agency_name` varchar(20) DEFAULT '',
  `identification_scheme_name` varchar(20) DEFAULT '',
  `fk_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `identifiers`
--

LOCK TABLES `identifiers` WRITE;
/*!40000 ALTER TABLE `identifiers` DISABLE KEYS */;
/*!40000 ALTER TABLE `identifiers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `installment_dues`
--

DROP TABLE IF EXISTS `installment_dues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `installment_dues` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `percent_of_payment_due` double DEFAULT 0,
  `payment_time_period` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `installment_dues`
--

LOCK TABLES `installment_dues` WRITE;
/*!40000 ALTER TABLE `installment_dues` DISABLE KEYS */;
/*!40000 ALTER TABLE `installment_dues` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventory_activity_line_items`
--

DROP TABLE IF EXISTS `inventory_activity_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventory_activity_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `line_item_number` int(10) unsigned DEFAULT 0,
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `inventory_item_location_information_id` int(10) unsigned DEFAULT 0,
  `inventory_report_id` int(10) unsigned DEFAULT 0,
  `reporting_period_begin` datetime DEFAULT current_timestamp(),
  `reporting_period_end` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventory_activity_line_items`
--

LOCK TABLES `inventory_activity_line_items` WRITE;
/*!40000 ALTER TABLE `inventory_activity_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `inventory_activity_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventory_activity_quantity_specifications`
--

DROP TABLE IF EXISTS `inventory_activity_quantity_specifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventory_activity_quantity_specifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inventory_activity_type_code` varchar(80) DEFAULT '',
  `inventory_movement_type_code` varchar(80) DEFAULT '',
  `quantity_of_units` double DEFAULT 0,
  `qou_measurement_unit_code` varchar(80) DEFAULT '',
  `qou_code_list_version` varchar(35) DEFAULT '',
  `inventory_status_line_item_id` int(10) unsigned DEFAULT 0,
  `inventory_activity_line_item_id` int(10) unsigned DEFAULT 0,
  `inventory_report_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventory_activity_quantity_specifications`
--

LOCK TABLES `inventory_activity_quantity_specifications` WRITE;
/*!40000 ALTER TABLE `inventory_activity_quantity_specifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `inventory_activity_quantity_specifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventory_item_location_informations`
--

DROP TABLE IF EXISTS `inventory_item_location_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventory_item_location_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inventory_location_id` int(10) unsigned DEFAULT 0,
  `inventory_report_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventory_item_location_informations`
--

LOCK TABLES `inventory_item_location_informations` WRITE;
/*!40000 ALTER TABLE `inventory_item_location_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `inventory_item_location_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventory_reports`
--

DROP TABLE IF EXISTS `inventory_reports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventory_reports` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `inventory_report_type_code` varchar(80) DEFAULT '',
  `structure_type_code` varchar(80) DEFAULT '',
  `inventory_report_identification` int(10) unsigned DEFAULT 0,
  `inventory_reporting_party` int(10) unsigned DEFAULT 0,
  `inventory_report_to_party` int(10) unsigned DEFAULT 0,
  `reporting_period_begin` datetime DEFAULT current_timestamp(),
  `reporting_period_end` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventory_reports`
--

LOCK TABLES `inventory_reports` WRITE;
/*!40000 ALTER TABLE `inventory_reports` DISABLE KEYS */;
/*!40000 ALTER TABLE `inventory_reports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventory_status_line_items`
--

DROP TABLE IF EXISTS `inventory_status_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventory_status_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `handling_unit_type` varchar(80) DEFAULT '',
  `inventory_unit_cost` double DEFAULT 0,
  `iuc_code_list_version` varchar(35) DEFAULT '',
  `iuc_currency_code` varchar(80) DEFAULT '',
  `line_item_number` int(10) unsigned DEFAULT 0,
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `inventory_status_owner` int(10) unsigned DEFAULT 0,
  `inventory_sub_location_id` int(10) unsigned DEFAULT 0,
  `logistic_unit_identification` int(10) unsigned DEFAULT 0,
  `returnable_asset_identification` int(10) unsigned DEFAULT 0,
  `inventory_report_type_code` varchar(80) DEFAULT '',
  `structure_type_code` varchar(80) DEFAULT '',
  `inventory_report_identification` int(10) unsigned DEFAULT 0,
  `inventory_reporting_party` int(10) unsigned DEFAULT 0,
  `inventory_report_to_party` int(10) unsigned DEFAULT 0,
  `inventory_item_location_information_id` int(10) unsigned DEFAULT 0,
  `inventory_report_id` int(10) unsigned DEFAULT 0,
  `first_in_first_out_date_time_begin` datetime DEFAULT current_timestamp(),
  `first_in_first_out_date_time_end` datetime DEFAULT current_timestamp(),
  `inventory_date_time_begin` datetime DEFAULT current_timestamp(),
  `inventory_date_time_end` datetime DEFAULT current_timestamp(),
  `reporting_period_begin` datetime DEFAULT current_timestamp(),
  `reporting_period_end` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventory_status_line_items`
--

LOCK TABLES `inventory_status_line_items` WRITE;
/*!40000 ALTER TABLE `inventory_status_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `inventory_status_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventory_sub_locations`
--

DROP TABLE IF EXISTS `inventory_sub_locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `inventory_sub_locations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `additional_party_identification` varchar(80) DEFAULT '',
  `additional_party_identification_type_code` varchar(80) DEFAULT '',
  `code_list_version` varchar(80) DEFAULT '',
  `gln` varchar(80) DEFAULT '',
  `gln_extension` varchar(80) DEFAULT '',
  `inventory_sub_location_function_code` varchar(80) DEFAULT '',
  `inventory_sub_location_type_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventory_sub_locations`
--

LOCK TABLES `inventory_sub_locations` WRITE;
/*!40000 ALTER TABLE `inventory_sub_locations` DISABLE KEYS */;
/*!40000 ALTER TABLE `inventory_sub_locations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_allowance_charges`
--

DROP TABLE IF EXISTS `invoice_allowance_charges`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoice_allowance_charges` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `levied_duty_fee_tax` int(10) unsigned DEFAULT 0,
  `allowance_charge` int(10) unsigned DEFAULT 0,
  `invoice_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_allowance_charges`
--

LOCK TABLES `invoice_allowance_charges` WRITE;
/*!40000 ALTER TABLE `invoice_allowance_charges` DISABLE KEYS */;
/*!40000 ALTER TABLE `invoice_allowance_charges` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_line_item_information_after_taxes`
--

DROP TABLE IF EXISTS `invoice_line_item_information_after_taxes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoice_line_item_information_after_taxes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `amount_exclusive_allowances_charges` double DEFAULT 0,
  `aeac_code_list_version` varchar(35) DEFAULT '',
  `aeac_currency_code` varchar(80) DEFAULT '',
  `amount_inclusive_allowances_charges` double DEFAULT 0,
  `aiac_code_list_version` varchar(35) DEFAULT '',
  `aiac_currency_code` varchar(80) DEFAULT '',
  `invoice_id` int(10) unsigned DEFAULT 0,
  `invoice_line_item_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_line_item_information_after_taxes`
--

LOCK TABLES `invoice_line_item_information_after_taxes` WRITE;
/*!40000 ALTER TABLE `invoice_line_item_information_after_taxes` DISABLE KEYS */;
/*!40000 ALTER TABLE `invoice_line_item_information_after_taxes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_line_items`
--

DROP TABLE IF EXISTS `invoice_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoice_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `amount_exclusive_allowances_charges` double DEFAULT 0,
  `aeac_code_list_version` varchar(35) DEFAULT '',
  `aeac_currency_code` varchar(80) DEFAULT '',
  `amount_inclusive_allowances_charges` double DEFAULT 0,
  `aiac_code_list_version` varchar(35) DEFAULT '',
  `aiac_currency_code` varchar(80) DEFAULT '',
  `credit_line_indicator` varchar(80) DEFAULT '',
  `credit_reason` varchar(80) DEFAULT '',
  `delivered_quantity` double DEFAULT 0,
  `dq_measurement_unit_code` varchar(80) DEFAULT '',
  `dq_code_list_version` varchar(35) DEFAULT '',
  `excluded_from_payment_discount_indicator` tinyint(1) DEFAULT 0,
  `extension` varchar(35) DEFAULT '',
  `free_goods_quantity` double DEFAULT 0,
  `fgq_measurement_unit_code` varchar(80) DEFAULT '',
  `fgq_code_list_version` varchar(35) DEFAULT '',
  `invoiced_quantity` double DEFAULT 0,
  `iq_measurement_unit_code` varchar(80) DEFAULT '',
  `iq_code_list_version` varchar(35) DEFAULT '',
  `item_price_base_quantity` double DEFAULT 0,
  `ipbq_measurement_unit_code` varchar(80) DEFAULT '',
  `ipbq_code_list_version` varchar(35) DEFAULT '',
  `item_price_exclusive_allowances_charges` double DEFAULT 0,
  `ipeac_code_list_version` varchar(35) DEFAULT '',
  `ipeac_currency_code` varchar(80) DEFAULT '',
  `item_price_inclusive_allowances_charges` double DEFAULT 0,
  `ipiac_code_list_version` varchar(35) DEFAULT '',
  `ipiac_currency_code` varchar(80) DEFAULT '',
  `legally_fixed_retail_price` double DEFAULT 0,
  `lfrp_code_list_version` varchar(35) DEFAULT '',
  `lfrp_currency_code` varchar(80) DEFAULT '',
  `line_item_number` int(10) unsigned DEFAULT 0,
  `margin_scheme_information` varchar(80) DEFAULT '',
  `owenrship_prior_to_payment` varchar(80) DEFAULT '',
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `recommended_retail_price` double DEFAULT 0,
  `rrp_code_list_version` varchar(35) DEFAULT '',
  `rrp_currency_code` varchar(80) DEFAULT '',
  `retail_price_excluding_excise` double DEFAULT 0,
  `rpee_code_list_version` varchar(35) DEFAULT '',
  `rpee_currency_code` varchar(80) DEFAULT '',
  `total_ordered_quantity` double DEFAULT 0,
  `toq_measurement_unit_code` varchar(80) DEFAULT '',
  `toq_code_list_version` varchar(35) DEFAULT '',
  `consumption_report` int(10) unsigned DEFAULT 0,
  `contract` int(10) unsigned DEFAULT 0,
  `delivery_note` int(10) unsigned DEFAULT 0,
  `despatch_advice` int(10) unsigned DEFAULT 0,
  `energy_quantity` int(10) unsigned DEFAULT 0,
  `inventory_location_from` int(10) unsigned DEFAULT 0,
  `inventory_location_to` int(10) unsigned DEFAULT 0,
  `promotional_deal` int(10) unsigned DEFAULT 0,
  `purchase_conditions` int(10) unsigned DEFAULT 0,
  `purchase_order` int(10) unsigned DEFAULT 0,
  `receiving_advice` int(10) unsigned DEFAULT 0,
  `returnable_asset_identification` int(10) unsigned DEFAULT 0,
  `sales_order` int(10) unsigned DEFAULT 0,
  `ship_from` int(10) unsigned DEFAULT 0,
  `ship_to` int(10) unsigned DEFAULT 0,
  `trade_agreement` int(10) unsigned DEFAULT 0,
  `invoice_id` int(10) unsigned DEFAULT 0,
  `transfer_of_ownership_date` datetime DEFAULT current_timestamp(),
  `actual_delivery_date` datetime DEFAULT current_timestamp(),
  `servicetime_period_line_level_begin` datetime DEFAULT current_timestamp(),
  `servicetime_period_line_level_end` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_line_items`
--

LOCK TABLES `invoice_line_items` WRITE;
/*!40000 ALTER TABLE `invoice_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `invoice_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoice_totals`
--

DROP TABLE IF EXISTS `invoice_totals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoice_totals` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `base_amount` double DEFAULT 0,
  `ba_code_list_version` varchar(35) DEFAULT '',
  `ba_currency_code` varchar(80) DEFAULT '',
  `prepaid_amount` double DEFAULT 0,
  `pa_code_list_version` varchar(35) DEFAULT '',
  `pa_currency_code` varchar(80) DEFAULT '',
  `tax_accounting_currency` varchar(80) DEFAULT '',
  `total_amount_invoice_allowances_charges` double DEFAULT 0,
  `taiac_code_list_version` varchar(35) DEFAULT '',
  `taiac_currency_code` varchar(80) DEFAULT '',
  `total_amount_line_allowances_charges` double DEFAULT 0,
  `talac_code_list_version` varchar(35) DEFAULT '',
  `talac_currency_code` varchar(80) DEFAULT '',
  `total_economic_value` double DEFAULT 0,
  `tev_code_list_version` varchar(35) DEFAULT '',
  `tev_currency_code` varchar(80) DEFAULT '',
  `total_goods_value` double DEFAULT 0,
  `tgv_code_list_version` varchar(35) DEFAULT '',
  `tgv_currency_code` varchar(80) DEFAULT '',
  `total_invoice_amount` double DEFAULT 0,
  `tia_code_list_version` varchar(35) DEFAULT '',
  `tia_currency_code` varchar(80) DEFAULT '',
  `total_invoice_amount_payable` double DEFAULT 0,
  `tiap_code_list_version` varchar(35) DEFAULT '',
  `tiap_currency_code` varchar(80) DEFAULT '',
  `total_line_amount_exclusive_allowances_charges` double DEFAULT 0,
  `tlaeac_code_list_version` varchar(35) DEFAULT '',
  `tlaeac_currency_code` varchar(80) DEFAULT '',
  `total_line_amount_inclusive_allowances_charges` double DEFAULT 0,
  `tlaiac_code_list_version` varchar(35) DEFAULT '',
  `tlaiac_currency_code` varchar(80) DEFAULT '',
  `total_payment_discount_basis_amount` double DEFAULT 0,
  `tpdba_code_list_version` varchar(35) DEFAULT '',
  `tpdba_currency_code` varchar(80) DEFAULT '',
  `total_retail_value` double DEFAULT 0,
  `trv_code_list_version` varchar(35) DEFAULT '',
  `trv_currency_code` varchar(80) DEFAULT '',
  `total_tax_amount` double DEFAULT 0,
  `tta_code_list_version` varchar(35) DEFAULT '',
  `tta_currency_code` varchar(80) DEFAULT '',
  `total_tax_basis_amount` double DEFAULT 0,
  `ttba_code_list_version` varchar(35) DEFAULT '',
  `ttba_currency_code` varchar(80) DEFAULT '',
  `total_vat_amount` double DEFAULT 0,
  `tva_code_list_version` varchar(35) DEFAULT '',
  `tva_currency_code` varchar(80) DEFAULT '',
  `invoice_line_item_id` int(10) unsigned DEFAULT 0,
  `invoice_id` int(10) unsigned DEFAULT 0,
  `prepaid_amount_date` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoice_totals`
--

LOCK TABLES `invoice_totals` WRITE;
/*!40000 ALTER TABLE `invoice_totals` DISABLE KEYS */;
/*!40000 ALTER TABLE `invoice_totals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `invoices`
--

DROP TABLE IF EXISTS `invoices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `invoices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `country_of_supply_of_goods` varchar(80) DEFAULT '',
  `credit_reason_code` varchar(80) DEFAULT '',
  `discount_agreement_terms` varchar(80) DEFAULT '',
  `invoice_currency_code` varchar(80) DEFAULT '',
  `invoice_type` varchar(80) DEFAULT '',
  `is_buyer_based_in_eu` tinyint(1) DEFAULT 0,
  `is_first_seller_based_in_eu` tinyint(1) DEFAULT 0,
  `supplier_account_receivable` varchar(80) DEFAULT NULL,
  `blanket_order` int(10) unsigned DEFAULT 0,
  `buyer` int(10) unsigned DEFAULT 0,
  `contract` int(10) unsigned DEFAULT 0,
  `delivery_note` int(10) unsigned DEFAULT 0,
  `despatch_advice` int(10) unsigned DEFAULT 0,
  `dispute_notice` int(10) unsigned DEFAULT 0,
  `inventory_location` int(10) unsigned DEFAULT 0,
  `inventory_report` int(10) unsigned DEFAULT 0,
  `invoice` int(10) unsigned DEFAULT 0,
  `invoice_identification` int(10) unsigned DEFAULT 0,
  `manifest` int(10) unsigned DEFAULT 0,
  `order_response` int(10) unsigned DEFAULT 0,
  `payee` int(10) unsigned DEFAULT 0,
  `payer` int(10) unsigned DEFAULT 0,
  `pickup_from` int(10) unsigned DEFAULT 0,
  `price_list` int(10) unsigned DEFAULT 0,
  `promotional_deal` int(10) unsigned DEFAULT 0,
  `purchase_order` int(10) unsigned DEFAULT 0,
  `receiving_advice` int(10) unsigned DEFAULT 0,
  `remit_to` int(10) unsigned DEFAULT 0,
  `returns_notice` int(10) unsigned DEFAULT 0,
  `sales_order` int(10) unsigned DEFAULT 0,
  `sales_report` int(10) unsigned DEFAULT 0,
  `seller` int(10) unsigned DEFAULT 0,
  `ship_from` int(10) unsigned DEFAULT 0,
  `ship_to` int(10) unsigned DEFAULT 0,
  `supplier_agent_representative` int(10) unsigned DEFAULT 0,
  `supplier_corporate_office` int(10) unsigned DEFAULT 0,
  `tax_currency_information` int(10) unsigned DEFAULT 0,
  `tax_representative` int(10) unsigned DEFAULT 0,
  `trade_agreement` int(10) unsigned DEFAULT 0,
  `ultimate_consignee` int(10) unsigned DEFAULT 0,
  `actual_delivery_date` datetime DEFAULT current_timestamp(),
  `invoicing_period_begin` datetime DEFAULT current_timestamp(),
  `invoicing_period_end` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `invoices`
--

LOCK TABLES `invoices` WRITE;
/*!40000 ALTER TABLE `invoices` DISABLE KEYS */;
/*!40000 ALTER TABLE `invoices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legal_registrations`
--

DROP TABLE IF EXISTS `legal_registrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legal_registrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `legal_registration_additional_information` varchar(80) DEFAULT '',
  `legal_registration_number` varchar(80) DEFAULT '',
  `legal_registration_type` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legal_registrations`
--

LOCK TABLES `legal_registrations` WRITE;
/*!40000 ALTER TABLE `legal_registrations` DISABLE KEYS */;
/*!40000 ALTER TABLE `legal_registrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `levied_duty_fee_taxes`
--

DROP TABLE IF EXISTS `levied_duty_fee_taxes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `levied_duty_fee_taxes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `duty_fee_tax_accounting_currency` varchar(80) DEFAULT '',
  `duty_fee_tax_agency_name` varchar(80) DEFAULT '',
  `duty_fee_tax_amount` double DEFAULT 0,
  `dfta_code_list_version` varchar(35) DEFAULT '',
  `dfta_currency_code` varchar(80) DEFAULT '',
  `duty_fee_tax_amount_in_accounting_currency` double DEFAULT 0,
  `dftaiac_code_list_version` varchar(35) DEFAULT '',
  `dftaiac_currency_code` varchar(80) DEFAULT '',
  `duty_fee_tax_basis_amount` double DEFAULT 0,
  `dftba_code_list_version` varchar(35) DEFAULT '',
  `dftba_currency_code` varchar(80) DEFAULT '',
  `duty_fee_tax_basis_amount_in_accounting_currency` double DEFAULT 0,
  `dftbaiac_code_list_version` varchar(35) DEFAULT '',
  `dftbaiac_currency_code` varchar(80) DEFAULT '',
  `duty_fee_tax_category_code` varchar(80) DEFAULT '',
  `duty_fee_tax_description` varchar(80) DEFAULT '',
  `duty_fee_tax_exemption_description` varchar(80) DEFAULT '',
  `duty_fee_tax_exemption_reason` varchar(80) DEFAULT '',
  `duty_fee_tax_percentage` double DEFAULT 0,
  `duty_fee_tax_type_code` varchar(80) DEFAULT '',
  `extension` varchar(80) DEFAULT '',
  `order_line_item_id` int(10) unsigned DEFAULT 0,
  `invoice_line_item_id` int(10) unsigned DEFAULT 0,
  `invoice_id` int(10) unsigned DEFAULT 0,
  `duty_fee_tax_point_date` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `levied_duty_fee_taxes`
--

LOCK TABLES `levied_duty_fee_taxes` WRITE;
/*!40000 ALTER TABLE `levied_duty_fee_taxes` DISABLE KEYS */;
/*!40000 ALTER TABLE `levied_duty_fee_taxes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logistic_locations`
--

DROP TABLE IF EXISTS `logistic_locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logistic_locations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `gln` varchar(50) DEFAULT '',
  `location_name` varchar(50) DEFAULT NULL,
  `sublocation_identification` varchar(50) DEFAULT NULL,
  `un_location_code` varchar(50) DEFAULT '',
  `utc_offset` double DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logistic_locations`
--

LOCK TABLES `logistic_locations` WRITE;
/*!40000 ALTER TABLE `logistic_locations` DISABLE KEYS */;
/*!40000 ALTER TABLE `logistic_locations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logistic_unit_inventory_events`
--

DROP TABLE IF EXISTS `logistic_unit_inventory_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logistic_unit_inventory_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_identifier` varchar(80) DEFAULT '',
  `inventory_business_step_code` varchar(80) DEFAULT '',
  `inventory_event_reason_code` varchar(80) DEFAULT '',
  `inventory_movement_type_code` varchar(80) DEFAULT '',
  `inventory_sub_location_id` int(10) unsigned DEFAULT 0,
  `event_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logistic_unit_inventory_events`
--

LOCK TABLES `logistic_unit_inventory_events` WRITE;
/*!40000 ALTER TABLE `logistic_unit_inventory_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `logistic_unit_inventory_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logistic_unit_inventory_statuses`
--

DROP TABLE IF EXISTS `logistic_unit_inventory_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logistic_unit_inventory_statuses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inventory_disposition_code` varchar(80) DEFAULT '',
  `inventory_sub_location_id` int(10) unsigned DEFAULT 0,
  `inventory_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logistic_unit_inventory_statuses`
--

LOCK TABLES `logistic_unit_inventory_statuses` WRITE;
/*!40000 ALTER TABLE `logistic_unit_inventory_statuses` DISABLE KEYS */;
/*!40000 ALTER TABLE `logistic_unit_inventory_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logistic_unit_references`
--

DROP TABLE IF EXISTS `logistic_unit_references`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logistic_unit_references` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `trade_item_quantity` double DEFAULT 0,
  `q_measurement_unit_code` varchar(80) DEFAULT '',
  `q_code_list_version` varchar(35) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logistic_unit_references`
--

LOCK TABLES `logistic_unit_references` WRITE;
/*!40000 ALTER TABLE `logistic_unit_references` DISABLE KEYS */;
/*!40000 ALTER TABLE `logistic_unit_references` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logistic_units`
--

DROP TABLE IF EXISTS `logistic_units`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logistic_units` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `child_package_type_code` varchar(80) DEFAULT '',
  `level_identification` int(10) unsigned DEFAULT 0,
  `package_type_code` varchar(80) DEFAULT '',
  `parent_level_identification` int(10) unsigned DEFAULT 0,
  `quantity_of_children` int(10) unsigned DEFAULT 0,
  `quantity_of_logistic_units` int(10) unsigned DEFAULT 0,
  `logistic_unit_identification` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logistic_units`
--

LOCK TABLES `logistic_units` WRITE;
/*!40000 ALTER TABLE `logistic_units` DISABLE KEYS */;
/*!40000 ALTER TABLE `logistic_units` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logistics_inventory_report_inventory_locations`
--

DROP TABLE IF EXISTS `logistics_inventory_report_inventory_locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logistics_inventory_report_inventory_locations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inventory_location_id` int(10) unsigned DEFAULT 0,
  `logistics_inventory_report_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logistics_inventory_report_inventory_locations`
--

LOCK TABLES `logistics_inventory_report_inventory_locations` WRITE;
/*!40000 ALTER TABLE `logistics_inventory_report_inventory_locations` DISABLE KEYS */;
/*!40000 ALTER TABLE `logistics_inventory_report_inventory_locations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `logistics_inventory_reports`
--

DROP TABLE IF EXISTS `logistics_inventory_reports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logistics_inventory_reports` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `structure_type_code` varchar(80) DEFAULT '',
  `type_of_service_transaction` varchar(80) DEFAULT '',
  `inventory_reporting_party` int(10) unsigned DEFAULT 0,
  `inventory_report_to_party` int(10) unsigned DEFAULT 0,
  `logistics_inventory_report_identification` int(10) unsigned DEFAULT 0,
  `logistics_inventory_report_request` int(10) unsigned DEFAULT 0,
  `reporting_period_begin` datetime DEFAULT current_timestamp(),
  `reporting_period_end` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logistics_inventory_reports`
--

LOCK TABLES `logistics_inventory_reports` WRITE;
/*!40000 ALTER TABLE `logistics_inventory_reports` DISABLE KEYS */;
/*!40000 ALTER TABLE `logistics_inventory_reports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_acidities`
--

DROP TABLE IF EXISTS `meat_acidities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_acidities` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `acidity_measurement_time` int(10) unsigned DEFAULT 0,
  `acidity_of_meat` double DEFAULT 0,
  `meat_slaughtering_detail_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_acidities`
--

LOCK TABLES `meat_acidities` WRITE;
/*!40000 ALTER TABLE `meat_acidities` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_acidities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_activity_histories`
--

DROP TABLE IF EXISTS `meat_activity_histories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_activity_histories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `activity_sub_step_identification` int(10) unsigned DEFAULT 0,
  `country_of_activity_code` varchar(80) DEFAULT '',
  `current_step_identification` int(10) unsigned DEFAULT 0,
  `meat_processing_activity_type_code` varchar(80) DEFAULT '',
  `movement_reason_code` varchar(80) DEFAULT '',
  `next_step_identification` int(10) unsigned DEFAULT 0,
  `meat_mincing_detail_id` int(10) unsigned DEFAULT 0,
  `meat_fattening_detail_id` int(10) unsigned DEFAULT 0,
  `meat_cutting_detail_id` int(10) unsigned DEFAULT 0,
  `meat_breeding_detail_id` int(10) unsigned DEFAULT 0,
  `meat_processing_party_id` int(10) unsigned DEFAULT 0,
  `meat_work_item_identification_id` int(10) unsigned DEFAULT 0,
  `meat_slaughtering_detail_id` int(10) unsigned DEFAULT 0,
  `meat_despatch_advice_line_item_extension_id` int(10) unsigned DEFAULT 0,
  `date_of_arrival` datetime DEFAULT current_timestamp(),
  `date_of_departure` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_activity_histories`
--

LOCK TABLES `meat_activity_histories` WRITE;
/*!40000 ALTER TABLE `meat_activity_histories` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_activity_histories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_breeding_details`
--

DROP TABLE IF EXISTS `meat_breeding_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_breeding_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `animal_type_code` varchar(80) DEFAULT '',
  `breed_code` varchar(80) DEFAULT '',
  `breed_of_father_code` varchar(80) DEFAULT '',
  `breed_of_mother_code` varchar(80) DEFAULT '',
  `cross_breed_indicator` tinyint(1) DEFAULT 0,
  `feeding_system_code` varchar(80) DEFAULT '',
  `gender_code` varchar(80) DEFAULT '',
  `housing_system_code` varchar(80) DEFAULT '',
  `breeding_dna_test` int(10) unsigned DEFAULT 0,
  `date_of_birth` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_breeding_details`
--

LOCK TABLES `meat_breeding_details` WRITE;
/*!40000 ALTER TABLE `meat_breeding_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_breeding_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_cutting_details`
--

DROP TABLE IF EXISTS `meat_cutting_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_cutting_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `meat_profile_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_cutting_details`
--

LOCK TABLES `meat_cutting_details` WRITE;
/*!40000 ALTER TABLE `meat_cutting_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_cutting_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_despatch_advice_line_item_extensions`
--

DROP TABLE IF EXISTS `meat_despatch_advice_line_item_extensions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_despatch_advice_line_item_extensions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `animal_identification_id` int(10) unsigned DEFAULT 0,
  `slaughter_number_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_despatch_advice_line_item_extensions`
--

LOCK TABLES `meat_despatch_advice_line_item_extensions` WRITE;
/*!40000 ALTER TABLE `meat_despatch_advice_line_item_extensions` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_despatch_advice_line_item_extensions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_fattening_details`
--

DROP TABLE IF EXISTS `meat_fattening_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_fattening_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `feeding_system_code` varchar(80) DEFAULT '',
  `housing_system_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_fattening_details`
--

LOCK TABLES `meat_fattening_details` WRITE;
/*!40000 ALTER TABLE `meat_fattening_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_fattening_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_mincing_details`
--

DROP TABLE IF EXISTS `meat_mincing_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_mincing_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `fat_content_percent` double DEFAULT 0,
  `mincing_type_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_mincing_details`
--

LOCK TABLES `meat_mincing_details` WRITE;
/*!40000 ALTER TABLE `meat_mincing_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_mincing_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_processing_parties`
--

DROP TABLE IF EXISTS `meat_processing_parties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_processing_parties` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `approval_number` varchar(80) DEFAULT '',
  `meat_processing_party_identification_type_code` varchar(80) DEFAULT '',
  `meat_processing_party_type_code` varchar(80) DEFAULT '',
  `transactional_party_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_processing_parties`
--

LOCK TABLES `meat_processing_parties` WRITE;
/*!40000 ALTER TABLE `meat_processing_parties` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_processing_parties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_slaughtering_details`
--

DROP TABLE IF EXISTS `meat_slaughtering_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_slaughtering_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `age_of_animal` int(10) unsigned DEFAULT 0,
  `fat_content_percent` double DEFAULT 0,
  `fat_cover_code` varchar(80) DEFAULT '',
  `meat_category_code` varchar(80) DEFAULT '',
  `meat_colour_code` varchar(80) DEFAULT '',
  `meat_conformation_code` varchar(80) DEFAULT '',
  `meat_profile_code` varchar(80) DEFAULT '',
  `slaughtering_system_code` varchar(80) DEFAULT '',
  `slaughtering_weight` double DEFAULT 0,
  `sw_code_list_version` varchar(35) DEFAULT '',
  `sw_measurement_unit_code` varchar(10) DEFAULT NULL,
  `bse_test_id` int(10) unsigned DEFAULT 0,
  `slaughtering_dna_test_id` int(10) unsigned DEFAULT 0,
  `date_of_slaughtering` datetime DEFAULT current_timestamp(),
  `optimum_maturation_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_slaughtering_details`
--

LOCK TABLES `meat_slaughtering_details` WRITE;
/*!40000 ALTER TABLE `meat_slaughtering_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_slaughtering_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_tests`
--

DROP TABLE IF EXISTS `meat_tests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_tests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `test_method` varchar(80) DEFAULT '',
  `test_result` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_tests`
--

LOCK TABLES `meat_tests` WRITE;
/*!40000 ALTER TABLE `meat_tests` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_tests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meat_work_item_identifications`
--

DROP TABLE IF EXISTS `meat_work_item_identifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meat_work_item_identifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `batch_number` varchar(80) DEFAULT '',
  `livestock_mob_identifier` varchar(80) DEFAULT '',
  `meat_work_item_type_code` varchar(80) DEFAULT '',
  `animal_identification_id` int(10) unsigned DEFAULT 0,
  `product_identification` int(10) unsigned DEFAULT 0,
  `slaughter_number_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meat_work_item_identifications`
--

LOCK TABLES `meat_work_item_identifications` WRITE;
/*!40000 ALTER TABLE `meat_work_item_identifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `meat_work_item_identifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `object_events`
--

DROP TABLE IF EXISTS `object_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `object_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_id` varchar(120) DEFAULT '',
  `event_time_zone_offset` varchar(80) DEFAULT '',
  `certification` varchar(80) DEFAULT '',
  `event_time` datetime DEFAULT current_timestamp(),
  `reason` varchar(80) DEFAULT '',
  `declaration_time` datetime DEFAULT current_timestamp(),
  `action` varchar(80) DEFAULT '',
  `biz_step` varchar(80) DEFAULT '',
  `disposition` varchar(80) DEFAULT '',
  `read_point` varchar(80) DEFAULT '',
  `biz_location` varchar(80) DEFAULT '',
  `ilmd` varchar(80) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `object_events`
--

LOCK TABLES `object_events` WRITE;
/*!40000 ALTER TABLE `object_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `object_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operating_hours`
--

DROP TABLE IF EXISTS `operating_hours`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `operating_hours` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `day_of_the_week_code` varchar(50) DEFAULT '',
  `is_operational` tinyint(1) DEFAULT 0,
  `logistic_location_id` int(10) unsigned DEFAULT 0,
  `closing_time` datetime DEFAULT current_timestamp(),
  `opening_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operating_hours`
--

LOCK TABLES `operating_hours` WRITE;
/*!40000 ALTER TABLE `operating_hours` DISABLE KEYS */;
/*!40000 ALTER TABLE `operating_hours` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_line_item_details`
--

DROP TABLE IF EXISTS `order_line_item_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_line_item_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `requested_quantity` double DEFAULT 0,
  `rq_measurement_unit_code` varchar(80) DEFAULT '',
  `rq_code_list_version` varchar(35) DEFAULT '',
  `order_line_item_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_line_item_details`
--

LOCK TABLES `order_line_item_details` WRITE;
/*!40000 ALTER TABLE `order_line_item_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_line_item_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_line_items`
--

DROP TABLE IF EXISTS `order_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `extension` varchar(80) DEFAULT '',
  `free_goods_quantity` double DEFAULT 0,
  `fgq_measurement_unit_code` varchar(80) DEFAULT '',
  `fgq_code_list_version` varchar(35) DEFAULT '',
  `item_price_base_quantity` double DEFAULT 0,
  `ipbq_measurement_unit_code` varchar(80) DEFAULT '',
  `ipbq_code_list_version` varchar(35) DEFAULT '',
  `item_source_code` varchar(80) DEFAULT '',
  `line_item_action_code` varchar(80) DEFAULT '',
  `line_item_number` int(10) unsigned DEFAULT 0,
  `list_price` double DEFAULT 0,
  `lp_code_list_version` varchar(35) DEFAULT '',
  `lp_currency_code` varchar(80) DEFAULT '',
  `monetary_amount_excluding_taxes` double DEFAULT 0,
  `maet_code_list_version` varchar(35) DEFAULT '',
  `maet_currency_code` varchar(80) DEFAULT '',
  `monetary_amount_including_taxes` double DEFAULT 0,
  `mait_code_list_version` varchar(35) DEFAULT '',
  `mait_currency_code` varchar(80) DEFAULT '',
  `net_amount` double DEFAULT 0,
  `na_code_list_version` varchar(35) DEFAULT '',
  `na_currency_code` varchar(80) DEFAULT '',
  `net_price` double DEFAULT 0,
  `np_code_list_version` varchar(35) DEFAULT '',
  `np_currency_code` varchar(80) DEFAULT '',
  `order_instruction_code` varchar(80) DEFAULT '',
  `order_line_item_instruction_code` varchar(80) DEFAULT '',
  `order_line_item_priority` varchar(80) DEFAULT '',
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `recommended_retail_price` double DEFAULT 0,
  `requested_quantity` double DEFAULT 0,
  `rq_measurement_unit_code` varchar(80) DEFAULT '',
  `rq_code_list_version` varchar(35) DEFAULT '',
  `return_reason_code` varchar(80) DEFAULT '',
  `contract` int(10) unsigned DEFAULT 0,
  `customer_document_reference` int(10) unsigned DEFAULT 0,
  `delivery_date_according_to_schedule` int(10) unsigned DEFAULT 0,
  `despatch_advice` int(10) unsigned DEFAULT 0,
  `material_specification` int(10) unsigned DEFAULT 0,
  `order_line_item_contact` int(10) unsigned DEFAULT 0,
  `preferred_manufacturer` int(10) unsigned DEFAULT 0,
  `promotional_deal` int(10) unsigned DEFAULT 0,
  `purchase_conditions` int(10) unsigned DEFAULT 0,
  `returnable_asset_identification` int(10) unsigned DEFAULT 0,
  `order_id` int(10) unsigned DEFAULT 0,
  `latest_delivery_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_line_items`
--

LOCK TABLES `order_line_items` WRITE;
/*!40000 ALTER TABLE `order_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_logistical_date_informations`
--

DROP TABLE IF EXISTS `order_logistical_date_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_logistical_date_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_response_id` int(10) unsigned DEFAULT 0,
  `requested_delivery_date_range_begin` datetime DEFAULT current_timestamp(),
  `requested_delivery_date_range_end` datetime DEFAULT current_timestamp(),
  `requested_delivery_date_range_at_ultimate_consignee_begin` datetime DEFAULT current_timestamp(),
  `requested_delivery_date_range_at_ultimate_consignee_end` datetime DEFAULT current_timestamp(),
  `requested_delivery_date_time` datetime DEFAULT current_timestamp(),
  `requested_delivery_date_time_at_ultimate_consignee` datetime DEFAULT current_timestamp(),
  `requested_pick_up_date_time` datetime DEFAULT current_timestamp(),
  `requested_ship_date_range_begin` datetime DEFAULT current_timestamp(),
  `requested_ship_date_range_end` datetime DEFAULT current_timestamp(),
  `requested_ship_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_logistical_date_informations`
--

LOCK TABLES `order_logistical_date_informations` WRITE;
/*!40000 ALTER TABLE `order_logistical_date_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_logistical_date_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_logistical_informations`
--

DROP TABLE IF EXISTS `order_logistical_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_logistical_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `commodity_type_code` varchar(80) DEFAULT '',
  `shipment_split_method_code` varchar(80) DEFAULT '',
  `intermediate_delivery_party` int(10) unsigned DEFAULT 0,
  `inventory_location` int(10) unsigned DEFAULT 0,
  `ship_from` int(10) unsigned DEFAULT 0,
  `ship_to` int(10) unsigned DEFAULT 0,
  `ultimate_consignee` int(10) unsigned DEFAULT 0,
  `order_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_logistical_informations`
--

LOCK TABLES `order_logistical_informations` WRITE;
/*!40000 ALTER TABLE `order_logistical_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_logistical_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_response_line_item_details`
--

DROP TABLE IF EXISTS `order_response_line_item_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_response_line_item_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `confirmed_quantity` double DEFAULT 0,
  `cq_measurement_unit_code` varchar(80) DEFAULT '',
  `cq_code_list_version` varchar(35) DEFAULT '',
  `return_reason_code` varchar(80) DEFAULT '',
  `order_response_line_item_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_response_line_item_details`
--

LOCK TABLES `order_response_line_item_details` WRITE;
/*!40000 ALTER TABLE `order_response_line_item_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_response_line_item_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_response_line_items`
--

DROP TABLE IF EXISTS `order_response_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_response_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `confirmed_quantity` double DEFAULT 0,
  `cq_measurement_unit_code` varchar(80) DEFAULT '',
  `cq_code_list_version` varchar(35) DEFAULT '',
  `line_item_action_code` varchar(80) DEFAULT '',
  `line_item_change_indicator` varchar(80) DEFAULT '',
  `line_item_number` int(10) unsigned DEFAULT 0,
  `monetary_amount_excluding_taxes` double DEFAULT 0,
  `maet_code_list_version` varchar(35) DEFAULT '',
  `maet_currency_code` varchar(80) DEFAULT '',
  `monetary_amount_including_taxes` double DEFAULT 0,
  `mait_code_list_version` varchar(35) DEFAULT '',
  `mait_currency_code` varchar(80) DEFAULT '',
  `net_amount` double DEFAULT 0,
  `na_code_list_version` varchar(35) DEFAULT '',
  `na_currency_code` varchar(80) DEFAULT '',
  `net_price` double DEFAULT 0,
  `np_code_list_version` varchar(35) DEFAULT '',
  `np_currency_code` varchar(80) DEFAULT '',
  `order_response_reason_code` varchar(80) DEFAULT '',
  `original_order_line_item_number` int(10) unsigned DEFAULT 0,
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `order_response_id` int(10) unsigned DEFAULT 0,
  `delivery_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_response_line_items`
--

LOCK TABLES `order_response_line_items` WRITE;
/*!40000 ALTER TABLE `order_response_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_response_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_responses`
--

DROP TABLE IF EXISTS `order_responses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_responses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `order_response_reason_code` varchar(80) DEFAULT '',
  `response_status_code` varchar(80) DEFAULT '',
  `total_monetary_amount_excluding_taxes` double DEFAULT 0,
  `tmaet_code_list_version` varchar(35) DEFAULT '',
  `tmaet_currency_code` varchar(80) DEFAULT '',
  `total_monetary_amount_including_taxes` double DEFAULT 0,
  `tmait_code_list_version` varchar(35) DEFAULT '',
  `tmait_currency_code` varchar(80) DEFAULT '',
  `total_tax_amount` double DEFAULT 0,
  `tta_code_list_version` varchar(35) DEFAULT '',
  `tta_currency_code` varchar(80) DEFAULT '',
  `amended_date_time_value` int(10) unsigned DEFAULT 0,
  `bill_to` int(10) unsigned DEFAULT 0,
  `buyer` int(10) unsigned DEFAULT 0,
  `order_response_identification` int(10) unsigned DEFAULT 0,
  `original_order` int(10) unsigned DEFAULT 0,
  `sales_order` int(10) unsigned DEFAULT 0,
  `seller` int(10) unsigned DEFAULT 0,
  `ship_to` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_responses`
--

LOCK TABLES `order_responses` WRITE;
/*!40000 ALTER TABLE `order_responses` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_responses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `orders` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `is_application_receipt_acknowledgement_required` tinyint(1) DEFAULT 0,
  `is_order_free_of_excise_tax_duty` tinyint(1) DEFAULT 0,
  `order_change_reason_code` varchar(80) DEFAULT '',
  `order_entry_type` varchar(80) DEFAULT '',
  `order_instruction_code` varchar(80) DEFAULT '',
  `order_priority` varchar(50) DEFAULT '',
  `order_type_code` varchar(80) DEFAULT '',
  `total_monetary_amount_excluding_taxes` double DEFAULT 0,
  `tmaet_code_list_version` varchar(35) DEFAULT '',
  `tmaet_currency_code` varchar(80) DEFAULT '',
  `total_monetary_amount_including_taxes` double DEFAULT 0,
  `tmait_code_list_version` varchar(35) DEFAULT '',
  `tmait_currency_code` varchar(80) DEFAULT '',
  `total_tax_amount` double DEFAULT 0,
  `tta_code_list_version` varchar(35) DEFAULT '',
  `tta_currency_code` varchar(80) DEFAULT '',
  `bill_to` int(10) unsigned DEFAULT 0,
  `buyer` int(10) unsigned DEFAULT 0,
  `contract` int(10) unsigned DEFAULT 0,
  `customer_document_reference` int(10) unsigned DEFAULT 0,
  `customs_broker` int(10) unsigned DEFAULT 0,
  `order_identification` int(10) unsigned DEFAULT 0,
  `pickup_from` int(10) unsigned DEFAULT 0,
  `promotional_deal` int(10) unsigned DEFAULT 0,
  `quote_number` varchar(35) DEFAULT '',
  `seller` int(10) unsigned DEFAULT 0,
  `trade_agreement` varchar(35) DEFAULT '',
  `delivery_date_according_to_schedule` datetime DEFAULT current_timestamp(),
  `latest_delivery_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `organisations`
--

DROP TABLE IF EXISTS `organisations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `organisations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `issued_capital` double DEFAULT 0,
  `ic_code_list_version` varchar(35) DEFAULT '',
  `ic_currency_code` varchar(80) DEFAULT '',
  `organisation_name` varchar(80) DEFAULT '',
  `official_address` int(10) unsigned DEFAULT 0,
  `transactional_party_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `organisations`
--

LOCK TABLES `organisations` WRITE;
/*!40000 ALTER TABLE `organisations` DISABLE KEYS */;
/*!40000 ALTER TABLE `organisations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `package_totals`
--

DROP TABLE IF EXISTS `package_totals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `package_totals` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `package_type_code` varchar(80) DEFAULT '',
  `total_gross_volume` varchar(80) DEFAULT '',
  `tgv_code_list_version` varchar(35) DEFAULT '',
  `tg_vmeasurement_unit_code` varchar(80) DEFAULT '',
  `total_gross_weight` varchar(80) DEFAULT '',
  `tgw_code_list_version` varchar(35) DEFAULT '',
  `tgw_measurement_unit_code` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `package_totals`
--

LOCK TABLES `package_totals` WRITE;
/*!40000 ALTER TABLE `package_totals` DISABLE KEYS */;
/*!40000 ALTER TABLE `package_totals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_methods`
--

DROP TABLE IF EXISTS `payment_methods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_methods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `automated_clearing_house_payment_format` varchar(50) DEFAULT '',
  `payment_method_code` varchar(50) DEFAULT '',
  `payment_method_identification` varchar(50) DEFAULT '',
  `payment_term_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_methods`
--

LOCK TABLES `payment_methods` WRITE;
/*!40000 ALTER TABLE `payment_methods` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_methods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_terms`
--

DROP TABLE IF EXISTS `payment_terms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_terms` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `payment_terms_event_code` varchar(50) DEFAULT '',
  `payment_terms_type_code` varchar(50) DEFAULT '',
  `proximo_cut_off_day` varchar(50) DEFAULT '',
  `order_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_terms`
--

LOCK TABLES `payment_terms` WRITE;
/*!40000 ALTER TABLE `payment_terms` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_terms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_terms_discounts`
--

DROP TABLE IF EXISTS `payment_terms_discounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_terms_discounts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `discount_amount` double DEFAULT 0,
  `da_code_list_version` varchar(35) DEFAULT '',
  `da_currency_code` varchar(80) DEFAULT '',
  `discount_percent` double DEFAULT 0,
  `discount_type` varchar(50) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_terms_discounts`
--

LOCK TABLES `payment_terms_discounts` WRITE;
/*!40000 ALTER TABLE `payment_terms_discounts` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_terms_discounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_time_periods`
--

DROP TABLE IF EXISTS `payment_time_periods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_time_periods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `day_of_month_due` varchar(50) DEFAULT '',
  `payment_term_id` int(10) unsigned DEFAULT 0,
  `payment_term_discount_id` int(10) unsigned DEFAULT 0,
  `date_due` datetime DEFAULT current_timestamp(),
  `time_period_due` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_time_periods`
--

LOCK TABLES `payment_time_periods` WRITE;
/*!40000 ALTER TABLE `payment_time_periods` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_time_periods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `persistent_dispositions`
--

DROP TABLE IF EXISTS `persistent_dispositions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `persistent_dispositions` (
  `set_disp` varchar(80) DEFAULT '',
  `unset_disp` varchar(80) DEFAULT '',
  `event_id` int(10) unsigned DEFAULT NULL,
  `type_of_event` varchar(80) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `persistent_dispositions`
--

LOCK TABLES `persistent_dispositions` WRITE;
/*!40000 ALTER TABLE `persistent_dispositions` DISABLE KEYS */;
/*!40000 ALTER TABLE `persistent_dispositions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `quantity_elements`
--

DROP TABLE IF EXISTS `quantity_elements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `quantity_elements` (
  `epc_class` varchar(80) DEFAULT '',
  `quantity` double DEFAULT 0,
  `uom` varchar(80) DEFAULT '',
  `event_id` int(10) unsigned DEFAULT NULL,
  `type_of_event` varchar(80) DEFAULT '',
  `type_of_quantity` varchar(80) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `quantity_elements`
--

LOCK TABLES `quantity_elements` WRITE;
/*!40000 ALTER TABLE `quantity_elements` DISABLE KEYS */;
/*!40000 ALTER TABLE `quantity_elements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `receiving_advice_line_items`
--

DROP TABLE IF EXISTS `receiving_advice_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `receiving_advice_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `line_item_number` int(10) unsigned DEFAULT 0,
  `parent_line_item_number` int(10) unsigned DEFAULT 0,
  `quantity_accepted` double DEFAULT 0,
  `qa_measurement_unit_code` varchar(80) DEFAULT '',
  `qa_code_list_version` varchar(35) DEFAULT '',
  `quantity_despatched` double DEFAULT 0,
  `qd_measurement_unit_code` varchar(80) DEFAULT '',
  `qd_code_list_version` varchar(35) DEFAULT '',
  `quantity_received` double DEFAULT 0,
  `qr_measurement_unit_code` varchar(80) DEFAULT '',
  `qr_code_list_version` varchar(35) DEFAULT '',
  `transactional_trade_item` int(10) unsigned DEFAULT 0,
  `ecom_consignment_identification` int(10) unsigned DEFAULT 0,
  `contract` int(10) unsigned DEFAULT 0,
  `customer_reference` int(10) unsigned DEFAULT 0,
  `delivery_note` int(10) unsigned DEFAULT 0,
  `despatch_advice` int(10) unsigned DEFAULT 0,
  `product_certification` int(10) unsigned DEFAULT 0,
  `promotional_deal` int(10) unsigned DEFAULT 0,
  `purchase_conditions` int(10) unsigned DEFAULT 0,
  `purchase_order` int(10) unsigned DEFAULT 0,
  `requested_item_identification` int(10) unsigned DEFAULT 0,
  `specification` int(10) unsigned DEFAULT 0,
  `receiving_advice_id` int(10) unsigned DEFAULT 0,
  `pick_up_date_time_begin` datetime DEFAULT current_timestamp(),
  `pick_up_date_time_end` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `receiving_advice_line_items`
--

LOCK TABLES `receiving_advice_line_items` WRITE;
/*!40000 ALTER TABLE `receiving_advice_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `receiving_advice_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `receiving_advices`
--

DROP TABLE IF EXISTS `receiving_advices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `receiving_advices` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `reporting_code` varchar(80) DEFAULT '',
  `total_accepted_amount` double DEFAULT 0,
  `taa_code_list_version` varchar(35) DEFAULT '',
  `taa_currency_code` varchar(80) DEFAULT '',
  `total_deposit_amount` double DEFAULT 0,
  `tda_code_list_version` varchar(35) DEFAULT '',
  `tda_currency_code` varchar(80) DEFAULT '',
  `total_number_of_lines` int(10) unsigned DEFAULT 0,
  `total_on_hold_amount` double DEFAULT 0,
  `toha_code_list_version` varchar(35) DEFAULT '',
  `toha_currency_code` varchar(80) DEFAULT '',
  `total_rejected_amount` double DEFAULT 0,
  `tra_code_list_version` varchar(35) DEFAULT '',
  `tra_currency_code` varchar(80) DEFAULT '',
  `receiving_advice_transport_information` int(10) unsigned DEFAULT 0,
  `bill_of_lading_number` int(10) unsigned DEFAULT 0,
  `buyer` int(10) unsigned DEFAULT 0,
  `carrier` int(10) unsigned DEFAULT 0,
  `consignment_identification` int(10) unsigned DEFAULT 0,
  `delivery_note` int(10) unsigned DEFAULT 0,
  `despatch_advice` int(10) unsigned DEFAULT 0,
  `inventory_location` int(10) unsigned DEFAULT 0,
  `purchase_order` int(10) unsigned DEFAULT 0,
  `receiver` int(10) unsigned DEFAULT 0,
  `receiving_advice_identification` int(10) unsigned DEFAULT 0,
  `seller` int(10) unsigned DEFAULT 0,
  `ship_from` int(10) unsigned DEFAULT 0,
  `shipment_identification` int(10) unsigned DEFAULT 0,
  `shipper` int(10) unsigned DEFAULT 0,
  `ship_to` int(10) unsigned DEFAULT 0,
  `despatch_advice_delivery_date_time_begin` datetime DEFAULT current_timestamp(),
  `despatch_advice_delivery_date_time_end` datetime DEFAULT current_timestamp(),
  `payment_date_time_begin` datetime DEFAULT current_timestamp(),
  `payment_date_time_end` datetime DEFAULT current_timestamp(),
  `receiving_date_time_begin` datetime DEFAULT current_timestamp(),
  `receiving_date_time_end` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `receiving_advices`
--

LOCK TABLES `receiving_advices` WRITE;
/*!40000 ALTER TABLE `receiving_advices` DISABLE KEYS */;
/*!40000 ALTER TABLE `receiving_advices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `returnable_packaging_inventory_events`
--

DROP TABLE IF EXISTS `returnable_packaging_inventory_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `returnable_packaging_inventory_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_identifier` varchar(80) DEFAULT '',
  `inventory_business_step_code` varchar(80) DEFAULT '',
  `inventory_event_reason_code` varchar(80) DEFAULT '',
  `inventory_movement_type_code` varchar(80) DEFAULT '',
  `inventory_sub_location_id` int(10) unsigned DEFAULT 0,
  `event_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `returnable_packaging_inventory_events`
--

LOCK TABLES `returnable_packaging_inventory_events` WRITE;
/*!40000 ALTER TABLE `returnable_packaging_inventory_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `returnable_packaging_inventory_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `returnable_packagings`
--

DROP TABLE IF EXISTS `returnable_packagings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `returnable_packagings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `current_holder_registration` int(10) unsigned DEFAULT 0,
  `new_holder_registration` int(10) unsigned DEFAULT 0,
  `packaging_condition_code` varchar(80) DEFAULT '',
  `packaging_quantity` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `returnable_packagings`
--

LOCK TABLES `returnable_packagings` WRITE;
/*!40000 ALTER TABLE `returnable_packagings` DISABLE KEYS */;
/*!40000 ALTER TABLE `returnable_packagings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sensor_elements`
--

DROP TABLE IF EXISTS `sensor_elements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sensor_elements` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `device_id` varchar(80) DEFAULT '',
  `device_metadata` varchar(80) DEFAULT '',
  `raw_data` varchar(80) DEFAULT '',
  `data_processing_method` varchar(80) DEFAULT '',
  `biz_rules` varchar(80) DEFAULT '',
  `sensor_time` datetime DEFAULT current_timestamp(),
  `start_time` datetime DEFAULT current_timestamp(),
  `end_time` datetime DEFAULT current_timestamp(),
  `event_id` int(10) unsigned DEFAULT NULL,
  `type_of_event` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sensor_elements`
--

LOCK TABLES `sensor_elements` WRITE;
/*!40000 ALTER TABLE `sensor_elements` DISABLE KEYS */;
/*!40000 ALTER TABLE `sensor_elements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sensor_reports`
--

DROP TABLE IF EXISTS `sensor_reports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sensor_reports` (
  `sensor_report_type` varchar(80) DEFAULT '',
  `device_id` varchar(80) DEFAULT '',
  `raw_data` varchar(80) DEFAULT '',
  `data_processing_method` varchar(80) DEFAULT '',
  `microorganism` varchar(80) DEFAULT '',
  `chemical_substance` varchar(80) DEFAULT '',
  `sensor_value` double DEFAULT 0,
  `component` varchar(80) DEFAULT '',
  `string_value` varchar(80) DEFAULT '',
  `boolean_value` tinyint(1) DEFAULT 0,
  `hex_binary_value` varchar(80) DEFAULT '',
  `uri_value` varchar(80) DEFAULT '',
  `min_value` double DEFAULT 0,
  `max_value` double DEFAULT 0,
  `mean_value` double DEFAULT 0,
  `perc_rank` double DEFAULT 0,
  `perc_value` double DEFAULT 0,
  `uom` varchar(80) DEFAULT '',
  `s_dev` double DEFAULT 0,
  `device_metadata` varchar(80) DEFAULT '',
  `sensor_element_id` int(10) unsigned DEFAULT NULL,
  `sensor_report_time` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sensor_reports`
--

LOCK TABLES `sensor_reports` WRITE;
/*!40000 ALTER TABLE `sensor_reports` DISABLE KEYS */;
/*!40000 ALTER TABLE `sensor_reports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_transportation_informations`
--

DROP TABLE IF EXISTS `shipment_transportation_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_transportation_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `route_id` int(10) unsigned DEFAULT 0,
  `transport_means_id` int(10) unsigned DEFAULT 0,
  `transport_means_type` varchar(80) DEFAULT '',
  `transport_service_category_type` varchar(80) DEFAULT '',
  `transport_service_level_code` varchar(80) DEFAULT '',
  `carrier` int(10) unsigned DEFAULT 0,
  `freight_forwarder` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_transportation_informations`
--

LOCK TABLES `shipment_transportation_informations` WRITE;
/*!40000 ALTER TABLE `shipment_transportation_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_transportation_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sources`
--

DROP TABLE IF EXISTS `sources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sources` (
  `source_type` varchar(80) DEFAULT '',
  `source` varchar(80) DEFAULT '',
  `event_id` int(10) unsigned DEFAULT NULL,
  `type_of_event` varchar(80) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sources`
--

LOCK TABLES `sources` WRITE;
/*!40000 ALTER TABLE `sources` DISABLE KEYS */;
/*!40000 ALTER TABLE `sources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `trade_item_identifications`
--

DROP TABLE IF EXISTS `trade_item_identifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trade_item_identifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `additional_trade_item_identification` varchar(80) DEFAULT '',
  `additional_trade_item_identification_type_code` varchar(80) DEFAULT '',
  `code_list_version` varchar(35) DEFAULT '',
  `gtin` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trade_item_identifications`
--

LOCK TABLES `trade_item_identifications` WRITE;
/*!40000 ALTER TABLE `trade_item_identifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `trade_item_identifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `trade_item_inventory_events`
--

DROP TABLE IF EXISTS `trade_item_inventory_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trade_item_inventory_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_identifier` varchar(80) DEFAULT '',
  `inventory_business_step_code` varchar(80) DEFAULT '',
  `inventory_disposition_code` varchar(80) DEFAULT '',
  `inventory_event_reason_code` varchar(80) DEFAULT '',
  `inventory_movement_type_code` varchar(80) DEFAULT '',
  `inventory_sub_location_id` int(10) unsigned DEFAULT 0,
  `liable_party` int(10) unsigned DEFAULT 0,
  `event_date_time` datetime DEFAULT current_timestamp(),
  `logistics_inventory_report_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trade_item_inventory_events`
--

LOCK TABLES `trade_item_inventory_events` WRITE;
/*!40000 ALTER TABLE `trade_item_inventory_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `trade_item_inventory_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `trade_item_inventory_statuses`
--

DROP TABLE IF EXISTS `trade_item_inventory_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trade_item_inventory_statuses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inventory_disposition_code` varchar(80) DEFAULT '',
  `inventory_sub_location_id` int(10) unsigned DEFAULT 0,
  `inventory_date_time` datetime DEFAULT current_timestamp(),
  `logistics_inventory_report_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trade_item_inventory_statuses`
--

LOCK TABLES `trade_item_inventory_statuses` WRITE;
/*!40000 ALTER TABLE `trade_item_inventory_statuses` DISABLE KEYS */;
/*!40000 ALTER TABLE `trade_item_inventory_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction_events`
--

DROP TABLE IF EXISTS `transaction_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transaction_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_id` varchar(120) DEFAULT '',
  `event_time_zone_offset` varchar(80) DEFAULT '',
  `certification` varchar(80) DEFAULT '',
  `event_time` datetime DEFAULT current_timestamp(),
  `reason` varchar(80) DEFAULT '',
  `declaration_time` datetime DEFAULT current_timestamp(),
  `parent_id` varchar(80) DEFAULT '',
  `action` varchar(80) DEFAULT '',
  `biz_step` varchar(80) DEFAULT '',
  `disposition` varchar(80) DEFAULT '',
  `read_point` varchar(80) DEFAULT '',
  `biz_location` varchar(80) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction_events`
--

LOCK TABLES `transaction_events` WRITE;
/*!40000 ALTER TABLE `transaction_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `transaction_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactional_item_certifications`
--

DROP TABLE IF EXISTS `transactional_item_certifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactional_item_certifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `item_certification_agency` varchar(35) DEFAULT '',
  `item_certification_standard` varchar(35) DEFAULT '',
  `item_certification_value` varchar(35) DEFAULT '',
  `organic_certification_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactional_item_certifications`
--

LOCK TABLES `transactional_item_certifications` WRITE;
/*!40000 ALTER TABLE `transactional_item_certifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactional_item_certifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactional_item_data`
--

DROP TABLE IF EXISTS `transactional_item_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactional_item_data` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `batch_number` varchar(35) DEFAULT '',
  `country_of_origin` varchar(35) DEFAULT '',
  `item_in_contact_with_food_product` tinyint(1) DEFAULT 0,
  `lot_number` varchar(35) DEFAULT '',
  `product_quality_indication` int(10) unsigned DEFAULT 0,
  `serial_number` varchar(35) DEFAULT '',
  `shelf_life` varchar(35) DEFAULT '',
  `trade_item_quantity` int(10) unsigned DEFAULT 0,
  `available_for_sale_date` datetime DEFAULT current_timestamp(),
  `best_before_date` datetime DEFAULT current_timestamp(),
  `item_expiration_date` datetime DEFAULT current_timestamp(),
  `packaging_date` datetime DEFAULT current_timestamp(),
  `production_date` datetime DEFAULT current_timestamp(),
  `sell_by_date` datetime DEFAULT current_timestamp(),
  `trade_item_inventory_status_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactional_item_data`
--

LOCK TABLES `transactional_item_data` WRITE;
/*!40000 ALTER TABLE `transactional_item_data` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactional_item_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactional_item_data_carrier_and_identifications`
--

DROP TABLE IF EXISTS `transactional_item_data_carrier_and_identifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactional_item_data_carrier_and_identifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `data_carrier` varchar(35) DEFAULT '',
  `gs1_transactional_item_identification_key` varchar(35) DEFAULT '',
  `transactional_item_data_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactional_item_data_carrier_and_identifications`
--

LOCK TABLES `transactional_item_data_carrier_and_identifications` WRITE;
/*!40000 ALTER TABLE `transactional_item_data_carrier_and_identifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactional_item_data_carrier_and_identifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactional_item_logistic_unit_informations`
--

DROP TABLE IF EXISTS `transactional_item_logistic_unit_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactional_item_logistic_unit_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `maximum_stacking_factor` int(10) unsigned DEFAULT 0,
  `number_of_layers` int(10) unsigned DEFAULT 0,
  `number_of_units_per_layer` int(10) unsigned DEFAULT 0,
  `number_of_units_per_pallet` int(10) unsigned DEFAULT 0,
  `package_type_code` varchar(35) DEFAULT '',
  `packaging_terms` varchar(35) DEFAULT '',
  `returnable_package_transport_cost_payment` varchar(35) DEFAULT '',
  `transactional_item_data_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactional_item_logistic_unit_informations`
--

LOCK TABLES `transactional_item_logistic_unit_informations` WRITE;
/*!40000 ALTER TABLE `transactional_item_logistic_unit_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactional_item_logistic_unit_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactional_item_organic_informations`
--

DROP TABLE IF EXISTS `transactional_item_organic_informations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactional_item_organic_informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `is_trade_item_organic` tinyint(1) DEFAULT 0,
  `transactional_item_data_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactional_item_organic_informations`
--

LOCK TABLES `transactional_item_organic_informations` WRITE;
/*!40000 ALTER TABLE `transactional_item_organic_informations` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactional_item_organic_informations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactional_parties`
--

DROP TABLE IF EXISTS `transactional_parties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactional_parties` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `gln` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactional_parties`
--

LOCK TABLES `transactional_parties` WRITE;
/*!40000 ALTER TABLE `transactional_parties` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactional_parties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactional_references`
--

DROP TABLE IF EXISTS `transactional_references`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactional_references` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `transactional_reference_type_code` varchar(80) DEFAULT '',
  `ecom_document_reference` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactional_references`
--

LOCK TABLES `transactional_references` WRITE;
/*!40000 ALTER TABLE `transactional_references` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactional_references` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transformation_events`
--

DROP TABLE IF EXISTS `transformation_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transformation_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_id` varchar(120) DEFAULT '',
  `event_time_zone_offset` varchar(80) DEFAULT '',
  `certification` varchar(80) DEFAULT '',
  `event_time` datetime DEFAULT current_timestamp(),
  `reason` varchar(80) DEFAULT '',
  `declaration_time` datetime DEFAULT current_timestamp(),
  `transformation_id` varchar(80) DEFAULT '',
  `biz_step` varchar(80) DEFAULT '',
  `disposition` varchar(80) DEFAULT '',
  `read_point` varchar(80) DEFAULT '',
  `biz_location` varchar(80) DEFAULT '',
  `ilmd` varchar(80) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transformation_events`
--

LOCK TABLES `transformation_events` WRITE;
/*!40000 ALTER TABLE `transformation_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `transformation_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_equipment_inventory_events`
--

DROP TABLE IF EXISTS `transport_equipment_inventory_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_equipment_inventory_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_identifier` varchar(80) DEFAULT '',
  `inventory_business_step_code` varchar(80) DEFAULT '',
  `inventory_disposition_code` varchar(80) DEFAULT '',
  `inventory_event_reason_code` varchar(80) DEFAULT '',
  `inventory_movement_type_code` varchar(80) DEFAULT '',
  `number_of_pieces_of_equipment` int(10) unsigned DEFAULT 0,
  `inventory_sub_location_id` int(10) unsigned DEFAULT 0,
  `event_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_equipment_inventory_events`
--

LOCK TABLES `transport_equipment_inventory_events` WRITE;
/*!40000 ALTER TABLE `transport_equipment_inventory_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_equipment_inventory_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_equipment_inventory_statuses`
--

DROP TABLE IF EXISTS `transport_equipment_inventory_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_equipment_inventory_statuses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `inventory_disposition_code` varchar(80) DEFAULT '',
  `number_of_pieces_of_equipment` int(10) unsigned DEFAULT 0,
  `inventory_sub_location_id` int(10) unsigned DEFAULT 0,
  `inventory_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_equipment_inventory_statuses`
--

LOCK TABLES `transport_equipment_inventory_statuses` WRITE;
/*!40000 ALTER TABLE `transport_equipment_inventory_statuses` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_equipment_inventory_statuses` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-22 15:08:12
