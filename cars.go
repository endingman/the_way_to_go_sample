//cars.go
package main

import (
	"fmt"
)

//空接口或者最小接口 不包含任何方法，它对实现不做任何要求：
type Any interface{}

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
	// ...
}

// 应用中定义了一个结构体，那么你也可能需要这个结构体的（指针）对象集合
type Cars []*Car //全局可用

func main() {

	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	// Cars []*Car 指针集合
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})

	// query，类似laravel集合fliter
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})

	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	//可以用它把汽车分类为独立的集合，像这样
	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}

	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)

	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}

// (空接口实现)定义一个通用的 Process() 函数，它接收一个作用于每一辆 car 的 f 函数作参数
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// (空接口实现)在上面的基础上，实现一个查找函数来获取子集合，并在 Process() 中传入一个闭包执行（这样就可以访问局部切片 cars）
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)

	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// (空接口实现)实现 Map 功能，产出除 car 对象以外的东西：
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

//我们也可以根据入参返回不同的函数。也许我们想根据不同的厂商添加汽车到不同的集合，但是这可能会是多变的。所以我们可以定义一个函数来产生特定的添加函数和 map 集：
func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) { //返回值==>(func(car *Car)\map[string]Cars
	// Prepare maps of sorted cars.
	sortedCars := make(map[string]Cars)

	/**
	 * sortedCars初始化
	 * @type {[type]}
	 */
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	// Prepare appender function（Car结构体指针集合）
	appender := func(c *Car) {
		//append car
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}

	return appender, sortedCars
}
