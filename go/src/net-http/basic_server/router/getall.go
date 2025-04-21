package router

import (
	"errors"
	"fmt"
	"log"
	"main/data"
	"net/http"
	"strings"
)

type ServeHttp func(w http.ResponseWriter, r *http.Request)

func GetAll(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	addJsonContentType(header)
	w.Write(data.ReadData())
	r.Body.Close()
}

func GetSpecific() ServeHttp {
	return func(w http.ResponseWriter, r *http.Request) {
		datas := data.DB()
		queryParams := r.URL.Query()
		r.Body.Close()
		// get the params
		switch {
		case len(queryParams.Get("id")) > 0:
			GetDataById(w, r, &datas, queryParams)
			return

		case len(queryParams.Get("names")) > 0:
			names := strings.Split(queryParams.Get("names"), ",")
			GetByName(names, datas)(w, r)
			return

		case len(queryParams.Get("type")) > 0:
			types := strings.Split(queryParams.Get("type"), ",")
			GetByType(types, datas)(w, r)
			return

		case len(queryParams.Get("tags")) > 0:
			tags := strings.Split(queryParams.Get("tags"), ",")
			log.Printf("The tags %v", tags)
			GetByTags(tags, datas)(w, r)
			return
		default:

			info := InvalidInfoT{
				Name:  fmt.Sprintf("No Data Found"),
				Cause: errors.New("Something Is Wrong"),
				Msg:   "Retry with the valid data",
			}

			InvalidInfo(w, r, info)
			return
		}

	}
}
