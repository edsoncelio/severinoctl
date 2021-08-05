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
	TaskDefinitionArn       string          `json:"taskDefinitionArn"`
	IpcMode                 string          `json:"ipcMode"`
	ExecutionRoleArn        string          `json:"executionRoleArn"`
	ContainerDefinitions    []ContainerOpts `json:"containerDefinitions"`
	CPU                     string          `json:"cpu"`
	Memory                  string          `json:"memory"`
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
	Cpu               int16             `json:"cpu,omitempty"`
	Memory            int16             `json:"memory,omitempty"`
	MemoryReservation int16             `json:"memoryReservation,omitempty"`
	Environment       []EnvironmentOpts `json:"environment"`
}

type EnvironmentOpts struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func updateTaskDefinition(imageName string, taskFamily string, region string) {
	var OutputTaskDefinition TaskDefinition
	var formattedContainerDefinitions []ContainerOpts
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
		if OutputTaskDefinition.TaskDefinitions.ContainerDefinitions[0].Image == imageName {
			fmt.Printf("⚠️  The image is already registered\n")
			return
		} else {
			fmt.Printf("🔍 Actual image found: %s - New image to use: '%s'\n", OutputTaskDefinition.TaskDefinitions.ContainerDefinitions[0].Image, imageName)
		}
	}

	OutputTaskDefinition.TaskDefinitions.ContainerDefinitions[0].Image = imageName

	if len(OutputTaskDefinition.TaskDefinitions.ContainerDefinitions) > 0 {
		formattedContainerDefinitions = prepareContainerDefinitions(OutputTaskDefinition.TaskDefinitions.ContainerDefinitions, imageName)
	} else {
		fmt.Printf("%s", "No container definition found")

	}
	if OutputTaskDefinition.TaskDefinitions.Memory == "" {
		OutputTaskDefinition.TaskDefinitions.Memory = "200" // set a default value to task memory
	}
	fmt.Printf("⌛  Creating the new revision...\n")
	outTask := registerTask(taskFamily, formattedContainerDefinitions, OutputTaskDefinition.TaskDefinitions.Memory)
	fmt.Printf("🎉 Revision '%s' created!\n", outTask.TaskDefinitions.TaskDefinitionArn)
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

func prepareContainerDefinitions(containerDefinition []ContainerOpts, image string) []ContainerOpts {
	containerItem := 0
	for i, definitions := range containerDefinition {
		if definitions.Image == image {
			containerItem = i
		}
	}

	if containerDefinition[containerItem].Command == nil {
		containerDefinition[containerItem].Command = []string{}
	}

	if containerDefinition[containerItem].EntryPoint == nil {
		containerDefinition[containerItem].EntryPoint = []string{}
	}

	return containerDefinition
}

func registerTask(Taskfamily string, containerDefinition []ContainerOpts, memory string) TaskDefinition {
	var taskDefinition TaskDefinition
	_, lookErr := exec.LookPath("aws")
	if lookErr != nil {
		panic(lookErr)
	}

	outJson, _ := json.Marshal(containerDefinition)
	formmatedJson := fmt.Sprintf(`%s`, string(outJson))

	cmd := exec.Command("aws", "ecs", "register-task-definition", "--family", Taskfamily, "--container-definitions", formmatedJson, "--memory", memory)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(out), &taskDefinition)
	return taskDefinition
}
