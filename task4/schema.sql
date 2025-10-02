
-- Create Master Produk table
CREATE TABLE master_produk (
    kode VARCHAR(10) PRIMARY KEY,
    nama_produk VARCHAR(100) NOT NULL,
    kategori VARCHAR(50) NOT NULL,
    brand VARCHAR(50) NOT NULL
);

-- Create Master Pelanggan table
CREATE TABLE master_pelanggan (
    kode VARCHAR(10) PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    level_pelanggan VARCHAR(20) NOT NULL CHECK (level_pelanggan IN ('VIP', 'Reguler')),
    asal_kota VARCHAR(50) NOT NULL
);

-- Create Master Harga Jual table
CREATE TABLE master_harga_jual (
    id SERIAL PRIMARY KEY,
    kode_produk VARCHAR(10) NOT NULL,
    harga DECIMAL(15,2) NOT NULL,
    level_pelanggan VARCHAR(20) NOT NULL CHECK (level_pelanggan IN ('VIP', 'Reguler')),
    periode_awal DATE NOT NULL,
    periode_akhir DATE NOT NULL,
    FOREIGN KEY (kode_produk) REFERENCES master_produk(kode) ON DELETE CASCADE,
    UNIQUE (kode_produk, level_pelanggan, periode_awal)
);

-- Create Transaksi Penjualan table
CREATE TABLE transaksi_penjualan (
    nomor_penjualan VARCHAR(10) PRIMARY KEY,
    tanggal DATE NOT NULL,
    customer VARCHAR(10) NOT NULL,
    nama_kasir VARCHAR(100) NOT NULL,
    FOREIGN KEY (customer) REFERENCES master_pelanggan(kode) ON DELETE RESTRICT
);

-- Create Detail Transaksi table
CREATE TABLE detail_transaksi (
    nomor_penjualan VARCHAR(10) NOT NULL,
    urut_item INTEGER NOT NULL,
    kode_produk VARCHAR(10) NOT NULL,
    qty INTEGER NOT NULL CHECK (qty > 0),
    harga DECIMAL(15,2) NOT NULL,
    diskon DECIMAL(15,2) DEFAULT 0,
    total_nilai DECIMAL(15,2) NOT NULL,
    PRIMARY KEY (nomor_penjualan, urut_item),
    FOREIGN KEY (nomor_penjualan) REFERENCES transaksi_penjualan(nomor_penjualan) ON DELETE CASCADE,
    FOREIGN KEY (kode_produk) REFERENCES master_produk(kode) ON DELETE RESTRICT
);
