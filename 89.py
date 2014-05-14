f = open('roman.txt')

s = {
    'IV': 4,
    'IX': 9,
    'XL': 40,
    'XC': 90,
    'CD': 400,
    'CM': 900,
}

c = {'I' : 1,
    'V' : 5,
    'X' : 10,
    'L' : 50,
    'C' : 100,
    'D' : 500,
    'M' : 1000,
    }

a =  dict(**s)
a.update(c)

rev_a = {v:k for k,v in a.items()}

savings = 0

for line in f:
    i = line.strip()
    for x in s.keys() + c.keys():
        if x in i:
            i = i.replace(x, ' %s ' % str(a[x]))
    i = i.split()
    num = sum(int(x) for x in i)
    n = str(num)
    l = len(n)
    n = [(int(x) * pow(10, l-i), int(x)) for i, x in enumerate(n, 1)]
    res = []
    for i, (x, y) in enumerate(n):
        if x == 0:
            continue
        key = x / y
        if x in rev_a:
            res.append(rev_a[x])
        else:
            p1, p2 = y / 5, y % 5
            if p1:
                res.append(rev_a[5 * key])
            res.append(p2 * rev_a[key])

    res = ''.join(res)    
    savings += len(line.strip()) - len(res)

print savings
