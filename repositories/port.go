package repositories

import (
	"ports/models"

	"github.com/upper/db/v4"
)

type PortRepository struct {
	db db.Session
}

func NewPortRepository(database db.Session) PortRepository {
	return PortRepository{
		db: database,
	}
}

func (r PortRepository) SavePorts(ports []models.Port) error {
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
