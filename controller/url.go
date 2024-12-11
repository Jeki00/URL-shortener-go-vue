package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/Jeki00/UrlShortener_go_vue/database"
)

type tautan struct {
	Id        int    `json:"id"`
	RealUrl   string `json:"realUrl"`
	ShortUrl  string `json:"shortUrl"`
	Timestamp string `json:"timestamp"`
}


func GetAll(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	rows, err := db.Query(`SELECT real_url, short_url FROM url ORDER BY id ASC`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	defer db.Close()
	var arr []tautan

	for rows.Next() {
		var data tautan
		err = rows.Scan(&data.RealUrl, &data.ShortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		arr = append(arr, data)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(arr)

}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	short_url := r.PathValue("url")
	var data tautan
	err = db.QueryRow(`SELECT id,real_url, short_url, timestamp FROM url WHERE short_url = $1`, short_url).Scan(&data.Id, &data.RealUrl, &data.ShortUrl, &data.Timestamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer db.Close()

	date := time.Now().Format("2006-01-02")
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	_, err = db.Exec("INSERT INTO accessed (url_id, date, timestamp) VALUES ($1,$2,$3)", data.Id, date, timestamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("https://%s", data.RealUrl), http.StatusSeeOther)
	return

}

func DetailUrl(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	query := `
	SELECT ur.short_url, ur.real_url, COALESCE(ac.date,CURRENT_DATE), COUNT(ac.date) FROM url AS ur
	LEFT JOIN accessed AS ac ON ur.id = ac.url_id
	WHERE ur.short_url = $1
	GROUP BY 1, 2,3 
	`

	rows, err := db.Query(query, r.PathValue("shortUrl"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	defer db.Close()

	type detail struct {
		ShortUrl string `json:"shortUrl"`
		RealUrl  string `json:"realUrl"`
		Date     string `json:"date"`
		Count    int    `json:"count"`
	}
	var arr []detail
	for rows.Next() {
		var data detail
		if err := rows.Scan(&data.ShortUrl, &data.RealUrl, &data.Date, &data.Count); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		arr = append(arr, data)
	}

	// result := json.Unmarshal([]byte(arr))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(arr)
}

func PostUrl(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	var url tautan
	defer db.Close()

	err = json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, "cant read body", http.StatusInternalServerError)
		return
	}

	// regex for trimming https link
	s := url.RealUrl
	reg := regexp.MustCompile(`^https?://`)
	trimmed := reg.ReplaceAllString(s, "")

	// get time now
	url.Timestamp = time.Now().Format("2006-01-02 15:04:05")

	_, err = db.Exec(`INSERT INTO url (real_url, short_url, timestamp) VALUES ($1,$2,$3)`, trimmed, url.ShortUrl, url.Timestamp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Berhasil")

}
