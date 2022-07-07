from socket import *
import string

SERVER_HOST = "localhost"
SERVER_PORT = 50505

def calcule_weight(height: str, gender: str):
    # Checks if height values are valid.
    try:
        float(height)
    except ValueError:
        print("Invalid height input!")
        return

    # Connect with the server
    client = socket(AF_INET, SOCK_STREAM)
    client.connect((SERVER_HOST, SERVER_PORT))

    # Send request to the server
    client.send(("IWC " + height + " " + gender.upper()).encode())

    # Receive response
    resp = client.recv(1024).decode()

    # Present results
    print("Ideal Weight: " + resp)

    client.close()

if __name__ == "__main__":
    print("-------------------- Ideal Weight Calculator --------------------")
    print("Element format: HEIGHT GENDER(M/F)")
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
        if len(data) != 2:
            print("Invalid input!")
            continue

        # Passes to the processing function
        calcule_weight(data[0], data[1])