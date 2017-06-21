# kolpa  [![Build Status](https://travis-ci.org/malisit/kolpa.svg?branch=master)](https://travis-ci.org/malisit/kolpa) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/malisit/kolpa) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/malisit/kolpa/master/LICENSE)
*kolpa* is a fake data generator written in and for Go.  

## usage
``` go
package main

import (
	"fmt"
	"github.com/malisit/kolpa"
)

func main() {
	k := kolpa.C() // Initiate kolpa
	
	fmt.Println(k.FirstName()) // Prints John
	fmt.Println(k.Address()) // Prints 729 Richmond Springs Suite 949, Luisborough, VT 85700-5554

}
```

You can set language when initiating *kolpa*.
``` go
k := kolpa.C("tr_TR")
```

Language can be setted afterwards as well.
``` go
k.SetLanguage("tr_TR")
```

