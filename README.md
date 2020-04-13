# Badger
Badger helps you find uncontrolled possible panics in your go functions.

It performs brute force executions on a given method using zero and random values. If a *panic* arises from the given method, an error will be returned with corresponding information.

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