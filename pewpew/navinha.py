import pygame, sys
from random import randint
from pygame.locals import *
from pygame.math import *
from math import sin,cos,pi

pygame.init()

from score import *
#set up icon
icon=pygame.Surface((32,32))
icon.set_colorkey((0,0,0))
rawicon=pygame.image.load('icon.png')
pygame.display.set_icon(icon)

#set up the display
FPS = 60
fpsClock = pygame.time.Clock()
graphwidth = 640
graphheight = 480
DISPLAYSURF = pygame.display.set_mode((graphwidth, graphheight))
DISPLAYSURF2 = pygame.display.set_mode((graphwidth, graphheight))
pygame.display.set_caption("PEW PEW!")

#set up sound
pygame.mixer.pre_init(44100,-16,2, 4096)
pygame.mixer.init()
main_menu_snd = pygame.mixer.Sound('farewell.ogg') 
shot_snd = pygame.mixer.Sound('shot.ogg')
take_dmg_snd = pygame.mixer.Sound('snare.ogg')
boss_is_dead_snd = pygame.mixer.Sound('boss_dies.ogg')
ost_snd = pygame.mixer.Sound('timbre_-6db.ogg')
exploding_snd = pygame.mixer.Sound('explosion.ogg')

#set up the colors
BLACK = ( 0, 0, 0)
WHITE = (255, 255, 255)
RED = (255, 0, 0)
GREEN = ( 0, 255, 0)
BLUE = ( 0, 0, 255)

#set up sprites
bg = pygame.image.load('BG.jpg')
playerImg = pygame.image.load('player_ship.png')
playershotImg = pygame.image.load('player_shot.png')
bossImg = pygame.image.load('bigboss.png')
boss_shotImg = pygame.image.load('shot_boss_missile.png')
radial_shotImg = pygame.image.load('shot_boss_radial.png')
straight_shotImg = pygame.image.load('shot_straight.png')
whip_shotImg = pygame.image.load('shot_whip.png')
wave_shotImg = pygame.image.load('shot_wave.png')
aimed_shotImg = pygame.image.load('shot_aimed.png')
mob1Img = pygame.image.load('mob1.png')
mob2Img = pygame.image.load('mob2.png')
mob3Img = pygame.image.load('mob3.png')
mob4Img = pygame.image.load('mob4.png')

#explosions
ex1 = pygame.image.load('explosion1.png')
ex2 = pygame.image.load('explosion2.png')
ex3 = pygame.image.load('explosion3.png')
ex4 = pygame.image.load('explosion4.png')#All explosion frames
ex5 = pygame.image.load('explosion5.png')
ex6 = pygame.image.load('explosion6.png')
ex7 = pygame.image.load('explosion7.png')
ex8 = pygame.image.load('explosion8.png')
ex9 = pygame.image.load('explosion9.png')
ex10 = pygame.image.load('explosion10.png')
ex11 = pygame.image.load('explosion11.png')
ex12 = pygame.image.load('explosion12.png')
ex13 = pygame.image.load('explosion13.png')
ex14 = pygame.image.load('explosion14.png')


##set up the variables
#mouse variables
mousex = 0
mousey = 0
mouseClicked = False
#variables used to store time and game-state
game_finish = 0
time_game_begun = 0
time_game_finished = 0
main_menu_music = False

#groups that will contain every object in the game
enemy_list = pygame.sprite.Group()
player_list = pygame.sprite.Group()
enemy_shots = pygame.sprite.Group()
player_shots = pygame.sprite.Group()
explosions = pygame.sprite.Group()

everything = [enemy_list,player_shots,enemy_shots,explosions]

gamemode = ''
dif = 1 #Difficulty, hard by standard.
kills = 0 #kills, used for the purpose of scoring
score = Score() #Score object, explained in detail at score.py
font = pygame.font.SysFont('impact', 50, False, False) #prepares a font according to pygame syntax

boss_spawned = False

###########################=========############################
######################## Classes Setup #########################
###########################=========############################
class shot(pygame.sprite.Sprite): #creates shot as an object, which ships will create with their shoot() method
    def __init__(self,x,y,spd,dmg,type1,sprite,team):
        pygame.sprite.Sprite.__init__(self)#Required for sprite objects as per pygame syntax
        self.image = sprite
        self.rect = self.image.get_rect()
        self.rect.x = x
        self.rect.y = y
        self.spd = spd
        self.dmg = dmg
        self.type = type1
        
        self.team = team

        self.cont = 0 #Cont is used by shots that have constantly variating sine/cosine
        if self.type == 'whip':
            self.direction_toggle = randint(0,1)#Controls the direction of spinning on whip shots
            
        if self.type == 'aimshot':#Defines a vector between the player and the position of the shot the moment it is created
            self.vector = Vector2(player.rect.centerx-self.rect.centerx,player.rect.centery-self.rect.centery)
            self.angle = self.vector.as_polar()[1]#Then gets the angle of this vector by converting to polar coordinates
            self.radians = self.angle * pi/180#And finally converts it to radians so it may be used with the cos and sin functions

            self.x = x
            self.y = y
        if self.team == 'red':
            enemy_shots.add(self)
        else:
            player_shots.add(self)
        shot_snd.play()
    def move(self):
        """Method used to define how each type of shot will move, each ship has a different type"""
        if self.type == 'simple':
            self.rect.y += self.spd*fpsmod*(dt/dtmod)#(dt/dtmod) is the solution to low fps, making the game run based on time
                                              # It is further explained in the play() loop
        elif self.type == 'wave':#Sine-like movement
            self.rect.y += self.spd*fpsmod*(dt/dtmod)
            self.rect.x += (self.spd*2*sin(self.cont) - self.spd*2*cos(self.cont))*fpsmod*(dt/dtmod)
            self.cont += 0.2*(dt/dtmod)*fpsmod*(dt/dtmod)
            
        elif self.type == 'whip':#Circular movement
            self.rect.y += (self.spd*sin(self.cont) + self.spd*cos(self.cont) + 1.5)*fpsmod*(dt/dtmod)
            self.rect.x += (self.spd*sin(self.cont) - self.spd*cos(self.cont))*fpsmod*(dt/dtmod)
            if self.direction_toggle == 1:#Controls whether the spin of the projectile is clockwise or anti-clockwise
                self.cont += 0.1*fpsmod*(dt/dtmod)
            else:
                self.cont -= 0.1*fpsmod*(dt/dtmod)
                
        elif self.type == 'aimshot':
            self.x += self.spd*cos(self.radians)*fpsmod*(dt/dtmod)
            self.rect.x = self.x
            self.y += self.spd*sin(self.radians)*fpsmod*(dt/dtmod)
            self.rect.y = self.y
            
        elif self.type == 'explosive':#Explosive sort of shot
            self.rect.y += self.spd*(dt/dtmod)*fpsmod
            if self.rect.y > 320-self.expl_y: #When it has travelled enough, explodes into 8 more different shots
                cont = 0
                for i in range(1,9):
                    spd = (self.spd * cos(cont))*(dt/dtmod)#Each shot speed is controlled by a factor of sine and cosine
                    spdy = (self.spd * sin(cont))*(dt/dtmod)#so that they move diagonally with the correct speed, as to mimic a circle.
                    instance = shot(self.rect.x,self.rect.y,spd,6/dif,'radial',radial_shotImg,'red')
                    instance.spdy = spdy
                    cont += pi/4 #Variation in degrees between each shot
                self.kill()#.kill() is a method that comes with the inheritance from pygame.sprite.Sprite, it completely deletes the object
        elif self.type == 'radial':#Shot created by the explosive shot
            self.rect.y += self.spdy*fpsmod#Moves based on the variation of cosine and sine
            self.rect.x += self.spd*fpsmod
    
class ship(pygame.sprite.Sprite): #sets up the ship class, which is the main class that represents all space-ships in the game.
    state = 'hunting'
    def __init__(self,x,y,hp,spd,shotspd,shotdmg,shottype,delay,sprite,spriteshot,team):
        pygame.sprite.Sprite.__init__(self)
        self.image = sprite #Stores the image from the sprite
        self.rect = self.image.get_rect() #Accomodates a rectangle of the same size as the sprite, will be used for collision and boundary checks
        self.rect.x = x #Sprite's position
        self.rect.y = y
        self.hp = hp #Hull points, the amount of damage it can sustain before exploding
        #variables that control the movement
        self.spd = 0 #variable that controls momentum / current speed
        self.maxspd = spd #max speed
        self.accel = spd/5
        #variables that store the shot information
        self.shotspd = shotspd
        self.shotdmg = shotdmg
        self.shottype = shottype
        self.delay = delay
        self.spriteshot = spriteshot
        self.maxdelay = delay
        #team is used to discern whether the object is an enemy or not
        self.team = team

        if self.team == 'red':
            enemy_list.add(self)
        else:
            player_list.add(self)
        #Positions for the fastguy class
        #Defines one random point on each third of the screen
        self.positions = {1:randint(self.rect.width+10,graphwidth/3),2:randint(graphwidth/3,(2*graphwidth)/3),3:randint((2*graphwidth)/3,graphwidth-self.rect.width - 10)}

    def shoot(self):
        """Creates an shot object based on the ship's current position and which one is it"""
        if self.team == 'red':#Enemy shots are created below their sprite, since they move downwards
            instance = shot(self.rect.midbottom[0],self.rect.midbottom[1],self.shotspd,self.shotdmg,self.shottype,self.spriteshot,self.team)
        else:#Player shots are created above the player sprite.
            instance = shot(self.rect.midtop[0]-6,self.rect.midtop[1]-8,self.shotspd,self.shotdmg,self.shottype,self.spriteshot,self.team)

    def takedamage(self,shotdmg):
        """Deals damage to the objects that is calling this method, whereas shotdmg is the damage dealt"""
        self.hp -= shotdmg
        if self.team == 'red':#Only changes the state if its an enemy
            self.state = 'fleeing'
        else:
            take_dmg_snd.play()#Only plays the sound if the player takes damage
        
    def aimov(self):
        """Uses states as the mean of choosing how the AI ships will move"""
        if self.state == "hunting": #follows the player
            if player.rect.centerx > self.rect.centerx+self.maxspd+self.rect.width/4 \
               and self.spd <= self.maxspd:
                self.spd += self.accel*fpsmod
            elif player.rect.centerx < self.rect.centerx-(self.maxspd+self.rect.width/4) \
               and self.spd > -self.maxspd:
                self.spd -= self.accel*fpsmod
                
        if self.state == "fleeing": #runs from the player
            if self.rect.centerx not in range(0,graphwidth): #when ship reaches either of the screen edges, go back to hunting
                self.state = "hunting"
                return
            if player.rect.centerx > self.rect.centerx \
               and self.rect.centerx in range(0,graphwidth)\
               and self.spd > -self.maxspd:#if the player is to the right of the ship, move towards the left
                self.spd -= self.accel*fpsmod
            elif player.rect.centerx < self.rect.centerx \
               and self.rect.centerx in range(0,graphwidth)\
               and self.spd < self.maxspd:#if its left, move right
                self.spd += self.accel*fpsmod
        self.rect.x += self.spd*(dt/dtmod)
            
    def aishoot(self):
        """The method which is used for controlling the AI's shot pattern"""
        self.delay -= dt/dtmod
        if self.delay <= 0:
                self.shoot()
                self.delay = self.maxdelay

    def ai(self):
        """The AI of the object """
        self.aishoot()
        self.aimov()
        
    def explode(self):
        """Method called when the ship dies, which plays the sound and creates an object explosion, where the ship was"""
        exploding_snd.play()
        self.exploding = True
        instance = explos(self.rect.x-self.rect.width,self.rect.y-self.rect.height)
        self.kill()
        return

class boss(ship): #boss has some different patterns, so we created a new object that inherits all the basics from the 'ship' object
    def explode(self): #calls 3 explosions, instead of one
        global boss_spawned
        boss_is_dead_snd.play()
        instance = explos(self.rect.x,self.rect.y)
        instance = explos(self.rect.x+self.rect.width/2,self.rect.y)
        instance = explos(self.rect.x+self.rect.width,self.rect.y)
        self.kill()
    def aimov(self): #changes the movement for the boss, so that it can spawn in the middle
        if self.rect.centery < self.rect.height:#Also never flees. The boss doesn't mind taking some shots
            self.rect.centery += self.accel*(dt/dtmod)
        if player.rect.centerx > self.rect.centerx+self.maxspd+self.rect.width/4 and self.spd <= self.maxspd/2:
            self.spd += self.accel*fpsmod
            self.rect.x += self.spd/4
        elif player.rect.centerx < self.rect.centerx-(self.maxspd+self.rect.width/4) and self.spd > -self.maxspd/2:
            self.spd -= self.accel*fpsmod
            self.rect.x += self.spd/4
        self.rect.x += self.spd*(dt/dtmod)
    def shoot(self): #works in a way so that the hp affects how often the boss attacks.
        if self.hp <= 300:
            instance = shot(self.rect.x+(0.8*self.rect.width/2),self.rect.y+(self.rect.height),self.shotspd,self.shotdmg,self.shottype,self.spriteshot,self.team)
            instance.expl_y = randint(0,80)
        if self.hp <= 220:
            self.maxdelay = 50
        if self.hp <= 180:
            self.maxdelay = 40
        if self.hp <= 120:
            self.maxdelay = 30

class fastguy(ship): #A quick enemy who bounces repeatedly from one point to the other and then stops to fire at the player.
    state = 'moving'
    shotcount = 10 #The amount of shots the ship will fire until it gets back to moving.
    nextpos = randint(1,3)
    def takedamage(self,shotdmg): #Overrides previous takedamage() method as to avoid the 'fleeing' state
        self.hp -= shotdmg

    def moveleft(self):
        if self.rect.centerx > self.positions[self.nextpos]+70 and self.spd > -self.maxspd:#Accelerates towards the next point
            self.spd -= self.accel*(dt/dtmod)
        if self.rect.centerx <= self.positions[self.nextpos]+70 and self.spd < 0:#Deaccelerates to stop near the next point
            self.spd += self.accel*(dt/dtmod)
            if self.spd > 0:#Countermeasure against low-fps, as to avoid spd being left at any value other than 0
                self.spd = 0
        if self.spd == 0:
            self.prevpos = self.nextpos #Saves the current pos
            self.nextpos = randint(1,3) #Redefines the next random pos
            while self.nextpos == self.prevpos: #If it's the same as the current pos, retries until it's different.
                self.nextpos = randint(1,3)
            self.state = 'shooting'#When it has fully stopped, switch to shooting.
        self.rect.x += self.spd
        self.delay -= dt/dtmod#delay is always decremented from so the next time the ship stops, it can always start firing right away

    def moveright(self):#Same thing as above but assumes the next position is to the right
        if self.rect.centerx < self.positions[self.nextpos]-70 and self.spd < self.maxspd:
            self.spd += self.accel*(dt/dtmod)
        if self.rect.centerx >= self.positions[self.nextpos]-70 and self.spd > 0:
            self.spd -= self.accel*(dt/dtmod)
            if self.spd < 0:
                self.spd = 0
        if self.spd == 0:
            self.prevpos = self.nextpos
            self.nextpos = randint(1,3)
            while self.nextpos == self.prevpos:
                self.nextpos = randint(1,3)
            self.state = 'shooting'#When it has bounced enough times, switches to shooting
        self.rect.x += self.spd
        self.delay -= dt/dtmod


    def aishoot(self):
        self.delay -= dt/dtmod
        if self.delay <= 0 and self.shotcount != 0:#Will shoot until shotcount is 0
                self.shoot()
                self.delay = self.maxdelay
                self.shotcount -= 1#Decrements from shotcount for every shot fired
        if self.shotcount == 0:
            self.state = 'moving'#When shot count is finally zero, will get back to moving
            self.shotcount = 10#Prepares shotcount for the next cycle

    def ai(self):
        if self.state == 'moving':
            if self.rect.centerx >= self.positions[self.nextpos]:     #A simple finite state machine responsible for the AI
                self.moveleft()
            else:
                self.moveright()
        elif self.state == 'shooting':
            self.aishoot()
            
class explos(pygame.sprite.Sprite): #The class which one calls when its hp is below 0, creating an explosion where it was
    explode_frame = 0
    explodeimg = [ex2,ex3,ex4,ex5,ex6,ex7,ex8,ex9,ex10,ex11,ex12,ex13,ex14]#List containing all sprites for the explosion
    def __init__(self,x,y,sprite=ex1):#Always sets the first sprite of the explosion automatically
        pygame.sprite.Sprite.__init__(self)
        self.image = sprite
        self.rect = self.image.get_rect()
        self.rect.x = x
        self.rect.y = y
        explosions.add(self)
    def cycle(self):#Everytime this method is called, it changes the current sprite to the next until its at the last one
        if self.explode_frame >= len(self.explodeimg):
            self.kill()#Then it deletes the whole object.
            return
        self.image = self.explodeimg[self.explode_frame]
        self.explode_frame += 1 * (dt/dtmod)
        self.explode_frame = int(self.explode_frame)
        


############################========############################
######################## Game mechanics ########################
###########################=========############################
def draw():
    """Calls the surface drawing functions, and updates the screen"""
    global levelwait
    DISPLAYSURF.blit(bg,(0,0))
    #DISPLAYSURF2.blit(font.render("FPS " + str(fpsClock.get_fps()), True, WHITE),(font.size("FPS" + str(fpsClock.get_fps()))[0]/8, font.size("FPS" + str(fpsClock.get_fps()))[1]))
    if levelwait > 0:
        levelwait -= dt
        DISPLAYSURF2.blit(font.render("Level " + str(level), True, WHITE),(font.size("Level" + str(level))[0]/8, graphheight - font.size("Level" + str(level))[1]))
    drawships(enemy_list)
    drawships(player_list)
    drawshots(enemy_shots,player_list)
    drawshots(player_shots,enemy_list)#Player shots are drawn after the enemy's so that if they overlap, the player's is the one visible
    drawexplosions()
    drawhp()
    pygame.display.update()
    
def drawhp():
    """Draws the hp bar for the boss and the player"""
    if boss_spawned == True \
        and (bigboss.rect.x >= -bigboss.rect.width/2 \
        and bigboss.rect.x <= graphwidth-bigboss.rect.width/2):
        pygame.draw.rect(DISPLAYSURF, RED,(0,0,((graphwidth*bigboss.hp)/300),5))#boss                                      
    pygame.draw.rect(DISPLAYSURF, GREEN, (0,graphheight-5,(graphwidth*player.hp)/100,graphheight)) #player

def drawships(shipt):
    """Draws all the different ships in a given list"""
    for i in shipt:
        DISPLAYSURF.blit(i.image, (i.rect.x, i.rect.y))#DISPLAYSURF.blit draws the given image on the given coordinates

def drawshots(shots_group,opposing_group):
    """Draws all the shots and checks collision"""
    for shot in shots_group:
        oldy = shot.rect.centery
        oldx = shot.rect.centerx#Stores the shot position before moving
        shot.move()
        newy = shot.rect.centery#Stores the shot position after moving
        newx = shot.rect.centerx

        if oldy - newy <=0:#Accomodates the code in case the variation is reversed
            yvariation = range(oldy,newy+1)
        else:
            yvariation = range(newy,oldy+1)#Will create a list with the trajectory of the shot for both axis
        if oldx - newx <= 0:
            xvariation = range(oldx,newx+1)#This means, all positions the shot has technically gone through.
        else:
            xvariation = range(newx,oldx+1)
        
        DISPLAYSURF.blit(shot.image, (shot.rect.x, shot.rect.y))
        if shot.rect.centery not in range(0,graphheight)\
           or shot.rect.centerx not in range(0,graphwidth): #If the shot has left the display, deletes it.
            shot.kill()
            continue
        
        for ship in opposing_group:
            x_is_true = False #Booleans used to check whether the shot has collided on the given axis
            y_is_true = False
            for number in xvariation:#Iterates through all values of the variation
                if number in range(ship.rect.left,ship.rect.right):#and checks if any of them has gone inside the collision box for the ship
                    x_is_true = True
                    break
            for number in yvariation:
                if number in range(ship.rect.top,ship.rect.bottom):#Does this for both axis
                    y_is_true = True
                    break
            if y_is_true and x_is_true: #If there has been collision on both axis, then the shot has collided.
                ship.takedamage(shot.dmg)#Deals damage to the ship hull points based on the damage value of the shot
                shot.kill()#finally, deletes it.
        
                
def cleargroup(group):
    """Deletes every object in the game."""
    for i in group:
        for j in i:
            j.kill()

def drawexplosions():
    """Draws all the explosions"""
    for explosion in explosions:
        DISPLAYSURF.blit(explosion.image, (explosion.rect.x, explosion.rect.y))

def spawn():
    """Controls where and when will the enemies appear"""
    time = pygame.time.get_ticks() + 30 #gets time after pygame.init was called, in ms
    global boss_spawned
    global level
    global bigboss #bigboss has to be global so that bigboss.hp and .x/.y may be checked along the code
    if (time - time_game_begun) % (1000 + (1000*dif)) < dt: #spawn random mobs, from either left or right
        if randint(-1,1) == 1:
            if level < 5:
                spawnrand = randint(1,level)
            else:
                spawnrand = randint(1,3)
            if spawnrand == 1:
                mob1 = ship(graphwidth+30,randint(15,graphheight/3),20,6,3,7/dif,'simple',20,mob1Img,straight_shotImg,"red")
            elif spawnrand == 2:
                mob2 = ship(graphwidth+30,randint(15,graphheight/3),20,5,3,7/dif,'wave',30,mob2Img,wave_shotImg,'red')
            elif spawnrand == 3:
                mob3 = ship(graphwidth+30,randint(15,graphheight/3),20,5,3,7/dif,'whip',25,mob3Img,whip_shotImg,'red')
            elif spawnrand == 5:
                mob4 = fastguy(graphwidth+30,randint(15,graphheight/3),20,10,8,7/dif,'aimshot',10,mob4Img,aimed_shotImg,'red')
        else:
            if level < 5: #may have more levels later
                spawnrand = randint(1,level)
            else:
                spawnrand = randint(1,3)
            if spawnrand == 1:
                mob1 = ship(-30,randint(15,graphheight/3),20,6,3,7/dif,'simple',20,mob1Img,straight_shotImg,"red")
            elif spawnrand == 2:
                mob2 = ship(-30,randint(15,graphheight/3),20,5,3,7/dif,'wave',30,mob2Img,wave_shotImg,'red')
            elif spawnrand == 3:
                mob3 = ship(-30,randint(15,graphheight/3),20,5,3,7/dif,'whip',25,mob3Img,whip_shotImg,'red')
            elif spawnrand == 4:
                mob4 = fastguy(-30,randint(15,graphheight/3),20,10,8,7/dif,'aimshot',10,mob4Img,aimed_shotImg,'red')
                
    if level >= 4 and boss_spawned == False and (leveltimechange - time) % 30000 <= dt: #if a certain time has passed, the boss has not appeared, and it's on the right level, make him appear
        boss_spawned = True
        if randint(0,1) == 1:
            bigboss = boss(graphwidth/2+90,-22,300,5,4,10/dif,'explosive',70,bossImg,boss_shotImg,"red")
        else:
            bigboss = boss(graphwidth/2-90,-22,300,5,4,10/dif,'explosive',70,bossImg,boss_shotImg,"red")
            
def levelmechanics():
    """Handles level transitions and wait time for it"""
    global level
    global levelwait
    global lastkills
    global leveltimechange
    global kills
    if gamemode == "arcade" and level < 4:
        if kills >= 15 or pygame.time.get_ticks() - leveltimechange >= 120000: #changes the level if the player has killed at least 15 mobs or if 2 minutes have passed
            leveltimechange = pygame.time.get_ticks() #resets the level timer
            lastkills += kills
            kills = 0 #saves the kills on laskills and resets the current kills
            level += 1
            levelwait = 3000 #sets the level changing timer to 3s


###########################=========############################
###########################  Menus  ############################
###########################=========############################
def inprint(inputs,scoreint): #centered print
    """ui for high-score name input"""
    global namestr
    global font
    DISPLAYSURF2.fill(BLACK)
    DISPLAYSURF2.blit(font.render(("Your score: " + str(scoreint)),True, WHITE),(graphwidth/2-(font.size(("Your score: " + str(scoreint)))[0])/2,(graphheight/2)-(font.size(("Your score: " + str(scoreint)))[1]/2)))
    DISPLAYSURF2.blit(font.render("Type your name: " + inputs,True, WHITE),(graphwidth/2-(font.size("Type your name: " + inputs)[0])/2,(graphheight/2)+(font.size("Type your name: " + inputs)[1])/2))
    pygame.display.update()

def prnt(dicti): #shows the highscore, centered
    """ui for showing high-scores"""
    DISPLAYSURF2.fill(BLACK) #clears the screen
    height = graphheight/1.6 
    for dicto in dicti: #gets the heights of all the scores, based on the graphheight
        for key in dicto:
            height -= (font.size(dicto[key])[1])/2
    height -= font.size("High scores:")[1] 
    DISPLAYSURF2.blit(font.render("High scores:",True, WHITE),(graphwidth/2-(font.size("High scores:")[0])/2,height-(font.size("High scores:")[1])/2))
    height += font.size("High scores:")[1]
    for dicto in dicti: #starts printing the scores, and change the height in which the next score will be printed
        for key in dicto:
            string = str(key) + ": " + str(dicto[key])
            DISPLAYSURF2.blit(font.render(string,True, WHITE),(graphwidth/2-(font.size(string)[0])/2,height-(font.size(string)[1])/2))
            height += (font.size(dicto[key])[1])
    pygame.display.update()

def endgame(scoreint):
    """loop that holds what happens after either you or the boss died"""
    name = []#stores all the letters the user types, one by one.
    namestr = ''#empty string, will store the final name as typed by the player.
    inprint(namestr, scoreint)
    if gamemode == "arcade":
        scorefile = "Scores.txt" #will call get and setscore based on the text file here defined
    elif gamemode == "survival":
        scorefile = "SurScores.txt"
    while True:
        for event in pygame.event.get(): # event handling loop
            if event.type == QUIT or (event.type == KEYUP and event.key == K_ESCAPE):
                pygame.quit()
                sys.exit()
            if event.type == KEYDOWN:
                if event.key != K_RETURN and event.key != K_BACKSPACE and event.key != K_MINUS and event.key != K_KP_MINUS:
                    name.append(pygame.key.name(event.key))
                if event.key == K_BACKSPACE and len(name)>0:
                    name.pop() #removes the last item on the list, if the player presses backspace
                for i in name:
                    namestr += i
                inprint(namestr,scoreint) #calls the print function, with the current string
                namestr = '' #clears the string, so it can be filled again, with the most recent input
                if event.key == K_RETURN and len(name)>=0:
                    for i in name: #calls setscore and shows the highscore
                        namestr += i #concatenates every letter in name to a single string
                    score.setScore(namestr, scoreint, scorefile) #adds the player name, and the score to the right file
                    prnt(score.getScore(scorefile))#shows the highscore
                    pygame.time.delay(3600)#waits 3.6 seconds
                    ost_snd.fadeout(1000)#finishes the song
                    return

def choiceprint(topstr,midstr,botstr): #funtion for general printing of 3 centered strings
    """Draws 3 strings, centered"""
    DISPLAYSURF2.blit(font.render(topstr, True, WHITE),(graphwidth/2 - font.size(topstr)[0]/2, graphheight/2 - font.size(topstr)[1]))
    DISPLAYSURF2.blit(font.render(midstr, True, WHITE),(graphwidth/2 - font.size(midstr)[0]/2, graphheight/2))
    DISPLAYSURF2.blit(font.render(botstr, True, WHITE),(graphwidth/2 - font.size(botstr)[0]/2, graphheight/2 + font.size(botstr)[1]))
    pygame.display.update()
    return

def choicepos(strg,pos):
    """Return dict with range in x(key 0) and range in y(key 1) of the given position, which can be top mid bot"""
    if pos == "top":
        posdict = {1:range(graphheight/2 - font.size(strg)[1],graphheight/2),\
                   0:range(graphwidth/2 - font.size(strg)[0]/2,graphwidth/2 + font.size(strg)[0]/2)}
    if pos == "mid":
        posdict = {1:range(graphheight/2, graphheight/2 + font.size(strg)[1]),\
                   0:range(graphwidth/2 - font.size(strg)[0]/2,graphwidth/2 + font.size(strg)[0]/2)}
    if pos == "bot":
        posdict = {1:range(graphheight/2 + font.size(strg)[1],graphheight/2 + font.size(strg)[1]*2),\
                   0:range(graphwidth/2 - font.size(strg)[0]/2, graphwidth/2 + font.size(strg)[0]/2)}
    return posdict
            
def choosedif():
    """ui for choosing difficulty, takes no arguments"""
    global dif
    global mousex
    global mousey
    pygame.mouse.set_visible(1)
    DISPLAYSURF2.fill(BLACK)
    mouseClicked = False
    while True:
        DISPLAYSURF2.blit(font.render("Back", True, WHITE),(font.size("Back")[0]/8, graphheight - font.size("Back")[1]))
        choiceprint("Easy","Medium","Hard") #calls the above mentioned function, so that it draws, from top to bottom, easy, medium and hard
        for event in pygame.event.get(): # event handling loop
            if event.type == QUIT or (event.type == KEYUP and event.key == K_ESCAPE):
                pygame.quit()
                sys.exit()
            elif event.type == MOUSEMOTION:
                mousex, mousey = event.pos
            elif event.type == MOUSEBUTTONDOWN:
                mousex, mousey = event.pos
                mouseClicked = True
            elif event.type == MOUSEBUTTONUP:
                mousex, mousey = event.pos
                mouseClicked = False
        if mousey in choicepos("Easy","top")[1] \
           and mousex in choicepos("Easy","top")[0]\
           and mouseClicked == True: #checks if there was a mouse click in the positions of the button
            dif = 3
            mouseClicked = False
            return
        if mousey in choicepos("Medium","mid")[1] \
           and mousex in choicepos("Medium","mid")[0] \
           and mouseClicked == True:
            dif = 2
            mouseClicked = False
            return
        if mousey in choicepos("Hard","bot")[1] \
           and mousex in choicepos("Hard","bot")[0] \
           and mouseClicked == True:
            dif = 1
            mouseClicked = False
            return
        if mousey in range(graphheight - font.size("Back")[1],graphheight)\
           and mousex in range(font.size("Back")[0]/8,(font.size("Back")[0]/8)+font.size("Back")[0])\
           and mouseClicked == True:
            dif = 0
            mouseClicked = False
            return
        
def choosescore():
    """ui for choosing which highscore to show, takes no arguments"""
    global mousex
    global mousey
    pygame.mouse.set_visible(1)
    DISPLAYSURF2.fill(BLACK)
    mouseClicked = False
    while True:
        DISPLAYSURF2.blit(font.render("Back", True, WHITE),(font.size("Back")[0]/8, graphheight - font.size("Back")[1]))
        choiceprint("Survival"," ","Arcade") #print survival and arcade, at the centered top and centered bottom positions
        for event in pygame.event.get(): # event handling loop
            if event.type == QUIT or (event.type == KEYUP and event.key == K_ESCAPE):
                pygame.quit()
                sys.exit()
            elif event.type == MOUSEMOTION:
                mousex, mousey = event.pos
            elif event.type == MOUSEBUTTONDOWN:
                mousex, mousey = event.pos
                mouseClicked = True
            elif event.type == MOUSEBUTTONUP:
                mousex, mousey = event.pos
                mouseClicked = False
        if mousey in choicepos("Survival","top")[1] \
           and mousex in choicepos("Survival","top")[0] \
           and mouseClicked == True: #checks if you clicked over the survival text
            prnt(score.getScore("SurScores.txt")) #calls the print score function
            pygame.time.delay(3600)
            mouseClicked = False
            return
        if mousey in choicepos("Arcade","bot")[1] \
           and mousex in choicepos("Arcade","bot")[0] \
           and mouseClicked == True:
            prnt(score.getScore("Scores.txt"))
            pygame.time.delay(3600)
            mouseClicked = False
            return
        if mousey in range(graphheight - font.size("Back")[1],graphheight)\
           and mousex in range(font.size("Back")[0]/8,(font.size("Back")[0]/8)+font.size("Back")[0])\
           and mouseClicked == True:
            mouseClicked = False
            return

def choosemode(): 
    """UI for choosing game mode, takes no arguments"""
    global gamemode
    global mousex
    global mousey
    global dif
    global level
    pygame.mouse.set_visible(1)
    DISPLAYSURF2.fill(BLACK)
    mouseClicked = False
    while True:
        DISPLAYSURF2.blit(font.render("Back", True, WHITE),(font.size("Back")[0]/8, graphheight - font.size("Back")[1]))
        choiceprint("Survival"," ","Arcade")
        for event in pygame.event.get(): # event handling loop
            if event.type == QUIT or (event.type == KEYUP and event.key == K_ESCAPE):
                pygame.quit()
                sys.exit()
            elif event.type == MOUSEMOTION:
                mousex, mousey = event.pos
            elif event.type == MOUSEBUTTONDOWN:
                mousex, mousey = event.pos
                mouseClicked = True
            elif event.type == MOUSEBUTTONUP:
                mousex, mousey = event.pos
                mouseClicked = False
        if mousey in choicepos("Survival","top")[1] \
           and mousex in choicepos("Survival","top")[0] \
           and mouseClicked == True:
            gamemode = 'survival'
            mouseClicked = False
            level = 4
            return play() #return to the last fucntion while calling play(), so that the  funtion still exits while moving to the next one
        if mousey in choicepos("Arcade","bot")[1] \
           and mousex in choicepos("Arcade","bot")[0] \
           and mouseClicked == True:
            gamemode = 'arcade'
            mouseClicked = False
            choosedif()
            if dif != 0: #checks if the function has not exited by the back button, checking the return given by the dif variable
                level = 1
                play()
            dif = 1
            return
        if mousey in range(graphheight - font.size("Back")[1],graphheight)\
           and mousex in range(font.size("Back")[0]/8,(font.size("Back")[0]/8)+font.size("Back")[0])\
           and mouseClicked == True:
            mouseClicked = False
            return

def scoreint():
    """Returns the player score based on the game mode"""
    if gamemode == "arcade":
        return ((player.hp*100000/(time_game_finished-time_game_begun)+((kills+lastkills)*100))/dif) #as kills are used for level changing, they are saved on laskills, and must be added to the score
    elif gamemode == "survival":
        return (time_game_finished-time_game_begun)/1000
    
def menu(): #main menu
    pygame.mouse.set_visible(1) #makes sure that the mouse is visible whilst the player is using the menu 
    global main_menu_music
    if not main_menu_music:
        main_menu_snd.play(-1)#Begins playing the main menu music. -1 as an argument means it loops indefinitely or until stopped.
        main_menu_music = True #Boolean used to check if main_menu is already playing, as to avoid multiple instances.
    global dif
    global mousex
    global mousey
    mouseClicked = False
    while True:
        DISPLAYSURF.fill(BLACK)
        DISPLAYSURF2.blit(font.render("PEW! PEW!",True, WHITE),(graphwidth/2-(font.size("PEW! PEW!")[0])/2,(graphheight/2)-(font.size("PEW! PEW!")[1]*3)))
        DISPLAYSURF2.blit(font.render("START",True, WHITE),(graphwidth/2-(font.size("START")[0])/2,(graphheight/2)+(font.size("START")[1])))
        DISPLAYSURF2.blit(font.render("HIGHSCORES",True, WHITE),(graphwidth/2-(font.size("HIGHSCORES")[0])/2,(graphheight/2)+(font.size("HIGHSCORES")[1]*2)))
        pygame.display.update()
        for event in pygame.event.get(): # event handling loop
            if event.type == QUIT or (event.type == KEYUP and event.key == K_ESCAPE):
                pygame.quit()
                sys.exit()
            elif event.type == MOUSEMOTION:
                mousex, mousey = event.pos
            elif event.type == MOUSEBUTTONDOWN:
                mousex, mousey = event.pos
                mouseClicked = True
            elif event.type == MOUSEBUTTONUP:
                mousex, mousey = event.pos
                mouseClicked = False
        if mousey in range(((graphheight/2)+(font.size("START")[1])),(graphheight/2)+(font.size("START")[1]*2)) \
           and mousex in range(graphwidth/2-(font.size("START")[0])/2,graphwidth/2+(font.size("START")[0])/2) \
           and mouseClicked == True: #check to see if the player has clicked the text
            mouseClicked = False
            choosemode()
            return 
        if mousey in range(((graphheight/2)+(font.size("HIGHSCORE")[1]*2)),(graphheight/2)+(font.size("HIGHSCORE")[1]*3)) \
           and mousex in range(graphwidth/2-(font.size("HIGHSCORE")[0])/2,graphwidth/2+(font.size("HIGHSCORE")[0])/2) \
           and mouseClicked == True: #check to see if the player has clicked the text
            mouseClicked = False
            choosescore()
            return

def play(): #main game loop
    global mousex
    global mousey
    global mouseClicked
    global player

    global game_finish
    global time_game_finished
    global time_game_begun
    global main_menu_music

    global kills
    global boss_spawned
    global prevtime
    global dt
    global nowtime
    global dtmod

    global levelwait
    global lastkills
    global leveltimechange
    global fpsmod
    #clearing or starting the variables used while playing
    main_menu_snd.fadeout(1000)#Fadeouts the main-menu music during 1 second
    main_menu_music = False
    pygame.mouse.set_visible(0)
    ost_snd.play(-1)
    nowtime = pygame.time.get_ticks()
    prevtime = pygame.time.get_ticks()
    dtmod = 16 #16 is 1000/60, hence, the amount of miliseconds between one frame and another in 60fps
    player = ship(mousex,mousey,100,0,-10,10,'simple',15,playerImg,playershotImg,"green")
    fpsmod = 1
    fpsClock.tick(FPS)
    player_shot_delay = 0 #Will be used at the game-loop in order to control the player rof
    player_rof = 15 #rate of fire, how many frames has to pass until the next shot is fired
    time_game_begun = pygame.time.get_ticks()
    time_game_finished = 0
    game_finish = 0
    boss_spawned = False
    kills = 0
    gamerun = False
    if gamemode == 'arcade':
        level = 1
    else:
        level = 4
    lastkills = 0
    levelwait = 3000 #timer so that the level is displayed when the game starts
    leveltimechange = 0
    cleargroup(everything) #deletes old shots and ships
    while True:
        nowtime = pygame.time.get_ticks()
        dt = nowtime-prevtime #delta time
        prevtime = pygame.time.get_ticks()
        #print str(dt) + " dt " + str(prevtime) + " prevtime " + str(nowtime) + " nowtime"
        if time_game_finished != 0:  
            if pygame.time.get_ticks() - time_game_finished > 2000:
                endgame(scoreint())
                ost_snd.fadeout(1000)
                return
        if player.hp <= 0 and game_finish == 0:
            game_finish = 1
            time_game_finished = pygame.time.get_ticks()
        elif boss_spawned == True and bigboss.hp <= 0 and game_finish == 0 and gamemode == 'arcade':
            game_finish = 1
            time_game_finished = pygame.time.get_ticks()
        elif boss_spawned == True and bigboss.hp <= 0 and gamemode == 'survival':
            boss_spawned = False
        for event in pygame.event.get(): # event handling loop
            if event.type == QUIT or (event.type == KEYUP and event.key == K_ESCAPE):
                pygame.quit()#Pressing the up and escape keys at the same time will quit the game.
                sys.exit()
            elif event.type == MOUSEMOTION:#Handles the mouse movement
                mousex, mousey = event.pos
            elif event.type == MOUSEBUTTONUP:
                mousex, mousey = event.pos #When the mouse buttom is released, it causes mouseClicked to be False.
                mouseClicked = False
            elif event.type == MOUSEBUTTONDOWN: #While the mouse buttom is pressed, mouse(is)Clicked.
                mousex, mousey = event.pos
                mouseClicked = True
        if mouseClicked == True and player_shot_delay <= 0 and player.hp > 0: #Checks whether the player is able to shoot
            player.shoot()
            player_shot_delay = player_rof#Resets the delay to the original rate of fire.
        player_shot_delay -= dt/dtmod#Decrements every loop so that if the player stops firing any second, he can always start firing.
        draw()
        levelmechanics()#Controls the level mechanics
        if not boss_spawned and levelwait <= 0: #stops spawn if the boss has spawned or if the game is displaying the level
            spawn()#passes the level to the spawn function
        for shipt in enemy_list:
            if shipt.hp < 1:
                kills += 1
                shipt.explode() #calls the explosion method if a ship has died
                continue
            shipt.ai()
        for shipt in player_list:
            if shipt.hp <= 0:
                shipt.explode()
            else:
                shipt.rect.x = mousex-shipt.rect.width/2
                shipt.rect.y = mousey-shipt.rect.height/2 #Centers the mouse at player's ship
        for explosion in explosions:
            explosion.cycle() #cicles the current explosions
        fpsClock.tick(FPS) #Ticks a single frame, serves as a time-controller
        if gamerun == True:
            fpsmod = (fpsClock.get_fps()/FPS)
        else:
            fpsmod = 1
        gemerun = True
while True: #so that even if the game ends,the player is taken back to the menu after.
    menu()
