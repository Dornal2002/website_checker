package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Data interface {
	Getdata()
	CreateData()
	CheckQuery()
}

type response struct {
	Web_name []string `json:"websites"`
}

var mp map[string]string

func Getdata(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(mp)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	mp = make(map[string]string)
	resp := response{}
	r.ParseForm()
	// fmt.Fprint(w, r.Form)

	err := json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, "Data posted Succesfully")
	for _, url := range resp.Web_name {
		go CheckStatus(url)
	}

}

func CheckStatus(url string) {
	for {
		res, err := http.Get(url)
		if err != nil {
			mp[url] = "DOWN"
		} else if res.StatusCode >= 200 && res.StatusCode < 300 {
			mp[url] = "UP"
		} else {
			mp[url] = "DOWN"
		}
		time.Sleep(1 * time.Minute)
	}
}

func CheckQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m := r.Form
	val := m["websites"][0]

	res, ok := mp[val]
	if ok != true {
		fmt.Fprintf(w, "%v : DOWN", val)
		return
	}
	fmt.Fprintf(w, "%v is :%v", val, res)

}
