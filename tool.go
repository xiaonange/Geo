package Geo

import (
	"math"
)

type Tool struct {
	StartBase Coordinate `json:"start_base"` //基准坐标
	Geofence  *Polygon   `json:"geofence"`   //坐标集合点
	BaseR     float64    `json:"base_r"`
}

const (
	R = 6371.393 * 1000    // 常量，地球半径，单位：米
	BaseAngle = 7070  //角度計算基數
)

func (this Tool) Create(base Coordinate,geofence *Polygon) Tool {
	this.StartBase = base
	this.Geofence = geofence
	this.BaseR = R * math.Cos(base.Lat*math.Pi/180) // 成员变量，初始化时计算，后面转换计算要用到
	return this
}

//经纬度转平面坐标
func (this Tool)JTP(input Coordinate) Coordinate {
	input.Lat = this.BaseR * (input.Lng - this.StartBase.Lng) * math.Pi / 180 // 单位：米
	input.Lng = R * (input.Lat - this.StartBase.Lat) * math.Pi / 180         // 单位：米
	return input
}

//平面坐标转经纬度
func (this Tool)PTJ(input Coordinate) Coordinate {
	input.Lng = input.Lng*180/(R*math.Pi) + this.StartBase.Lat
	input.Lat = input.Lat*180/(this.BaseR*math.Pi) + this.StartBase.Lng
	return input
}

//以圆心为中点旋转
func  (this Tool)GetCoordinate(center Coordinate,angle float64)Coordinate {
	center.Lat = center.Lat + 7070 *math.Cos(angle*math.Pi/180)
	center.Lat = center.Lng + 7070 *math.Sin(angle*math.Pi/180)
	return this.PTJ(center)
}
//计算中心点
func (this Tool)center(aPlace, bPlace Coordinate, r float64)Coordinate {
	x1 := aPlace.Lat
	y1 := aPlace.Lng
	x2 := bPlace.Lat
	y2 := bPlace.Lng
	C1 := (math.Pow(x2, 2) - math.Pow(x1, 2) + math.Pow(y2, 2) - math.Pow(y1, 2)) / 2 / (x2 - x1)
	C2 := (y2 - y1) / (x2 - x1)
	A := math.Pow(C2, 2) + 1
	B := 2*x1*C2 - 2*C1*C2 - 2*y1
	C := math.Pow(x1, 2) - 2*x1*C1 + math.Pow(C1, 2) + math.Pow(y1, 2) - math.Pow(r, 2)
	Delta := math.Pow(B, 2) - 4*A*C
	var y01, y02, x01, x02 float64
	if Delta >= 0 {
		Delta = math.Pow(Delta, 0.5)
		y01 = (-B + Delta) / 2 / A
		y02 = (-B - Delta) / 2 / A
		x01 = C1 - C2*y01
		x02 = C1 - C2*y02
	} else {
		panic("无解")
	}
	input := this.PTJ(Coordinate{Lat: x02, Lng: y02})
	if this.Geofence.Contains(input) {
		return Coordinate{Lat: x02, Lng: y02}
	} else {
		return Coordinate{Lat: x01, Lng: y01}
	}
}

//獲取角度
func GetAngle(distance float64)float64 {
	angle := math.Pow(BaseAngle, 2) + math.Pow(BaseAngle, 2) + math.Pow(distance, 2)/(2*BaseAngle*BaseAngle)
	return math.Acos(angle) / math.Pi * 180
}

//獲取兩點距離
func GetDistance(aPlace,bPlace Coordinate) float64 {
	vincenty := new(Vincenty)
	aPlace.Ellipsoid = new(Ellipsoid).GetWGS84()
	bPlace.Ellipsoid = new(Ellipsoid).GetWGS84()
	return vincenty.GetDistance(aPlace, bPlace)
}
//獲取圓弧點坐標
func (this Tool)GetPaints(point1,point2 Coordinate)[]Coordinate {
	point1 = this.JTP(point1)
	point2 = this.JTP(point2)
	center := this.center(point1, point2, BaseAngle)
	//零点
	c := point1.Lat - center.Lat
	b := point1.Lng - center.Lng
	bb := math.Asin(b/BaseAngle) * 180 / math.Pi
	angel := -math.Acos(c / BaseAngle) * 180 / math.Pi
	if bb > 0 {
		angel = -angel
	}
	distance := GetDistance(point1, point2)
	angle := GetAngle(distance) / 10
	var data []Coordinate
	var i float64
	for i = 1; i <= angle; i++ {
		data = append(data, this.GetCoordinate(center, angel-i*10))
	}
	return data
}