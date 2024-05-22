package obat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seno-medika.com/model/common"
	"seno-medika.com/model/station/pharmacystation"
	"seno-medika.com/query/obat/resep"
)

func AddResep(c *gin.Context) {
	var resVal pharmacystation.Resep
	if err := c.ShouldBindJSON(&resVal); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	if err := resep.AddResep(resVal); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Successfully add resep",
		Status:     "Status Created",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})

	return
}

func DeleteResep(c *gin.Context) {
	deleteBy := c.Query("delete_by")
	target := c.Query("target")

	switch deleteBy {
	case "id":
		if err := resep.DeleteResepById(target); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})

			c.JSON(http.StatusOK, common.Response{
				Message:    "Successfully delete resep",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Invalid delete by",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

}

func PatchResep(c *gin.Context) {
	updateBy := c.Query("update_by")
	target := c.Query("target")
	var resVal pharmacystation.Resep

	if err := c.ShouldBindJSON(&resVal); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
	}

	switch updateBy {
	case "id":
		if err := resep.PatchResepById(target, resVal); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully update resep",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nil,
		})
		return

	default:
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Invalid update by",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
}

func GetResep(c *gin.Context) {
	findBy := c.Query("find_by")
	target := c.Query("target")

	switch findBy {
	case "id":
		res, err := resep.FindResepById(target)
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
			Message:    "Successfully get resep",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       res,
		})
		return

	default:
		res, err := resep.FindAllResep()
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
			Message:    "Successfully get resep",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       res,
		})
		return
	}
}
