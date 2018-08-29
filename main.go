package main

import (
  // "fmt"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

type Person struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Age int `json:"age"`
  Gender string `json:"gender"`
}

var db *gorm.DB

func dbInit() {
  var err error
  db, err = gorm.Open("mysql", "root:password@tcp(db)/sampledb")

  if err != nil {
    panic("failed to connect database\n")
  }

  db.AutoMigrate(&Person{})
}

func main() {
  r := gin.Default()

  dbInit()

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ping",
    })
  })

  r.GET("/peoples", func(c *gin.Context) {
    var peoples []Person

    db.Find(&peoples)

    c.JSON(200, peoples)
  })

  r.POST("/peoples", func(c *gin.Context) {
    var person Person
    c.BindJSON(&person)
    // fmt.Printf(person.Name)
    db.Save(&person)
    c.JSON(200, person)
  })

  r.Run(":8080")
}
