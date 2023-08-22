# clanwar-spreadsheet-gen

## Getting Clan War data

1. Get an API key from https://developer.clashofclans.com/.
2. Create a `.env`:
```env
TOKEN=clash api key goes here
```
3. Run `./fetch_current.sh` and it will fetch the current war into `data/`.
This only works before a new war has been started.

## Generating spreadsheet

1. Run `make`.
2. Run `./run.sh [data json file]`.
This will fetch the latest war data if you do not specify the path of said json file.

#### Using generated data

Optionally, nice formatting to steal can be found at https://frogg.ie/llama

1. Add Sheet
2. <kbd>Ctrl</kbd> + <kbd>V</kbd>
3. Data → Split text to columns
4. Format → Alternating Columns
5. Select Column E
6. <kbd>Ctrl</kbd> + <kbd>H</kbd>
7. Find `~`, Replace with `=`, Replace All
8. Select Columns A to P → Right Click → Resize column → Fit to data
9. Select Columns D to E → Insert → Checkbox
