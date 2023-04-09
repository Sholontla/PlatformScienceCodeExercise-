import json
from contextlib import contextmanager


class Config:
    def __init__(self):
        pass

    @contextmanager
    def save_file(self, filename):
        try:
            file = open(filename, 'w')
            yield file
        finally:
            file.close()

    @contextmanager
    def read_config_file(self, file_path):
        with open(file_path, 'r') as f:
            yield json.load(f)


def ws_config(filepath):
    config = Config()
    with config.read_config_file(filepath) as config:
        return config["ws_server"]


def redis_config(filepath):
    config = Config()
    with config.read_config_file(filepath) as config:
        return config["redis"]
