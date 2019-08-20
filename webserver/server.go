package webserver

import (
	"../types"
	"./handlers"
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
)

func Init(myConfig types.MyConfig) {

	router := pat.New()
	router.Get("/", http.HandlerFunc(indexHandler))
	router.Get("/sensors", http.HandlerFunc(getSensorHandler))
	router.Get("/sensor/:id/type/:type", http.HandlerFunc(sensorGetDataHandler))
	router.Post("/sensor/upload", http.HandlerFunc(sensorUploadHandler))

	router.Post("/sensor/:id/type/:type", http.HandlerFunc(sensorPostDataHandler))

	fmt.Println("====================================================")
	fmt.Println("ListenAndServe : " + myConfig.Env.HttpPort)

	http.Handle("/", router)
	//http.Handle("/images/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	err := http.ListenAndServe(":"+myConfig.Env.HttpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func indexHandlerOk(w http.ResponseWriter, r *http.Request) {
	handlers.HandlerSuccess(w, "ok")
}

func optionsHandler(w http.ResponseWriter, _ *http.Request) {
	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.WriteHeader(200)
}
