package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	HTML = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>GO-DOJO</title><script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script><link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic"><link rel="stylesheet" href="//cdn.rawgit.com/necolas/normalize.css/master/normalize.css"><link rel="stylesheet" href="//cdn.rawgit.com/milligram/milligram/master/dist/milligram.min.css"><style type="text/css">.order_btn {width: 100%;}</style></head><body><div class="container"><div class="row"><h2>Order Coffee</h2></div><div class="row"><input id="customer_name" type="text" name="customer" placeholder="put your name.."/></div><div class="row"><button class="order_btn" type="button" name="name" value="Espresso">Espresso</button></div><div class="row"><button class="order_btn" type="button" name="name" value="Cappuccino">Cappuccino</button></div><div class="row"><button class="order_btn" type="button" name="name" value="Latte">Latte</button></div><div class="row"><p id="result"></p></div></div><script>$(function(){$('.order_btn').click(function(){$.ajax({type: "POST",url: "/",contentType: "application/json; charset=utf-8",data: JSON.stringify({"name": this.value,"customer": $('#customer_name').val()}),success: function(result){ console.log(result); $("#result").text(result); },error: function(xhr, resp, text){console.log(xhr, resp, text);$("#result").text(resp);}});});})</script></body></html>`
)

type OrderTransfer struct {
	Name     string `json:"name"`
	Customer string `json:"customer"`
}

type Handler struct {
	Queue IQueue
}

func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// return HTML page for GET requests
	if req.Method == http.MethodGet {
		rw.Write([]byte(HTML))
		return
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println("ERROR: cannot get request body")
		rw.Write([]byte(err.Error()))
		return
	}

	var orderRequested OrderTransfer

	err = json.Unmarshal(body, &orderRequested)

	if err != nil {
		fmt.Printf("ERROR: cannot parse request body: %#v Error details: %#v\n", body, err)
		rw.Write([]byte(err.Error()))
		return
	}

	err = h.Queue.Put(
		NewOrder(orderRequested.Name, orderRequested.Customer),
	)

	if err != nil {
		fmt.Printf("ERROR: cannot add %#v to queue because of: %#v\n", orderRequested, err)
		rw.Write([]byte(err.Error()))
		return
	}

	fmt.Printf("Added order for %s from %s\n", orderRequested.Name, orderRequested.Customer)
	rw.Write([]byte("OK"))
	return
}

func StartServer(queue IQueue) {
	mux := http.NewServeMux()
	mux.Handle(
		"/",
		Handler{
			Queue: queue,
		},
	)

	err := http.ListenAndServe(":7777", mux)

	if err != nil {
		fmt.Printf("%#v", err)
	}
}
