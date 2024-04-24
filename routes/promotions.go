package routes

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"	
	"go-http-server/utils"
)

type vw_promotions struct {
	Id int 		`json:"id"`;
	ProCode string `json:"pro_code"`;
	Name string `json:"name"`;
	Email string `json:"email"`;
	HowManyRanked int `json:"how_many_ranked"`;
	WebsiteLink string `json:"website_link"`;
}

// Helper function to convert sql.NullString to a plain string
func NullStringToString(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }
    return "" // Return an empty string if the value is invalid (NULL)
}

func Promotions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// connect to the database
			db := utils.PgConnect()
			defer db.Close()
		
			//query the promotions view
			id := r.PathValue("id")
			rows, err := db.Query("SELECT * FROM vw_promotions WHERE id = $1 LIMIT 1;", id)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()

			//turn the data into proper json
			var promotions []vw_promotions
			for rows.Next() {
				var promo vw_promotions
				var websiteLink sql.NullString
				if err := rows.Scan(
					&promo.Id,
					&promo.ProCode,
					&promo.Name,
					&promo.Email,
					&promo.HowManyRanked,
					&websiteLink,
				); err != nil {
					http.Error(w, "Data extraction error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				promo.WebsiteLink = NullStringToString(websiteLink)
				promotions = append(promotions, promo)
			}

			//return data
			jsonData, err := json.Marshal(promotions)
			if err != nil {
				http.Error(w, "JSON serialization failed", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}

			// Set the content type and write the JSON data
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		default:
  			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
 	}
}
