// generated by 'threeport-codegen api-model' - do not edit

package v0

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	v0 "github.com/threeport/threeport/pkg/api/v0"
	client "github.com/threeport/threeport/pkg/client"
	"net/http"
)

// GetDomainNameDefinitionByID feteches a domain name definition by ID
func GetDomainNameDefinitionByID(id uint, apiAddr, apiToken string) (*v0.DomainNameDefinition, error) {
	var domainNameDefinition v0.DomainNameDefinition

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/domain-name-definitions/%d", apiAddr, ApiVersion, id),
		apiToken,
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &domainNameDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &domainNameDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&domainNameDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API", err)
	}

	return &domainNameDefinition, nil
}

// GetDomainNameDefinitionByName feteches a domain name definition by name
func GetDomainNameDefinitionByName(name, apiAddr, apiToken string) (*v0.DomainNameDefinition, error) {
	var domainNameDefinitions []v0.DomainNameDefinition

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/domain-name-definitions?name=%s", apiAddr, ApiVersion, name),
		apiToken,
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &v0.DomainNameDefinition{}, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.DomainNameDefinition{}, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&domainNameDefinitions); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API", err)
	}

	switch {
	case len(domainNameDefinitions) < 1:
		return &v0.DomainNameDefinition{}, errors.New(fmt.Sprintf("no workload definitions with name %s", name))
	case len(domainNameDefinitions) > 1:
		return &v0.DomainNameDefinition{}, errors.New(fmt.Sprintf("more than one workload definition with name %s returned", name))
	}

	return &domainNameDefinitions[0], nil
}

// CreateDomainNameDefinition creates a new domain name definition
func CreateDomainNameDefinition(domainNameDefinition *v0.DomainNameDefinition, apiAddr, apiToken string) (*v0.DomainNameDefinition, error) {
	jsonDomainNameDefinition, err := client.MarshalObject(domainNameDefinition)
	if err != nil {
		return domainNameDefinition, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/domain-name-definitions", apiAddr, ApiVersion),
		apiToken,
		http.MethodPost,
		bytes.NewBuffer(jsonDomainNameDefinition),
		http.StatusCreated,
	)
	if err != nil {
		return domainNameDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return domainNameDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&domainNameDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API", err)
	}

	return domainNameDefinition, nil
}

// UpdateDomainNameDefinition updates a domain name definition
func UpdateDomainNameDefinition(domainNameDefinition *v0.DomainNameDefinition, apiAddr, apiToken string) (*v0.DomainNameDefinition, error) {
	// capture the object ID then remove it from the object since the API will not
	// allow an update the ID field
	domainNameDefinitionID := *domainNameDefinition.ID
	domainNameDefinition.ID = nil

	jsonDomainNameDefinition, err := client.MarshalObject(domainNameDefinition)
	if err != nil {
		return domainNameDefinition, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/domain-name-definitions/%d", apiAddr, ApiVersion, domainNameDefinitionID),
		apiToken,
		http.MethodPatch,
		bytes.NewBuffer(jsonDomainNameDefinition),
		http.StatusOK,
	)
	if err != nil {
		return domainNameDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return domainNameDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&domainNameDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API", err)
	}

	return domainNameDefinition, nil
}

// GetDomainNameInstanceByID feteches a domain name instance by ID
func GetDomainNameInstanceByID(id uint, apiAddr, apiToken string) (*v0.DomainNameInstance, error) {
	var domainNameInstance v0.DomainNameInstance

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/domain-name-instances/%d", apiAddr, ApiVersion, id),
		apiToken,
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &domainNameInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &domainNameInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&domainNameInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API", err)
	}

	return &domainNameInstance, nil
}

// GetDomainNameInstanceByName feteches a domain name instance by name
func GetDomainNameInstanceByName(name, apiAddr, apiToken string) (*v0.DomainNameInstance, error) {
	var domainNameInstances []v0.DomainNameInstance

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/domain-name-instances?name=%s", apiAddr, ApiVersion, name),
		apiToken,
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &v0.DomainNameInstance{}, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.DomainNameInstance{}, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&domainNameInstances); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API", err)
	}

	switch {
	case len(domainNameInstances) < 1:
		return &v0.DomainNameInstance{}, errors.New(fmt.Sprintf("no workload definitions with name %s", name))
	case len(domainNameInstances) > 1:
		return &v0.DomainNameInstance{}, errors.New(fmt.Sprintf("more than one workload definition with name %s returned", name))
	}

	return &domainNameInstances[0], nil
}

// CreateDomainNameInstance creates a new domain name instance
func CreateDomainNameInstance(domainNameInstance *v0.DomainNameInstance, apiAddr, apiToken string) (*v0.DomainNameInstance, error) {
	jsonDomainNameInstance, err := client.MarshalObject(domainNameInstance)
	if err != nil {
		return domainNameInstance, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/domain-name-instances", apiAddr, ApiVersion),
		apiToken,
		http.MethodPost,
		bytes.NewBuffer(jsonDomainNameInstance),
		http.StatusCreated,
	)
	if err != nil {
		return domainNameInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return domainNameInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&domainNameInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API", err)
	}

	return domainNameInstance, nil
}

// UpdateDomainNameInstance updates a domain name instance
func UpdateDomainNameInstance(domainNameInstance *v0.DomainNameInstance, apiAddr, apiToken string) (*v0.DomainNameInstance, error) {
	// capture the object ID then remove it from the object since the API will not
	// allow an update the ID field
	domainNameInstanceID := *domainNameInstance.ID
	domainNameInstance.ID = nil

	jsonDomainNameInstance, err := client.MarshalObject(domainNameInstance)
	if err != nil {
		return domainNameInstance, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/domain-name-instances/%d", apiAddr, ApiVersion, domainNameInstanceID),
		apiToken,
		http.MethodPatch,
		bytes.NewBuffer(jsonDomainNameInstance),
		http.StatusOK,
	)
	if err != nil {
		return domainNameInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return domainNameInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&domainNameInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API", err)
	}

	return domainNameInstance, nil
}
