package Golang_Validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	user := "Hanif"

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()
	password := "Benar"
	confPass := "Benar"

	err := validate.VarWithValue(password, confPass, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "02020"

	err := validate.Var(user, "required,number,min=5,max=10")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	res := LoginRequest{
		Username: "hanif@gmail.com",
		Password: "12",
	}

	err := validate.Struct(res)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestCrossField(t *testing.T) {
	type LoginRequest struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}

	validate := validator.New()

	res := LoginRequest{
		Username:        "hanif@gmail.com",
		Password:        "12345",
		ConfirmPassword: "12345",
	}

	err := validate.Struct(res)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		ID      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()

	res := User{
		ID:   "",
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	err := validate.Struct(res)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}
func TestValidationCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		ID      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
	}

	validate := validator.New()

	res := User{
		ID:   "",
		Name: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	err := validate.Struct(res)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		ID      string    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
		Hobbies []string  `validate:"required,dive,required,min=3"`
	}

	validate := validator.New()

	res := User{
		ID:   "",
		Name: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"X",
		}}

	err := validate.Struct(res)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		ID      string            `validate:"required"`
		Name    string            `validate:"required"`
		Address []Address         `validate:"required,dive"`
		Hobbies []string          `validate:"dive,required,min=3"`
		Schools map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	}

	validate := validator.New()

	res := User{
		ID:   "1",
		Name: "Hanif",
		Address: []Address{
			{
				City:    "Jakarta",
				Country: "Indonesia",
			},
			{
				City:    "Bandung",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"Reading",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD Indonesia",
			},
			"SMP": {
				Name: "SMP Indonesia",
			},
			"SMA": {
				Name: "SMA Indonesia",
			},
		},
	}

	err := validate.Struct(res)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=200")

	type Seller struct {
		Id     string `validate:"varchar"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "",
		Name:   "",
		Owner:  "",
		Slogan: "",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "HANIF",
		Password: "",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestValidationParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	req := Login{
		Phone: "64783912065781239",
		Pin:   "273812",
	}

	err := validate.Struct(req)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	validate := validator.New()

	req := Login{
		Username: "hanif",
		Password: "",
	}

	err := validate.Struct(req)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("Field not ok")
	}

	firstValue := strings.ToUpper(field.Field().Interface().(string))
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("ignore_case", MustEqualsIgnoreCase)

	type Login struct {
		Username string `validate:"required,ignore_case=Email|ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := Login{
		Username: "748291478921",
		Email:    "hanif@gmail.com",
		Phone:    "748291478921",
		Name:     "Hanif",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}

}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegistration(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {

	} else {
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegistration, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "hanif@ex.com",
		Email:    "hanif@ex.com",
		Phone:    "1546436347",
		Password: "rahasia",
	}

	err := validate.Struct(registerRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}
