package main

import (
	"fmt"
	"log"
	"time"
	"vanishingnote/database"
	"vanishingnote/handlers"
	"vanishingnote/models"

	"github.com/gofiber/fiber/v2"
)

func startCleanupJob() {
	ticker := time.NewTicker(30 * time.Minute) // Cleanup every 30 minutes
	defer ticker.Stop()                        // Stop the ticker if the function stop for some reason

	for {
		<-ticker.C // Wait for the next tick sent in the channel
		result := database.DB.Where("expires_at < ?", time.Now()).Delete(&models.Note{})
		fmt.Printf("Cleanup : %d notes deleted at %v\n", result.RowsAffected, time.Now())
	}
}
func main() {
	app := fiber.New()
	database.ConnectDB()

	// Create the table if it doesn't exist already
	database.DB.AutoMigrate(&models.Note{})

	// Start the cleanup job in a separate goroutine
	go startCleanupJob()

	app.Static("/", "./static")

	// Api routes
	app.Post("/api/createNote", handlers.CreateNote)
	app.Get("/api/getNote/:uid", handlers.GetNote)

	// Route to read a note
	app.Get("/note/:uid", func(c *fiber.Ctx) error {
		return c.SendFile("static/views/read.html")
	})

	// Default route to create a new note
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("static/views/create.html")
	})

	// Redirect all other routes to the default route
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Redirect("/", fiber.StatusMovedPermanently)
	})

	port := ":9000"
	log.Printf("Server listening on port %s", port)
	log.Fatal(app.Listen(port))
}
