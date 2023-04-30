`range` is a keyword in Go used to iterate over data structures such as arrays, slices, maps, and strings. It provides a convenient way to access each element of a collection, one at a time, in a loop.

The syntax of `range` varies slightly depending on the type of data structure being iterated over. For example, when used with an array or a slice, `range` returns the index and the value of each element in turn:

``` go
nums := []int{1, 2, 3, 4, 5}
for index, value := range nums {
    fmt.Println(index, value)
}
```

Output:
```
0 1
1 2
2 3
3 4
4 5
```

When used with a map, `range` returns the key and the value of each entry in turn:

``` go
m := map[string]int{"a": 1, "b": 2, "c": 3}
for key, value := range m {
    fmt.Println(key, value)
}
```

Output:
```
a 1
b 2
c 3
```

`range` can also be used with strings, in which case it returns the index and the Unicode code point of each character in turn:

``` go
s := "hello"
for index, char := range s {
    fmt.Println(index, char)
}
```

Output:
```
0 104
1 101
2 108
3 108
4 111
```

Note that the second variable returned by `range` is the value for arrays, slices and maps, but it's the Unicode code point for strings.