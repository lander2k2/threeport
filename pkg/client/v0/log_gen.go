// generated by 'threeport-codegen api-model' - do not edit

package v0

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	util "github.com/threeport/threeport/internal/util"
	v0 "github.com/threeport/threeport/pkg/api/v0"
	"net/http"
)

// GetLogBackends fetches all log backends.
// TODO: implement pagination
func GetLogBackends(apiClient *http.Client, apiAddr string) (*[]v0.LogBackend, error) {
	var logBackends []v0.LogBackend

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-backends", apiAddr, ApiVersion),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logBackends, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &logBackends, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logBackends); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logBackends, nil
}

// GetLogBackendByID fetches a log backend by ID.
func GetLogBackendByID(apiClient *http.Client, apiAddr string, id uint) (*v0.LogBackend, error) {
	var logBackend v0.LogBackend

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-backends/%d", apiAddr, ApiVersion, id),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logBackend, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &logBackend, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logBackend); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logBackend, nil
}

// GetLogBackendByName fetches a log backend by name.
func GetLogBackendByName(apiClient *http.Client, apiAddr, name string) (*v0.LogBackend, error) {
	var logBackends []v0.LogBackend

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-backends?name=%s", apiAddr, ApiVersion, name),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &v0.LogBackend{}, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.LogBackend{}, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logBackends); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	switch {
	case len(logBackends) < 1:
		return &v0.LogBackend{}, errors.New(fmt.Sprintf("no workload definitions with name %s", name))
	case len(logBackends) > 1:
		return &v0.LogBackend{}, errors.New(fmt.Sprintf("more than one workload definition with name %s returned", name))
	}

	return &logBackends[0], nil
}

// CreateLogBackend creates a new log backend.
func CreateLogBackend(apiClient *http.Client, apiAddr string, logBackend *v0.LogBackend) (*v0.LogBackend, error) {
	jsonLogBackend, err := util.MarshalObject(logBackend)
	if err != nil {
		return logBackend, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-backends", apiAddr, ApiVersion),
		http.MethodPost,
		bytes.NewBuffer(jsonLogBackend),
		http.StatusCreated,
	)
	if err != nil {
		return logBackend, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return logBackend, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logBackend); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return logBackend, nil
}

// UpdateLogBackend updates a log backend.
func UpdateLogBackend(apiClient *http.Client, apiAddr string, logBackend *v0.LogBackend) (*v0.LogBackend, error) {
	// capture the object ID then remove fields that cannot be updated in the API
	logBackendID := *logBackend.ID
	logBackend.ID = nil
	logBackend.CreatedAt = nil
	logBackend.UpdatedAt = nil

	jsonLogBackend, err := util.MarshalObject(logBackend)
	if err != nil {
		return logBackend, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-backends/%d", apiAddr, ApiVersion, logBackendID),
		http.MethodPatch,
		bytes.NewBuffer(jsonLogBackend),
		http.StatusOK,
	)
	if err != nil {
		return logBackend, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return logBackend, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logBackend); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return logBackend, nil
}

// DeleteLogBackend deletes a log backend by ID.
func DeleteLogBackend(apiClient *http.Client, apiAddr string, id uint) (*v0.LogBackend, error) {
	var logBackend v0.LogBackend

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-backends/%d", apiAddr, ApiVersion, id),
		http.MethodDelete,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logBackend, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &logBackend, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logBackend); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logBackend, nil
}

// GetLogStorageDefinitions fetches all log storage definitions.
// TODO: implement pagination
func GetLogStorageDefinitions(apiClient *http.Client, apiAddr string) (*[]v0.LogStorageDefinition, error) {
	var logStorageDefinitions []v0.LogStorageDefinition

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-definitions", apiAddr, ApiVersion),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logStorageDefinitions, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &logStorageDefinitions, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageDefinitions); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logStorageDefinitions, nil
}

// GetLogStorageDefinitionByID fetches a log storage definition by ID.
func GetLogStorageDefinitionByID(apiClient *http.Client, apiAddr string, id uint) (*v0.LogStorageDefinition, error) {
	var logStorageDefinition v0.LogStorageDefinition

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-definitions/%d", apiAddr, ApiVersion, id),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logStorageDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &logStorageDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logStorageDefinition, nil
}

// GetLogStorageDefinitionByName fetches a log storage definition by name.
func GetLogStorageDefinitionByName(apiClient *http.Client, apiAddr, name string) (*v0.LogStorageDefinition, error) {
	var logStorageDefinitions []v0.LogStorageDefinition

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-definitions?name=%s", apiAddr, ApiVersion, name),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &v0.LogStorageDefinition{}, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.LogStorageDefinition{}, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageDefinitions); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	switch {
	case len(logStorageDefinitions) < 1:
		return &v0.LogStorageDefinition{}, errors.New(fmt.Sprintf("no workload definitions with name %s", name))
	case len(logStorageDefinitions) > 1:
		return &v0.LogStorageDefinition{}, errors.New(fmt.Sprintf("more than one workload definition with name %s returned", name))
	}

	return &logStorageDefinitions[0], nil
}

// CreateLogStorageDefinition creates a new log storage definition.
func CreateLogStorageDefinition(apiClient *http.Client, apiAddr string, logStorageDefinition *v0.LogStorageDefinition) (*v0.LogStorageDefinition, error) {
	jsonLogStorageDefinition, err := util.MarshalObject(logStorageDefinition)
	if err != nil {
		return logStorageDefinition, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-definitions", apiAddr, ApiVersion),
		http.MethodPost,
		bytes.NewBuffer(jsonLogStorageDefinition),
		http.StatusCreated,
	)
	if err != nil {
		return logStorageDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return logStorageDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return logStorageDefinition, nil
}

// UpdateLogStorageDefinition updates a log storage definition.
func UpdateLogStorageDefinition(apiClient *http.Client, apiAddr string, logStorageDefinition *v0.LogStorageDefinition) (*v0.LogStorageDefinition, error) {
	// capture the object ID then remove fields that cannot be updated in the API
	logStorageDefinitionID := *logStorageDefinition.ID
	logStorageDefinition.ID = nil
	logStorageDefinition.CreatedAt = nil
	logStorageDefinition.UpdatedAt = nil

	jsonLogStorageDefinition, err := util.MarshalObject(logStorageDefinition)
	if err != nil {
		return logStorageDefinition, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-definitions/%d", apiAddr, ApiVersion, logStorageDefinitionID),
		http.MethodPatch,
		bytes.NewBuffer(jsonLogStorageDefinition),
		http.StatusOK,
	)
	if err != nil {
		return logStorageDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return logStorageDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return logStorageDefinition, nil
}

// DeleteLogStorageDefinition deletes a log storage definition by ID.
func DeleteLogStorageDefinition(apiClient *http.Client, apiAddr string, id uint) (*v0.LogStorageDefinition, error) {
	var logStorageDefinition v0.LogStorageDefinition

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-definitions/%d", apiAddr, ApiVersion, id),
		http.MethodDelete,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logStorageDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &logStorageDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logStorageDefinition, nil
}

// GetLogStorageInstances fetches all log storage instances.
// TODO: implement pagination
func GetLogStorageInstances(apiClient *http.Client, apiAddr string) (*[]v0.LogStorageInstance, error) {
	var logStorageInstances []v0.LogStorageInstance

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-instances", apiAddr, ApiVersion),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logStorageInstances, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &logStorageInstances, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageInstances); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logStorageInstances, nil
}

// GetLogStorageInstanceByID fetches a log storage instance by ID.
func GetLogStorageInstanceByID(apiClient *http.Client, apiAddr string, id uint) (*v0.LogStorageInstance, error) {
	var logStorageInstance v0.LogStorageInstance

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-instances/%d", apiAddr, ApiVersion, id),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logStorageInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &logStorageInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logStorageInstance, nil
}

// GetLogStorageInstanceByName fetches a log storage instance by name.
func GetLogStorageInstanceByName(apiClient *http.Client, apiAddr, name string) (*v0.LogStorageInstance, error) {
	var logStorageInstances []v0.LogStorageInstance

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-instances?name=%s", apiAddr, ApiVersion, name),
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &v0.LogStorageInstance{}, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.LogStorageInstance{}, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageInstances); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	switch {
	case len(logStorageInstances) < 1:
		return &v0.LogStorageInstance{}, errors.New(fmt.Sprintf("no workload definitions with name %s", name))
	case len(logStorageInstances) > 1:
		return &v0.LogStorageInstance{}, errors.New(fmt.Sprintf("more than one workload definition with name %s returned", name))
	}

	return &logStorageInstances[0], nil
}

// CreateLogStorageInstance creates a new log storage instance.
func CreateLogStorageInstance(apiClient *http.Client, apiAddr string, logStorageInstance *v0.LogStorageInstance) (*v0.LogStorageInstance, error) {
	jsonLogStorageInstance, err := util.MarshalObject(logStorageInstance)
	if err != nil {
		return logStorageInstance, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-instances", apiAddr, ApiVersion),
		http.MethodPost,
		bytes.NewBuffer(jsonLogStorageInstance),
		http.StatusCreated,
	)
	if err != nil {
		return logStorageInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return logStorageInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return logStorageInstance, nil
}

// UpdateLogStorageInstance updates a log storage instance.
func UpdateLogStorageInstance(apiClient *http.Client, apiAddr string, logStorageInstance *v0.LogStorageInstance) (*v0.LogStorageInstance, error) {
	// capture the object ID then remove fields that cannot be updated in the API
	logStorageInstanceID := *logStorageInstance.ID
	logStorageInstance.ID = nil
	logStorageInstance.CreatedAt = nil
	logStorageInstance.UpdatedAt = nil

	jsonLogStorageInstance, err := util.MarshalObject(logStorageInstance)
	if err != nil {
		return logStorageInstance, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-instances/%d", apiAddr, ApiVersion, logStorageInstanceID),
		http.MethodPatch,
		bytes.NewBuffer(jsonLogStorageInstance),
		http.StatusOK,
	)
	if err != nil {
		return logStorageInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return logStorageInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return logStorageInstance, nil
}

// DeleteLogStorageInstance deletes a log storage instance by ID.
func DeleteLogStorageInstance(apiClient *http.Client, apiAddr string, id uint) (*v0.LogStorageInstance, error) {
	var logStorageInstance v0.LogStorageInstance

	response, err := GetResponse(
		apiClient,
		fmt.Sprintf("%s/%s/log-storage-instances/%d", apiAddr, ApiVersion, id),
		http.MethodDelete,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &logStorageInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &logStorageInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&logStorageInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &logStorageInstance, nil
}
