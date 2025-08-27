package routes

import (
	"github.com/akhilr007/fem-api-project/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(app *app.Application) *chi.Mux {

	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutById)
	return r
}
