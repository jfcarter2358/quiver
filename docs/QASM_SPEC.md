# Instructions

- `ADD <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `SUBTRACT <source> <source> <dest source>`
- `MULTIPLY <source> <source> <dest source>`
- `DIVIDE <source> <source> <dest source>`
- `MODULO <source> <source> <dest source>`
- `POWER <source> <source> <dest source>`
- `BIN_ADD <source> <source> <dest source>`
- `BIN_SUBTRACT <source> <source> <dest source>`
- `GREATER <source> <source> <dest source>`
- `GREATER_EQUAL <source> <source> <dest source>`
- `EQUAL <source> <source> <dest source>`
- `LESS_EQUAL <source> <source> <dest source>`
- `LESS <source> <source> <dest source>`
- `INPUT <source>`
- `OUTPUT <source>`
- `BRANCH_POSITIVE <source> <label>`
- `BRANCH_NOT_POSITIVE <source> <label>`
- `BRANCH_ZERO <source> <label>`
- `BRANCH_NOT_ZERO <source> <label>`
- `BRANCH_NEGATIVE <source> <label>`
- `BRANCH_NOT_NEGATIVE <source> <label>`
- `GOTO <label>`
- `STOP <code>`
    - ```
      [op code][return code]
      [ 1     ][ 1         ]
      ```


# Dot codes

- `.LITERAL <source> <data>`
    - ```
      [dot code][data type][source length][source][length][data...]
      [ 1      ][ 1       ][ 1           ][ ?    ][ 8    ][ ?     ]
      ```
- `.LABEL <source>`
    - ```
      [dot code][source length][source][program line]
      [ 1      ][ 1           ][ ?    ][ 8          ]
      ```
