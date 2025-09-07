from pydantic import BaseModel


############################
###### Request Models ######
class UserCreate(BaseModel):
    username: str
    name: str
    surname: str
    email: str


class UserRead(BaseModel):
    username: str
    communication_key: str


#########################
#### Response Models ####

class RUser(BaseModel):
    id: int
    username: str
    name: str
    surname: str
    email: str