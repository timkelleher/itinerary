# Itinerary

## Background
Itinerary is a development tool which provides quick way to stand up an API whose endpoints and payloads can be defined from a GUI.

Two common use cases for Itinerary:
- You're in a hurry and you need to develop against an API you don't yet have access to (yes, it happens...)
- End-to-end testing where you can easily provide mock data from a real service

## Running
Run Itinerary via docker compose:

```
docker compose up -d
```

Visit localhost:3000 to start using Itinerary.

## Usage
Create a Collection.  The Collection is a set of API Endpoints grouped by a single path.

Then create an Endpoint.  You can define an HTTP method, and a payload for this Endpoint.

## Development
For local development, you can run the backend and frontend separately.  You'll also want to fire up a MySQL/mariadb server for the backend to connect to.  Use ./localDatabase.sh to run a standalone mariadb instance in Docker.

Run ./localBackend.sh to build and run a local copy of the backend on port 3000.

Run ./localFrontend.sh to build and run a local copy of the frontend on port 8080.  Hit that with your browser to bring up the application.

### Backend
```
cd backend
go run cmd/server/main.go
```

Visit localhost:3000 to hit the backend API.

Via curl:
```
curl localhost:3000/health
```

Via wget:
```
wget -nv -O - localhost:3000/health
```

Via httpie:
```
http localhost:3000/health
```

### Frontend

Itinerary runs Vue 3.  You will need vue-cli to run the frontend.

```
brew install vue-cli
```

After vue-clie is installed, you can start the frontend service:

```
cd frontend
npm install
npm run serve
```

Visit localhost:8080 to view the frontend.

# Roadmap

Future features planned:
- Request input matching
- Response Delays/Jitter
- Data import & export
- Authentication simulation
- Analytics & stress test simulation