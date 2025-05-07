CREATE TABLE IF NOT EXISTS assets (
                                      id UUID PRIMARY KEY,
                                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      user_id UUID NOT NULL,
                                      name VARCHAR(255) NOT NULL,
                                      ticker VARCHAR(50) NOT NULL,
                                      quantity INT NOT NULL,
                                      current_price DECIMAL(10, 2),
                                      purchase_price DECIMAL(10, 2) NOT NULL,
                                      mid_price DECIMAL(10, 2),
                                      percentage DECIMAL(5, 2)
);
