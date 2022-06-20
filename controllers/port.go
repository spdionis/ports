package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"ports/models"
	"ports/repositories"
	"ports/services"
)

type PortController struct {
	repo              repositories.PortRepository
	fileImportService services.PortFileImportService
}

func NewController(repo repositories.PortRepository, fileImportService services.PortFileImportService) PortController {
	return PortController{
		repo:              repo,
		fileImportService: fileImportService,
	}
}

type updatePortsRequest map[string]models.Port

type importPortsRequest struct {
	Filename string `json:"filename"`
}

func (c PortController) UpdatePorts(w http.ResponseWriter, r *http.Request) {
	bodyRaw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ports := make(updatePortsRequest)
	err = json.Unmarshal(bodyRaw, &ports)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//could make the save call go through a service layer as well, depending on the style of the project
	//I considered it unnecessary complexity in our case
	err = c.repo.SavePorts(updatePortsRequestToPortList(ports))
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

	var request importPortsRequest
	err = json.Unmarshal(bodyRaw, &request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.fileImportService.ImportFile(request.Filename)
	if err != nil {
		log.Println(err)
		//this is not a good assumption, could also be a 5xx error
		//leaving correct application error handling as an exercise for the reader
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updatePortsRequestToPortList(ports updatePortsRequest) []models.Port {
	sl := make([]models.Port, 0, len(ports))
	for portID, port := range ports {
		port.PortID = portID
		sl = append(sl, port)
	}

	return sl
}
