package main

import (
	"fmt"
)

type block struct {
	x int
	y int
	z int
}

var block_type_1 = []block{
	// {x_length, y_length, z_length}
	// 2x2x3
	{2, 2, 3},
	{3, 2, 2},
	{2, 3, 2},
}

var block_type_2 = []block{
	// 2x4x1
	{1, 2, 4},
	{2, 1, 4},
	{4, 2, 1},
	{2, 4, 1},
	{4, 1, 2},
	{1, 4, 2},
}

var block_type_3 = []block{
	// 1x1x1
	{1, 1, 1},
}

var blocks = [3]int{6, 6, 5}

var box [5][5][5]bool
var levels = [5]int{25, 25, 25, 25, 25}
var current_z = 0

func get_block(block_type int) []block {
	if blocks[block_type] > 0 {
		blocks[block_type]--
		switch block_type {
		case 0:
			return block_type_1
		case 1:
			return block_type_2
		case 2:
			return block_type_3
		}
	}
	return nil
}

func return_block(block_type int) {
	blocks[block_type]++
}

func no_more_blocks() bool {
	if (blocks[0] == 0) && (blocks[1] == 0) && (blocks[2] == 0) {
		return true
	}
	return false
}

func insert_block(x int, y int, z int, b block) bool {

	if (x+b.x > 5) ||
		(y+b.y > 5) ||
		(z+b.z > 5) ||
		(box[x][y][z] == true) ||
		(box[x][y][z+b.z-1] == true) {
		return false
	}
	// make sure the base is empty
	for y_idx := y; y_idx < y+b.y; y_idx++ {
		for x_idx := x; x_idx < x+b.x; x_idx++ {
			if box[x_idx][y_idx][z] == true {
				return false
			}
		}
	}

	for z_idx := z; z_idx < z+b.z; z_idx++ {
		for y_idx := y; y_idx < y+b.y; y_idx++ {
			for x_idx := x; x_idx < x+b.x; x_idx++ {
				box[x_idx][y_idx][z_idx] = true
			}
		}
		levels[z_idx] = levels[z_idx] - (b.x * b.y)
	}
	//fmt.Printf("removing -%d: %d\n", (b.x * b.y), levels[0])
	return true
}

func remove_block(x int, y int, z int, b block) {
	if (x+b.x > 5) ||
		(y+b.y > 5) ||
		(z+b.z > 5) ||
		(box[x][y][z] == false) {
		return
	}

	for z_idx := z; z_idx < z+b.z; z_idx++ {
		for y_idx := y; y_idx < y+b.y; y_idx++ {
			for x_idx := x; x_idx < x+b.x; x_idx++ {
				box[x_idx][y_idx][z_idx] = false
			}
		}
		levels[z_idx] = levels[z_idx] + (b.x * b.y)
	}
	//fmt.Printf("adding %d: %d\n", (b.x * b.y), levels[0])
}

func print_box(level int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if box[x][y][level] {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println(levels[level])
}

func next_block(x int, y int, z int) bool {

	open_slot := false
	for !open_slot {
		if x >= 5 {
			x = 0
			y++
		}
		if y >= 5 {
			// current z-level must be filled completely before moving to the next z-level
			if levels[z] == 0 {
				y = 0
				z++
			} else {
				return false
			}
		}

		// check if everything is complete
		if (z == 5) && (no_more_blocks()) {
			// solution found!
			fmt.Println("======= Solution found!!! =========")
			return true
		}

		// move on to the next slot if it's filled
		if box[x][y][z] == true {
			x++
		} else {
			open_slot = true
		}
	}

	fmt.Printf("inserting to x:%d y:%d z:%d \n", x, y, z)

	for block_type := 0; block_type < 3; block_type++ {
		b := get_block(block_type)
		//fmt.Print(b)
		if b != nil {
			count := cap(b)
			for direction := 0; direction < count; direction++ {
				//fmt.Printf("empty slots is %d \n", levels[z])
				if insert_block(x, y, z, b[direction]) {
					// print_box(0)
					// fmt.Println(block_type)
					// fmt.Printf("calling next with x:%d y:%d z:%d \n", x+b[direction].x, y, z)
					// time.Sleep(time.Second)
					if next_block(x+b[direction].x, y, z) {
						fmt.Println(b[direction])
						fmt.Printf("(x:%d y:%d z:%d) block:%d dir:%d \n", x, y, z, block_type, direction)
						return true
					}
					remove_block(x, y, z, b[direction])
					// fmt.Printf("removed %d\n", block_type)
					// print_box(0)
				}
			}
			return_block(block_type)
		}
	}

	// return false - no more options
	return false

}

func main() {

	next_block(0, 0, 0)
}
