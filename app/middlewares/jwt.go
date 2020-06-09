package middlewares

//func Jwt() echo.MiddlewareFunc {
//	secret := os.Getenv("SUPER_MARIO_BROS")
//	return middleware.JWTWithConfig(middleware.JWTConfig{
//		SigningKey:    []byte(secret),
//		SigningMethod: "HS256",
//		ContextKey:    "token",
//		Claims:        &models.JwtClaims{},
//		TokenLookup:   "header:Authorization",
//		AuthScheme:    "Bearer",
//		Skipper: func(c echo.Context) bool {
//			if strings.Contains(c.Path(), "/login") {
//				return true
//			}
//			return false
//		},
//	})
//}
//
//func SetToken (user *models.User) (string, error){
//	secret := os.Getenv("SUPER_MARIO_BROS")
//	claim := models.JwtClaims{
//		ID:             user.Id,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
//	tokenString, err := token.SignedString([]byte(secret))
//
//	return tokenString, err
//}
//
//func GetToken (c echo.Context) *models.User{
//	token := c.Get("token").(*jwt.Token)
//	claim := token.Claims.(*models.JwtClaims)
//
//	if user := models.UserShowById(claim.ID) ; user != nil {
//		return user
//	}
//
//	return nil
//}

