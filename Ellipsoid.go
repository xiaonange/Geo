package Geo

type Ellipsoid struct {
	Name string  `json:"name"`
	A    float64 `json:"a"`
	F    float64 `json:"f"`
}
/**
 * Some often used ellipsoids
 * @var array
 */
func (Ellipsoid)GetWGS84() *Ellipsoid {
	return &Ellipsoid{
		Name: "World Geodetic System  1984",
		A:    6378137.0,
		F:    298.257223563}
}


func (Ellipsoid)GetGRS80()  *Ellipsoid {
	return &Ellipsoid{
		Name: "Geodetic Reference System 1980",
		A:    6378137.0,
		F:    298.257222100}
}

/**
 * @return string
 */
func (this Ellipsoid)GetName() string {
	return this.Name
}

/**
 * @return float
 */
func (this Ellipsoid)GetA() float64 {
	return this.A
}
/**
 * Calculation of the semi-minor axis
 *
 * @return float
 */
func (this Ellipsoid)GetB() float64 {
	return this.A* (1 - 1 / this.F)
}
/**
 * @return float
 */
func (this Ellipsoid)GetF() float64 {
	return this.F
}

/**
 * Calculates the arithmetic mean radius
 *
 * @see http://home.online.no/~sigurdhu/WGS84_Eng.html
 *
 * @return float
 */
func (this Ellipsoid)GetArithmeticMeanRadius() float64 {
	return this.A * (1 - 1/this.F/3)
}

