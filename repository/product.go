package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/manish-id/manish-backend/model"
)

func CreateProduct(db *mongo.Database, product *model.Product) error {
	collection := db.Collection("products")

	// Cek apakah product dengan nama yang sama sudah ada
	count, err := collection.CountDocuments(context.TODO(), bson.M{"product_name": product.ProductName})
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("product already exists")
	}

	// Set ID dan waktu otomatis
	product.ProductID = primitive.NewObjectID()
	product.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Validasi harga dan stock
	if product.ProductPrice <= 0 {
		return fmt.Errorf("invalid product price")
	}
	if product.Stock < 0 {
		return fmt.Errorf("stock cannot be negative")
	}

	// Validasi link gambar (harus diawali dengan https://)
	if !strings.HasPrefix(product.ProductImage, "https://") {
		return fmt.Errorf("invalid product image URL")
	}

	// Validasi ukuran (optional)
	if product.Size == "" {
		return fmt.Errorf("size is required")
	}

	// Simpan ke database
	_, err = collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}

	return nil
}