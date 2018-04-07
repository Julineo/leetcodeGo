#def letterCombinations(self, digits):
digits = '123'
"""
:type digits: str
:rtype: List[str]
"""
if digits =='':
	print []
	raise SystemExit(0)
mapping = {'1':[''],
            '2':['a','b','c'],
            '3':['d','e','f'],
            '4':['g','h','i'],
            '5':['j','k','l'],
            '6':['m','n','o'],
            '7':['p','q','r','s'],
            '8':['t','u','v'],
            '9':['w','x','y','z'],
            '0':[' ']}
res =['']
    
for d in digits:
	print "d:", d
	temp = []
	print "res:", res
	for item in res:
		print "item:", item
		for x in mapping[d]:
			print "x:", x
			temp.append(item+x)
			print temp
	res = temp
print res
