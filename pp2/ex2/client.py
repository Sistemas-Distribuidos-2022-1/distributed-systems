from socket import *
import string

SERVER_HOST = "localhost"
SERVER_PORT = 50505

def majority(name: str, gender: str, age: str):
    # Connect with the server
    client = socket(AF_INET, SOCK_STREAM)
    client.connect((SERVER_HOST, SERVER_PORT))

    # Send request for readjustment
    client.send(("MAJ " + name + " " + gender.upper() + " " + age).encode())

    # Receive response
    resp = client.recv(1024).decode()

    # Present results
    if resp == "TRUE":
        print(name + " has reached the age of majority.")
    elif resp == "FALSE":
        print(name + " has not reached the age of majority.")
    else:
        print("It was not possible to check the age of majority for " + name + ", " + gender + ", " + age)

    client.close()

if __name__ == "__main__":
    print("-------------------- Majority Age Check --------------------")
    print("Element format: NAME GENDER(M/F) AGE")
    print("Multiple data input: NAME1 GENDER1 AGE1 NAME2 GENDER2 AGE2 ...")
    print("type \'exit\' to close")

    while True:
        # Collect input data from terminal
        input_message = input("> ")
        input_message = input_message.replace("\n", "")
        input_message = input_message.replace("\r", "")

        if input_message == "exit":
            break

        # Check that the input string has the correct format
        data = input_message.split(" ")
        if len(data)%3 != 0:
            print("Invalid input format")
            continue
        
        # Counts the number of elements present in the input data.
        numPeoples = len(data)//3
        for i in range(numPeoples):
            name = data[i*3]
            gender = data[i*3+1]
            age = data[i*3+2]

            # Checks if age value is valid.
            try:
                int(age)
            except ValueError:
                print("Invalid age input!")
                continue

            # Passes to the processing function
            majority(name, gender, age)