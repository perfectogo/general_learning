#!/bin/bash

let num=$1
current_datetime=$(date +"%Y-%m-%d %T")
echo current_datetime

for number in {1..5}
do
    if [ $2 == 'd' ]
    then 
        mkdir $3_$number_$current_datetime
    elif [ $2 == 'f' ]
    then
        touch $3_$number_$current_datetime.txt
    else
        echo thms is wrong
    fi
done

for file in *
do
    echo "File: $file"
    # Additional actions on each file
done
