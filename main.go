// punto de estrada del programa go
package main

//Esta sección importa el paquete fmt,
//que proporciona funciones de formato de entrada y salida, como Println y Printf.
import (
	"fmt"
)

/*
Definición de la estructura User
objetos en formato go
ID: de tipo string
Name: de tipo string
Email: de tipo string
*/
type User struct {
	ID    string
	Name  string
	Email string
}

/*
variable global users que es una lista (slice) de User.
Esta lista se usará para almacenar todos los usuarios creados.
*/
// Función para crear un usuario
/*Llamamos a createUser("1", "John Doe", "john@example.com").
Se crea un nuevo User con los valores proporcionados.Este User se agrega al slice users
Ejemplo:
users: [
{
ID: "1", Name: "John Doe", Email: "john@example.com"
}
]

*/
var users []User

func createUser(id, name, email string) {
	user := User{ID: id, Name: name, Email: email}
	//con la funcion append se agrega a la lista
	users = append(users, user)
	/*"User created: ...": Es una cadena literal que se imprimirá tal cual. En este caso, indica que se ha creado un usuario.
	%+v: Es un verbo de formato específico para Go que se utiliza para imprimir valores de estructuras.
	Este verbo imprimirá los nombres de los campos de la estructura junto con sus valores.
	\n: Es un carácter de nueva línea (newline) que mueve el cursor a la siguiente línea después de imprimir la cadena. Esto asegura que cualquier texto posterior se imprima en una nueva línea.*/
	fmt.Printf("User created: %+v\n", user)
}

// Función para obtener todos los usuarios
func getUsers() {
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

// Función para obtener un usuario por ID
func getUserByID(id string) (*User, bool) {
	for _, user := range users {
		if user.ID == id {
			return &user, true
		}
	}
	return nil, false
}

// Función para actualizar un usuario por ID

func updateUser(id, name, email string) bool {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = name
			users[i].Email = email
			fmt.Printf("User updated: %+v\n", users[i])
			return true
		}
	}
	return false
}

// Función para eliminar un usuario por ID
func deleteUser(id string) bool {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			fmt.Printf("User deleted: %+v\n", user)
			return true
		}
	}
	return false
}

func main() {

	// Inicialización
	fmt.Println("Inicializando aplicación...")

	//buscar o obtener todos los usuarios
	createUser("1", "John Doe", "john@example.com")
	createUser("2", "Jane Doe", "jane@example.com")

	fmt.Println("All Users:")
	getUsers()

	// Buscar un usuario por ID
	idToFind := "1"
	user, found := getUserByID(idToFind)
	if found {
		fmt.Printf("User found: %+v\n", *user)
	} else {
		fmt.Println("User not found")
	}

	// Actualizar un usuario por id

	idToUpdate := "1"
	updated := updateUser(idToUpdate, "John Smith", "johnsmith@example.com")
	if updated {
		fmt.Printf("User with ID %s updated successfully\n", idToUpdate)
	} else {
		fmt.Printf("User with ID %s not found\n", idToUpdate)
	}

	// Mostrar todos los usuarios después de la actualización
	fmt.Println("All Users after update:")
	getUsers()

	// Eliminar un usuario
	idToDelete := "2"
	deleted := deleteUser(idToDelete)
	if deleted {
		fmt.Printf("User with ID %s deleted successfully\n", idToDelete)
	} else {
		fmt.Printf("User with ID %s not found\n", idToDelete)
	}

	// Mostrar todos los usuarios después de la eliminación
	fmt.Println("All Users after deletion:")
	getUsers()

}
