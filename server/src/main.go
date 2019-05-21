package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	gorm.Model
	Name     string `sql:"not null"`
	Password string `sql:"not null"`
	Tasks    []Task `gorm:"foreignkey:UserRefer"`
}

type Task struct {
	gorm.Model
	Description string `sql:"not null"`
	Completed   bool   `sql:"not null"`
	UserRefer   uint   `sql:"not null"`
}

var db *gorm.DB
var allUsers []User

func main() {
	// instantiate echo
	e := echo.New()

	db := gormConnect()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	defer db.Close()

	// routing
	e.POST("/signup", Signup())

	// launch server
	e.Start(":1313")
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "todoapp"

	CONNECT := USER + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(DBMS, CONNECT)

	db.AutoMigrate(User{})
	db.AutoMigrate(Task{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

func Users() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, allUsers)
	}
}

func Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)

		if err := c.Bind(user); err != nil {
			return err
		}

		if user.Name == "" || user.Password == "" {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "invalid name or password",
			}
		}

		if u := FindUser(&User{Name: user.Name}); u.ID != 0 {
			return &echo.HTTPError{
				Code:    http.StatusConflict,
				Message: "name already exists",
			}
		}

		CreateUser(user)
		user.Password = ""
		return c.JSON(http.StatusCreated, user)
	}
}

func CreateUser(user *User) {
	db.Create(&user)
}

func FindUser(u *User) User {
	var user User
	k := db.Where(u).First(&user)
	return user
}
