package main
import (
	"fmt"
	"net/http"
	"io"
)



// we define a struct that will store home page sizes

type HomePageSize struct {
	URL string
	Size int
}



func main() {

	// create a list of URLs
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}



	// create a channel
	results := make(chan HomePageSize)

	for _, url := range urls {

		// start a new goroutine for each URL (so we will be fetching 4 URLs simultaneously)
		
		// notice this is an unnamed function that is immediately invoked. This 
		// is a common pattern with goroutines., but also notice that we
		// defined the function as taking a single parameter (the url).

		go func(url string) {

			// for each URL we make an HTTP request
			res, err := http.Get(url)

			if err != nil {
				panic(err)
			}
		
		defer res.Body.Close()
	
		bs, err := io.ReadAll(res.Body)

		if err != nil {
			panic(err)
		}

		// we store the size of the response body
		results <- HomePageSize{
			URL: url,
			Size: len(bs),
		}

	}(url)

}



// loop 4 times to pull the 4 results off of the channel, compare that result to the 
// current biggest, and swap if its bigger

	var biggest HomePageSize

	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}


	fmt.Println("The biggest home page:", biggest.URL)


}