# macros
GB := $(shell command -v gb)
GULP := $(shell command -v gulp)

# default
all: clean app

# archivos requeridos
bin/registro:
ifndef GB
		$(error "gb no está instalado. Obtenlo en http://getgb.io")
endif
		gb build

bin/registro-linux-amd64:
ifndef GB
		$(error "gb no está instalado. Obtenlo en http://getgb.io")
endif
		env GOOS=linux GOARCH=amd64 gb build

static/dist:
ifndef GULP
		$(error "gulp no está instalado. Obtenlo en https://webpack.github.io")
endif
		cd static && \
		gulp build

# aliases
frontend: static/dist
backend: bin/registro
backend_linux: bin/gato-linux-amd64

app: frontend backend
	mkdir -p build
	cp -r bin/registro build/registro
	cp -r static/dist build/static

#run: app
#	cd build && \
#	./registro

docker: frontend backend_linux
		docker build -t gato .

clean:
		rm -rf bin
		rm -rf pkg
		rm -rf static/dist
		rm -rf build
