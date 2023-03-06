# Instructions

- `ADD <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `SUBTRACT <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `MULTIPLY <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `DIVIDE <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `MODULO <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `POWER <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `BINARY_ADD <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `BINARY_SUBTRACT <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `GREATER <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `GREATER_EQUAL <source> <source> <dest>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `EQUAL <source> <source> <dest>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `LESS_EQUAL <source> <source> <dest>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `LESS <source> <source> <dest source>`
    - ```
      [op code][source length 1][source 1][source length 2][source 2][dest length][dest]
      [ 1     ][ 1             ][ ?      ][ 1             ][ ?      ][ 1         ][ ?  ]
      ```
- `INPUT_BLOCK <dest>`
    - ```
      [op code][dest length][dest]
      [ 1     ][ 1         ][ ?  ]
      ```
- `INPUT_NON_BLOCK <dest_ascii> <dest_keycode>`
    - ```
      [op code][dest_ascii length][dest_ascii][dest_keycode length][dest_keycode]
      [ 1     ][ 1               ][ ?        ][ 1                 ][ ?          ]
- `OUTPUT <source>`
    - ```
      [op code][source length][source]
      [ 1     ][ 1           ][ ?    ]
      ```
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
