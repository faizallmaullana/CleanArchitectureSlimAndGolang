package main

import (
	"log"

	"github.com/faizallmaullana/rekrutment-gbu-go/internal/config"
	"github.com/faizallmaullana/rekrutment-gbu-go/internal/infrastructure/postgres"
	"gorm.io/gorm"
)

type seedProduct struct {
	ID    int64
	Name  string
	Price float64
}

func main() {
	settings := config.Load()
	database, err := postgres.NewDatabase(settings.Database.DSN())
	if err != nil {
		log.Fatal(err)
	}
	sqlDatabase, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDatabase.Close()

	products := []seedProduct{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 19.99},
		{ID: 3, Name: "Product 3", Price: 5.99},
	}

	err = database.Transaction(func(transaction *gorm.DB) error {
		for _, product := range products {
			if err := transaction.Exec(`
				INSERT INTO products (id, name, price)
				VALUES (?, ?, ?)
				ON CONFLICT (id) DO UPDATE
				SET name = EXCLUDED.name, price = EXCLUDED.price, updated_at = NOW()
			`, product.ID, product.Name, product.Price).Error; err != nil {
				return err
			}
		}

		return transaction.Exec(`
			SELECT setval(
				pg_get_serial_sequence('products', 'id'),
				GREATEST((SELECT COALESCE(MAX(id), 1) FROM products), 1),
				true
			)
		`).Error
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("seeded %d products", len(products))
}
