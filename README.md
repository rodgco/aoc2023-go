# Advent of Code 2023 - Go Lang

My attempt to solve [Advent of Code 202](https://adventofcode.com/2023) using the (Go)[https://go.dev] Language.

## Tools

* (Advent of Code Downloader - aocdl)[https://github.com/GreenLightning/advent-of-code-downloader]

## Configuration

* Follow aocdl instructions to get your Session Cookie
* Create a `.env`, simply `cp .env.example .env`
* Edit `.env` and add your Session Cookie

## Usage

The `start.sh` will automatically create a directory for each day and download your input file.

Without extra parameters it will create this structure for the current day, you can specify a number
as an argument, so the strucuture for that day will be created.

```
./start.sh 3
```

This will create the `day3` folder and initialize it with a template and input for the third day of AOC.

