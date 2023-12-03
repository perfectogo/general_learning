#!/bin/bash
 
# to see all the files in the current directory and subdirectories. 
ls -R 

#!
gzip files.tar 

#!
zip filename.zim -r dir_name
zip riddles.zip -r riddles

#!
tar -cf filename.tar dir_name
tar -cf text_files.tar hello.txt bye.txt poem.txt 

tar -xf files.tar 

# options
tar -cjf riddles.tar.bz2 riddles # riddles.tar.bz2
tar -cJf riddles.tar.xz riddles  # riddles.tar.xz


# network
curl -O https://static-assets.codecademy.com/Courses/learn-linux/secret-file.txt o
wget www.codecademy.com
host codecademy.com
ping 104.17.212.81
ifconfig