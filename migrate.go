package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func dbMigrate(db *sql.DB) (sql.Result, error) {

	if db == nil {
		return nil, errors.New("tidak ada koneksi")
	}

	// Migrasi tabel-tabel ke database
	result, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS products (
            id VARCHAR(36) PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            price BIGINT NOT NULL,
			is_deleted BOOLEAN NOT NULL DEFAULT FALSE
        );
		CREATE TABLE IF NOT EXISTS orders (
		 	id VARCHAR(36) PRIMARY KEY,
            email VARCHAR(255) NOT NULL,
			address VARCHAR(255) NOT NULL,
			passcode VARCHAR(255),
			paid_at TIMESTAMP,
			paid_bank VARCHAR(255),
			paid_account VARCHAR(255),
			grand_total BIGINT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS order_details (
            id VARCHAR(36) PRIMARY KEY,
			order_id VARCHAR(36),
            product_id VARCHAR(36),
            quantity INT NOT NULL,
            price BIGINT NOT NULL,
            total BIGINT NOT NULL,
			FOREIGN KEY (order_id) REFERENCES orders(id) ON UPDATE CASCADE ON DELETE RESTRICT,
            FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE RESTRICT
		);
	`)

	if err != nil {
		return nil, fmt.Errorf("gagal melakukan migrasi database: %w", err)
	}

	fmt.Println("Migrasi database sukses!")
	return result, nil
}
