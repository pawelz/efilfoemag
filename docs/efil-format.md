# Efil data format

Efil is a data format used to encode the state of the Game of Life used by
efilfoemag.

## Limitations

Width and height of the game must be divisible for 8.

The reason for that is strictly an implementation detail and I am sorry for
that. Efilfoemag makes extensive use of bitwise operations. It stores the
states of cells in individual bits. Having each row represented by a complete
number of bytes makes the code simpler.

## File format

The file consists of a list of lines. Lines are separated with the Unix newline
character. The last line must end ith the new line character.

### Header line

The first line of file consists of two integers separated by 'x' character.
Those integers represent the width and height of the Game respectively. Both
integers must be non-negative, divisible by 8, andcoded in base 10.

### Data lines

Each of the following lines encodes a single row of the game. Alive cell is
rendered as '#' character, dead cell is rendered as '+' character.

### Example

Here is a very simple example of a valid file:

```
8x8
++++++++
++++++++
++++#+++
+++#+#++
++++#+++
++++++++
++++++++
++++++++
```

See the `src/examples` directory for more examples.

## File names

By convention the efil files have suffix `.efil`.
