package main

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	foundObstacle := false
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			foundObstacle = true
		}
		if foundObstacle {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = 1
		}
	}
	foundObstacle = obstacleGrid[0][0] == 0
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			foundObstacle = true
		}
		if foundObstacle {
			obstacleGrid[0][i] = 0
		} else {
			obstacleGrid[0][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
