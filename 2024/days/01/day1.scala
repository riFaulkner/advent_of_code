//> using scala 3.5.2
//> using toolkit 0.6.0
import scala.util.matching.Regex
import scala.compiletime.ops.int
import scala.util.Try
import scala.collection.mutable.ArrayBuffer
import scala.collection.mutable.Map
import scala.compiletime.ops.boolean

@main
def day01(fileName: String, part2: Boolean): Unit =
  // First parse the input file into lines
  println(s"Starting day 1 with file: $fileName")
  val path: os.Path = os.pwd / fileName
  val lines: Seq[String] = os.read.lines(path)
  // This can also be done as a stream, may be helpful to know later
  // val lineStream: geny.Generator[String] = os.read.lines.stream(path)
  // val firstLine: String = lineStream.head
  val left: ArrayBuffer[Int] = ArrayBuffer[Int]()
  val right: ArrayBuffer[Int] = ArrayBuffer[Int]()
  val rightMap: Map[Int, Int] = Map()

  for l <- lines do
    // for each line add the first item to the left and second item to the right
    val (x, y) = processLine(l)
    if x.isEmpty || y.isEmpty then
      println(s"Bad Line: $l produced x: $x -- y: $y")

    val xVal = x.get
    val yVal = y.get
    left.addOne(xVal)
    if !part2 then right.addOne(yVal)
    else
      val yOccurances = rightMap.getOrElse(yVal, 0)
      rightMap(yVal) = yOccurances + 1

  // sort the two lists
  if !part2 then
    left.sortInPlace()
    right.sortInPlace()

  // print(left)
  // println(right)

  // iterate through the lists, summing the difference
  var sum = 0
  left.indices.foreach(i =>
    if !part2 then
      val v = right(i) - left(i)
      sum = sum + math.abs(v)
    else
      // Get the occurances to create the similarity score
      val key = left(i)
      val v = rightMap.getOrElse(key, 0)
      sum = sum + (key * v)
  )
  println(sum)

def processLine(line: String): (Option[Int], Option[Int]) =
  val pattern: Regex = "\\d+".r
  val splits = pattern.findAllMatchIn(line).toList

  return (tryToInt(splits(0).toString()), tryToInt(splits(1).toString()))

def tryToInt(s: String) = Try(s.toInt).toOption
