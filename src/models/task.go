package task

import (
	"fmt"
	"time"
	"strconv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB

func init(){
    var err error

	DBMS     := "mysql"
	USER     := "root"
	PASS     := "root"
	PROTOCOL := "tcp(todoapp_db_1:3306)"
	DBNAME   := "todoapp_db"
  
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=True"
	db,err = gorm.Open(DBMS, CONNECT)
	
	db.DropTableIfExists(&Task{})
	db.CreateTable(&Task{})
  
	if err != nil {
	  panic(err.Error())
	}

}
  
type Task struct {
	ID   int    `gorm:"primary_key"`
	Title string `gorm:"size:140"`
	Done bool
    DeadlineDate   time.Time `sql:"not null;type:datetime"`
}
type TaskRepository struct {
}

func NewTaskRepository() TaskRepository {
    return TaskRepository{}
}

// データベースに一行登録する
func (m TaskRepository) Create(title string, deadline time.Time) {
	task := Task{Title: title, DeadlineDate: deadline}
    db.Create(&task)
    db.Save(&task)
}

func (m TaskRepository) ChangeDone(id int, done bool){
	task := Task{}
	db.Find(&task, id)
	fmt.Println("task changing! done: " + strconv.FormatBool(task.Done))
	fmt.Println(task)
	task.Done = done
	fmt.Println("task changed! done: " + strconv.FormatBool(done))
	fmt.Println(task)
    db.Save(&task)
}

func (m TaskRepository) Update(id int, title string){
	task := Task{}
	db.Find(&task, id)
	task.Title = title
    db.Save(&task)
}


func (m TaskRepository) Delete(id int){
	task := Task{}
	db.Find(&task, id)
	fmt.Println(task);
    db.Delete(&task)
}

type DisplayTask struct {
		
	ID   int
	Title string
	Done bool
	DeadlineDate string
}

type Tasks []Task
type DisplayTasks []DisplayTask

func (m TaskRepository) GetAll() DisplayTasks {
	var tasks = Tasks{}
    db.Find(&tasks)

	var dts = DisplayTasks{}
	for i:= 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
		dt := DisplayTask{}
		
		dt.ID = tasks[i].ID
		dt.Title = tasks[i].Title
		dt.Done = tasks[i].Done
		dt.DeadlineDate =  tasks[i].DeadlineDate.Format("2006-01-02")
		dts = append(dts, dt)
		
	}



    return dts
}