// Is n divisible by x and y?
// https://www.codewars.com/kata/5545f109004975ea66000086

package kata

func IsDivisible(n, x, y int) bool {
    if n % x != 0 {
      return false;
    }

    if n % y != 0 {
      return false;
    }

    return true;
}
