package handlers

import (
	"example/hello/helpers"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to homepage"))

	data := map[string]string{
		"title":   "I'm learn Golang",
		"content": "I'm on my way",
		"footer":  "\u00A9 Selly Supriyatin 2025",
	}

	helpers.LoadTemplate(w, "./pages/index.html", data)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Product Pages"))
	id := r.URL.Query().Get("id")

	if id == "" {
		// http.Error(w, "Product id is required", http.StatusBadRequest)
		data := map[string]interface{}{
			"content": "",
			"text":    "Product id is required",
		}

		helpers.LoadTemplate(w, "./pages/invalidPage.html", data)
		return
	}

	idNumb, _ := strconv.Atoi(id)

	if idNumb < 1 {
		// http.Error(w, "Invalid product Id", http.StatusBadRequest)
		data := map[string]interface{}{
			"content": idNumb,
			"text":    "Invalid Product Id",
		}

		helpers.LoadTemplate(w, "./pages/invalidPage.html", data)
		return
	}
	// fmt.Fprintf(w, "Product pages ke %d", idNumb)

	data := map[string]int{
		"content": idNumb,
	}

	helpers.LoadTemplate(w, "product.html", data)
}
