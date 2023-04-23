FROM loadimpact/k6

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

CMD ["k6", "run", "--config", "/k6config.js", "/test.js"]