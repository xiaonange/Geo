package Geo

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	//判斷是否在區域內
	polygon := new(Polygon)
	polygon = polygon.AddPoint(Coordinate{Lng: 116.248445, Lat:39.997954})
	polygon = polygon.AddPoint(Coordinate{Lng: 116.52585, Lat:40.002162})
	polygon = polygon.AddPoint(Coordinate{Lng: 116.536836, Lat:39.791445})
	polygon = polygon.AddPoint(Coordinate{Lng: 116.226472, Lat:39.812546})
	fmt.Println(polygon.Contains(Coordinate{Lng:116.586275, Lat:39.869486}))//圈外116.586275,39.869486  圈内116.39676,39.896885
}

func TestGetDistance(t *testing.T) {
	//獲取兩點之間距離
	vincenty := new(Vincenty)
	lat1 := Coordinate{Lng: 116.248445, Lat:39.997954,Ellipsoid:new(Ellipsoid).GetWGS84()}
	lat2 :=Coordinate{Lng:116.586275, Lat:39.869486,Ellipsoid:new(Ellipsoid).GetWGS84()}
	fmt.Println(vincenty.GetDistance(lat1,lat2))	
}




func TestGetCenter(t *testing.T)  {
	//獲取多個點的中心點
	var points []Coordinate
	points = append(points,Coordinate{Lng: 116.398183, Lat:39.928847})
	points = append(points,Coordinate{Lng: 116.407885, Lat:39.928847})
	points = append(points,Coordinate{Lng: 116.408388, Lat:39.919938})
	points = append(points,Coordinate{Lng: 116.39883, Lat:39.919938})
	lats,lngs :=SortByPoints(points)
	northWest := Coordinate{Lat:lats[len(lats)-1], Lng:lngs[0]}
	southEast := Coordinate{Lat:lats[0], Lng:lngs[len(lngs)-1]}
	bounds := new(Bounds).Create(northWest, southEast).GetCenter() //圆心
	radius := new(Vincenty).GetDistance(Coordinate{Lat:points[0].Lat, Lng:points[0].Lng,Ellipsoid:new(Ellipsoid).GetWGS84()},bounds) //半径
	fmt.Println(bounds,radius)
}
