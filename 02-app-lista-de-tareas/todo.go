package main

import "fmt"

type taskList struct {
	tasks []*task // Almacenar el apuntador a la estructura task
}

func (t *taskList) addTask(task *task) {
	t.tasks = append(t.tasks, task)
}

func (t *taskList) removeTask(index int) {
	// Eliminar el ultimo elemento
	// t.tasks = t.tasks[:len(t.tasks)-1]

	// Eliminar el elemento en la posicion index
	// t.tasks[:index] -> Agregamos el slice de la posicion 0 hasta la posicion index (no incluye el index)
	// t.tasks[index+1:] -> Agregamos el slice de la posicion index+1 hasta el final
	// ... operador ellipsis, desempaqueta el slice del segundo parametro y lo transforma en argumentos append(t.tasks, task1, task2, task3)
	// De esa forma eliminamos el elemento en la posicion index, nos saltamos esa posicion
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printTasks() {
	for i, task := range t.tasks {
		fmt.Printf("%d. %s\n", i, task.name)
	}
}

func (t *taskList) printTasksCompleted() {
	for i, task := range t.tasks {
		if task.completed {
			fmt.Printf("%d. %s\n", i, task.name)
		}
	}
}

type task struct {
	name      string
	decs      string
	completed bool
}

func (t *task) markComplete() {
	t.completed = true
}

func (t *task) newDecs(decs string) {
	t.decs = decs
}

func (t *task) newName(name string) {
	t.name = name
}

func main() {
	fmt.Println("======= 1 Ejercicio =======")
	t := task{
		name: "Comprar pan",
		decs: "Comprar pan para el dia de hoy",
	}

	fmt.Printf("%+v\n", t)

	t.markComplete()
	t.newDecs("Comprar pan para el dia de mañana")
	t.newName("Comprar pan mañana")

	fmt.Printf("%+v\n", t)

	fmt.Println("======= 2 Ejercicio =======")
	t1 := task{
		name: "Comprar pan Blanco",
		decs: "Comprar pan para el dia de hoy",
	}

	t2 := task{
		name: "Comprar pan Integral",
		decs: "Comprar pan para el dia de hoy",
	}

	t3 := task{
		name: "Comprar pan de hamburguesa",
		decs: "Comprar pan para el dia de hoy",
	}

	lista := taskList{
		tasks: []*task{&t1, &t2},
	}

	fmt.Println("Lista de tareas, elemento 1", lista.tasks[0])

	lista.addTask(&t3)
	fmt.Println("Longitud de la lista", len(lista.tasks))
	fmt.Println("Lista de tareas, elemento 3", lista.tasks[2])

	fmt.Println("Remover el elemento 2 de la lista")
	lista.removeTask(1)
	fmt.Println("Longitud de la lista", len(lista.tasks))

	fmt.Println("======= 3 Ejercicio Recorrer =======")
	for i, task := range lista.tasks {
		fmt.Printf("Elemento %d: %+v\n", i, *&task.name)
	}

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Printf("Elemento %d: %+v\n", i, *&lista.tasks[i].name)
	}

	fmt.Println("======= 3.1 Ejercicio printTasks =======")
	lista.printTasks()

	fmt.Println("======= 3.2 Ejercicio printTasks Completed =======")
	lista.tasks[0].markComplete()
	lista.printTasksCompleted()

	fmt.Println("======= 4 Ejercicio Break =======")

	for i := 0; i < 10; i++ {
		// Break -> Termina el ciclo
		if i == 5 {
			fmt.Println("Break", i)
			break
		}
		fmt.Println("Iteracion", i)
	}

	fmt.Println("======= 5 Ejercicio continue =======")

	for i := 0; i < 10; i++ {
		// Break -> Termina el ciclo
		if i == 5 {
			fmt.Println("continue", i)
			continue
		}
		fmt.Println("Iteracion", i)
	}

	fmt.Println("======= 6 Ejercicio Mapas Tareas =======")

	mapTasks := make(map[string]*taskList)
	mapTasks["Jesus"] = &lista

	t4 := task{
		name: "Comprar pan Blanco 2",
		decs: "Comprar pan para el dia de hoy",
	}

	t5 := task{
		name: "Comprar pan Integral 2",
		decs: "Comprar pan para el dia de hoy",
	}

	t6 := task{
		name: "Comprar pan de hamburguesa 2",
		decs: "Comprar pan para el dia de hoy",
	}

	lista2 := taskList{
		tasks: []*task{&t4, &t5, &t6},
	}

	mapTasks["Oriana"] = &lista2

	fmt.Println("__ Mapas tareas de Jesus")
	mapTasks["Jesus"].printTasks()

	fmt.Println("__ Mapas tareas de Oriana")
	mapTasks["Oriana"].printTasks()
}
