package task

import(
	task "app/models"
)

type Task struct {
}

func NewTask() Task {
	return Task{}
}

func (c Task) GetAll() interface{} {
    repo := task.NewTaskRepository()
    tasks := repo.GetAll()

    return tasks
}

func (c Task) Create(title string){
	repo := task.NewTaskRepository()
	repo.Create(title)//
}

func (c Task) ChangeDone(id int, done bool){
	repo := task.NewTaskRepository()
	repo.ChangeDone(id, done)//
}

func (c Task) Update(id int, title string){
	repo := task.NewTaskRepository()
	repo.Update(id, title)
}

func (c Task) Delete(id int){
	repo := task.NewTaskRepository()
	repo.Delete(id)
}