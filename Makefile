.PHONY: build build-image

build:
	$(MAKE) -C ./rpm build

build-image:
	$(MAKE) -C ./app test-app.tar.xz
