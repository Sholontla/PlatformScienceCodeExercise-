import redis
import json


class DataCache:
    def __init__(self, host, port, password, db=0):
        self.host = host
        self.port = port
        self.password = password
        self.redis = redis.Redis(host=self.host, port=self.port,
                                 db=db, password=self.password)

    def cache_data(self, data):
        for record in data:
            # Convert dict to JSON string and set it as a value in Redis with a key equal to the record ID
            record_id = record.get('id')
            if record_id:
                existing_data = self.redis.get(record_id)
                if existing_data:
                    existing_record = json.loads(existing_data)
                    existing_record.update(record)
                    self.redis.set(record_id, json.dumps(existing_record))
                else:
                    self.redis.set(record_id, json.dumps(record))
            else:
                # Handle case where record ID is missing from data
                pass

    def get_data(self, record_ids):
        data = []
        for record_id in record_ids:
            # Get JSON string from Redis and convert it to a dict
            record_data = self.redis.get(record_id)
            if record_data:
                data.append(json.loads(record_data))
        return data

    def get_all_data(self):
        keys = self.redis.keys()
        data = self.redis.mget(keys)
        return [json.loads(record_data) for record_data in data if record_data is not None]

    def delete_all_data(self):
        self.redis.flushall()
