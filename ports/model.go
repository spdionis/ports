package ports

import "github.com/upper/db/v4"

type Port struct {
	ID          int       `db:"id"`
	PortID      string    `db:"port_id"`
	Name        string    `json:"name" db:"name"`
	City        string    `json:"city" db:"city"`
	Country     string    `json:"country" db:"country"`
	Alias       []string  `json:"alias" db:"alias"`
	Regions     []string  `json:"regions" db:"regions"`
	Coordinates []float64 `json:"coordinates" db:"coordinates"`
	Province    string    `json:"province" db:"province"`
	Timezone    string    `json:"timezone" db:"timezone"`
	Unlocs      []string  `json:"unlocs" db:"unlocs"`
	Code        string    `json:"code" db:"code"`
}

type PortRepository struct {
	db db.Session
}

func NewRepository(database db.Session) PortRepository {
	return PortRepository{
		db: database,
	}
}

func (r *PortRepository) SavePorts(ports map[string]Port) error {
	return r.SavePortList(flattenPorts(ports))
}

func (r *PortRepository) SavePortList(ports []Port) error {
	query := r.db.SQL().
		InsertInto("ports").
		Columns(
			"port_id",
			"name",
			"city",
			"country",
			"alias",
			"regions",
			"coordinates",
			"province",
			"timezone",
			"unlocs",
			"code",
		)

	for _, port := range ports {
		query = query.
			Values(
				port.PortID,
				port.Name,
				port.City,
				port.Country,
				port.Alias,
				port.Regions,
				port.Coordinates,
				port.Province,
				port.Timezone,
				port.Unlocs,
				port.Code,
			)
	}
	_, err := query.Exec()

	if err != nil {
		return err
	}

	return nil
}

func flattenPorts(ports map[string]Port) []Port {
	sl := make([]Port, 0, len(ports))
	for portID, port := range ports {
		port.PortID = portID
		sl = append(sl, port)
	}

	return sl
}
