SHELL=/bin/zsh # necesario porque bash en OS X no tiene globstar

# macros
GB := $(shell command -v gb)
GULP := $(shell command -v gulp)

# source files
gosrc := $(shell ls src/**/*.go)
staticsrc := $(shell ls static/app/**/*.*)

# default
all: app

# archivos requeridos
bin/registro: $(gosrc)
ifndef GB
		$(error "gb no está instalado. Obtenlo en http://getgb.io")
endif
		gb build

bin/registro-linux-amd64: $(gosrc)
ifndef GB
		$(error "gb no está instalado. Obtenlo en http://getgb.io")
endif
		env GOOS=linux GOARCH=amd64 gb build

static/dist/scripts/bundle.js: $(staticsrc)
ifndef GULP
		$(error "gulp no está instalado. Obtenlo en https://webpack.github.io")
endif
		cd static && \
		gulp build

# aliases
frontend: static/dist
backend: bin/registro
backend_linux: bin/registro-linux-amd64

app: frontend backend
	mkdir -p build
	cp -r bin/registro build/registro
	cp -r static/dist build/static

docker: frontend backend_linux
	docker-compose build

clean:
		rm -rf bin
		rm -rf pkg
		rm -rf static/dist
		rm -rf build
