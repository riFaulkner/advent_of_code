//> using scala 3.5.2
//> using toolkit 0.6.0
import scala.util.matching.Regex
import scala.collection.mutable.Map
import scala.annotation.switch
import scala.compiletime.ops.double
import scala.collection.mutable.ArrayBuffer

@main
def getAnswer(fileName: String, pt2: Boolean): Unit =
  println("Starting Day 4")

  val path: os.Path = os.pwd / fileName
  val lines: Seq[String] = os.read.lines(path)
  val coorMap: Map[Int, ArrayBuffer[Int]] = Map()
  val digitPatter = "\\d+".r
  var sum = 0

  // first read in the coordinates
  //  this is a coorMap
  // then read in the lines and verify them
  for l <- lines do
    // we may not need any indicaters for where we are in the ordering.println
    // If the digit parser has 2 items, its section 1. If it has 0 skip and if it has >2 its in the eval
    val splits =
      digitPatter.findAllMatchIn(l).map(s => s.toString().toInt).toList
    // println(s"parsing line $l -- has ${splits.length} splits")
    splits.length match
      case 2 =>
        val coor = splits(0)
        val arrBuf = coorMap.getOrElse(coor, ArrayBuffer())
        arrBuf.addOne(splits(1))
        coorMap(coor) = arrBuf

      case 0 => println("Length of 0 -- NOOP")
      case v =>
        val isVerified = verifyLine(splits, coorMap)
        if !pt2 && isVerified then
          val num = splits(
            splits.length / 2
          ) // get the middle item in the array
          // println(s"Line $l verified -- adding $num to the sum")
          sum = sum + num
        if pt2 && !isVerified then
          // do part 2. Figure out how to order the string, just return the number
          val num = orderLineAndReturnMiddle(splits, coorMap)
          sum = sum + num

  println(s"Sum: $sum")

def orderLineAndReturnMiddle(
    list: List[Int],
    coorMap: Map[Int, ArrayBuffer[Int]]
): Int =
  println(s"Invalid line provided: $list")
  val ordered = ArrayBuffer.fill(list.length)(0)

  return 0

// Brute force it and get it over with
def generateOrder(
    orig: List[Int],
    generated: ArrayBuffer[Int],
    coorMap: Map[Int, ArrayBuffer[Int]]
): Option[Int] =
  if orig.length == 0
  then // We're at the end of this generation, verify the generated list
    if verifyLine(generated.toList, coorMap)
    then // if it works, then this is the correct combo, return the middle value
      return Option.apply(generated.length / 2)
    return Option.empty
  breakable {
    for i <- orig do
      generated.append(i)
      val remaining = orig.filter(p => p != i).toList

  }

def verifyLine(list: List[Int], coorMap: Map[Int, ArrayBuffer[Int]]): Boolean =
  // println(s"list: $list -- coorMap: $coorMap")
  var isVerified = true
  list.indices.foreach(i =>
    // check the coordinate map, if there is a value (which there always should right?)
    // then verify that the response value has a later index in the list than the current i
    val mapping = coorMap.get(list(i))
    if !mapping.isEmpty then
      // Since each number can map to multiple items..
      mapping.get.foreach(num =>
        val indexOf = list.indexOf(num)
        if indexOf >= 0 && indexOf < i then
          isVerified = false
          return false
      )
  )
  return isVerified
