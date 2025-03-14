package handlers

import (
	"strconv"
	"time"
	"vanishingnote/database"
	"vanishingnote/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateNote(c *fiber.Ctx) error {

	var request struct {
		Content string `json:"content"`
		Expire  string `json:"expiration"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON("Invalid request")
	}

	//convert expiration time in int
	expireTime, err := strconv.Atoi(request.Expire)
	if err != nil {
		return c.Status(400).JSON("Invalid expiration time")
	}

	// Create the note
	now := time.Now()
	note := models.Note{
		Id:        uuid.New(),
		Content:   request.Content,
		CreatedAt: now,
		ExpiresAt: now.Add(time.Duration(expireTime) * time.Minute),
	}

	// Save the note in db
	if err := database.DB.Create(&note).Error; err != nil {
		return c.Status(500).SendString("Erreur lors de la création du note")
	}

	return c.JSON(fiber.Map{
		"message": "note created successfully",
		"uid":     note.Id,
	})
}

func GetNote(c *fiber.Ctx) error {
	uid := c.Params("uid")
	var note models.Note

	// Search if note exists
	result := database.DB.First(&note, "id = ?", uid)
	if result.Error != nil {
		return c.Status(404).JSON("Note not found")
	}
	// Check if note is expired
	if time.Now().After(note.ExpiresAt) {
		// Supprimer le paste expiré
		database.DB.Delete(&note)
		return c.Status(410).JSON("Note expired")
	}
	// Keep the content before deleting the note
	content := note.Content

	// Delete the note after use
	database.DB.Delete(&note)
	return c.JSON(fiber.Map{
		"content": content,
	})
}
