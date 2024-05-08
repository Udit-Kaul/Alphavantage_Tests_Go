### The AlphaVantage API's GlobalQuote function retrieves real-time price information about a specific stock symbol. The following test cases verify positive, negative and boundary cases handling 

```
Example API : https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=IBM&apikey=demo

Example Positive Output :

    "Global Quote": {
        "01. symbol": "IBM",
        "02. open": "166.4900",
        "03. high": "166.7600",
        "05. price": "166.2000",
        "06. volume": "6011634",
        "07. latest trading day": "2024-04-30",
        "08. previous close": "167.4300",
        "09. change": "-1.2300",
        "10. change percent": "-0.7346%"
        }
```  
    Test Case 1 – Positive Test with Valid Symbol: 
```
    Test Case Description: Verify that the API returns accurate stock data for a valid symbol.
    Pre-Conditions: API access and a known valid stock symbol (e.g., "AAPL").
    Test Steps:
        Send a request to the GlobalQuote function with the symbol "AAPL".
    Test Data: Symbol: "AAPL"
    Expected Result: The API returns real-time stock data for Apple Inc., including the price, volume, and other relevant information.
```
        
    Test Case 2 – Positive Test with Valid Symbol with Low Activity: 
```
    Test Case Description: Verify that the API returns accurate stock data for a stock symbol known for low trading activity 
    Pre-Conditions: API access and a known valid stock symbol (e.g., "AEAE").
    Test Steps:
        Send a request to the GlobalQuote function with the symbol "AEAE".
    Test Data: Symbol: "AEAE"
    Expected Result: The API returns real-time stock data for Altenergy Acquisition Corp, including the price, volume, and other relevant information.
```
    
    Test Case 3 – Positive Test with Market Open

```
    Test Case Description: Verify that the API returns accurate stock data during market hours.
    Pre-Conditions: API access and market hours.
    Test Steps:
        Send a request to the GlobalQuote function during regular market hours.
    Test Data: Symbol: "AAPL" (or any valid symbol).
    Expected Result: The API returns real-time stock data for the specified symbol, reflecting current market conditions.
```

    Test Case 4 – Boundary Test with High Frequency Requests
```
    Test Case Description: Test the API's behavior when subjected to a high frequency of requests.
    Pre-Conditions: API access and capability to send multiple simultaneous requests.
    Test Steps:
        Rapidly send multiple requests to the GlobalQuote function for the same or different symbols.
    Test Data: Multiple symbols or repeated requests.
    Expected Result: The API responds to each request without delays or errors, maintaining accurate and consistent data retrieval.
```

    Test Case 5 – Negative Test with Invalid Symbol
```
    Test Case Description: Verify that the API handles an invalid symbol gracefully.
    Pre-Conditions: API access.
    Test Steps:
        Send a request to the GlobalQuote function with an invalid symbol (e.g., "XXXX").
    Test Data: Invalid Symbol: "XXXX"
    Expected Result: The API returns an error message indicating that the symbol is invalid or not found.
```
    
    Test Case 6 – Negative Test with Invalid API Key
```
    Test Case Description: Verify how the API handles requests with an invalid API key.
    Pre-Conditions: API access and a known invalid API key.
    Test Steps:
        -Replace the valid API key in the request with an invalid or expired API key.
        -Send a request to the GlobalQuote function with any valid stock symbol (e.g., "AAPL").

    Test Data: Valid stock symbol: "AAPL"
    Expected Result: The API returns an error response indicating that the provided API key is invalid or expired.
```
