from peewee import SqliteDatabase, Model, IntegerField, CharField

db = SqliteDatabase("auth.sqlite")

class Auth(Model):
    user_id = IntegerField(null=False, index=True)
    password = CharField(null=False)
    
    def get_user(self): # Get User information from User Microservice
        pass
    
    class Meta:
        database = db

db.create_tables((Auth, ))