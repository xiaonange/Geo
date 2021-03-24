package Geo

import (
	"sort"
)

type Bounds struct {
	NorthWest Coordinate `json:"north_west"`
	SouthEast Coordinate `json:"south_east"`
 }

/**
 * @param Coordinate northWest
 * @param Coordinate southEast
 */
func (this Bounds)Create(northWest,southEast Coordinate)Bounds {
	this.NorthWest = northWest
	this.SouthEast = southEast
	return this
}

/**
 * Getter
 * @return Coordinate
 */
func (this Bounds)GetNorthWest()Coordinate {
	return this.NorthWest
}

/**
 * Getter
 *
 * @return Coordinate
 */
func (this Bounds)GetSouthEast()Coordinate {
	return this.SouthEast
}
/**
 * @return float
 */
func (this Bounds)GetNorth() float64{
return this.NorthWest.GetLat()
}

/**
 * @return float
 */
func (this Bounds)GetSouth() float64{
	return this.SouthEast.GetLat()
}
/**
 * @return float
 */
func (this Bounds)GetWest() float64{
	return this.NorthWest.GetLng()
}
/**
 * @return float
 */
func (this Bounds)GetEast() float64{
	return this.SouthEast.GetLng()
}
/**
 * Calculates the center of this bounds object and returns it as a
 * Coordinate instance.
 *
 * @return Coordinate
 */
func (this Bounds)GetCenter() Coordinate {

	centerLat := (this.GetNorth() + this.GetSouth()) / 2
	if this.NorthWest.Ellipsoid == nil {
		this.NorthWest.Ellipsoid = new(Ellipsoid).GetWGS84()
	}
	return Coordinate{Lat: centerLat, Lng: this.GetCenterLng(),Ellipsoid:this.NorthWest.Ellipsoid}
}

/**
 * @return float
 */
func (this Bounds)GetCenterLng() float64 {
	centerLng := (this.GetEast() + this.GetWest()) / 2
	overlap := this.GetWest() > 0 && this.GetEast() < 0
	if overlap && centerLng > 0 {
		return -180 + centerLng
	}
	if overlap && centerLng < 0 {
		return 180 + centerLng
	}
	if overlap && centerLng == 0 {
		return 180
	}
	return centerLng
}

/*
sort by points
 */
func SortByPoints(points []Coordinate)([]float64,[]float64) {
	var lat []float64
	var lng []float64
	for _, v := range points {
		lat = append(lat, v.Lat)
		lng = append(lng, v.Lng)
	}
	sort.Float64s(lat)
	sort.Float64s(lng)
	return lat, lng
}
