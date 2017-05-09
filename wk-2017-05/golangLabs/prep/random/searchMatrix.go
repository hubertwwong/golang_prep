func searchMatrix(matrix [][]int, target int) bool {
    lenX := len(matrix)
    lenY := len(matrix[0])
    
    x := 0
    y := 0
    for {
        if matrix[x][y] == target {
            return true
        } else if matrix[x][y] > target {
            return false
        } else if x + 1 == lenX && y + 1 == lenY {
            // you are at the bottom right
            // you are done.
            return false
        }
        
        // travel
        if x + 1 == lenX {
            y++
            continue
        } else if y + 1 == lenY {
            x++
            continue
        } else {
            posX := matrix[x+1][y]
            posY := matrix[x][y+1]
            
            if posX > target {
                y++
            } else if posY > target {
                x++
            } else if posX > posY {
                x++
            } else {
                y++
            }
        }
    }
}