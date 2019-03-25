package task

import(
	task "app/models"
	"fmt"
	"strings"
	"strconv"
	"time"
    "net/http"
    "github.com/gin-gonic/gin"
)

type Task struct {
}

func NewTask() Task {
	return Task{}
}

func (t Task) GetAll(c *gin.Context) {
    repo := task.NewTaskRepository()
    tasks := repo.GetAll()

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "TODO List",
		"tasks": tasks,
	})
}

func (t Task) Create(c *gin.Context) {

	title := c.PostForm("title")
	deadline := c.PostForm("deadline")
	datestring := strings.Split(deadline, "-")
	year, _ :=  strconv.Atoi(datestring[0])
	month, _ := strconv.Atoi(datestring[1])
	day, _ := strconv.Atoi(datestring[2])
	deadlineDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	repo := task.NewTaskRepository()
	repo.Create(title, deadlineDate)//

	c.Redirect(http.StatusMovedPermanently, "/")

}

func (t Task) ChangeDone(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Println("changedone: " + c.PostForm("id"))
	var done bool
	if c.PostForm("done") == "true" {
		done = true
	} else {
		done = false
	}
	repo := task.NewTaskRepository()
	repo.ChangeDone(id, done)//
	c.Redirect(http.StatusMovedPermanently, "/")

}

func (t Task) Update(c *gin.Context) {


	id, _ := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	

	repo := task.NewTaskRepository()
	repo.Update(id, title)
	
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (t Task) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Println("delete target id = " + c.PostForm("id"));


	repo := task.NewTaskRepository()
	repo.Delete(id)

	c.Redirect(http.StatusMovedPermanently, "/")

}