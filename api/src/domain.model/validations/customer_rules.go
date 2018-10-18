package customerdomainvalidation

import (
	"github.com/thedevsaddam/govalidator"
)

//Validate return customers validate
func Validate(obj interface{}) (map[string]interface{}, bool) {

	rules := govalidator.MapData{
		"name": []string{"required", "between:3,150"},
		//"email": []string{"required", "min:4", "max:20", "email"},
		//"web":   []string{"url"},
		//"age":   []string{"numeric_between:18,56"},
	}

	options := govalidator.Options{
		Data:  &obj,
		Rules: rules,
	}

	validator := govalidator.New(options)

	validateJSON := validator.ValidateJSON()
	mapsErrors := map[string]interface{}{"validationError": validateJSON}

	hasErros := len(mapsErrors) > 0

	return mapsErrors, hasErros
}
