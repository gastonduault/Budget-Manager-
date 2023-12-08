package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

/*Ce code fourni un serveur web écrit en Go qui expose des endpoints pour gérer un budget.
Le serveur utilise la bibliothèque Gorilla Mux pour router les requêtes HTTP.
Le serveur expose plusieurs endpoints, notamment /login, /transaction, /transaction/add, et /solde.
Le endpoint /login gère les requêtes GET pour connecter un utilisateur. Si la requête est valide, l'utilisateur est connecté et un entier(1) est renvoyé en tant que réponse.
Le endpoint /transaction renvoie toutes les transactions enregistrées dans la base de données en fonction d'un user (identifiant).
Le endpoint /transaction/add permet d'ajouter une nouvelle transaction en utilisant la méthode HTTP GET. Les paramètres de la transaction sont passés dans l'URL en tant que paramètres de requête.
Le endpoint /solde renvoie le solde total ainsi que le solde par catégorie.
Le serveur utilise une base de données MySQL pour stocker les données. Les informations de connexion sont codées en dur dans le code, mais elles peuvent être modifiées pour se connecter à une autre base de données.



*/
// LA LISTE DES STRUCTURES DONT ON AURA BESOIN
/*type New_Transaction struct {
	Login             string `json:"Login"`
	Amount            int    `json:"Amount"`
	Transaction_Name  string `json:"Transaction_Name"`
	Sub_Category_Name string `json:"Sub_Category_Name"`
}*/

// STRUCTURE D'UN UTILISATEUR
type User struct {
	UserID    int    `json:"userID"`
	Login     string `json:"login"`
	Pass_word string `json:"pass_word"`
}

// STRUCTURE D'UNE TRANSACTION
type Transaction struct {
	ID              int     `json:"id"`
	Label           string  `json:"label"`
	Price           float64 `json:"price"`
	PrincipalCat    string  `json:"principal_cat"`
	SubCat          string  `json:"sub_cat"`
	DateTransaction string  `json:"date_transaction"`
}

// sTRUCTURE D'UN SOLDE
type Solde struct {
	Total      float64            `json:"total"`
	ByCategory map[string]float64 `json:"by_category"`
}

// MAIN
func main() {

	// INITIALISE LE ROUTEUR
	router := mux.NewRouter()
	router.HandleFunc("/login", handler_login)
	router.HandleFunc("/transaction", handler_transaction)
	router.HandleFunc("/transaction/add", handler_newTransaction)
	router.HandleFunc("/solde", handler_solde)
	log.Fatal(http.ListenAndServe("localhost:80", router))

}

// FONCTION qui recupere en parametre de la requet get, le login, l'amount, le nom de la transaction et de la sous categorie.
// la fonction retourne la reponse de la requete si l'insertion à été faite ou pas.
func handler_newTransaction(w http.ResponseWriter, r *http.Request) {
	//connexion à la base de donnée
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/budget_manager")
	//S'il y a une erreur, panic va nous renvoyer l'erreur explicitement
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Test connections
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Connnection à la base de donnée
	fmt.Println("Database connection successful.")

	// REQUETE GET POUR AJOUTER UNE TRANSACTION
	if r.Method == "GET" {
		fmt.Println("GET NEW TRANSACTION")
		ctx := context.Background()
		// Récupération des arguments de la méthode GET
		keys := r.URL.Query()
		// On récupère le premier argument
		loginTab, ok := keys["login"]
		if !ok || len(loginTab[0]) < 1 {
			fmt.Fprintf(w, "Argument login manquant\n")
			return
		}

		//ces instructions récupèrent la première valeur d'un tableau de chaînes de caractères appelé loginTab
		//et l'affichent dans la réponse HTTP via la variable w
		login := loginTab[0]
		fmt.Fprintf(w, "Argument login : = %s\n", login)

		// On récupère le second argument
		amountTab, ok := keys["amount"]
		if !ok || len(amountTab[0]) < 1 {
			fmt.Fprintf(w, "Argument amount manquant\n")
			return
		}
		amount := amountTab[0]
		fmt.Fprintf(w, "Argument amount : = %s\n", amount)
		// On récupère le troisieme argument
		transactionNameTab, ok := keys["transactionName"]
		if !ok || len(transactionNameTab[0]) < 1 {
			fmt.Fprintf(w, "Argument transactionName manquant\n")
			return
		}
		transactionName := transactionNameTab[0]
		fmt.Fprintf(w, "Argument transactionName : = %s\n", transactionName)
		// On récupère le quatrieme argument
		subCategoryNameTab, ok := keys["subCategoryName"]
		if !ok || len(subCategoryNameTab[0]) < 1 {
			fmt.Fprintf(w, "Argument subCategoryName manquant\n")
			return
		}
		subCategoryName := subCategoryNameTab[0]
		fmt.Fprintf(w, "Argument subCategoryName : = %s\n", subCategoryName)

		// On prepare la requête (vérifie si elle est correct)
		stmt, err := db.Prepare("SELECT `new_transaction`(?,?,?,?) AS `new_transaction`;")
		if err != nil {
			fmt.Println("requete NOK")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		/// On exécute la requête avec les valeur récupérées dans via la méthode GET
		row := stmt.QueryRowContext(ctx, login, amount, transactionName, subCategoryName)

		//Ces instructiions enregistrent le résultat d'une requête SQL dans une variable result de type string.
		var result string
		err = row.Scan(&result)

		// gere les erreurs de requete
		if err != nil {
			fmt.Println("Erreur row")
			fmt.Printf("REPONSE REQUETE (ERREUR) : %s\n", result)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", result)
		fmt.Printf("REPONSE REQUETE : %s\n", result)

		fmt.Println("FIN")

	}

}

// Fonction qui connecte l'user à son compte
func handler_login(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/budget_manager")
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
		fmt.Println("GET login")
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

		/// On exécute la requête avec les valeur récupérées dans via la méthode GET
		row := stmt.QueryRowContext(ctx, login, pass_word)

		var result int64
		err = row.Scan(&result)

		if err != nil {
			fmt.Println("Erreur row")
			fmt.Printf("REPONSE REQUETE (ERREUR) : %d\n", result)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%d", result)
		fmt.Printf("REPONSE REQUETE : %d\n", result)

		fmt.Println("FIN")

		//partie pour renvoyer en json
	}

}

// FONCTION POUR CREER UNE TRANSACTION
func handler_transaction(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/budget_manager")
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

	/*// Gérer la demande POST pour créer une transaction
	if r.Method == "POST" {
		fmt.Println("POST transaction")
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Insérer l'utilisateur dans la base de données
		stmt, err := db.Prepare("INSERT INTO transaction(amount,sub_categoryID) VALUES (?, ?)") //login, montant, libélé, sub-catégorue name
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

	}*/
	//GET

	if r.Method == "GET" {

		fmt.Println("GET transaction")
		ctx := context.Background()

		// Récupération des arguments de la méthode GET
		keys := r.URL.Query()
		// On récupère le premier argument
		transactionTab, ok := keys["accountID"]
		if !ok || len(transactionTab[0]) < 1 {
			fmt.Fprintf(w, "Argument accountID manquant")
			return
		}
		accountID := transactionTab[0]
		fmt.Fprintf(w, "Argument acccountID : = %s\n", accountID)
		// On prepare la requête (vérifie si elle est correct)

		stmt, err := db.Prepare("select user.userID, category.category_name, sub_category.sub_category_name, transaction.amount, transaction.transaction_name, transaction.date FROM transaction INNER JOIN user ON transaction.accountID = user.userID INNER JOIN sub_category ON transaction.sub_categoryID = sub_category.sub_categoryID INNER JOIN category ON category.categoryID = sub_category.categoryID WHERE user.userID = ?")
		if err != nil {
			fmt.Println("requete NOK")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// On exécute la requête avec les valeur récupérées dans via la méthode GET
		row := stmt.QueryRowContext(ctx, accountID)

		var result Transaction
		err = row.Scan(&result.ID, &result.PrincipalCat, &result.SubCat, &result.Price, &result.Label, &result.DateTransaction)

		if err != nil {
			fmt.Println("Erreur row")
			fmt.Printf("REPONSE REQUETE (ERREUR) : %s\n", accountID)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("REPONSE REQUETE SUR ACCOUNT : %s\n", accountID)

		fmt.Println("FIN")

		// Exécuter une requête SELECT pour avoir l'historique d'une personne selon son id
		rows, err := db.Query("select user.userID, category.category_name, sub_category.sub_category_name, transaction.amount, transaction.transaction_name, transaction.date FROM transaction INNER JOIN user ON transaction.accountID = user.userID INNER JOIN sub_category ON transaction.sub_categoryID = sub_category.sub_categoryID INNER JOIN category ON category.categoryID = sub_category.categoryID WHERE user.userID = " + accountID)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		// Créer une slice pour stocker les résultats de la requête
		results := []Transaction{}

		// Parcourir les résultats de la requête et ajouter chaque ligne à la slice
		for rows.Next() {
			var result Transaction
			err := rows.Scan(&result.ID, &result.PrincipalCat, &result.SubCat, &result.Price, &result.Label, &result.DateTransaction)
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
}

func handler_solde(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/budget_manager")
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
	if r.Method == "POST" {
		fmt.Println("POST")
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if r.Method == "GET" {
		fmt.Println("GET SOLDE")
		ctx := context.Background()
		// Récupération des arguments de la méthode GET
		keys := r.URL.Query()
		// On récupère le premier argument
		transactionTab, ok := keys["accountID"]
		if !ok || len(transactionTab[0]) < 1 {
			fmt.Fprintf(w, "Argument accountID manquant")
			return
		}
		accountID := transactionTab[0]
		fmt.Fprintf(w, "Argument acccountID : = %s\n", accountID)
		// On prepare la requête (vérifie si elle est correct)
		stmt, err := db.Prepare("select user.userID, category.category_name, sub_category.sub_category_name, transaction.amount, transaction.transaction_name, transaction.date FROM transaction INNER JOIN user ON transaction.accountID = user.userID INNER JOIN sub_category ON transaction.sub_categoryID = sub_category.sub_categoryID INNER JOIN category ON category.categoryID = sub_category.categoryID WHERE user.userID = ?")
		if err != nil {
			fmt.Println("requete NOK")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// On exécute la requête avec les valeur récupérées dans via la méthode GET
		row := stmt.QueryRowContext(ctx, accountID)
		var result Solde
		err = row.Scan(&result.Total, &result.ByCategory)
		// Exécuter une requête SELECT pour avoir l'historique d'une personne selon son id
		rows, err := db.Query("select user.userID, category.category_name, sub_category.sub_category_name, transaction.amount, transaction.transaction_name, transaction.date FROM transaction INNER JOIN user ON transaction.accountID = user.userID INNER JOIN sub_category ON transaction.sub_categoryID = sub_category.sub_categoryID INNER JOIN category ON category.categoryID = sub_category.categoryID WHERE user.userID = " + accountID)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		// Créer une slice pour stocker les résultats de la requête
		results := []Solde{}

		// Parcourir les résultats de la requête et ajouter chaque ligne à la slice
		for rows.Next() {
			var result Solde
			err := rows.Scan(&result.Total, &result.ByCategory)
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
}

// Call stored procedures
/*
	_, err = db.Exec("CALL add_user(?, ?)", "testuser", "testpass")
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

	// Parce que la requete peut renvoyer un element nul
	rows2, err := db.Query("SELECT get_category_name_from_transaction_ID(?)", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()

	for rows2.Next() {
		var categoryName sql.NullString

		if err := rows2.Scan(&categoryName); err != nil {
			log.Fatal(err)
			fmt.Println("Category name is NULL")
		}
		fmt.Printf("Category name: %s\n", categoryName.String)
	}
	if err := rows2.Err(); err != nil {
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

	lignes, err := db.Query("SELECT new_transaction(?, ?, ?, ?)", "testuser", 100, "Groceries", "Food")
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
