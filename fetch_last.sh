#!/bin/bash

for f in data/war_#*; do
    T0="$(date -d "$(echo "$f" | cut -c 21-36)" +%s)"
    TAG="$(jq -r '.opponent.tag' "$f")"

    # If we haven't set F1, or the current unix T0 is greater than unix T1
    if [[ -z "$F1" ]] || [[ "$T0" -gt "$T1" ]]; then
        # Filter if current tag matches arg 1 filter
        if [[ -n "$1" ]] && [[ "$TAG" == "$1" ]]; then
            continue
        fi
 
        F1="$f"
        T1="$T0" 
    fi
done

jq -r '.opponent.tag' "$F1"
