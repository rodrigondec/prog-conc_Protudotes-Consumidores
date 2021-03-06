#!/bin/bash

#constants
MAX_ITERATIONS=10
FILE_NAME_START="bmk_"
FILE_EXTENSION=".txt"
QTT_GOROUTINES=(0001 0005 0010 0050 0100 0500 1000 5000)


for i in "${QTT_GOROUTINES[@]}";
do
	FILE=$FILE_NAME_START$i$FILE_EXTENSION
	#script
	touch $FILE #guarantees that file exists
	rm $FILE #so it doesnt print unnecessary info
	echo $FILE

	#runs the program the desired number of times
	for j in `seq 1 $MAX_ITERATIONS`;
	do
		start=$( date +"%T" )
		go run etapa1.go 5000 $i
		end=$( date +"%T" )
		echo "-------------------------ITERATION #"$j"-------------------------" >> $FILE
		echo "Start time: "$start >> $FILE
		echo "Finished time: "$end >> $FILE
		echo "Duration: "$(($(date -d $end +%s) - $(date -d $start +%s))) "second(s)" >> $FILE
		echo"" >> $FILE
	done
done