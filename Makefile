DOCKER_IMAGE_NAME := swrc

###################
# Docker          #
###################
.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 ${DOCKER_IMAGE_NAME}

###################
# App             #
###################
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/swrc ./cmd/gapi/main.go

###################
# Utils           #
###################
.PHONY: clear-notes
clear-notes:
	find . -type f -name "*_notes.md" -exec rm -f {} \;