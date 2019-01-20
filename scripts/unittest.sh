#!/usr/bin/env bash

if [[ -z "${REPORTER}" ]]
then
  export REPORTER=cc-test-reporter
fi

function unittest() {
    EXIT_CODE=0
    # unit tests all the packages
    for dir in `ls . | grep -v vendor`;
    do
        if [[ -d ${dir} && ! "${dir}" =~ ^(scripts|tmp)$ ]]
        then
            go test ./${dir} -race -v -coverprofile cover.${dir}.out || EXIT_CODE=1
        fi

    done

    # unit test the main package
    go test . -race -v -coverprofile cover.out || EXIT_CODE=1

    # exit poorly if any test failed
    exit ${EXIT_CODE}
}


function generate_coverage() {
    echo "Using $REPORTER to generate JSON files"

    for cover_file in `find . -name "cover.*.out"`;
    do
        report_file=$(echo ${cover_file} | sed 's/\.\/cover\.\(.*\)\.out/cc.\1.json/')
        echo "Generating $report_file from ${cover_file}"
        ${REPORTER} format-coverage -t gocov -o tmp/${report_file} ${cover_file}
    done

    echo "Generating cc.main.json from ./cover.out"
    ${REPORTER} format-coverage -t gocov -o tmp/cc.main.json ./cover.out
}

function sum_coverage() {
    count=$(find tmp -name "cc.*.json" | wc -l)
    ${REPORTER} sum-coverage tmp/cc.*.json -o tmp/cc.total.json -p ${count}
}


# Execute the scripts
case $1 in

run)
  unittest
;;

coverage)
  generate_coverage
;;

sum)
  sum_coverage
;;

*)
  exit 1

esac