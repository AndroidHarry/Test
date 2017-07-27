package main

/*
方法
Go 没有类。然而，仍然可以在结构体类型上定义方法。

方法接收者 出现在 func 关键字和方法名之间的参数中。

====================================================================
方法即函数
记住：方法只是个带接收者参数的函数。

====================================================================
方法（续）
你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。

但是，不能对来自其他包的类型或基础类型定义方法。

====================================================================
方法与指针重定向
比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：

var v Vertex
ScaleFunc(v)  // 编译错误！
ScaleFunc(&v) // OK
而以指针为接收者的方法被调用时，接收者既能为值又能为指针：

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
对于语句 v.Scale(5) ，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5) 。

====================================================================
方法与指针重定向（续）
同样的事情也发生在相反的方向。

接受一个值作为参数的函数必须接受一个指定类型的值：

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！
而以值为接收者的方法被调用时，接收者既能为值又能为指针：

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs() 。
*/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main1() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs1())
}

/********************************************************************************
方法（续）
你也可以为非结构体类型声明方法。

在此例中，我们看到了一个带 Abs 方法的数值类型 MyFloat 。

你只能为在同一包内定义的类型的接收者声明方法， 而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。

（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）
********************************************************************************/
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

/********************************************************************************
接收者为指针的方法
方法可以与命名类型或命名类型的指针关联。

刚刚看到的两个 Abs 方法。一个是在 *Vertex 指针类型上，而另一个在 MyFloat 值类型上。 有两个原因需要使用指针接收者。首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。其次，方法可以修改接收者指向的值。

尝试修改 Abs 的定义，同时 Scale 方法使用 Vertex 代替 *Vertex 作为接收者。

当 v 是 Vertex 的时候 Scale 方法没有任何作用。Scale 修改 v。当 v 是一个值（非指针），方法看到的是 Vertex 的副本，并且无法修改原始值。

Abs 的工作方式是一样的。只不过，仅仅读取 v。所以读取的是原始值（通过指针）还是那个值的副本并没有关系。
********************************************************************************/
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

////func (v *Vertex) Abs() float64 {
////	return math.Sqrt(v.X*v.X + v.Y*v.Y)
////}

//func main3() {
//	v := &Vertex{3, 4}
//	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
//	v.Scale(5)
//	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
//}

/********************************************************************************
方法与指针重定向
比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：

var v Vertex
ScaleFunc(v)  // 编译错误！
ScaleFunc(&v) // OK
而以指针为接收者的方法被调用时，接收者既能为值又能为指针：

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
对于语句 v.Scale(5) ，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。
也就是说，由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5) 。
********************************************************************************/
//func (v *Vertex) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main4() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

/********************************************************************************
方法与指针重定向（续）
同样的事情也发生在相反的方向。

接受一个值作为参数的函数必须接受一个指定类型的值：

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // 编译错误！
而以值为接收者的方法被调用时，接收者既能为值又能为指针：

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs() 。
********************************************************************************/
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

//func AbsFunc(v Vertex) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
//	fmt.Println(AbsFunc(v))

//	p := &Vertex{4, 3}
//	fmt.Println(p.Abs())
//	fmt.Println(AbsFunc(*p))
//}

/********************************************************************************
选择值或指针作为接收者
使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

在本例中， Scale 和 Abs 接收者的类型为 *Vertex ，即便 Abs 并不需要修改其接收者。

通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。 （我们会在接下来几页中明白为什么。）
********************************************************************************/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
