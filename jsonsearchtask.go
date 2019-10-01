package main
//hello
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type manager struct {
	ID    string `json:"ID"`
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
}

type allManager []manager

var managers = allManager{
	{
		ID:    "1",
		Fname: "Harshit",
		Lname: "Mohan",
	},
	{
		ID:    "2",
		Fname: "Unnati",
		Lname: "Kala",
	},
	{
		ID:    "3",
		Fname: "Shivangi",
		Lname: "Varshney",
	},
	{
		ID:    "4",
		Fname: "Harshit",
		Lname: "Mohaan",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent manager
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the manager id and full name")
	}

	json.Unmarshal(reqBody, &newEvent)

	// Add the newly created manager to the array of managers
	managers = append(managers, newEvent)

	// Return the 201 created status code
	w.WriteHeader(http.StatusCreated)
	// Return the newly created manager
	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	eventFname := mux.Vars(r)["Fname"]

	// Get the details from an existing manager
	// Use the blank identifier to avoid creating a value that will not be used
	for _, singleEvent := range managers {
		if singleEvent.Fname == eventFname {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllManager(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(managers)
}

/*func updateEvent(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	eventID := mux.Vars(r)["id"]
	var updatedEvent manager
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the manager title and description only in order to update")
	}

	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range managers {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			managers = append(managers[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}*/

/*func deleteEvent(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	eventID := mux.Vars(r)["id"]

	// Get the details from an existing manager
	// Use the blank identifier to avoid creating a value that will not be used
	for i, singleEvent := range managers {
		if singleEvent.ID == eventID {
			managers = append(managers[:i], managers[i+1:]...)
			fmt.Fprintf(w, "The manager with ID %v has been deleted successfully", eventID)
		}
	}
}*/
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/createmanager", createEvent).Methods("POST")
	router.HandleFunc("/managers", getAllManager).Methods("GET")
	router.HandleFunc("/managers/{Fname}", getOneEvent).Methods("GET")

	log.Fatal(http.ListenAndServe(":8043", router))
}
