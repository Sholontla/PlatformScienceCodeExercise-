FROM python:3.11

WORKDIR /app

COPY local.txt .
RUN pip install -r local.txt

COPY . .

CMD ["python", "http/main.py"]