import assembler.bak.config as config

import sys


REGISTERS = {f'R{i}': i for i in range(0, config.REGISTER_COUNT)}
OP_CODES = {key: {'code': i, 'length': config.OP_CODES_DEF[key]} for i, key in enumerate(config.OP_CODES_DEF)}
DEF_CODES = {key: {'code': i, 'length': config.DOT_CODES_DEF[key]} for i, key in enumerate(config.DOT_CODES_DEF)}

class OP:
    def __init__(self, op):
        self.op = op
        self.arg1 = ''
        self.arg2 = ''
        self.arg2_type = ''

def first_pass(lines: list[str]) -> bytes:
    data_block = b''

    for line in lines:
        line = line.strip()
        if len(line) == 0:



def main():
    bytes = b''
    length = 0

    file_parts = sys.argv[1].split('.')
    file_name = '.'.join(file_parts[:-1])
    file_extension = file_parts[-1]

    with open(f'{file_name}.{file_extension}', 'r', encoding='utf-8') as input_file:
        lines = input_file.read().split('\n')

    for line in lines:
        line = line.strip()
        if len(line) == 0:
            continue

        if line.startswith('#'):
            continue

        parts = line.split(' ')
        op_code = parts[0]

        bytes += OP_CODES[op_code]['code'].to_bytes(1, config.ENDIAN)
        length += OP_CODES[op_code]['length'] + 1

    length_bytes = length.to_bytes(8, config.ENDIAN)

    bytes = length_bytes + bytes

    with open(f'{file_name}.qvc', 'wb') as output_file:
        output_file.write(bytes)
        
if __name__ == '__main__':
    main()
