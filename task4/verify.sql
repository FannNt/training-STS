-- Verify data insertion
SELECT 'Master Produk' as table_name, COUNT(*) as record_count FROM master_produk
UNION ALL
SELECT 'Master Pelanggan', COUNT(*) FROM master_pelanggan
UNION ALL
SELECT 'Master Harga Jual', COUNT(*) FROM master_harga_jual
UNION ALL
SELECT 'Transaksi Penjualan', COUNT(*) FROM transaksi_penjualan
UNION ALL
SELECT 'Detail Transaksi', COUNT(*) FROM detail_transaksi;

-- Check relationships
SELECT 
    t.nomor_penjualan,
    t.tanggal,
    p.nama as customer_name,
    p.level_pelanggan,
    COUNT(d.urut_item) as total_items,
    SUM(d.total_nilai) as total_transaction
FROM transaksi_penjualan t
JOIN master_pelanggan p ON t.customer = p.kode
JOIN detail_transaksi d ON t.nomor_penjualan = d.nomor_penjualan
GROUP BY t.nomor_penjualan, t.tanggal, p.nama, p.level_pelanggan
ORDER BY t.nomor_penjualan;
