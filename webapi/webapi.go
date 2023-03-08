package webapi

import (
	"DummyGameBackend/internal/resolver"
	_ "DummyGameBackend/webapi/docs"
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
	w.run()
}

// @Title     Application Api
// @Version   1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
	router.POST("/register", w.Register())
	authGroup.GET("/logout", authMiddleware.LogoutHandler)

	authGroup.GET("/characters", w.Characters())
	authGroup.GET("/character/{id}", w.GetCharacter())
	authGroup.GET("/character/create", w.CreateCharacter())
	authGroup.GET("/character/update", w.UpdateCharacter())
	authGroup.GET("/character/delete", w.DeleteCharacter())

	err = router.Run(address)
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

// Login  godoc
// @Summary     Login user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       credentials body models.Login true "credentials"
// @Success     200 {object} models.User "logged in user"
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /login [post]
func (w *WebApi) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// Register  godoc
// @Summary     register user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       user body models.AddUser true "user"
// @Success     200
// @Error       500 {string} string
// @Router      /register [post]
func (w *WebApi) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.AddUser
		err := c.BindJSON(&user)
		if err != nil {
			w.logger.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to parse params"})
		}

		err = w.resolver.CreateUser(user)
		if err != nil {
			w.logger.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to register"})
		}
	}
}

// Characters  godoc
// @Summary     get all user characters
// @Tags        Character
// @Accept      json
// @Produce     json
// @Success     200
// @Error       500 {string} string
// @Router      /characters [get]
func (w *WebApi) Characters() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

// GetCharacter  godoc
// @Summary     get character with id
// @Tags        Character
// @Accept      json
// @Produce     json
// @Param       id path int true "character id"
// @Success     200
// @Error       500 {string} string
// @Router      /character{id} [get]
func (w *WebApi) GetCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

// CreateCharacter  godoc
// @Summary     create character
// @Tags        Character
// @Accept      json
// @Produce     json
// @Param       character body models.AddCharacter true "character"
// @Success     200
// @Error       500 {string} string
// @Router      /character/create [post]
func (w *WebApi) CreateCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

// UpdateCharacter  godoc
// @Summary     update character
// @Tags        Character
// @Accept      json
// @Produce     json
// @Param       character body models.UpdateCharacter true "character"
// @Success     200
// @Error       500 {string} string
// @Router      /character/update [post]
func (w *WebApi) UpdateCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

// DeleteCharacter  godoc
// @Summary     delete character
// @Tags        Character
// @Accept      json
// @Produce     json
// @Param       id path int true "character id"
// @Success     200
// @Error       500 {string} string
// @Router      /character/delete/{id} [post]
func (w *WebApi) DeleteCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
