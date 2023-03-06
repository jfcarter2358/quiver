class StructureObject:
    def __init__(self, contents: str = ""):
        self.m_contents = contents
        self.m_children = []

    def __repr__(self):
        cleaned_contents = self.m_contents.replace("\"", "\\\"")
        
        children_contents = '['
        for child, i in enumerate(self.m_children):
            children_contents += str(child)
            if i < len(self.m_children):
                children_contents += ','
        children_contents += ']'
        
        return f'{{"contents": "{cleaned_contents}", "children": {children_contents}}}'

    def __str__(self):
        return self.__repr__()

    def parse_structure(self, lines: list, indent: int) -> list:
        output = []
        parse_buffer = []
        child_indent = -1

        for line in lines:
            if len(line.strip()) == 0:
                continue
            line_indent = len(line) - len(line.strip())
            if child_indent == -1:
                child_indent = line_indent
            if line_indent > indent:
                parse_buffer.append(line)
                continue
            if len(parse_buffer) > 0:
                output[len(output) - 1].m_children = self.parse_structure(parse_buffer, child_indent)
                parse_buffer = []
                child_indent = -1
            output.append(StructureObject(line.strip()))
        if len(parse_buffer) > 0:
            if len(output) == 0:
                output = self.parse_structure(parse_buffer, child_indent)
            else:
                output[len(output) - 1].m_children = self.parse_structure(parse_buffer, child_indent)
        return output

    def build_structure(self, path: str) -> list:
        lines = []
        with open(path, 'r', encoding='utf-8') as input_file:
            for line in input_file:
                lines.append(line)
        
        return parse_structure(lines, 0)
