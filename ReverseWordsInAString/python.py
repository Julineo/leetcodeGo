//works firs or second version depending on python version

s = "   a   b"
if s == " ":
    print s.strip()
rr = s[::-1]
rr = rr.split(" ")
res = ""
for r in rr:
    r = r[::-1]
    res = res + " " + r
print res.strip()



if s == " ":
    print s.strip()
        
data = s.split(" ")
new = []
for i in data:
    if i == " " or i == "":
        continue
    new.append(i)

print " ".join(new[::-1])
