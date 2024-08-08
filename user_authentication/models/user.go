package models
import(

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/dgrijalva/jwt-go"
	
)

type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"password"`
	Role        string             `bson:"role" json:"role"`
}


// Claims represents the JWT claims
type Claims struct {
	ID string `json:"userid"`
	Role   string `json:"role"`
	jwt.StandardClaims
}