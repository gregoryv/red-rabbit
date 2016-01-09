[![Build Status](https://travis-ci.org/gregoryv/red-rabbit.svg)](https://travis-ci.org/gregoryv/red-rabbit)

![](redrabbit.png)

Sharing assorted go modules as I write my text editor.

cursor
------

Moving the index of a byte slice conceptually up and down 
between rows as you would when using your arrow keys in a
text file. The methods that you might miss in this module
are not missing, rather they are a part of the bytes module 
that comes with go.

The cursor methods are forgiving and handle errors in a manner 
you would expect from a text navigation point of view. Eg.
if you are on the last line and use the DOWN arrow key 
you will get the last index.
