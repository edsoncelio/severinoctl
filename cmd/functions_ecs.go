package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type TaskDefinitionList struct {
	TaskDefinitionArns []string `json:"taskDefinitionArns"`
}

type TaskDefinition struct {
	TaskDefinitions TaskDefinitions `json:"taskDefinition"`
}

type TaskDefinitions struct {
	TaskDefinitionArn string `json:"taskDefinitionArn"`
	IpcMode           string `json:"ipcMode"`
	ExecutionRoleArn  string `json:"executionRoleArn"`

	ContainerDefinitions    []ContainerOpts `json:"containerDefinitions"`
	CPU                     string          `json:"cpu"`
	Family                  string          `json:"family"`
	NetworkMode             string          `json:"networkMode"`
	RequiresCompatibilities []string        `json:"requiresCompatibilities"`
}

type ContainerOpts struct {
	Name              string            `json:"name"`
	Command           []string          `json:"command"`
	EntryPoint        []string          `json:"entryPoint"`
	Essential         bool              `json:"essential"`
	Image             string            `json:"image"`
	Cpu               int16             `json:"cpu"`
	Memory            int16             `json:"memory"`
	MemoryReservation int16             `json:"memoryReservation"`
	Environment       []EnvironmentOpts `json:"environment"`
}

type EnvironmentOpts struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func updateTaskDefinition(imageName string, taskFamily string, region string) {
	var OutputTaskDefinition TaskDefinition

	_, lookErr := exec.LookPath("aws")
	if lookErr != nil {
		panic(lookErr)
	}

	out, err := exec.Command("aws", "ecs", "describe-task-definition", "--task-definition", taskFamily, "--region", region).Output()

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(out), &OutputTaskDefinition)

	if len(OutputTaskDefinition.TaskDefinitions.ContainerDefinitions) > 0 {
		fmt.Printf("Docker image found: %s, the new revision will use the image '%s'\n", OutputTaskDefinition.TaskDefinitions.ContainerDefinitions[0].Image, imageName)
	}

	OutputTaskDefinition.TaskDefinitions.ContainerDefinitions[0].Image = imageName

	outJson, _ := json.Marshal(OutputTaskDefinition.TaskDefinitions)
	fmt.Println(string(outJson))
	registerTask(string(outJson), region)

}

func listTaskDefinition(familyPrefix string) {

	var listTaskDefinition TaskDefinitionList

	_, lookErr := exec.LookPath("aws")
	if lookErr != nil {
		panic(lookErr)
	}

	out, err := exec.Command("aws", "ecs", "list-task-definitions", "--family-prefix", familyPrefix).Output()

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(out), &listTaskDefinition)

	for task, taskValue := range listTaskDefinition.TaskDefinitionArns {
		fmt.Printf("TaskDefinition ARN to revision %d: %s\n", task, taskValue)
	}
}

func registerTask(taskDefinition string, region string) {
	_, lookErr := exec.LookPath("aws")
	if lookErr != nil {
		panic(lookErr)
	}

	out, err := exec.Command("aws", "ecs", "register-task-definition", "--cli-input-json", taskDefinition, "--region", region).Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
