package main

import (
	"backer-startup/auth"
	"backer-startup/campaign"
	"backer-startup/handler"
	"backer-startup/helper"
	"backer-startup/payment"
	"backer-startup/transaction"
	"backer-startup/user"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"

	webHandler "backer-startup/web/handler"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:rifaimartin123@tcp(127.0.0.1:3306)/backer_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	// user, _ := userService.GetUserByID(1)

	// input := transaction.CreateTransactionInput{
	// 	CampaignID: 9,
	// 	Amount:     500000,
	// 	User:       user,
	// }

	// transactionService.CreateTransaction(input)

	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0fQ.gPQZuzlzKZsJevcIpEX_M5rjWbfhw_ZdPIdjaHd6IKE")
	// if err != nil {
	// 	fmt.Println("ERROR")
	// }

	// if token.Valid {
	// 	fmt.Println("VALID")
	// } else {
	// 	fmt.Println("INVALID")
	// }

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	userWebHandler := webHandler.NewUserHandler(userService)

	router := gin.Default()
	router.Use(cors.Default())

	// router.LoadHTMLGlob("web/templates/**/*") temp
	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")


	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailabilty)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	api.GET("campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransaction)
	api.GET("transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	router.GET("/users", userWebHandler.Index)
	router.GET("/users/new", userWebHandler.New)
	router.POST("/users", userWebHandler.Create)
	router.GET("/users/edit/:id", userWebHandler.Edit)
	router.POST("/users/update/:id", userWebHandler.Update)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			fmt.Println("Bearer empity")
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// split by space example Bearer = Token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// context currentUser
		c.Set("currentUser", user)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
