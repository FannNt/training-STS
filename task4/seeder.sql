-- Insert Master Produk data
INSERT INTO master_produk (kode, nama_produk, kategori, brand) VALUES
('PRD001', 'iPhone 15 Pro', 'Smartphone', 'Apple'),
('PRD002', 'Samsung Galaxy S24', 'Smartphone', 'Samsung'),
('PRD003', 'MacBook Air M2', 'Laptop', 'Apple'),
('PRD004', 'Dell XPS 13', 'Laptop', 'Dell'),
('PRD005', 'Xiaomi Redmi Note 12', 'Smartphone', 'Xiaomi'),
('PRD006', 'ASUS ROG Strix', 'Laptop', 'ASUS'),
('PRD007', 'OnePlus 11', 'Smartphone', 'OnePlus'),
('PRD008', 'Lenovo ThinkPad X1', 'Laptop', 'Lenovo'),
('PRD009', 'Google Pixel 8', 'Smartphone', 'Google'),
('PRD010', 'HP Pavilion', 'Laptop', 'HP');

-- Insert Master Pelanggan data
INSERT INTO master_pelanggan (kode, nama, level_pelanggan, asal_kota) VALUES
('CUST001', 'Ahmad Wijaya', 'VIP', 'Jakarta'),
('CUST002', 'Siti Nurhaliza', 'Reguler', 'Bandung'),
('CUST003', 'Budi Santoso', 'VIP', 'Surabaya'),
('CUST004', 'Dewi Kartika', 'Reguler', 'Medan'),
('CUST005', 'Rizki Pratama', 'VIP', 'Yogyakarta'),
('CUST006', 'Maya Sari', 'Reguler', 'Semarang'),
('CUST007', 'Agus Setiawan', 'VIP', 'Makassar'),
('CUST008', 'Indah Permata', 'Reguler', 'Palembang'),
('CUST009', 'Fajar Nugroho', 'VIP', 'Denpasar'),
('CUST010', 'Lina Marlina', 'Reguler', 'Balikpapan');

-- Insert Master Harga Jual data
INSERT INTO master_harga_jual (kode_produk, harga, level_pelanggan, periode_awal, periode_akhir) VALUES
('PRD001', 15000000, 'VIP', '2024-01-01', '2024-12-31'),
('PRD001', 16000000, 'Reguler', '2024-01-01', '2024-12-31'),
('PRD002', 12000000, 'VIP', '2024-01-01', '2024-12-31'),
('PRD002', 13000000, 'Reguler', '2024-01-01', '2024-12-31'),
('PRD003', 18000000, 'VIP', '2024-01-01', '2024-12-31'),
('PRD003', 19000000, 'Reguler', '2024-01-01', '2024-12-31'),
('PRD004', 15000000, 'VIP', '2024-01-01', '2024-12-31'),
('PRD004', 16000000, 'Reguler', '2024-01-01', '2024-12-31'),
('PRD005', 3000000, 'VIP', '2024-01-01', '2024-12-31'),
('PRD005', 3200000, 'Reguler', '2024-01-01', '2024-12-31');

-- Insert Transaksi Penjualan data (inferred from Detail Transaksi)
INSERT INTO transaksi_penjualan (nomor_penjualan, tanggal, customer, nama_kasir) VALUES
('TRX001', '2024-01-15', 'CUST001', 'Kasir A'),
('TRX002', '2024-01-16', 'CUST005', 'Kasir B'),
('TRX003', '2024-01-17', 'CUST003', 'Kasir A'),
('TRX004', '2024-01-18', 'CUST002', 'Kasir C'),
('TRX005', '2024-01-19', 'CUST004', 'Kasir A'),
('TRX006', '2024-01-20', 'CUST007', 'Kasir B'),
('TRX007', '2024-01-21', 'CUST006', 'Kasir C'),
('TRX008', '2024-01-22', 'CUST010', 'Kasir A'),
('TRX009', '2024-01-23', 'CUST008', 'Kasir B'),
('TRX010', '2024-01-24', 'CUST003', 'Kasir C'),
('TRX011', '2024-01-25', 'CUST002', 'Kasir A'),
('TRX012', '2024-01-26', 'CUST001', 'Kasir B'),
('TRX013', '2024-01-27', 'CUST004', 'Kasir C'),
('TRX014', '2024-01-28', 'CUST006', 'Kasir A'),
('TRX015', '2024-01-29', 'CUST010', 'Kasir B'),
('TRX016', '2024-01-30', 'CUST003', 'Kasir C'),
('TRX017', '2024-02-01', 'CUST005', 'Kasir A'),
('TRX018', '2024-02-02', 'CUST008', 'Kasir B'),
('TRX019', '2024-02-03', 'CUST007', 'Kasir C'),
('TRX020', '2024-02-04', 'CUST004', 'Kasir A');

-- Insert Detail Transaksi data
INSERT INTO detail_transaksi (nomor_penjualan, urut_item, kode_produk, qty, harga, diskon, total_nilai) VALUES
('TRX001', 1, 'PRD001', 1, 15000000, 50000, 14500000),
('TRX002', 1, 'PRD005', 2, 3000000, 100000, 5900000),
('TRX003', 1, 'PRD003', 1, 18000000, 100000, 17000000),
('TRX003', 2, 'PRD001', 1, 15000000, 50000, 14500000),
('TRX004', 1, 'PRD002', 1, 12000000, 0, 12000000),
('TRX005', 1, 'PRD004', 1, 15000000, 75000, 14250000),
('TRX006', 1, 'PRD007', 1, 8000000, 200000, 7800000),
('TRX007', 1, 'PRD006', 1, 12000000, 500000, 11500000),
('TRX007', 2, 'PRD009', 1, 10000000, 30000, 9700000),
('TRX008', 1, 'PRD010', 1, 8000000, 0, 8000000),
('TRX009', 1, 'PRD008', 1, 14000000, 100000, 13000000),
('TRX010', 1, 'PRD003', 1, 19000000, 50000, 18500000),
('TRX011', 1, 'PRD002', 1, 13000000, 0, 13000000),
('TRX011', 2, 'PRD005', 1, 3200000, 200000, 3000000),
('TRX012', 1, 'PRD001', 1, 16000000, 100000, 15000000),
('TRX012', 2, 'PRD007', 1, 8000000, 0, 8000000),
('TRX013', 1, 'PRD004', 1, 16000000, 50000, 15000000),
('TRX014', 1, 'PRD006', 1, 12000000, 100000, 11000000),
('TRX014', 2, 'PRD009', 1, 10000000, 50000, 9500000),
('TRX015', 1, 'PRD010', 1, 8000000, 0, 8000000),
('TRX016', 1, 'PRD003', 1, 19000000, 150000, 17500000),
('TRX016', 2, 'PRD001', 1, 16000000, 0, 16000000),
('TRX017', 1, 'PRD005', 2, 3200000, 200000, 6000000),
('TRX018', 1, 'PRD008', 1, 14000000, 100000, 13000000),
('TRX018', 2, 'PRD002', 1, 13000000, 0, 13000000),
('TRX019', 1, 'PRD007', 1, 8000000, 0, 8000000),
('TRX020', 1, 'PRD004', 1, 16000000, 100000, 15000000),
('TRX020', 2, 'PRD006', 1, 12000000, 50000, 11500000);
