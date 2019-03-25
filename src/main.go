package main

import (
    "github.com/gin-gonic/gin"
    "app/controllers"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("./views/*")
    router.Static("/assets", "./assets")

    ctrl := task.NewTask()
    router.GET("", ctrl.GetAll)

    router.POST("/", ctrl.Create)

    router.POST("/changedone", ctrl.ChangeDone)

    router.POST("/update", ctrl.Update)

    router.POST("/delete", ctrl.Delete)

    /*
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
    */
    router.Run()
}