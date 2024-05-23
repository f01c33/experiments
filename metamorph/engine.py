out = ""
def dexec(code):
    print(toStr(code))
    print("DEXEC")
    encoded = finalSplit(fromb64(code), sepFinal)
    k = encoded[-1]
    for i in range(len(encoded)-1):
        d,nk = dec(encoded[i],k)
        print(eval(d))
        k = fromb64(nk)

def main(inp):
    # print(inp)
    encoded = byteArrArr()
    k = key 
    # lines = finalSplit(toByteArr(lines,"\n"))
    lines = inp.split("\n")
    print(len(lines))
    for i in range(len(lines)):
        # print(lines[i])
        if lines[i] == "":
            continue
        nk = newkey()
        e = enc(toByteArr(lines[i]+" # "+tob64(nk)),k)
        print(toStr(e))
        encoded.append(e)
        g,nk1 = dec(encoded[-1],k)
        g2,nk2 = dec(e,k)
        print("E > ",toStr(g))
        if g != g2 or toStr(nk1) != toStr(tob64(nk)):
            # # print(">")
            # print("eval(g):",eval(g))
            print("g:\t",toStr(tob64(g)))
            print("g2:\t",toStr(tob64(g2)))
            print("nk1:\t",len(toStr(nk1)),toStr(nk1))
            print("nk:\t",len(toStr(tob64(nk))),toStr(tob64(nk)))
        k = nk
    final = tob64(finalJoin(encoded,key,sepFinal))
    out = toStr(final)
    # return toStr(final)
    return dexec(final)

main(inp)