package utility

import (
	"policymanagement/models"

	"github.com/mitchellh/mapstructure"
)

func StructToMapVehicle(v *models.Vehicle) map[string]interface{} {
	var vehicleMap map[string]interface{}
	_ = mapstructure.Decode(v, &vehicleMap)

	return vehicleMap

}

func StructToMapLocation(p *models.Location) map[string]interface{} {
	var locationMap map[string]interface{}
	_ = mapstructure.Decode(p, &locationMap)
	return locationMap
}
