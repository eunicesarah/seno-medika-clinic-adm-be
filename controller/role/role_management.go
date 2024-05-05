package role

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"seno-medika.com/config/db"
	"seno-medika.com/helper"
	"seno-medika.com/model/common"
	"seno-medika.com/model/person"
	superadmin2 "seno-medika.com/query/role/superadmin"
	"strconv"
	"sync"
)

type filterResponse struct {
	User []person.UserWithoutPassword `json:"user"`
	Size    int                    `json:"size"`
}

func AddUser(c *gin.Context) {
	var userInput person.User
	var wg sync.WaitGroup

	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	errChan := make(chan error, 3)
	wg.Add(3)
	go func() {
		defer wg.Done()
		helper.ValidationEmail(userInput.Email, errChan)
	}()
	go func() {
		defer wg.Done()
		helper.IsEmailExists(userInput.Email, errChan)
	}()
	go func() {
		defer wg.Done()
		helper.ValidationPassword(userInput.Password, errChan)
	}()
	wg.Wait()

	close(errChan)

	if err := <-errChan; err != nil {
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

	_, err = db.DB.Query(
		"INSERT INTO users(user_uuid, nama, password, email, role) VALUES($1,$2,$3,$4,$5)", userInput.UserUUID, userInput.Nama, userInput.Password,
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
		err := superadmin2.DeleteUserById(val)
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
		err := superadmin2.DeleteUserByUuid(target)
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
		err := superadmin2.DeleteUserByName(target)
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
		err := superadmin2.DeleteUserByEmail(target)
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
		err := superadmin2.DeleteUserByRole(target)
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

func PutUser(c *gin.Context) {
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
		err := superadmin2.UpdateUserById(val, userInput)
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
		err := superadmin2.UpdateUserByUuid(target, userInput)
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

func PatchUser(c *gin.Context) {
	var putInput common.PatchInput

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

			err := superadmin2.ChangeNameByUuid(putInput.Key.(string), putInput.Value.(string))
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
			err := superadmin2.ChangeNameById(id, putInput.Value.(string))
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

			err := superadmin2.ChangeEmailByUuid(putInput.Key.(string), putInput.Value.(string))
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
			err := superadmin2.ChangeEmailById(strconv.Itoa(id), putInput.Value.(string))
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

			err := superadmin2.ChangePasswordByUuid(putInput.Key.(string), putInput.Value.(string))
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
			err := superadmin2.ChangePasswordById(id, putInput.Value.(string))
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

			err := superadmin2.ChangeRoleByUuid(putInput.Key.(string), putInput.Value.(string))
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
			err := superadmin2.ChangeRoleById(id, putInput.Value.(string))
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


	switch findBy {
	case "dashboard":
		page := c.Query("page")
		if page == "" {
			page = "1"
		}
		val, _ := strconv.Atoi(page)
		user, size, err := superadmin2.FindByFilter(strconv.Itoa(10*val - 10), "10")

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
			Data:       filterResponse{User: user, Size: size},
		})
		return
	case "id":
		val, _ := strconv.Atoi(target)
		user, err := superadmin2.FindById(val)
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
		user, err := superadmin2.FindByUuid(target)
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
		user, err := superadmin2.FindByName(target)
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
		user, err := superadmin2.FindByEmail(target)
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
		user, err := superadmin2.FindByRole(target)
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
	default:
		user, err := superadmin2.FindAll()
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
