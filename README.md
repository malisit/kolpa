# kolpa  [![Build Status](https://travis-ci.org/malisit/kolpa.svg?branch=master)](https://travis-ci.org/malisit/kolpa) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/malisit/kolpa) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/malisit/kolpa/master/LICENSE)
*kolpa* is a fake data generator written in and for Go.  

## usage
``` go
k := kolpa.C() // Initiate kolpa
k.Name() // Returns John Doe
k.FirstName() // Returns Jane
k.NameFemale() // Returns Jane Doe
```

You can set language setting when initiating `kolpa`.
``` go
k := kolpa.C("en_US")
```

Language can be setted afterwards as well.
``` go
k.SetLanguage("tr_TR")
```

