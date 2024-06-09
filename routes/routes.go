package routes

import (
	"github.com/gin-gonic/gin"
	"tutoringsystem/controllers"
	"tutoringsystem/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	public := r.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.Auth())
	{
		protected.GET("/tutors/:id", controllers.GetTutor)
		protected.POST("/tutors", controllers.CreateOrUpdateTutor)
		protected.GET("/students/:id", controllers.GetStudent)
		protected.POST("/students", controllers.CreateOrUpdateStudent)
		protected.GET("/schedules", controllers.GetSchedules)
		protected.POST("/schedules", controllers.UpdateSchedules)
		protected.GET("/matches", controllers.GetMatches)
		protected.POST("/matches", controllers.MatchStudentsAndTutors)
		protected.GET("/notifications", controllers.GetNotifications)
	}
}
