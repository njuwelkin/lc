pi = 3.14
G = 6.67 * 10**(-11)
h = 6.62607015 * 10**(-34)
c = 3 * 10**8 # m/s

hc4 = 3 * 6.62607015 * 10**2 # hC^4

def evaporation_time(m): # (520πG^2 M^3)/(hC^4 )
	return (5120 * pi * m**3 * G**2) / hc4

def evaporation_t(m):
    return 8.41 * (m**3) * (10**(-17))

def evaporation_g(s):
    return (hc4 * s) / (520 * pi * G**2)

def e(m):
    return m * c**2

def power(m):
    return e(m) / evaporation_t(m) / 1000

def sch_r(m):
    return 2 * m * G / c**2
