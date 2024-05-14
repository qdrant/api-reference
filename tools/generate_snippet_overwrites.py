
import argparse
from collections import defaultdict
import json
import glob
import yaml
import os


FILE_EXTENSION_BY_SDK = {
    'python': '.py',
    'javascript': '.js',
    'typescript': '.ts',
    'java': '.java',
    'csharp': '.cs',
    'go': '.go',
    'php': '.php',
    'ruby': '.rb',
    'swift': '.swift',
    'kotlin': '.kt',
    'rust': '.rs',
    'scala': '.scala',
    'dart': '.dart',
    'bash': '.sh',
    'txt': '.txt',
    'plaintext': '.txt',
    'curl': '.sh',
}


def parse_openapi_spec(path: str):
    """
    Load the OpenAPI specs and generate mapping of the endpoints to operation operationId
    """
    
    with open(path, 'r') as f:
        openapi_spec = json.load(f)
    
    paths = openapi_spec['paths']
    operation_mapping = {}
    for path, path_spec in paths.items():
        for method, method_spec in path_spec.items():
            operation_mapping[method_spec['operationId']] = [path, method]
    return operation_mapping


def create_empty_file(path: str):
    # make sure directory exists
    os.makedirs(os.path.dirname(path), exist_ok=True)

    with open(path, 'w') as f:
        f.write('')

def get_file_content_if_exists(path: str):
    try:
        with open(path, 'r') as f:
            return f.read()
    except FileNotFoundError:
        return ''
    

def convert_snippets_to_yaml_format(overwrites: dict) -> dict:
    """
    Convert to someting like this:
    paths:
        /collections/{collection_name}/shards:
            put:
            x-fern-examples:
                - code-samples:
                    - sdk: python
                    code: CODE
    """
    yaml_overwrites = {}
    for path, methods in overwrites.items():
        yaml_overwrites[path] = {}
        for method, sdks in methods.items():
            yaml_overwrites[path][method] = {
                'x-fern-examples': []
            }
            for sdk, code in sdks.items():
                yaml_overwrites[path][method]['x-fern-examples'].append({
                    'code-samples': [
                        {
                            'language': sdk,
                            'code': code
                        }
                    ]
                })
    return {
        "paths": yaml_overwrites
    }


def main():
    parser = argparse.ArgumentParser(description='Generate snippet overwrites for the OpenAPI spec')
    parser.add_argument('--openapi', type=str, help='Path to the OpenAPI spec')
    parser.add_argument('--output', type=str, help='Path to the snippet overwrite file')
    args = parser.parse_args()

    operation_mapping = parse_openapi_spec(args.openapi)

    # reverse_mapping = defaultdict(dict)

    # for operation_id, endpoint in operation_mapping.items():
    #     reverse_mapping[endpoint[0]][endpoint[1]] = operation_id


    # with open("fern/apis/master/openapi-overrides.yml", 'r') as f:
    #     existing_mappings = yaml.load(f, Loader=yaml.FullLoader)
    

    # for path, path_obj in existing_mappings['paths'].items():
    #     for method, method_obj in path_obj.items():
    #         code_samples = method_obj.get('x-fern-examples', [{}])[0].get('code-samples', [])

    #         for code_sample in code_samples:
    #             sdk = code_sample['language']
    #             code = code_sample['code']

    #             operation_id = reverse_mapping.get(path, {}).get(method)
    #             if not operation_id:
    #                 print(f"Operation not found for {path} {method}")
    #                 continue

    #             snippet_file = f"snippets/{sdk}/{operation_id}{FILE_EXTENSION_BY_SDK.get(sdk, '.txt')}"

    #             create_empty_file(snippet_file)

    #             with open(snippet_file, 'w') as f:
    #                 f.write(code)

    

    all_sdks = [sdk.split('/')[-1] for sdk in glob.glob("snippets/*")]

    overwrites = defaultdict(lambda: defaultdict(dict))

    for sdk in all_sdks:
        file_extension = FILE_EXTENSION_BY_SDK.get(sdk, '.txt')
        for operation_id, endpoint in operation_mapping.items():

            snippet_file = f"snippets/{sdk}/{operation_id}{file_extension}"

            file_content = get_file_content_if_exists(snippet_file)

            if not file_content or file_content.strip() == '':
                print(f"No snippet found for {operation_id} in {sdk}")
                continue

            api_path, method = endpoint

            overwrites[api_path][method][sdk] = file_content
    
    yaml_overwrites = convert_snippets_to_yaml_format(overwrites)

    with open(args.output, 'w') as f:
        yaml.dump(yaml_overwrites, f)
 

if __name__ == '__main__':
    main()
