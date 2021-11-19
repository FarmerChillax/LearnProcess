# -*- coding: utf-8 -*-
'''
    :file: __init__.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/05 15:08:17
'''
import os
from portScan.utils import make_res
import time
from apiflask import APIFlask, input
from celery import Celery
from portScan.settings import redis_url, backend_url
from portScan.scan import PortScan
from portScan.schemas import ScanInSchema

app = APIFlask(__name__)

redis_url=os.getenv('REDIS_URL', None)
backend_url = redis_url
celery = Celery(app.name, broker=redis_url, backend=backend_url)


@celery.task(bind=True)
def scan_task(self, host, all_port=False):

    scanner = PortScan(all_ports=all_port, rate=1024)
    scanner.async_port_scan(host=host)
    return {"result": scanner.open_list}


@app.post("/longtask")
@input(ScanInSchema)
def index(data):
    task = scan_task.delay(data['host'], data['all_port'])
    return {"task_id": task.id}


@app.get("/<string:task_id>")
def status(task_id):
    task = scan_task.AsyncResult(task_id)

    if task.state == 'PENDING':
        return make_res(code=200, message=task.state, data=task.info)
    if task.state == 'SUCCESS':
        return make_res(code=200, message=task.state, data=task.info)
    return {"msg": "err"}, 500
