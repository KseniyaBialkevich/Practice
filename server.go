package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// // HTTP-сервер
//*************************************************************************************
// func main () {
// 	// http://localhost:8080/ - слэш игнорируется браузером
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // обработчик, возвращающий "Hello World!"
// 		fmt.Fprintf(w, "Hello World!")
// 	})

// 	http.ListenAndServe(":8080", nil) // http.ListenAndServe принимает входящие запросы; задаем слушать порт 8080
// }
//*************************************************************************************
// func main () {
// 	// http://localhost:8080/hello
// 	helloHandler := func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello World!")
// 	}

// 	http.HandleFunc("/hello", helloHandler) // для обработки запросов передаем объект helloHandler;
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
//*************************************************************************************
// type msg string
// func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
//    fmt.Fprint(resp, m)
// }
// func main() {
// 	// http://localhost:8181/
//     msgHandler := msg("Hello from Web Server in Go")
//     fmt.Println("Server is listening...")
//     http.ListenAndServe("localhost:8181", msgHandler)
// }
//*************************************************************************************
// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
//     http.HandleFunc("/", handler)
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }
//*************************************************************************************
// type Page struct {
// 	Title string `json:"title"`
// 	Body  []byte `json:"body"`
// }

// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return ioutil.WriteFile(filename, p.Body, 0600)
// }

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".txt"
// 	body, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):] // len - 6, title - "test".

// 	p, _ := loadPage(title)

// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

// //http://localhost:8080/view/test
// func main() {

// 	// p1 := &Page{
// 	// 	Title: "test",
// 	// 	Body:  []byte("Hello World!")}
// 	// p1.save()
// 	// p2, _ := loadFile(p1.Title)
// 	// fmt.Println(string(p2.Body))

// 	http.HandleFunc("/view/", viewHandler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

//*************************************************************************************
// func main() {
// 	//http.HandleFunc("/", handler)

// 	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintf(resp, "I love %s!", req.URL.Path[1:])
// 	})

// 	fmt.Println("Server is listening ...")

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
//*************************************************************************************
// func main() {
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "hello.html")
// 	})
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }
//*************************************************************************************
// type Handler struct {
// 	Message string `json:"message"`
// }

// func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, h.Message)
// }

// func main() {
// 	h1 := Handler{Message: "Index"}
// 	h2 := Handler{Message: "About"}

// 	http.Handle("/", h1)
// 	http.Handle("/about", h2)

// 	fmt.Println("Server is listening ...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }
//*************************************************************************************
// func main() {

// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/", fs)

// 	http.HandleFunc("/extra", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "extra/extra.html")
// 	})

// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "hello.html")
// 	})

// 	http.HandleFunc("/string", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello! I'm STRING!")
// 	})

// 	fmt.Println("Server is listening...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }
//*************************************************************************************
//  http://localhost:8181/products/2
// func productsHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r) // map[id:2]
// 	id := vars["id"]    // 2
// 	response := fmt.Sprintf("Product %s", id)
// 	fmt.Fprintf(w, response)
// }

// func main() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
// 	http.Handle("/", router)

// 	fmt.Println("Server listening...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }
//*************************************************************************************
//  http://localhost:8181/products/phones/2
// func productsHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	category := vars["category"]
// 	response := fmt.Sprintf("Product category=%s id=%s", category, id)
// 	fmt.Fprintf(w, response)
// }

// func main() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/products/{category}/{id:[0-9]+}", productsHandler)
// 	http.Handle("/", router)

// 	fmt.Println("Server is listening...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }
//*************************************************************************************
// http://localhost:8181/products/2
// http://localhost:8181/articles/5
// http://localhost:8181/
// func productsHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	response := fmt.Sprintf("id=%s", id)
// 	fmt.Fprintf(w, response)
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Index Page")
// }

// func main() {

// 	router := mux.NewRouter()
// 	router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
// 	router.HandleFunc("/articles/{id:[0-9]+}", productsHandler)
// 	router.HandleFunc("/", indexHandler)
// 	http.Handle("/", router)

// 	fmt.Println("Server is listening...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }
//*************************************************************************************
// http://localhost:8181/World
// func helloName(w http.ResponseWriter, r *http.Request) {
// 	name := chi.URLParam(r, "name")
// 	response := fmt.Sprintf("Hello %s!", name)

// 	fmt.Fprintf(w, response)
// }

// func main() {

// 	router := chi.NewRouter()
// 	router.Get("/one/{name}", helloName)
// 	router.Get("/two/{name}", helloName)
// 	http.Handle("/", router)

// 	fmt.Println("Server is listening...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }
//*************************************************************************************
// http://localhost:8181/user?name=Same&age=21
// func main() {
// 	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
// 		url := r.URL // /user?name=Same&age=21
// 		query := r.URL.Query() // map[age:[21] name:[Same]]

// 		name := r.URL.Query().Get("name") // Same
// 		age := r.URL.Query().Get("age") // 21

// 		fmt.Fprintf(w, "Name: %s and Age: %s", name, age)
// 	})

// 	fmt.Println("Server is listening...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }

//*************************************************************************************
// http://localhost:8181/
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "user.html")
// 	})

// 	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {

// 		name := r.FormValue("username")
// 		age := r.FormValue("userage")

// 		fmt.Fprintf(w, "Name: %s and Age: %s", name, age)
// 	})

// 	fmt.Println("Server is listening...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }

//******************************SQL*******************************************************
// type Product struct {
// 	ID      int
// 	Model   string
// 	Company string
// 	Price   int
// }

// func main() {

// 	db, err := sql.Open("mysql", "root:11111111@/productdb")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM productdb.products")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	pProducts := []Product{}

// 	for rows.Next() {
// 		p := Product{}
// 		err := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
// 		if err != nil {
// 			fmt.Println(err)

// 			continue
// 		}

// 		pProducts = append(pProducts, p)
// 	}

// 	for _, p := range pProducts {
// 		fmt.Println(p.ID, p.Model, p.Company, p.Price)
// 	}

// }
//******************************SQL*******************************************************

// func main() {

// 	db, err := sql.Open("mysql", "root:11111111@/productdb")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	result, err := db.Exec("INSERT INTO productdb.products (model, company, price) VALUES ('LG G6', 'LG', 400)")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result.LastInsertId()) // id добавленного объекта
// 	fmt.Println(result.RowsAffected()) // количество затронутых строк

// }

//******************************SQL*******************************************************

// type Product struct {
// 	ID      int
// 	Model   string
// 	Company string
// 	Price   int
// }

// func main() {

// 	db, err := sql.Open("mysql", "root:11111111@/productdb")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	row := db.QueryRow("SELECT * FROM productdb.products WHERE id = ?", 4)

// 	p := Product{}

// 	err = row.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(p.ID, p.Model, p.Company, p.Price)

// }

//******************************SQL*******************************************************

// type Product struct {
// 	ID      int
// 	Model   string
// 	Company string
// 	Price   int
// }

// func main() {

// 	db, err := sql.Open("mysql", "root:11111111@/productdb")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	result, err := db.Exec("UPDATE productdb.products SET company = ?, price = ? WHERE id = ?", "Samsung", 1300, 2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result.LastInsertId())
// 	fmt.Println(result.RowsAffected())
// }

//******************************SQL*******************************************************

// type Product struct {
// 	ID      int
// 	Model   string
// 	Company string
// 	Price   int
// }

// func main() {

// 	db, err := sql.Open("mysql", "root:11111111@/productdb")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	result, err := db.Exec("DELETE FROM productdb.products WHERE id = 3")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result.LastInsertId()) // id последнего удаленого объекта
// 	fmt.Println(result.RowsAffected()) // количество затронутых строк
// }

//******************************TEMPLATE*******************************************************

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		data := "Go Template"

// 		tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")

// 		tmpl.Execute(w, data)

// 	})

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)

// }

//******************************TEMPLATE*******************************************************

// type ViewData struct {
// 	Title   string
// 	Message string
// }

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		data := ViewData{
// 			Title:   "World Cup",
// 			Message: "FIFA will never regret it",
// 		}

// 		tmpl := template.Must(template.New("data").Parse(`<div>
// <h1>{{ .Title}}</h1>
// <p>{{ .Message}}</p>
// </div>`))

// 		tmpl.Execute(w, data)
// 	})

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)
// }

//******************************TEMPLATE*******************************************************

// type ViewData struct {
// 	Title   string
// 	Message string
// }

// func main() {

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		data := ViewData{
// 			Title:   "World Cup",
// 			Message: "FIFA will never regret it",
// 		}

// 		tmpl, _ := template.ParseFiles("templates/index.html")

// 		tmpl.Execute(w, data)

// 	})

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)

// }

//******************************TEMPLATE*******************************************************

// type ViewData struct {
// 	Title string
// 	Users []string
// }

// func main() {

// 	data := ViewData{
// 		Title: "Users List",
// 		Users: []string{"Tom", "Bob", "Sam"},
// 	}

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		tmpl, _ := template.ParseFiles("templates/range_index.html")

// 		tmpl.Execute(w, data)
// 	})

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)
// }

//******************************TEMPLATE*******************************************************

// type ViewData struct {
// 	Title string
// 	Users []User
// }

// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {

// 	data := ViewData{
// 		Title: "Users List",
// 		Users: []User{
// 			{Name: "Tom", Age: 21},
// 			{Name: "Kate", Age: 23},
// 			{Name: "Alice", Age: 25},
// 		},
// 	}

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		tmpl, _ := template.ParseFiles("templates/complex_data.html")
// 		tmpl.Execute(w, data)
// 	})

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)
// }

//******************************TEMPLATE*******************************************************

// type ViewData struct {
// 	Available bool
// }

// func main() {

// 	data := ViewData{
// 		Available: false,
// 	}

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		tmpl, _ := template.ParseFiles("templates/condition.html")

// 		tmpl.Execute(w, data)
// 	})

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)
// }

//******************************TEMPLATE*******************************************************

// type ViewData struct {
// 	Hour int
// }

// func main() {

// 	data := ViewData{
// 		Hour: time.Now().Hour(),
// 	}

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 		tmpl, _ := template.ParseFiles("templates/compare.html")

// 		tmpl.Execute(w, data)
// 	})

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)
// }

//******************************SERVER & SQL*******************************************************

// type Product struct {
// 	ID      int
// 	Model   string
// 	Company string
// 	Price   int
// }

// var database *sql.DB

// func IndexHandler(w http.ResponseWriter, r *http.Request) {

// 	rows, err := database.Query("SELECT * FROM productdb.products")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	pProducts := []Product{}

// 	for rows.Next() {

// 		p := Product{}

// 		err := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
// 		if err != nil {
// 			fmt.Println(err)

// 			continue
// 		}

// 		pProducts = append(pProducts, p)
// 	}

// 	tmpl, _ := template.ParseFiles("tmpls/index.html")
// 	tmpl.Execute(w, pProducts)

// }

// func main() {

// 	db, err := sql.Open("mysql", "root:11111111@/productdb")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	database = db

// 	defer db.Close()

// 	http.HandleFunc("/", IndexHandler)

// 	fmt.Println("Server is listening...")
// 	log.Fatal(http.ListenAndServe(":8181", nil))
// }

//******************************SERVER & SQL*******************************************************

// type Product struct {
// 	ID      int
// 	Model   string
// 	Company string
// 	Price   int
// }

// var database *sql.DB

// // функция добавления данных
// func CreateHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		model := r.FormValue("model")
// 		company := r.FormValue("company")
// 		price := r.FormValue("price")

// 		_, err = database.Exec("INSERT INTO productdb.products (model, company, price) VALUES (?, ?, ?)", model, company, price)
// 		if err != err {
// 			log.Println(err)
// 		}

// 		http.Redirect(w, r, "/", 301) // переадресация на корень сайта

// 	} else {
// 		http.ServeFile(w, r, "tmpls/create.html")
// 	}
// }

// func IndexHandler(w http.ResponseWriter, r *http.Request) {

// 	rows, err := database.Query("SELECT * FROM productdb.products")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer rows.Close()

// 	pProducts := []Product{}

// 	for rows.Next() {
// 		p := Product{}

// 		err := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
// 		if err != nil {
// 			log.Println(err)

// 			continue
// 		}

// 		pProducts = append(pProducts, p)
// 	}

// 	tmpl, _ := template.ParseFiles("tmpls/index.html")
// 	tmpl.Execute(w, pProducts)
// }

// func main() {

// 	db, err := sql.Open("mysql", "root:11111111@tcp(localhost:3306)/productdb") // Порт 3306 - это порт по умолчанию MySql
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	database = db

// 	defer db.Close()

// 	http.HandleFunc("/", IndexHandler)
// 	http.HandleFunc("/create", CreateHandler)

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)

// }

//******************************SERVER & SQL*******************************************************

// type Product struct {
// 	ID      int
// 	Model   string
// 	Company string
// 	Price   int
// }

// var database *sql.DB

// // возвращаем пользователю страницу для редактирования объекта
// func EditPage(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	row := database.QueryRow("SELECT * FROM productdb.products WHERE id = ?", id)

// 	prodt := Product{}

// 	err := row.Scan(&prodt.ID, &prodt.Model, &prodt.Company, &prodt.Price)
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, http.StatusText(404), http.StatusNotFound)
// 	} else {
// 		tmpl, _ := template.ParseFiles("tmpls/edit.html")
// 		tmpl.Execute(w, prodt)
// 	}
// }

// // получаем измененные данные и сохраняем их в БД
// func EditHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	id := r.FormValue("id")
// 	model := r.FormValue("model")
// 	company := r.FormValue("company")
// 	price := r.FormValue("price")

// 	_, err = database.Exec("UPDATE productdb.products set model = ?, company = ?, price = ? WHERE id = ?", model, company, price, id)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	http.Redirect(w, r, "/", 301)
// }

// func CreateHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {

// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		model := r.FormValue("model")
// 		company := r.FormValue("company")
// 		price := r.FormValue("price")

// 		_, err = database.Exec("INSERT INTO productdb.products (model, company, price) VALUES (?, ?, ?)", model, company, price)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		http.Redirect(w, r, "/", 301)
// 	} else {
// 		http.ServeFile(w, r, "tmpls/create.html")
// 	}
// }

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	rows, err := database.Query("SELECT * FROM productdb.products")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer rows.Close()

// 	pProducts := []Product{}

// 	for rows.Next() {
// 		prodt := Product{}

// 		err := rows.Scan(&prodt.ID, &prodt.Model, &prodt.Company, &prodt.Price)
// 		if err != nil {
// 			log.Println(err)

// 			continue
// 		}

// 		pProducts = append(pProducts, prodt)
// 	}

// 	tmpl, _ := template.ParseFiles("tmpls/indexedt.html")
// 	tmpl.Execute(w, pProducts)
// }

// func main() {

// 	db, err := sql.Open("mysql", "root:11111111@tcp(localhost:3306)/productdb")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	database = db

// 	defer db.Close()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/", IndexHandler)
// 	router.HandleFunc("/create", CreateHandler)
// 	router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
// 	router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")

// 	http.Handle("/", router)

// 	fmt.Println("Server is listening...")
// 	http.ListenAndServe(":8181", nil)
// }

//******************************SERVER & SQL*******************************************************

type Product struct {
	ID      int
	Model   string
	Company string
	Price   int
}

var database *sql.DB

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec("DELETE FROM productdb.products WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := database.QueryRow("SELECT * FROM productdb.products WHERE id = ?", id)

	p := Product{}

	err := row.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("templates/edit.html")
		tmpl.Execute(w, p)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	id := r.FormValue("id")
	model := r.FormValue("model")
	company := r.FormValue("company")
	price := r.FormValue("price")

	_, err = database.Exec("UPDATE productdb.products set model = ?, company = ?, price = ? WHERE id = ?", model, company, price, id)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		model := r.FormValue("model")
		company := r.FormValue("company")
		price := r.FormValue("price")

		_, err = database.Exec("INSERT INTO productdb.products (model, company, price) VALUES (?, ?, ?)", model, company, price)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 301)

	} else {
		http.ServeFile(w, r, "templates/create.html")
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("SELECT * FROM productdb.products")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	pProducts := []Product{}

	for rows.Next() {
		p := Product{}

		err := rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)

			continue
		}

		pProducts = append(pProducts, p)
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, pProducts)
}

func main() {

	db, err := sql.Open("mysql", "root:11111111@tcp(localhost:3306)/productdb")
	if err != nil {
		log.Println(err)
	}

	database = db

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/create", CreateHandler)
	router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", DeleteHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
