package pkg

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if err := registerValidators(); err != nil {
		log.Printf("Failed to register validators: %v", err)
	}
	log.Println("Validators registered successfully")
}

// ethWalletValidator validates Ethereum address format (starts with 0x, followed by 40 hexadecimal characters)
func ethWalletValidator(fl validator.FieldLevel) bool {
	wallet := fl.Field().String()
	matched, _ := regexp.MatchString("^0x[0-9a-fA-F]{40}$", wallet)
	return matched
}

func registerValidators() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("ethwallet", ethWalletValidator)
		if err != nil {
			return fmt.Errorf("error registering custom validators: %v", err)
		}
	}
	return nil
}
