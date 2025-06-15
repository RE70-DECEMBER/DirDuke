import os

print("Welcome To DirDuke Setup!\n\n1. Make Work Outside Of DIR\n2. Quit\n")

iN = input("enter choice: ")

if iN == "1" or iN == "one" or iN == "One" or iN == "y":
	print("\nCopying DirDuke to --> /usr/local/bin")
	os.system("sudo cp DirDuke /usr/local/bin")
	print("Script Should Work Now OutSide Of Script DIR!\nTry Typing DirDuke in home DIR")
else:
	print("\nScript Exited!")
