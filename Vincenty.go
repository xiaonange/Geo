package Geo

import (
	"fmt"
	"math"
)

type Vincenty struct {

}

/**
 * @param Coordinate point1
 * @param Coordinate point2
 * @return float
 */
func (Vincenty)GetDistance(point1, point2 Coordinate)float64 {
	if point1.Ellipsoid == nil || point2.Ellipsoid == nil|| point1.Ellipsoid.A != point2.Ellipsoid.A  {
		panic("The ellipsoids for both coordinates must match")
	}
	//角度转化为弧度
	lat1 := deg2rad(point1.Lat)
	lat2 := deg2rad(point2.Lat)
	lng1 := deg2rad(point1.Lng)
	lng2 := deg2rad(point2.Lng)

	a := point1.Ellipsoid.GetA()
	b := point1.Ellipsoid.GetB()
	f := 1 / point1.Ellipsoid.GetF()

	L := lng2 - lng1
	U1 := math.Atan((1 - f) * math.Tan(lat1))
	U2 := math.Atan((1 - f) * math.Tan(lat2))

	iterationLimit := 100
	lambda := L

	sinU1 := math.Sin(U1)
	sinU2 := math.Sin(U2)
	cosU1 := math.Cos(U1)
	cosU2 := math.Cos(U2)

	var cosSqAlpha float64
	var sinSigma float64
	var cos2SigmaM float64
	var cosSigma float64
	var sigma float64
	for {
		sinLambda := math.Sin(lambda)
		cosLambda := math.Cos(lambda)
		sinSigma = math.Sqrt((cosU2*sinLambda)*(cosU2*sinLambda) + (cosU1*sinU2-sinU1*cosU2*cosLambda)*(cosU1*sinU2-sinU1*cosU2*cosLambda))
		if sinSigma == 0 {
			return 0.0
		}
		cosSigma = sinU1*sinU2 + cosU1*cosU2*cosLambda
		sigma = math.Atan2(sinSigma, cosSigma)
		sinAlpha := cosU1 * cosU2 * sinLambda / sinSigma
		cosSqAlpha = 1 - sinAlpha*sinAlpha
		if cosSqAlpha != 0 {
			cos2SigmaM = cosSigma
		}
		C := f / 16 * cosSqAlpha * (4 + f*(4-3*cosSqAlpha))
		lambdaP := lambda
		lambda = L + (1-C)*f*sinAlpha*(sigma+C*sinSigma*(cos2SigmaM+C*cosSigma*(- 1+2*cos2SigmaM*cos2SigmaM)))
		iterationLimit--
		if math.Abs(lambda-lambdaP) > 1e-12 && iterationLimit > 0 {
			continue
		} else {
			break
		}
	}
	if iterationLimit == 0 {
		panic("")
	}
	uSq := cosSqAlpha * (a*a - b*b) / (b * b)
	A := 1 + uSq/16384*(4096+uSq*(- 768+uSq*(320-175*uSq)))
	B := uSq / 1024 * (256 + uSq*(- 128+uSq*(74-47*uSq)))
	deltaSigma := B * sinSigma * (cos2SigmaM + B/4*(cosSigma*(- 1+2*cos2SigmaM*cos2SigmaM)-B/6*cos2SigmaM*(- 3+4*sinSigma*sinSigma)*(- 3+4*cos2SigmaM*cos2SigmaM)))
	s := b * A * (sigma - deltaSigma)
	fmt.Println(s)
	return math.Floor(s)
}

//角度轉化為弧度
func deg2rad(deg float64) float64 {
	return deg * math.Pi / 180
}

//弧度轉化為角度
func  rad2deg(rad float64) float64 {
	return 180 * rad / math.Pi
}