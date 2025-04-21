package router

import (
	"encoding/json"
	"errors"
	"log"
	"main/data"
	"net/http"
	"strconv"
)

func DeleteProduct() ServeHttp {

	return func(w http.ResponseWriter, r *http.Request) {
		datas := data.DB()
		idstr := r.PathValue("id")
		var (
			id  int = 0
			err error
		)

		if id, err = strconv.Atoi(idstr); err != nil || len(idstr) <= 0 {
			if err == nil {
				err = errors.New("No Id Provided")
			}

			info := InvalidInfoT{
				Name:  "No Id Provided",
				Cause: errors.New("No Id Provided"),
				Msg:   "Try Again Later",
			}
			log.Printf("No Id Provided")
			InvalidInfo(w, r, info)
			return
		}

		sample := datas.Each(func(p *data.Product) bool {
			if p.Id == id {
				return false
			}
			return true
		})

		single := datas.Each(func(p *data.Product) bool {
			if p.Id == id {
				return true
			}
			return false
		})

		if len(single) <= 0 {
			w.WriteHeader(202)
			w.Write([]byte("{}"))
			return
		}

		buf, err := json.Marshal(datas[id-1])
		datas = sample
		err = datas.Save()
		if err != nil {
			ServerError(w, r)
		}

		if err != nil {

			info := InvalidInfoT{
				Name:  "Error On Marshaling",
				Cause: err,
				Msg:   "Try Again Later",
			}
			log.Printf("Error Marshaling")
			InvalidInfo(w, r, info)
			return
		}

		w.WriteHeader(202)
		w.Write(buf)
	}

}
