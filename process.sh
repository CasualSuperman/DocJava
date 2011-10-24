#!/bin/bash
for file in `ls html`
do
	"./DocJava -f test/$file > done/"${$file%.html}".java"
done
