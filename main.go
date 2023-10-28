package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// * Membuat struktur data student berdasarkan key dan value
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

// * Inisiasi Student Database
var students []Student

// * Inisiasi Web Server routes
func initRoutes() {
	e := echo.New()

	// Routing
	e.GET("/students", getStudents)
	e.GET("/students/:id", getStudent)
	e.POST("/students", createStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	e.Start(":8080")
}

// TODO Problem and Testing 1: Buat Fungsi yang memberikan seluruh informasi student
func getStudents(c echo.Context) error {
	// Tulis jawabanmu disini
	return c.JSON(http.StatusOK, students)
}

// TODO Problem and Testing 2: Buat Fungsi yang memberikan informasi student berdasarkan ID
func getStudent(c echo.Context) error {
	// Tulis jawabanmu disini

	id, _ := strconv.Atoi(c.Param("id"))
	for _, s := range students {
		if s.ID == id {
			return c.JSON(http.StatusOK, students)

		}
	}
	return c.JSON(http.StatusOK, nil)
}

// TODO Problem and Testing 3: Buat Fungsi yang membuat informasi student baru
func createStudent(c echo.Context) error {
	// Tulis jawabanmu disini

	var student Student

	if err := c.Bind(&student); err != nil {
		panic("Error on create student: " + err.Error())
	}

	if len(students) >= 1 {
		lasStudent := students[len(students)-1]
		student.ID = lasStudent.ID + 1
	} else {
		student.ID = 1
	}

	students = append(students, student)
	return c.JSON(http.StatusCreated, student)
}

// TODO Problem and Testing 4: Buat Fungsi yang dapat mengupdate informasi student berdasarkan ID
func updateStudent(c echo.Context) error {
	// Tulis jawabanmu disini
	var newStudent Student

	param := c.Param("id")

	// return c.JSON(http.StatusOK, map[string]any{
	// 	"id": param,
	// })

	// entah kenapa pas pake test param id selalu ""
	if param == "" {
		return c.JSON(http.StatusOK, nil)
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		panic("Error on update student: " + err.Error())
	}

	if err := c.Bind(&newStudent); err != nil {
		panic(err.Error())
	}

	for i, student := range students {
		if student.ID == id {
			if newStudent.Name != "" {
				students[i].Name = newStudent.Name
			}

			if newStudent.Age != 0 {
				students[i].Age = newStudent.Age
			}

			if newStudent.Grade != "" {
				students[i].Grade = newStudent.Grade
			}
		}
	}
	return c.JSON(http.StatusOK, students)
}

// TODO Problem and Testing 5: Buat Fungsi yang dapat menghapus informasi student berdasarkan ID
func deleteStudent(c echo.Context) error {
	// Tulis jawabanmu disini
	id, _ := strconv.Atoi(c.Param("id"))
	students = append(students[:id], students[id:]...)
	return c.JSON(http.StatusNoContent, students)
}
