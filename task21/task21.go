package main

import "fmt"

// Патерн адаптер необходим для совместной работы несовместимых интерфейсов.

// Интерфейс дорога, который содержит метод движения по дроге DriveOnTheRoad.
type road interface {
	DriveOnTheRoad()
}

// Класс водитель, который может двигаться по этой дороге с помощью метода DrivingByCar.
type driver struct {
}

func (d driver) DrivingByCar(r road) {
	r.DriveOnTheRoad()
}

// Класс автомобильные колёса. Обычно с помощью этих колёс происходит движение по дороге.
type carWheels struct {
}

func (c carWheels) DriveOnTheRoad() {
	fmt.Println("Traffic started on the road")
}

// Интерфейс железная дорога, который содержит метод движения по рельсам DriveOnTheRail.
type railway interface {
	DriveOnTheRail()
}

// Класс водитель поезда, который может двигаться по железной дороге с помощью метода DrivingByTrain.
type trainDriver struct {
}

func (td trainDriver) DrivingByTrain(rw railway) {
	rw.DriveOnTheRail()
}

// Класс колёсная пара. С помощью метода колёсной пары DriveOnTheRail можно перемещаться по железной дороге.
type wheelset struct {
}

func (w wheelset) DriveOnTheRail() {
	fmt.Println("Railway traffic started")
}

// Адаптер, который позволяет сменить колёса с автомобильных на колёсную пару и двигаться по железной дороге.
// То есть использовать метод DriveOnTheRail водителем driver, который не описан в методах интерфейса road.
type RailwayAdapter struct {
	carTrain *wheelset
}

// Для того, чтобы можно было использовать движение по джелезной дороге DriveOnTheRail классом driver,
// необходимо, чтобы у нас присутствовал одноимённый метод DriveOnTheRoad.
// Таким образом интерфейс road не будет знать о том, что произошла подмена методов.
// Для него всё будет происходить в штатном режиме.
func (adapter *RailwayAdapter) DriveOnTheRoad() {
	fmt.Println("Installing a wheelset instead of car wheels")
	adapter.carTrain.DriveOnTheRail()
}

func main() {
	// Движение по дороге DriveOnTheRoad на автомобильных колёсах carWheels водителем driver.
	driver := &driver{}
	carWheels := &carWheels{}
	driver.DrivingByCar(carWheels)

	// Движение по железной дороге DriveOnTheRail на колёсной паре wheelset машинистом поезда trainDriver.
	trainDriver := &trainDriver{}
	trainWheels := &wheelset{}
	trainDriver.DrivingByTrain(trainWheels)

	// Движение по дороге на автомобиле-поезде с колёсной парой вместо колёс водителем driver.
	carTrainAdapter := &RailwayAdapter{
		carTrain: trainWheels,
	}
	driver.DrivingByCar(carTrainAdapter)
}
