package tasks

func numTilePossibilities(tiles string) int {
	//factorial based dynamic counting
	arr := make([]int, 26)
	for i := 0; i < len(tiles); i++ {
		arr[tiles[i]-'A']++
	}
	return 0
}
