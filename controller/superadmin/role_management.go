package superadmin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/person"
	"seno-medika.com/service/superadmin"
	"strconv"
)

// AddUser TODO: Add validation for email
func AddUser(c *gin.Context) {
	var userInput person.User

	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	userInput.UserUUID = uuid.New()
	hashPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusForbidden, common.Response{
			Message:    err.Error(),
			Status:     "Forbidden",
			StatusCode: http.StatusForbidden,
			Data:       nil,
		})
		return
	}

	userInput.Password = string(hashPass)

	val, err := db.DB.Query(
		"SELECT * FROM users"+
			" WHERE nama = $1 AND email = $2 AND role = $3",
		userInput.Nama, userInput.Email, userInput.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
	}

	if val.Next() {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	_, err = db.DB.Query(
		"INSERT INTO users(user_uuid, nama, password, email, role)"+
			" VALUES($1,$2,$3,$4,$5)", userInput.UserUUID, userInput.Nama, userInput.Password,
		userInput.Email, userInput.Role)

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
		Message:    "Successfully insert user",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}

func DeleteUser(c *gin.Context) {
	deleteBy := c.Query("delete_by")
	target := c.Query("target")

	if deleteBy == "" || target == "" {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
			Status:     "Bad Request",
			Data:       nil,
		})
		return
	}

	switch deleteBy {
	case "id":
		val, _ := strconv.Atoi(target)
		err := superadmin.DeleteUserById(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

	case "uuid":
		err := superadmin.DeleteUserByUuid(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

	case "nama":
		err := superadmin.DeleteUserByName(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

	case "email":
		err := superadmin.DeleteUserByEmail(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

	case "role":
		err := superadmin.DeleteUserByRole(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    fmt.Sprintf("Successfully delete user by %s = %s", deleteBy, target),
		StatusCode: http.StatusOK,
		Status:     "ok",
		Data:       nil,
	})
	return
}

func UpdateUser(c *gin.Context) {
	updateBy := c.Query("update_by")
	target := c.Query("target")
	var userInput person.User

	if updateBy == "" || target == "" {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
			Status:     "Bad Request",
			Data:       nil,
		})
		return
	}

	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	switch updateBy {
	case "id":
		val, _ := strconv.Atoi(target)
		err := superadmin.UpdateUserById(val, userInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

	case "uuid":
		err := superadmin.UpdateUserByUuid(target, userInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    fmt.Sprintf("Successfully update user by %s = %s", updateBy, target),
		StatusCode: http.StatusOK,
		Status:     "ok",
		Data:       nil,
	})
	return
}

func PutUser(c *gin.Context) {
	var putInput common.PutInput

	changeType := c.Query("change_type")
	changeBy := c.Query("change_by")

	if changeType == "" || changeBy == "" {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
			Status:     "Bad Request",
		})
		return
	}

	switch changeType {
	case "name":
		if changeBy == "uuid" {
			if err := c.ShouldBind(&putInput); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := superadmin.ChangeNameByUuid(putInput.Key.(string), putInput.Value.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					StatusCode: http.StatusBadRequest,
					Status:     "Bad Request",
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    fmt.Sprintf("Successfully change name by uuid = %s", putInput.Key.(string)),
				StatusCode: http.StatusOK,
				Status:     "ok",
			})
			return
		}
		if changeBy == "id" {
			if err := c.ShouldBind(&putInput); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			id := putInput.Key.(int)
			err := superadmin.ChangeNameById(id, putInput.Value.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					StatusCode: http.StatusBadRequest,
					Status:     "Bad Request",
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    fmt.Sprintf("Successfully change name by id = %d", id),
				StatusCode: http.StatusOK,
				Status:     "ok",
			})
			return
		}
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
			Status:     "Bad Request",
		})
		return
	case "email":
		if changeBy == "uuid" {
			if err := c.ShouldBind(&putInput); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := superadmin.ChangeEmailByUuid(putInput.Key.(string), putInput.Value.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					StatusCode: http.StatusBadRequest,
					Status:     "Bad Request",
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    fmt.Sprintf("Successfully change email by uuid = %s", putInput.Key.(string)),
				StatusCode: http.StatusOK,
				Status:     "ok",
			})
			return
		}
		if changeBy == "id" {
			if err := c.ShouldBind(&putInput); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			id := putInput.Key.(int)
			err := superadmin.ChangeEmailById(strconv.Itoa(id), putInput.Value.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					StatusCode: http.StatusBadRequest,
					Status:     "Bad Request",
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    fmt.Sprintf("Successfully change email by id = %d", id),
				StatusCode: http.StatusOK,
				Status:     "ok",
			})
			return
		}

		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
			Status:     "Bad Request",
		})
		return
	case "password":
		if changeBy == "uuid" {
			if err := c.ShouldBind(&putInput); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := superadmin.ChangePasswordByUuid(putInput.Key.(string), putInput.Value.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					StatusCode: http.StatusBadRequest,
					Status:     "Bad Request",
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    fmt.Sprintf("Successfully change password by uuid = %s", putInput.Key.(string)),
				StatusCode: http.StatusOK,
				Status:     "ok",
			})
			return
		}
		if changeBy == "id" {
			if err := c.ShouldBind(&putInput); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			id := putInput.Key.(int)
			err := superadmin.ChangePasswordById(id, putInput.Value.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					StatusCode: http.StatusBadRequest,
					Status:     "Bad Request",
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    fmt.Sprintf("Successfully change password by id = %d", id),
				StatusCode: http.StatusOK,
				Status:     "ok",
			})
			return
		}
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
			Status:     "Bad Request",
		})
		return
	case "role":
		if changeBy == "uuid" {
			if err := c.ShouldBind(&putInput); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := superadmin.ChangeRoleByUuid(putInput.Key.(string), putInput.Value.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					StatusCode: http.StatusBadRequest,
					Status:     "Bad Request",
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    fmt.Sprintf("Successfully change role by uuid = %s", putInput.Key.(string)),
				StatusCode: http.StatusOK,
				Status:     "ok",
			})
			return
		}
		if changeBy == "id" {
			if err := c.ShouldBind(&putInput); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			id := putInput.Key.(int)
			err := superadmin.ChangeRoleById(id, putInput.Value.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					StatusCode: http.StatusBadRequest,
					Status:     "Bad Request",
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    fmt.Sprintf("Successfully change role by id = %d", id),
				StatusCode: http.StatusOK,
				Status:     "ok",
			})
			return
		}
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
			Status:     "Bad Request",
		})
		return
	}
}

func GetUser(c *gin.Context) {
	findBy := c.Query("find_by")
	target := c.Query("target")

	if findBy == "" || target == "" {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
			Status:     "Bad Request",
			Data:       nil,
		})
		return
	}

	switch findBy {
	case "id":
		val, _ := strconv.Atoi(target)
		user, err := superadmin.FindById(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    fmt.Sprintf("Successfully get user by id = %d", val),
			StatusCode: http.StatusOK,
			Status:     "ok",
			Data:       user,
		})
		return
	case "uuid":
		user, err := superadmin.FindByUuid(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    fmt.Sprintf("Successfully get user by uuid = %s", target),
			StatusCode: http.StatusOK,
			Status:     "ok",
			Data:       user,
		})
		return
	case "nama":
		user, err := superadmin.FindByName(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    fmt.Sprintf("Successfully get user by name = %s", target),
			StatusCode: http.StatusOK,
			Status:     "ok",
			Data:       user,
		})
		return
	case "email":
		user, err := superadmin.FindByEmail(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    fmt.Sprintf("Successfully get user by email = %s", target),
			StatusCode: http.StatusOK,
			Status:     "ok",
			Data:       user,
		})
		return
	case "role":
		user, err := superadmin.FindByRole(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    fmt.Sprintf("Successfully get user by role = %s", target),
			StatusCode: http.StatusOK,
			Status:     "ok",
			Data:       user,
		})
		return
	case "all":
		user, err := superadmin.FindAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
				Status:     "Bad Request",
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully get all user",
			StatusCode: http.StatusOK,
			Status:     "ok",
			Data:       user,
		})
		return
	}

}
