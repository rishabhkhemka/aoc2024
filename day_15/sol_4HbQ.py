grid, moves = open('input_day_15.txt').read().split('\n\n')

G = grid.replace('#', '##').replace('.', '..').\
         replace('O', '[]').replace('@', '@.')

G = {i+j*1j: c for j,r in enumerate(G.split())
               for i,c in enumerate(r)}


def move(p, d, test=False):
    if not test and not move(p, d, True): return False

    ps = [p]
    if d.imag:
        if G[p] == '[': ps+= [(p+1)]
        if G[p] == ']': ps+= [(p-1)]

    for p in ps:
        if test:
            if G[p+d] in '[]' and not move(p+d, d, True)\
            or G[p+d] in '#': return False

        else:
            if G[p+d] in '[]': move(p+d, d)
            G[p+d], G[p] = G[p], G[p+d]

    return True


p = min({p for p,c in G.items() if c == '@'})

for m in moves.replace('\n', ''):
    d = {'<':-1, '>':+1, '^':-1j, 'v':+1j}[m]
    if move(p, d): p += d

print(sum(p.real+100*p.imag for p in G if G[p]=='['))