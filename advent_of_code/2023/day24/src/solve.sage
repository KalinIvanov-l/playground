var('x y z vx vy vz t1 t2 t3 ans')
eq1 = x + (vx * t1) == 320870677764563 + (-40 * t1)
eq2 = y + (vy * t1) == 335750934489987 + (-24 * t1)
eq3 = z + (vz * t1) == 282502845957937 + (10 * t1)
eq4 = x + (vx * t2) == 219235623600942 + (127 * t2)
eq5 = y + (vy * t2) == 408022798608755 + (-45 * t2)
eq6 = z + (vz * t2) == 245679379684914 + (66 * t2)
eq7 = x + (vx * t3) == 171834827764229 + (-122 * t3)
eq8 = y + (vy * t3) == 225154401936948 + (-521 * t3)
eq9 = z + (vz * t3) == 232302441670972 + (95 * t3)
eq10 = ans == x + y + z
print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9,eq10],x,y,z,vx,vy,vz,t1,t2,t3,ans))