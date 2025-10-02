-- ============================================
-- Query 1: Customers without transactions between 2024-01-15 and 2024-01-20
-- ============================================
SELECT 
    p.kode,
    p.nama,
    p.level_pelanggan,
    p.asal_kota
FROM master_pelanggan p
WHERE p.kode NOT IN (
    SELECT DISTINCT t.customer
    FROM transaksi_penjualan t
    WHERE t.tanggal BETWEEN '2024-01-15' AND '2024-01-20'
)
ORDER BY p.kode;

-- ============================================
-- Query 2: Sales summary for each brand
-- ============================================
SELECT 
    mp.brand,
    COUNT(DISTINCT dt.nomor_penjualan) as total_transactions,
    SUM(dt.qty) as total_quantity_sold,
    SUM(dt.total_nilai) as total_sales_amount,
    AVG(dt.total_nilai) as average_sales_per_item,
    COUNT(DISTINCT dt.kode_produk) as products_sold,
    SUM(dt.diskon) as total_discount_given
FROM detail_transaksi dt
JOIN master_produk mp ON dt.kode_produk = mp.kode
GROUP BY mp.brand
ORDER BY total_sales_amount DESC;

-- Detailed brand summary with product breakdown
SELECT 
    mp.brand,
    mp.kode,
    mp.nama_produk,
    mp.kategori,
    COUNT(dt.nomor_penjualan) as times_sold,
    SUM(dt.qty) as total_qty,
    SUM(dt.total_nilai) as total_revenue,
    SUM(dt.diskon) as total_discount
FROM master_produk mp
LEFT JOIN detail_transaksi dt ON mp.kode = dt.kode_produk
GROUP BY mp.brand, mp.kode, mp.nama_produk, mp.kategori
ORDER BY mp.brand, total_revenue DESC;
