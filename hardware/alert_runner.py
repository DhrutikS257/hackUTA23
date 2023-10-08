import ctypes
import requests
import json

my_c_lib = ctypes.CDLL("./alert_handling.dll")

url = 'http://localhost:8080/insert'

def helpMe(str):
    # Convert Python string to a bytearray
    buffer = ctypes.create_string_buffer(len(str.encode('utf-8')) + 50)  # Allocate enough space
    
    buffer.value = str.encode('utf-8')

    # Pass the bytearray to the C function
    my_c_lib.getAlert(buffer)
    
    # Convert the modified bytearray back to a string
    return buffer.value.decode('utf-8')

headers = {"Content-Type": "application/json"}


# data = {'sensor':'a Sonic triggered','time':'Sun Oct 08 08:39:19 2023','alert':'warning','email':'josego0716@gmail.com'}

# for i in range(4):
#     data[thing] = otherthing

# response = requests.post(url, json=data, headers=headers)


while True:
    data = "\0"*256

    data = helpMe(data)
    data = data.split("'")
    json_obj = {data[0]:data[1],data[2]:data[3],data[4]:data[5],data[6]:data[7]}
    # json_obj = json.loads("\"" + data + "\"")
    # print(json_obj)
    response = requests.post(url, json=json_obj, headers=headers)
    print(response)
