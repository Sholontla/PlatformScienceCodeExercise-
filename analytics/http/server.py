from typing import List
from fastapi import WebSocket, Request, WebSocketDisconnect
import websockets

from typing import Dict

from main import app, cache, templates


class ConnectionManager:
    def __init__(self):
        self.active_connections = []
        self.waiting_connections = []

    async def connect(self, websocket: WebSocket):
        if not self.waiting_connections:
            await websocket.accept()
            self.active_connections.append(websocket)
        else:
            waiting_websocket = self.waiting_connections.pop(0)
            waiting_websocket.__dict__.update(websocket.__dict__)
            await self.connect(waiting_websocket)

    def disconnect(self, websocket: WebSocket):
        self.active_connections.remove(websocket)

    async def broadcast(self, data: Dict):
        for connection in self.active_connections:
            await connection.send_json(data)

    async def wait_for_connection(self):
        websocket = WebSocket(scope={}, receive=None, send=None)
        self.waiting_connections.append(websocket)
        await websocket.receive()
        return websocket


manager = ConnectionManager()


@app.websocket("/ws")
async def websocket_endpoint(websocket: WebSocket):
    await manager.connect(websocket)
    try:
        while True:
            data = await websocket.receive_json()
            cache.cache_data(data)
            await manager.broadcast(data)
    except WebSocketDisconnect:
        manager.disconnect(websocket)
    except websockets.exceptions.ConnectionClosedError:
        manager.disconnect(websocket)


@app.get("/")
async def root(request: Request):
    return templates.TemplateResponse("index.html", {"request": request})


@app.get("/profile")
async def root(request: Request):
    return templates.TemplateResponse("./profile.html", {"request": request})


def count_driver_repeats(messages: List[dict]) -> dict:
    driver_counts = {}

    for message in messages:
        driver = message["driver"]
        if driver not in driver_counts:
            driver_counts[driver] = 0
        driver_counts[driver] += 1

    return driver_counts


@ app.get("/drivers")
async def driver(request: Request):
    all_data_by_time = cache.get_all_data()
    print("...........: ", all_data_by_time)

    # process = PlatformScienceCodeExercise(shipments, drivers)
    # process.run()
    driver_counts = count_driver_repeats(all_data_by_time)
    return {"fleet_data": all_data_by_time}


@ app.get("/drivers/all")
async def sales_month():
    all_data_by_time = cache.get_all_data()
    driver_counts = count_driver_repeats(all_data_by_time)
    return driver_counts
