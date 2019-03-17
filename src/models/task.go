package task

import (
	"fmt"
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
  
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
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
}
type Tasks []Task

type TaskRepository struct {
}

func NewTaskRepository() TaskRepository {
    return TaskRepository{}
}

// データベースに一行登録する
func (m TaskRepository) Create(title string) {
	task := Task{Title: title}
    db.Create(&task)
    db.Save(&task)
}

func (m TaskRepository) ChangeDone(id int, done bool){
	task := Task{}
	db.Find(&task, id)
	task.Done = done
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

func (m TaskRepository) GetAll() Tasks {
    var tasks = Tasks{}
    db.Find(&tasks)

    return tasks
}