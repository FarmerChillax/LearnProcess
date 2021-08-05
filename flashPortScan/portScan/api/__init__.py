# -*- coding: utf-8 -*-
'''
    :file: __init__.py
    :author: -Farmer
    :url: https://blog.farmer233.top
    :date: 2021/08/05 15:08:17
'''

import time
from apiflask import APIFlask
from celery import Celery
from api.settings import redis_url, backend_url
from scan.portScan import PortScan

app = APIFlask(__name__)
celery = Celery(app.name, broker=redis_url, backend=backend_url)

@celery.task(bind=True)
def scan_task(self, arg=1, arg2=2):
    STATUS = ["active", "finish"]
    total = 5
    for i in range(total):
        self.update_state(state="PROGRESS", meta={"total": i})
        time.sleep(i)
    
    return arg+arg2

@app.post("/longtask")
def index():
    task = scan_task.apply_async()
    return {"task_id": task.id}

@app.get("/<string:task_id>")
def status(task_id):
    task = scan_task.AsyncResult(task_id)
    print(task.state)
    # if task.state == "PENDING":
    #     print("未开始")
    print(task, type(task))

    return {"state": task.state, "status": str(task.info)}