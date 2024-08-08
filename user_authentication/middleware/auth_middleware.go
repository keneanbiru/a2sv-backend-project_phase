package middleware
import(
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"user_authentication/models"
	"github.com/dgrijalva/jwt-go"

)



func AuthMiddleware() gin.HandlerFunc {
	var JwtSecret = []byte("your_secret_key") 
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }

        // Extract token part (after "Bearer ")
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Parse and validate the token
        token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
            return JwtSecret, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Extract claims
        if claims, ok := token.Claims.(*models.Claims); ok {
            c.Set("ID", claims.ID)  // Ensure this key matches what you use in the AddTask method
            c.Set("Role", claims.Role)
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract claims"})
            c.Abort()
            return
        }

        c.Next()
    }
}
