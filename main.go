package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"MonsterIndia.com/role"
)

var tpl *template.Template

type person struct {
	First      string
	Last       string
	Subscribed bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func Service(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribed") == "on"

	p1 := person{
		f,
		l,
		s,
	}
	tpl.ExecuteTemplate(w, "index.html", p1)
}
func main() {
	// items, total, err := companies.FindCompaniesByCharacter()
	// if err != nil {
	// 	fmt.Println("failed Companies Function..: ", err)
	// }
	// fmt.Println(total)
	// for _, item := range items {
	// 	fmt.Println(item)
	// }
	http.HandleFunc("/", Service)
	// http.HandleFunc("/results", role.FindJobsByRoleSkill)
	http.HandleFunc("/results", role.FindJobsByKeywordAndLocation)
	// role.FindFrelanceJobs()
	// role.FindJobsByLocation()
	err := http.ListenAndServe(":3000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("====================================================================================================")

}
