import json

sample_data = {
    "example data..",
    "...."
}

num_objects = 10000
filename = "path/to/example_file.txt"

with open(filename, 'w') as file:
    for i in range(num_objects):
        modulefinder = sample_data.copy()
        modulefinder("fileds") += str(i)
        data_str = json.dumps(modulefinder)
        # file.write(data_str + "\n---\n")
