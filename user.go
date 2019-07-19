package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"github.com/domodwyer/mailyak"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var mySigningKey = []byte("secret")
var db *gorm.DB
var err error

type BookedDetails struct {
	gorm.Model
	Userid        string
	Eventid       string
	Eventname     string
	Eventtype     string
	Stadiumtype   string
	Eventdate     string
	Eventlocation string
	Ticketbooked  string
	Amountpaid    string
}

type EventBookDetails struct {
	gorm.Model
	Eventid       string
	Eventname     string
	Aboutevent    string
	Eventtype     string
	Stadiumtype   string
	Numberofseats string
	Price         string
	Discount      string
	Location      string
	Dateofevent   string
	Booked        string
	Bimage        string
}
type User struct {
	gorm.Model
	Userid      string
	FirstName   string
	LastName    string
	DateOfBirth string
	Gender      string
	Typeoflogin string
	Mobile      string
	Email       string
	Password    string
}

func InitialMigration() {
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user_details password=bgsathish@20 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Data base not connected")
	} else {
		fmt.Println("Connected DB Successfully")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&EventBookDetails{})
	db.AutoMigrate(&BookedDetails{})
}

// func Alluser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user_details password=bgsathish@20 sslmode=disable")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Could not connect")
// 	} else {
// 		fmt.Println("Connected DB Successfully")
// 	}
// 	defer db.Close()

// 	var dispUser []User
// 	db.Find(&dispUser)
// 	json.NewEncoder(w).Encode(dispUser)
// 	fmt.Fprint(w, "All user displayed")
// }

func BookEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user_details password=bgsathish@20 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect")
	} else {
		fmt.Println("Connected DB Successfully")
	}
	defer db.Close()

	vars := mux.Vars(r)
	userid := vars["userid"]
	usermail := vars["usermail"]
	fname := vars["fname"]
	lname := vars["lname"]
	bimage := vars["bimage"]
	eventid := vars["eventid"]
	eventname := vars["eventname"]
	eventtype := vars["eventtype"]
	stadiumtype := vars["stadiumtype"]
	eventdate := vars["eventdate"]
	eventlocation := vars["eventlocation"]
	ticketbooked := vars["ticketbooked"]
	amountpaid := vars["amountpaid"]
	db.Create(&BookedDetails{Userid: userid, Eventid: eventid, Eventname: eventname, Eventtype: eventtype, Stadiumtype: stadiumtype, Eventdate: eventdate, Eventlocation: eventlocation, Ticketbooked: ticketbooked, Amountpaid: amountpaid})
	var ticketToBook EventBookDetails
	db.Where("eventid =?", eventid).Find(&ticketToBook)
	alreadyBooked := ticketToBook.Booked
	alreadyBookedToInt, _ := strconv.Atoi(alreadyBooked)
	nowBooked, _ := strconv.Atoi(ticketbooked)
	finalVal := alreadyBookedToInt + nowBooked
	ticketToBook.Booked = strconv.Itoa(finalVal)
	db.Save(&ticketToBook)
	mailSending(fname, lname, usermail, bimage, eventname, eventdate, eventlocation, ticketbooked, amountpaid)
}
func mailSending(fname string, lname string, email string, bimage string, eventname string, eventdate string, eventlocation string, ticketbooked string, amountpaid string) {
	bimage = strings.ReplaceAll(bimage, "-", "/")
	mail := mailyak.New("smtp.gmail.com:587", smtp.PlainAuth("", "sathishbalucs@gmail.com", "98425588833", "smtp.gmail.com"))

	mail.To(email)
	mail.From("oops@itsallbroken.com")
	mail.Subject("Event Booking")
	mail.HTML().Set(
		"<html>" + "\r\n" +
			"<head> " + "\r\n" +
			"<link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css'>" + "\r\n" +
			"<script src='https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js'></script>" + "\r\n" +
			"<script src='https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js'></script>" + "\r\n" +
			"<script src='https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js'></script>" + "\r\n" +
			"</head>" + "\r\n" +
			"<body>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"<div class='col-md-6 col-xl-6'>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"<div class='col-md-12 col-xl-12'>" + "\r\n" +
			"<img src='https://www.brandcrowd.com/gallery/brands/pictures/picture14483012951724.png' height='200' width='200' />" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"<div class='col-md-6 col-xl-6'>" + "\r\n" +
			"<h5> Dear " + fname + " " + lname + ", </h5>" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"<div class='col-md-6 col-xl-6'>" + "\r\n" +
			"<h5> Event Name :" + eventname + "<h5>" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"<div class='col-md-6 col-xl-6'>" + "\r\n" +
			"<h5> Event Date :" + eventdate + "<h5>" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"<div class='col-md-6 col-xl-6'>" + "\r\n" +
			"<h5> Event Location :" + eventlocation + " <h5>" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"<div class='col-md-6 col-xl-6'>" + "\r\n" +
			"<h5> Ticket Booked :" + ticketbooked + " <h5>" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"<div class='row'>" + "\r\n" +
			"<div class='col-md-6 col-xl-6'>" + "\r\n" +
			"<h5> Amount Paid :" + amountpaid + " <h5>" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"</div>" + "\r\n" +
			"</body>" + "\r\n" +
			"</html>")

	// input can be a bytes.Buffer, os.File, os.Stdin, etc.
	// call multiple times to attach multiple files

	if err := mail.Send(); err != nil {
		panic(" 💣 ")
	}
}

func ShowDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user_details password=bgsathish@20 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect")
	} else {
		fmt.Println("Connected DB Successfully")
	}
	defer db.Close()

	var dispDetails []EventBookDetails
	db.Find(&dispDetails)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = &dispDetails
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := token.SignedString(mySigningKey)
	w.Write([]byte(tokenString))
}
func MyBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user_details password=bgsathish@20 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect")
	} else {
		fmt.Println("Connected DB Successfully")
	}
	defer db.Close()

	var myBookingDetails []BookedDetails
	vars := mux.Vars(r)
	Userid := vars["userid"]
	db.Where("userid=?", Userid).Find(&myBookingDetails)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = &myBookingDetails
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := token.SignedString(mySigningKey)
	w.Write([]byte(tokenString))
}
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user_details password=bgsathish@20 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect")
	} else {
		fmt.Println("Connected DB Successfully")
	}
	defer db.Close()

	var user []User
	// db.Find(&user)
	vars := mux.Vars(r)
	mail := vars["mail"]
	db.Where("email=?", mail).Find(&user)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = &user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := token.SignedString(mySigningKey)
	w.Write([]byte(tokenString))

}

func Newuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user_details password=bgsathish@20 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect")
	} else {
		fmt.Println("Connected DB Successfully")
	}
	defer db.Close()

	vars := mux.Vars(r)
	userid := vars["userid"]
	fname := vars["fname"]
	lname := vars["lname"]
	dob := vars["dob"]
	gender := vars["gender"]
	typeoflogin := vars["typeoflogin"]
	mobile := vars["mobile"]
	mail := vars["mail"]
	pass := vars["password"]

	db.Create(&User{Userid: userid, FirstName: fname, LastName: lname, DateOfBirth: dob, Mobile: mobile, Gender: gender, Typeoflogin: typeoflogin, Email: mail, Password: pass})

}

func EventBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	db, err = gorm.Open("postgres", "port=5432 user=postgres dbname=user_details password=bgsathish@20 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect")
	} else {
		fmt.Println("Connected DB Successfully")
	}
	defer db.Close()

	// vars := mux.Vars(r)
	// eventid := vars["eventid"]
	// eventname := vars["eventname"]
	// aboutevent := vars["aboutevent"]
	// eventtype := vars["eventname"]
	// stadiumtype := vars["eventtype"]
	// numberofseats := vars["numberofseats"]
	// price := vars["price"]
	// discount := vars["discount"]
	// location := vars["location"]
	// dateofevent := vars["dateofevent"]

	// Image upload
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error formed")
		panic(err)

	}
	defer file.Close()

	fmt.Println("MIME Header: ", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.jpg")
	if err != nil {
		fmt.Println(err)

	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	var newEvent EventBookDetails
	newEvent.Eventid = r.FormValue("eventid")
	newEvent.Eventname = r.FormValue("eventname")
	newEvent.Aboutevent = r.FormValue("aboutevent")
	newEvent.Eventtype = r.FormValue("eventtype")
	newEvent.Stadiumtype = r.FormValue("stadiumtype")
	newEvent.Numberofseats = r.FormValue("numberofseats")
	newEvent.Price = r.FormValue("price")
	newEvent.Discount = r.FormValue("discount")
	newEvent.Location = r.FormValue("location")
	newEvent.Dateofevent = r.FormValue("dateofevent")
	newEvent.Booked = r.FormValue("booked")
	newEvent.Bimage = tempFile.Name()

	db.Create(&newEvent)
	fmt.Fprint(w, "Event Booked")

}

// func Deleteuser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprint(w, "Delete user displayed")
// }
// func Updateuser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprint(w, "Update user displayed")
// }

/* Handlers */
// var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	/* Create the token */
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	/* Create a map to store our claims
// 	claims := token.Claims.(jwt.MapClaims)

// 	/* Set token claims */
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["admin"] = true
// 	claims["name"] = "Ado Kukic"
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 	/* Sign the token with our secret */
// 	tokenString, _ := token.SignedString(mySigningKey)

// 	/* Finally, write the token to the browser window */
// 	w.Write([]byte(tokenString))
// })
