from flask import request, jsonify

from app import app
from app.mails import send_mail
from app.models import write_db_email

@app.route('/mail', methods=['POST'])
def view_mail():
    print('send...')
    req_data = request.get_json()
    thr = send_mail(subject=req_data.get('subject'), to=req_data.get('user_email'), body=req_data.get('body'))
    print(thr)
    write_db_email(req_data)
    return jsonify(msg='ok')


    # if req_data.get("user") == None:
    #     message = send_mail(subject=req_data.get('subject'), to=req_data.get('user_email'), body=req_data.get('body'))
    # else:
    #     message = send_mail(subject=req_data.get('subject'), to=req_data.get('user') + " " + req_data.get('user_email'), body=req_data.get('body'))
