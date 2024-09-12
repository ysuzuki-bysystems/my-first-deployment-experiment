.PHONY: build build-image clean

build:
	$(MAKE) -C ./rpm build

build-image:
	$(MAKE) -C ./app test-app.tar.xz

clean:
	$(MAKE) -C ./rpm clean
	$(MAKE) -C ./app clean
