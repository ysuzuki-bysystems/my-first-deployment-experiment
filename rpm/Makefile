NFPM_VERSION := 2.40.0
NFPM_DOWNLOAD_URL := https://github.com/goreleaser/nfpm/releases/download/v$(NFPM_VERSION)/nfpm_$(NFPM_VERSION)_Linux_x86_64.tar.gz

.PHONY: build clean

build: bin/nfpm
	mkdir -p dist
	./bin/nfpm pkg --packager rpm --target dist/

bin/nfpm:
	mkdir -p bin
	curl -sSfL $(NFPM_DOWNLOAD_URL) | tar -Cbin -zxf- nfpm

clean:
	$(RM) -r dist/ bin/
