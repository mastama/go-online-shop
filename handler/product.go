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

// ListProducts adalah handler function yang digunakan untuk menampilkan semua data list products
func ListProducts(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// mengambil semua data list products dan di simpan ke variable products
		products, err := model.SelectProduct(db)
		if err != nil {
			// Jika terjadi kesalahan saat mengambil data list products
			log.Printf("Terjadi kesalahan saat mengambil data list product: %v\n", err)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Gagal mengambil data list produk",
				"data":            products,
			})
			return
		}

		// jika semua data list products berhasil diambil, maka kirim response 200 dan data products
		log.Printf("Berhasil mengambil data list produk: %v\n", products)
		ctx.JSON(200, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": "Berhasil mengambil data list produk",
			"data":            products,
		})
	}

}

// GetProducts adalah handler function untuk menampilkan data product
func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// mengambil parameter id product dan disimpan pada variable id
		id := ctx.Param("id")

		// ambil product by id dan disimpan pada variable product
		product, err := model.SelectProductById(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// Jika terjadi kesalahan saat mengambil data produk pada database karena tidak ada
				log.Printf("Terjadi kesalahan saat mengambil data produk: %v\n", err)
				ctx.JSON(404, gin.H{
					"responseCode":    "013",
					"responseMessage": "Produk tidak ditemukan",
					"data":            nil,
				})
				return
			}
			// Jika terjadi kesalahan pada server atau query yang salah
			log.Printf("Terjadi kesalahan pada server saya mengambil data produk: %v\n", err)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Terjadi kesalahan pada server. Gagal mengambil data produk",
				"data":            nil,
			})
			return
		}

		// jika data produk berhasil di tampilkan maka berikan response 200
		log.Printf("Berhasil mengambil product dengan id: %v\n", product)
		ctx.JSON(200, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": fmt.Sprintf("Berhasil mengambil produk dengan id: %s", product.ID),
			"data":            product,
		})
	}
}

// CreateProduct adalah handler function yang digunakan untuk membuat product
func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// membuat variable product untuk menyimpan requestBody yang akan di-bind ke struct product
		var product model.Product
		// melakukan binding requestBody ke struct product
		if err := ctx.Bind(&product); err != nil {
			// Jika terjadi kesalahan saat binding
			log.Printf("Terjadi kesalahan saat binding data product: %v\n", err)
			ctx.JSON(400, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Gagal binding data product",
				"data":            nil,
			})
			return
		}

		// mengatur id product secara otomatis menggunakan UUID
		product.ID = string(uuid.New().String())

		// menyimpan data product ke database
		if err := model.InsertProduct(db, product); err != nil {
			// Jika terjadi kesalahan saat menyimpan data product ke database, kirim response 500 (server error)
			log.Printf("Terjadi kesalahan saat membuat data product: %v\n", err)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Terjadi kesalahan pada server. Gagal menyimpan data produk",
				"data":            nil,
			})
			return
		}
		// Jika data product berhasil disimpan, kirim response 201 (created) dan data product yang baru sudah dibuat
		log.Printf("Berhasil membuat product baru: %+v\n", product)
		ctx.JSON(201, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": "Berhasil membuat product baru",
			"data":            product,
		})
	}
}

// UpdateProduct adalah handler function yang digunakan untuk memperbarui data produk berdasarkan ID yang diberikan
func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Mengambil parameter "id" dari URL request
		id := ctx.Param("id")

		// Membuat variabel productReq untuk menyimpan request body yang akan di-bind ke struct Product
		var productReq model.Product

		// Mengikat (bind) data dari request body ke productReq
		if err := ctx.Bind(&productReq); err != nil {
			// Jika terjadi kesalahan saat binding (misal format data salah), kirim response 400 (bad request)
			log.Printf("Terjadi kesalahan saat membaca request body data product: %v\n", err)
			ctx.JSON(400, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Data produk tidak valid", // Pesan error yang akan ditampilkan
				"data":            nil,                       // Data kosong karena terjadi kesalahan
			})
			return
		}

		// Mengambil data produk dari database berdasarkan ID yang diberikan
		product, err := model.SelectProductById(db, id)
		if err != nil {
			// Jika terjadi kesalahan saat mengambil data produk, kirim response 500 (server error)
			log.Printf("Terjadi kesalahan saat mengambil data produk: %+v\n", product)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Terjadi kesalahan pada server", // Pesan error jika ada masalah server
				"data":            nil,                             // Data kosong karena terjadi kesalahan
			})
			return
		}

		// Jika field Name dari productReq tidak kosong, update field Name di product
		if productReq.Name != "" {
			product.Name = productReq.Name
		}

		// Jika field Price dari productReq tidak bernilai 0, update field Price di product
		if productReq.Price != 0 {
			product.Price = productReq.Price
		}

		// Memperbarui data produk di database
		if err := model.UpdateProduct(db, product); err != nil {
			// Jika terjadi kesalahan saat menyimpan perubahan, kirim response 500 (server error)
			log.Printf("Terjadi kesalahan saat menyimpan atau memperbarui data produk: %v\n", err)
			ctx.JSON(500, gin.H{
				"responseCode":    "ERROR SERVER",
				"responseMessage": "Terjadi kesalahan pada server. Gagal menyimpan data produk",
				"data":            nil,
			})
			return
		}

		// Jika berhasil memperbarui produk, kirim response 201 (created) dengan data produk yang sudah diperbarui
		log.Printf("Berhasil update product: %+v\n", product)
		ctx.JSON(201, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": "Berhasil update product", // Pesan sukses jika update berhasil
			"data":            product,                   // Mengembalikan data produk yang telah diperbarui
		})
	}
}

func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// mengambil paramter id dari request
		id := ctx.Param("id")

		// cek ke database apakah product dengan id tersebut ada atau tidak
		_, err := model.SelectProductById(db, id)
		if err != nil {
			// Jika product tidak ditemukan, kirim response 400 (bad request)
			log.Printf("Product dengan id %s tidak ditemukan\n", id)
			ctx.JSON(400, gin.H{
				"responseCode":    "ERROR",
				"responseMessage": "Product dengan id tersebut tidak ditemukan",
				"data":            nil,
			})
			return // keluar function jika product tidak ditemukan sebelum menghapus data product tersebut
		} else {
			// menghapus data product berdasarkan id yang diberikan
			if err := model.DeleteProduct(db, id); err != nil {
				// Jika terjadi kesalahan saat menghapus data product, kirim response 500 (server error)
				log.Printf("Terjadi kesalahan saat menghapus data produk: %v\n", err)
				ctx.JSON(500, gin.H{
					"responseCode":    "ERROR SERVER",
					"responseMessage": "Terjadi kesalahan pada server. Gagal menghapus data produk",
					"data":            nil,
				})
				return
			}
		}
		// Jika berhasil menghapus produk, kirim response 204 (no content)
		log.Printf("Berhasil menghapus product dengan ID: %s\n", id)
		ctx.JSON(200, gin.H{
			"responseCode":    "SUCCESS",
			"responseMessage": "Berhasil menghapus product", // Pesan sukses jika produk berhasil dihapus
			"data":            nil,                          // Data kosong karena produk telah dihapus
		})
	}
}
