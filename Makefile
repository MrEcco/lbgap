DOCKERNAME=mrecco/lbgap
DOCKERTAG=v1.0.0
DOCKERFILE=Dockerfile

run:
	@go run main.go

build:
	@docker build -f $(DOCKERFILE) -t $(DOCKERNAME):$(DOCKERTAG) .

push:
	@docker push $(DOCKERNAME):$(DOCKERTAG)
	@docker tag $(DOCKERNAME):$(DOCKERTAG) $(DOCKERNAME):latest
	@docker push $(DOCKERNAME):latest

native-build:
	@go build -o lbgap .

test:
	@echo I really have no idea how to mistake in THIS...
	# @go test .
