# severinoctl

A tool to automate some tasks in ECS/ECR.
> Work in progress...

## Prerequisites
* awscli
* working aws credentials
* environment `AWS_REGION` exported

## How to install
Get a binary from https://github.com/edsoncelio/severinoctl/releases 

## Features
 - [x] Check if a tag exists in a ECR repository
 - [x] List task definition ARN revisions to a family
 - [x] Create a task definition revision to a new image version

## Usage

### To check if a specific image tag exists in a specific ECR repository:   
`severinoctl checkTag --registry my-ecr-registry --tag 1.0`

If the tag exists, you will receive the following message:
```
✅ Tag '1.0' found! digest: <IMAGE DIGEST HERE> and repository: my-ecr-registry
```

Otherwise:
```
❌ Tag '1.0' not found, reason: <REASON> - try again!
```

### To list all the revisions related to a specific task definition:   
`severinoctl listTask --name sample-app`

And the output:
```
TaskDefinition ARN to revision 0: arn:aws:ecs:us-east-1:<AWS ACCOUNT>:task-definition/sample-app:1
TaskDefinition ARN to revision 1: arn:aws:ecs:us-east-1:<AWS ACCOUNT>:task-definition/sample-app:2
...
```

### To create a new task definition revision to a specific image:   
`severinoctl updateTask --family sample-app --image myimage:v1.0.0`

If the image and tag already exists in the task definition:   
`⚠️  The image is already registered`

Otherwise:   
```
🔍 Actual image found: myimage:v1.0.0 - New image to use: 'myimage:v2.0.0'
⌛  Creating the new revision...
🎉 Revision 'arn:aws:ecs:us-east-1:<AWS ACCOUNT>:task-definition/sample-app:19' created!
```

## TODO
 - [ ] Add tests
 - [ ] Use aws sdk instead of aws cli
 - [x] Configure goreleaser
 - [ ] Add option to create a ECS task definition revision (and show a diff) - partial done
