/*
 * @lc app=leetcode.cn id=812 lang=golang
 *
 * [812] 最大三角形面积
 */

// @lc code=start
func largestTriangleArea(points [][]int) float64 {
// 根据三顶点坐标求三角形面积 鞋带公式: 1/2[(x1y2-x2y1)+(x2y3-x3y2)+(x3y1-x1y3)]
	var res float64
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			for k := j+1; k < n; k++ {
				res = math.Max(res, 0.5*math.Abs(float64(
				(points[i][0]*points[j][1]-points[j][0]*points[i][1])+
				(points[j][0]*points[k][1]-points[k][0]*points[j][1])+
				(points[k][0]*points[i][1]-points[i][0]*points[k][1]))))
			}
		}
	}
	return res
}
// @lc code=end

