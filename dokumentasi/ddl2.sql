CREATE TABLE users (
	id INT PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(20) UNIQUE,
	name VARCHAR(50),
	address TEXT,
	email VARCHAR(100),
	password VARCHAR(255),
	isAdmin BOOLEAN DEFAULT FALSE
);

CREATE TABLE menu (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR(100) NOT NULL,
    harga INT NOT NULL,
    deskripsi TEXT,
    category VARCHAR(20),
    status ENUM('available', 'unavailable') DEFAULT 'available'
);

CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    order_date DATETIME NOT NULL,
    discount INT,
    amount INT,
    status ENUM('on progress', 'berhasil', 'gagal') NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)   
);

-- Tabel Detail Pesanan
CREATE TABLE order_details (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT NOT NULL,
    menu_id INT NOT NULL,
    qty INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (menu_id) REFERENCES menu(id)
);

CREATE TABLE carts (
    user_id INT NOT NULL,
    menu_id INT NOT NULL,
    qty INT NOT NULL DEFAULT 1,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (menu_id) REFERENCES menu(id)
);


-- Data
   
INSERT INTO menu (nama, harga, deskripsi, category, status)
VALUES
	('Espresso', 12000, 'A strong black coffee made by forcing steam through ground coffee beans.','Drink', 'available'),
	('Latte', 20000, 'Espresso with a larger amount of steamed milk and a small amount of froth.','Drink', 'available'),
	('Cireng', 10000, 'Deskripsi Cireng','Food', 'available');

INSERT INTO orders  (user_id, order_date, discount, amount, status)
VALUES
    (1, '2023-08-09', 0, 12000, 'on progress'),
    (1, '2023-08-10', 0, 20000,'berhasil');

INSERT INTO order_details  (order_id, menu_id, qty)
VALUES
    (3, 3, 1);


INSERT INTO carts (user_id, menu_id, qty)
VALUES
    (1, 1, 1),
    (1, 3, 2);
 


