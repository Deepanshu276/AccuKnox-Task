package router

import (
	"github.com/Deepanshu276/AccuKnox-Task/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controller.Signup)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)

	app.Get("/api/notes", controller.GetNotes)
	app.Get("/api/note/:NoteID", controller.GetNoteById)
	app.Put("/api/update-note/:NoteID", controller.UpdateNoteById)
	app.Post("/api/create-note", controller.CreateNote)
	app.Delete("/api/delete-note/:NoteID", controller.DeleteNote)

}