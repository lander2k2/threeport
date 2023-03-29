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

// GetClusterDefinitionByID feteches a cluster definition by ID
func GetClusterDefinitionByID(id uint, apiAddr, apiToken string) (*v0.ClusterDefinition, error) {
	var clusterDefinition v0.ClusterDefinition

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/cluster_definitions/%d", apiAddr, ApiVersion, id),
		apiToken,
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &clusterDefinition, err
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &clusterDefinition, err
	}

	if err = json.Unmarshal(jsonData, &clusterDefinition); err != nil {
		return &clusterDefinition, err
	}

	return &clusterDefinition, nil
}

// GetClusterDefinitionByName feteches a cluster definition by name
func GetClusterDefinitionByName(name, apiAddr, apiToken string) (*v0.ClusterDefinition, error) {
	var clusterDefinitions []v0.ClusterDefinition

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/cluster_definitions?name=%s", apiAddr, ApiVersion, name),
		apiToken,
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &v0.ClusterDefinition{}, err
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.ClusterDefinition{}, err
	}

	if err = json.Unmarshal(jsonData, &clusterDefinitions); err != nil {
		return &v0.ClusterDefinition{}, err
	}

	switch {
	case len(clusterDefinitions) < 1:
		return &v0.ClusterDefinition{}, errors.New(fmt.Sprintf("no workload definitions with name %s", name))
	case len(clusterDefinitions) > 1:
		return &v0.ClusterDefinition{}, errors.New(fmt.Sprintf("more than one workload definition with name %s returned", name))
	}

	return &clusterDefinitions[0], nil
}

// CreateClusterDefinition creates a new cluster definition
func CreateClusterDefinition(clusterDefinition *v0.ClusterDefinition, apiAddr, apiToken string) (*v0.ClusterDefinition, error) {
	jsonClusterDefinition, err := client.MarshalObject(clusterDefinition)
	if err != nil {
		return clusterDefinition, err
	}

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/cluster_definitions", apiAddr, ApiVersion),
		apiToken,
		http.MethodGet,
		bytes.NewBuffer(jsonClusterDefinition),
		http.StatusCreated,
	)
	if err != nil {
		return clusterDefinition, err
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return clusterDefinition, err
	}

	if err = json.Unmarshal(jsonData, &clusterDefinition); err != nil {
		return clusterDefinition, err
	}

	return clusterDefinition, nil
}

// UpdateClusterDefinition updates a cluster definition
func UpdateClusterDefinition(clusterDefinition *v0.ClusterDefinition, apiAddr, apiToken string) (*v0.ClusterDefinition, error) {
	jsonClusterDefinition, err := client.MarshalObject(clusterDefinition)
	if err != nil {
		return clusterDefinition, err
	}

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/cluster_definitions/%d", apiAddr, ApiVersion, *clusterDefinition.ID),
		apiToken,
		http.MethodPatch,
		bytes.NewBuffer(jsonClusterDefinition),
		http.StatusOK,
	)
	if err != nil {
		return clusterDefinition, err
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return clusterDefinition, err
	}

	if err = json.Unmarshal(jsonData, &clusterDefinition); err != nil {
		return clusterDefinition, err
	}

	return clusterDefinition, nil
}

// GetClusterInstanceByID feteches a cluster instance by ID
func GetClusterInstanceByID(id uint, apiAddr, apiToken string) (*v0.ClusterInstance, error) {
	var clusterInstance v0.ClusterInstance

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/cluster_instances/%d", apiAddr, ApiVersion, id),
		apiToken,
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &clusterInstance, err
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &clusterInstance, err
	}

	if err = json.Unmarshal(jsonData, &clusterInstance); err != nil {
		return &clusterInstance, err
	}

	return &clusterInstance, nil
}

// GetClusterInstanceByName feteches a cluster instance by name
func GetClusterInstanceByName(name, apiAddr, apiToken string) (*v0.ClusterInstance, error) {
	var clusterInstances []v0.ClusterInstance

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/cluster_instances?name=%s", apiAddr, ApiVersion, name),
		apiToken,
		http.MethodGet,
		new(bytes.Buffer),
		http.StatusOK,
	)
	if err != nil {
		return &v0.ClusterInstance{}, err
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.ClusterInstance{}, err
	}

	if err = json.Unmarshal(jsonData, &clusterInstances); err != nil {
		return &v0.ClusterInstance{}, err
	}

	switch {
	case len(clusterInstances) < 1:
		return &v0.ClusterInstance{}, errors.New(fmt.Sprintf("no workload definitions with name %s", name))
	case len(clusterInstances) > 1:
		return &v0.ClusterInstance{}, errors.New(fmt.Sprintf("more than one workload definition with name %s returned", name))
	}

	return &clusterInstances[0], nil
}

// CreateClusterInstance creates a new cluster instance
func CreateClusterInstance(clusterInstance *v0.ClusterInstance, apiAddr, apiToken string) (*v0.ClusterInstance, error) {
	jsonClusterInstance, err := client.MarshalObject(clusterInstance)
	if err != nil {
		return clusterInstance, err
	}

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/cluster_instances", apiAddr, ApiVersion),
		apiToken,
		http.MethodGet,
		bytes.NewBuffer(jsonClusterInstance),
		http.StatusCreated,
	)
	if err != nil {
		return clusterInstance, err
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return clusterInstance, err
	}

	if err = json.Unmarshal(jsonData, &clusterInstance); err != nil {
		return clusterInstance, err
	}

	return clusterInstance, nil
}

// UpdateClusterInstance updates a cluster instance
func UpdateClusterInstance(clusterInstance *v0.ClusterInstance, apiAddr, apiToken string) (*v0.ClusterInstance, error) {
	jsonClusterInstance, err := client.MarshalObject(clusterInstance)
	if err != nil {
		return clusterInstance, err
	}

	response, err := GetResponse(
		fmt.Sprintf("%s/%s/cluster_instances/%d", apiAddr, ApiVersion, *clusterInstance.ID),
		apiToken,
		http.MethodPatch,
		bytes.NewBuffer(jsonClusterInstance),
		http.StatusOK,
	)
	if err != nil {
		return clusterInstance, err
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return clusterInstance, err
	}

	if err = json.Unmarshal(jsonData, &clusterInstance); err != nil {
		return clusterInstance, err
	}

	return clusterInstance, nil
}
