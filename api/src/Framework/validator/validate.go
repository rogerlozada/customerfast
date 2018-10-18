package validate

import (
	"customerfast/api/src/framework/common"

	"github.com/thedevsaddam/govalidator"
)

//Validate return customers validate
func Validate(rules govalidator.MapData, obj interface{}) (map[string]interface{}, bool) {

	options := govalidator.Options{
		Data:  &obj,
		Rules: rules,
	}

	validator := govalidator.New(options)

	validateJSON := validator.ValidateJSON()
	mapsErrors := map[string]interface{}{"validationError": validateJSON}

	return mapsErrors, common.ContainsValue(mapsErrors)
}
