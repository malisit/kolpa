# kolpa  
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

