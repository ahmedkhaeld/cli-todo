package todo_test

import (
	todo "github.com/ahmedkhaeld/cli-todo"
	"io/ioutil"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	task := "New Task"
	l.Add(task)

	if l[0].Task != task {
		t.Errorf("Expected %q got %q instead", task, l[0].Task)
	}

}

func TestComplete(t *testing.T) {
	l := todo.List{}
	task := "New Task"
	l.Add(task)

	if l[0].Task != task {
		t.Errorf("Expected %q got %q instead", task, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("new task should not be completed; Done should be false.")
	}
	//mark done to true by calling Complete
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("new task should be completed; Done should be true")
	}

}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	// add tasks to the l List
	for _, v := range tasks {
		l.Add(v)
	}
	// now l and tasks are the same

	if l[0].Task != tasks[0] {
		t.Errorf("expected %q, got %q instead", tasks[0], l[0].Task)
	}
	//before calling delete, length of l is 3
	l.Delete(2) // remove from l task2
	// now # of  tasks in l are 2

	if len(l) != 2 {
		t.Errorf("expected list length %d, got %d instead", 2, len(l))
	}

	//l []{"Task 1", "Task 3"}
	// tasks  []{"Task 1", "Task 2", "Task 3"}
	// so l[1] eq to task[2]
	if l[1].Task != tasks[2] {
		t.Errorf("expected %q, got %q instead", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	//In this test case, you’re creating two variables,
	//l1 and l2, both of type todo.List.
	//You’re adding a task to l1 and saving it.
	//Then you’re loading it into l2 and comparing both values.
	//The test fails if the values don’t match,
	//in which case you provide an error message
	//showing the values you got.
	l1 := todo.List{}
	l2 := todo.List{}

	task := "New Task"
	l1.Add(task)
	// l1 first task should be eq to task
	if l1[0].Task != task {
		t.Errorf("expected %q, got %q instead", task, l1[0].Task)
	}
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name()) //ensure the temp file is deleted at the end of the test

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("error saving list  to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	//after saving task to tf, and getting task from tf l1 & l2 should be equal
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l1[0].Task, l2[0].Task)
	}
}
