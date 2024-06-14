package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	NO_OF_THREAD = 1
	API_ENDPOINT = "http://169.47.92.210:30763"
	API_REQUEST  = "tea"
)

func getRequest(thread_no int, wg *sync.WaitGroup) {
	req_time := time.Now()
	URL_STR := API_ENDPOINT + "/" + API_REQUEST

	client := http.Client{}

	req, err := http.NewRequest("GET", URL_STR, nil)
	if err != nil {
		fmt.Printf("TH: [%d] Error: [%s]\n", thread_no, err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("TH: [%d] Error: [%s]\n", thread_no, err)
		return
	}
	resp_time := time.Now()
	defer func(){
		resp.Body.Close()
		wg.Done()
		fmt.Println()
		fmt.Println()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("TH: [%d] Error: [%s]\n", thread_no, err)
	}
	//fmt.Printf("TH: [%d] URL: [%s] Respose Code: [%d]\n", thread_no, req.URL.String(), resp.StatusCode)
	//fmt.Printf("TH: [%d] Response Body: %s\n", thread_no, string(body))
	headers := resp.Header
	timeDiff := resp_time.Sub(req_time)
	fmt.Printf("TH: [%d]\nURL: [%s]\nRespose Code: [%d]\nResponse Body: [%s]\nHeaders: [%s]\nRequest Time: [%s]\nResponse Time: [%s]\nTime Taken: [%s]\n\n\n\n", thread_no, req.URL.String(), resp.StatusCode, string(body), headers, req_time.Format("2006-01-02 3:4:5 pm"), resp_time.Format("2006-01-02 3:4:5 pm"), timeDiff)
}

func main() {
	if len(os.Args[:]) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error: [%s]\n", err)
			os.Exit(1)
		}
		NO_OF_THREAD = i
	}

	wg := new(sync.WaitGroup)

	fmt.Printf("Testing with NO_OF_THREAD: [%d]\n", NO_OF_THREAD)

	for true {
		for i := 0; i < NO_OF_THREAD; i++ {
			wg.Add(1)
			go getRequest(i, wg)
		}
		time.Sleep(10 * time.Second)
		fmt.Println("Next Iteration")
	}
	wg.Wait()
}
