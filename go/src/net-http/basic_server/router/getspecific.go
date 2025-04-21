package router

import (
	"encoding/json"
	"fmt"
	"log"
	"main/data"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetDataById(w http.ResponseWriter, r *http.Request, datas *data.Datas, queryParams url.Values) {
	// like me
	single := datas.Each(func(p *data.Product) bool {
		id, err := strconv.Atoi(queryParams.Get("id"))
		log.Printf("has ID of %d", id)
		if err != nil {
			return false
		} else if id == p.Id {
			log.Printf("FoundMatching %d", id)
			return true
		}
		return false
	})

	if len(single) > 1 {
		log.Print("Single len is ")
		buf, err := json.Marshal(single)
		if err != nil {
			info := InvalidInfoT{
				Name:  fmt.Sprintf("Something Not Found"),
				Cause: err,
				Msg:   "Retry with the valid data",
			}

			InvalidInfo(w, r, info)
			return
		} else {
			w.Write(buf)
			return
		}
	} else if len(single) > 0 {
		buf, err := json.Marshal(single[0])
		log.Print("Single len is 0")
		if err != nil {
			info := InvalidInfoT{
				Name:  fmt.Sprintf("Something Not Found"),
				Cause: err,
				Msg:   "Retry with the valid data",
			}
			InvalidInfo(w, r, info)
			return
		} else {
			w.Write(buf)
			return
		}
	} else {
		w.Write([]byte("[]"))
	}

}

func GetByName(names []string, datas data.Datas) ServeHttp {
	return func(w http.ResponseWriter, r *http.Request) {
		sample := datas.Each(func(p *data.Product) bool {
			infonames := strings.Fields(p.Info.Name)
			for _, infoname := range infonames {
				for _, name := range names {
					if strings.Compare(strings.ToLower(name), strings.ToLower(infoname)) == 0 {
						log.Printf("A matching name is found")
						return true
					}

					log.Printf("A matching name is not found %q", strings.ToLower(infoname))

				}
			}
			return false
		})

		if len(sample) > 0 {
			var buf []byte
			var err error
			if len(sample) > 1 {
				buf, err = json.Marshal(sample)
			} else if len(sample) > 0 {
				buf, err = json.Marshal(sample[0])
			} else {
				w.Write([]byte("[]"))
				return
			}

			// error happend
			if err != nil {
				info := InvalidInfoT{
					Name:  "Fail Marshaling",
					Cause: err,
					Msg:   "Try Again Later",
				}

				InvalidInfo(w, r, info)
				return
			}
			// it is sucessful
			if buf != nil {
				w.Write(buf)
			} else {
				ServerError(w, r)
			}
			return
		}
	}
}

// Processing type

func GetByType(types []string, datas data.Datas) ServeHttp {

	return func(w http.ResponseWriter, r *http.Request) {
		sample := datas.Each(func(p *data.Product) bool {
			for _, t := range types {
				if strings.Compare(strings.ToLower(t),
					strings.ToLower(p.Info.Type)) == 0 {
					log.Printf("GetType Match: %v", t)
					return true
				}

				log.Printf("GetType: %v", t)
			}

			return false
		})

		var buf []byte
		var err error

		switch {
		case len(sample) > 1:
			buf, err = json.Marshal(sample)
			break
		case len(sample) > 0:
			buf, err = json.Marshal(sample[0])
			break

		default:
			w.Write([]byte("[]"))
			return
		}

		// process the data

		if err != nil {
			info := InvalidInfoT{
				Name:  "Fail Marshaling",
				Cause: err,
				Msg:   "Try Again Later",
			}

			InvalidInfo(w, r, info)
			return
		}
		// it is sucessful
		if buf != nil {
			w.Write(buf)
		} else {
			ServerError(w, r)
		}
		return

	}

}

// processing tags

func GetByTags(input []string, datas data.Datas) ServeHttp {

	return func(w http.ResponseWriter, r *http.Request) {
		sample := datas.Each(func(p *data.Product) bool {
			for _, tag := range p.Tags {
				for _, t := range input {
					if strings.Compare(strings.ToLower(t),
						strings.ToLower(tag)) == 0 {
						log.Printf("GetTags Match: %v", t)
						return true
					}

					log.Printf("GetTags: %v", t)
				}
			}

			return false
		})

		var buf []byte
		var err error

		switch {
		case len(sample) > 1:
			buf, err = json.Marshal(sample)
		case len(sample) > 0:
			buf, err = json.Marshal(sample[0])
			goto forward

		default:
			w.Write([]byte("[]"))
			return
		}

	forward:

		// process the data

		if err != nil {
			info := InvalidInfoT{
				Name:  "Fail Marshaling",
				Cause: err,
				Msg:   "Try Again Later",
			}

			InvalidInfo(w, r, info)
			return
		}
		// it is sucessful
		if buf != nil {
			w.Write(buf)
		} else {
			ServerError(w, r)
		}
		return

	}
}
