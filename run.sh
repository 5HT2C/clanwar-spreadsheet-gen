#!/bin/bash

if [[ -z "$2" ]]; then
    F="$(./fetch_current.sh)"
else
    F="$2"
fi

if [[ -n "$1" ]]; then
    T="$1"
else
    T="$(./fetch_last.sh "$(jq -r '.opponent.tag' "$F")")"
fi

./csg -file "$F" -lastTag "$T" | sed "s/,\"=/,\"~/g"
