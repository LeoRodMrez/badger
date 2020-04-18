# Badger
Badger helps you find uncontrolled panics in your go functions.

It analyzes the parameters of a given go function and performs executions using zero, limit-range and conflicting values. If a *panic* arises from the function, an error will be returned with the corresponding information.

The ultimate goal of Badger is to enhance defensive programming on your code.

### How to use me: 

Download me: 

```bash
go get github.com/LeoRodMrez/badger
```

Import me into your project:

```golang
import "github.com/LeoRodMrez/badger"
```

##### Example: 

Consider this go function: 

```golang
func getLetter(a string,b int) string {
    return string(a[b])
}
```

We can look for panics with Badger like this: 

```golang
err := badger.Sniff(getLetter)
```

This will return a descriptive error if Badger was able to cause an uncontrolled panic in the function. As simple as that...

So you could add a test to your project as follows:

```golang
func TestPanicsWithBadger(t *testing.T) {
    t.Run("BadgerTest"), func(t *testing.T) {
        hasPanicked := badger.Sniff(getLetter)
        require.Nil(t, hasPanicked, "Oops! Looks like this function is panicking...")
    })
}
```
Making sure your go functions stay away from uncontrolled panics, which can cause a lot of trouble.

You can also perform brute force executions on your go function with:

```golang
err := badger.BruteForceSniff(getLetter)
```

But keep in mind that brute force panic search is strongly based on random value generation so it may not be the best approach if you want a reproducible test environment.

### Currently supported parameter types:

```golang
int
bool
string
*int
*bool
*string
```

The maximum number of parameter supported is 5. Any violation on the number or the type of the parameters will also return an error.