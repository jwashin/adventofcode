# https://www.reddit.com/user/Marce_Villarino/
# https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/

datos = list()
moleculas = list()
steps = int()
with open("input.txt") as ficheiro:
    for liña in ficheiro:
        liña = liña.strip()
        if liña:
            key,_,value = liña.split()
            datos.append( (key,value) )
        else:
            break
    ### Yes, I had to take a look over there to guess this
    datos = sorted(datos, key =  lambda x: len(x[1]), reverse = True)
    ###
    moleculas.append(ficheiro.read().strip())

def backintime(molecula, transforms):
    out = list()
    for item in transforms:
        tail = molecula
        seed = ""
        while tail:
            head, sep, tail = tail.partition(item[1])
            if sep == item[1]:
                seed += head
                intermediate = seed + item[0] + tail
                seed += sep
                if "e" in intermediate and len(intermediate)>1: continue
                out.append(str(intermediate))
    return out

while "e" not in moleculas:
    steps += 1
    tmp = set()
    for item in moleculas:
        tmp.update( backintime(item, datos) )
    moleculas = sorted(list(tmp),key = len)[:5] #Consider only an arbitrary # of the shorter intermediate results
    if moleculas == []: break
print(steps)    
