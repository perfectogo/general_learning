#!/bin/bash

select fruit in Apple Orange Banana Exit 
do
    case $fruit in
        Apple)
            echo "You selected Apple."
            ;;
        Orange)
            echo "You selected Orange."
            ;;
        Banana)
            echo "You selected Banana."
            ;;
        Exit)
            echo "Exiting the program."
            break
            ;;
        *)
            echo "Invalid option. Please try again."
            ;;
    esac
done


1 berilgan argumentlarga file va directoriya

2 file_name from words with date

3 permissions 777

