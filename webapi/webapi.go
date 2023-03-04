package webapi

import (
	"DummyGameBackend/internal/resolver"
	"DummyGameBackend/webapi/models"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type WebApi struct {
	address  string
	logger   *log.Logger
	database *gorm.DB
	resolver *resolver.Resolver
}

func NewWebapi(db *gorm.DB) *WebApi {
	return &WebApi{
		address:  fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port")),
		database: db,
		logger:   log.New(os.Stderr, "webapi", log.LstdFlags),
	}
}

func (w *WebApi) Start() {
	w.resolver = resolver.NewResolver(w.database)
	go w.run()
}

func (w *WebApi) run() {
	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))

	authMiddleware, err := jwt.New(newJwtMiddleware(w, true))
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("Auth middleware init error:" + errInit.Error())
	}

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(newCors()))

	authGroup := router.Group("")
	authGroup.Use(authMiddleware.MiddlewareFunc())

	router.GET("/", w.HandlePing())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/register", w.RegisterHandler())
	authGroup.GET("/logout", authMiddleware.LogoutHandler)
	authGroup.GET("/user", w.GetUser())

	err = router.Run(address)
	//err := router.RunTLS(address, "./server-cert.pem", "./server-key.pem")
	if err != nil {
		log.Fatal(err)
	}
}

// HandlePing   godoc
// @Summary		Health check
// @Tags        General
// @Accept      json
// @Produce     json
// @Success     200 {string} string "healthy"
// @Router      / [get]
func (w *WebApi) HandlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

func (w *WebApi) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

// RegisterHandler  godoc
// @Summary     register user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       user body models.AddUser true "user"
// @Success     200
// @Error       500 {string} string
// @Router      /register [post]
func (w *WebApi) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.AddUser
		err := c.BindJSON(&user)
		if err != nil {
			w.logger.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to parse params"})
		}

		err = w.resolver.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to register"})
		}
	}
}
