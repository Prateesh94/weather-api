# weather-api

A simple RESTful API that provides weather data for cities using the Visual Crossing Weather API, built in Go. The API implements caching and rate limiting for efficient and fair usage.

## Features

- **Current weather endpoint**: Get weather data for any city.
- **Caching**: Weather results for each city are cached for 15 minutes to reduce external API calls and improve speed.
- **Rate limiting**: Each client IP is limited to 50 requests per minute.
- **Environment-based API key**: The Visual Crossing API key is loaded from a local JSON file.

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Prateesh94/weather-api.git
   cd weather-api
   ```

2. **Install dependencies**
   - Go 1.18+ required
   - The following Go packages must be installed:
     ```
     go get github.com/gorilla/mux
     go get github.com/patrickmn/go-cache
     go get golang.org/x/time/rate
     ```

3. **Create API key file**
   - Create a file called `env.json` in the root directory with the following content:
     ```json
     {
       "key": "YOUR_VISUAL_CROSSING_API_KEY"
     }
     ```

## Usage

1. **Run the server**
   ```bash
   go run main.go
   ```

2. **API Endpoint**

   - **GET /add?city={city_name}**
     - Returns the weather data for the specified city.
     - Example:
       ```
       http://localhost:5050/add?city=London
       ```

     - Response: JSON weather data (either from cache or fetched from Visual Crossing)

## How It Works

- **Weather Fetching**: When `/add?city={city}` is called, the API first checks if the city’s weather is cached. If not, it fetches from Visual Crossing, caches the result, and returns it.
- **Rate Limiting**: Requests are limited per client IP to avoid abuse.
- **Environment Key**: The API key is loaded at startup from `env.json`.

## Project Structure

```
main.go           # Main application and server logic
cacher/cache.go   # Implements caching and rate limiting
env.json          # (Created by user) Stores the API key
```

## License

MIT License

## Author

[Prateesh94](https://github.com/Prateesh94)
