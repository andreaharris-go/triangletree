package triangletree

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func TreeSum(input interface{}) (int, error) {
	var triangle [][]int

	switch v := input.(type) {
	case string:
		data, err := fetchFromURL(v)

		if err != nil {
			return 0, fmt.Errorf("error fetching data from URL: %w", err)
		}

		if !isValidTriangle(data) {
			return 0, errors.New("invalid triangle structure from URL")
		}

		triangle = data

	case [][]int:
		if !isValidTriangle(v) {
			return 0, errors.New("invalid triangle structure from [][]int")
		}

		triangle = v

	default:
		return 0, errors.New("unsupported input type (must be string URL or [][]int)")
	}

	return maxPathSum(triangle), nil
}

func fetchFromURL(url string) ([][]int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch from URL: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %w", err)
	}

	var triangle [][]int
	if err := json.Unmarshal(body, &triangle); err != nil {
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	return triangle, nil
}

func isValidTriangle(triangle [][]int) bool {
	for i, row := range triangle {
		if len(row) != i+1 {
			return false
		}

		for _, val := range row {
			if reflect.TypeOf(val).Kind() != reflect.Int {
				return false
			}
		}
	}

	return true
}

func maxPathSum(triangle [][]int) int {
	dp := make([][]int, len(triangle))

	for i := range triangle {
		dp[i] = make([]int, len(triangle[i]))
		copy(dp[i], triangle[i])
	}

	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j < len(dp[i]); j++ {
			left := dp[i+1][j]
			right := dp[i+1][j+1]
			if left > right {
				dp[i][j] += left
			} else {
				dp[i][j] += right
			}
		}
	}

	return dp[0][0]
}
