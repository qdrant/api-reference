
import argparse
import json
import yaml

def main():
    parser = argparse.ArgumentParser(description='Remove api methods from the overwrite file if they are not present in the OpenAPI spec')
    parser.add_argument('--openapi', type=str, help='Path to the OpenAPI spec')
    parser.add_argument('--overrides', type=str, help='Path to the overwrite file')
    args = parser.parse_args()


    with open(args.overrides, 'r') as f:
        existing_mappings = yaml.load(f, Loader=yaml.FullLoader)

    with open(args.openapi, 'r') as f:
        openapi = json.load(f)

    openapi_paths = openapi['paths']
    existing_mappings_paths = existing_mappings['paths']

    to_remove = []

    for path, methods in existing_mappings_paths.items():
        for method in methods.keys():
            if path not in openapi_paths or method not in openapi_paths[path]:
                print(f"Removing {path} {method}")
                to_remove.append((path, method))

    for path, method in to_remove:
        del existing_mappings['paths'][path][method]

    with open(args.overrides, 'w') as f:
        yaml.dump(existing_mappings, f)
    
if __name__ == '__main__':
    main()
    
