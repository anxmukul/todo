package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anxmukul/todo/controller"
)

func getStringInput() string {
	titleScanner := bufio.NewScanner(os.Stdin)
	titleScanner.Scan()
	err := titleScanner.Err()
	if err != nil {
		panic(err)
	}
	return titleScanner.Text()

}
func main() {
	var contr controller.Controller
	contr = controller.NewTodoController()
	var choice int
	fmt.Scanf("%d", &choice)
	if choice == 1 {
		fmt.Println("Enter title:")
		title := getStringInput();
		fmt.Println("Enter content:")
		content := getStringInput()
		_, err := contr.Create(title, content);
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("New Todo :", newTodo)
	} else if choice == 2 {
		fmt.Println("Enter id:")
		var id int64
		fmt.Scanf("%d", &id)
		_, err := contr.SearchById(id)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("My Todo is:", myTodo)
	} else if choice == 3 {
		fmt.Println("Enter title:")
		title := getStringInput()
		_, err := contr.SearchByTitle(title)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("My Todos are :", myTodo)
	} else if choice == 4 {
		fmt.Println("Enter title:")
		title := getStringInput()
		_, err := contr.DeleteByTitle(title)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("Deleted Todo detail :", myTodo)
	} else {
		fmt.Println(choice)
	}
}
