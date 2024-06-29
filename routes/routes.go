package routes

import (
	"github.com/gin-gonic/gin"
	"tutoringsystem/controllers"
)

func RegisterRoutes(r *gin.RouterGroup) {
	public := r.Group("/")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	protected := r.Group("/")
	//protected.Use(middlewares.Auth())
	protected.Use()
	{
		protected.GET("/tutors/:id", controllers.GetTutor)
		protected.POST("/tutors", controllers.CreateOrUpdateTutor)
		protected.GET("/tutors/certification", controllers.GetTutor)
		protected.GET("/students/:id", controllers.GetStudent)
		protected.POST("/students", controllers.CreateOrUpdateStudent)
		protected.GET("/schedules", controllers.GetSchedules)
		protected.POST("/schedules", controllers.UpdateSchedules)
		protected.GET("/matches", controllers.GetMatches)
		protected.POST("/matches", controllers.MatchStudentsAndTutors)
		protected.GET("/notifications", controllers.GetNotifications)
		protected.GET("/message/:userid", controllers.GetMessages)
		protected.POST("/send", controllers.SendMessage)
	}
}
