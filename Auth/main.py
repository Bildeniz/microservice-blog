from fastapi import APIRouter

from func import *
from api_models import UserRead, UserCreate, RToken
from models import *

import os

import requests

from fastapi import FastAPI

app = FastAPI()

@app.post("/auth/login", response_model=RToken)
def read_login(user: UserRead):
    username = user.username
    password = user.password

    res = requests.post(
        os.getenv("USER_SERVICE")+"/auth/login",
        json={
            "username": username,
            "communication_key": COMMUNICATION_KEY,
        }
    )

    if res.status_code != 200:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="User is not founded",
        )

    user_json = res.json()

    auth = Auth.select().where(Auth.user_id == user_json["id"]).get_or_none()
    if not pwd_context.verify(password, auth.password):
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Incorrect password",
        )

    data = {
        "auth_id": auth.id,
        "name": user_json['name'],
        "surname": user_json['surname'],
        "email": user_json['email'],
        "user_id": user_json['id'],
        "username": username
    }

    token = create_jwt_token(data)

    return {
        "token": token,
        "token_type": "bearer"
    }


@app.post("/auth/register", response_model=RToken)
async def read_register(user: UserCreate):
    username = user.username
    password = user.password
    name = user.name
    surname = user.surname
    email = user.email

    # Hash password
    encrypted_password = pwd_context.hash(password)

    res = requests.post(
                os.getenv("USER_SERVICE")+"/auth/register", # Signup path of user microservice
                json={
                    "username": username,
                    "name": name,
                    "surname": surname,
                    "email": email,
                }
            )

    if res.status_code != 201:
        raise HTTPException(status_code=status.HTTP_400_BAD_REQUEST, detail=res.json())

    user_id = res.json()["id"]

    auth = Auth(
        user_id = user_id,
        password = encrypted_password,
    )
    auth.save()

    data = {
        "auth_id": auth.id,
        "name": name,
        "surname": surname,
        "email": email,
        "user_id": user_id,
        "username": username
    }

    token = create_jwt_token(data)

    return {
        "token": token,
        "token_type": "bearer"
    }

@app.get("/secret")
def read_secret(user:dict = Depends(verify_token)):
    return "protected value"
