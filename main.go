// @author: Nicolas Vasquez Murcia
package main
import ("fmt"
		"database/sql"
	_ "github.com/lib/pq")

var db sql.DB

func main() {
    fmt.Println("Recipes Truora")
    fmt.Println("Nicolas Vasquez Murcia")

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


}

func create(int id, string imgURL, string description, string name, string ingredients)
{
	// Insert two rows into the "accounts" table.
    if _, err := db.Exec(
        "INSERT INTO Cook.Recipes VALUES ("+id+"," name+","+imgURL+","+description+","+ingredients+")"; err != nil {
        log.Fatal(err)
    }

}

func update(string fieldToUpdate, string name, string update)
{
	if _, err := db.Exec(
        "UPDATE Cook.Recipes SET "+fieldToUpdate+"="+update+"where name="+name; err != nil {
        log.Fatal(err)
    }
}

func read()
{
    rows, err := db.Query("SELECT * FROM Book.Recipes")
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

func delete(string name)
{
	if _, err := db.Exec(
        "Delete * from Cook.Recipes where name ="+name; err != nil {
        log.Fatal(err)
    }
}