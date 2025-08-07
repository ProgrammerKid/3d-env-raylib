package main

// import rl "github.com/gen2brain/raylib-go/raylib"

// type Player struct {
// 	friction     float32
// 	acceleration float32
// 	speed        rl.Vector2
// 	pos          rl.Vector2
// 	rotation     float32
// }

// func (p *Player) update() {
// 	// Movement input
// 	var dir rl.Vector2
// 	if rl.IsKeyDown(rl.KeyD) {
// 		dir.X += 1
// 	}
// 	if rl.IsKeyDown(rl.KeyA) {
// 		dir.X -= 1
// 	}
// 	if rl.IsKeyDown(rl.KeyW) {
// 		dir.Y -= 1
// 	}
// 	if rl.IsKeyDown(rl.KeyS) {
// 		dir.Y += 1
// 	}

// 	dir = rl.Vector2Normalize(dir)

// 	p.speed.X += dir.X * p.acceleration
// 	p.speed.Y += dir.Y * p.acceleration

// 	p.speed.X *= p.friction
// 	p.speed.Y *= p.friction

// 	p.pos.X += p.speed.X * 0.02
// 	p.pos.Y += p.speed.Y * 0.02

// 	texture := textures[0]
// 	halfWidth := float32(texture.Width) / 2
// 	halfHeight := float32(texture.Height) / 2

// 	screenW := float32(rl.GetScreenWidth())
// 	screenH := float32(rl.GetScreenHeight())

// 	if p.pos.X < halfWidth {
// 		p.pos.X = halfWidth
// 	}
// 	if p.pos.X > screenW-halfWidth {
// 		p.pos.X = screenW - halfWidth
// 	}
// 	if p.pos.Y < halfHeight {
// 		p.pos.Y = halfHeight
// 	}
// 	if p.pos.Y > screenH-halfHeight {
// 		p.pos.Y = screenH - halfHeight
// 	}

// 	// Rotation to face mouse
// 	dir = rl.Vector2Subtract(rl.GetMousePosition(), p.pos)
// 	p.rotation = float32(rl.Vector2Angle(rl.Vector2{X: 1, Y: 0}, dir)) * rl.Rad2deg
// }

// func (p *Player) draw() {
// 	texture := textures[0]
// 	center := rl.Vector2{X: float32(texture.Width) / 2, Y: float32(texture.Height) / 2}
// 	rl.DrawTexturePro(
// 		texture,
// 		rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)),
// 		rl.NewRectangle(p.pos.X, p.pos.Y, float32(texture.Width), float32(texture.Height)),
// 		center,
// 		p.rotation,
// 		rl.White,
// 	)
// }