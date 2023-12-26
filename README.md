# piece-table
A basic implementation of piece table in Go.

A piece table is a data structure generally used for representing text document while it is edited in a text editor. 
In this we maintain two buffers, one for storing the original text (immutable), and the second for storing text that we added (append-only).
We also maintain a table (list) of pieces where each piece represents a substring from buffers.
