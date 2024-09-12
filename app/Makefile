sources := ./main.go ./go.mod

test-app.tar.xz: $(sources)
	docker build . -t test-app:latest
	docker image save test-app:latest | xz -c > $@
