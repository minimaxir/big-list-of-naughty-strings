import json
from scripts import parse_blns

def test_categories_and_strings(tmp_path):
    txt = tmp_path / "blns.txt"
    txt.write_text("# Cat1\nfoo\n# Cat2\nbar\n")
    out = tmp_path / "blns.cat.json"
    parse_blns(str(txt), str(out))
    data = json.loads(out.read_text())
    assert data == [
        {"category": "Cat1", "string": "foo"},
        {"category": "Cat2", "string": "bar"},
    ]
