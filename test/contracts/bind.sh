#!/bin/sh

set -e

gen() {
    local package=$1

    abigen --bin bin/${package}.bin --abi abi/${package}.abi --pkg=${package} --out=${package}/${package}.go
}

gen verifybatchesmock
gen claimmock
gen claimmockcaller
gen claimmocktest