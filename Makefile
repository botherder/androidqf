BUILD_FOLDER  = "$(shell pwd)/build"
ASSETS_FOLDER = "$(shell pwd)/assets"

FLAGS_LINUX   = GOOS=linux
FLAGS_DARWIN  = GOOS=darwin
FLAGS_WINDOWS = GOOS=windows GOARCH=amd64 CC=i686-w64-mingw32-gcc CGO_ENABLED=1

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
	go mod tidy
	@echo "[deps] Dependencies installed."

windows:
	@mkdir -p $(BUILD_FOLDER)

	@if [ ! -f /tmp/$(PLATFORMTOOLS_WINDOWS) ]; then \
		echo "Downloading Windows Android Platform Tools..."; \
		wget $(PLATFORMTOOLS_URL)$(PLATFORMTOOLS_WINDOWS) -O /tmp/$(PLATFORMTOOLS_WINDOWS); \
	fi

	@rm -rf $(PLATFORMTOOLS_FOLDER)
	@cd /tmp && unzip -u $(PLATFORMTOOLS_WINDOWS)
	@cp $(PLATFORMTOOLS_FOLDER)/AdbWinApi.dll $(ASSETS_FOLDER)
	@cp $(PLATFORMTOOLS_FOLDER)/AdbWinUsbApi.dll $(ASSETS_FOLDER)
	@cp $(PLATFORMTOOLS_FOLDER)/adb.exe $(ASSETS_FOLDER)

	@echo "[builder] Building Windows binary for amd64"

	$(FLAGS_WINDOWS) go build --ldflags '-s -w -extldflags "-static"' -o $(BUILD_FOLDER)/androidqf_windows_amd64.exe .

	@echo "[builder] Done!"

darwin:
	@mkdir -p $(BUILD_FOLDER)

	@if [ ! -f /tmp/$(PLATFORMTOOLS_DARWIN) ]; then \
		echo "Downloading Darwin Android Platform Tools..."; \
		wget $(PLATFORMTOOLS_URL)$(PLATFORMTOOLS_DARWIN) -O /tmp/$(PLATFORMTOOLS_DARWIN); \
	fi

	@rm -rf $(PLATFORMTOOLS_FOLDER)
	@cd /tmp && unzip -u $(PLATFORMTOOLS_DARWIN)
	@cp $(PLATFORMTOOLS_FOLDER)/adb $(ASSETS_FOLDER)

	@echo "[builder] Building Darwin binary for amd64"

	$(FLAGS_DARWIN) GOARCH=amd64 go build --ldflags '-s -w' -o $(BUILD_FOLDER)/androidqf_darwin_amd64 .
	$(FLAGS_DARWIN) GOARCH=arm64 go build --ldflags '-s -w' -o $(BUILD_FOLDER)/androidqf_darwin_arm64 .

	@echo "[builder] Done!"

linux:
	@mkdir -p $(BUILD_FOLDER)

	@if [ ! -f /tmp/$(PLATFORMTOOLS_LINUX) ]; then \
		echo "Downloading Linux Android Platform Tools..."; \
		wget $(PLATFORMTOOLS_URL)$(PLATFORMTOOLS_LINUX) -O /tmp/$(PLATFORMTOOLS_LINUX); \
	fi

	@rm -rf $(PLATFORMTOOLS_FOLDER)
	@cd /tmp && unzip -u $(PLATFORMTOOLS_LINUX)
	@cp $(PLATFORMTOOLS_FOLDER)/adb $(ASSETS_FOLDER)

	@echo "[builder] Building Linux binary for amd64"

	@$(FLAGS_LINUX) GOARCH=amd64 go build --ldflags '-s -w' -o $(BUILD_FOLDER)/androidqf_linux_amd64 .
	@$(FLAGS_LINUX) GOARCH=arm64 go build --ldflags '-s -w' -o $(BUILD_FOLDER)/androidqf_linux_arm64 .

	@echo "[builder] Done!"

clean:
	rm -rf $(BUILD_FOLDER)
	rm -f $(ASSETS_FOLDER)/adb $(ASSETS_FOLDER)/adb.exe $(ASSETS_FOLDER)/AdbWinApi.dll $(ASSETS_FOLDER)/AdbWinUsbApi.dll
