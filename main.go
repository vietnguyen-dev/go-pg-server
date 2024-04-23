package main
import (
	"net/http"
	"database/sql"
	"fmt"	
_	"github.com/lib/pq"
)

type vw_promotions struct {
	id int;
	pro_code string;
	name string;
	email string;
	how_many_ranked string;
	website_link string;
}

func pgConnect() *sql.DB { 
	connStr := "host=localhost port=5432 user=postgres password=Transformers22! dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func promotions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			db := pgConnect()
			defer db.Close()
			
			id := r.PathValue("id")
			rows, err := db.Query("SELECT * FROM vw_promotions WHERE id = $1;", id)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(rows)
		
			var promo vw_promotions
			for rows.Next() {
				err := rows.Scan(&promo.id, &promo.pro_code, &promo.name, &promo.email, &promo.how_many_ranked, &promo.website_link)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				// Log out the data
				fmt.Printf("Promotion ID: %d, Code: %s, Name: %s, Email: %s, How Many Ranked: %s, Website Link: %s\n",
					promo.id, promo.pro_code, promo.name, promo.email, promo.how_many_ranked, promo.website_link)
			}
		default:
  			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
 	}
}


func main() {
	http.HandleFunc("/{id}", promotions)
	http.ListenAndServe(":8080", nil);
}
