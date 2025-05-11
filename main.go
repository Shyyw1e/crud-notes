package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Note struct{
	ID int `json:"id"`
	Text string `json:"text,omitempty"`
}

var(
	notes []Note
	nextID = 1
	mu sync.Mutex
)

// func deleteNote(id int) {
	
// }

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		mu.Lock()
	
		data := make([]Note, len(notes))
		copy(data, notes)
		mu.Unlock()

		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Failed to encode JSON slice", http.StatusBadRequest)
			return
		}
	case "POST":
		defer r.Body.Close()
	
		var req struct {
			Text string `json:"text"`
		}
	
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if req.Text == "" {
			http.Error(w, "Empty JSON text", http.StatusBadRequest)
			return 
		}
		mu.Lock()
		note := Note{ID: nextID, Text: req.Text}
		nextID++
		notes = append(notes, note)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(note)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// findIndex возвращает индекс заметки с Note.ID == id или -1, если нет.
func findIndex(id int) int {
    mu.Lock()
    defer mu.Unlock()
    for i, note := range notes {
        if note.ID == id {
            return i
        }
    }
    return -1
}


func noteByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/notes/"))
	if err != nil {
		http.Error(w, "invalid URL", http.StatusNotFound)
		return
	}
	idx := findIndex(id)
	if idx < 0 {
		http.Error(w, "Not found", http.StatusNotFound)
   	 	return
	}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		mu.Lock()
		note := notes[idx]
		mu.Unlock()
		err := json.NewEncoder(w).Encode(note)
		if err != nil {
			http.Error(w, "failed to encode JSON", http.StatusBadRequest)
			return
		}
	case http.MethodPut:
		defer r.Body.Close()
		var req struct {
			Text string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode JSON", http.StatusBadRequest)
			return
		}
		if req.Text == "" {
			http.Error(w, "Empty note", http.StatusBadRequest)
			return
		}
		mu.Lock()
		notes[idx].Text = req.Text
		refreshedNote := notes[idx]
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(refreshedNote)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		mu.Lock()
		defer mu.Unlock()

		copy(notes[idx:], notes[idx+1:])
		notes = notes[:len(notes)-1]

		w.WriteHeader(http.StatusNoContent)

	default:
    	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func main() {
	http.HandleFunc("/api/notes", NotesHandler)
	http.HandleFunc("/api/notes/", noteByIDHandler)
	fmt.Println("Server listening at :8080")
	http.ListenAndServe(":8080", nil)
	
}