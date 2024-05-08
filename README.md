# AlphaVantage Test Automation Framework

This test automation framework is designed to test the AlphaVantage API endpoints using Jest, Babel, and Jest-Cucumber.

## Installation

1. Clone the repository:
    ```
    git clone https://github.com/Udit-Kaul/Alpha_Vantage_Tests.git
    ```
2. Install Node Modules
   ```
   go build ./...
    ```
## Configuration
1. Obtain an API key from AlphaVantage : https://www.alphavantage.co/support/#api-key
2. Set the API key as an environment variable. Create a .env file in the root directory
```
   API_KEY= <Your API Key>
```
## Running Tests

To run the test suite, execute the following command:
```
go test ./...
```



## Manual Test Cases

The manual test cases can be found in the ManualTestCases.md file in the root directory