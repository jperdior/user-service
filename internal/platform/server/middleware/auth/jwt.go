package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// JWTMiddleware is a middleware that checks for a valid JWT token in the Authorization header
func JWTMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		// remove Bearer
		tokenString = tokenString[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method (optional, but recommended)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			// Return the secret key for validation
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set claims in context
			c.Set("claims", claims) // You can access claims later using c.MustGet("claims")
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("ID", c.MustGet("claims").(jwt.MapClaims)["ID"])
		c.Set("roles", c.MustGet("claims").(jwt.MapClaims)["roles"])
		c.Set("name", c.MustGet("claims").(jwt.MapClaims)["name"])
		c.Set("email", c.MustGet("claims").(jwt.MapClaims)["email"])

		c.Next()
	}
}
