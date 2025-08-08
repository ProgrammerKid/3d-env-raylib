package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func get_hitbox(pos rl.Vector3, model rl.Model) rl.BoundingBox {
	bounding_box := rl.GetMeshBoundingBox(model.GetMeshes()[0])
	return rl.NewBoundingBox(
		rl.Vector3Add(pos, bounding_box.Min),
		rl.Vector3Add(pos, bounding_box.Max),
	)
}

func get_obstacles(tank Tank) []rl.BoundingBox {
	obstacles := make([]rl.BoundingBox, 0, 5)

	for _, hitbox := range tank.hitboxes {
		obstacles = append(obstacles, hitbox)
	}

	return obstacles
}

// Tank
type Tank struct {
	dim        rl.Vector3
	pos        rl.Vector3
	texture    rl.Texture2D
	models_dim [5]rl.Vector3
	models_pos [5]rl.Vector3
	models     [5]rl.Model
	hitboxes   [5]rl.BoundingBox
}

func (tank *Tank) init(dim rl.Vector3, pos rl.Vector3) {
	tank.dim = dim
	tank.pos = pos

	tank.texture = rl.LoadTextureFromImage(rl.GenImageGradientRadial(int(dim.X)/3, int(dim.Z)/3, 1.5, rl.DarkBlue, rl.SkyBlue))
	tank.models_dim = [5]rl.Vector3{
		rl.NewVector3(dim.X, 2, dim.Z),
		rl.NewVector3(2, dim.Y, dim.Z),
		rl.NewVector3(2, dim.Y, dim.Z),
		rl.NewVector3(dim.X, dim.Y, 2),
		rl.NewVector3(dim.X, dim.Y, 2),
	}

	for i, model_dim := range tank.models_dim {
		tank.models[i] = rl.LoadModelFromMesh(rl.GenMeshCube(model_dim.X, model_dim.Y, model_dim.Z))
		rl.SetMaterialTexture(&tank.models[i].GetMaterials()[0], rl.MapAlbedo, tank.texture)
	}

	tank.models_pos = [5]rl.Vector3{
		rl.NewVector3(tank.pos.X, tank.pos.Y-1, tank.pos.Z),
		rl.NewVector3(-tank.dim.X/2-1+tank.pos.X, tank.dim.Y/2+tank.pos.Y, tank.pos.Z),
		rl.NewVector3(tank.dim.X/2+1+tank.pos.X, tank.dim.Y/2+tank.pos.Y, tank.pos.Z),
		rl.NewVector3(tank.pos.X, tank.dim.Y/2+tank.pos.Y, -tank.dim.Z/2-1+tank.pos.Z),
		rl.NewVector3(tank.pos.X, tank.dim.Y/2+tank.pos.Y, tank.dim.Z/2+1+tank.pos.Z),
	}

	for i, pos := range tank.models_pos {
		tank.hitboxes[i] = get_hitbox(rl.NewVector3(pos.X, pos.Y, pos.Z), tank.models[i])
	}

}

func (tank *Tank) draw() {
	for i, model := range tank.models {
		if i != 4 {
			rl.DrawModel(model,
				rl.NewVector3(tank.models_pos[i].X, tank.models_pos[i].Y, tank.models_pos[i].Z),
				1, rl.White)
		}
	}
}

// Player
type Player struct {
	friction     float32
	acceleration float32
	speed        rl.Vector3
	pos          rl.Vector3
	length       float32
	model        rl.Model
	hitbox       rl.BoundingBox
	// rotation     float32
}

func (p *Player) init(friction float32, acceleration float32, pos rl.Vector3, length float32) {
	p.friction = friction
	p.acceleration = acceleration
	p.pos = pos
	p.length = length
	p.speed = rl.NewVector3(0, 0, 0)

	texture := rl.LoadTextureFromImage(rl.GenImageGradientRadial(int(p.length)*3, int(p.length)*3, 1.4, rl.Purple, rl.DarkPurple))
	p.model = rl.LoadModelFromMesh(rl.GenMeshCube(p.length, p.length, p.length))
	rl.SetMaterialTexture(&p.model.GetMaterials()[0], rl.MapAlbedo, texture)

}

func (p *Player) update(obstacles []rl.BoundingBox) {
	p.hitbox = get_hitbox(p.pos, p.model)

	// Move
	// if !p.check_collisions(nil, nil, false) {
	var dir rl.Vector3
	if rl.IsKeyDown(rl.KeyD) {
		dir.X = 1
	} else if rl.IsKeyDown(rl.KeyA) {
		dir.X = -1
	} else {
		dir.X = 0
	}
	if rl.IsKeyDown(rl.KeyW) {
		dir.Z = -1
	} else if rl.IsKeyDown(rl.KeyS) {
		dir.Z = 1
	} else {
		dir.Z = 0
	}

	dir = rl.Vector3Normalize(dir)

	p.speed.X += dir.X * p.acceleration
	p.speed.Z += dir.Z * p.acceleration

	p.speed.X *= p.friction
	if math.Abs(float64(p.speed.X)) < 0.1 {
		p.speed.X = 0
	}
	p.speed.Z *= p.friction
	if math.Abs(float64(p.speed.Z)) < 0.1 {
		p.speed.Z = 0
	}
	// }

	p.pos.X += p.speed.X
	p.hitbox = get_hitbox(p.pos, p.model)
	p.check_collisions("x", obstacles, true)
	p.pos.Z += p.speed.Z
	p.hitbox = get_hitbox(p.pos, p.model)
	p.check_collisions("z", obstacles, true)

}

func (p *Player) check_collisions(axis any, obstacles []rl.BoundingBox, update bool) bool {

	for _, obstacle := range obstacles {
		if rl.CheckCollisionBoxes(obstacle, p.hitbox) {

			fmt.Println("collision")

			if update {
				if axis == "x" {
					if p.speed.X > 0 {
						p.pos.X = obstacle.Min.X - p.length/2 - 0.0001
						p.speed.X = 0
					} else if p.speed.X < 0 {
						p.pos.X = obstacle.Max.X + p.length/2 + 0.0001
						p.speed.X = 0
					}
				}
				if axis == "z" {
					if p.speed.Z > 0 {
						p.pos.Z = obstacle.Min.Z - p.length/2 - 0.0001
						p.speed.Z = 0
					} else if p.speed.Z < 0 {
						p.pos.Z = obstacle.Max.Z + p.length/2 + 0.0001
						p.speed.Z = 0
					}
				}
				if axis == "y" {

				}
			}
			return true

		} else {
			fmt.Println("no collision")
		}
	}
	return false
}

func (p *Player) draw() {
	rl.DrawModel(p.model, p.pos, 1, rl.White)
}

// Main
func main() {
	fmt.Println("Started")
	rl.InitWindow(1500, 800, "Game")
	rl.SetTargetFPS(60)

	camera := rl.NewCamera3D(
		rl.NewVector3(0, 80, 100), //pos
		rl.NewVector3(0, 0, 0),    //target
		rl.NewVector3(0, 1, 0),
		80, rl.CameraPerspective)

	var tank Tank
	tank.init(rl.NewVector3(150, 30, 150), rl.NewVector3(0, 0, 0))

	var player Player
	player.init(0.9, 0.2, rl.NewVector3(0, 6, 0), 10)

	var obstacles []rl.BoundingBox = get_obstacles(tank)

	for !rl.WindowShouldClose() {

		// dt := rl.GetFrameTime()
		player.update(obstacles)
		// player.update(tank.hitboxes)

		rl.ClearBackground(rl.White)
		rl.BeginDrawing()
		rl.DrawFPS(500, 50)

		rl.BeginMode3D(camera)
		// rl.DrawGrid(20, 5)

		tank.draw()
		player.draw()

		rl.EndMode3D()
		rl.EndDrawing()

	}
	rl.CloseWindow()
}
