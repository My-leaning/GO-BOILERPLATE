package auth

import (
	"context"
	"go_boilerplate/internal/api/user"
	mongoDB "go_boilerplate/internal/database/mongodb"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type ResponseToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// @Summary Register a new user
// @Description Register a new user with a username and password
// @Accept  json
// @Produce  json
// @Param   credentials  body  auth.Credentials  true  "User credentials"
// @Success 201 {string} string "User created"
// @Failure 400 {string} string "User already exists"
// @Router /register [post]
func Register(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	user := user.User{Username: creds.Username, Password: string(hashedPassword)}

	_, err := mongoDB.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// @Summary Login a user
// @Description Login a user with username and password
// @Accept  json
// @Produce  json
// @Param   credentials  body  auth.Credentials  true  "User credentials"
// @Success 200 {string} string "Token"
// @Failure 401 {string} string "Unauthorized"
// @Router /login [post]
func Login(c *gin.Context) {
	var creds Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var storedCreds user.User
	err := mongoDB.UserCollection.FindOne(context.Background(), bson.M{"username": creds.Username}).Decode(&storedCreds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, err := createAccessToken(storedCreds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
		return
	}

	refreshToken, err := createRefreshToken(storedCreds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create refresh token"})
		return
	}

	responseToken := &ResponseToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, responseToken)
}

// NOTE - create access token
func createAccessToken(storedCreds user.User) (string, error) {
	jwtAccessKey := []byte(os.Getenv("JWT_ACCESS_SECRET"))

	accessExpirationTime := time.Now().Add(1 * time.Minute)
	payloadAccessToken := &Claims{
		Id:       storedCreds.ID.Hex(),
		Username: storedCreds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payloadAccessToken).SignedString(jwtAccessKey)
}

// NOTE - create refresh token
func createRefreshToken(storedCreds user.User) (string, error) {
	jwtRefreshKey := []byte(os.Getenv("JWT_REFRESH_SECRET"))

	accessExpirationTime := time.Now().Add(5 * time.Minute)
	payloadAccessToken := &Claims{
		Id:       storedCreds.ID.Hex(),
		Username: storedCreds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payloadAccessToken).SignedString(jwtRefreshKey)
}
