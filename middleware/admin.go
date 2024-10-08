package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminOnly(ctx *gin.Context) {
	key := os.Getenv("SECRET_KEY")

	// Ambil token yang dikirim melalui header Authorization
	token := ctx.GetHeader("Authorization")

	// Cek apakah token yang dikirim sesuai dengan secret key yang dibutuhkan
	// Jika token salah, kirim response 401 (unauthorized)
	if token == "" {
		ctx.JSON(401, gin.H{
			"responseCode":    "ERROR",
			"responseMessage": "Unauthorized access",
			"data":            nil,
		})
		ctx.Abort()
		return
	}

	// validasi token yang dikirim pada header Authorization
	// Jika bukan admin, kirim response 403 (forbidden)
	if token != key {
		ctx.JSON(403, gin.H{
			"responseCode":    "ERROR",
			"responseMessage": "Anda bukan admin",
			"data":            nil,
		})
		ctx.Abort()
		return
	}
	// Jika token valid, lanjutkan ke proses lain
	ctx.Next()
}
