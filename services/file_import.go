package services

import (
	"ports/repositories"
)

type PortFileImportService struct {
	parser         StreamingPortJSONParser
	portRepository repositories.PortRepository
	batchSize      int
}

func NewPortFileImportService(p StreamingPortJSONParser, r repositories.PortRepository, batchSize int) PortFileImportService {
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
