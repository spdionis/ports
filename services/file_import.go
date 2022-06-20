package services

import (
	"ports/models"

	"github.com/bcicen/jstream"
)

type portRepository interface {
	SavePorts(ports []models.Port) error
}

type jsonParser interface {
	OpenStream(filename string) (chan *jstream.MetaValue, error)
	Parse(stream chan *jstream.MetaValue, limit int) ([]models.Port, bool, error)
}

type PortFileImportService struct {
	parser         jsonParser
	portRepository portRepository
	batchSize      int
}

func NewPortFileImportService(p jsonParser, r portRepository, batchSize int) PortFileImportService {
	return PortFileImportService{
		parser:         p,
		portRepository: r,
		batchSize:      batchSize,
	}
}

func (s PortFileImportService) ImportFile(filename string) error {
	stream, err := s.parser.OpenStream(filename)
	if err != nil {
		return err
	}

	for {
		ports, streamEOF, err := s.parser.Parse(stream, s.batchSize)
		if err != nil {
			return nil
		}

		err = s.portRepository.SavePorts(ports)
		if err != nil {
			return err
		}

		if streamEOF {
			break
		}
	}

	return nil
}
