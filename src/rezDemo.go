package main

import (
	"model"
)
import "time"
import "math/rand"

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
)

func main() {
	fmt.Print(model.HelloWord())
	response, err := http.Get("http://ets-res3-1130:ets1130@www2.cooptel.qc.ca/services/temps/?mois=4&cmd=Visualiser")
	defer response.Body.Close()
	if err != nil {

		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {

			fmt.Printf("%s", err)
			os.Exit(1)
		}

		r := bufio.NewReader(bytes.NewReader(contents))

		line, err := r.ReadBytes('\n')

		var wg sync.WaitGroup
		results := make(chan string, 30)
		//mapNoLigneEtChmbr := make(chan map[int]model.TrfParDateParChmbr)
		for i := 0; err != io.EOF; i++ {

			line, err = r.ReadBytes('\n')

			if err != nil {

				fmt.Println(err)

			}
			wg.Add(1)
			//go processText(mapNoLigneEtChmbr, line, results, &wg)
			go processText(line, results, &wg)

			// Wait for the goroutine to finish

		}
		go func() {
			wg.Wait()

			close(results)

		}()

		for element := range results {
			fmt.Println(element)

		}

	}

}

func processText(lineAsByte []byte, results chan string, wg *sync.WaitGroup) {

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	line := string(lineAsByte[:])

	rp := regexp.MustCompile("<TR><TD>([0-9]+)</TD><TD>([0-9]+-[0-9]+-[0-9]+)</TD><TD ALIGN=\x22RIGHT\x22>[ ]+([0-9]+\x2e[0-9]+)</TD><TD ALIGN=\x22RIGHT\x22>[ ]+([0-9]+\x2e[0-9]+)</TD></TR>")

	match := rp.FindStringSubmatch(line)

	if match != nil {

		//var trfParDateParChambre model.TrfParDateParChmbr
		var strConcat string

		for k, v := range match {

			switch k {
			case 1:
				//trfParDateParChambre.
				strConcat = "chambre : " + v + "\n"

				//fmt.Printf("chambre : %d", chambreCourrante)

			case 2:

				strConcat += "date : " + v + "\n"
				//fmt.Printf(" date : %s", date)

			case 3:

				strConcat += "upload : " + v + "\n"
				//fmt.Printf(" upload : %f", bw)

			case 4:

				strConcat += "download : " + v + "\n"
				//fmt.Printf(" download : %s\n", v)

			default:
				strConcat += "??? \n"

			}

		}
		results <- strConcat

		print(line)

	}

	wg.Done()

}
