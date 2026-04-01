package darts

func Score(x, y float64) int {
    distanceSquared := x*x + y*y
    switch {
    case distanceSquared <= 1*1:
        return 10
    case distanceSquared <= 5*5:
        return 5
    case distanceSquared <= 10*10:
        return 1
    default:
        return 0
    }
}
