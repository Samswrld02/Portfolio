package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Samswrld02/Portfolio/dto"
	"github.com/Samswrld02/Portfolio/models"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ProjectController struct {
	db *gorm.DB
}

func NewProjectController(gorm *gorm.DB) *ProjectController {
	return &ProjectController{
		db: gorm,
	}
}

// basic crud
func (p *ProjectController) Create(c *echo.Context) error {
	var project models.Project

	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&project); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})

	}

	if err := p.db.Create(&project).Error; err != nil {
		return c.JSON(http.StatusExpectationFailed, map[string]string{"error": "could not make project"})
	}

	return c.JSON(http.StatusCreated, map[string]any{"created": project})
}

func (p *ProjectController) Read(c *echo.Context) error {
	var projectsDTO = make([]dto.ReadProjectDTO, 0)
	if err := p.db.Model(&models.Project{}).Find(&projectsDTO).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, projectsDTO)
}

func (p *ProjectController) Show(c *echo.Context) error {
	id := c.Param("id")
	var projectDTO dto.ReadProjectDTO

	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "Invalid, use integer"})
	}

	if err := p.db.Model(&models.Project{}).First(&projectDTO, id).Error; err != nil {
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "Project does not exist"})
	}

	return c.JSON(http.StatusOK, projectDTO)
}

func (p *ProjectController) Update(c *echo.Context) error {
	id := c.Param("id")
	var project models.Project

	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "Invalid, use integer"})
	}

	if err := p.db.First(&project, id).Error; err != nil {
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "Project does not exist"})
	}

	var input dto.UpdateProjectDTO
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "updating went wrong"})
	}

	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if input.Title != nil {
		project.Title = *input.Title
	}
	if input.Description != nil {
		project.Description = *input.Description
	}

	if err := p.db.Save(&project).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Updating went wrong"})
	}

	return c.JSON(http.StatusOK, map[string]string{"succes": fmt.Sprintf("project %d has been updated succesfully", project.ID)})
}

func (p *ProjectController) Delete(c *echo.Context) error {

	return nil
}
