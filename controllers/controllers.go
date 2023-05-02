package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/luongquochai/ecommerce_golang/database"
	"github.com/luongquochai/ecommerce_golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")

var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")

var Validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	//compare password from backend <db> with given password
	
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""

	if err!= nil {
        valid = false
        msg = "Login or Password is incorrect"
    }
	return valid, msg
}

func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone no is already in use"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshToken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, *&user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, inserterr := UserCollection.InsertOne(ctx, user)

		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not  get created"})
			return
		}

		defer cancel()

		c.JSON(http.StatusCreated, "Successfully sign in!!!")

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		// check if the user exists in the database
		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancel()

		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		// after password is valid
		token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, founduser.User_ID)
		defer cancel()

		generate.UpdateAllTokens(token, refreshToken, founduser.User_ID)

		c.JSON(http.StatusFound, founduser)

	}
}

func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

	return func(c *gin.Context) {
		// create product List
		var productList []models.Product
		var ctx, cancel = context.WithTimeOut(context.Background(), 100*time.Second)

		defer cancel()

		// BSON data 'D': An ordered representation of aBSON document (slice)
		cursor. err := ProductCollection.Find(ctx, bson.D{{}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong, please try after some time")
			return
		}

		err = cursor.All(ctx, &productList)
		if err!= nil {
            log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
        }

		defer cursor.Close()

		if err := cursor.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		defer cancel()

		c.IndentedJSON(200, productList)
	}
}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchProduct []models.Product
		queryParam := c.Query("name")

		// you want to check if it's empty

		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(HTTP.StatusNotFound, gin.H{"Error": "Invalid search index"})
			c.Abort()
			return
        }

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		searchQueryDB, err := ProductCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex": queryParam}})

		if err != nil {
			c.IndentedJSON(404, "Something went wrong while fetching the data")
			return
		}

		err = searchQueryDB.All(ctx, &searchProduct)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "Invalid")
			return
		}
		// close search query
		defer searchQueryDB.Close()

		if err := searchQueryDB.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "Invalid request")
			return
		}

		defer cancel()

		c.IndentedJSON(200, searchProduct)

	}
}
