ENDIAN = 'big'

UINT64_MAX = 2 ** 64

REGISTER_COUNT = 64

OP_CODES_DEF = {
    'ADD': 2,
    'SUBTRACT': 2,
    'MULTIPLY': 2,
    'DIVIDE': 2,
    'MODULO': 2,
    'POWER': 2,
    'BIN_ADD': 2,
    'BIN_SUBTRACT': 2,
    'GREATER': 2,
    'GREATER_EQUAL': 2,
    'EQUAL': 2,
    'LESS_EQUAL': 2,
    'LESS': 2,
    'INPUT': 1,
    'OUTPUT': 1,
    'BRANCH_POSITIVE': 2,
    'BRANCH_NOT_POSITIVE': 2,
    'BRANCH_ZERO': 2,
    'BRANCH_NOT_ZERO': 2,
    'BRANCH_NEGATIVE': 2,
    'BRANCH_NOT_NEGATIVE': 2,
    'GOTO': 1,
}

DOT_CODES_DEF = {
    '.STOP': 0,
    '.LITERAL': 2,
    '.RESERVE': 2,
    '.LABEL': 1,
}
