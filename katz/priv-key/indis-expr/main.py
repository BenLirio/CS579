class Set():
    def __init__(self, data):
        self.data = data
    def __mul__(self, other):
        A, B = self.data, other.data
        return Set(set([ (a, b) for a in A for b in B ]))
    def __or__(self, other):
        A, B = self.data, other.data
        return Set(A|B)
    def __and__(self, other):
        A, B = self.data, other.data
        return Set(A&B)
    def __str__(self):
        return str(self.data)
    def map(self, f):
        print(set(map(lambda x: 1, self.data)))
        return Set(set(map(f, self.data)))

def Enc(S):
    def f(xy):
        x,y = xy
        return (x + y)%2
    return S.map(f)

M = Set(set([0, 1]))
K = Set(set([0, 1]))
print(M*K)
C = Enc(M*K)
print(C)
