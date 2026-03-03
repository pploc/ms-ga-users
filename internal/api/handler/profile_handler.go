package handler

import (
	"net/http"

	"ms-ga-user/internal/api/generated"
	"ms-ga-user/internal/domain/entity"
	"ms-ga-user/internal/service"
	"ms-ga-user/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileHandler struct {
	profileService service.ProfileService
}

func NewProfileHandler(s service.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: s}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid user id format")
		return
	}

	profile, addresses, err := h.profileService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, ToGeneratedUserProfile(profile, addresses), nil)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid user id format")
		return
	}

	var req generated.UpdateUserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	profile := &entity.UserProfile{
		UserID:                userID,
		Department:            req.Department,
		EmergencyContactName:  req.EmergencyContactName,
		EmergencyContactPhone: req.EmergencyContactPhone,
	}

	if req.HireDate != nil {
		profile.HireDate = &req.HireDate.Time
	}

	var addresses []*entity.UserAddress
	if req.Addresses != nil {
		for _, a := range *req.Addresses {
			addr := &entity.UserAddress{
				UserID:    userID,
				Street:    *a.Street,
				City:      *a.City,
				State:     a.State,
				ZipCode:   a.ZipCode,
				Country:   *a.Country,
				IsPrimary: *a.IsPrimary,
			}
			if a.Id != nil {
				addr.ID = uuid.UUID(*a.Id)
			} else {
				addr.ID = uuid.New()
			}
			addresses = append(addresses, addr)
		}
	}

	updated, err := h.profileService.UpdateProfile(c.Request.Context(), profile, addresses)
	if err != nil {
		utils.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, ToGeneratedUserProfile(updated, addresses), nil)
}
