from socket import *
import string

SERVER_HOST = "localhost"
SERVER_PORT = 50505

def majority(n1: str, n2: str, n3: str):
    # Checks if grade values are valid.
    try:
        int(n1)
        int(n2)
        int(n3)
    except ValueError:
        print("Invalid grade input!")
        return

    # Connect with the server
    client = socket(AF_INET, SOCK_STREAM)
    client.connect((SERVER_HOST, SERVER_PORT))

    # Send request for readjustment
    client.send(("APC " + n1 + " " + n2 + " " + n3).encode())

    # Receive response
    resp = client.recv(1024).decode()

    # Present results
    if resp == "TRUE":
        print("Approved!")
    elif resp == "FALSE":
        print("Reproved!")
    else:
        print("It was not possible to check approval for grades " + n1 + ", " + n2 + ", " + n3)

    client.close()

if __name__ == "__main__":
    print("-------------------- Grade Check --------------------")
    print("Element format: N1 N2 N3")
    print("Multiple data input: NOT SUPPORTED!")
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
        if len(data) != 3:
            print("Invalid input format")
            continue

        # Passes to the processing function
        majority(data[0], data[1], data[2])