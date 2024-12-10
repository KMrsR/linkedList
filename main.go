package main

func main() {
	// Создаем пустой связный список
	myList := LinkedList{}

	// Сначала вставляем данные=
	myList.insertAtBack(1)
	myList.insertAtBack(10)
	myList.insertAtBack(10)
	myList.insertAtBack(20)
	myList.insertAtBack(20)
	myList.insertAtBack(20)
	myList.insertAtBack(30)
	myList.insertAtBack(30)
	myList.insertAtBack(30)

	/*
		// Выводим содержимое связного списка
		fmt.Println("Linked List Contents:")
		myList.print()

		// Подсчитываем узлы в связном списке
		count := myList.countNodes()
		fmt.Printf("Total number of nodes: %d\n", count)

		// Находим и выводим узел по заданному индексу
		indexToFind := 2
		foundNode := myList.findNodeAt(indexToFind)
		if foundNode != nil {
			fmt.Printf("Node at index %d has data: %d\n", indexToFind, foundNode.data)
		} else {
			fmt.Printf("Node at index %d not found\n", indexToFind)
		}

		// При необходимости выполняем другие операции...

		// Удаляем последний узел
		myList.deleteLast()

		// Выводим обновленный связный список
		fmt.Println("Linked List After Deletion:")
	*/
	myList.print()
	myList.deleteDuplicate()
	myList.print()
}
