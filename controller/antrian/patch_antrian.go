package antrian

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "seno-medika.com/model/antrian"
    "seno-medika.com/model/common"
    antrian2 "seno-medika.com/service/antrian"
)

func ConfirmPasien(c *gin.Context) {
    var antrianPasien antrian.Antrian
    antrianID := c.Query("AntrianID")

    if err := c.ShouldBind(&antrianPasien); err != nil {
        c.JSON(http.StatusBadRequest, common.Response{
            Message:    err.Error(),
            Status:     "Bad Request",
            StatusCode: http.StatusBadRequest,
            Data:       nil,
        })
        return
    }

    val, err := strconv.Atoi(antrianID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, common.Response{
            Message:    "Invalid Antrian ID",
            Status:     "Bad Request",
            StatusCode: http.StatusBadRequest,
            Data:       nil,
        })
        return
    }

    err = antrian2.UpdateStatusAntrianById(val, antrianPasien)
    if err != nil {
        c.JSON(http.StatusInternalServerError, common.Response{
            Message:    err.Error(),
            Status:     "Internal Server Error",
            StatusCode: http.StatusInternalServerError,
            Data:       nil,
        })
        return
    }

    c.JSON(http.StatusOK, common.Response{
        Message:    "Successfully update antrian",
        Status:     "ok",
        StatusCode: http.StatusOK,
        Data:       nil,
    })
    return
}