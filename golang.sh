#!/bin/bash
grep 'Text: (string)' output.txt | cut -d "\"" -f2 | grep 'func'