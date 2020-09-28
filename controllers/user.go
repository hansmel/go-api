package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"hanmel.com/webservice/fileio"
	"hanmel.com/webservice/models"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var httpRequestTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of http requests.",
	},
	[]string{"handler", "method"},
)

var httpRequestDuration = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "http_request_duration_seconds",
		Help: "The duration for the http request.",
	},
	[]string{"handler", "method"},
)

var httpRequestDurationsHistogram = promauto.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_histogram_seconds",
		Help:    "The request durations histogram",
		Buckets: prometheus.DefBuckets,
	},
)

type userController struct {
	userIDPattern *regexp.Regexp
}

var start time.Time

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start = time.Now()

	// Simulate request latency
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Request took %s\n", elapsed)
	fmt.Printf("Request took %d microseconds\n", elapsed.Microseconds())

	elapsedSeconds := float64(time.Since(start).Microseconds()) / 1000000
	fmt.Printf("Request took %f seconds\n", elapsedSeconds)
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Http GET (all)")
	httpRequestTotal.With(prometheus.Labels{"handler": "/users", "method": "GETALL"}).Inc()
	elapsedSeconds := float64(time.Since(start).Microseconds()) / 1000000
	httpRequestDuration.With(prometheus.Labels{"handler": "/users", "method": "GETALL"}).Set(elapsedSeconds)
	httpRequestDurationsHistogram.Observe(elapsedSeconds)
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	fmt.Println("Http GET (by id)")
	httpRequestTotal.With(prometheus.Labels{"handler": "/users", "method": "GET"}).Inc()
	elapsedSeconds := float64(time.Since(start).Microseconds()) / 1000000
	httpRequestDuration.With(prometheus.Labels{"handler": "/users", "method": "GET"}).Set(elapsedSeconds)
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Http POST")
	httpRequestTotal.With(prometheus.Labels{"handler": "/users", "method": "POST"}).Inc()
	elapsedSeconds := float64(time.Since(start).Microseconds()) / 1000000
	httpRequestDuration.With(prometheus.Labels{"handler": "/users", "method": "POST"}).Set(elapsedSeconds)
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fileio.WriteUsers(models.GetUsers())
	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Http PUT")
	httpRequestTotal.With(prometheus.Labels{"handler": "/users", "method": "PUT"}).Inc()
	elapsedSeconds := float64(time.Since(start).Microseconds()) / 1000000
	httpRequestDuration.With(prometheus.Labels{"handler": "/users", "method": "PUT"}).Set(elapsedSeconds)
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fileio.WriteUsers(models.GetUsers())
	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	fmt.Println("Http DELETE")
	httpRequestTotal.With(prometheus.Labels{"handler": "/users", "method": "DELETE"}).Inc()
	elapsedSeconds := float64(time.Since(start).Microseconds()) / 1000000
	httpRequestDuration.With(prometheus.Labels{"handler": "/users", "method": "DELETE"}).Set(elapsedSeconds)
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fileio.WriteUsers(models.GetUsers())
	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
