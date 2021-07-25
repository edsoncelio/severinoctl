# severinoctl

A tool to automate some tasks in ECS/ECR.
> Work in progress...

## Prerequisites
* awscli
* working aws credentials
* environment `AWS_REGION` exported

## How to install
TODO

## Features
 - [x] Check if a tag exists in a ECR repository
 - [x] List task definition ARN revisions to a family

## Usage

1. To check if a specific image tag exists in a specific ECR repository:   
`severinoctl checkTag --registry my-ecr-registry --tag 1.0`

If the tag exists, you will receive the following message:
```
✅ Tag '1.0' found! digest: <IMAGE DIGEST HERE> and repository: my-ecr-registry
```

Otherwise:
```
❌ Tag '1.0' not found, reason: <REASON> - try again!
```

2. To list all the revisions related to a specific task definition:   
`severinoctl listTask --name sample-app`

Then, you i'll see:
```
TaskDefinition ARN to revision 0: arn:aws:ecs:us-east-1:<AWS ACCOUNT>:task-definition/sample-app:1
TaskDefinition ARN to revision 1: arn:aws:ecs:us-east-1:<AWS ACCOUNT>:task-definition/sample-app:2
...
```

## TODO
 - [ ] Add tests
 - [ ] Use aws sdk instead of aws cli
 - [ ] Configure goreleaser
 - [ ] Add option to create a ECS task definition revision (and show a diff)
