#Determines if a string has all unique characters
def isUnique(string):
    characters = {}
    for c in string:
        if characters.get(c) == None:
            characters[c] = 1
        else:
            return False
    return True

string = input("Enter a string:")
print(isUnique(string))