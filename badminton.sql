-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 10, 2025 at 05:46 PM
-- Server version: 8.0.30
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `badminton`
--

-- --------------------------------------------------------

--
-- Table structure for table `kehadirans`
--

CREATE TABLE `kehadirans` (
  `id` int NOT NULL,
  `tanggal` longtext,
  `nama` longtext,
  `npm` longtext,
  `prodi` longtext,
  `status` longtext,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `kehadirans`
--

INSERT INTO `kehadirans` (`id`, `tanggal`, `nama`, `npm`, `prodi`, `status`, `created_at`) VALUES
(2, '2025-07-09', 'Budi Santoso', '2105011002', 'informatika', 'Hadir', '2025-07-09 17:01:33.000'),
(3, '2025-07-03', 'Dewi Lestari', '2105011003', 'Teknik Informatika', 'Hadir', '2025-07-09 17:01:33.000'),
(5, '2025-07-09', 'Rafaeza', '23670080', 'Teknik Informatika', 'Tidak Hadir', '2025-07-09 23:59:08.834'),
(6, '2025-07-09', 'aditya', '23670064', 'manajemen', 'Hadir', '2025-07-10 00:13:32.075'),
(8, '2025-07-11', 'Erika Puji Suhartanti', '23670084', 'Teknik Informatika', 'Hadir', '2025-07-11 00:21:54.396');

-- --------------------------------------------------------

--
-- Table structure for table `keuangans`
--

CREATE TABLE `keuangans` (
  `id` int NOT NULL,
  `tanggal` datetime(3) DEFAULT NULL,
  `deskripsi` longtext,
  `tipe` longtext,
  `jumlah` double DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `keuangans`
--

INSERT INTO `keuangans` (`id`, `tanggal`, `deskripsi`, `tipe`, `jumlah`, `created_at`, `updated_at`) VALUES
(3, '2025-07-04 00:00:00.000', 'Donasi alumni', 'Pemasukan', 100000, '2025-07-09 17:02:04.000', NULL),
(4, '2025-07-05 07:00:00.000', 'membeli minum', 'Pengeluaran', 10000, '2025-07-09 17:02:04.000', '2025-07-11 00:23:42.302'),
(5, '2025-07-09 07:00:00.000', 'kas tanggal 9 kegiatan latihan rutin', 'Pemasukan', 15000, '2025-07-10 00:20:09.931', '2025-07-10 00:20:09.931'),
(6, '2025-07-01 07:00:00.000', 'BELI SUTTLECOCK', 'Pengeluaran', 80000, '2025-07-10 16:02:29.165', '2025-07-10 16:02:29.165'),
(7, '2025-07-01 07:00:00.000', 'kas', 'Pemasukan', 20000, '2025-07-10 18:57:17.913', '2025-07-10 18:57:17.913'),
(11, '2025-07-11 07:00:00.000', 'kas kegiatan latihan rutin', 'Pemasukan', 20000, '2025-07-11 00:23:07.651', '2025-07-11 00:23:07.651');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `username` varchar(191) DEFAULT NULL,
  `password` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`) VALUES
(1, 'admin', 'admin123');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `kehadirans`
--
ALTER TABLE `kehadirans`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `keuangans`
--
ALTER TABLE `keuangans`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `kehadirans`
--
ALTER TABLE `kehadirans`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `keuangans`
--
ALTER TABLE `keuangans`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
