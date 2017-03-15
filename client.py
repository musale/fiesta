#!/usr/bin/python2.7

import urllib
import urllib2
from datetime import datetime

url = "http://127.0.0.1:8018/"


def get_report(idx=None):
    st = str(datetime.now())[:10] + " 00:00:00"
    sto = str(datetime.now())[:10] + " 24:59:59"
    payload = {
        'start': st, 'mail': True, 'stop': sto
    }
    return urllib2.urlopen(url + 'range', urllib.urlencode(payload)).read()


if __name__ == '__main__':
    print get_report()
