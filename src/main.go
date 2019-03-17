package main

import (
    //"fmt"
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
        ctrl := task.NewTask()
        ctrl.Create(title)

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
    router.Run()
}