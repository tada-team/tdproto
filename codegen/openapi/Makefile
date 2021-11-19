# Set theese variables
SWAGGER_VERSION ?= 4.1.0
DESTDIR ?= ./build
TDPROTO_OPEN_API_DIR ?= ./
OPENAPI_FILE ?= v4.json

SWAGGER_URL = https://github.com/swagger-api/swagger-ui/archive/refs/tags/v${SWAGGER_VERSION}.tar.gz
SWAGGER_ARCHIVE = swagger.tar.gz
SWAGGER_DIST_PATH = swagger-ui-${SWAGGER_VERSION}/dist
# Escape dots for sed
SED_OPENAPI_FILE := $(shell printf '%s' '${OPENAPI_FILE}' | sed 's/\./\\\./g' )

$(SWAGGER_ARCHIVE):
	wget "${SWAGGER_URL}" --output-document "${SWAGGER_ARCHIVE}"

$(DESTDIR):
	mkdir --parents "${DESTDIR}"

install: $(SWAGGER_ARCHIVE) $(DESTDIR)
	tar --extract --file "./${SWAGGER_ARCHIVE}" "${SWAGGER_DIST_PATH}"
	cp -r "./${SWAGGER_DIST_PATH}" "${DESTDIR}"
	mv "${DESTDIR}/dist/index.html" "${DESTDIR}/"
	echo ${SED_OPENAPI_FILE} ${OPENAPI_FILE}
	sed --in-place \
		-e 's/\.\//dist\//g' \
		-e 's/url\: ".*"/url\: "${SED_OPENAPI_FILE}"/g' \
		"${DESTDIR}/index.html"
	go run "${TDPROTO_OPEN_API_DIR}" > "${DESTDIR}/${OPENAPI_FILE}"
