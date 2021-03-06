package logentries

import (
	"encoding/json"
	"net/http"
)

type LogSetCreateResponse struct {
	LogSet `json:"logset"`
}

type LogSetReadResponse struct {
	LogSet `json:"logset"`
}

type LogSetUpdateResponse struct {
	LogSet `json:"logset"`
}

type LogSetCreateRequest struct {
	LogSet LogSetFields `json:"logset"`
}

type LogSetReadRequest struct {
	ID string
}

type LogSetUpdateRequest struct {
	ID     string       `json:"id"`
	LogSet LogSetFields `json:"logset"`
}

type LogSetDeleteRequest struct {
	ID string
}

type LogSetUpdateRequestWrapper struct {
	LogSet LogSetFields `json:"logset"`
}

type LogSetFields struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	UserData    struct{}  `json:"user_data,omitempty"`
	LogsInfo    []LogInfo `json:"logs_info,omitempty"`
}

type LogSet struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name"`
	Description interface{} `json:"description,omitempty"`
	UserData    UserData    `json:"user_data"`
	LogsInfo    []LogInfo   `json:"logs_info"`
}

type LogInfo struct {
	ID    string `json:"id"`
	Links []Link `json:"links"`
	Name  string `json:"name"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type UserData struct {
	LeDistname string `json:"le_distname"`
	LeDistver  string `json:"le_distver"`
	LeNameintr string `json:"le_nameintr"`
}

func (l *LogSetClient) Create(createRequest *LogSetCreateRequest) (LogSetCreateResponse, error) {
	url := l.getUrl(logentriesLogsetsResource)

	payload, err := json.Marshal(createRequest)
	if err != nil {
		return LogSetCreateResponse{}, err
	}

	resp, err := l.postLogentries(url, payload, http.StatusCreated)
	if err != nil {
		return LogSetCreateResponse{}, err
	}

	var logset LogSetCreateResponse

	err = json.Unmarshal(resp, &logset)
	if err != nil {
		return LogSetCreateResponse{}, err
	}

	return logset, nil
}

func (l *LogSetClient) Read(readRequest *LogSetReadRequest) (LogSetReadResponse, error) {
	url := l.getUrl(logentriesLogsetsResource + readRequest.ID)

	resp, err := l.getLogentries(url, http.StatusOK)
	if err != nil {
		return LogSetReadResponse{}, err
	}

	var logset LogSetReadResponse
	err = json.Unmarshal(resp, &logset)
	if err != nil {
		return LogSetReadResponse{}, err
	}

	return logset, nil
}

func (l *LogSetClient) Update(updateRequest *LogSetUpdateRequest) (LogSetUpdateResponse, error) {
	url := l.getUrl(logentriesLogsetsResource + updateRequest.ID)

	payload, err := json.Marshal(&LogSetUpdateRequestWrapper{LogSet: updateRequest.LogSet})
	if err != nil {
		return LogSetUpdateResponse{}, err
	}

	resp, err := l.putLogentries(url, payload, http.StatusOK)
	if err != nil {
		return LogSetUpdateResponse{}, err
	}

	var logset LogSetUpdateResponse

	err = json.Unmarshal(resp, &logset)
	if err != nil {
		return LogSetUpdateResponse{}, err
	}

	return logset, nil

}
func (l *LogSetClient) Delete(deleteRequest *LogSetDeleteRequest) (bool, error) {
	url := l.getUrl(logentriesLogsetsResource + deleteRequest.ID)

	success, err := l.deleteLogentries(url, http.StatusNoContent)
	if err != nil {
		return false, err
	}

	return success, nil
}
