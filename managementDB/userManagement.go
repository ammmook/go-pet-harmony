package managementdb

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	if r.Method == http.MethodGet {
		db := initializers.OpenConnection()
		result, _ := db.Query("select * from users")
		for result.Next() {
			var user model.User
			result.Scan(&user.Email, &user.Firstname, &user.Lastname, &user.Password, &user.PhoneNumber, &user.Role)
			users = append(users, user)
		}
		res, _ := json.Marshal(users)
		fmt.Fprint(w, (string(res)))

		defer db.Close()
	}
}

func PostUsers(w http.ResponseWriter, r *http.Request) {

}

func PutUsers(w http.ResponseWriter, r *http.Request) {

}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {

}
