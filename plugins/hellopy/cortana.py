#!/usr/bin/env python
# coding: utf8
from bottle import route, run, request
import json
import requests


plug_host = plug_port = None

@route('/register', method='POST')
def reg():
    plug_data = request.json
    print "Registered: %s at %s:%s" % (plug_data['plug'], plug_data['host'], plug_data['port'])
    global plug_host, plug_port
    plug_host = plug_data['host']
    plug_port = plug_data['port']
    print "Pinging"
    response = _plug('ping', {'ping': plug_data['plug']})
    if 'pong' in response and response['pong'] == plug_data['plug']:
        print "Pong OK"
    else:
        print "Pong ERROR, got: %s" % response
        
    print "Calling process"
    response = _plug('process', {'data': "Florin"})
    print "Got: %s" % response['result']
    
    
def _plug(action, data):
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
    return requests.post("http://%s:%d/%s" % (plug_host, plug_port, action), data=json.dumps(data), headers=headers).json()

if __name__ == '__main__':
    run(host='localhost', port=8000, debug=True)
