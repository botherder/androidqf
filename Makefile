BUILD_FOLDER	= $(shell pwd)/build
ASSETS_FOLDER	= $(shell pwd)/assets

FLAGS_DARWIN	= GOOS=darwin
FLAGS_WINDOWS	= GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CGO_ENABLED=1

PLATFORMTOOLS_URL     = https://dl.google.com/android/repository/
PLATFORMTOOLS_WINDOWS = platform-tools-latest-windows.zip
PLATFORMTOOLS_DARWIN  = platform-tools-latest-darwin.zip
PLATFORMTOOLS_LINUX   = platform-tools-latest-linux.zip
PLATFORMTOOLS_FOLDER  = /tmp/platform-tools

lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...

deps:
	@echo "[deps] Installing dependencies..."
	go mod download
	go get -u github.com/go-bindata/go-bindata/...
	@echo "[deps] Dependencies installed."

windows: deps
	@mkdir -p $(BUILD_FOLDER)
	@mkdir -p $(ASSETS_FOLDER)

	@if [ ! -f /tmp/$(PLATFORMTOOLS_WINDOWS) ]; then \
		echo "Downloading Windows Android Platform Tools..."; \
		wget $(PLATFORMTOOLS_URL)$(PLATFORMTOOLS_WINDOWS) -O /tmp/$(PLATFORMTOOLS_WINDOWS); \
	fi

	@rm -rf $(PLATFORMTOOLS_FOLDER)
	@cd /tmp && unzip -u $(PLATFORMTOOLS_WINDOWS)
	@cp $(PLATFORMTOOLS_FOLDER)/AdbWinApi.dll $(ASSETS_FOLDER)
	@cp $(PLATFORMTOOLS_FOLDER)/AdbWinUsbApi.dll $(ASSETS_FOLDER)
	@cp $(PLATFORMTOOLS_FOLDER)/adb.exe $(ASSETS_FOLDER)
	@go-bindata -pkg adb -o adb/bindata.go -prefix $(ASSETS_FOLDER) $(ASSETS_FOLDER)

	$(FLAGS_WINDOWS) go build --ldflags '-s -w -extldflags "-static"' -o $(BUILD_FOLDER)/androidqf_windows.exe ./cmd/

darwin: deps
	@mkdir -p $(BUILD_FOLDER)
	@mkdir -p $(ASSETS_FOLDER)

	@if [ ! -f /tmp/$(PLATFORMTOOLS_DARWIN) ]; then \
		echo "Downloading Darwin Android Platform Tools..."; \
		wget $(PLATFORMTOOLS_URL)$(PLATFORMTOOLS_DARWIN) -O /tmp/$(PLATFORMTOOLS_DARWIN); \
	fi

	@rm -rf $(PLATFORMTOOLS_FOLDER)
	@cd /tmp && unzip -u $(PLATFORMTOOLS_DARWIN)
	@cp $(PLATFORMTOOLS_FOLDER)/adb $(ASSETS_FOLDER)
	@go-bindata -pkg adb -o adb/bindata.go -prefix $(ASSETS_FOLDER) $(ASSETS_FOLDER)

	$(FLAGS_DARWIN) go build --ldflags '-s -w' -o $(BUILD_FOLDER)/androidqf_darwin ./cmd/

linux: deps
	@mkdir -p $(BUILD_FOLDER)
	@mkdir -p $(ASSETS_FOLDER)

	@if [ ! -f /tmp/$(PLATFORMTOOLS_LINUX) ]; then \
		echo "Downloading Linux Android Platform Tools..."; \
		wget $(PLATFORMTOOLS_URL)$(PLATFORMTOOLS_LINUX) -O /tmp/$(PLATFORMTOOLS_LINUX); \
	fi

	@rm -rf $(PLATFORMTOOLS_FOLDER)
	@cd /tmp && unzip -u $(PLATFORMTOOLS_LINUX)
	@cp $(PLATFORMTOOLS_FOLDER)/adb $(ASSETS_FOLDER)
	@go-bindata -pkg adb -o adb/bindata.go -prefix $(ASSETS_FOLDER) $(ASSETS_FOLDER)

	@go build --ldflags '-s -w' -o $(BUILD_FOLDER)/androidqf_linux ./cmd/

clean:
	rm -rf $(ASSETS_FOLDER)
	rm -rf $(BUILD_FOLDER)
	rm -f ./*/bindata.go
