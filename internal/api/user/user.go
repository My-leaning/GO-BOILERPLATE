package user

import (
	postgresDB "go_boilerplate/internal/database/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SECTION - mongoDB
// type User struct {
// 	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
// 	Username string             `bson:"username" json:"username"`
// 	Password string             `bson:"password,omitempty" json:"password"`
// }

// @Summary Get list of users
// @Description Get list of users
// @Produce  json
// @Success 200 {array} user.User
// @Router /api/users [get]
// func GetUsers(c *gin.Context) {
// 	var users []User

// 	cursor, err := database.UserCollection.Find(context.Background(), bson.D{})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	defer cursor.Close(context.Background())

// 	for cursor.Next(context.Background()) {
// 		var user User
// 		cursor.Decode(&user)
// 		user.Password = ""
// 		users = append(users, user)
// 	}

// 	c.JSON(http.StatusOK, users)
// }

// @Summary Create a new user
// @Description Create a new user with username and password
// @Accept  json
// @Produce  json
// @Param   user  body  user.User  true  "User data"
// @Success 201 {string} string "User created"
// @Router /api/user [post]
// func CreateUser(c *gin.Context) {
// 	var user User
// 	//NOTE: validate body request
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	fmt.Println("user", user)

// 	//NOTE: validate input key
// 	if user.Username == "" || user.Password == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
// 		return
// 	}

// 	//NOTE: validate user already exists
// 	var findUser ResPoneUser
// 	database.UserCollection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&findUser)

// 	if findUser.Username != "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
// 		return
// 	}

// 	//NOTE:  hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash password error"})
// 		return
// 	}

// 	user.Password = string(hashedPassword)

// 	//NOTE: create user
// 	result, err := database.UserCollection.InsertOne(context.Background(), user)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		return
// 	}

// 	user.ID = result.InsertedID.(primitive.ObjectID)

// 	//NOTE: set response
// 	responseUser := ResPoneUser{
// 		ID:       user.ID,
// 		Username: user.Username,
// 		Phone:    user.Phone,
// 	}

// 	//NOTE - create response
// 	c.JSON(http.StatusCreated, responseUser)
// }

// func GetUserById(c *gin.Context) {
// 	id := c.Param("id")
// 	fmt.Println("id", id)
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// 		return
// 	}

// 	var user ResPoneUser
// 	err = database.UserCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// func UpdateUser(c *gin.Context) {
// 	// id := c.Param("id")

// }
//SECTION - End of mongoDB

// !SECTION - postgres

func CreateUser(c *gin.Context) {
	var user postgresDB.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := postgresDB.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := postgresDB.DB.Delete(&postgresDB.User{}, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}

//!SECTION - End of postgres
