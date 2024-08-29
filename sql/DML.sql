-- Insert data into users
INSERT INTO users (full_name, username, email, password, status, role, book_count)
VALUES
    ('Adi Prasetya', 'adi_prasetya', 'adi@example.com', 'password123', 'active', 'reader', 5),
    ( 'Budi Santoso', 'budi_santoso', 'budi@example.com', 'password123', 'active', 'admin', 10),
    ('Citra Dewi', 'citra_dewi', 'citra@example.com', 'password123', 'inactive', 'reader', 2),
    ( 'Dewi Lestari', 'dewi_lestari', 'dewi@example.com', 'password123', 'active', 'reader', 7),
    ( 'Eko Nugroho', 'eko_nugroho', 'eko@example.com', 'password123', 'banned', 'reader', 0);

-- Insert data into wallets
INSERT INTO wallets (user_id, balance)
VALUES
    ( 1, 150000.00),
    ( 2, 500000.00),
    ( 3, 75000.00),
    ( 4, 200000.00),
    ( 5, 0.00);

-- Insert data into deposits
INSERT INTO deposits (id, amount, wallet_id, invoice_url, status)
VALUES
    (1, 50000.00, 1, 'https://example.com/invoice/1', 'confirmed'),
    (2, 200000.00, 2, 'https://example.com/invoice/2', 'confirmed'),
    (3, 25000.00, 3, 'https://example.com/invoice/3', 'pending'),
    (4, 100000.00, 4, 'https://example.com/invoice/4', 'pending'),
    (5, 0.00, 5, 'https://example.com/invoice/5', 'refunded');

-- Insert data into categories
INSERT INTO categories (name)
VALUES
    ('Fiksi'),
    ('Non-Fiksi'),
    ('Komik'),
    ('Biografi'),
    ('Sejarah');

-- Insert data into books
INSERT INTO books ( category_id, isbn, sku, author, title, image, pages, language, description, stock, status, borrowed_count, published_at, base_price)
VALUES
    (1, '978-602-03-1677-7', 'SKU123456', '{"name": "Tere Liye"}', 'Hujan', 'https://example.com/image/hujan.jpg', 300, 'Indonesia', 'Novel tentang cinta dan hujan.', 10, 'available', 3, '2023-01-01', 100000.00),
    (2, '978-602-03-1678-4', 'SKU123457', '{"name": "Ayu Utami"}', 'Saman', 'https://example.com/image/saman.jpg', 350, 'Indonesia', 'Novel yang menggambarkan kehidupan sosial dan politik.', 5, 'available', 2, '2023-02-01', 120000.00),
    (3, '978-602-03-1679-1', 'SKU123458', '{"name": "Dewi Lestari"}', 'Supernova', 'https://example.com/image/supernova.jpg', 400, 'Indonesia', 'Seri novel yang memadukan sains dan fantasi.', 8, 'available', 5, '2023-03-01', 150000.00),
    (4, '978-602-03-1680-8', 'SKU123459', '{"name": "Pramoedya Ananta Toer"}', 'Bumi Manusia', 'https://example.com/image/bumi_manusia.jpg', 250, 'Indonesia', 'Novel tentang perjuangan hidup dan politik.', 3, 'not available', 7, '2023-04-01', 90000.00),
    (5, '978-602-03-1681-5', 'SKU123460', '{"name": "Andrea Hirata"}', 'Laskar Pelangi', 'https://example.com/image/laskar_pelangi.jpg', 500, 'Indonesia', 'Kisah inspiratif tentang anak-anak di Belitong.', 12, 'available', 4, '2023-05-01', 110000.00);

-- Insert data into book_physical_details
INSERT INTO book_physical_details ( book_id, weight, height, width)
VALUES
    (1, 0.5, 21, 14),
    (2, 0.6, 22, 15),
    (3, 0.7, 23, 16),
    (4, 0.4, 20, 13),
    (5, 0.8, 24, 17);

-- Insert data into borrows
INSERT INTO borrows ( user_id, book_id, status, start_date, end_date, total_price, returned_at)
VALUES
    ( 1, 1, 'pending', '2024-08-01', '2024-08-15', 100000.00, NULL),
    ( 2, 2, 'success', '2024-07-15', '2024-08-01', 120000.00, '2024-08-01'),
    ( 3, 3, 'pending', '2024-08-10', '2024-08-24', 150000.00, NULL),
    ( 4, 4, 'returned', '2024-07-01', '2024-07-15', 90000.00, '2024-07-15'),
    ( 5, 5, 'cancel', '2024-08-05', '2024-08-20', 110000.00, NULL);

-- Insert data into borrow_prices
INSERT INTO borrow_prices ( book_id, duration_type, price_multiplier)
VALUES
    (1, 'daily', 1.00),
    (2, 'weekly', 5.00),
    (3, 'monthly', 20.00),
    (4, 'daily', 0.75),
    (5, 'weekly', 4.00);

-- Insert data into payments
INSERT INTO payments (borrow_id, amount, payment_date, status)
VALUES
    (1, 100000.00, '2024-08-01', 'confirmed'),
    (2, 120000.00, '2024-08-01', 'confirmed'),
    (3, 150000.00, '2024-08-10', 'pending'),
    (4, 90000.00, '2024-07-15', 'pending'),
    (5, 110000.00, '2024-08-05', 'refunded');
