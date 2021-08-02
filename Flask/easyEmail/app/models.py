from app import db

class Mail(db.Model):
    id = db.Column(db.Integer,  primary_key=True)
    recipient = db.Column(db.String(50), index = True)
    recipient_email = db.Column(db.String(50))
    subject = db.Column(db.String(50))
    body = db.Column(db.Text)