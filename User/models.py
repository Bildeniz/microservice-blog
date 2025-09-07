from peewee import SqliteDatabase, Model, CharField

db = SqliteDatabase('users.sqlite')

class User(Model):
    name = CharField()
    surname = CharField()
    username = CharField(unique=True)
    email = CharField(unique=True)

    class Meta:
        database = db
        table_name = 'users'

db.create_tables([User])