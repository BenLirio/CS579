# Just track to coefficients
def f(a, b):
    if b < a:
        c, d = f(b, a)
        return d, c
    if b%a == 1:
        return -(b//a), 1
    c, d = f(b%a, a)
    return d-c*(b//a), c

p = 7
q = 3
e1 = q*f(p,q)[1]
e2 = p*f(p,q)[0]

def gcd(a, b):
    if a > b:
        return gcd(b, a)
    if a == 0:
        return b
    return gcd(b%a, a)

def to_base_p(x):
    return (x%p, x%q)

def from_base_p(x):
    a, b = x
    return (a*e1 + b*e2) % (p*q)

for i in range(0, p*q):
    print(from_base_p(to_base_p(i)))

