# Badger
Badger helps you find uncontrolled possible panics in your go functions.

###How to use me: 

Consider this Sum function: 

```golang
func Sum(a,b int) int {
    return a + b
}
```

Once imported, we can look for panics with Badger like this: 

```golang
err := badger.Sniff(Sum)
```

This will return a descriptive error if Badger was able to cause an uncontrolled panic in the function. As simple as that...