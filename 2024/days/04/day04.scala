//> using scala 3.5.2
//> using toolkit 0.6.0

@main
def day04(fileName: String, pt2: Boolean): Unit =
  println("Starting Day 4")

  val path: os.Path = os.pwd / fileName
  val lines: Seq[String] = os.read.lines(path)

  var matches = 0
  lines.indices.foreach(y =>
    lines(y).indices.foreach(x =>
      if pt2 then
        if lines(y)(x) == 'A' then matches = matches + searchForMas(x, y, lines)
      else
        // check for terminating or triggering characters 'X' or 'S'
        // when a match is found, check in the eligible directions
        // right, down, diagonal to the right, diagonal to the left
        // to avoid dupes, never check up or to the left.
        if lines(y)(x) == 'X' then
          matches = matches + searchForXmas(x, y, "XMAS", lines)

        if lines(y)(x) == 'S' then
          matches = matches + searchForXmas(x, y, "SAMX", lines)
    )
  )
  println(s"Matches: $matches")

def searchForMas(x: Int, y: Int, lines: Seq[String]): Int =
  // println(s"Evaluating A at x: $x -- y: $y")
  var matches = 0
  val mas = "MAS"
  val sam = "SAM"

  // start up and to the left
  matches = matches + findMatch(x - 1, y - 1, 1, 1, mas, lines)
  matches = matches + findMatch(x - 1, y - 1, 1, 1, sam, lines)
  // start down and to the right
  matches = matches + findMatch(x - 1, y + 1, 1, -1, mas, lines)
  matches = matches + findMatch(x - 1, y + 1, 1, -1, sam, lines)

  // println(s"A matches $matches")

  if matches >= 2 then return 1

  return 0

def searchForXmas(x: Int, y: Int, s: String, lines: Seq[String]): Int =
  // check right
  var matches = findMatch(x, y, 1, 0, s, lines) // 5
  // var matches = 0

  // check down
  matches = matches + findMatch(x, y, 0, 1, s, lines)

  // check diagonal down right
  matches = matches + findMatch(x, y, 1, 1, s, lines)

  // check diagonal down left
  matches = matches + findMatch(x, y, 1, -1, s, lines) // 3
  return matches

def findMatch(
    x: Int,
    y: Int,
    xStep: Int,
    yStep: Int,
    s: String,
    lines: Seq[String]
): Int =
  var tmpX = x
  var tmpY = y

  var matches = 1
  s.foreach(c =>
    if matches == 0 then return 0
    else if tmpY < 0 || lines.length - 1 < tmpY || tmpX < 0 || lines(tmpY)
        .length() - 1 < tmpX
    then
      matches = 0
      return 0
    else if lines(tmpY)(tmpX) != c then
      // println(s"Match no longer eligible expected $")
      matches = 0
    else
      tmpX = tmpX + xStep
      tmpY = tmpY + yStep
  )
  return matches
