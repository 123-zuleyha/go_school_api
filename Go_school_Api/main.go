package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var students = []Student{
	{ID: 1, Name: "ayse", Class: "1-b", Teacher: "kemal"},
	{ID: 2, Name: "merve", Class: "1-c", Teacher: "kemal"},
}

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Teacher string `json:"teacher"`
}

func listStudents(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, students)

}

func createStudent(context *gin.Context) {
	var studentByUser Student
	err := context.BindJSON(&studentByUser)

	if err == nil && studentByUser.ID != 0 && studentByUser.Class != "" && studentByUser.Name != "" && studentByUser.Teacher != "" {
		students = append(students, studentByUser)
		context.IndentedJSON(http.StatusCreated, gin.H{"message": "student has been created", "student_id": studentByUser.ID})
		return

	} else {

		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "student cannot  be created"})

		return

	}
}
func getStudentByID(int_id int) (*Student, error) {

	for i, s := range students {
		if s.ID == int_id {
			return &students[i], nil
		}

	}

	return nil, errors.New("Student cannot be found ")

}

func getStudent(context *gin.Context) {
	str_id := context.Param("id")
	int_id, err := strconv.Atoi(str_id)
	if err != nil {
		panic(err)
	}
	student, err := getStudentByID(int_id)
	if err == nil {
		context.IndentedJSON(http.StatusOK, student)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "student cannot be found"})
	}

}

func main() {
	router := gin.Default()
	router.GET("/students", listStudents)
	router.POST("/students", createStudent)
	router.GET("/students/:id", getStudent)
	router.Run("localhost:9090")

}
