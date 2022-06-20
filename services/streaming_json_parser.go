package services

import (
	"errors"
	"os"

	"ports/models"

	"github.com/bcicen/jstream"
)

type StreamingPortJSONParser struct{}

func (p StreamingPortJSONParser) OpenStream(filename string) (chan *jstream.MetaValue, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return jstream.NewDecoder(fh, 1).EmitKV().Stream(), nil
}

func (p StreamingPortJSONParser) Parse(stream chan *jstream.MetaValue, limit int) ([]models.Port, bool, error) {
	ports := make([]models.Port, 0, limit)
	for i := 0; i < limit; i++ {
		value, ok := <-stream
		if !ok {
			return ports, true, nil
		}

		kv, ok := value.Value.(jstream.KV)
		if ok != true {
			return ports, false, errors.New("invalid json")
		}

		raw, err := kv.Value.(map[string]interface{})
		if err != true {
			return ports, false, errors.New("invalid json")
		}

		port := streamValueToPort(raw)
		port.PortID = kv.Key
		ports = append(ports, port)
	}

	return ports, false, nil
}

//quick and dirty function to convert the jstream value to a Port model
//beware, some types of JSON files will make this panic
func streamValueToPort(value map[string]interface{}) models.Port {
	port := models.Port{}

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
