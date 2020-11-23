import os

def run(data):
    s = "go run main.go encrypt.go random.go run.go struct.go -s "
    s += data[1]
    result = os.system(s)
    print(data[0])

a = open("./account.txt", 'r')
for lines in a.readlines():
    data = lines.split(' ')
    run(data)
a.close()

