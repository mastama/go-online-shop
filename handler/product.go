package handler

import (
	"database/sql"
	"errors"
	"go-online-shop/model"
	"log"

	"github.com/gin-gonic/gin"
)

func ListProducts(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := model.SelectProduct(db)
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil data product: %v\n", err)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Gagal mengambil data produk",
				"data":            products,
			})
			return
		}
		log.Printf("Berhasil mengambil data produk: %v\n", products)
		ctx.JSON(200, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": "Berhasil mengambil data produk",
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
			"responseMessage": "Berhasil mengambil product dengna id: %v",
			"data":            product,
		})
	}
}
