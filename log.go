package logentries

import (
	"encoding/json"
	"net/http"
)

type LogCreateRequest struct {
	Log LogCreateRequestFields `json:"log"`
}

type LogReadRequest struct {
	ID string
}

type LogUpdateRequest struct {
	ID  string                 `json:"id"`
	Log LogUpdateRequestFields `json:"log"`
}

type LogDeleteRequest struct {
	ID string `json:"id"`
}

type LogCreateResponse struct {
	Log `json:"log"`
}

type LogReadResponse struct {
	Log `json:"log"`
}

type LogUpdateResponse struct {
	Log Log `json:"log"`
}

type LogUpdateRequestWrapper struct {
	Log LogUpdateRequestFields `json:"logset"`
}

type LogUserData struct {
	LeAgentFilename string `json:"le_agent_filename"`
	LeAgentFollow   string `json:"le_agent_follow"`
}

type LogsetsInfo struct {
	ID    string `json:"id"`
	Links []Link `json:"links,omitempty"`
	Name  string `json:"name,omitempty"`
}
type Log struct {
	LogsetsInfo []LogsetsInfo `json:"logsets_info"`
	Name        string        `json:"name"`
	UserData    LogUserData   `json:"user_data,omitempty"`
	Tokens      []string      `json:"tokens,omitempty"`
	SourceType  string        `json:"source_type,omitempty"`
	TokenSeed   interface{}   `json:"token_seed,omitempty"`
	Structures  []interface{} `json:"structures,omitempty"`
	ID          string        `json:"id"`
}

type LogCreateRequestFields struct {
	Name        string        `json:"name"`
	Structures  []interface{} `json:"structures,omitempty"`
	UserData    LogUserData   `json:"user_data,omitempty"`
	SourceType  string        `json:"source_type,omitempty"`
	TokenSeed   interface{}   `json:"token_seed,omitempty"`
	LogsetsInfo []LogsetsInfo `json:"logsets_info,omitempty"`
}

type LogUpdateRequestFields struct {
	LogsetsInfo []LogsetsInfo `json:"logsets_info"`
	Name        string        `json:"name"`
	UserData    LogUserData   `json:"user_data"`
	Tokens      []string      `json:"tokens"`
	TokenSeed   interface{}   `json:"token_seed"`
	Structures  []interface{} `json:"structures"`
}

func (l *LogClient) Create(createRequest *LogCreateRequest) (LogCreateResponse, error) {
	url := l.getUrl(logentriesLogsResource)

	payload, err := json.Marshal(createRequest)
	if err != nil {
		return LogCreateResponse{}, err
	}

	resp, err := l.postLogentries(url, payload, http.StatusCreated)
	if err != nil {
		return LogCreateResponse{}, err
	}

	var log LogCreateResponse

	err = json.Unmarshal(resp, &log)
	if err != nil {
		return LogCreateResponse{}, err
	}

	return log, nil
}

func (l *LogClient) Read(readRequest *LogReadRequest) (LogReadResponse, error) {
	url := l.getUrl(logentriesLogsResource + readRequest.ID)

	resp, err := l.getLogentries(url, http.StatusOK)
	if err != nil {
		return LogReadResponse{}, err
	}
	var log LogReadResponse

	err = json.Unmarshal(resp, &log)
	if err != nil {
		return LogReadResponse{}, err
	}

	return log, nil
}

func (l *LogClient) Update(updateRequest *LogUpdateRequest) (LogUpdateResponse, error) {
	url := l.getUrl(logentriesLogsResource + updateRequest.ID)

	payload, err := json.Marshal(&LogUpdateRequestWrapper{Log: updateRequest.Log})
	if err != nil {
		return LogUpdateResponse{}, err
	}

	resp, err := l.putLogentries(url, payload, http.StatusOK)
	if err != nil {
		return LogUpdateResponse{}, err
	}

	var log LogUpdateResponse

	err = json.Unmarshal(resp, &log)
	if err != nil {
		return LogUpdateResponse{}, err
	}

	return log, nil

}
func (l *LogClient) Delete(deleteRequest *LogDeleteRequest) (bool, error) {
	url := l.getUrl(logentriesLogsResource + deleteRequest.ID)

	success, err := l.deleteLogentries(url, http.StatusNoContent)
	if err != nil {
		return false, err
	}

	return success, nil
}
