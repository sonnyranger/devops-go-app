package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	tasksAdded = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_tasks_added_total",
			Help: "Total number of tasks added",
		},
	)
)

func init() {
	prometheus.MustRegister(tasksAdded)
}

// Todo структура задачи
type Todo struct {
	Item string
}

// PageData данные для шаблона
type PageData struct {
	Title string
	Todos []Todo
}

var (
	// Используем мьютекс для безопасной работы с данными из разных потоков
	mu    sync.Mutex
	todos []Todo
	// Шаблон HTML
	tmpl = template.Must(template.ParseFiles("index.html"))
)

// Обработчик главной страницы
func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	data := PageData{
		Title: "Мой список задач",
		Todos: todos,
	}
	tmpl.Execute(w, data)
}

// Обработчик добавления задачи
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	task := r.FormValue("task")
	if task != "" {
		mu.Lock()
		todos = append(todos, Todo{Item: task})
		mu.Unlock()
	}
	tasksAdded.Inc()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
