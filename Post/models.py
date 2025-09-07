from datetime import datetime

from peewee import SqliteDatabase, Model, CharField, IntegerField, TextField, DateField

db = SqliteDatabase('posts.db')

class Post(Model):
    user_id = IntegerField()

    title = CharField(max_length=255)
    content = TextField()
    date = DateField(default=datetime.date.today)

    class Meta:
        database = db

db.create_tables([Post])