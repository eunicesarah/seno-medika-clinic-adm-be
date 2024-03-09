package konfirmasi

import (
    "github.com/gin-gonic/gin"
    "seno-medika.com/controller/konfirmasi"
)

func KonfirmasiPasienRouter(r *gin.Engine) {
    r.PATCH("/konfirmasipasien", konfirmasi.ConfirmPasien)
}