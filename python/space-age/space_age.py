

class SpaceAge(object):
    onEarth = 31557600
    onMercury = 0.2408467*onEarth
    onVenus = 0.61519726*onEarth
    onMars = 1.8808158*onEarth
    onJupiter = 11.862615*onEarth
    onSaturn = 29.447498*onEarth
    onUranus = 84.016846*onEarth
    onNeptune = 164.79132*onEarth


    def __init__(self, seconds):
        self.space_age = seconds
    def on_earth(self):
        return round(self.space_age/self.onEarth, 2)

    def on_mercury(self):
        return round(self.space_age/self.onMercury, 2)

    def on_venus(self):
        return round(self.space_age/self.onVenus, 2)

    def on_mars(self):
        return round(self.space_age/self.onMars, 2)

    def on_jupiter(self):
        return round(self.space_age/self.onJupiter, 2)

    def on_saturn(self):
        return round(self.space_age/self.onSaturn, 2)

    def on_uranus(self):
        return round(self.space_age/self.onUranus, 2)

    def on_neptune(self):
        return round(self.space_age/self.onNeptune, 2)

    


