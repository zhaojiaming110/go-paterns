---
title: "建造者模式-Builder"
date: 2019-10-14T20:01:23+08:00
tags: ["创作模式"]
categories: ["设计模式-Go"]
author: "Beyourself"
---

无论是在现实世界中还是在软件系统中，都存在一些复杂的对象，它们拥有多个组成部分，如汽车，它包括车轮、方向盘、发送机等各种部件。而对于大多数用户而言，无须知道这些部件的装配细节，也几乎不会使用单独某个部件，而是使用一辆完整的汽车，可以通过建造者模式对其进行设计与描述，**建造者模式可以将部件和其组装过程分开，一步一步创建一个复杂的对象。用户只需要指定复杂对象的类型就可以得到该对象，而无须知道其内部的具体构造细节.**

![](https://tva1.sinaimg.cn/large/006y8mN6ly1g7y0k0kfm2j30c008k3zb.jpg)

# 建造者模式介绍

## 概述

**`建造者模式（Builder Pattern）`**又名**`生成器模式`**，是一种对象构建模式。它可以将复杂对象的建造过程抽象出来（抽象类别），使这个抽象过程的不同实现方法可以构造出不同表现（属性）的对象。

**`建造者模式`** 是一步一步创建一个复杂的对象，它允许用户只通过指定复杂对象的类型和内容就可以构建它们，用户不需要知道内部的具体构建细节。

直白一点的说，就是将我们在开发过程中遇到的大型对象，拆分成多个小对象，然后将多个小对象组装成大对象，并且对外部隐藏建造过程。

![建造过程](https://tva1.sinaimg.cn/large/006y8mN6ly1g7y0o6wl0uj30lx0bsmzp.jpg)

## 为什么要用建造者模式（**`优点`**）？

- 客户端不必知道产品内部组成的细节，将产品本身与产品的创建过程解耦，使得相同的创建过程可以创建不同的产品对象。
- 每一个具体建造者都相对独立，而与其他的具体建造者无关，因此可以很方便地替换具体建造者或增加新的具体建造者， 用户使用不同的具体建造者即可得到不同的产品对象 。
- 可以更加精细地控制产品的创建过程。将复杂产品的创建步骤分解在不同的方法中，使得创建过程更加清晰，也更方便使用程序来控制创建过程。
- 增加新的具体建造者无须修改原有类库的代码，指挥者类针对抽象建造者类编程，系统扩展方便，符合 **“开闭原则”**

## 哪些情况不要用建造者模式（**`缺点`**）？

- **产品之间差异性很大的情况：** 建造者模式所创建的产品一般具有较多的共同点，其组成部分相似，如果产品之间的差异性很大，则不适合使用建造者模式，因此其使用范围受到一定的限制。
- **产品内部变化很复杂的情况：** 如果产品的内部变化复杂，可能会导致需要定义很多具体建造者类来实现这种变化，导致系统变得很庞大。

## 抽象工厂模式vs建造者模式

抽象工厂模式实现对产品家族的创建，一个产品家族是这样的一系列产品：具有不同分类维度的产品组合，采用抽象工厂模式不需要关心构建过程，只关心什么产品由什么工厂生产即可。而建造者模式则是要求按照指定的蓝图建造产品，它的主要目的是通过组装零配件而产生一个新产品。

## 模式结构

建造者模式主要由以下四部分构成：

- **`Product（产品角色`**）：一个具体的产品对象。
- **`Builder（抽象建造者）`**：创建一个Product对象的各个部件指定的抽象接口。
- **`ConCreateBuilder（具体建造者`**）：实现抽象接口，构建和装配各个配件。
- **`Director（指挥者）`**：构建一个使用Builder接口的对象。它主要是用于创建一个复杂的对象。它主要有两个作用，一是：隔离了客户与对象的生产过程，二是：负责控制产品对象的生产过程。

## 类图&&时序图

![](https://tva1.sinaimg.cn/large/006y8mN6ly1g7y13q9x9sj30fa0790sz.jpg)

![](https://tva1.sinaimg.cn/large/006y8mN6ly1g7y14dntsbj30em0bdglu.jpg)

# 建造者模式分析

- Product（一个典型的复杂对象），其结构代码表示如下：

  ```go
  package car
  
  import "fmt"
  
  type Car struct {
  	Speed  float64
  	Color  string
  	Wheels string
  }
  
  func (c *Car) Show() {
  	fmt.Println("this car's Color:"+c.Color+" wheels:"+c.Wheels+" speed:", c.Speed)
  }
  ```

- Builder（抽象建造者），其中定义了产品的创建方法和返回方法，其典型代码如下：

  ```go
  package car
  
  type Builder interface {
  	BuildSpeed(speed float64) Builder
  	BuildColor(color string) Builder
  	BuildWheels(wheels string) Builder
  	GetResult() interface{}
  }
  ```

- ConCreateBuilder（具体建造者），实现抽象接口构建和装配各个部件实现代码如下：

  ```go
  package car
  
  type CarBuilder struct {
  	Car *Car
  }
  
  func (cb *CarBuilder) BuildSpeed(speed float64) Builder {
  
  	if cb.Car == nil {
  		cb.Car = new(Car)
  	}
  	cb.Car.Speed = speed
  
  	return cb
  }
  
  func (cb *CarBuilder) BuildColor(color string) Builder {
  
  	if cb.Car == nil {
  		cb.Car = new(Car)
  	}
  
  	cb.Car.Color = color
  
  	return cb
  }
  
  func (cb *CarBuilder) BuildWheels(wheels string) Builder {
  	if cb.Car == nil {
  		cb.Car = new(Car)
  	}
  
  	cb.Car.Wheels = wheels
  
  	return cb
  }
  
  func (cb *CarBuilder) GetResult() interface{} {
  	return cb.Car
  }
  ```

- Director（指挥者），其代码实现如下：

  建造者模式的结构中还引入了一个指挥者类Director，该类的作用主要有两个：一方面它隔离了客户与生产过程；另一方面它负责控制产品的生成过程。指挥者针对抽象建造者编程，客户端只需要知道具体建造者的类型，即**可通过指挥者类调用建造者的相关方法，返回一个完整的产品对象。**

  ```go
  package car
  
  type Director struct {
  	builder Builder
  }
  
  func (d *Director) SetBuilder(builder Builder) {
  	d.builder = builder
  }
  
  func (d *Director) Create(color, wheels string, speed float64) *Car {
  	d.builder.BuildColor(color)
  	d.builder.BuildSpeed(speed)
  	d.builder.BuildWheels(wheels)
  	return d.builder.GetResult().(*Car)
  
  }
  ```

- 客户端实现代码如下：

  在客户端代码中，无须关心产品对象的具体组装过程，只需确定具体建造者的类型即可，**建造者模式将复杂对象的构建与对象的表现分离开来，这样使得同样的构建过程可以创建出不同的表现**。

  ```go
  package main
  
  import (
  	"go-paterns/creational/builder/car"
  )
  
  func main() {
  	builder := new(car.CarBuilder)
  	director := new(car.Director)
  	director.SetBuilder(builder)
  	car := director.Create("red", "sport", 180.64)
  	car.Show()
  }
  ```

- 全部代码请见[Github](https://github.com/zhaojiaming110/go-paterns/tree/master/creational/builder/car).

