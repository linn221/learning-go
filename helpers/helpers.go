package helpers

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Dump(i interface{}) {
	fmt.Printf("%#v\n\n", i)
}

// returns zero for invalid input
func StrToUInt(id string) uint {
	id = strings.TrimSpace(id)
	ID, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return uint(ID)
}

func ProcessValidationErrors(err error) map[string]string {

	validationErrors := err.(validator.ValidationErrors)

	errorResponse := make(map[string]string)

	for _, ve := range validationErrors {
		errorResponse[ve.Field()] = ve.Tag()
	}

	return errorResponse
}

func SanitizeStr(unsafe string) string {
	return html.EscapeString(strings.TrimSpace(unsafe))
}

func VerifyPassword(password, hashedPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashStr(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// var input models.PurchaseOrder

// 	if err := ctx.ShouldBindJSON(&input); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
// 	}

// 	if err := validator.New().Struct(input); err != nil {
// 		errorResponse := helper.ProcessValidationErrors(err)

//         ctx.JSON(http.StatusBadRequest, gin.H{"error": errorResponse})
//         return
// 	}
