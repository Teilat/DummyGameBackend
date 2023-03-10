package webapi

import (
	"DummyGameBackend/webapi/converters"
	"DummyGameBackend/webapi/helpers"
	"DummyGameBackend/webapi/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

var LoginUser = ""

func newJwtMiddleware(webapi *WebApi, SendCookie bool) *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		SendCookie:  SendCookie,
		CookieName:  "jwt",
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TimeFunc:    time.Now,
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		// -------------------- login pipeline
		Authenticator: authFunc(webapi),
		PayloadFunc:   payloadFunc,
		LoginResponse: loginResponseFunc(webapi),
		//---------------------- auth pipeline
		IdentityHandler: identityFunc,
		Authorizator:    authorizatorFunc,
		Unauthorized:    unauthorizedFunc,
		LogoutResponse:  logoutResponseFunc,
	}
}

func authFunc(w *WebApi) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals models.Login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		if !helpers.CheckUserPass(w.database, loginVals) {
			return "", jwt.ErrFailedAuthentication
		}
		return &models.Login{
			Login: loginVals.Login,
		}, nil
	}
}

// структура внутри jwt
func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.Login); ok {
		LoginUser = v.Login
		return jwt.MapClaims{
			jwt.IdentityKey: v.Login,
		}
	}
	return jwt.MapClaims{}
}

func loginResponseFunc(w *WebApi) func(c *gin.Context, code int, message string, time time.Time) {
	return func(c *gin.Context, code int, message string, time time.Time) {
		c.Writer.Header().Add("Access-Token", message)
		c.Writer.Header().Add("Expire-Token", time.Format("2006-01-02 15:04:05"))
		c.JSON(code, converters.UserToApiUser(w.resolver.GetUserByUsername(LoginUser), message, time.Format("2006-01-02 15:04:05")))
	}
}

func identityFunc(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.Login{
		Login: claims[jwt.IdentityKey].(string),
	}
}

func authorizatorFunc(data interface{}, c *gin.Context) bool {
	return true
}

func unauthorizedFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func logoutResponseFunc(c *gin.Context, code int) {
	c.JSON(code, "")
}
