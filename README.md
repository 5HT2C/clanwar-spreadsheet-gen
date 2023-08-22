# clanwar-spreadsheet-gen

## Getting Clan War data

1. Get an API key from https://developer.clashofclans.com/
2. Create `.env`:
```env
TOKEN=clash api key goes here
```
3. Run `./fetch_current.sh` and it will fetch the current war into `data/`.
This only works before a new war has been started.

## Generating spreadsheet

1. Run `make`
2. Run `./run.sh [data json file]`. This will fetch the latest war data if you do not specify the path of said json file.
