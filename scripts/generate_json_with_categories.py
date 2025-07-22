import json
import base64
import os

def parse_blns(input_txt, output_json, add_base64=False):
    result = []
    current_cat = "Uncategorized"

    with open(input_txt, encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            if line.startswith("#"):
                current_cat = line.lstrip("#").strip()
            else:
                entry = {
                    "category": current_cat,
                    "string": line
                }
                if add_base64:
                    entry["base64"] = base64.b64encode(line.encode()).decode()
                result.append(entry)

    with open(output_json, "w", encoding="utf-8") as f:
        json.dump(result, f, ensure_ascii=False, indent=2)


if __name__ == "__main__":
    root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    blns_txt = os.path.join(root, "blns.txt")
    blns_json = os.path.join(root, "blns.json")
    blns_base64_json = os.path.join(root, "blns.base64.json")

    parse_blns(blns_txt, blns_json)
    parse_blns(blns_txt, blns_base64_json, add_base64=True)

    print("✅ JSON files generated with categories.")
