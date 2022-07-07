from socket import *
import string

SERVER_HOST = "localhost"
SERVER_PORT = 50505

def calcule_weight(age: str,):
    # Checks if age values are valid.
    try:
        int(age)
    except ValueError:
        print("Invalid age input!")
        return

    # Connect with the server
    client = socket(AF_INET, SOCK_STREAM)
    client.connect((SERVER_HOST, SERVER_PORT))

    # Send request to the server
    client.send(("SAC " + age).encode())

    # Receive response
    resp = client.recv(1024).decode()

    # Present results
    print("Category: " + resp)

    client.close()

if __name__ == "__main__":
    print("-------------------- Swimmer Age Category --------------------")
    print("Element format: AGE")
    print("Multiple data input: NOT SUPPORTED!")
    print("type \'exit\' to close")

    while True:
        # Collect input data from terminal
        input_message = input("> ")
        input_message = input_message.replace("\n", "")
        input_message = input_message.replace("\r", "")

        if input_message == "exit":
            break

        # Passes to the processing function
        calcule_weight(input_message)