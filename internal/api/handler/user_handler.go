package handler

import (
	"net/http"
	"strconv"

	"ms-ga-user/internal/api/generated"
	"ms-ga-user/internal/domain/entity"
	"ms-ga-user/internal/domain/repository"
	"ms-ga-user/internal/service"
	"ms-ga-user/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req generated.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := uuid.Parse(*req.Id)
	if err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid user id format")
		return
	}

	user := &entity.User{
		ID:        id,
		FirstName: *req.FirstName,
		LastName:  *req.LastName,
		Email:     *req.Email,
		Phone:     req.Phone,
	}

	created, err := h.userService.CreateUser(c.Request.Context(), user)
	if err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, ToGeneratedUser(created), nil)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid user id format")
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		utils.ErrResponse(c, http.StatusNotFound, "user not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, ToGeneratedUser(user), nil)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	search := c.Query("search")

	filter := repository.UserFilter{
		Status: status,
		Search: search,
		Page:   page,
		Limit:  limit,
	}

	users, total, err := h.userService.ListUsers(c.Request.Context(), filter)
	if err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var generatedUsers []generated.User
	for _, u := range users {
		generatedUsers = append(generatedUsers, *ToGeneratedUser(u))
	}

	meta := generated.ResponseMeta{
		Page:  &page,
		Limit: &limit,
		Total: func(i int) *int { return &i }(int(total)),
	}

	utils.SuccessResponse(c, http.StatusOK, generatedUsers, meta)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid user id format")
		return
	}

	var req generated.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	existingUser, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if existingUser == nil {
		utils.ErrResponse(c, http.StatusNotFound, "user not found")
		return
	}

	if req.FirstName != nil {
		existingUser.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		existingUser.LastName = *req.LastName
	}
	if req.Phone != nil {
		existingUser.Phone = req.Phone
	}
	if req.AvatarUrl != nil {
		existingUser.AvatarURL = req.AvatarUrl
	}

	updated, err := h.userService.UpdateUser(c.Request.Context(), existingUser)
	if err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, ToGeneratedUser(updated), nil)
}

func (h *UserHandler) DeactivateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid user id format")
		return
	}

	if err := h.userService.DeactivateUser(c.Request.Context(), id); err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, map[string]string{"message": "User deactivated."}, nil)
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	q := c.Query("q")
	if len(q) < 2 {
		utils.ErrResponse(c, http.StatusBadRequest, "search query must be at least 2 characters")
		return
	}

	filter := repository.UserFilter{Search: q, Limit: 50}
	users, _, err := h.userService.ListUsers(c.Request.Context(), filter)
	if err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var generatedUsers []generated.User
	for _, u := range users {
		generatedUsers = append(generatedUsers, *ToGeneratedUser(u))
	}

	utils.SuccessResponse(c, http.StatusOK, generatedUsers, nil)
}
