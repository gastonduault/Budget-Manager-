package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mux"

	_ "github.com/go-sql-driver/mysql"
)

// STRUCTURE'S LIST
type Result struct {
	Column1 int    `json:"column1"`
	Column2 int    `json:"column2"`
	Column3 string `json:"column3"`
}

type User struct {
	UserID    int    `json:"userID"`
	Login     string `json:"login"`
	Pass_word string `json:"pass_word"`
}

type Transaction struct {
	ID              int     `json:"id"`
	Label           string  `json:"label"`
	Price           float64 `json:"price"`
	PrincipalCat    string  `json:"principal_cat"`
	SubCat          string  `json:"sub_cat"`
	DateTransaction string  `json:"date_transaction"`
}

type Solde struct {
	Total      float64            `json:"total"`
	ByCategory map[string]float64 `json:"by_category"`
}

// MAIN
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/test", handler)
	router.HandleFunc("/login", handler_login)
	http.HandleFunc("/transaction", handler_transation)
	http.HandleFunc("/solde", handler_solde)
	log.Fatal(http.ListenAndServe("localhost:80", router))

}

func handler_login(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/db_budget_manager")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Test connectionsf
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection successful.")

	// Gérer la demande POST pour créer un utilisateur
	/*if r.Method == "POST" {
		fmt.Println("POST")
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insérer l'utilisateur dans la base de données
		stmt, err := db.Prepare("INSERT INTO users (login, pass_word) VALUES (?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = stmt.Exec(u.Login, u.Pass_word)
		if err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				if mysqlErr.Number == 1062 {
					http.Error(w, "Username already exists", http.StatusConflict)
					return
				}

			}
		}
	}*/

	// Méthode GET
	if r.Method == "GET" {
		fmt.Println("GET")
		ctx := context.Background()

		// Récupération des arguments de la méthode GET
		keys := r.URL.Query()
		// On récupère le premier argument
		loginTab, ok := keys["login"]
		if !ok || len(loginTab[0]) < 1 {
			fmt.Fprintf(w, "Argument login manquant")
			return
		}
		login := loginTab[0]
		fmt.Fprintf(w, "Argument login : = %s", login)

		// On récupère le second argument
		pass_wordTab, ok := keys["pass_word"]
		if !ok || len(pass_wordTab[0]) < 1 {
			fmt.Fprintf(w, "Argument pass_word manquant")
			return
		}
		pass_word := pass_wordTab[0]
		fmt.Fprintf(w, "Argument pass_word : = %s", pass_word)

		// On prepare la requête (vérifie si elle est correct)
		stmt, err := db.Prepare("SELECT `connection`(?, ?) AS `connection`;")
		if err != nil {
			fmt.Println("requete NOK")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// On exécute la requête avec les valeur récupérées dans via la méthode GET
		row := stmt.QueryRowContext(ctx, login, pass_word)

		var result int64
		err = row.Scan(&result)

		if err != nil {
			fmt.Println("Erreur row")
			fmt.Printf("REPONSE REQUETE (ERREUR) : %d\n", result)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("REPONSE REQUETE : %d\n", result)

		fmt.Println("FIN")
	}
}

func handler_transation(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/db_budget_manager")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Test connectionsf
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection successful.")
}

func handler_solde(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/db_budget_manager")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Test connectionsf
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection successful.")
}
func handler(w http.ResponseWriter, r *http.Request) {
	// Créer une connexion à la base de données
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/db_budget_manager")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Test connectionsf
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection successful.")

	// Initialister le routeur Gorilla mux
	//router := mux.NewRouter()

	// Gestionnaire d'authentification
	//router.HandleFunc("/test",handler_login).Methods("POST","GET")

	// Exécuter une requête SELECT
	rows, err := db.Query("SELECT * FROM sub_category")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	// Parcourir les résultats de la requête
	var tableau []string = []string{}
	for rows.Next() {
		var column1 int
		var column2 int
		var column3 string
		err := rows.Scan(&column1, &column2, &column3)
		if err != nil {
			panic(err.Error())
		}
		requete := fmt.Sprintf("|%d, |%d, |%s", column1, column2, column3)
		//requete := printf("%d | %d| %s"| column1, column2, column3)
		tableau = append(tableau, requete)
	}
	fmt.Println(tableau)

	// Gérer les erreurs de parcours des résultats
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	// Créer une slice pour stocker les résultats de la requête
	results := []Result{}

	// Parcourir les résultats de la requête et ajouter chaque ligne à la slice
	for rows.Next() {
		var result Result
		err := rows.Scan(&result.Column1, &result.Column2, &result.Column3)
		if err != nil {
			panic(err.Error())
		}
		results = append(results, result)

	}

	// Transformer la slice en JSON
	jsonResults, err := json.Marshal(results)
	if err != nil {
		panic(err.Error())
	}

	// Envoyer le JSON en tant que réponse HTTP
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResults)
}

//-------------------------------------------------------------------------------------------------------
// Call stored procedures
/*_, err = db.Exec("CALL add_user(?, ?)", "testuser", "testpass")
if err != nil {
	log.Fatal(err)
}

_, err = db.Exec("CALL change_pass_word(?, ?)", "newpass", 1)
if err != nil {
	log.Fatal(err)
}

_, err = db.Exec("CALL delete_transaction(?)", 1)
if err != nil {
	log.Fatal(err)
}

_, err = db.Exec("CALL modify_transaction_amount(?, ?)", 1, 500)
if err != nil {
	log.Fatal(err)
}

// Call functions
var accountID int
err = db.QueryRow("SELECT get_accountID(?)", "testuser").Scan(&accountID)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Account ID for testuser: %d\n", accountID)

//Parce que la requete peut renvoyer un element nul
rows, err := db.Query("SELECT get_category_name_from_transaction_ID(?)", 1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
	var categoryName sql.NullString

	if err := rows.Scan(&categoryName); err != nil {
		log.Fatal(err)
		fmt.Println("Category name is NULL")
	}
	fmt.Printf("Category name: %s\n", categoryName.String)
}
if err := rows.Err(); err != nil {
	log.Fatal(err)
}

var subCategoryID int
err = db.QueryRow("SELECT get_sub_categoryID(?)", "Food").Scan(&subCategoryID)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Sub-category ID for Food: %d\n", subCategoryID)

var userID int
err = db.QueryRow("SELECT get_userID(?)", "testuser").Scan(&userID)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("User ID for testuser: %d\n", userID)

// Insert new transaction

/*lignes, err := db.Query("SELECT new_transaction(?, ?, ?, ?)", "testuser", 100, "Groceries", "Food")
if err != nil {
	log.Fatal("erreur1")
}
defer lignes.Close()

for lignes.Next() {
	var result sql.NullString
	err := lignes.Scan(&result)
	if err != nil {
		log.Fatal("erreur2")
	}
	fmt.Println(result.String)
}

if err := lignes.Err(); err != nil {
	log.Fatal("erreur3")
}*/
