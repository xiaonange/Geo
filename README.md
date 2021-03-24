# Geo
##golang地理实用计算包
* 点对点距离计算
* 获取一串坐标点中心点
* 获取两点之间的圆弧坐标点
* 判断点是否在多边形内
```
     //判斷是否在區域內
	polygon := new(Geo.Polygon)
	polygon = polygon.AddPoint(Geo.Coordinate{Lng: 116.248445, Lat:39.997954})
	polygon = polygon.AddPoint(Geo.Coordinate{Lng: 116.52585, Lat:40.002162})
	polygon = polygon.AddPoint(Geo.Coordinate{Lng: 116.536836, Lat:39.791445})
	polygon = polygon.AddPoint(Geo.Coordinate{Lng: 116.226472, Lat:39.812546})
	fmt.Println(polygon.Contains(Geo.Coordinate{Lng:116.586275, Lat:39.869486}))//圈外116.586275,39.869486  圈内116.39676,39.896885 


```
	//獲取兩點之間距離
	vincenty := new(Geo.Vincenty)
	lat1 := Geo.Coordinate{Lng: 116.248445, Lat:39.997954,Ellipsoid:new(Geo.Ellipsoid).GetWGS84()}
	lat2 :=Geo.Coordinate{Lng:116.586275, Lat:39.869486,Ellipsoid:new(Geo.Ellipsoid).GetWGS84()}
	fmt.Println(vincenty.GetDistance(lat1,lat2))

```
	//獲取多個點的中心點
	var points []Geo.Coordinate
	points = append(points,Geo.Coordinate{Lng: 116.398183, Lat:39.928847})
	points = append(points,Geo.Coordinate{Lng: 116.407885, Lat:39.928847})
	points = append(points,Geo.Coordinate{Lng: 116.408388, Lat:39.919938})
	points = append(points,Geo.Coordinate{Lng: 116.39883, Lat:39.919938})
	lats,lngs :=Geo.SortByPoints(points)
	northWest := Geo.Coordinate{Lat:lats[len(lats)-1], Lng:lngs[0]}
	southEast := Geo.Coordinate{Lat:lats[0], Lng:lngs[len(lngs)-1]}
	bounds := new(Geo.Bounds).Create(northWest, southEast).GetCenter() //圆心
	radius := new(Geo.Vincenty).GetDistance(Geo.Coordinate{Lat:points[0].Lat, Lng:points[0].Lng,Ellipsoid:new(Geo.Ellipsoid).GetWGS84()},bounds) //半径
	fmt.Println(bounds,radius)
