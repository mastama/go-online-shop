package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"go-online-shop/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListProducts(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := model.SelectProduct(db)
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil data list product: %v\n", err)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Gagal mengambil data list produk",
				"data":            products,
			})
			return
		}
		log.Printf("Berhasil mengambil data list produk: %v\n", products)
		ctx.JSON(200, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": "Berhasil mengambil data list produk",
			"data":            products,
		})
	}

}

func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		// ambil product by id
		product, err := model.SelectProductById(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Terjadi kesalahan saat mengambil data produk: %v\n", err)
				ctx.JSON(404, gin.H{
					"responseCode":    "013",
					"responseMessage": "Produk tidak ditemukan",
					"data":            nil,
				})
				return
			}
			log.Printf("Terjadi kesalahan pada server saya mengambil data produk: %v\n", err)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR SERVER",
				"responseMessage": "Terjadi kesalahan pada server. Gagal mengambil data produk",
				"data":            nil,
			})
			return
		}
		log.Printf("Berhasil mengambil product dengan id: %v\n", product)
		ctx.JSON(200, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": fmt.Sprintf("Berhasil mengambil produk dengan id: %s", product.ID),
			"data":            product,
		})
	}
}

func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product model.Product
		if err := ctx.Bind(&product); err != nil {
			log.Printf("Terjadi kesalahan saat binding data product: %v\n", err)
			ctx.JSON(400, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Gagal binding data product",
				"data":            nil,
			})
			return
		}

		// simpan id
		product.ID = string(uuid.New().String())

		// simpan ke database
		if err := model.InsertProduct(db, product); err != nil {
			log.Printf("Terjadi kesalahan saat membuat data product: %v\n", err)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR SERVER",
				"responseMessage": "Terjadi kesalahan pada server. Gagal menyimpan data produk",
				"data":            nil,
			})
			return
		}
		log.Printf("Berhasil membuat product baru: %+v\n", product)
		ctx.JSON(201, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": "Berhasil membuat product baru",
			"data":            product,
		})
	}
}
func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO Implement update Product
	}
}
func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO Implement delete Product
	}
}
