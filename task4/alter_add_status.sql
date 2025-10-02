-- ============================================
-- Add status_penjualan column to transaksi_penjualan table
-- ============================================

-- Add the column with CHECK constraint
ALTER TABLE transaksi_penjualan 
ADD COLUMN status_penjualan VARCHAR(10) CHECK (status_penjualan IN ('cancel', 'done'));

-- Update all transactions with random status (cancel or done)
UPDATE transaksi_penjualan
SET status_penjualan = CASE 
    WHEN random() < 0.7 THEN 'done'
    ELSE 'cancel'
END;

-- Make the column NOT NULL after populating data
ALTER TABLE transaksi_penjualan 
ALTER COLUMN status_penjualan SET NOT NULL;

-- Verify the update
SELECT 
    status_penjualan,
    COUNT(*) as count,
    ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM transaksi_penjualan), 2) as percentage
FROM transaksi_penjualan
GROUP BY status_penjualan
ORDER BY status_penjualan;
