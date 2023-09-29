package services

import (
	"context"

	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/20-VIGNESH-K/crud/config"
	"github.com/20-VIGNESH-K/crud/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Ctx         context.Context
	MongoClient *mongo.Client
)

// CustomValidator is a custom validation function
func CustomValidator(fl validator.FieldLevel) bool {
	// Use a regular expression to check if the field contains only letters
	name := fl.Field().String()
	match, _ := regexp.MatchString("^[A-Za-z]+$", name)
	return match
}

func CheckValidation(user models.Profile) bool {
	validate := validator.New()
	validate.RegisterValidation("customValidator", CustomValidator)

	err := validate.Struct(user)
	if err != nil {
		// Handle validation errors
		for _, validationErr := range err.(validator.ValidationErrors) {
			log.Printf("Validation Error in field %s: %s\n", validationErr.Field(), validationErr.Tag())
			return false
		}
	}

	log.Println("Validation successful")
	return true

}

func Create(context *gin.Context) {

	var user models.Profile
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}
	filter := bson.M{"name": user.Name}
	existingProfile := config.ProfileCollection.FindOne(Ctx, filter)
	if existingProfile.Err() == nil {
		context.JSON(http.StatusConflict, gin.H{"message": "Profile with the same name already exists"})
		return // Return early if the profile already exists
	}
	fmt.Println(user)
	err := CheckValidation(user)
	if err == false {
		context.JSON(http.StatusNotAcceptable, gin.H{"message": "Validation error"})
	} else {

		create, err := config.ProfileCollection.InsertOne(Ctx, &user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(create.InsertedID)
		context.JSON(http.StatusOK, gin.H{"message": "success"})
	}

}
