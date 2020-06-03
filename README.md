# Graffle

### Compiler usage
```bash
make antlr4gen

go run ./cmd/graffle/main.go [PATH TO SRC FILE]
```

### Syntax
Graffle aims to be convenient for those, who knows other languages,
like python, lua, prolog etc, and to look like regular mathematical text.
This means, that graffle users have a lot of ways to write code,
so it can look more like natural text(with a lot of words,
capitalisation and punktuation), or like shorthand friendly code.

1. Hello world
    ```python
    print "Hello, World!"
    ``` This is equivalent to:
    <<< "Hello, World!"
    ```
   You can use both `'` and `"` characters to write strings.
1. Comments
    ```c++
    ``` One-line comment
    ` Multi-line
      comment `
    ```
1. Function declaration
    ```lisp
    ``` Shorthand examples:
    factorial(value) = answer
        if value == 0
            answer = 1
        else
            answer = value * factorial(value - 1)
        end
    end

    f1(arg) = arg * z
        z = 9
    end
   
    procedure()
           print "Procedure called"
       end

    ``` Natural text-like examples:
    f2(k) = a, where if k is false, then a = "NO". Else a = "YES". end. end.

    f3(i) = a, where
    If i is 1, then a = 1. Is 2, then do print "2". Default a = "RRRRARRRR!". end.
    end.
    ```
1. Conditional statements
    ```lisp
    if a
        <<< "a"
    elif b
        <<< "b"
    else
        <<< "c"
    end

    ``` Alternative to switch operator(like in C, Go etc):
    if a
        is 1
            <<< "KO"
        is 2
            <<< "LOL"
        default
            <<< "Well..."
    end
    ```
1. Cycles
    ```lisp
    for i = 0, i < 10, i += 1
        print "Test for"
    end

    until true do print "This text will be not printed"

    for i in range from 0 to 10 print "Test for in range"

    i = 0
    while i < 10
        <<< 'Test while'
    end

    Until i not less than 10 do
        Print 'Test until'
    end

    from 0 to 10 print 'Test from to'
    ```
1. Built-in types
    ```lisp
    v1 = 12 @ int vertice
    v2 = true @ boolean s here
    v3 = "Why do we all must to wear those ridiculous ties?!" @[ This was
    a question]

    e1 = v1 -> v2 @ not weighted oriented arc
    e2 = v2 -- v3 @ not weighted not oriented arc
    e3 = v1 -[1]-> v2 @ weighted oriented arc

    g1 = G(e1) @ graph converted from arc
    g2 = G(g1 + v3) @ graph converted from summ of graph and vertice

    g = (g1 + g2) @ summ of graphs
    g = g1 * g2 @ let's change the label
    ```
1. Standard functions
    ```lisp
    ``` Print operator
    <<< "Why do we all must to wear those ridiculous ties?!"
    Print "Freeman, yo fool!".

    ``` Input operator
    >>> a
    input b
    ```
