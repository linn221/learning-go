package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Dump(i interface{}) {
	fmt.Printf("%+v\n", i)
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

// var input models.PurchaseOrder

// 	if err := context.ShouldBindJSON(&input); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
// 	}

// 	if err := validator.New().Struct(input); err != nil {
// 		errorResponse := helper.ProcessValidationErrors(err)

//         context.JSON(http.StatusBadRequest, gin.H{"error": errorResponse})
//         return
// 	}
