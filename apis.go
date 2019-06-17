package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
	"net/http"
)

func getUnit(c *gin.Context) {
	b:=All.Connection.Collection(UNIT)
	results:=b.Find(bson.M{})
	unit:=&Unit{}
	units:=[]*Unit{}
	for results.Next(unit) {
		units=append(units, unit)
		unit=&Unit{}
	}

	c.JSON(200, units)
	c.Abort()
	return
}

func getUnprocessed(c *gin.Context)  {
	b:=All.Connection.Collection(ORDER)
	results:=b.Find(bson.M{"processed":false})
	order:=&Order{}
	orders:=[]*Order{}

	for results.Next(order) {
		orders=append(orders, order)
		order=&Order{}
	}

	c.JSON(200, orders)
	c.Abort()
	return
}

func getItem(c *gin.Context)  {
	b:=All.Connection.Collection(ITEM)
	results:=b.Find(bson.M{})
	item:=&Item{}
	items:=[]*Item{}

	for results.Next(item) {
		items=append(items, item)
		item=&Item{}
	}

	c.JSON(200, items)
	c.Abort()
	return
}

func postOrder(c *gin.Context)  {
	has,tel:=hasUser(c)
	if !has{
		c.JSON(http.StatusUnauthorized,gin.H{"message":"密码错误"})
		c.Abort()
		return
	}

	b:=All.Connection.Collection(RESTAURANT)
	restaurant := &Restautant{}
	err:=b.FindOne(bson.M{"tel":tel},restaurant)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误1"})
		c.Abort()
		return
	}
	order:=&Order{}
	err=c.ShouldBindBodyWith(order,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误2"})
		c.Abort()
		return
	}

	b=All.Connection.Collection(ORDER)
	order.Processed=false
	err=b.Save(order)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误3"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"成功"})
	c.Abort()

	return
}

func register(c *gin.Context)  {
	user := &UserLogin{}
	err:=c.ShouldBindBodyWith(user,binding.JSON)
	b:=All.Connection.Collection(USER)
	realUser :=&UserLogin{}
	errr:=b.FindOne(bson.M{"tel":user.Tel},realUser)
	if errr==nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"已经注册过了"})
		c.Abort()
		return

	}
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误4"})
		c.Abort()
		return
	}
	restaurant:=&Restautant{}
	err=c.ShouldBindBodyWith(restaurant,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误5"})
		c.Abort()
		return
	}

	user.Password,err=HashPassword(user.Password)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误6"})
		c.Abort()
		return
	}

	err=b.Save(user)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误7"})
		c.Abort()
		return
	}

	b=All.Connection.Collection(RESTAURANT)
	err=b.Save(restaurant)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误8"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"成功"})
	c.Abort()
	return

}

//func getLast(c *gin.Context){
//
//	has,tel:=hasUser(c)
//	if !has {
//		c.JSON(http.StatusUnauthorized,gin.H{"message":"密码错误"})
//		c.Abort()
//		return
//	}
//
//
//
//}

func hasUser(c *gin.Context) (bool,string){
	user := &UserLogin{}
	err:=c.ShouldBindBodyWith(user,binding.JSON)
	if err!=nil {
		println(err)
	}
	b:=All.Connection.Collection(USER)
	realUser := &UserLogin{}
	err=b.FindOne(bson.M{"tel":user.Tel},realUser)
	if err!=nil {
		fmt.Println(err.Error())
		return false,""
	}
	same:=CheckPasswordHash(user.Password,realUser.Password)
	if !same {
		return false,""
	}
	return true,user.Tel
}

func addUnit(c *gin.Context)  {
	unit:= &Unit{}
	err:=c.ShouldBindBodyWith(unit,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误9"})
		c.Abort()
		return
	}
	b:=All.Connection.Collection(UNIT)
	err=b.Save(unit)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误10"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"成功"})
	c.Abort()
	return
}

func addItem(c *gin.Context)  {
	item:= &Item{}
	err:=c.ShouldBindBodyWith(item,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误11"})
		c.Abort()
		return
	}
	b:=All.Connection.Collection(ITEM)
	err=b.Save(item)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误12"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"成功"})
	c.Abort()
	return
}

func orderDone(c *gin.Context)  {
	id := &OrderID{}
	err:=c.ShouldBindBodyWith(id,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误13"})
		c.Abort()
		return
	}
	b:=All.Connection.Collection(ORDER)
	order := &Order{}
	if !bson.IsObjectIdHex(id.Id){
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误16"})
		c.Abort()
		return
	}
	err=b.FindById(bson.ObjectIdHex(id.Id),order)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		fmt.Println("document not found")
	}
	if err!=nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误14"})
		c.Abort()
		return
	}

	order.Processed=true
	err=b.Save(order)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误15"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"成功"})
	c.Abort()
	return

}

func removeUnit(c *gin.Context)  {
	id := &OrderID{}
	err:=c.ShouldBindBodyWith(id,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误16"})
		c.Abort()
		return
	}
	b:=All.Connection.Collection(UNIT)
	unit := &Unit{}
	if !bson.IsObjectIdHex(id.Id){
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误17"})
		c.Abort()
		return
	}
	err=b.FindById(bson.ObjectIdHex(id.Id),unit)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		fmt.Println("document not found")
	}
	if err!=nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误18"})
		c.Abort()
		return
	}

	err=b.DeleteDocument(unit)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误19"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"成功"})
	c.Abort()
	return
}


func removeItem(c *gin.Context)  {
	id := &OrderID{}
	err:=c.ShouldBindBodyWith(id,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误20"})
		c.Abort()
		return
	}
	b:=All.Connection.Collection(ITEM)
	item := &Item{}
	if !bson.IsObjectIdHex(id.Id){
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误21"})
		c.Abort()
		return
	}
	err=b.FindById(bson.ObjectIdHex(id.Id),item)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		fmt.Println("document not found")
	}
	if err!=nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误22"})
		c.Abort()
		return
	}

	err=b.DeleteDocument(item)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误23"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"成功"})
	c.Abort()
	return
}

func findByPhone (c *gin.Context)  {
	tel:=c.Param("tel")
	b:=All.Connection.Collection(ORDER)
	results:=b.Find(bson.M{"tel":tel})
	order:=&Order{}
	orders:=[]*Order{}
	for results.Next(order) {
		orders=append(orders, order)
		order=&Order{}
	}

	c.JSON(200, orders)
	c.Abort()
	return
}

func getClass(c *gin.Context)  {
	b:=All.Connection.Collection(CLASS)
	results:=b.Find(bson.M{})

	class:=&Class{}
	classes:=[]*Class{}
	for results.Next(class) {
		classes=append(classes, class)
		class=&Class{}
	}

	c.JSON(200, classes)
	c.Abort()
	return
}

func addClass(c *gin.Context)  {

	class:=&Class{}
	err:=c.ShouldBindBodyWith(class,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误24"})
		c.Abort()
		return
	}
	b:=All.Connection.Collection(CLASS)

	err=b.Save(class)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误25"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message":"成功"})
	c.Abort()
	return
}

func removeClass(c *gin.Context)  {

	id:=&OrderID{}
	err:=c.ShouldBindBodyWith(id,binding.JSON)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误26"})
		c.Abort()
		return
	}
	b:=All.Connection.Collection(CLASS)

	if !bson.IsObjectIdHex(id.Id){
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误27"})
		c.Abort()
		return
	}
	class :=&Class{}
	err=b.FindById(bson.ObjectIdHex(id.Id),class)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		fmt.Println("document not found")
	}
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误28"})
		c.Abort()
		return
	}
	err=b.DeleteDocument(class)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"错误29"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message":"成功"})
	c.Abort()
	return
}

