#!/bin/bash

source .env || exit 1
[[ -d data/ ]] || mkdir data/

F="data/war_tmp.json"
curl -s -H "Authorization: Bearer $TOKEN" https://api.clashofclans.com/v1/clans/%23PV9P9JQ8/currentwar | jq > "$F"

T="$(jq -r '.opponent.tag' "$F")"
D="$(jq -r '.preparationStartTime' "$F" | sed "s/T/ /g")"
D="${D:0:4}-${D:4}"
D="${D:0:7}-${D:7}"
D="${D:0:13}:${D:13}"
D="${D:0:16}:${D:16}"
D="data/war_${T}_$(date -d "$D" -u +%Y-%m-%dT%H:%M).json"

mv "$F" "$D"
echo "$D"
