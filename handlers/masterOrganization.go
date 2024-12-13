package handlers

import (
	"belajar-api/models"
	"belajar-api/response"
	"belajar-api/services"

	"github.com/gofiber/fiber/v2"
)

type MasterOrganizationHandler struct {
	MasterOrganizationService *services.MasterOrganizationService
}

// GetAllMasterOrganizations handles the request to fetch all MasterOrganizations
func (h *MasterOrganizationHandler) GetAllMasterOrganizations(c *fiber.Ctx) error {
	MasterOrganizations, err := h.MasterOrganizationService.GetAllMasterOrganizations()
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve Organizations")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Organizations retrieved successfully", MasterOrganizations)
}

// GetMasterOrganizationByUID handles the request to fetch a MasterOrganization by its UID
func (h *MasterOrganizationHandler) GetMasterOrganizationByUID(c *fiber.Ctx) error {
	uid := c.Params("uid")
	MasterOrganization, err := h.MasterOrganizationService.GetMasterOrganizationByUID(uid)
	if err != nil {
		return response.InternalServerError(c, err, "Cannot retrieve Organization")
	}
	if MasterOrganization == nil {
		return response.BadRequest(c, err, "Organization not found")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Organization retrieved successfully", MasterOrganization)
}

// CreateMasterOrganization handles the request to create a new MasterOrganization
func (h *MasterOrganizationHandler) CreateMasterOrganization(c *fiber.Ctx) error {
	var MasterOrganization models.MasterOrganization
	if err := c.BodyParser(&MasterOrganization); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.MasterOrganizationService.CreateMasterOrganization(&MasterOrganization); err != nil {
		return response.InternalServerError(c, err, "Cannot create Organization")
	}
	return response.SuccessHandler(c, fiber.StatusCreated, "Organization created successfully", MasterOrganization)
}

// UpdateMasterOrganization handles the request to update a MasterOrganization by its UID
func (h *MasterOrganizationHandler) UpdateMasterOrganization(c *fiber.Ctx) error {
	uid := c.Params("uid")
	var updatedMasterOrganization models.MasterOrganization
	if err := c.BodyParser(&updatedMasterOrganization); err != nil {
		return response.BadRequest(c, err, "Cannot parse JSON")
	}
	if err := h.MasterOrganizationService.UpdateMasterOrganization(uid, &updatedMasterOrganization); err != nil {
		return response.InternalServerError(c, err, "Cannot update Organization")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Organization updated successfully", updatedMasterOrganization)
}

// DeleteMasterOrganization handles the request to delete a MasterOrganization by its UID
func (h *MasterOrganizationHandler) DeleteMasterOrganization(c *fiber.Ctx) error {
	uid := c.Params("uid")
	if err := h.MasterOrganizationService.DeleteMasterOrganization(uid); err != nil {
		return response.InternalServerError(c, err, "Cannot delete Organization")
	}
	return response.SuccessHandler(c, fiber.StatusOK, "Organization deleted successfully", nil)
}
