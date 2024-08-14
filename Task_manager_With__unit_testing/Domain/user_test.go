package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUser_Initialization(t *testing.T) {
	// Set up test data
	id := primitive.NewObjectID()

	// Initialize the entity
	user := User{
		ID:       id,
		Email:    "test@example.com",
		Password: "hashedpassword",
		IsAdmin:  false,
	}

	// Verify that the entity fields are set correctly
	assert.Equal(t, id, user.ID)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "hashedpassword", user.Password)
	assert.Equal(t, false, user.IsAdmin)
}

func TestUser_JSONSerialization(t *testing.T) {
	// Set up test data
	id := primitive.NewObjectID()

	// Initialize the entity
	user := User{
		ID:       id,
		Email:    "test@example.com",
		Password: "hashedpassword",
		IsAdmin:  true,
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(user)
	assert.NoError(t, err)

	// Deserialize back to a User
	var deserializedUser User
	err = json.Unmarshal(jsonData, &deserializedUser)
	assert.NoError(t, err)

	// Ensure the deserialized entity matches the original
	assert.Equal(t, user.ID.Hex(), deserializedUser.ID.Hex())
	assert.Equal(t, user.Email, deserializedUser.Email)
	assert.Equal(t, user.Password, deserializedUser.Password)
	assert.Equal(t, user.IsAdmin, deserializedUser.IsAdmin)
}
