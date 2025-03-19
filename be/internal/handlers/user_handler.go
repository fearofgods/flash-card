package handlers

import "github.com/gin-gonic/gin"

type UserController struct{}

// GetUsersHandler retrieves all users
// @Summary Get all users
// @Description Fetch all users from the database
// @Tags Users
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /users [get]
func (h UserController) GetUserById(ctx *gin.Context) {
	if ctx.Param("id") != "" {

	}

	return
}
