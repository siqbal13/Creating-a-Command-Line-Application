# CSV to JSON Lines Converter

This is a command-line application written in Go to convert a comma-delimited CSV file to a JSON lines file. It takes input in the form of a CSV file and produces output in the form of a JSON lines file.

## Usage

To run the application, use the following command:

```
csvtojl input.csv output.jl
```

Where `input.csv` is the path to the input CSV file and `output.jl` is the path to the output JSON lines file.

## Installation

1. Make sure you have Go installed on your system. If not, you can download it from the [official Go website](https://golang.org/).
2. Clone this repository to your local machine:

```
git clone https://github.com/yourusername/csv-to-json-lines.git
```

3. Navigate to the project directory:

```
cd csv-to-json-lines
```

4. Build the application using the following command:

```
go build
```

This will generate an executable file named `csvtojl` in the current directory.

## Testing

To run the unit tests for the application, use the following command:

```
go test
```

This will execute all the test functions in the project and display the results.

## Usage Notes

- The input CSV file should have the following headers: "value", "income", "age", "rooms", "bedrooms", "pop", "hh".
- The output JSON lines file will contain one JSON object per line, corresponding to each row in the input CSV file.

## Author

[Your Name](https://github.com/yourusername)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

