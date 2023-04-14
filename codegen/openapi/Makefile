# Set theese variables
SWAGGER_VERSION ?= 4.1.0
DESTDIR ?= ./build
TDPROTO_OPEN_API_DIR ?= ./
OPENAPI_FILE ?= v4.json
SWAGGER_SERVER_URL ?= https://web.tada.team

SWAGGER_URL = https://github.com/swagger-api/swagger-ui/archive/refs/tags/v${SWAGGER_VERSION}.tar.gz
SWAGGER_ARCHIVE = swagger.tar.gz
SWAGGER_DIST_PATH = swagger-ui-${SWAGGER_VERSION}/dist
DESTDIR_ABSOLUTE := $(shell readlink --canonicalize "${DESTDIR}")
# Escape dots for sed
SED_OPENAPI_FILE := $(shell printf '%s' '${OPENAPI_FILE}' | sed 's/\./\\\./g' )
WORKING_DIR := $(shell pwd)
TDPROTO_GIT_VERSION := $(shell git describe --tags)

$(SWAGGER_ARCHIVE):
	wget "${SWAGGER_URL}" --output-document "${SWAGGER_ARCHIVE}"

$(DESTDIR_ABSOLUTE):
	mkdir --parents "${DESTDIR_ABSOLUTE}"

install: $(SWAGGER_ARCHIVE) $(DESTDIR_ABSOLUTE)
	tar --extract --file "./${SWAGGER_ARCHIVE}" "${SWAGGER_DIST_PATH}"
	cp -r "./${SWAGGER_DIST_PATH}" "${DESTDIR_ABSOLUTE}"
#	mv "${DESTDIR_ABSOLUTE}/dist/index.html" "${DESTDIR_ABSOLUTE}/"
	echo ${SED_OPENAPI_FILE} ${OPENAPI_FILE}
	sed --in-place \
		-e 's/\.\//dist\//g' \
		-e 's/url\: ".*"/url\: "${SED_OPENAPI_FILE}"/g' \
		"${DESTDIR_ABSOLUTE}/index.html"
	
	echo WORKING DIRECTORY "${WORKING_DIR}"
	cd "${TDPROTO_OPEN_API_DIR}" && \
		go run  ./ \
			-server "${SWAGGER_SERVER_URL}" -version "${TDPROTO_GIT_VERSION}" \
		> "${DESTDIR_ABSOLUTE}/${OPENAPI_FILE}"
