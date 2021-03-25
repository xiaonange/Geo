package Geo

type Polygon struct {
	Points []Coordinate `json:"points"`
}

/**
 * @param Coordinate point
 */
func (this Polygon)AddPoint(point Coordinate) *Polygon {
	this.Points = append(this.Points, point)
	return &this
}

/**
 * @return array
 */
func (this Polygon)GetPoints()[]Coordinate {
	return this.Points
}

/**
 * Return all polygon point's latitudes.
 * @return float
 */
func (this Polygon)GetLats()[]float64 {
	var lats []float64
	for _, v := range this.Points {
		lats = append(lats, v.Lat)
	}
	return lats
}

/**
 * Return all polygon point's longitudes.
 *
 * @return float
 */
func (this Polygon)GetLngs()[]float64 {
	var lngs []float64
	for _, v := range this.Points {
		lngs = append(lngs, v.Lng)
	}
	return lngs
}

/**
 * @return int
 */
func (this Polygon)getNumberOfPoints()int {
	return len(this.Points)
}

/**
 * Determine if given point is contained inside the polygon. Uses the PNPOLY
 * algorithm by W. Randolph Franklin. Therfore some edge cases may not give the
 * expected results, e. g. if the point resides on the polygon boundary.
 * @see http://www.ecse.rpi.edu/Homepages/wrf/Research/Short_Notes/pnpoly.html
 * For special cases this calculation leads to wrong results:
 * - if the polygons spans over the longitude boundaries at 180/-180 degrees
 * @param Coordinate point
 * @return boolean
 */
func (this Polygon)Contains(point Coordinate) bool {
	numberOfPoints := this.getNumberOfPoints()
	polygonLats := this.GetLats()  //经度
	polygonLngs := this.GetLngs() //纬度
	polygonContainsPoint := false
	altNode := numberOfPoints - 1
	for node := 0; node < numberOfPoints; node++ {
		if ((polygonLngs[node] > point.Lng != (polygonLngs[altNode] > point.Lng) &&
			(point.Lat < (polygonLats[altNode]-polygonLats[node])*(point.Lng-polygonLngs[node])/(polygonLngs[altNode]-polygonLngs[node])+polygonLats[node]))) {
			polygonContainsPoint = ! polygonContainsPoint
		}
		altNode = node
	}

	return polygonContainsPoint
}