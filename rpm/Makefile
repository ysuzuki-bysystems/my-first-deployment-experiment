NFPM_VERSION := 2.40.0
NFPM_DOWNLOAD_URL := https://github.com/goreleaser/nfpm/releases/download/v$(NFPM_VERSION)/nfpm_$(NFPM_VERSION)_Linux_x86_64.tar.gz

.PHONY: build clean

# RPM をビルドする
build: bin/nfpm usr/local/share/test/test-proxy.tar.xz
	mkdir -p dist
	./bin/nfpm pkg --packager rpm --target dist/

# nfpm をダウンロードする
bin/nfpm:
	mkdir -p bin
	curl -sSfL $(NFPM_DOWNLOAD_URL) | tar -Cbin -zxf- nfpm

usr/local/share/test/test-proxy.tar.xz:
	mkdir -p usr/local/share/test
	docker pull caddy:latest
	docker image save caddy:latest | xz > $@

clean:
	$(RM) -r dist/ bin/
