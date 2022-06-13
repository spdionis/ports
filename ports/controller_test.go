package ports

import (
	"strings"
	"testing"

	"github.com/bcicen/jstream"
)

const portsJSONTest = `
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }
}`

func TestJstreamToPort(t *testing.T) {
	decoder := jstream.NewDecoder(strings.NewReader(portsJSONTest), 1).EmitKV()

	ports := make([]Port, 0)
	for value := range decoder.Stream() {
		kv, _ := value.Value.(jstream.KV)
		ports = append(ports, jstreamToPort(kv, kv.Value.(map[string]interface{})))
	}

	if len(ports) != 2 {
		t.Fail()
	}
}
