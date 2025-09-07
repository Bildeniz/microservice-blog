from fastapi import HTTPException, status, Depends
from fastapi.security import oauth2, OAuth2PasswordBearer

from jose import jwt, JWTError

from pathlib import Path

from datetime import timedelta, datetime

from dotenv import load_dotenv

from passlib.context import CryptContext

# Configurations
PRIVATE_KEY_PATH = Path("private_key.pem")
PUBLIC_KEY_PATH = Path("public_key.pem")

TOKEN_EXP_DELTA = timedelta(days=30)

with open(PRIVATE_KEY_PATH, "r") as f:
    PRIVATE_KEY = f.read()

with open(PUBLIC_KEY_PATH, "r") as f:
    PUBLIC_KEY = f.read()


ALGORITHM = "RS256"

with open('communication_key', 'r') as file:
    COMMUNICATION_KEY = file.read()

load_dotenv()

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="/auth/login")

pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

def verify_token(token: str = Depends(oauth2_scheme)):
    try:
        payload = jwt.decode(token, PUBLIC_KEY, algorithms=[ALGORITHM])
        if not payload.get("user_id"):
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Unauthorized",
            )
        if datetime.fromtimestamp(payload.get("exp")) <= datetime.now():
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Expired token",
            )

        return payload
    except JWTError as e:
        raise HTTPException(status_code=401, detail="Unvalid Token")

def create_jwt_token(data: dict, expires_delta: timedelta = TOKEN_EXP_DELTA):
    expire = (datetime.now() + expires_delta).timestamp()
    data.update({"exp": expire})
    return jwt.encode(data, PRIVATE_KEY, algorithm=ALGORITHM)


