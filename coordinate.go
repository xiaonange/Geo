package Geo

import "fmt"

type Coordinate struct {
 Lat float64 `json:"lat"`
 Lng float64 `json:"lng"`
 Ellipsoid *Ellipsoid `json:"ellipsoid"`
}
/**
 * @param float $lat           -90.0 .. +90.0
 * @param float $lng           -180.0 .. +180.0
 * @param Ellipsoid $ellipsoid if omitted, WGS-84 is used
 *
 * @throws \InvalidArgumentException
 */
func (this Coordinate)Create(lat, lng float64, ellipsoid Ellipsoid) Coordinate {
	if ! isValidLatitude(lat) {
		panic(fmt.Sprintf("Latitude value must be numeric -90.0 .. +90.0 (given: {%f})", lat))
	}
	if ! isValidLongitude(lng) {
		panic(fmt.Sprintf("Longitude value must be numeric -180.0 .. +180.0 (given: {%f})", lat))
	}
	this.Lat = lat
	this.Lng = lng
	if ellipsoid.F != 0 {
		this.Ellipsoid = &ellipsoid
	} else {
		this.Ellipsoid = new(Ellipsoid).GetWGS84()
	}
	return this
}


/**
 * Validates latitude
 *
 * @param mixed latitude
 *
 * @return bool
 */
func isValidLatitude(latitude float64) bool{
return isNumericInBounds(latitude, - 90.0, 90.0)
}


/**
 * Validates longitude
 *
 * @param mixed longitude
 *
 * @return bool
 */
func isValidLongitude(longitude float64)bool {
	return isNumericInBounds(longitude, - 180.0, 180.0)
}

/**
 * Checks if the given value is (1) numeric, and (2) between lower
 * and upper bounds (including the bounds values).
 * @param float value
 * @param float lower
 * @param float upper
 * @return bool
 */
func isNumericInBounds(value, lower,upper float64) bool {
	if value < lower || value > upper {
		return false
	}
	return true
}
/**
 * @return float
 */
func (this Coordinate)GetLat() float64{
return this.Lat
}

/**
 * @return float
 */
func (this Coordinate)GetLng() float64{
	return this.Lng
}