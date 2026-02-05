package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake Database -file
var courses []Course

// Middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""

	//if courseId is not empty, User should move further
	return c.CourseName == ""

}

func main() {
	fmt.Println("API - CattyVibes.in")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "BlackBac",
		CoursePrice: 200,
		Author:      &Author{Fullname: "Vibes", Website: "Vibes.dev"},
	})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "WhiteBac",
		CoursePrice: 1100,
		Author:      &Author{Fullname: "Vibes", Website: "go.dev"},
	})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllcourse).Methods("Get")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("Get")
	r.HandleFunc("/courses", CreateOneCours).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", DeleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//Controllers -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Api by VibesAI</h1>"))
}

func getAllcourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	//to set headerss
	w.Header().Set("content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-Type", "applicaton/json")

	//Grap Id from request
	params := mux.Vars(r)
	// fmt.Println(params)

	// loop throuhg the courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func CreateOneCours(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("content-Type", "applicaton/json")

	//what if: body is empthy
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what abt data send like this -{}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside Json")
		return
	}
	for _, existingcourse := range courses {
		if existingcourse.CourseName == course.CourseName {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}
	//Todo Check only if title is duplicate,
	// loop, title matches with course name, json (exact code already exist )

	//Generate a Unique Id, convert to strings
	//append course to courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("content-Type", "applicaton/json")

	// First grab id from request
	params := mux.Vars(r)

	// Loop, after get ID, Remove operation, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//Todo send a response when Id is not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"Message": "Course not found",
	})
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("content-Type", "applicaton/json")

	param := mux.Vars(r)

	//loop, id, remove(INDEX, INDEX+1)
	for index, course := range courses {
		if course.CourseId == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			// todo send a confirm or deny response
			json.NewEncoder(w).Encode("Course Deleted Successfully")
			break

		}
	}
}
