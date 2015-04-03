package main

import (
	"time"

	"github.com/mikefaille/rezDemo/model"
)

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
	"strconv"
	"sync"
)

func main() {

	const dateLayout = "2006-01-02"

	response, err := http.Get("http://ets-res3-1130:ets1130@www2.cooptel.qc.ca/services/temps/?mois=3&cmd=Visualiser")

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
		results := make(chan model.Chambre, 30)
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
			fmt.Printf("chambre : %d \n", element.ChambreNo)
			fmt.Println("date : " + element.Date.Format(dateLayout))
			fmt.Printf("upload : %0.2f \n", element.Upload)
			fmt.Printf("download : %0.2f \n", element.Download)
			fmt.Println()

		}

	}

}

func processText(lineAsByte []byte, results chan model.Chambre, wg *sync.WaitGroup) {
	const dateLayout = "2006-01-02"
	var chambre model.Chambre
	isRandomDelay := false
	if isRandomDelay {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}

	line := string(lineAsByte[:])

	rp := regexp.MustCompile("<TR><TD>([0-9]+)</TD><TD>([0-9]+-[0-9]+-[0-9]+)</TD><TD ALIGN=\x22RIGHT\x22>[ ]+([0-9]+\x2e[0-9]+)</TD><TD ALIGN=\x22RIGHT\x22>[ ]+([0-9]+\x2e[0-9]+)</TD></TR>")

	match := rp.FindStringSubmatch(line)

	if match != nil {

		//var trfParDateParmodel.Chambre model.TrfParDateParChmbr
		var strConcat string

		for k, v := range match {

			switch k {
			case 1:
				//trfParDateParmodel.Chambre.
				strConcat = "chambre : " + v + "\n"
				chambreNo, _ := strconv.ParseInt(v, 0, 64)
				chambre.ChambreNo = chambreNo

				//fmt.Printf("chambre : %d", chambreCourrante)

			case 2:

				strConcat += "date : " + v + "\n"
				chambre.Date, _ = time.Parse(dateLayout, v)
				//fmt.Printf(" date : %s", date)

			case 3:

				strConcat += "upload : " + v + "\n"
				chambre.Upload, _ = strconv.ParseFloat(v, 32)
				//fmt.Printf(" upload : %f", bw)

			case 4:

				strConcat += "download : " + v + "\n"
				chambre.Download, _ = strconv.ParseFloat(v, 32)
				//fmt.Printf(" download : %s\n", v)

			default:
				strConcat += "??? \n"

			}

		}
		results <- chambre

		print(line)

	}

	wg.Done()

}

// type model.Chambre struct {
// 	chambreNo int64
// 	date      time.Time
// 	upload    float64
// 	download  float64
// }
