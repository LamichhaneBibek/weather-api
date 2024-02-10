# Weather API

This is a simple Go application that provides weather information using the [WeatherAPI](https://www.weatherapi.com/). It includes a web server and a frontend component written in [Templ](https://github.com/a-h/templ) for generating dynamic HTML.

## Project Structure

- `main.go`: Contains the main entry point of the application.
- `weather.templ`: Templ file for the weather forecast frontend.
- `Makefile`: Makefile for generating Templ code and running the application.
- `go.mod`: Go module file specifying project dependencies.

## Getting Started

1. Install Go on your machine: [https://golang.org/doc/install](https://golang.org/doc/install)
2. Clone this repository:

   ```bash
   git clone https://github.com/LamichhaneBibek/weather-api.git
   ```

3. Navigate to the project directory:

   ```bash
   cd weather-api
   ```

4. Run the following command to generate Templ code and start the application:

   ```bash
   make run
   ```

   This command will generate Templ code and run the web server. Open your browser and go to [http://localhost:8080](http://localhost:8080) to see the weather forecast.

## Project Dependencies

- [Templ](https://github.com/a-h/templ): Templ is used for generating dynamic HTML content in this project.
- [WeatherAPI](https://www.weatherapi.com/): The external API used to fetch weather information.

## Contributing

Feel free to contribute to this project by opening issues or creating pull requests. Your feedback and contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
