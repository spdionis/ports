package services

import (
	"errors"
	"testing"

	"ports/models"

	"github.com/bcicen/jstream"
)

//THIS IS A VERY VERY DUMB WAY TO WRITE TESTS :D
//Leaving the implementation of a testing suite with something like stretchr/testify as an exercise to the reader
type fileImportTestSuite struct {
	savePorts  int
	openStream int
	parse      int
}

//yes, these are global
var suite fileImportTestSuite
var testBatchSize = 100
var testFilename = "test.json"

type portRepositoryMock struct{}

func (portRepositoryMock) SavePorts(ports []models.Port) error {
	suite.savePorts++
	return nil
}

type jsonParserMock struct{}

func (j jsonParserMock) OpenStream(filename string) (chan *jstream.MetaValue, error) {
	if filename != testFilename {
		return nil, errors.New("test failed")
	}

	suite.openStream++
	return nil, nil
}

func (j jsonParserMock) Parse(stream chan *jstream.MetaValue, limit int) ([]models.Port, bool, error) {
	suite.parse++
	if limit != testBatchSize {
		return nil, false, errors.New("test failed")
	}

	if suite.parse == 2 {
		return nil, true, nil
	}

	return nil, false, nil
}

func TestImportFile(t *testing.T) {
	service := PortFileImportService{
		parser:         jsonParserMock{},
		portRepository: portRepositoryMock{},
		batchSize:      testBatchSize,
	}

	err := service.ImportFile("test.json")
	if err != nil {
		t.FailNow()
	}

	if suite.openStream != 1 || suite.parse != 2 || suite.savePorts != 2 {
		t.FailNow()
	}
}
