/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:   "updateTask",
	Short: "Create a new task definition revision in a specific ECS cluster and service",
	Run: func(cmd *cobra.Command, args []string) {
		imageName, _ := cmd.Flags().GetString("image")
		taskFamily, _ := cmd.Flags().GetString("family")
		region, _ := cmd.Flags().GetString("region")
		updateTaskDefinition(imageName, taskFamily, region)
	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)

	updateTaskCmd.Flags().String("family", "", "ECS task definition family name")
	updateTaskCmd.MarkFlagRequired("service")
	updateTaskCmd.Flags().String("image", "", "Docker image to be update in the task definition with tag")
	updateTaskCmd.MarkFlagRequired("image")
	updateTaskCmd.Flags().String("region", "us-east-1", "AWS Region")

}
