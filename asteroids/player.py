import pygame
from circleshape import CircleShape
from constants import PLAYER_SHOOT_COOLDOWN_SECONDS, PLAYER_SHOT_SPEED, PLAYER_RADIUS, LINE_WIDTH, PLAYER_TURN_SPEED, PLAYER_SPEED
from shot import Shot

class Player(CircleShape):
    def __init__(self, x: float, y: float):
        super().__init__(x=x, y=y, radius=PLAYER_RADIUS)
        self.rotation: float = 0.0
        self.shoot_cooldown_timer: float = 0.0
    
    def triangle(self) -> list[pygame.Vector2]:
        forward: pygame.Vector2 = pygame.Vector2(0,1).rotate(self.rotation)
        right: pygame.Vector2 = pygame.Vector2(0,1).rotate(self.rotation + 90) * self.radius / 1.5
        a: pygame.Vector2 = self.position + forward * self.radius
        b: pygame.Vector2 = self.position - forward * self.radius - right
        c: pygame.Vector2 = self.position - forward * self.radius + right
        return [a,b,c]
    
    def rotate(self, dt: float) -> None:
        self.rotation += PLAYER_TURN_SPEED * dt
    
    def move(self, dt: float) -> None:
        self.position += pygame.Vector2(0,1).rotate(self.rotation) * PLAYER_SPEED * dt
    
    def shoot(self) -> None:
        if self.shoot_cooldown_timer <= 0:
            shot = Shot(x=self.position[0], y=self.position[1])
            shot.velocity = pygame.Vector2(0,1).rotate(self.rotation) * PLAYER_SHOT_SPEED
            self.shoot_cooldown_timer = PLAYER_SHOOT_COOLDOWN_SECONDS


    def update(self, dt: float) -> None:
        self.shoot_cooldown_timer -= dt
        keys = pygame.key.get_pressed()
        if keys[pygame.K_a]:
            self.rotate(dt=-dt)
        if keys[pygame.K_d]:
            self.rotate(dt=dt)
        if keys[pygame.K_w]:
            self.move(dt=dt)
        if keys[pygame.K_s]:
            self.move(dt=-dt)
        if keys[pygame.K_SPACE]:
            self.shoot()

    def draw(self, screen: pygame.Surface) -> None:
        pygame.draw.polygon(surface=screen, color="white", points=self.triangle(), width=LINE_WIDTH)