-- 1. Table for Affiliates
CREATE TABLE affiliates (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    master_affiliate UUID,  -- References the affiliate's parent (NULL for level 1)
    balance DOUBLE PRECISION DEFAULT 0,
    level INTEGER,  -- Indicates the affiliate level (e.g., 1 for L1, 2 for L2, etc.)
    FOREIGN KEY (master_affiliate) REFERENCES affiliates(id) ON DELETE CASCADE
);

-- 2. Table for Products
CREATE TABLE products (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    price DOUBLE PRECISION NOT NULL
);

-- 3. Table for Orders
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    affiliate_id UUID NOT NULL,
    product_id UUID NOT NULL,
    total_amount DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (affiliate_id) REFERENCES affiliates(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- 4. Table for Users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    balance DOUBLE PRECISION NOT NULL,
    affiliate_id UUID,
    FOREIGN KEY (affiliate_id) REFERENCES affiliates(id) ON DELETE SET NULL
);

-- 5. Table for Commissions
CREATE TABLE commissions (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    affiliate_id UUID NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (affiliate_id) REFERENCES affiliates(id) ON DELETE CASCADE
);
