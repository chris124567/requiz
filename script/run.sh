#!/bin/sh
set -eu

npx prettier --single-quote false --html-whitespace-sensitivity strict --no-bracket-spacing --quote-props preserve --trailing-comma es5 --print-width 100 --write --tab-width 4 src/*
npx tsc