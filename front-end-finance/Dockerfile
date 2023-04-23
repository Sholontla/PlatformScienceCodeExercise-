# Use the official lightweight Node.js 16 image as the base image
FROM node:18-alpine

# Set the working directory to /app
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install

# Copy the rest of the application code to the container
COPY . .

# Build the Next.js app for production
# RUN yarn run build

# Expose port 3000 for the app
EXPOSE 3001

# Start the app with npm

CMD ["yarn", "dev"]