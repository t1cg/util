import csv
import sys

def goMessages(objs, structType):
    messageString = "var Messages = struct { \n"
    # initializing all objects
    for obj in objs:
        messageString += "\t"+ obj['name'] +"\t"+ structType +"\n"

    messageString += "}{\n"
    # defining them
    for obj in objs:
        messageString += "\t"+ obj['name'] +":\t"+ structType +"{"
        for key in obj.keys():
            messageString += "\""+obj[key]+"\", "
        messageString += "},\n"
    messageString += "}"
    return messageString

# converts obj to a golang struct
def goStruct(objs, structType):
    # actually define the new message type
    structString = "type "+ structType + " struct { \n"

    # need all the keys for this message type. can use any
    # object for this so I just got the first one.
    for key in objs[0].keys():
        uKey = key[0].upper() + key[1:]
        structString += "\t"+uKey+"\tstring\n"

    structString += "}\n\n"
    return structString

def go(objs, path):
    f = open("./"+ path +"/"+ path +".go", "w+")
    # define package
    f.write("package "+ path + "\n\n")

    # define struct specific to this type of message (need to capitalize the
    # first letter in order to export.)
    structType = path[0].upper() + path[1:] + "Message"
    f.write(goStruct(objs, structType))
    f.write(goMessages(objs, structType))
    f.close()
    print(objs)

def reactObjToString(obj):
    stringObj = "{\n"
    for key in obj.keys():
        stringObj += "\t\""+ key +"\": \""+obj[key]+"\"\n"
    stringObj += "}\n"
    return stringObj

def react(objs, path):
    f = open("./"+ path +"/"+ path +".js", "w+")
    for obj in objs:
        f.write("export const " + obj['name'] + " = " + reactObjToString(obj) + "\n")
    f.close()
    print(obj)

def csvCopy(path):
    obj = []
    with open('./' + path + '/' + path + '.csv') as csvfile:
        reader = csv.DictReader(csvfile, delimiter=',')
        for row in reader:
            rowObj = {}
            for key in row.keys():
                rowObj[key] = row[key]
            obj.append(rowObj)
    return obj

def main():
    functions = [go, react]
    messageTypes = ['user', 'application']
    csv = {}

    if len(sys.argv) > 1 :
        # TODO: only generate code for specified messages
        print(sys.argv)

    for m in messageTypes:
        print("==== " + m + " ====")
        obj = csvCopy(m)
        for f in functions:
            print("\t==== "+ f.__name__ + " ====")
            f(obj, m)
            print("\t==== done with "+ f.__name__  +" ====")
        print("==== done with "+ m +" ====")

if __name__== "__main__":
    main()
