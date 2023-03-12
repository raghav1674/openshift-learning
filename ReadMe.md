## Openshift learning

* *build-agent*: Contains the azure devops debian build agents Dockerfile
* *tasks*: Contains Sample images for use in openshift env
* *main.go*: a sample app which is used for learning about s2i along with .s2i 


### Dockerhub Images:

#### Inside tasks folder:

- raggupta/rhelubi8-devops-base:v1: It uses shell `sleep command` and has terraform and python3 installed

- raggupta/rhelubi8-devops-base:supervisord-v1: It uses supervisord and has terraform and python3 installed

#### Inside build-agent/rhelubi8 folder:

- raggupta/azdo-rhelubi8-build-docker-agent:v1-latest: It will pull the latest azure devops build agent binary and run it 

- raggupta/azdo-rhelubi8-build-docker-agent:v1-2.217.2: It will run 2.217.2 azure devops build agent

