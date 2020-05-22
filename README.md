### Compiler usage
```bash
go run ./cmd/graffle/main.go [PATH TO SRC FILE]
```

### Syntax
1. Comments
```python
    ``` One-line comment
    ` Multi-line
      comment `
```
1. Function declaration
```python
factorial(value) = answer, where
    if value equals to 0, then
        answer = 1.
    Else answer = value * factorial(value - 1)
    end.
end

f1(arg) = arg * z, where
    z = 9
end

f2(k) = a, where if k == false, then a = "NO". Else a = "YES". end. end

f3(i) = a, where if i is 1 then a = 1. is 2 then a = 10. default a = "RRRRARRRR!". end. end

procedure()
    <<< "I am the dragon!"
end
```
1. Conditional statements
```python
if a
    <<< "OK"
else
    <<< "Not OK"
end

if a
    is 1
        <<< "KO"
    is 2, then do
        <<< "LOL"
    default
        <<< "Well..."
end
```
1. Cycles
```python
for i = 0, i less than 10, i+=1
    print "Test for"
end.

until true do print "Well"

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
```python
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
```python
    ``` Print operator
    <<< "Why do we all must to wear those ridiculous ties?!"
    Print "Freeman, yo fool!"

    ``` Input operator
    >>> a
    input b
```

## Вариант задания
Реализация — `antlr4`. Целевой код — `go`(`gccgo`).

### Язык
Графовый язык.
Встроенные типы — `node`, `arc`, `graph`.
Переопределить операции — `+`, `-`, `*`, `\`, и т.д. для встроенных типов.

#### Дополнительные свойства языка

_(2)_ Неявное объявление переменных  
_(1)_ Явное преобразование `a=(int)b`  
_(1)_ Одноцелевое присваивание `a=b`  
_(1)_ Структуры, ограничивающие область видимости — подпрограммы  
~~_(2)_ Неявный маркер блока(как в python)~~  
_(1)_ Двух вариантный оператор условий `if-then-else`  
_(1)_ Перегрузка подпрограмм отсутствует  
_(1)_ Передача параметров в подпрограмму только по значению и возращаемому значению  
_(1)_ Объявление подпрограмм только в начале файла  
