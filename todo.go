// Package todo is an API contains the business logic to work with to-do items
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//item represent a single to-do item
//not exported, so it can't be used by API users directly
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

//List represents a list of to-do items
type List []item

//String custom format when calling println on l
func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}

		//adjust the item number k to print number starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}

	return formatted
}

//Add creates a new to-do item and append it to the list
//as it only require an item.Task field to create new to-do
func (l *List) Add(task string) {
	i := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}
	*l = append(*l, i) //dereference the pointer to the List to access the underlying slice
}

//Complete marks an item as completed by setting Done to true
// and CompletedAt to now
func (l *List) Complete(i int) error {
	ls := *l
	//check if task # provided in range of the list
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exists", i)
	}
	// Adjusting index for 0 based index
	//because index starts from 0
	// e.g. task #3 at index 2
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

//Delete remove an item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	//e.g. task #4 at index 3, so we want to remove index3
	//ls[:i-1] from 0 to 3 this excludes index3
	// ls[4:] start from index 4 to end and append them
	// this removes index 3

	return nil
}

//Save encodes the List as json and saves it to a file using the provided file name
func (l *List) Save(fn string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fn, js, 06444)

}

//Get obtains a list of items from a saved json file
//opens the provided file name, decodes the json data
//and store it into a list
func (l *List) Get(fn string) error {
	file, err := ioutil.ReadFile(fn)
	if err != nil {
		// handles when file does not exist
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	//handles when file is empty
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
