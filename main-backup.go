package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // var number int64 = 10
// // var reflectValue = reflect.ValueOf(number)

// type Users struct {
// 	Name     string `required:"true"`
// 	Password string `required:"false"`
// 	Username string `required:"true"`
// 	Address  string `required:"true"`
// 	Age      int    `required:"true"`
// }

// func ValidateStruct(s interface{}) error {
// 	val := reflect.TypeOf(s)
// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)
// 		if field.Tag.Get("required") == "true" {
// 			value := reflect.ValueOf(s).Field(i).Interface()
// 			if value == "" {
// 				return fmt.Errorf("%s is required", field.Name)
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	newUser := Users{
// 		Name:     "",
// 		Password: "",
// 		Username: "Ags",
// 		Address:  "Jakarta",
// 		Age:      28,
// 	}

// 	userValidate := ValidateStruct(newUser)
// 	fmt.Println(userValidate)
// 	// fmt.Println("tipe :", reflectValue.Type())
// 	// fmt.Println("tipe :", reflectValue.Kind())
// 	// if reflectValue.Kind() == reflect.Int64 {
// 	// 	fmt.Println(reflectValue.Int())
// 	// }
// }
