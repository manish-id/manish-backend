package controller

import (
	"context"
	// "fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/manish-id/manish-backend/config"
	"github.com/manish-id/manish-backend/model"
)

func CreateProduct(c *fiber.Ctx) error {
	// Ambil koleksi produk dari koneksi database
	collection := config.MongoClient.Database("dbmanish").Collection("products")

	// Bind request ke struct Product
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validasi: nama, harga, dan stok
	if product.ProductName == "" || product.ProductPrice <= 0 || product.Stock < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please fill in valid product_name, product_price, and stock",
		})
	}

	// Validasi: product_image harus link https
	if !strings.HasPrefix(product.ProductImage, "https://") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product image URL. Must start with https://",
		})
	}

	// Validasi: size wajib diisi
	if product.Size == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product size is required",
		})
	}

	// Cek apakah product dengan nama yang sama sudah ada
	count, err := collection.CountDocuments(context.TODO(), bson.M{
		"product_name": product.ProductName,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to check product existence",
		})
	}
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product already exists",
		})
	}

	// Set ID dan waktu
	product.ProductID = primitive.NewObjectID()
	product.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Simpan ke MongoDB
	_, err = collection.InsertOne(context.TODO(), product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}

	// Response sukses
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"data":    product,
	})
}
