#!/usr/bin/env python
# coding: utf8
from bottle import route, run, request
import socket
from multiprocessing import Process
import json
import requests

@route('/ping', method=['POST'])    
def ping():
    return json.dumps({'pong': request.json['ping']})
    
@route('/process', method=['POST'])        
def process():
    return json.dumps({'result': "Hello, %s" % request.json['data']})
    
def _cortana(action, data):
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
    return requests.post(cortana_url % action, data=json.dumps(data), headers=headers)
    
def _start_bottle(port):
    """ Wrapper for Bottle.run """
    try:
        run(host='localhost', port=port, debug=True)
    except KeyboardInterrupt:
        print "Keyboard interrupt, exiting"
    
def _get_free_port():
    """ Allocate random port, read its value, then close socket and return port """
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.bind(('localhost', 0))
    port = sock.getsockname()[1]
    sock.close()
    return port
        
if __name__ == '__main__':
    cortana_url = "http://localhost:8000/%s"
    # open port
    port=_get_free_port()
    p=Process(target=_start_bottle, args=(port,))
    p.start()
    # register
    _cortana('register', {'plug':'bottlepy', 'host': 'localhost', 'port': port})
    try:
        p.join()
    except KeyboardInterrupt:
        print "Keyboard interrupt, exiting"
