#!/bin/bash
for file in `ls html`
do
	./DocJava -f html/$file > done/${file%.html}.java
done
