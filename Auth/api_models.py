from pydantic import BaseModel, EmailStr

#################################
#########Request Models##########

class UserRead(BaseModel):
    username: str
    password: str

class UserCreate(BaseModel):
    username: str
    password: str
    name : str
    surname: str
    email: EmailStr



#################################
####### Response Model ##########

class RToken(BaseModel):
    token: str
    token_type: str = "bearer"