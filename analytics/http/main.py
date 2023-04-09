import os
import uvicorn
from fastapi.staticfiles import StaticFiles
from fastapi.templating import Jinja2Templates
from fastapi import FastAPI
from config import ws_config, redis_config

from data_cache import DataCache

path = "http/config.json"
ws = ws_config(path)
r = redis_config(path)

host = ws["host"]
port = ws["port"]

redis_host = r["host"]
redis_port = r["port"]
redis_password = r["password"]

app = FastAPI()

cache = DataCache(redis_host, redis_port, redis_password)

script_dir = os.path.dirname(__file__)
st_abs_file_paath = os.path.join(script_dir, "static/")
app.mount("/static", StaticFiles(directory=st_abs_file_paath), name="static")
templates = Jinja2Templates(directory="http/templates")


if __name__ == "__main__":

    uvicorn.run("server:app",
                host=host, port=port, reload=True)
