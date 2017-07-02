# kolpa  [![Build Status](https://travis-ci.org/malisit/kolpa.svg?branch=master)](https://travis-ci.org/malisit/kolpa) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/malisit/kolpa) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/malisit/kolpa/master/LICENSE)
*kolpa* is a fake data generator written in and for Go.  
It's capable of generating fake data for following instances for now,
- Name
- Address
- Color
- Datetime
- User Agent

## Installation
run ```go get github.com/malisit/kolpa``` on your command line.

## Usage
``` go
package main

import (
	"fmt"
	"github.com/malisit/kolpa"
	"time"
)

func main() {
	k := kolpa.C() // Initiate kolpa
	
	fmt.Println(k.FirstName()) // Prints John
	fmt.Println(k.Address()) // Prints 729 Richmond Springs Suite 949, Luisborough, VT 85700-5554
	fmt.Println(k.UserAgent()) // Prints Mozilla/5.0 (compatible; MSIE 5.0; Windows 98; Win 9x 4.90; Trident/4.0)
	fmt.Println(k.Color()) // Prints Lime #00FF00
	fmt.Println(k.DateTimeAfter(time.Date(2015, 1, 0, 0, 0, 0, 0, time.UTC))) // Prints 2015-09-08 15:34:29 +0300 EEST

}
```
List of all possible functions can be seen on godoc.

You can set language when initiating *kolpa*.
``` go
k := kolpa.C("tr_TR")
```

Language can be setted afterwards as well.
``` go
k.SetLanguage("tr_TR")
```

