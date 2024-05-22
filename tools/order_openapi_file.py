import argparse
from collections import OrderedDict
import yaml
import json

def load_yaml(path: str):
    with open(path, 'r') as f:
        return yaml.load(f, Loader=yaml.FullLoader)
    
def load_json(path: str):
    with open(path, 'r') as f:
        return json.load(f)
    
def save_json(data, path: str):
    with open(path, 'w') as f:
        json.dump(data, f, indent=2)


def order_openapi_file(openapi, ordering):
    paths = openapi['paths']

    orderred_paths = OrderedDict()

    for path in ordering:
        if path in paths:
            orderred_paths[path] = paths[path]
    
    for path in paths.keys():
        if path not in orderred_paths:
            orderred_paths[path] = paths[path]

    openapi['paths'] = orderred_paths


def main():
    parser = argparse.ArgumentParser(description='Generate snippet overwrites for the OpenAPI spec')
    parser.add_argument('--openapi', type=str, help='Path to the OpenAPI spec')
    parser.add_argument('--output', type=str, help='Path to the snippet overwrite file')
    args = parser.parse_args()

    ordering = load_yaml("fern/api-ordering.yml")

    ordering_list = [
        item['path']
        for item in ordering['order']
    ]

    openapi = load_json(args.openapi)

    order_openapi_file(openapi, ordering_list)

    save_json(openapi, args.output)


if __name__ == '__main__':
    main()
