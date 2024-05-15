# Go GET with Multiple thread

This code is an improved version of the goGet.go one, designed to make multiple concurrent GET requests to an API endpoint. Here's a breakdown:

1. Packages:

    Same as the previous code (fmt, net/http, os, time, io/ioutil, sync, strconv)

2. Variables:

    NO_OF_THREAD: This variable defines the number of concurrent requests to make (default 1, can be overridden with a command line argument).
    API_ENDPOINT: Stores the base URL of the API server (http://169.47.92.210:30763).
    API_REQUEST: Stores the specific API request path (tea).

3. getRequest function:

    This function takes two arguments:
        thread_no (int): Identifies the current thread number.
        wg (*sync.WaitGroup): A WaitGroup object used for synchronization.
    It performs the following actions:
        Records the request time using time.Now().
        Constructs the full URL by combining API_ENDPOINT and API_REQUEST.
        Creates an HTTP client and a new GET request object similar to the previous code.
        Handles potential errors during request creation.
        Sends the request and stores the response in resp.
        Records the response time using time.Now().
        Uses a defer statement to ensure the response body is closed and the WaitGroup counter is decremented even if errors occur.
        Reads the response body into a byte array (body).
        Prints detailed information about the request and response:
            Thread number
            URL
            Response code
            Response body (as a string)
            Response headers
            Request time (formatted)
            Response time (formatted)
            Time taken for the request (calculated from request and response times)
        Prints extra newlines for better readability.

4. main function:

    Checks if there's a command-line argument provided.
        If yes, it tries to convert it to an integer using strconv.Atoi.
        If conversion fails, it prints an error and exits.
        Otherwise, it updates the NO_OF_THREAD variable with the provided value.
    Creates a new sync.WaitGroup object wg.
    Prints the number of threads being used.
    Enters an infinite loop:
        Creates NO_OF_THREAD concurrent goroutines using go.
            Each goroutine calls getRequest with its thread number and the WaitGroup.
        Waits for 10 seconds using time.Sleep.
        Prints a message indicating the next iteration.
    Waits for all goroutines to finish using wg.Wait().

In summary:

This code defines a function to make a single GET request and print detailed information about it. The main function allows specifying the number of concurrent requests via a command line argument. It then creates multiple goroutines to send these requests simultaneously and waits for them to complete. This approach enables fetching data from the API endpoint much faster compared to sending requests sequentially.