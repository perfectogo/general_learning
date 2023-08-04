import os

def clear_screen():
   
  #  os.system('cls' if os.name == 'nt' else 'clear')
    print(os.cpu_count())

# Test the function
print("This is some content on the screen.")
input("Press Enter to clear the screen.")
clear_screen()