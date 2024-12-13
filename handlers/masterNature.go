package handlers

import (
	"belajar-api/models"
	"belajar-api/response"
	"belajar-api/services"

	"github.com/gofiber/fiber/v2"
)

type MasterNatureHandler struct {
	MasterNatureService *services.MasterNatureService
}

// GetAllMasterPriorities handles the request to fetch all MasterPriorities
func (h *MasterNatureHandler) GetAllMasterNatures(c *fiber.Ctx) error {
	priorities, err := h.MasterNatureService.GetAllMasterNatures()
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve Nature")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Nature retrieved successfully", priorities)
}

// GetMasterNatureByUID handles the request to fetch a MasterNature by its UID
func (h *MasterNatureHandler) GetMasterNatureByUID(c *fiber.Ctx) error {
	uid := c.Params("uid")
	Nature, err := h.MasterNatureService.GetMasterNatureByUID(uid)
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve Nature")
	}
	if Nature == nil {
		return response.BadRequest(c, err, "Nature not found")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Nature retrieved successfully", Nature)
}

// CreateMasterNature handles the request to create a new MasterNature
func (h *MasterNatureHandler) CreateMasterNature(c *fiber.Ctx) error {
	var Nature models.MasterNature
	if err := c.BodyParser(&Nature); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.MasterNatureService.CreateMasterNature(&Nature); err != nil {
		return response.InternalServerError(c, err, "Cannot create Nature")
	}
	return response.SuccessHandler(c, fiber.StatusCreated, "Nature created successfully", Nature)
}

// UpdateMasterNature handles the request to update a MasterNature by its UID
func (h *MasterNatureHandler) UpdateMasterNature(c *fiber.Ctx) error {
	uid := c.Params("uid")
	var updatedNature models.MasterNature
	if err := c.BodyParser(&updatedNature); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.MasterNatureService.UpdateMasterNature(uid, &updatedNature); err != nil {
		return response.InternalServerError(c, err, "Cannot update Nature")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Nature updated successfully", updatedNature)
}

// DeleteMasterNature handles the request to delete a MasterNature by its UID
func (h *MasterNatureHandler) DeleteMasterNature(c *fiber.Ctx) error {
	uid := c.Params("uid")
	if err := h.MasterNatureService.DeleteMasterNature(uid); err != nil {
		return response.InternalServerError(c, err, "Cannot delete Nature")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Nature deleted successfully", nil)
}
