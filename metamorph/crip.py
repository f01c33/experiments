import random
import zlib
import base64
import gc
import dis
import marshal

def t():
    x = 0
    print("Exploit 1=", x+1)
    # gerar outro
    global op
    op = f_enc(op)

    n_gerado = random.randint(len(op)//2,len(op))
    listid = sorted(random.sample(range(0, len(op)), n_gerado))

    out = [0 for i in range(len(op)+n_gerado)]
    i_code = 0
    for i in range(len(out)):
        if i in listid:
            out[i] = random.choice(op)
        else:
            out[i] = op[i_code]
            i_code += 1
    listid = f_enc(listid)
    out = ''.join(out)
    op = out+listid+'.'+str(len(listid)+3)

    novo = input("Novo arquivo:")
    #novo = "piroca.py"
    out =["""import random
import zlib
import base64
import marshal
pyt
op =""",'"',op,'"',
"""
f_enc = lambda f: str(base64.b64encode(zlib.compress(marshal.dumps(f),9)))[2:-1]
f_dec = lambda f: marshal.loads(zlib.decompress(base64.b64decode(f.encode())))

def dec(s):
    i = len(s)-1
    list_size = ""
    while s[i] != '.':
        list_size += s[i]
        i -= 1
    list_size = int(list_size[::-1])-3
    list = f_dec(s[i-list_size:i])
    return ''.join([s[i] for i in range(i-list_size) if i not in list])

op = f_dec(dec(op))
exec(op)
"""]
    out = ''.join(out)
    out = str(base64.b64encode(out.encode()))[2:-1]
    out = 'import base64\nexec(base64.b64decode('+'"'+out+'"'+'.encode()))'

    # eval(out)
    with open(novo, 'w') as file:
        file.write(out)


def enc(s):
    n_gerado = random.randint(len(s)//2,len(s))
    listid = sorted(random.sample(range(0, len(s)), n_gerado))

    out = [0 for i in range(len(s)+n_gerado)]
    i_code = 0
    for i in range(len(out)):
        if i in listid:
            out[i] = random.choice(s)
        else:
            out[i] = s[i_code]
            i_code += 1
    listid = f_enc(listid)
    out = ''.join(out)
    return out+listid+'.'+str(len(listid)+3)


f_enc = lambda f: str(base64.b64encode(zlib.compress(marshal.dumps(f),9)))[2:-1]
f_dec = lambda f: marshal.loads(zlib.decompress(base64.b64decode(f.encode())))

def make(code_name):

    #out = ['import zlib, marshal, base64\nexec(marshal.loads(zlib.decompress(base64.b64decode(','"',f_enc(t.__code__),'"','.encode()','))))']
    #out = compile(''.join(out), code_name, 'exec')
    with open(code_name, 'wb') as file:
        file.write(out)



def dec(s):
    #s = arg
    i = len(s)-1
    list_size = ""
    while s[i] != '.':
        list_size += s[i]
        i -= 1
    list_size = int(list_size[::-1])-3
    list = f_dec(s[i-list_size:i])
    return ''.join([s[i] for i in range(i-list_size) if i not in list])
    #out = ''.join([s[i] for i in range(i-list_size) if i not in list])

#make("novo.py")
print(enc(f_enc(t.__code__)))
