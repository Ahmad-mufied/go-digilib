-- Function to update timestamp
CREATE OR REPLACE FUNCTION update_modified_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Function to update book status based on stock
CREATE OR REPLACE FUNCTION update_book_status()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.stock > 0 THEN
        NEW.status = 'available';
    ELSE
        NEW.status = 'not available';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Function to update borrow count
CREATE OR REPLACE FUNCTION update_book_borrow_count()
    RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        IF NEW.status != 'cancel' THEN
            UPDATE books SET borrowed_count = borrowed_count + 1 WHERE id = NEW.book_id;
        END IF;
    ELSIF TG_OP = 'UPDATE' THEN
        IF OLD.status != 'cancel' AND NEW.status = 'cancel' THEN
            UPDATE books SET borrowed_count = borrowed_count - 1 WHERE id = NEW.book_id;
        ELSIF OLD.status = 'cancel' AND NEW.status != 'cancel' THEN
            UPDATE books SET borrowed_count = borrowed_count + 1 WHERE id = NEW.book_id;
        END IF;
    ELSIF TG_OP = 'DELETE' THEN
        IF OLD.status != 'cancel' THEN
            UPDATE books SET borrowed_count = borrowed_count - 1 WHERE id = OLD.book_id;
        END IF;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER update_users_modtime
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();


CREATE TRIGGER update_wallets_modtime
    BEFORE UPDATE
    ON wallets
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_deposits_modtime
    BEFORE UPDATE
    ON deposits
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_books_modtime
    BEFORE UPDATE
    ON books
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_books_status
    BEFORE INSERT OR UPDATE OF stock ON books
    FOR EACH ROW
EXECUTE FUNCTION update_book_status();

CREATE TRIGGER update_book_physical_details_modtime
    BEFORE UPDATE
    ON book_physical_details
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_borrows_modtime
    BEFORE UPDATE
    ON borrows
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_book_borrow_count
    AFTER INSERT OR UPDATE OR DELETE ON borrows
    FOR EACH ROW
EXECUTE FUNCTION update_book_borrow_count();

CREATE TRIGGER update_borrow_prices_modtime
    BEFORE UPDATE
    ON borrow_prices
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_payments_modtime
    BEFORE UPDATE
    ON payments
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();




