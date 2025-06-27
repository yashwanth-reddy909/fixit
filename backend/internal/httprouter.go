package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fixit.com/backend/internal/auth"
	chttp "fixit.com/backend/src/controllers/http"
	mrepo "fixit.com/backend/src/repo/mongo"
	service "fixit.com/backend/src/service"

	// mongo driver imports

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func StartHttpServer(mongoClient *mongo.Client) error {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	})

	// google auth
	googleAuth := auth.CreateGoogleAuth(
		"54384433304-cq2q6nh3hukkf5b2s0ppsmpos6rn9i9h.apps.googleusercontent.com",
		"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",
		"http://localhost:8080/auth/google/callback",
	)

	// repos
	userRepo := mrepo.CreateUserRepo(mongoClient, "fixit", "users")

	// services
	userSvc := service.CreateUserSvc(userRepo, googleAuth)

	// controllers
	userController := chttp.NewUserController(*userSvc)

	// routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", userController.Signup)
		authRoutes.POST("/login", userController.Login)
		authRoutes.POST("/google/callback", userController.GoogleAuthCallback)
	}

	return r.Run(":8000")
}
