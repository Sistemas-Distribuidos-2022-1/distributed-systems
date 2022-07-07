from socket import *
import string

SERVER_HOST = "localhost"
SERVER_PORT = 50505

def calculate_salary(name: str, level: str, salary: str, num_deps: str):
    # Checks if values are valid.
    try:
        int(num_deps)
        float(salary)
    except ValueError:
        print("Invalid input!")
        return

    # Connect with the server
    client = socket(AF_INET, SOCK_STREAM)
    client.connect((SERVER_HOST, SERVER_PORT))

    # Send request for readjustment
    client.send(("CNS " + level.upper() + " " + salary + " " + num_deps).encode())

    # Receive response
    resp = client.recv(1024).decode()

    # Present results
    print("NAME: " + name + " | LEVEL: " + level + " | NET_SALARY: " + resp)

    client.close()

if __name__ == "__main__":
    print("-------------------- Net Salary Calculator --------------------")
    print("Element format: NAME LEVEL SALARY NUM_DEPS")
    print("Multiple data input: NAME1 LEVEL1 SALARY1 NUM_DEPS1 NAME2 LEVEL2 SALARY2 NUM_DEPS2 ...")
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
        if len(data)%4 != 0:
            print("Invalid input format")
            continue
        
        # Counts the number of elements present in the input data.
        numEmployees = len(data)//4
        for i in range(numEmployees):
            # Passes to the processing function
            calculate_salary(data[i*3], data[i*3+1], data[i*3+2], data[i*3+3])