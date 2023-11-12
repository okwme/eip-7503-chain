import random
import json
alen = 256
blen = 1 * 136 * 8

b = [random.choice([0, 1]) for _ in range(blen)]

subInput = b[100:356]
# subInput[240] = 1 - subInput[240]

mainInput = b
print(json.dumps({
    'subInput': subInput,
    'mainInput': mainInput,
    'numBlocks': 1
},indent=3))