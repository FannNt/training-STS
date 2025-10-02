# Task 4 - PostgreSQL Database Implementation

## 📁 Files Overview

### **schema.sql**
Database structure definition with all tables, foreign keys, and constraints.

### **seeder.sql**
Initial data insertion for all tables (includes 10 products, 10 customers, 20 transactions).

### **query.sql**
Contains the two main queries requested:
1. Customers without transactions in date range (2024-01-15 to 2024-01-20)
2. Sales summary grouped by brand

### **alter_add_status.sql**
Script to add `status_penjualan` column to existing database and populate with random values.

### **verify.sql**
Verification queries to check data integrity and relationships.

---

## 🚀 How to Execute

### Step 1: Initial Setup (Fresh Database)
```bash
# Create schema
sudo -u postgres psql < schema.sql

# Insert data
sudo -u postgres psql < seeder.sql

# Verify
sudo -u postgres psql < verify.sql
```

### Step 2: Modify Existing Schema (Add Status Column)
```bash
# Add status_penjualan column and update records
sudo -u postgres psql < alter_add_status.sql
```

### Step 3: Run Custom Queries
```bash
# Run custom queries
sudo -u postgres psql < query.sql
```

**Note:** The `alter_add_status.sql` demonstrates ALTER TABLE operations on an existing database. It should be run AFTER the initial setup, not included in schema.sql.

---

## 📊 Database Structure

### Tables
1. **master_produk** - Product master data
2. **master_pelanggan** - Customer master data
3. **master_harga_jual** - Product pricing by customer level
4. **transaksi_penjualan** - Sales transaction header
5. **detail_transaksi** - Sales transaction details

### Relationships
- Master Produk ↔ Master Harga Jual: **1-N**
- Master Pelanggan ↔ Transaksi Penjualan: **1-N**
- Transaksi Penjualan ↔ Detail Transaksi: **1-N**
- Master Produk ↔ Detail Transaksi: **1-N**

---

## 📈 Query Results Summary

### Query 1: Customers Without Transactions (2024-01-15 to 2024-01-20)
**Result:** 4 customers
- CUST006 - Maya Sari (Reguler, Semarang)
- CUST008 - Indah Permata (Reguler, Palembang)
- CUST009 - Fajar Nugroho (VIP, Denpasar)
- CUST010 - Lina Marlina (Reguler, Balikpapan)

### Query 2: Sales Summary by Brand (Top 3)
1. **Apple**: 113M (7 units, 5 transactions)
2. **Dell**: 44.25M (3 units, 3 transactions)
3. **Samsung**: 38M (3 units, 3 transactions)

### Transaction Status Distribution
- **Done**: 16 transactions (80%)
- **Cancel**: 4 transactions (20%)

---

## 🔧 Additional Notes

- All monetary values use DECIMAL(15,2) for precision
- Date format: YYYY-MM-DD
- Status values: 'done' or 'cancel'
- Foreign key constraints ensure referential integrity
- Transactions with status 'cancel' are still in the database but marked as cancelled
