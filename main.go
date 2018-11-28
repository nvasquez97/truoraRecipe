// @author: Nicolas Vasquez Murcia
package main
import ("fmt"
    "encoding/json"
	"database/sql"
	"net/http"
	"github.com/lib/pq"
    "github.com/gorilla/mux"
	)

var db sql.DB

type Recipe struct{
	id 			int 	'json:"id,omitempty"'
	name 		string 	'json:"name,omitempty"'
	imgURL 		string 	'json:"imgURL,omitempty"'
	description string 	'json:"description,omitempty"'
	ingredients string 	'json:"ingredients,omitempty"'
}

var recipes []Recipe

func main() {
    fmt.Println("Recipes Truora")
    fmt.Println("Nicolas Vasquez Murcia")

    router := mux.NewRouter()
    router.HandleFunc("/recipes/", readAll).Methods("GET")
	router.HandleFunc("/recipe/{id}", read).Methods("GET")
	router.HandleFunc("/recipe/{name}", search).Methods("GET")
	router.HandleFunc("/recipe/", update).Methods("POST")
	router.HandleFunc("/recipe/", create).Methods("PUT")
    router.HandleFunc("/recipe/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))

    //Conectarse a la DB
    db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/Cook?sslmode=disable")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

    // Crear la tabla de recetas
    if _, err := db.Exec(
        "create table if not exists Cook.Recipes(id INT PRIMARY KEY, name STRING, imgURL STRING, description STRING, ingredients STRING)"); err != nil {
        log.Fatal(err)
    }

}

func create(writer http.ResponseWriter, request *http.Request)
{
	int id, string imgURL, string description, string name, string ingredients
	// Insert two rows into the "accounts" table.
    if _, err := db.Exec(
        "INSERT INTO Cook.Recipes VALUES ("+id+"," name+","+imgURL+","+description+","+ingredients+")"; err != nil {
        log.Fatal(err)
    })

}

func update(writer http.ResponseWriter, request *http.Request)
{
	if _, err := db.Exec(
        "UPDATE Cook.Recipes SET "+fieldToUpdate+"="+update+"where name="+name; err != nil {
        log.Fatal(err)
    })
}

func readAll(writer http.ResponseWriter, request *http.Request)
{
    rows, err := db.Query("SELECT * FROM Book.Recipes")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int, imgURL,name, description,ingredients string
        if err := rows.Scan(&id, &imgURL, &name, &description, &ingredients); err != nil {
            log.Fatal(err)
        }
        var rec = Recipe{id: id, name:name, imgURL: imgURL, description:description, ingredients:ingredients}
        recipes.append(recipes, rec)
    }
    json.NewEncoder(writer).Encode(recipes)
}

func read(writer http.ResponseWriter, request *http.Request)
{
	param:=
    rows, err := db.Query("SELECT * FROM Book.Recipes")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    fmt.Println("Initial balances:")
    for rows.Next() {
        var id, imgURL,name, description,ingredients int
        if err := rows.Scan(&id, &imgURL, &name, &description, &ingredients); err != nil {
            log.Fatal(err)
        }
    }
}

func search(writer http.ResponseWriter, request *http.Request)
{
    rows, err := db.Query("SELECT * FROM Book.Recipes WHERE name LIKE "+"U+0025"+NAME+"U+0025")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    fmt.Println("Initial balances:")
    for rows.Next() {
        var id, balance int
        if err := rows.Scan(&id, &balance); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%d %d\n", id, balance)
    }
}

func delete(writer http.ResponseWriter, request *http.Request)
{
	if _, err := db.Exec(
        "Delete * from Cook.Recipes where name ="+name; err != nil {
        log.Fatal(err)
    }
}