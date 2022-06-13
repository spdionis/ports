package ports

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bcicen/jstream"
)

type PortController struct {
	repo PortRepository
}

func NewController(repo PortRepository) PortController {
	return PortController{
		repo: repo,
	}
}

func (c PortController) UpdatePorts(w http.ResponseWriter, r *http.Request) {
	bodyRaw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ports := make(map[string]Port)
	err = json.Unmarshal(bodyRaw, &ports)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repo.SavePorts(ports)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c PortController) ImportPorts(w http.ResponseWriter, r *http.Request) {
	bodyRaw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestBody := make(map[string]string)
	err = json.Unmarshal(bodyRaw, &requestBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if requestBody["filename"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fh, err := os.Open(requestBody["filename"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ports := make([]Port, 0, 100)
	i := 0
	decoder := jstream.NewDecoder(fh, 1).EmitKV()
	for value := range decoder.Stream() {
		kv, err := value.Value.(jstream.KV)
		if err != true {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		value, err := kv.Value.(map[string]interface{})
		if err != true {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		port := jstreamToPort(kv, value)

		i++
		ports = append(ports, port)
		if i%1000 == 0 {
			err := c.repo.SavePortList(ports)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ports = make([]Port, 100)
		}
	}
}

//quick and dirty function to convert the jstream value to a Port model
//beware, some types of JSON files will make this panic
func jstreamToPort(kv jstream.KV, value map[string]interface{}) Port {
	port := Port{
		PortID: kv.Key,
	}

	if value["name"] != nil {
		port.Name = value["name"].(string)
	}
	if value["city"] != nil {
		port.City = value["city"].(string)
	}
	if value["country"] != nil {
		port.Country = value["country"].(string)
	}
	if value["province"] != nil {
		port.Province = value["province"].(string)
	}
	if value["timezone"] != nil {
		port.Timezone = value["timezone"].(string)
	}
	if value["code"] != nil {
		port.Code = value["code"].(string)
	}

	if value["alias"] != nil {
		port.Alias = make([]string, 0)
		for _, v := range value["alias"].([]interface{}) {
			port.Alias = append(port.Alias, v.(string))
		}
	}

	if value["regions"] != nil {
		port.Regions = make([]string, 0)
		for _, v := range value["regions"].([]interface{}) {
			port.Regions = append(port.Regions, v.(string))
		}
	}

	if value["coordinates"] != nil {
		port.Coordinates = make([]float64, 0)
		for _, v := range value["coordinates"].([]interface{}) {
			port.Coordinates = append(port.Coordinates, v.(float64))
		}
	}

	if value["unlocs"] != nil {
		port.Unlocs = make([]string, 0)
		for _, v := range value["unlocs"].([]interface{}) {
			port.Unlocs = append(port.Unlocs, v.(string))
		}
	}
	return port
}
