# Имя приложения
APP_NAME := image-converter

# Компилятор Go
GO := go

# Определение платформы
ifeq ($(OS),Windows_NT)
    OS_FLAG := windows
    BIN_EXT := .exe
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        OS_FLAG := linux
        BIN_EXT :=
    endif
    ifeq ($(UNAME_S),Darwin)
        OS_FLAG := darwin
        BIN_EXT :=
    endif
endif

# Флаги компиляции
BUILD_FLAGS := -ldflags="-s -w"

# Сборка приложения
.PHONY: build
build:
	GOOS=$(OS_FLAG) GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o $(APP_NAME)$(BIN_EXT)

# Очистка
.PHONY: clean
clean:
	rm -f $(APP_NAME)$(BIN_EXT)
