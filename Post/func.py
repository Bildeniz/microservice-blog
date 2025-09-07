from jose import jwt, JWTError

from fastapi import HTTPException, status, Depends, Header

from datetime import datetime

from dotenv import load_dotenv

ALGORITHM = "RS256"

load_dotenv()

with open("public_key.pem", "r") as f:
    PUBLIC_KEY = f.read()

with open('communication_key', 'r') as f:
    COMMUNICATION_KEY = f.read()


async def get_token_from_header(authorization: str = Header(...)) -> str:
    if not authorization.startswith("Bearer "):
        raise HTTPException(status_code=401, detail="Invalid authorization header")
    return authorization[7:]

def verify_token(token: str = Depends(get_token_from_header)):
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
        raise HTTPException(status_code=401, detail=f"Unvalid Token {e}")

def check_token(token: str = Depends(get_token_from_header)):
    pass