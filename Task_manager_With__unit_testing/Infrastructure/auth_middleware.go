package infrastructure

import (
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Replace with your actual secret key
const JWT_SECRET = "your_secret_key_here"

// The middleware for Authentication
func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization") // Extracting the authentication value from the header
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.JSON(401, gin.H{"error": "Invalid authorization header"})
		c.Abort()
		return
	}

	tokenString := authParts[1]             // The token string
	token, err := TokenClaimer(tokenString) // Verifying the token

	if err != nil {
		log.Println("Token parsing error:", err.Error())
		c.JSON(401, gin.H{"error": "Invalid JWT"})
		c.Abort()
		return
	}

	// Extracting the map claims from the token
	if role, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(role["_id"])
		c.Set("isadmin", role["isadmin"])
		c.Set("userid", role["_id"])
	} else {
		c.JSON(401, gin.H{"error": "Invalid JWT"})
		c.Abort()
		return
	}

	c.Next()
}

// The middleware for Authentication
func AdminMiddleware(c *gin.Context) {
	isAdmin, exists := c.Get("isadmin") // Fetching the data from the context
	if !exists || !isAdmin.(bool) {
		c.JSON(403, gin.H{"error": "Forbidden: You don't have admin privileges"})
		c.Abort()
		return
	}

	c.Next()
}

// TokenGenerator generates a JWT token with the provided details.
func TokenGenerator(id primitive.ObjectID, email string, isadmin bool) (string, error) {
	var jwtSecret = []byte(JWT_SECRET)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":     id.Hex(),
		"email":   email,
		"isadmin": isadmin,
	})

	return token.SignedString(jwtSecret)
}

// TokenClaimer parses and validates the JWT token.
func TokenClaimer(tokenstr string) (*jwt.Token, error) {
	return jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_SECRET), nil
	})
}
