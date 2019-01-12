package lib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
	"github.com/julienschmidt/httprouter"
	"time"
	"log"
)


/**
 * 从处理程序填充模型
 * param w http.ResponseWriter
 * param r *http.Request
 * model interface{}
 */
func PopulateModelFromHandler(w http.ResponseWriter, r *http.Request, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}

	if err := r.Body.Close(); err != nil {
		return err
	}

	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}


type fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)

func Logger(slef fn) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		slef(w, r, param)
		log.Printf("Done in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	}
}
