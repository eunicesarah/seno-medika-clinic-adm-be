package apotek 

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/query/apotek"
)

func GetApotek(c *gin.Context){
	find_by := c.Query("find_by")
	target := c.Query("target")

	switch find_by {
	case "today":
		apotek, err := apotek.FindAllAntrianApotekToday()
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
			Message:    "Successfully get apotek data today",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       apotek,
		})
		return
	case "date":
		apotek, err := apotek.FindAllAntrianApotekByDate(target)
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
			Message:    "Successfully get apotek data on " + target,
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       apotek,
		})
		return
	case "detail_antrian":
		val, _ := strconv.Atoi(target)
		apotek, err := apotek.FindDetailResepByNoAntrian(val)
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
			Message:    "Successfully get apotek data detail obat",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       apotek,
		})
		return
	default:
		apotek, err := apotek.FindAllAntrianApotek()
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
			Message:    "Successfully get apotek data",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       apotek,
		})
		return
	}
}
