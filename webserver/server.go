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

	//assistant
	router.Get("/assistant", http.HandlerFunc(assistantPostHandler))
	router.Post("/assistant", http.HandlerFunc(assistantPostHandler))
	//alisa
	router.Get("/alisa", http.HandlerFunc(alisaPostHandler))
	router.Post("/alisa", http.HandlerFunc(alisaPostHandler))
	router.Get("/alisa/v1.0/user/devices", http.HandlerFunc(alisaGetDevicesHandler))
	router.Post("/alisa/v1.0/user/devices/action", http.HandlerFunc(alisaDevicesActionHandler))

	//security
	router.Get("/security/on", http.HandlerFunc(securityOnHandler))
	router.Get("/security/off", http.HandlerFunc(securityOffHandler))
	router.Get("/security/alarm/:zone", http.HandlerFunc(securityAlarmHandler))



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
