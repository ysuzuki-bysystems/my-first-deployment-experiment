sources := ./main.go ./go.mod

.PHONY: clean

test-app.tar.xz: $(sources)
	docker build . -t test-app:latest
	docker image save test-app:latest | xz -c > $@

clean:
	$(RM) test-app.tar.xz
