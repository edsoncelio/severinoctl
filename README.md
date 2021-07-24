# severinoctl

A tool to automate some tasks in ECS/ECR.
> Work in progress...

## Prerequisites
* awscli
* working aws credentials

## How to install
TODO

## How to use

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

## TODO
 - [ ] Add unit tests
 - [ ] Use aws sdk instead of aws cli
 - [x] Add option to check if a tag exists in a ECR repository
 - [ ] Add option to create a ECS task definition revision (and show a diff)
