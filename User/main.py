from fastapi import FastAPI, HTTPException, status, Depends

from api_models import UserCreate, UserRead, RUser
from models import *
from func import COMMUNICATION_KEY, verify_token
app = FastAPI()


#################################
#Endpoints for auth microservice#
@app.post("/auth/register", status_code=status.HTTP_201_CREATED, response_model=RUser)
def create_register(user: UserCreate):
    username = user.username
    name = user.name
    surname = user.surname
    email = user.email

    # if user already registered raise error
    if user:= User.select().where(User.username == username).get_or_none():
        raise HTTPException(status_code=status.HTTP_400_BAD_REQUEST,
                             detail=f"Username already registered as {user.email}")

    if user := User.select().where(User.email == email).get_or_none():
        raise HTTPException(status_code=status.HTTP_400_BAD_REQUEST,
                             detail=f"email already registered as {user.username}")

    user = User(
        username=username,
        name=name,
        surname=surname,
        email=email,
    )

    user.save()

    return {
        "id": user.id,
        "username": user.username,
        "name": user.name,
        "surname": user.surname,
        "email": user.email,
    }

@app.post("/auth/login", response_model=RUser)
def read_login(user: UserRead):
    username = user.username
    communication_key = user.communication_key

    if communication_key != COMMUNICATION_KEY:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Communication key is invalid"
        )

    user = User.select().where(User.username == username).get_or_none()
    if not user:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="User not found")

    return {
        "id": user.id,
        "username": user.username,
        "name": user.name,
        "surname": user.surname,
        "email": user.email,
    }

###########################
##### User Endpoints ######
@app.post("/user/info")
def read_user_info(user: dict = Depends(verify_token)):
    return user