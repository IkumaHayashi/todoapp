package main

import (
    "fmt"
    "time"
    "strings"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "app/controllers"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("./views/*")
    router.Static("/assets", "./assets")

    router.GET("/", func(c *gin.Context) {
        
        controller := task.NewTask()
        tasks := controller.GetAll()

        c.HTML(http.StatusOK, "index.tmpl", gin.H{
                    "title": "TODO List",
                    "tasks": tasks,
                })
    })

    router.POST("/", func(c *gin.Context){
        title := c.PostForm("title")
        deadline := c.PostForm("deadline")
        datestring := strings.Split(deadline, "-")
        year, _ :=  strconv.Atoi(datestring[0])
        month, _ := strconv.Atoi(datestring[1])
        day, _ := strconv.Atoi(datestring[2])
        deadlineDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

        ctrl := task.NewTask()
        ctrl.Create(title, deadlineDate)

        c.Redirect(http.StatusMovedPermanently, "/")
    })

    router.POST("/changedone", func(c *gin.Context){
        id, _ := strconv.Atoi(c.PostForm("id"))
        var done bool
        if c.PostForm("done") == "true" {
            done = true
        } else {
            done = false
        }
        ctrl := task.NewTask()
        ctrl.ChangeDone(id, done)

        c.Redirect(http.StatusMovedPermanently, "/")
    })

    router.POST("/update", func(c *gin.Context){
        id, _ := strconv.Atoi(c.PostForm("id"))
        title := c.PostForm("title")
        
        ctrl := task.NewTask()
        ctrl.Update(id, title)

        c.Redirect(http.StatusMovedPermanently, "/")
    })

    router.POST("/delete", func(c *gin.Context){
        id, _ := strconv.Atoi(c.PostForm("id"))
        fmt.Println("delete target id = " + c.PostForm("id"));
        ctrl := task.NewTask()
        ctrl.Delete(id)

        c.Redirect(http.StatusMovedPermanently, "/")
    })
    router.Run()
}