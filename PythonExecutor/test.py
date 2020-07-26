import json
import sys


parameters = None

if len(sys.argv) > 0:
    count = 1
    for i in sys.argv:
        if count == 1:
            continue
        
    parameters = json.loads(i)

    count += 1


# print(parameters)


try:

    item = {
        'success': True,
        'message' :'This is a test',
        'result' : {
            'subitem':'some more data',
            'second':parameters[0]['Value']
        }
    }




    print(json.dumps(item))

except Exception as e:
    print(e)