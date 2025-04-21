package router

import (
	"encoding/json"
	"fmt"
	"log"
	"main/data"
	"net/http"
	"strconv"
)

func PatchUpdateProduct() ServeHttp {
	return func(w http.ResponseWriter, r *http.Request) {
		datas := data.DB()
		sample := new(data.Product)
		idstr := r.PathValue("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			log.Print("Path Invalid Error")
			info := InvalidInfoT{
				Name:  "Parsing Error",
				Cause: err,
				Msg:   "Try With A valid number string",
			}
			InvalidInfo(w, r, info)
			return
		}
		// minus
		if id > 0 {
			id--
		}
		err = json.NewDecoder(r.Body).Decode(sample)
		// invalid marshaling
		if err != nil {
			log.Print("Path Marshaling Error")
			info := InvalidInfoT{
				Name:  "Marshaling Error",
				Cause: err,
				Msg:   "Do it Again Later",
			}
			InvalidInfo(w, r, info)
			return
		}
		// edit the datas
		old, err := datas.Edit(id, *sample)
		if err != nil {
			ServerError(w, r)
			return
		}
		oldbuf, err := json.MarshalIndent(old, "", " ")

		if err != nil {
			log.Print("Path Marshaling Old")
			info := InvalidInfoT{
				Name:  "Marshaling Error",
				Cause: err,
				Msg:   "Do it Again Later",
			}
			InvalidInfo(w, r, info)
			return
		}

		newbuf, err := json.MarshalIndent(datas[id], "", " ")
		if err != nil {
			log.Print("Path Marshaling New")
			info := InvalidInfoT{
				Name:  "Marshaling Error",
				Cause: err,
				Msg:   "Do it Again Later",
			}
			InvalidInfo(w, r, info)
			return
		}

		toret := fmt.Sprintf("{\n%q : %v, %q: %v\n}", "old", string(oldbuf), "new", string(newbuf))

		w.Write([]byte(toret))
	}
}
