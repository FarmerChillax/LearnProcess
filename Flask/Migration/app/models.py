from collections import UserList
from enum import unique
from operator import index
from os import name
from sqlalchemy.orm import backref, relation
from sqlalchemy.sql.elements import True_

from sqlalchemy.sql.schema import ForeignKey, PrimaryKeyConstraint
from app import db

# 一对多 model demo
class Author(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(20), unique = True)
    phone = db.Column(db.String(20))
    # 一个author有多篇文章
    articles = db.relationship('Article')

class Article(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    title = db.Column(db.String(50), index = True)
    body = db.Column(db.Text)
    # articles外键, 关联表
    author_id = db.Column(db.Integer, db.ForeignKey("author.id"))

# Base on 一对多关系的双向关系
class Writer(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(64), unique = True)
    books = db.relationship('Book', back_populates = "writer")

class Book(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(50), index = True)
    writer_id = db.Column(db.Integer, db.ForeignKey('writer.id'))
    writer = db.relationship('Writer', back_populates = "books")

# use backref to simple relationship
# class Singer(db.Model):
#     id = db.Column(db.Integer, primary_key = True)
#     name = db.Column(db.String(70), unique = True)
#     songs = db.relationship('Song', backref='singer')

# class Song(db.Model):
#     id = db.Column(db.Integer, primary_key = True)
#     name = db.Column(db.String(50), index = True)
#     singer_id = db.Column(db.Integer, db.ForeignKey('singer.id'))
#     songs = db.relationship('Song', backref=backref('singer', uselist=False))

# 多对一
class Citizen(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(30), unique = True)
    city_id = db.Column(db.Integer, db.ForeignKey('city.id'))
    city = db.relationship('City')

class City(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(30), unique = True)


# one by one
class Country(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(30), unique = True)
    capital = db.relationship('Capital', uselist=False)

class Capital(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(30), unique = True)
    country_id = db.Column(db.Integer, db.ForeignKey('country.id'))
    country = db.relationship('Country')

# 多对多
# 外键只能存储一个记录，但多对多模型需要存储一组外键，所以还需要一张关联表，这张表不存储数据，只用来存储关系两侧模型的外键对应关系
association_tabel = db.Table('association', 
    db.Column('student_id', db.Integer, db.ForeignKey('student.id')),
    db.Column('teacher_id', db.Integer, db.ForeignKey('teacher.id'))
    )

class Student(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(70), unique = True)
    grade = db.Column(db.String(20))
    teachers = db.relationship('Teacher',
        secondary = association_tabel,
        back_populates='students')

class Teacher(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(70), unique = True)
    office = db.Column(db.String(20))
    students = db.relationship('Student',
        secondary = association_tabel,
        back_populates='teachers')