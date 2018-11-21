import csv
import sys
import generate

validPaths = ["user", "application"]
pathError = 'need to choose a valid message type. user or application'
prompts = {
    "application": {
        "code": "http code if relevant, otherwise put none. \n",
        "name": "descriptive name \n",
        "message": "message\n"
    },
    "user": {
        "name": "descriptive name of error for developers. Not to going to be shown to users.\n",
        "header": "brief description of error\n",
        "message": "tell the user what they can do about the error\n",
    }
}
def add(path):
    print("Adding to "+ path +" messages")
    objs = generate.csvCopy(path)
    headers = objs[0].keys()
    newObj = {}
    if 'id' in headers:
        newObj['id'] = len(objs)
        headers.remove('id')
    for ele in headers:
        p = prompts[path][ele]
        val = raw_input(ele+": "+p)
        newObj[ele] = val
    print(newObj)

def main():
    if len(sys.argv) > 1:
        path = sys.argv[1]
        if path in validPaths:
            add(path)
        else:
            print(pathError)

    else:
        print(pathError)


if __name__== "__main__":
    main()
