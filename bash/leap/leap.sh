#!/usr/bin/env bash

set -o errexit
set -o nounset

main(){
    local year=${1:-}
    if [[  "$#" -ne "1" || ! "$year" =~ ^[0-9]+$ ]]
    then
        echo 'Usage: leap.sh <year>'
        exit 1
    
    elif (( $year % 4 == 0)) && (($year % 400 == 0))
    then
        echo 'true'
    else
        if (( $year % 4 == 0 )) && (($year % 100 != 0))
        then
            echo 'true'
        else
            echo 'false'
        fi
    fi
    
}

main "$@"
