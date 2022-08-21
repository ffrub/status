package cmd

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

func GetStatus(url string, res Response) <-chan error {
    ch := make(chan error) 

    go func() {
        defer close(ch)


        resp, err := http.Get(url)
        if err != nil {
            ch <- err
            return
        }

        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            ch <- err
            return
        }

        err = json.Unmarshal(body, res)
        if err != nil {
            ch <- err
            return
        }
    }()

    return ch
}

func Merge(errChans ...<-chan error) <-chan error {
    mergedChan := make(chan error)

    var wg sync.WaitGroup
    wg.Add(len(errChans))
    go func() {
        // When all of the err errChans are closed, clode the mergedChan
        wg.Wait()
        close(mergedChan)
    }()

    for i := range errChans {
        go func(errChan <-chan error) {
            // Wait for each errChans to close
            for err := range errChan {
                if err != nil {
                    // Fan in contents of each errChan into the mergedChan
                    mergedChan <- err
                }
            }
            wg.Done()
        }(errChans[i])
    }

    return mergedChan
}
