#!/bin/sh
#
# Code coverage generation
name=$(date +%s)
COVERAGE_DIR="${COVERAGE_DIR:-coverage}"
PKG_LIST=$(go list ./... | grep -v /vendor/)
ppath=${COVERAGE_DIR}/ullr-go-${name}.cov

# Create the coverage files directory
mkdir -p "$COVERAGE_DIR";

# Create a coverage file for each package
#for package in ${PKG_LIST}; do
#    go test -covermode=count -coverprofile="${COVERAGE_DIR}/${package##*/}.cov" "$package" ;
#done ;

go test -covermode=count -coverprofile=${ppath} ${PKG_LIST};

# Merge the coverage profile files
# echo 'mode: count' > "${COVERAGE_DIR}"/coverage.cov ;
# tail -q -n +2 "${COVERAGE_DIR}"/*.cov >> "${COVERAGE_DIR}"/coverage.cov ;

# Display the global code coverage
go tool cover -func ${ppath} ;

# If needed, generate HTML report
if [ "$1" == "html" ]; then
    go tool cover -html ${ppath} -o coverage.html ;
fi

# Remove the coverage files directory
rm -rf "$COVERAGE_DIR";
