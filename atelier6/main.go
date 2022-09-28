package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

}

func num1() {
	prime := make(chan int, 10)
	for len(prime) < 10 {
		go isPrime(genPremier(), prime)
	}
	var input string
	fmt.Scanln(&input)
}

func genPremier() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}
func isPrime(n int, c chan int) {
	if n < 2 {
		return
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	println(n)
	c <- n
}

func num2() {
	searchDir := "C:\\Users\\xboxm\\Documents\\AUT2022\\reseaux2\\image"
	fileList := make([]string, 0)

	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		filetoEdit, err := os.Open(file)

		if err != nil {
			log.Fatal(err)
		}
		// decode jpeg into image.Image
		if filepath.Ext(file) == ".jpg" || filepath.Ext(file) == ".jpeg" {
			img, err := jpeg.Decode(filetoEdit)

			if err != nil {
				log.Fatal(err)
			}
			filetoEdit.Close()

			m := resize.Resize(80, 0, img, resize.Lanczos3)
			filename := strings.ReplaceAll(file, "C:\\Users\\xboxm\\Documents\\AUT2022\\reseaux2\\image\\", "")

			out, err := os.Create("C:\\Users\\xboxm\\Documents\\AUT2022\\reseaux2\\image\\imageResize_" + filename)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			jpeg.Encode(out, m, nil)
		}
		if filepath.Ext(file) == ".png" {
			img, err := png.Decode(filetoEdit)

			if err != nil {
				log.Fatal(err)
			}
			filetoEdit.Close()

			m := resize.Resize(80, 0, img, resize.Lanczos3)
			filename := strings.ReplaceAll(file, "C:\\Users\\xboxm\\Documents\\AUT2022\\reseaux2\\image\\", "")

			out, err := os.Create("C:\\Users\\xboxm\\Documents\\AUT2022\\reseaux2\\image\\imageResize_" + filename)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to fil
			png.Encode(out, m)
		}
	}
}