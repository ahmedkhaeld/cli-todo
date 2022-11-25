# Todo CLI Program
Command line tool for managing a list of "to-do" items. it will keep track of items left in an activity
.<br> It will save the list of items in a file using `json` format 

* `todo` package provides an API to interact with todo-list functionalities (Add, Save, Get, Delete) items
* `main` package provides CLI to implement the API
   * `-list`: A boolean flag to list all to-do items
   * `-task`: a String flag to add new task to the list
   * `-complete`: An integer flag, to mark a task number as completed using task number
    
---
1.  add new task
<br>`$ go run . -task study`<br>
`$cat .todo.json` <br>
```json
[{"Task":"task","Done":false,"CreatedAt":"2022-11-25T04:19:46.621972588+02:00","CompletedAt":"2022-11-25T04:19:46.621972708+02:00"}]
```
2. mark task as completed
<br> `$ go run . -complete 1` done value modified to true<br>
```json
[{"Task":"task","Done":true,"CreatedAt":"2022-11-25T04:19:46.621972588+02:00","CompletedAt":"2022-11-25T04:47:40.058318093+02:00"}]
```
3. list current list<br>
`$ go run . -list`
```
X 1: study
  2: play

```
note: completed task has sign X 

---
`$ go build .` <br>
`$ ./todo -h` display usage<br>
``` 
./todo tool. Developed for The Pragmatic Bookshelf
Copyright 2022
Usage information:
  -complete int
        Item to be completed
  -list
        List all tasks
  -task string
        Task to be included in the ToDo list

```